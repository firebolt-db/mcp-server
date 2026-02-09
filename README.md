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

Before you start, ensure you have either:
- A Firebolt [service account](https://docs.firebolt.io/Guides/managing-your-organization/service-accounts.html) with a client ID and client secret.
- A [Firebolt Core](https://www.firebolt.io/core) instance URL.

### Installing the MCP Server

You can run the Firebolt MCP Server either via Docker or by downloading the binary.

#### Option 1: Run with Docker

[//]: # (x-release-please-start-version)
```bash
docker run \
  --rm \
  --network host \
  -e FIREBOLT_MCP_CLIENT_ID=your-client-id \
  -e FIREBOLT_MCP_CLIENT_SECRET=your-client-secret \
  ghcr.io/firebolt-db/mcp-server:0.6.0
```
[//]: # (x-release-please-end)

#### Option 2: Run the Binary

[//]: # (x-release-please-start-version)
```bash
# Download the binary for your OS from:
# https://github.com/firebolt-db/mcp-server/releases/tag/v0.6.0

./firebolt-mcp-server \
  --client-id your-client-id \
  --client-secret your-client-secret
```
[//]: # (x-release-please-end)

### Connecting Your LLM

Once the MCP Server is installed, you can connect various LLM clients. Choose your client below.

#### Cursor

Install the Firebolt MCP server in Cursor with one click. You will need to set your credentials in **Cursor Settings > Tools & MCP** (edit the added server or `mcp.json`).

**Firebolt Cloud** (service account):

[![Add to Cursor](https://cursor.com/deeplink/mcp-install-dark.svg)](https://cursor.com/en-US/install-mcp?name=firebolt&config=eyJjb21tYW5kIjoiZG9ja2VyIiwiYXJncyI6WyJydW4iLCItaSIsIi0tcm0iLCItLW5ldHdvcmsiLCJob3N0IiwiLWUiLCJGSVJFQk9MVF9NQ1BfQ0xJRU5UX0lEIiwiLWUiLCJGSVJFQk9MVF9NQ1BfQ0xJRU5UX1NFQ1JFVCIsIi1lIiwiRklSRUJPTFRfTUNQX0RJU0FCTEVfUkVTT1VSQ0VTIiwiZ2hjci5pby9maXJlYm9sdC1kYi9tY3Atc2VydmVyOmxhdGVzdCJdLCJlbnYiOnsiRklSRUJPTFRfTUNQX0NMSUVOVF9JRCI6InlvdXItY2xpZW50LWlkIiwiRklSRUJPTFRfTUNQX0NMSUVOVF9TRUNSRVQiOiJ5b3VyLWNsaWVudC1zZWNyZXQiLCJGSVJFQk9MVF9NQ1BfRElTQUJMRV9SRVNPVVJDRVMiOiJ0cnVlIn19)

**Firebolt Core** (local or self-hosted):

[![Add to Cursor](https://cursor.com/deeplink/mcp-install-dark.svg)](https://cursor.com/en-US/install-mcp?name=firebolt-core&config=eyJjb21tYW5kIjoiZG9ja2VyIiwiYXJncyI6WyJydW4iLCItaSIsIi0tcm0iLCItLW5ldHdvcmsiLCJob3N0IiwiLWUiLCJGSVJFQk9MVF9NQ1BfQ09SRV9VUkwiLCItZSIsIkZJUkVCT0xUX01DUF9ESVNBQkxFX1JFU09VUkNFUyIsImdoY3IuaW8vZmlyZWJvbHQtZGIvbWNwLXNlcnZlcjpsYXRlc3QiXSwiZW52Ijp7IkZJUkVCT0xUX01DUF9DT1JFX1VSTCI6Imh0dHA6Ly9sb2NhbGhvc3Q6MzQ3MyIsIkZJUkVCT0xUX01DUF9ESVNBQkxFX1JFU09VUkNFUyI6InRydWUifX0=)

> [!IMPORTANT]
> Cursor does not support MCP resources yet. This MCP server uses resources heavily, so it will not work in Cursor unless resource usage is disabled. The buttons above already set `FIREBOLT_MCP_DISABLE_RESOURCES=true` for you. If you configure manually, you must set this environment variable.

#### VS Code (Copilot Chat)

To integrate MCP with Copilot Chat in VS Code, see the official guide: [Extending Copilot Chat with the Model Context Protocol](https://docs.github.com/en/copilot/customizing-copilot/extending-copilot-chat-with-mcp).

Add the Firebolt MCP server to your MCP configuration (e.g. via Command Palette > **MCP: Open User Configuration**). Example for **Firebolt Cloud** with Docker:

```json
{
  "mcpServers": {
    "firebolt": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "--network", "host",
        "-e", "FIREBOLT_MCP_CLIENT_ID",
        "-e", "FIREBOLT_MCP_CLIENT_SECRET",
        "ghcr.io/firebolt-db/mcp-server:latest"
      ],
      "env": {
        "FIREBOLT_MCP_CLIENT_ID": "your-client-id",
        "FIREBOLT_MCP_CLIENT_SECRET": "your-client-secret"
      }
    }
  }
}
```

Replace `your-client-id` and `your-client-secret` with your [Firebolt service account](https://docs.firebolt.io/Guides/managing-your-organization/service-accounts.html) credentials.

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
            "--network", "host",
            "-e", "FIREBOLT_MCP_CLIENT_ID",
            "-e", "FIREBOLT_MCP_CLIENT_SECRET",
            "ghcr.io/firebolt-db/mcp-server:0.6.0"
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

    **Connecting to Firebolt Core**

    If you are using [Firebolt Core](https://docs.firebolt.io/firebolt-core), you can connect by providing the Core URL instead of service account credentials.

    To use Firebolt Core with **Claude Desktop**:
    
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
            "--network", "host",
            "-e", "FIREBOLT_MCP_CORE_URL",
            "ghcr.io/firebolt-db/mcp-server:0.6.0"
          ],
          "env": {
            "FIREBOLT_MCP_CORE_URL": "http://localhost:3473"
          }
        }
      }
    }
    ```
    [//]: # (x-release-please-end)

4. Save the config and restart Claude Desktop.

More details: [Claude MCP Quickstart Guide](https://modelcontextprotocol.io/quickstart/user)

#### Manual configuration (other MCP clients)

For any MCP client that uses a JSON config (e.g. `mcpServers`), you can add the Firebolt server with the following canonical configuration. **Firebolt Cloud** (Docker):

```json
{
  "mcpServers": {
    "firebolt": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "--network", "host",
        "-e", "FIREBOLT_MCP_CLIENT_ID",
        "-e", "FIREBOLT_MCP_CLIENT_SECRET",
        "ghcr.io/firebolt-db/mcp-server:latest"
      ],
      "env": {
        "FIREBOLT_MCP_CLIENT_ID": "your-client-id",
        "FIREBOLT_MCP_CLIENT_SECRET": "your-client-secret"
      }
    }
  }
}
```

For **Firebolt Core**, use `FIREBOLT_MCP_CORE_URL` (e.g. `http://localhost:3473`) in `env` and pass `-e FIREBOLT_MCP_CORE_URL` in `args` instead of the client ID/secret. For Cursor, also set `FIREBOLT_MCP_DISABLE_RESOURCES=true` in `env` and in `args`.

#### Using SSE Transport

By default, the MCP Server uses STDIO as the transport mechanism.  
However, Server-Sent Events (SSE) are also supported and require additional configuration.

To enable SSE, set the `--transport` CLI flag (or the `FIREBOLT_MCP_TRANSPORT` environment variable) to `sse`.

Optionally, you can specify the address the server should listen on by setting the `--transport-sse-listen-address` CLI flag (or the `FIREBOLT_MCP_TRANSPORT_SSE_LISTEN_ADDRESS` environment variable).


#### Requiring LLMs to present docs read proof before connecting

To provide wider context to LLMs before connecting to Firebolt and running queries, by default `firebolt_connect` tool
requires the LLM to present a read proof of the Firebolt documentation (by querying the `firebolt_docs_overview` tool).

While this provides a good starting point for LLMs ensuring it has full context of Firebolt documentation, at the same time this may lead to slower responses and higher token consumption.

To disable this requirement, set the `--skip-docs-proof` CLI bool flag (or the `FIREBOLT_MCP_SKIP_DOCS_PROOF` environment variable) to `false`.

## Architecture

Firebolt MCP Server implements the [Model Context Protocol](https://modelcontextprotocol.io/introduction), providing:

1. **Tools** - Task-specific capabilities provided to the LLM:
    - `firebolt_docs_overview`: Access basic Firebolt documentation overview
    - `firebolt_connect`: Establish connections to Firebolt engines and databases
    - `firebolt_query`: Execute SQL queries against Firebolt
    - `firebolt_docs_search`: Search Firebolt documentation for any details

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
