# MCPChain MVP Roadmap

MCPChain is a decentralized infrastructure layer that builds on the Model Context Protocol (MCP) to provide on-chain trust, registries, and verifiable interactions for AI agents and tools. This MVP roadmap outlines the core milestones needed to demonstrate its core loop: tool registration, invocation, and attestation.

---

## MVP Goals

- **Decentralized MCP Server Registry**  
  MCP tool providers can register their servers on-chain with metadata, schema hashes, and stake deposits.

- **Basic Stake-to-Register Mechanism**  
  Server providers must lock a minimal stake to list their tools, providing skin in the game.

- **On-Chain Attestation Logging**  
  AI agents interacting with registered MCP tools can submit verifiable interaction logs (tool used, input/output hash, timestamps).

---

## Phase 1: Chain Bootstrap (Cosmos SDK + CometBFT)

**Objective**: Set up MCPChain blockchain using the Cosmos SDK.

### Tasks:
- Scaffold new chain using `ignite` or manual SDK setup.
- Configure:
  - Chain ID and genesis params
  - Validator node(s)
  - CometBFT networking
- Start local single-node testnet

**Tech**:  
- Cosmos SDK (Golang)  
- CometBFT

---

## Phase 2: Core Modules Development

### 1. MCP Server Registry Module
Register and manage MCP tool metadata on-chain.

#### Structures:
- `RegisteredServer`:
  - `id`, `owner_address`, `description`, `endpoint_url`
  - `mcp_schema_hash`, `staked_amount`, `registration_timestamp`

#### Messages:
- `MsgRegisterServer`
- `MsgUpdateServerDetails`
- `MsgRemoveServer`

#### Queries:
- `QueryServerById(id)`
- `QueryServersByOwner(owner_address)`
- `QueryAllServers()`

---

### 2. Simplified Staking Module
Require a stake deposit to register an MCP server.

#### Structures:
- `StakeRecord`:
  - `server_id`, `amount`, `staker_address`

#### Logic:
- Lock testnet tokens during `MsgRegisterServer`
- Optional: allow unstaking upon removal

---

### 3. Attestation Logging Module
Log interaction between agent and tool.

#### Structures:
- `AttestationRecord`:
  - `interaction_id`, `agent_id`, `tool_id`
  - `tool_name`, `request_hash`, `response_hash`, `timestamp`

#### Messages:
- `MsgLogAttestation`

#### Queries:
- `QueryAttestationById(id)`
- `QueryAttestationsByAgent(agent_id)`
- `QueryAttestationsByToolServer(tool_id)`

---

## Phase 3: MCP Tooling & Reference Stack

### Reference MCP Server (Golang)
- MCP-compliant server exposing:
  - `echo(text)`, `getCurrentTime()`
- JSON schema served via HTTP

### Reference MCP Client CLI
- Discovers tool metadata from MCPChain
- Invokes MCP server endpoint
- Logs interaction via `MsgLogAttestation`

**Goal**: Full loop → `discover → interact → attest`

---

## Phase 4: End-to-End Flow + Demo

### Demo Steps:
1. Launch MCPChain node
2. Run MCP Server
3. Stake + Register Server via CLI
4. Run MCP Client → invoke tool
5. Log attestation
6. Query chain for server metadata and attestation

### Testing:
- Unit tests for each Cosmos module
- Manual CLI-driven integration test
- Optional: Write demo script for automation

---

## Deliverables

- Cosmos SDK appchain: MCPChain with 3 core modules
- CLI client to register, stake, query, and attest
- MCP-compliant server + test tools
- Demo script showing end-to-end invocation
- Documentation + API schema

---

## Post-MVP Stretch Goals

- DID + VC integration for agent and server identities
- Full tokenomics + slashing for malicious tools
- TEE or ZK-based interaction proof system
- Frontend dashboard (React or Tauri)
- OAuth-style MCP authorization layer

---
