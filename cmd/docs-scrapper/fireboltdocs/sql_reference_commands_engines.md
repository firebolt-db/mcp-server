## [](#engine-commands)Engine commands

Engine commands allow you to create, modify, start, stop, and manage Firebolt engines. Engines are compute resources that process queries and handle data ingestion. These commands help control engine availability, configuration, and resource allocation based on workload needs.

Use the following to manage Firebolt engines:

- [ALTER ENGINE](/sql_reference/commands/engines/alter-engine.html) – Modify an existing engine’s configuration, such as node type, cluster count, or concurrency settings.
- [CREATE ENGINE](/sql_reference/commands/engines/create-engine.html) – Define and provision a new Firebolt engine with specific compute resources.
- [DROP ENGINE](/sql_reference/commands/engines/drop-engine.html) – Remove an existing engine.
- [START ENGINE](/sql_reference/commands/engines/start-engine.html) – Activate an engine to make it available for running queries.
- [STOP ENGINE](/sql_reference/commands/engines/stop-engine.html) – Shut down an engine to free up resources when not in use.
- [USE ENGINE](/sql_reference/commands/engines/use-engine.html) – Set an active engine for running queries in the current session.

Each command provides fine-grained control over compute resources, ensuring that queries run efficiently while managing costs. Select a command from the list for syntax and usage details.