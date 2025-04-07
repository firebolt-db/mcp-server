# [](#synchronous-queries)Synchronous queries

Synchronous queries in Firebolt process SQL statements and wait for a response before proceeding with other operations. These queries are best suited for interactive analytics, dashboards, and data retrieval where low-latency performance is essential. Synchronous queries complete within a single request-response cycle.

Synchronous queries are the default query mode for submitting SQL statements in Firebolt. All statements in the [SQL reference](/sql_reference/) guide can be used inside a synchronous query.

## [](#how-to-submit-a-synchronous-query)How to submit a synchronous query

You can submit a synchronous query using the user interface (UI) in the Firebolt **Develop Space**. Every SQL statement submitted using the UI is a synchronous query. For more information about how to submit a SQL statement using the UI, see [Get started using SQL](/Guides/getting-started/get-started-sql.html).

You can also submit a synchronous query programmatically using the Firebolt API. The following are required prerequisites to submit a query programmatically:

1. **A Firebolt account** – Ensure that you have access to an active Firebolt account. If you don’t have access, you can [sign up for an account](https://www.firebolt.io/sign-up). For more information about how to register with Firebolt, see [Get started with Firebolt](/Guides/getting-started/).
2. **A Firebolt service account** – You must have access to an active Firebolt [service account](/Guides/managing-your-organization/service-accounts.html), which facilitates programmatic access to Firebolt.
3. **A user associated with the Firebolt service account** – You must associate a [user](/Guides/managing-your-organization/managing-users.html#-users) with your service account, and the user must have the necessary permissions to run the query on the specified database using the specified engine.
4. **Sufficient permissions** If you want to query user data through a specific engine, you must have sufficient permissions on the engine, as well as on any tables and databases you access.

To submit a synchronous query programatically, use a Firebolt Driver to send an HTTP request with the SQL statement to Firebolt’s API endpoint.

### [](#use-a-firebolt-driver)Use a Firebolt driver

Use a Firebolt driver to connect to a Firebolt database, authenticate securely, and run SQL statements with minimal setup. The driver provides built-in methods for running SQL statements, handling responses, and managing connections. All Firebolt drivers support synchronous queries. See the documentation for each driver for specific details on how to submit synchronous queries programmatically:

- [Node.js SDK](/Guides/developing-with-firebolt/connecting-with-nodejs.html) – Firebolt Node.js SDK
- [Python SDK](/Guides/developing-with-firebolt/connecting-with-Python.html) – Firebolt Python SDK
- [JDBC Driver](/Guides/developing-with-firebolt/connecting-with-jdbc.html) – Firebolt JDBC Driver
- [SQLAlchemy](/Guides/developing-with-firebolt/connecting-with-sqlalchemy.html) – Firebolt SQLAlchemy adapter
- [.NET SDK](/Guides/developing-with-firebolt/connecting-with-net-sdk.html) – Firebolt .NET SDK
- [Go SDK](/Guides/developing-with-firebolt/connecting-with-go.html) – Firebolt Go SDK

### [](#submit-a-query)Submit a query

After setting up a Firebolt driver, submit a query to verify connectivity and validate your credentials.

Submitting a query through a Firebolt drivers and SDKs have similar formats. The following code example shows how to establish a connection to a Firebolt database using a service account’s credentials, runs a simple `SELECT` statement, retrieves and prints the result using the [Python SDK](/Guides/developing-with-firebolt/connecting-with-Python.html). For other languages, consult the specific driver for details.

```
from firebolt.db import connect
from firebolt.client.auth import ClientCredentials

id = "service_account_id"
secret = "service_account_secret"
engine_name = "your_engine_name"
database_name = "your_test_db"
account_name = "your_account_name"

query = """
    SELECT 42;
    """

with connect(
    engine_name=engine_name,
    database=database_name,
    account_name=account_name,
    auth=ClientCredentials(id, secret),
) as connection:
    cursor = connection.cursor()

    cursor.execute(query)
    for row in cursor.fetchall():
        print(row)
```

#### [](#handling-long-running-synchronous-queries)Handling long-running synchronous queries

Synchronous queries maintain an open HTTP connection for the duration of the query, and stream results back as they become available. While there is no strict time limit, queries running longer than one hour may experience connectivity interruptions. If the HTTP connection is lost, some SQL statements, including `INSERT`, continue to run by default, while `SELECT` statements are cancelled. You can modify this behavior using the [cancel\_query\_on\_connection\_drop](/Reference/system-settings.html#query-cancellation-mode-on-connection-drop) setting.

To avoid connection issues, consider submitting long-running queries as [asynchronous](/API-reference/using-async-queries.html) queries.

#### [](#check-query-status)Check query status

The queries running on an engine are available in the [engine\_running\_queries](/sql_reference/information-schema/engine-running-queries.html) view.

#### [](#cancel-a-query)Cancel a query

A running synchronous query can be cancelled using the [CANCEL](/sql_reference/commands/queries/cancel.html) statement as follows:

```
CANCEL QUERY '<query_id>';
```

Use the query ID retrieved from the [engine\_running\_queries](/sql_reference/information-schema/engine-running-queries.html) view to cancel a specific query.

## [](#error-handling)Error handling

Common errors and solutions when using synchronous queries:

Error Type Cause Solution **Connection loss** The HTTP connection is interrupted. Depending on the type of query, the query may still be running. Check [engine\_running\_queries](/sql_reference/information-schema/engine-running-queries.html) to verify, and use the `cancel_query_on_connection_drop` setting to modify behavior. **Engine does not exist or you don’t have permission to access it** The user lacks required permissions. Ensure the user has `USAGE` permission on the engine and that the engine exists.