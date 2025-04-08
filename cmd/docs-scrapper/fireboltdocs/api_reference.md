# [](#api-reference)API reference

The Firebolt API enables programmatic interaction with Firebolt databases for running SQL statements, retrieving data, and managing engines. Use API calls to submit queries, retrieve results, and perform administrative tasks without the user interface (UI).

Firebolt offers official SDKs and drivers to simplify API usage. These drivers interface between your application and Firebolt, handling authentication, SQL statement submission, and result processing.

![Use a service account and a driver to connect to the Firebolt API which returns a result.](../assets/images/API-workflow.png)

To submit an API request, set up a Firebolt driver and use it to send a query to Firebolt, as explained in the following sections.

**Topics:**

- [Prerequisites](#prerequisites) – Set up your account and credentials before submitting an API request.
- [Set up a driver](#set-up-a-driver) – Download, install, and configure a Firebolt driver to send queries using the Firebolt API.
- [Submit a query](#submit-a-query) – Use a driver to connect to Firebolt and submit a query.

## [](#prerequisites)Prerequisites

Before you submit API queries, you need the following:

1. **A Firebolt account** – Ensure that you have access to an active Firebolt account. If you don’t have access, you can [sign up for an account](https://www.firebolt.io/sign-up). For more information about how to register with Firebolt, see [Get started with Firebolt](/Guides/getting-started/).
2. **A Firebolt service account** – You must have access to an active Firebolt [service account](/Guides/managing-your-organization/service-accounts.html), which facilitates programmatic access to Firebolt.
3. **A user associated with the Firebolt service account** – You must associate a [user](/Guides/managing-your-organization/managing-users.html#-users) with your service account, and the user must have the necessary permissions to run the query on the specified database using the specified engine.
4. **Sufficient permissions** If you want to query user data through a specific engine, you must have sufficient permissions on the engine, as well as on any tables and databases you access.

## [](#set-up-a-driver)Set up a driver

Drivers are software components that facilitate communication between applications and databases. Use a Firebolt driver to connect to a Firebolt database, authenticate securely, and run SQL statements with minimal setup.

Use a Firebolt driver for the following:

- **Simplified API access** – Manage authentication and request formatting, eliminating the need for manual API calls. Requires only installation and basic configuration to connect and run SQL statements.
- **Optimized performance** – Improve query processing and connection management for faster response times.
- **Secure authentication** – Use service accounts and industry-standard methods to ensure secure access.

Firebolt provides multiple drivers and SDKs. Refer to the following [driver documentation](/Guides/integrations/integrations.html) for installation instructions:

- [Node.js SDK](/Guides/developing-with-firebolt/connecting-with-nodejs.html) – For JavaScript-based applications.
- [Python SDK](/Guides/developing-with-firebolt/connecting-with-Python.html) – For Python-based applications and data workflows.
- [JDBC Driver](/Guides/developing-with-firebolt/connecting-with-jdbc.html) – For Java applications.
- [SQLAlchemy](/Guides/developing-with-firebolt/connecting-with-sqlalchemy.html) – For ORM-based integrations in Python.
- [.NET SDK](/Guides/developing-with-firebolt/connecting-with-net-sdk.html) – For applications running on the .NET framework.
- [Go SDK](/Guides/developing-with-firebolt/connecting-with-go.html) – For applications using the Go programming language.

## [](#submit-a-query)Submit a query

After setting up a Firebolt driver, submit a query to verify connectivity and validate your credentials.

Submitting a query through a Firebolt drivers and SDKs have similar formats. The following code example shows how to submit a query using the [Python SDK](/Guides/developing-with-firebolt/connecting-with-Python.html). For other languages, consult the specific driver for details:

```
from firebolt.db import connect
from firebolt.client.auth import ClientCredentials

id = "service_account_id"
secret = "service_account_secret"
engine_name = "your_engine_name"
database_name = "your_test_db"
account_name = "your_account_name"

firstQuery = """
    SELECT 42;
    """
secondQuery = """
    SELECT 'my second query';
"""
    
with connect(
    engine_name=engine_name,
    database=database_name,
    account_name=account_name,
    auth=ClientCredentials(id, secret),
) as connection:
    cursor = connection.cursor()
    cursor.execute(firstQuery)
    for row in cursor.fetchall():
        print(row)
    # The cursor can be reused for multiple queries.
    cursor.execute(secondQuery)
    for row in cursor.fetchall():
        print(row)
```

### [](#query-types)Query types

Firebolt supports two types of query modes: **synchronous** and **asynchronous** queries.

A [synchronous query](/API-reference/using-sync-queries.html) waits for a response before proceeding. This mode is ideal for interactive queries that require immediate results, such as dashboard queries or user-initiated requests. Firebolt maintains an open HTTP connection for the duration of the query and streams results back as they become available.

An [asynchronous query](/API-reference/using-async-queries.html) runs in the background, allowing your application to continue executing other tasks. This is useful for long-running queries, such as [INSERT](/sql_reference/commands/data-management/insert.html), or [ALTER ENGINE](/sql_reference/commands/engines/alter-engine.html), where waiting for a response is unnecessary. The query status can be checked periodically using a query token.