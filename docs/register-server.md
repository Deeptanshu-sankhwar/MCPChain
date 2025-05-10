# Steps to register an MCP server on chain

## 1. **Wipe old chain state**

```bash
rm -rf ~/.mcpchain
```

---

## 2. **Initialize chain**

```bash
mcpchaind init validator --chain-id mcpchain
```

---

## 3. **Create key**

```bash
mcpchaind keys add mykey --keyring-backend test
```

---

## 4. **Fix denom in genesis**

```bash
sed -i '' 's/"stake"/"umcp"/g' ~/.mcpchain/config/genesis.json
```

---

## 5. **Add funds to account**

```bash
mcpchaind genesis add-genesis-account $(mcpchaind keys show mykey -a --keyring-backend test) 1000000000umcp
```

---

## 6. **Generate validator gentx**

```bash
mcpchaind genesis gentx mykey 500000000umcp --chain-id mcpchain --keyring-backend test
```

---

## 7. **Collect and validate genesis**

```bash
mcpchaind genesis collect-gentxs
mcpchaind genesis validate
```

---

## 8. **Start the chain**

```bash
mcpchaind start --minimum-gas-prices=0.001umcp
```

---

## 9. **Send the transaction**

```bash
mcpchaind tx toolregistry register-server \
  $(mcpchaind keys show mykey -a --keyring-backend test) \
  "My AI Server" \
  "https://my-server.com" \
  "0x123..." \
  "1000000umcp" \
  --from mykey \
  --keyring-backend test \
  --chain-id mcpchain \
  --gas 200000 \
  --gas-prices 0.1umcp \
  -y
```

---
