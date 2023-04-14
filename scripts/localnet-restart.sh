#!/bin/sh

# Set localnet settings
BINARY=crescentd
CHAIN_ID=localnet
CHAIN_DIR=./data
RPC_PORT=26657
GRPC_PORT=9090
MNEMONIC_1="guard cream sadness conduct invite crumble clock pudding hole grit liar hotel maid produce squeeze return argue turtle know drive eight casino maze host"
MNEMONIC_2="friend excite rough reopen cover wheel spoon convince island path clean monkey play snow number walnut pull lock shoot hurry dream divide concert discover"
MNEMONIC_3="fuel obscure melt april direct second usual hair leave hobby beef bacon solid drum used law mercy worry fat super must ritual bring faculty"
GENESIS_COINS=10000000000000stake,10000000000000airdrop,10000000000000uatom

$BINARY start --home $CHAIN_DIR/$CHAIN_ID --pruning=nothing --grpc.address="0.0.0.0:$GRPC_PORT" > $CHAIN_DIR/$CHAIN_ID.log 2>&1 &
