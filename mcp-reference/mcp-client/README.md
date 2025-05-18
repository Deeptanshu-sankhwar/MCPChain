# MCP Client

---

## What Does `mcp-client` Do?

It’s a command-line tool that interacts with your **MCPChain** and **registered MCP servers**. Specifically, it supports:

### 1. **Discover registered MCP servers on-chain**

```bash
mcp-client discover
```

* Uses: `mcpchaind query toolregistry all-servers`
* Output: Pretty-printed list of all on-chain registered tools.

---

### 2. **Invoke MCP server endpoints**

```bash
mcp-client invoke --endpoint http://localhost:8080/echo --payload '{"text": "hello"}'
```

```bash
./mcp-client invoke \
  --endpoint http://localhost:8080/echo \
  --payload '{"text": "hello"}' \
  --agent-id cosmos1xyz... \
  --tool-id 0 \
  --tool-name echo-tool \
  --from mykey
```

or for `/time`:

```bash
mcp-client invoke --endpoint http://localhost:8080/time
```

* Makes a **GET** or **POST** request to an MCP-compliant server.
* Then, it submits an on-chain attestation by calling `MsgLogAttestation` with:

  * agent_id
  * tool_id
  * tool_name
  * request_hash
  * response_hash
  * timestamp

---

### 3. **Hash utility**

```bash
mcp-client hash "some string"
```

* Returns the SHA256 hash of a string (useful for debug/test).

---

## How to Run It Locally

### 1. Build the binary:

```bash
cd mcp-client
go build -o mcp-client
```

### 2. Run MCPChain and register a test server

```bash
# In your MCPChain directory
mcpchaind start --minimum-gas-prices=0.001umcp
```

Register a tool using:

```bash
mcpchaind tx toolregistry register-server \
  $(mcpchaind keys show mykey -a --keyring-backend test) \
  "Echo Tool" \
  "http://localhost:8080/echo" \
  "schemahash123" \
  "1000000umcp" \
  --from mykey \
  --keyring-backend test \
  --chain-id mcpchain \
  --fees 1000umcp \
  -y
```

### 3. Start the reference MCP server (in another terminal):

```bash
cd mcp-server
go run main.go
```

### 4. Discover registered tools:

```bash
./mcp-client discover
```

### 5. Call a tool endpoint:

```bash
./mcp-client invoke --endpoint http://localhost:8080/echo --payload '{"text": "hello MCP!"}'
```

```bash
./mcp-client invoke \
  --endpoint http://localhost:8080/echo \
  --payload '{"text": "hello"}' \
  --agent-id cosmos1xyz... \
  --tool-id 0 \
  --tool-name echo-tool \
  --from mykey
```