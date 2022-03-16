# Extending Tendermint-BFT with VRF-based Election

## Consensus Overview

Ostracon's block generation mechanism based on Tendermint-BFT consists of the following three phases. We here refer to the generations of blocks as *height*, and a single approval round consisting of the following three processes as *round*.

**Election**. Elect one Proposer and several Voters from a candidate node set. This is the same as a Leader Election in a general distributed system, but in blockchain, it must be designed to prevent artificial selection so that malicious interference doesn't degrade the overall performance of the system. Also note that there is no centralized authority involved in Ostracon elections to ensure fairness. Since the election results can be computed deterministically by all nodes, each node can autonomously determine whether it has been elected as a Proposer or Voter.

**Block Generation**. The elected Proposer proposes a block. Unapproved transactions that have not yet been included in the blockchain are shared among nodes in the network via P2P and stored in an area of each node called the mempool. The node selected as the Proposer generates a block from the unapproved transactions remaining in its mempool and proposes it to the Voters.

**Block Verification**. The block proposed by the Proposer is verified by elected Voters. Each Voter votes on whether the block is correct or not, and the votes are replicated by Tendermint-BFT to the other Voters, and if more than 2/3+1 of all Voters vote in favor of the block, the block is officially approved. On the other hand, if a quorum is not reached, the proposed block is rejected and a new round of elections or voting is started over (Tendermint-BFT has several shortcuts depending on the reason for rejection).

![VRF-based Block Generation Round](vrf_based_round.png)

## VRF-based Consensus Group Election

VRF is an algorithm for generating a hash value $t$ that can be used as a cryptographic pseudo-random number. The differences between VRF and typical hash functions or pseudo-random number generators are that only the owner of the private key can generate the hash value $t$, and anyone with the corresponding public key can verify the correctness of the hash value.

A VRF hash generator $k$ generates a proof $\pi$ (VRF Proof) from the message $m$ using its private key $S_k$ as in Equation (1). Here, the hash value $t$ can be acquired from the proof $pi$ using Equation. (2). On the other hand, to verify that the hash value $t$ was generated by the owner of the private key $S_k$ based on the message $m$, the verifier applies the public key $P_k$ for $S_k$, $m$, and $\pi$ to Equation (3) to verify that both hash values are identical.

![VRF Expression](math_expression.png)

```math
\begin{eqnarray}
\pi & = & {\rm vrf\_prove}(S_k, m) \\
t & = & {\rm vrf\_proof\_to\_hash}(\pi)
\end{eqnarray}
\begin{equation}
{\rm vrf\_proof\_to\_hash}(\pi) \overset{\text{?}}{=} {\rm vrf\_verify}(P_k, m, \pi)
\end{equation}
```

With Ostracon, the Proposer and Voters of the next block are selected randomly by a verifiable random number from the Proposer that created the previous block. A VRF Proof field $pi$ is added to the block for this purpose.

The node that receives the new block initiates the election phase. In this phase, it verifies the VRF Proof $\pi$ contained in the block, calculates the VRF hash $t$, which is a "fair pseudo-random number," and selects the Proposer and Voters for this round based on that value. This is done by a simple and fast weighted random sampling based on the probability of selection according to Stake holdings (in other words, based on PoS).

![VRF-based Proposer/Voter Election](vrf_election.png)

The node selected as the Proposer by this phase picks up the unapproved transactions from its own mempool and generates a proposal block (at this point, the block is not confirmed yet). Then, the Proposer calculates VRF Proof $\pi'$ using the previous VRF Hash $t$ that selected itself, the new block height $h$, and the current round $r$ and sets it to the block.

![VRF Prove](math_prove.png)

```math
\begin{eqnarray*}
m_h & = & {\rm SHA256}(h \,\|\, r \,\|\, t_{h-1}) \\
\pi_h & = & {\rm vrf\_prove}(S_i, m_h) \\
t_h & = & {\rm vrf\_proof\_to\_hash}(\pi_h)
\end{eqnarray*}
```

Note that the message $m$ used to calculate the new VRF Proof $\pi$ doesn't involve the hash value of the block itself. We consider that the hash value of the block is inherently insecure because the Proposer who generates the block can obtain a favorable value by trial and error.

![VRF-based Block Generation](vrf_block_generation.png)

A node that is selected as a Voter in the election phase verifies the received Proposal block and votes on it. The votes are replicated by Tendermint-BFT through prevote, precommit, and commit. The block is confirmed if more than a quorum of valid votes are collected.

![VRF-based Block Validation](vrf_block_validation.png)

During the verification phase, the following VRF-related verifications are performed in addition to block verification:

* The Proposer that generated the block must be a node selected based on the VRF hash of its previous block. This can be determined by matching the node that actually generated the block with the Proposer selected by weighted random sampling using the VRF hash $t$.
* The $\pi$ contained in the block must be a VRF Proof generated using the private key of the Proposer. If the $t$ calculated from the VRF Proof $\pi$ matches the $t$ calculated using the `vrf_verify()` function, we can conclude that $\pi$ is not forged.

![VRF Verify](math_verify.png)

```math
{\rm vrf\_verify}(P_i, m_h, \pi_h) \overset{\text{?}}{=} {\rm vrf\_proof\_to\_hash}(\pi_h)
```

By repeating this sequence of rounds, fair random sampling can be chained across all block generation.

![BFT-based Block Generation](bft_round.png)

Recall here that the node that receives the block can deterministically calculate which nodes are the next Proposer and Voters. By revealing the nodes that are responsible for generating and verifying blocks in a given round, we can penalize nodes that are elected but don't actually perform their responsibility or that behave malicious actions such as Eclipse attacks. On the other hand, it's still difficult to predict the Proposer and Voters beyond one block, as they are only revealed for the minimum necessary time.

VRF is currently implemented using Ed25519, and even if a node uses BLS signatures, it also has an Ed25519 key to calculate VRF.

## Voters

In the Ostracon network, Validators mean candidate nodes that hold Stakes and can be elected as Proposers or Voters. The Voters (a subset of Validators) are a new concept introduced in Ostracon for two reasons; first, to make the distribution of rewards to nodes elected as Voters flexible, and second, to allow the ratio of Byzantine assumptions to be changed in networks with different trust policies for the participant nodes (as a result of the configuration, if the number of Voters is set to match the number of Validators, the behavior will be the same as in Tendermint).

Voter selections use a pseudo-random function $r$ to generate a sequence of random numbers in order to randomly select multiple nodes from a single VRF hash $t$. It's more important that $r$ is fast, simple to implement, has no variant by different interpretations, and saves memory since $t$ already has the properties of a cryptographic pseudo-random number. Ostracon uses a fast shift-register type pseudo-random number generation algorithm, called SplitMix64, for Voter selection.

## Disciplinary Scheme for Failures

Although Ostracon's consensus scheme works correctly even if a few nodes fail, it's ideal that failed nodes aren't selected for the consensus group in order to avoid wasting network and CPU resources. In particular, for cases that aren't caused by general asynchronous messaging problems, such as intentional malpractice, evidence of the behavior (whether malicious or not) will be shared and action will be taken to eliminate the candidate from the selection process by forfeiting the Stake.