# [](#engine-permissions)Engine permissions

In Firebolt, an **engine** is a compute resource that processes data and serves queries. Engines provide **full workload isolation**, allowing multiple workloads to run independently while sharing access to the same data. Engines are also **decoupled from databases**, which means:

- An engine can connect to multiple databases.
- A database can be accessed by multiple engines.

The following table outlines the privileges that can be granted for engines within a particular account:

Privilege Description GRANT Syntax REVOKE Syntax USAGE Allows using an engine to run queries. `GRANT USAGE ON ENGINE <engine_name> TO <role>;` `REVOKE USAGE ON ENGINE <engine_name> FROM <role>;` OPERATE Allows stopping and starting an engine. `GRANT OPERATE ON ENGINE <engine_name> TO <role>;` `REVOKE OPERATE ON ENGINE <engine_name> FROM <role>;` MODIFY Allows altering engine properties or dropping the engine. `GRANT MODIFY ON ENGINE <engine_name> TO <role>;` `REVOKE MODIFY ON ENGINE <engine_name> FROM <role>;` MONITOR \[USAGE] Enables the tracking of engine queries through the `engine_running_queries` view for active queries and the `engine_query_history` view for past queries in `information_schema`. `GRANT MONITOR USAGE ON ENGINE <engine_name> TO <role>;` `REVOKE MONITOR USAGE ON ENGINE <engine_name> FROM <role>;` ALL \[PRIVILEGES] Grants all privileges over the engine to a role. `GRANT ALL ON ENGINE <engine_name> TO <role>;` `REVOKE ALL ON ENGINE <engine_name> FROM <role>;`

If a user lacks **USAGE** and **OPERATE** privileges for an engine, they will not be able to select or interact with the engine via the Firebolt UI.

## [](#examples-of-granting-engine-permissions)Examples of granting engine permissions

### [](#usage-permission)USAGE permission

The following code example grants the role `developer_role` permission to use the `myEngine` engine for executing queries:

```
GRANT USAGE ON ENGINE "myEngine" TO developer_role;
```

### [](#operate-permission)OPERATE permission

The following code example gives the role `developer_role` permission to start and stop the `myEngine` engine:

```
GRANT OPERATE ON ENGINE "myEngine" TO developer_role;
```

### [](#modify-permission)MODIFY permission

The following code example grants the role `developer_role` permission to alter properties or drop the `myEngine` engine:

```
GRANT MODIFY ON ENGINE "myEngine" TO developer_role;
```

### [](#monitor-usage-permission)MONITOR \[USAGE] permission

The following code example grants the role `developer_role` permission to see the query history and currently running queries for the engine `myEngine`:

```
GRANT MONITOR USAGE ON ENGINE "myEngine" TO developer_role;
```

### [](#all-permissions)ALL permissions

The following code example grants the role `developer_role` with all engine permissions on `myEngine`:

```
GRANT ALL ON ENGINE "myEngine" TO developer_role;
```