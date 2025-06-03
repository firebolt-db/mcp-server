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
  <a href="#key-features">Key Features</a> |
  <a href="#how-to-use">How To Use</a> |
  <a href="#connecting-your-llm">Connecting Your LLM</a> |
  <a href="#architecture">Architecture</a> |
  <a href="#development">Development</a>
</p>

## Key Features

**LLM Integration with Firebolt**

- Connect your AI assistants directly to your data warehouse  
- Enable AI agents to autonomously query data and generate insights  
- Provide LLMs with deep knowledge of Firebolt SQL, features, and documentation

**SQL Query Execution**

- Support for multiple query types and execution modes  
- Direct access to Firebolt databases

**Documentation Access**

- Grant LLMs access to comprehensive Firebolt docs, SQL reference, function lists, and more

**Account Management**

- Seamless authentication with Firebolt service accounts  
- Connect to different engines and workspaces

**Multi-platform Support**

- Runs anywhere Go binaries are supported  
- Official Docker image available for easy deployment

## How To Use

Before you start, ensure you have a Firebolt [service account](https://docs.firebolt.io/Guides/managing-your-organization/service-accounts.html) with a client ID and client secret.

### Installing the MCP Server

You can run the Firebolt MCP Server either via Docker or by downloading the binary.

#### Option 1: Run with Docker

[//]: # (x-release-please-start-version)
```bash
docker run \
  --rm \
  -e FIREBOLT_MCP_CLIENT_ID=your-client-id \
  -e FIREBOLT_MCP_CLIENT_SECRET=your-client-secret \
  ghcr.io/firebolt-db/mcp-server:0.3.1
```
[//]: # (x-release-please-end)

#### Option 2: Run the Binary

[//]: # (x-release-please-start-version)
```bash
# Download the binary for your OS from:
# https://github.com/firebolt-db/mcp-server/releases/tag/v0.3.1

./firebolt-mcp-server \
  --client-id your-client-id \
  --client-secret your-client-secret
```
[//]: # (x-release-please-end)

### Connecting Your LLM

Once the MCP Server is installed, you can connect various LLM clients.

Below are integration examples for **Claude Desktop**.
For other clients like **VSCode Copilot Chat** and **Cursor**, please refer to their official documentation.

#### Claude Desktop

To integrate with Claude Desktop using **Docker**:

1. Open the Claude menu and select **Settings…**.
2. Navigate to **Developer** > **Edit Config**.
3. Update the configuration file (`claude_desktop_config.json`) to include:

    [//]: # (x-release-please-start-version)
    ```json
    {
      "mcpServers": {
        "firebolt": {
          "command": "docker",
          "args": [
            "run",
            "-i",
            "--rm",
            "-e", "FIREBOLT_MCP_CLIENT_ID",
            "-e", "FIREBOLT_MCP_CLIENT_SECRET",
            "ghcr.io/firebolt-db/mcp-server:0.3.1"
          ],
          "env": {
            "FIREBOLT_MCP_CLIENT_ID": "your-client-id",
            "FIREBOLT_MCP_CLIENT_SECRET": "your-client-secret"
          }
        }
      }
    }
    ```
    [//]: # (x-release-please-end)

    To use the **binary** instead of Docker:

    ```json
    {
      "mcpServers": {
        "firebolt": {
          "command": "/path/to/firebolt-mcp-server",
          "env": {
            "FIREBOLT_MCP_CLIENT_ID": "your-client-id",
            "FIREBOLT_MCP_CLIENT_SECRET": "your-client-secret"
          }
        }
      }
    }
    ```

4. Save the config and restart Claude Desktop.

More details: [Claude MCP Quickstart Guide](https://modelcontextprotocol.io/quickstart/user)

#### GitHub Copilot Chat (VSCode)

To integrate MCP with Copilot Chat in VSCode, refer to the official documentation:

👉 [Extending Copilot Chat with the Model Context Protocol](https://docs.github.com/en/copilot/customizing-copilot/extending-copilot-chat-with-mcp)

#### Cursor Editor

To set up MCP in Cursor, follow their guide:

👉 [Cursor Documentation on Model Context Protocol](https://docs.cursor.com/context/model-context-protocol)

#### Using SSE Transport

By default, the MCP Server uses STDIO as the transport mechanism.  
However, Server-Sent Events (SSE) are also supported and require additional configuration.

To enable SSE, set the `--transport` CLI flag (or the `FIREBOLT_MCP_TRANSPORT` environment variable) to `sse`.

Optionally, you can specify the address the server should listen on by setting the `--transport-sse-listen-address` CLI flag (or the `FIREBOLT_MCP_TRANSPORT_SSE_LISTEN_ADDRESS` environment variable).

## Architecture

Firebolt MCP Server implements the [Model Context Protocol](https://modelcontextprotocol.io/introduction), providing:

1. **Tools** - Task-specific capabilities provided to the LLM:
    - `firebolt_docs`: Access Firebolt documentation
    - `firebolt_connect`: Establish connections to Firebolt engines and databases
    - `firebolt_query`: Execute SQL queries against Firebolt

2. **Resources** - Data that can be referenced by the LLM:
    - Documentation articles
    - Lists of Accounts, Databases, Engines

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

# Run the tests
task test
```
