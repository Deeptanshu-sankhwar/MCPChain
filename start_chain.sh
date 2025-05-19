#!/bin/bash

set -e

echo "📦 Building MCPChain..."
cd mcpchain
rm -rf ~/.mcpchain
ignite chain build

echo "🔧 Initializing chain..."
mcpchaind init validator --chain-id mcpchain

echo "🔐 Creating key..."
mcpchaind keys add validator --keyring-backend test

echo "💱 Replacing 'stake' with 'umcp' in genesis.json..."
sed -i '' 's/"stake"/"umcp"/g' ~/.mcpchain/config/genesis.json

echo "💸 Adding genesis account..."
ADDR=$(mcpchaind keys show validator -a --keyring-backend test)
mcpchaind genesis add-genesis-account $ADDR 1000000000umcp

echo "🪙 Generating gentx..."
mcpchaind genesis gentx validator 500000000umcp --chain-id mcpchain --keyring-backend test

echo "📬 Collecting and validating genesis transactions..."
mcpchaind genesis collect-gentxs
mcpchaind genesis validate

echo "🚀 Starting the chain..."
mcpchaind start --minimum-gas-prices=0.001umcp
