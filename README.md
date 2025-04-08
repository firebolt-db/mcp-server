<h1 align="center">
  <br>
  <a href="https://www.firebolt.io"><img src="https://cdn.prod.website-files.com/5e8a264ceaf4870394477fc7/5e8a264ceaf4879f75477fdd_logo_website.svg" alt="Firebolt" width="300"></a>
  <br>
  MCP Server
  <br>
</h1>

<h4 align="center">
A Model Context Protocol implementation that connects your LLM to Firebolt Data Warehouse
</h4>

<p align="center">
  <a href="https://github.com/firebolt-db/mcp-server/releases">
    <img src="https://img.shields.io/github/v/release/firebolt-db/mcp-server" alt="Release">
  </a>
  <a href="https://github.com/firebolt-db/mcp-server/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/firebolt-db/mcp-server" alt="License">
  </a>
  <a href="https://go.dev">
    <img src="https://img.shields.io/badge/go-1.24.1-blue" alt="Go Version">
  </a>
  <a href="https://github.com/firebolt-db/mcp-server/actions/workflows/build.yml">
    <img src="https://img.shields.io/github/actions/workflow/status/firebolt-db/mcp-server/build.yml" alt="Build Status">
  </a>
</p>

<p align="center">
  <a href="#key-features">Key Features</a> |
  <a href="#how-to-use">How To Use</a> |
  <a href="#requirements">Requirements</a> |
  <a href="#architecture">Architecture</a> |
  <a href="#development">Development</a> |
  <a href="#license">License</a>
</p>

![screenshot](https://img.example.firebolt.io/mcp-server-demo.gif)

## Key Features

* **LLM Integration with Firebolt** - Connect your AI assistants directly to your data warehouse
  - Enable AI agents to autonomously query your data and build analytics solutions
  - Provide LLMs with specialized knowledge of Firebolt's capabilities and features

* **SQL Query Execution** 
  - Direct query execution against Firebolt databases
  - Support for multiple query types and execution modes

* **Documentation Access**
  - Comprehensive Firebolt documentation available to the LLM
  - SQL reference, function reference, and more

* **Account Management**
  - Connect to different accounts and engines
  - Manage authentication seamlessly

* **Multi-platform Support**
  - Run on any platform supporting Go binaries
  - Docker container support for easy deployment

## How To Use

To get started with the Firebolt MCP Server, you'll need a Firebolt service account. If you don't have a Firebolt account yet, [sign up here](https://www.firebolt.io/signup).

### Option 1: Use the Docker image

```bash
# Run with Docker
docker run -p 8080:8080 \
  -e FIREBOLT_MCP_CLIENT_ID=your-client-id \
  -e FIREBOLT_MCP_CLIENT_SECRET=your-client-secret \
  -e FIREBOLT_MCP_TRANSPORT=sse \
  firebolt/mcp-server:latest
```

### Option 2: Download and run the binary

```bash
# Download the latest release for your platform from:
# https://github.com/firebolt-db/mcp-server/releases

# Run the server
./firebolt-mcp-server \
  --client-id your-client-id \
  --client-secret your-client-secret \
  --transport sse
```

### Connecting your LLM

Once the server is running, you can connect to it using any MCP-compatible client. For example:

```bash
# Using the OpenAI API with MCP extension
curl -X POST https://api.openai.com/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -d '{
    "model": "gpt-4",
    "messages": [
      {"role": "system", "content": "You are a data analyst working with Firebolt."},
      {"role": "user", "content": "How many users did we have last month?"}
    ],
    "tools": [
      {
        "type": "mcp",
        "mcp": {
          "endpoint": "http://localhost:8080",
          "auth": {
            "type": "bearer",
            "token": "YOUR_TOKEN"
          }
        }
      }
    ]
  }'
```

## Requirements

- Firebolt service account credentials (client ID and client secret)
- For development: Go 1.24.1 or later
- For deployment: Docker (optional)

## Architecture

The Firebolt MCP Server implements the [Model Context Protocol](https://github.com/anthropics/anthropic-cookbook/tree/main/model_context_protocol) specification, providing:

1. **Tools** - Task-specific capabilities provided to the LLM:
   - `Connect`: Establish connections to Firebolt engines and databases
   - `Docs`: Access Firebolt documentation
   - `Query`: Execute SQL queries against Firebolt

2. **Resources** - Data that can be referenced by the LLM:
   - Documentation articles
   - Account information
   - Database schema
   - Engine statistics

3. **Prompts** - Predefined instructions for the LLM:
   - Firebolt Expert: Prompts the model to act as a Firebolt specialist

## Development

To set up the development environment:

```bash
# Clone this repository
git clone https://github.com/firebolt-db/mcp-server.git

# Go into the repository
cd mcp-server

# Install Task (if you don't have it already)
go install github.com/go-task/task/v3/cmd/task@latest

# Update Go dependencies
task mod

# Build the application
task build
```

### Running tests

```bash
go test ./...
```

### Building Docker image

```bash
docker build -t firebolt-mcp-server .
```

## License

MIT

---

> [firebolt.io](https://www.firebolt.io) &nbsp;&middot;&nbsp;
> GitHub [@firebolt-db](https://github.com/firebolt-db) &nbsp;&middot;&nbsp;
> Twitter [@FireboltDB](https://twitter.com/FireboltDB)
