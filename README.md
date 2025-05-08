# MCPChain
MCPChain brings the Model Context Protocol to Web3 by anchoring agent tools, trust registries, and attestable executions on-chain.

## Start the chain

```bash
cd mcpchain
ignite chain build
```

### [Optional] Delete Previous Chain State (Safe Dev Reset)
```bash
rm -rf ~/.mcpchain
```

### Init Chain
```bash
mcpchaind init validator --chain-id mcpchain
```

### Create Key
```bash
mcpchaind keys add validator
```

### Replace `"stake"` with `"umcp"` in `genesis.json`
```bash
sed -i '' 's/"stake"/"umcp"/g' ~/.mcpchain/config/genesis.json
```

### Add Funds
```bash
mcpchaind genesis add-genesis-account $(mcpchaind keys show validator -a) 1000000000umcp
```

### Generate gentx using updated denom
```bash
mcpchaind genesis gentx validator 500000000umcp --chain-id mcpchain
```

### Collect and validate
```bash
mcpchaind genesis collect-gentxs
mcpchaind genesis validate
```

### Start the chain with a minimum gas price
```bash
mcpchaind start --minimum-gas-prices=0.001umcp
```