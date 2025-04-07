# [](#system-engine)System Engine

Firebolt’s system engine enables running various metadata-related queries without having to start an engine. The system engine is always available for you in all databases to select and use.

The system engine supports running the following commands:

- All [access control](/sql_reference/commands/access-control/) commands
- All [engine](/sql_reference/commands/engines/) commands
- Most [data definition](/sql_reference/commands/data-definition/) commands. The following commands are not supported:
  
  - [ALTER TABLE DROP PARTITION](/sql_reference/commands/data-definition/alter-table.html)
  - [CREATE AGGREGATING INDEX](/sql_reference/commands/data-definition/create-aggregating-index.html)
  - [CREATE EXTERNAL TABLE](/sql_reference/commands/data-definition/create-external-table.html)
  - [CREATE TABLE AS SELECT](/sql_reference/commands/data-definition/create-fact-dimension-table-as-select.html)
- Most [metadata](/sql_reference/commands/metadata/) commands. The following commands are not supported:
  
  - [SHOW CACHE](/sql_reference/commands/metadata/show-cache.html)
- Non-data-accessing [SELECT](/sql_reference/commands/queries/select.html) queries like `SELECT CURRENT_TIMESTAMP()`
- [SELECT](/sql_reference/commands/queries/select.html) queries on some [information\_schema](/sql_reference/information-schema/) views:
  
  - [information\_schema.accounts](/sql_reference/information-schema/accounts.html)
  - [information\_schema.applicable\_roles](/sql_reference/information-schema/applicable-roles.html)
  - [information\_schema.transitive\_applicable\_roles](/sql_reference/information-schema/transitive-applicable-roles.html)
  - [information\_schema.columns](/sql_reference/information-schema/columns.html)
  - [information\_schema.catalogs](/sql_reference/information-schema/catalogs.html)
  - [information\_schema.enabled\_roles](/sql_reference/information-schema/enabled-roles.html)
  - [information\_schema.engines](/sql_reference/information-schema/engines.html)
  - [information\_schema.indexes](/sql_reference/information-schema/indexes.html)
  - [information\_schema.logins](/sql_reference/information-schema/logins.html)
  - [information\_schema.network\_policies](/sql_reference/information-schema/network_policies.html)
  - [information\_schema.service\_accounts](/sql_reference/information-schema/service-accounts.html)
  - [information\_schema.tables](/sql_reference/information-schema/tables.html)
  - [information\_schema.users](/sql_reference/information-schema/users.html)
  - [information\_schema.views](/sql_reference/information-schema/views.html)

## [](#using-the-system-engine-via-the-firebolt-manager)Using the system engine via the Firebolt manager

1. In the Firebolt manager, choose the Databases icon in the navigation pane.
2. Click on the SQL Workspace icon for the desired database. In case you have no database in your account - create one first.
3. From the engine selector in the SQL Workspace, choose System Engine, then run one of the supported queries.

## [](#using-the-system-engine-via-sdks)Using the system engine via SDKs

### [](#python-sdk)Python SDK

Connect via the connector without specifying the engine\_name. Database parameter is optional.

System engine does not need a database defined. If you wish to connect to an existing database and run metadata queries with the system engine, just specify the name of your database.

**Example**

```
from firebolt.db import connect
from firebolt.client import DEFAULT_API_URL
from firebolt.client.auth import ClientCredentials

client_id = "<service_account_id>"
client_secret = "<service_account_secret>"
account_name = "<your_account_name>"

with connect(
   database="<any_db_here>", # Omit this parameter if you don't need db-specific operations
   auth=ClientCredentials(client_id, client_secret),
   account_name=account_name,
   api_endpoint=DEFAULT_API_URL,
) as connection:

   cursor = connection.cursor()

   cursor.execute("SHOW CATALOGS")

   print(cursor.fetchall())
```

Guidance on creating service accounts can be found in the [service account](/Guides/managing-your-organization/service-accounts.html) section.

### [](#other-sdks)Other SDKs

Any other Firebolt connector can also be used similarly, as long as the engine name is omitted.

## [](#system-engine-limitations)System Engine Limitations

### [](#supported-queries-for-system-engine)Supported queries for system engine

System engine only supports running the metadata-related queries listed above. Additional queries will be supported in future versions.

### [](#rate-limits-for-system-engines)Rate Limits for System Engines

To ensure fair and consistent access to the System Engine for all users, we have introduced rate limits that govern resource usage per account. These limits are designed to prevent resource contention and ensure optimal performance for everyone.

When the rate limits are exceeded on the system engine, the system will return the following error: `429: Account system engine resources usage limit exceeded`. This error typically occurs when an account submits an excessive number of queries or executes highly complex queries that surpass the allocated resource thresholds.

**What to Do If You Encounter Rate Limits**

If you receive the 429 error, consider these steps to resolve the issue:

- Switch to a User Engine: Offload your workloads to a dedicated User Engine if possible. User Engines do not have the same rate limits, making them better suited for higher workloads or complex operations.
- Review your query patterns and ensure they are not unnecessarily complex or resource-intensive. Use best practices to write efficient queries that minimize resource consumption.
- Contact Support: If you believe your account has been rate-limited unfairly or you anticipate requiring higher limits, reach out to our support team to discuss adjusting your account’s thresholds.

**Best Practices to Avoid Rate Limits**

- Avoid running multiple concurrent queries that heavily use system resources.
- Leverage Firebolt’s indexing and other optimization features to streamline your queries.
- Regularly audit your workloads and usage patterns to align with the system’s best practices.

**Why This Matters**

These rate limits are critical for maintaining a fair and robust environment where all users can achieve reliable performance without disruption from resource-heavy neighbors. This measure aligns with our commitment to delivering consistent and high-quality service across all accounts.

For additional support or questions, please contact our support team or refer to our documentation on optimizing query performance.