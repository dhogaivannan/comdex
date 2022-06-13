package types

/*
comdex tx gov submit-proposal add-assets ATOM,CMDX,CMST,OSMO uatom,ucmdx,ucmst,uosmo 1000000,1000000,1000000,1000000 0,0,0,0 --from cooluser --chain-id test-1 --keyring-backend test --title "nothing" --description "not" --deposit 100000000stake -y
comdex q asset assets

comdex tx gov submit-proposal add-lend-pool cmdx 1,2,3 1,0,1 --from cooluser --chain-id test-1 --keyring-backend test --title "nothing" --description "not" --deposit 100000000stake -y
comdex q lend pools

comdex tx gov submit-proposal add-lend-pairs 2 1 0 1 1.5 1.6 1 1 1 1 --from cooluser --chain-id test-1 --keyring-backend test --title "nothing" --description "not" --deposit 100000000stake -y
comdex q lend pairs

comdex tx gov submit-proposal add-asset-to-pair-mapping 2 1 --from cooluser --chain-id test-1 --keyring-backend test --title "nothing" --description "not" --deposit 100000000stake
comdex q lend asset-pair-mappings

comdex tx lend lend 2 100000000ucmdx 1 --from cooluser --chain-id test-1 --keyring-backend test

*/
