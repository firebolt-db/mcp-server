# [](#asynchronous-queries)Asynchronous queries

An asynchronous query runs in the background and returns a successful response once it is accepted by the computing cluster, so that a client can proceed with other tasks without waiting for the statement to finish. The status of an asynchronous query can be checked at specified intervals, which provides flexibility, so that you can check the query’s status at meaningful times based on the expected duration of the operation. For example, a user can avoid unnecessary resource consumption by only checking the status periodically, rather than maintaining an open connection for the entire duration of the query, which might be unreliable or unnecessary for certain tasks.

Asynchronous queries are ideal for long-running SQL statements, such as `INSERT`, `STOP ENGINE`, and `ALTER ENGINE`, where keeping an HTTP connection open is both unreliable and unnecessary, and where the statement might return zero rows. In addition, tracking them can be challenging. Using an asynchronous query allows you to check the status of operations at intervals, based on the expected duration.

You should use asynchronous queries for any supported operation that may take more than a few minutes for which there are no results.

**Supported asynchronous queries**

- [INSERT](/sql_reference/commands/data-management/insert.html) – Inserts one or more values into a specified table.
- [COPY FROM](/sql_reference/commands/data-management/copy-from.html) – Loads data from an Amazon S3 bucket into Firebolt.
- [COPY TO](/sql_reference/commands/data-management/copy-to.html) – Copies the result of a `SELECT` query to an Amazon S3 location.
- [VACUUM](/sql_reference/commands/data-management/vacuum.html) – Optimizes tablets for query performance.
- [CREATE AGGREGATING INDEX](/sql_reference/commands/data-definition/create-aggregating-index.html) – Creates an index for precomputing and storing frequent aggregations.
- [CREATE AS SELECT](/sql_reference/commands/data-definition/create-fact-dimension-table-as-select.html) – Creates a table and loads data into it based on a `SELECT` query.
- [Engine commands](/sql_reference/commands/engines/) including [ALTER ENGINE](/sql_reference/commands/engines/alter-engine.html), [STOP ENGINE](/sql_reference/commands/engines/stop-engine.html), and [START ENGINE](/sql_reference/commands/engines/start-engine.html). By default, Firebolt engines finish running queries before returning results, which can take significant time. Starting an engine can also take more than a few minutes.

## [](#how-to-submit-an-asynchronous-query)How to submit an asynchronous query

You can only submit a synchronous query programmatically using the Firebolt API or the following listed drivers. Every SQL statement submitted using the Firebolt **Develop Space** user interface is a synchronous query.

The following are required prerequisites to submit a query programmatically:

1. **A Firebolt account** – Ensure that you have access to an active Firebolt account. If you don’t have access, you can [sign up for an account](https://www.firebolt.io/sign-up). For more information about how to register with Firebolt, see [Get started with Firebolt](/Guides/getting-started/).
2. **A Firebolt service account** – You must have access to an active Firebolt [service account](/Guides/managing-your-organization/service-accounts.html), which facilitates programmatic access to Firebolt.
3. **A user associated with the Firebolt service account** – You must associate a [user](/Guides/managing-your-organization/managing-users.html#-users) with your service account, and the user must have the necessary permissions to run the query on the specified database using the specified engine.
4. **Sufficient permissions** If you want to query user data through a specific engine, you must have sufficient permissions on the engine, as well as on any tables and databases you access.

To submit an asynchronous query via a raw HTTP request, you must use Firebolt protocol version 2.3 or later, while query status can be checked with any client. You can verify the protocol version by checking the X-Firebolt-Protocol-Version header in API response.

## [](#use-a-firebolt-driver)Use a Firebolt Driver

Use a Firebolt driver to connect to a Firebolt database, authenticate securely, and run SQL statements with minimal setup. The driver provides built-in methods for running SQL statements, handling responses, and managing connections. Only some Firebolt drivers support synchronous queries. See the documentation for each driver for specific details on how to submit asynchronous queries programmatically:

- [Python SDK](/Guides/developing-with-firebolt/connecting-with-Python.html) – Firebolt Python SDK
- [Node.js](/Guides/developing-with-firebolt/connecting-with-nodejs.html) – Firebolt Node SDK

## [](#submit-a-query)Submit a query

Submitting a query through a Firebolt drivers and SDKs have similar formats. The following code example shows how to submit an asynchronous query using the [Python SDK](/Guides/developing-with-firebolt/connecting-with-Python.html). For other languages, consult the specific driver for details:

The following code example establishes a connection to a Firebolt database using a service account, submits an asynchronous `INSERT` statement that groups generated numbers, periodically checks its run status, and then retrieves the row count from the `example` table:

```
from time import sleep

from firebolt.db import connect
from firebolt.client.auth import ClientCredentials

id = "service_account_id"
secret = "service_account_secret"
engine_name = "your_engine_name"
database_name = "your_test_db"
account_name = "your_account_name"

query = """
    INSERT INTO example SELECT idMod7 as id
    FROM (
        SELECT id%7 as idMod7
        FROM GENERATE_SERIES(1, 10000000000) s(id)
    )
    GROUP BY idMod7;
    """

with connect(
    engine_name=engine_name,
    database=database_name,
    account_name=account_name,
    auth=ClientCredentials(id, secret),
) as connection:
    cursor = connection.cursor()

    cursor.execute_async(query) # Needs firebolt-sdk 1.9.0 or later
    # Token lets us check the status of the query later
    token = cursor.async_query_token
    print(f"Query Token: {token}")

    # Block until the query is done
    # You can also do other work here
    while connection.is_async_query_running(token):
        print("Checking query status...")
        sleep(5)

    status = "Success" if connection.is_async_query_successful(token) else "Failed"
    print(f"Query Status: {status}")

    cursor.execute("SELECT count(*) FROM example;")  # Should contain 7 rows
    for row in cursor.fetchall():
        print(row)
```

### [](#check-query-statusy)Check query statusy

The query status token is included in the initial response when the query is submitted. If needed, you can also retrieve the token from the [engine\_running\_queries](/sql_reference/information-schema/engine-running-queries.html) view.

To check the status of an asynchronous query, use the token with the `CALL fb_GetAsyncStatus` function as follows:

```
CALL fb_GetAsyncStatus('<async_token>');
```

The previous code example returns a single row with the following schema:

Column Name Data Type Description account\_name TEXT The name of the account where the asynchronous query was submitted. user\_name TEXT The name of the user who submitted the asynchronous query. request\_id TEXT Unique ID of the request which submitted the asynchronous query. query\_id TEXT Unique ID of the asynchronous query. status TEXT Current status of the query: SUSPENDED, RUNNING, CANCELLED, FAILED, SUCCEEDED or IN\_DOUBT. submitted\_time TIMESTAMPTZ The time the asynchronous query was submitted. start\_time TIMESTAMPTZ The time the async query was most recently started. end\_time TIMESTAMPTZ If the asynchronous query is completed, the time it finished. error\_message TEXT If the asynchronous query failed, the error message from the failure. retries LONG The number of times the asynchronous query has retried. scanned\_bytes LONG The number of bytes scanned by the asynchronous query. scanned\_rows LONG The number of rows scanned by the asynchronous query.

### [](#cancel-a-query)Cancel a query

A running asynchronous query can be cancelled using the [CANCEL](/sql_reference/commands/queries/cancel.html) statement as follows:

```
CANCEL QUERY '<query_id>';
```

In the previous code example, retrieve the query ID from the [engine\_running\_queries](/sql_reference/information-schema/engine-running-queries.html) view or from the original query submission response.

## [](#error-handling)Error handling

Error Type Cause Solution **Protocol version mismatch** Using an outdated Firebolt protocol version. Make sure your driver supports async queries. **Query failure** The query encounters an execution error. Check the error message in `fb_GetAsyncStatus` and validate the query syntax. **Token not found** The provided async query token is invalid or expired. Verify that the correct token is being used and that the query has not expired. **Engine does not exist or you don’t have permission to access it** The specified Firebolt engine is not running or you don’t have permission to access it. Start the engine before submitting the query and double check permissions.