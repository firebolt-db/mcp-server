# [](#release-notes)Release notes

Firebolt continuously releases updates so that you can benefit from the latest and most stable service. These updates might happen daily, but we aggregate release notes to cover a longer time period for easier reference. The most recent release notes from the latest version are below.

- See the [Release notes archive](/Reference/release-notes/release-notes-archive.html) for earlier-version release notes.

Firebolt might roll out releases in phases. New features and changes may not yet be available to all accounts on the release date shown.

## [](#firebolt-release-notes---version-418)Firebolt Release Notes - Version 4.18

### [](#new-features)New Features

**Users can now ALTER their corresponding USER object without administrative or RBAC permissions**  
Users can now [ALTER](/sql_reference/commands/access-control/alter-user.html) their corresponding [USER](/Overview/organizations-accounts.html#users) object and change its properties without needing role-based access control permissions ([RBAC](/Overview/Security/Role-Based%20Access%20Control/)). This enhancement simplifies user self-management by reducing the dependency on administrative permissions. Restrictions remain for sensitive properties including [logins or service accounts](/Overview/organizations-accounts.html#organizations), which require higher-level permissions.

**Use a LOCATION object to store credentials for authentication**  
You can now use [CREATE LOCATION](/sql_reference/commands/data-definition/create-location.html) to create a `LOCATION` object in your Firebolt account. Use `LOCATION` to store credentials and authenticate to external systems without needing to provide static credentials each time you run a query or create a table. `LOCATION` works with ([RBAC](/Overview/Security/Role-Based%20Access%20Control/)) so you can manage permissions securely. You can view detailed information about your locations including source type, URL, description, owner, and creation time in [information\_schema.locations](/sql_reference/information-schema/locations.html).

**Added creation timestamps for tables, views, indexes, and locations**  
Use creation timestamps in `information_schema` views for [tables](/sql_reference/information-schema/tables.html), [views](/sql_reference/information-schema/views.html), [indexes](/sql_reference/information-schema/indexes.html), and [locations](/sql_reference/information-schema/locations.html) to help track objects for data management.

**Added support for SQL pipe syntax**  
Firebolt now supports [SQL Pipe syntax](/sql_reference/commands/queries/pipe.html), an alternative way to structure SQL queries using the `|>` operator. This syntax allows for a linear, step-by-step flow of query transformations, improving readability and simplifying query composition. It supports all standard SQL operations and can be combined with traditional SQL syntax.

**Added wildcard character functionality to `READ_PARQUET` and `READ_CSV` to simultaneously read multiple files**  
You can use wildcard characters such as `*` or `?` to specify a file URL as a [glob pattern](https://en.wikipedia.org/wiki/Glob_%28programming%29) in the [READ\_PARQUET](/sql_reference/functions-reference/table-valued/read_parquet.html) and [READ\_CSV](/sql_reference/functions-reference/table-valued/read_csv.html) table-valued functions to read multiple files simultaneously. This enhancement simplifies managing large datasets by reducing the need to make multiple function calls.

**Added functionality to transfer ownership of objects in the Firebolt Workspace**  
You can now [transfer ownership](/Guides/security/ownership.html#transfer-ownership-using-the-firebolt-workspace) of Firebolt objects through the **Firebolt Workspace** user interface (UI). You can transfer ownership of individual objects or bulk transfer owned by a specific user. You can also delete individual objects or in bulk, helping to simplify the management of object ownership within the UI.

### [](#performance-improvements)Performance Improvements

**Enabled result and subresult caching for queries with window functions**  
Enabled [result and subresult caching](/Overview/queries/understand-query-performance-subresult.html) for queries that contain [window functions](/sql_reference/functions-reference/window/), which can reduce query runtimes by storing previous results and enhance overall query performance and efficiency.

### [](#bug-fixes)Bug Fixes

**Fixed an issue where `CREATE VIEW` statements did not preserve the order of named function parameters**  
An issue was resolved where [CREATE VIEW](/sql_reference/commands/data-definition/create-view.html) statements did not maintain the correct order of named function parameters, which could lead to syntax errors when querying the view. This fix improves query reliability by ensuring the proper order of function parameters.