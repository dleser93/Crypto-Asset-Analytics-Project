# VASP Solvency Assessment
Repository of our group project about "VASP Solvency Assessment" for the lecture "Crypto Asset Analytics". The group members of this group project were Dennis Leser, Peter Lahnsteiner and Michael Kashofer.

## Research Questions
* How can we automatically measure asset holdings of cryptoasset exchanges?
* How could we compare holdings vs. user balances and assess solvency?

## Proof of Solvency
Total CEX Reserves ≥ Total CEX Liabilities
* Proof of Reserves from the Homepage of the selected exchange (BitMEX)
* Proof of Liabilities with Merkle Tree

## Binance Proof
The sample_users0_Binance_test.csv is a synthetic dataset to be fed into the Binance zk-merkle proof (https://github.com/binance/zkmerkle-proof-of-solvency/tree/main). The sample_users0_test.csv is an input for the preprocessing notebook. 
