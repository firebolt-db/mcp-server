# [](#release-notes-archive)Release notes archive

We provide an archive of release notes for your historical reference.

- [Firebolt Release Notes - Version 4.17](#firebolt-release-notes---version-417)
  
  - [New Features](#new-features)
  - [Performance Improvements](#performance-improvements)
  - [Bug Fixes](#bug-fixes)
- [Firebolt Release Notes - Version 4.16](#firebolt-release-notes---version-416)
  
  - [New Features](#new-features-1)
  - [Performance Improvements](#performance-improvements-1)
  - [Bug Fixes](#bug-fixes-1)
- [Firebolt Release Notes - Version 4.15](#firebolt-release-notes---version-415)
  
  - [New Features](#new-features-2)
  - [Performance Improvements](#performance-improvements-2)
  - [Behavior Changes](#behavior-changes)
  - [Bug Fixes](#bug-fixes-2)
- [Firebolt Release Notes - Version 4.14](#firebolt-release-notes---version-414)
  
  - [New Features](#new-features-3)
  - [Performance Improvements](#performance-improvements-3)
  - [Bug Fixes](#bug-fixes-3)
- [Firebolt Release Notes - Version 4.13](#firebolt-release-notes---version-413)
  
  - [New Features](#new-features-4)
  - [Behavior Changes](#behavior-changes-1)
  - [Bug Fixes](#bug-fixes-4)
- [Firebolt Release Notes - Version 4.12](#firebolt-release-notes---version-412)
  
  - [New Features](#new-features-5)
  - [Performance Improvements](#performance-improvements-4)
  - [Behavior Changes](#behavior-changes-2)
  - [Bug Fixes](#bug-fixes-5)
- [Firebolt Release Notes - Version 4.11](#firebolt-release-notes---version-411)
  
  - [New Features](#new-features-6)
- [Firebolt Release Notes - Version 4.10](#firebolt-release-notes---version-410)
  
  - [New Features](#new-features-7)
  - [Behavior Changes](#behavior-changes-3)
  - [Bug Fixes](#bug-fixes-6)
- [Firebolt Release Notes - Version 4.9](#firebolt-release-notes---version-49)
  
  - [New Features](#new-features-8)
  - [Performance Improvements](#performance-improvements-5)
  - [Bug Fixes](#bug-fixes-7)
- [Firebolt Release Notes - Version 4.8](#firebolt-release-notes---version-48)
  
  - [New Features](#new-features-9)
  - [Performance Improvements](#performance-improvements-6)
  - [Bug Fixes](#bug-fixes-8)
- [DB version 4.7](#db-version-47)
  
  - [New Features](#new-features-10)
  - [Performance Improvements](#performance-improvements-7)
  - [Behavior Changes](#behavior-changes-4)
  - [Bug Fixes](#bug-fixes-9)
- [DB version 4.6](#db-version-46)
  
  - [New Features](#new-features-11)
  - [Behavior Changes](#behavior-changes-5)
  - [Bug Fixes](#bug-fixes-10)
- [DB version 4.5](#db-version-45)
  
  - [New Features](#new-features-12)
  - [Bug Fixes](#bug-fixes-11)
- [DB version 4.4](#db-version-44)
  
  - [New Features](#new-features-13)
  - [Breaking Changes](#breaking-changes)
  - [Bug Fixes](#bug-fixes-12)
- [DB version 4.3](#db-version-43)
  
  - [New Features](#new-features-14)
  - [Performance Improvements](#performance-improvements-8)
  - [Bug Fixes](#bug-fixes-13)
  - [Breaking Changes](#breaking-changes-1)
- [DB version 4.2](#db-version-42)
  
  - [New features](#new-features-15)
  - [Enhancements, changes and new integrations](#enhancements-changes-and-new-integrations)
  - [Breaking Changes](#breaking-changes-2)
- [DB version 4.1](#db-version-41)
  
  - [Resolved issues](#resolved-issues)
- [DB version 4.0](#db-version-40)
  
  - [Enhancements, changes and new integrations](#enhancements-changes-and-new-integrations-1)
  - [Breaking Changes](#breaking-changes-3)
- [DB version 3.34](#db-version-334)
  
  - [Enhancements, changes and new integrations](#enhancements-changes-and-new-integrations-2)
  - [Resolved issues](#resolved-issues-1)
- [DB version 3.33](#db-version-333)
  
  - [Enhancements, changes and new integrations](#enhancements-changes-and-new-integrations-3)
  - [Resolved issues](#resolved-issues-2)
- [DB version 3.32](#db-version-332)
  
  - [New features](#new-features-16)
  - [Enhancements, changes and new integrations](#enhancements-changes-and-new-integrations-4)
  - [Resolved issues](#resolved-issues-3)
- [DB version 3.31](#db-version-331)
  
  - [New features](#new-features-17)
  - [Enhancements, changes and new integrations](#enhancements-changes-and-new-integrations-5)
  - [Resolved issues](#resolved-issues-4)
- [DB version 3.30](#db-version-330)
- [DB version 3.29](#db-version-329)
- [DB version 3.28](#db-version-328)

## [](#firebolt-release-notes---version-417)Firebolt Release Notes - Version 4.17

### [](#new-features)New Features

**Introduced the `IF` function to enhance query readability and simplify conditional expressions**  
The new [`IF`](/sql_reference/functions-reference/conditional-and-miscellaneous/if.html) function simplifies query writing as a more concise alternative to the `CASE WHEN` expression.  
You can now use `IF(<cond_expr>, <then_expr>, <else_expr>)` as a shorter equivalent to `CASE WHEN <cond_expr> THEN <then_expr> ELSE <else_expr> END`.

**Added `INCREMENTAL` index optimization with `VACUUM`**  
The [`VACUUM`](/sql_reference/commands/data-management/vacuum.html) statement now supports an `INDEXES = INCREMENTAL` option, allowing incremental optimization of related indexes. This new mode uses fewer resources compared to a full reevaluation, improving index layouts. Although incremental optimization may not achieve the optimal layout of a full reevaluation, it maintains a balance between performance and resource usage.

**Added `MAX_CONCURRENCY` option to `VACUUM` statement**  
The `VACUUM` command now supports the `MAX_CONCURRENCY` option, enabling you to limit concurrent processes during optimization. This allows for control of the number of concurrent processes in a `VACUUM` operation, optimizing resource usage and improving performance in multi-threaded environments.

**Added longitude wrapping for `GEOGRAPHY` data**  
Firebolt now automatically wraps longitude values outside the range of -180 to 180 degrees when parsing `GEOGRAPHY` data from WKT, GeoJSON, WKB, or using the `ST_GeogPoint` function. For example, `POINT(180.5 1)` is now correctly interpreted as `POINT(-179.5 1)`. This improvement simplifies geographic data handling.

**Enhanced the `EXPLAIN` function to support all SQL statements except for DDL and DCL**  
The [`EXPLAIN`](/sql_reference/commands/queries/explain.html) feature now supports analysis of all SQL statements. However, it does not provide output details for DDL (Data Definition Language) and DCL (Data Control Language) statements.

### [](#performance-improvements)Performance Improvements

**Optimized `COPY FROM` filtering performance**  
Filters applied to pseudo columns, such as `$SOURCE_FILE_NAME` and `$SOURCE_FILE_TIMESTAMP`, are now pushed down to the file listing during the `COPY FROM` process when using multiple URL and pattern locations. This enhancement improves performance by reducing unnecessary data processing and speeds up data loading operations.

### [](#bug-fixes)Bug Fixes

**Fixed latitude handling for `LineString` in WKT**  
Fixed an issue where latitudes outside the valid range of -90 to 90 degrees, in `LineString` data were incorrectly accepted when parsing from WKT. For example, `LINESTRING(0.5 1, 1 90.5)` now correctly returns an error instead of being interpreted as `LINESTRING(0.5 1, -179 89.5)`. This fix enhances data integrity and prevents erroneous geographic entries.

## [](#firebolt-release-notes---version-416)Firebolt Release Notes - Version 4.16

### [](#new-features-1)New Features

**Added `MAX_CONCURRENCY` option to the `VACUUM` statement for enhanced concurrency control**  
The [VACUUM](/sql_reference/commands/data-management/vacuum.html) statement now includes the `MAX_CONCURRENCY` option, allowing users to limit the number of concurrent streams. This improves control over resource usage during `VACUUM` operations.

**Introduced the `INDEXES = ALL | NONE` for the `VACUUM` statement**  
The [VACUUM](/sql_reference/commands/data-management/vacuum.html) statement now supports the `INDEXES = ALL | NONE` option, giving users control over whether indexes are optimized during `VACUUM` operations.

**`VACUUM` now runs automatically**  
Firebolt now automatically evaluates the data layout of tables and runs [VACUUM](/sql_reference/commands/data-management/vacuum.html) to optimize performance and storage efficiency. After [INSERT](/sql_reference/commands/data-management/insert.html), [UPDATE](/sql_reference/commands/data-management/update.html), or [DELETE](/sql_reference/commands/data-management/delete.html) operations modify data, the engine that performed the operation determines whether `VACUUM` is required. This decision is based on factors such as the number of deleted rows and the need to consolidate storage for faster query performance and reduced disk space usage.

**Added support for casting text literals to interval literals**  
Firebolt now supports casting text literals to interval literals using expressions like `'1 month'::INTERVAL`, making it easier to define time intervals in queries.

**Added default value support for `GEOGRAPHY` columns**  
Firebolt now supports default values for columns with the [GEOGRAPHY](/sql_reference/geography-data-type.html#geography-data-type) data type. For example, `CREATE TABLE geo_table (geo_column GEOGRAPHY DEFAULT 'GEOMETRYCOLLECTION EMPTY')` ensures consistency across database entries when no explicit value is provided.

**Added `MIN_CLUSTERS` and `MAX_CLUSTERS` columns to `INFORMATION_SCHEMA.ENGINES`**  
The [INFORMATION\_SCHEMA.ENGINES](/sql_reference/information-schema/engines.html) table now includes `MIN_CLUSTERS` and `MAX_CLUSTERS` columns, providing visibility into cluster configuration for improved database management.

**Added support for `STATEMENT_TIMEOUT` to manage query run time limits**  
Added support for `STATEMENT_TIMEOUT`. This feature specifies the number of milliseconds a statement is allowed to run. Any statement or query exceeding the specified time is canceled. A value of zero disables the timeout by default. Using `STATEMENT_TIMEOUT` helps prevent excessively long-running queries, improving system efficiency and resource use.

**Added the PostgreSQL function `DATE(<arg>)` as an alternative to `<arg>::DATE`**  
Firebolt now supports the `DATE(<arg>)` function, offering an alternative to the `<arg>::DATE` syntax for improved readability and usability in SQL queries.

**Added support for `FROM` first syntax**  
SQL queries can now use `FROM` before `SELECT`, allowing for more flexible query structures such as `FROM t SELECT a, SUM(b) GROUP BY a` or even `FROM t` without a `SELECT` clause.

**Support for AWS PrivateLink is now in public preview**  
[Firebolt now supports AWS PrivateLink](/Guides/security/privatelink.html), allowing Firebolt Enterprise customers to securely access the Firebolt API without exposing traffic to the public internet. AWS PrivateLink enhances security, minimizes data exposure, and improves network reliability by keeping traffic within AWS.

**Added concurrency auto-scaling**  
Engines can now be created with concurrency auto-scaling enabled, or modified to enable concurrency auto-scaling. Setting the `MIN_CLUSTERS` and `MAX_CLUSTERS` parameters on [CREATE ENGINE](/sql_reference/commands/engines/create-engine.html) and [ALTER ENGINE](/sql_reference/commands/engines/alter-engine.html) commands turns on concurrency auto-scaling: the engine will dynamically resize between the specified `MIN_CLUSTERS` and `MAX_CLUSTERS` values to match demand.

**Firebolt introduces three fully managed editions**

Firebolt now offers **Standard, Enterprise, and Dedicated editions**, each designed for different capabilities, security, and scalability needs.

- **Standard**: High-performance, elastic scaling – in and out, up and down – for cost-efficient, fully managed analytics on a single cluster.
- **Enterprise &amp; Dedicated**: Includes scaling capabilities like **multi-cluster scaling**, as well as advanced security features like **AWS PrivateLink**.
- **Dedicated**: Built for regulated industries (finance, healthcare) with **single-tenant infrastructure** and compliance with **HIPAA, SOC 2, ISO**.

Enterprise and Dedicated customers also get **24/7 support** with **faster support response times**, **Slack-based support**, and support from a **designated engineer**. For more information on Firebolt’s editions, refer to the [Pricing and billing](/Overview/billing/) page.

### [](#performance-improvements-1)Performance Improvements

**Introduced pruning for `GEOGRAPHY` columns at the tablet level to enhance query performance**  
Firebolt now prunes [GEOGRAPHY](/sql_reference/geography-data-type.html#geography-data-type) data at the tablet level to enhance query performance. To activate spatial pruning on tables created before this release, run `VACUUM`. For additional details, see our [blog post](https://www.firebolt.io/blog/architecture-and-internal-representation-of-the-geography-data-type).

**Added `INDEX_GRANULARITY` storage parameter to optimize table storage**  
The `CREATE TABLE` statement now supports the `INDEX_GRANULARITY` storage parameter, allowing users to configure internal tablet range sizes for better performance based on query patterns.

### [](#bug-fixes-1)Bug Fixes

**Fixed permission conflicts on public schemas across multiple databases**  
Resolved an issue where granting identical permissions on public schemas in different databases caused conflicts. This fix ensures correct permission application for improved database security.

## [](#firebolt-release-notes---version-415)Firebolt Release Notes - Version 4.15

### [](#new-features-2)New Features

**Improved `EXPLAIN (STATISTICS)` to include estimated row counts and column distinct counts**

The [EXPLAIN (STATISTICS)](/sql_reference/commands/queries/explain.html) function now provides estimated row counts and column distinct counts, when available. This enhancement offers more detailed insights for analyzing query performance.

**Added a Tableau connector for the current version of Firebolt**

[Tableau](https://www.tableau.com/) is a visual analytics platform that empowers users to explore, analyze, and present data through interactive visualizations. The current Firebolt connector in Tableau Exchange supports only an older version of Firebolt. You can now download the latest connector directly from Firebolt and integrate it with [Tableau Desktop](https://www.tableau.com/products/desktop) or [Tableau Server](https://www.tableau.com/products/server). Follow the installation instructions in [Integrate with Tableau](/Guides/integrations/tableau.html) to set up the updated connector.

**Added a DBeaver connector for the current version of Firebolt**

[DBeaver](https://dbeaver.io/) is a free, open-source database administration tool that supports multiple database types, provides a graphical interface for managing databases, running queries, and analyzing data. You can now connect to DBeaver using the [Firebolt JDBC driver](https://docs.firebolt.io/Guides/developing-with-firebolt/connecting-with-jdbc.html). Follow the instructions in [Integrate with DBeaver](/Guides/integrations/dbeaver.html) to set up a connection to DBeaver.

**Added the Firebolt Resource Center to the Firebolt Workspace**

The [Firebolt Resource Center](https://www.firebolt.io/resources) is now accessible from the **Firebolt Workspace**. Select the Firebolt icon in the bottom-right corner to access resources including the [Get started guide](/Guides/getting-started/), [Knowledge Center](https://www.firebolt.io/knowledge-center), [Documentation](https://docs.firebolt.io/), [Release notes](/Reference/release-notes/release-notes.html), Announcements, and a unified search tool covering all Firebolt resources.

### [](#performance-improvements-2)Performance Improvements

**Optimized `LEFT JOIN` conversion for better query performance**

A nested `LEFT JOIN` can now be automatically replaced with a more efficient join when its results are not needed due to filtering in a later step. This optimization occurs when a `LEFT JOIN` removes rows where the right-hand side contains `NULL` values, effectively discarding the extra rows introduced by the earlier `LEFT JOIN`. In such cases, simplifying the join structure improves efficiency without altering query results. This conversion reduces unnecessary operations, lowering computational overhead and enhancing performance.

**Improved performance by allowing multiple `INSERT INTO <tbl> VALUES ...` statements to be combined in a single request**

Workloads that send multiple consecutive `INSERT INTO <tbl> VALUES ...` statements into the same table can now run much faster by sending all statements in a single request separated by semicolons. These statements are now automatically merged and processed together on the server within a single transaction, which means that either all of them succeed or fail. This improvement reduces network overhead and enhances performance for batch data insertion.

### [](#behavior-changes)Behavior Changes

**Use `NULL` instead of empty strings for passing unset TVF parameters**

Table-valued functions (TVFs) such as [LIST\_OBJECTS](/sql_reference/functions-reference/table-valued/list-objects.html), [READ\_PARQUET](/sql_reference/functions-reference/table-valued/read_parquet.html), and [READ\_CSV](/sql_reference/functions-reference/table-valued/read_csv.html) that accept string named parameters like `aws_access_key_id` and `aws_role_arn` will no longer treat empty strings (`''`) as unset arguments. The empty strings will instead be forwarded to the credential provider and may return errors. If you want to pass an explicitly unset parameter, use `NULL` instead.

### [](#bug-fixes-2)Bug Fixes

**Resolved issue in distributed `GROUP BY` and `JOIN` planning**

Resolved a bug in the optimization process for distributed `GROUP BY` and `JOIN` operators. This bug sometimes led to missed optimization opportunities and, in rare cases, incorrect results.

**Fixed a bug in correlated `EXISTS` subqueries that caused duplicated outer tuples in query results**

Fixed a bug with non-trivial correlated `EXISTS` subquery, which is a dependent subquery inside an `EXISTS` condition that references a column from an outer query. An example of this kind of query follows:

```
SELECT *,
  EXISTS(SELECT 1 FROM table2 where COALESCE(table2.col_1, table2.col_2) = table1.col_1)
FROM table1
```

Previously, if an outer table contained a value, and the inner table had two matching values, the outer table’s row would appear twice in the final result instead of just once. This happened because the query checked for matches individually for each row in the inner table, rather than treating the condition as a simple existence check.

This bug fix corrected this issue by ensuring that the `EXISTS` condition only determines whether at least one match exists, without duplicating rows in the outer table. Now, each row in the outer table correctly appears once, with `TRUE` if a match exists and `FALSE` otherwise, improving the accuracy of query results.

## [](#firebolt-release-notes---version-414)Firebolt Release Notes - Version 4.14

### [](#new-features-3)New Features

**Added `E2E_DURATION_US` to include total query time in Firebolt infrastructure for enhanced performance monitoring and optimization**

Added a new column `E2E_DURATION_US` in the system tables `INFORMATION_SCHEMA.ENGINE_RUNNING_QUERIES`, `INFORMATION_SCHEMA.ENGINE_QUERY_HISTORY`, and `INFORMATION_SCHEMA.ENGINE_USER_QUERY_HISTORY` which shows the total time a query has spent within the Firebolt infrastructure. In contrast, `DURATION_US` measures only the time spent using the engine without considering retries or routing. The `E2E_DURATION_US` metric measures the total time a query takes from initiation to final result delivery, and includes all sub-components of latency such as routing, preparation, queuing, compilation, retries, and runtimes. For example, if a query starts a stopped engine, the engine’s startup time is included in the query’s end-to-end duration. This update provides a more accurate representation of total query latency, for performance monitoring and optimization.

**Unhid `scanned_storage_bytes` and `scanned_cache_bytes` from information schema views**

Unhid `scanned_storage_bytes` and `scanned_cache_bytes` columns from `information_schema.engine_query_history` and `information_schema.engine_user_query_history` views. These columns were previously accessible when explicitly used in a `SELECT` clause, but will now appear by default when you use `SELECT *`.

### [](#performance-improvements-3)Performance Improvements

**Enhanced data ingestion performance for `GEOGRAPHY` objects of type `POINT`**

Improved data loading performance for `GEOGRAPHY` objects of type `POINT`, enabling up to four times faster loading of geographical point data for more efficient data integration and analysis.

**Improved file listing times for large external scans**

In operations that read data from Amazon S3 buckets such as external table scans or `COPY FROM` queries, Firebolt lists files in a URL to an Amazon S3 bucket. This process is constrained by the AWS API, which limits file listing to 1,000 files per request. Firebolt has increased the number of concurrent operations so that listing a large number of files is up to 3.5 times faster.

**Added result cache support for cross and complex joins for improved performance**

The [query result cache](/Reference/system-settings.html#result-cache) now supports queries using cross joins or complex joins with `OR` conditions and inequalities. This change reduces redundant calculations, improving query performance.

### [](#bug-fixes-3)Bug Fixes

**`USAGE` permissions are now required to access `INFORMATION_SCHEMA` views**

Accessing `INFORMATION_SCHEMA` views now requires `USAGE` permissions on the database. Queries to `INFORMATION_SCHEMA` will fail if these permissions are missing, ensuring consistent enforcement across permission-restricted queries. Ensure that your database has the necessary permissions to prevent access issues.

**Improved `EXPLAIN` command accuracy for default values of `DATE`, `TIMESTAMP`, and `TIMESTAMPTZ` columns**

The `EXPLAIN` command now displays default values for columns of type `DATE`, `TIMESTAMP`, and `TIMESTAMPTZ` columns. This update fixes a bug that previously caused default values to be shown incompletely, improving clarity and accuracy in query plan analysis.

**Resolved filtering issue for views in `information_schema.tables` to enforce user permissions**

Fixed a bug in `information_schema.tables` which previously listed views that users were not authorized to access. Even though querying these views would fail, users could still see that they existed. Now `information_schema.tables` only lists views that users are allowed to access.

## [](#firebolt-release-notes---version-413)Firebolt Release Notes - Version 4.13

### [](#new-features-4)New Features

**`GRANT ALL ON ACCOUNT` and `REVOKE ALL ON ACCOUNT` statements for role-based privileges**  
The statements `GRANT ALL ON ACCOUNT account_name TO role_name` and `REVOKE ALL ON ACCOUNT account_name FROM role_name` are now supported. They grant or revoke all account-related privileges to the specified role `role_name`.

**Support for nested arrays in Parquet files**  
You can now ingest Parquet files containing nested array structures at any depth. For example: `array(array(array(string)))`.

### [](#behavior-changes-1)Behavior Changes

**Removed secured objects from `information_schema` views**  
Users can now only access information about objects for which they have the appropriate permissions or ownership for in [information\_schema views](/sql_reference/information-schema/views.html).

### [](#bug-fixes-4)Bug Fixes

**`@` character support restored in usernames**  
The usage of character `@` is allowed in usernames again, which was previously restricted. The following statements are now valid and will not cause errors:

```
CREATE USER "user@example.com";
ALTER USER user_name RENAME TO "user@example.com";
```

**Resolved memory overuse during CSV import**  
Resolved a memory overconsumption problem that occurred when importing CSV files into existing tables.

**Resolved `EXPLAIN VACUUM` and `EXPLAIN` to improve error handling and result accuracy**  
The following behavior of `EXPLAIN VACUUM` has been updated:

1. If a table is fully vacuumed, no further actions are performed, and the message “Table is fully vacuumed, no vacuum jobs were executed” is returned to the user.
2. The `EXPLAIN VACUUM` output no longer returns an empty result when the vacuumed object is an aggregating index.
3. `EXPLAIN` has been updated to show an error if the specified relation does not exist.

**Fixed incorrect evaluation of `IS NULL` in outer joins**  
Fixed an issue where `IS NULL` predicates on non-nullable columns from the non-preserving side of an outer join were incorrectly reduced to `FALSE` during common table expression (CTE) optimization. When the optimizer attempted to fuse multiple CTEs, it mistakenly replaced `t2.x IS NULL` with `FALSE`, altering query semantics and producing incorrect results. This occurred because `t2.x`, though defined as non-nullable, became nullable when used in a left join. The fix ensures that `IS NULL` predicates are correctly preserved during optimization.

## [](#firebolt-release-notes---version-412)Firebolt Release Notes - Version 4.12

### [](#new-features-5)New Features

**Added `ST_S2CELLIDFROMPOINT` to retrieve the [S2 Cell ID](http://s2geometry.io/devguide/s2cell_hierarchy) of a `GEOGRAPHY` Point**

You can now use [ST\_S2CELLIDFROMPOINT](/sql_reference/functions-reference/geospatial/st_s2cellidfrompoint.html) to retrieve the S2 cell ID, which identifies the region on Earth that fully contains, or covers, a single Point `GEOGRAPHY` object. You can also specify a cell resolution level.

**Added keyboard shortcuts to the Firebolt Develop Space**

The Firebolt **Develop Space** user interface added the following Windows/Mac [keyboard shortcuts](/Guides/query-data/using-the-develop-workspace.html#keyboard-shortcuts-for-the-develop-space):

- Ctrl + Alt + E / Ctrl + ⌘ + E – Toggle expanding or collapsing query results.
- Ctrl + Alt + N / Ctrl + ⌘ + N – Create a new script.
- Ctrl + Alt + [ / Ctrl + ⌘ + [ – Jump to the previous script.
- Ctrl + Alt + ] / Ctrl + ⌘ + ] – Jump to the next script.

**Introduced the `INFORMATION_SCHEMA.ROUTINES` view for built-in functions and operators**

Added the [INFORMATION\_SCHEMA.ROUTINES](/sql_reference/information-schema/routines.html) view to return information about all of Firebolt’s built-in functions and operators including their database, schema, name, type, return data type, parameter data types, and whether they are deterministic.

**Added support for the `GEOGRAPHY` data type in external tables using CSV and JSON formats**

Firebolt can now read columns of type `GEOGRAPHY` from external tables in CSV or JSON format, which allows the querying of geospatial data including Points and Polygons.

**Added a new `MONITOR USAGE` privilege**

You can use the `MONITOR USAGE` privilege to view all queries running on an engine using [information\_schema.engine\_query\_history](/sql_reference/information-schema/engine-query-history.html) or [information\_schema.engine\_running\_queries](/sql_reference/information-schema/engine-running-queries.html) views.

**Introduced support for network policy `ADD`/`REMOVE` commands**  
Admins can now append or remove specific IP addresses in `ALLOW` or `BLOCK` lists without overriding existing values. This update simplifies network policy management when handling large IP lists and reduces the risk of concurrent updates overwriting each other.

### [](#performance-improvements-4)Performance Improvements

**Improved performance of the `ST_COVERS`, `ST_CONTAINS`, and `ST_INTERSECTS` functions**

Optimized the [ST\_COVERS](/sql_reference/functions-reference/geospatial/st_covers.html), [ST\_CONTAINS](/sql_reference/functions-reference/geospatial/st_contains.html), and [ST\_INTERSECTS](/sql_reference/functions-reference/geospatial/st_intersects.html) functions to improve performance when processing LineStrings and Points with non-intersecting inputs, and Polygons with inputs that do not intersect their boundaries.

**Improved performance of the `REGEXP_LIKE_ANY` function**

The [REGEXP\_LIKE\_ANY](/sql_reference/functions-reference/string/regexp-like-any.html) function now performs more efficiently when matching against multiple patterns by compiling a single combined [RE2](https://github.com/google/re2/) regular expression object instead of evaluating each pattern separately.

### [](#behavior-changes-2)Behavior Changes

**Updated user name rules to improve consistency and validation**

The following changes affect the use of user names in [CREATE USER](/sql_reference/commands/access-control/create-user.html) AND [ALTER USER](/sql_reference/commands/access-control/alter-user.html):

- The `@` character is no longer allowed in user names.
- The range of permissible characters in user names is expanded. For more information, see [CREATE USER](/sql_reference/commands/access-control/create-user.html).
- When renaming a user with [ALTER USER](/sql_reference/commands/access-control/alter-user.html) `old_name RENAME TO new_name`, the `new_name` must now comply with the updated user name rules.
- Any new names created with [CREATE USER](/sql_reference/commands/access-control/create-user.html) must now comply with the updated user name rules.

### [](#bug-fixes-5)Bug Fixes

**Fixed an error where `APACHE_DATASKETCHES_HLL_ESTIMATE` failed for `NULL` inputs**

Resolved an error in the [APACHE\_DATASKETCHES\_HLL\_ESTIMATE](/sql_reference/functions-reference/datasketches/apache-datasketches-hll-estimate.html) function that occurred if any of its input values were `NULL`. The function can now process `NULL` inputs.

**Resolved issue that allowed account lockout on last login**

Fixed an issue where the `ALTER USER SET LOGIN/SERVICE_ACCOUNT=...` statement could lock out the only active login in an account, rendering the account inaccessible. The operation now fails with an explicit error message in such cases.

**Fixed incorrect ownership modification for `information_schema`**

The statement `ALTER SCHEMA information_schema SET OWNER owner_name;` previously succeeded, which was incorrect, because `information_schema` cannot be modified. The operation now fails with an explicit error message.

**Fixed an out-of-memory error during large CSV imports**

Updated the ingestion pipeline for [COPY FROM](/sql_reference/commands/data-management/copy-from.html) to ensure that large CSV files without a predefined schema can load into new tables without causing memory errors. This error did not affect external tables.

**Prevent running queries when using a dropped database**

When the current database does not exist, such as when it has been dropped, most queries fail as expected. We fixed a bug where some queries against specific `information_schema` views, such as `engines`, `catalogs`, `applicable_roles`, would still succeed in such cases. These queries now fail consistently, like all other queries against a non-existent database. For example, running `SELECT * FROM information_schema.engines` when the database is dropped previously worked, but now fails.

## [](#firebolt-release-notes---version-411)Firebolt Release Notes - Version 4.11

### [](#new-features-6)New Features

**Introduced the `GEOGRAPHY` data type and functions for geospatial data handling \[public preview]**

Added a new [GEOGRAPHY](/sql_reference/geography-data-type.html) data type and functions for working with geospatial data. Firebolt supports the three industry standard formats [Well-Known Text (WKT)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry), [Well-Known Binary (WKB)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Well-known_binary), and [GeoJSON](https://datatracker.ietf.org/doc/html/rfc7946) for geospatial data.

This public preview release includes the following functions:

- [ST\_ASBINARY](/sql_reference/functions-reference/geospatial/st_asbinary.html) – Converts shapes of the `GEOGRAPHY` data type to the [Well-Known Binary (WKB)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Well-known_binary) format for geographic objects.
- [ST\_ASEWKB](/sql_reference/functions-reference/geospatial/st_asewkb.html) – Converts shapes of the `GEOGRAPHY` data type to the [extended Well-Known Binary (EWKB)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Format_variations) format using Spatial Reference Identifier (SRID) 4326, which corresponds to the [WGS84](https://en.wikipedia.org/wiki/World_Geodetic_System#WGS_84) coordinate system.
- [ST\_ASGEOJSON](/sql_reference/functions-reference/geospatial/st_asgeojson.html) – Converts shapes of the `GEOGRAPHY` data type to the [GeoJSON](https://datatracker.ietf.org/doc/html/rfc7946) format.
- [ST\_ASTEXT](/sql_reference/functions-reference/geospatial/st_astext.html) – Converts shapes of the `GEOGRAPHY` data type to the [Well-Known Text (WKT)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry) format.
- [ST\_CONTAINS](/sql_reference/functions-reference/geospatial/st_contains.html) – Determines if one `GEOGRAPHY` object fully contains another.
- [ST\_COVERS](/sql_reference/functions-reference/geospatial/st_covers.html) – Determines if one `GEOGRAPHY` object fully encompasses another.
- [ST\_DISTANCE](/sql_reference/functions-reference/geospatial/st_distance.html) – Calculates the shortest distance, measured as a geodesic arc between two `GEOGRAPHY` objects, measured in meters.
- [ST\_GEOGFROMGEOJSON](/sql_reference/functions-reference/geospatial/st_geogfromgeojson.html) – Constructs a `GEOGRAPHY` object from a [GeoJSON](https://datatracker.ietf.org/doc/html/rfc7946) string.
- [ST\_GEOGFROMTEXT](/sql_reference/functions-reference/geospatial/st_geogfromtext.html) – Constructs a `GEOGRAPHY` object from a [Well-Known Text (WKT)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry) string.
- [ST\_GEOGFROMWKB](/sql_reference/functions-reference/geospatial/st_geogfromwkb.html) – Constructs a `GEOGRAPHY` object from a [Well-Known Binary](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Well-known_binary) (WKB) byte string.
- [ST\_GEOGPOINT](/sql_reference/functions-reference/geospatial/st_geogpoint.html) – Constructs a Point in the `GEOGRAPHY` data type created from specified longitude and latitude coordinates.
- [ST\_INTERSECTS](/sql_reference/functions-reference/geospatial/st_intersects.html) – Determines whether two input `GEOGRAPHY` objects intersect each other.
- [ST\_X](/sql_reference/functions-reference/geospatial/st_x.html) – Extracts the longitude coordinate of a `GEOGRAPHY` Point.
- [ST\_Y](/sql_reference/functions-reference/geospatial/st_y.html) – Extracts the latitude coordinate of a `GEOGRAPHY` Point.

**Added keyboard shortcuts to the Firebolt Develop Space**

The user interface in the Firebolt **Develop Space** added the following [keyboard shortcuts](/Guides/query-data/using-the-develop-workspace.html#keyboard-shortcuts-for-the-develop-space):

- Cmd + Enter – Runs the current query.
- Cmd+Shift+Enter – Runs all queries in a script.

**Added the window function `FIRST_VALUE`**

Added a new [FIRST\_VALUE](/sql_reference/functions-reference/window/first-value.html) window function that returns the first value evaluated in a specified window frame.

## [](#firebolt-release-notes---version-410)Firebolt Release Notes - Version 4.10

### [](#new-features-7)New Features

**Added `CREATE TABLE CLONE` to clone an existing table in a database**

You can create a clone of an existing table in a database using `CREATE TABLE CLONE`, which is extremely fast because it copies the table structure and references without duplicating the data. The clone functions independently of the original table. Any changes to the data or schema of either table will not affect the other.

**Added 3-part identifier support for specifying databases in queries**

You can now reference a database other than the current one in queries by using 3-part identifiers, which specify the database, schema, and object. For example, even if you previously selected a database `db` by using `USE DATABASE db`, you can still query a different database by using a query such as `SELECT * FROM other_db.public.t`. The limitation still exists that every query only addresses a single database.

**Added `ALTER TABLE ADD COLUMN` to add a column to an existing table**

You can now use `ALTER TABLE ADD COLUMN` to add columns to Firebolt-managed tables. This functionality is temporarily limited to tables that were created on Firebolt version 4.10 or higher.

**Added support of `ALTER TABLE RENAME` command** You can use `ALTER TABLE RENAME` to change the name of Firebolt-managed tables. This functionality is temporarily limited to tables created on Firebolt version 4.10 or higher.

**Added support for external file access using AWS session tokens**

You can now use `<AWS_SESSION_TOKEN>` with access keys to securely authenticate and access external files on AWS with the following features:

- The [COPY TO](/sql_reference/commands/data-management/copy-to.html) and [COPY FROM](/sql_reference/commands/data-management/copy-from.html) commands.
- [External tables](/Guides/loading-data/working-with-external-tables.html) located in an Amazon S3 bucket.
- The following table-valued functions: `read_parquet`, `read_csv`, and `list_objects`.

### [](#behavior-changes-3)Behavior Changes

**Enhanced PostgreSQL compliance for casting data types from text to float**

Cast from text to floating-point types is now compliant with PostgreSQL with the following improvements:

1. **The correct parsing of positive floats** – A plus sign (`+`) preceding a float is now handled correctly. Example: `'+3.4'`.
2. **Exponent-only input** – Float values starting with an exponent `'e'` or `'E'` are rejected. Example: `'E4'`.
3. **Incomplete exponents** – Float values ending with an exponent without a subsequent exponent value are rejected. Example: `'4e+'`.

**Account-level rate limits implemented for the system engine**

Firebolt has implemented account-level rate limits to ensure equitable resource usage among all users of the system engine. When these limits are exceeded, requests will be rejected with the following error message: `429: Account system engine resources usage limit exceeded`. This rate limit targets accounts with exceptionally high resource consumption. Accounts with typical resource usage should not be affected and require no further action.

### [](#bug-fixes-6)Bug Fixes

**Corrected runtime reporting**

Resolved an issue where the runtime displayed in Firebolt’s user interface and JSON responses omitted including processing times for some query steps.

**Resolved “Invalid Input Aggregate State Type” error with aggregating indexes**

Fixed an issue where the “invalid input aggregate state type” error could occur when queries read from aggregating indexes that defined a `COUNT(*)` aggregate function before other aggregate functions. After this fix, such aggregating indexes can now be queried correctly without needing to be rebuilt.

**Fixed a rare bug in subresult caching logic**

Addressed a rare issue in the logic for caching and reusing subresults that could cause query failures with specific query patterns. This issue did not impact the correctness of query results.

**Resolved issue preventing schema owners from granting “ANY” privileges**

Fixed an issue where schema owners were unable to grant “ANY” privileges on their schema to other users.  
For example:

```
GRANT SELECT ANY ON SCHEMA public TO ...
```

Schema owners can now execute this command which allows the specified user or role to perform SELECT operations on any table.

## [](#firebolt-release-notes---version-49)Firebolt Release Notes - Version 4.9

### [](#new-features-8)New Features

**Added the `enable_result_cache` setting for controlling query result caching during benchmarking**

You can set `enable_result_cache` to `FALSE` to disable the use of Firebolt’s result cache, which is set to `TRUE` by default. Disabling result cashing can be useful for benchmarking query performance. When `enable_result_cache` is disabled, resubmitting the same query will recompute the results rather than retrieving them from cache. For more information, see [Result Cache](/Reference/system-settings.html#result-cache).

**Added `LAG` and `LEAD` support for negative offsets.**

The second parameter in both [LAG](/sql_reference/functions-reference/window/lag.html) and [LEAD](/sql_reference/functions-reference/window/lead.html) can now accept negative numbers. Given a negative number, a `LAG` will become a `LEAD` and vice versa. For example, `LAG(x,-5,3)` is the same as `LEAD(x,5,3)`.

### [](#performance-improvements-5)Performance Improvements

**Faster string searches for case-insensitive simple regular expressions in `REGEXP_LIKE`**

Simple regular expressions in [REGEXP\_LIKE](/sql_reference/functions-reference/string/regexp-like.html) with case-insensitive matching, using the `i` flag, now use the same optimized string search implementation as [ILIKE](/sql_reference/functions-reference/string/ilike.html), achieving up to three times faster runtimes in observed cases.

### [](#bug-fixes-7)Bug Fixes

**Empty character classes in regular expressions**

Fixed a rare case where empty character classes were mistakenly interpreted as valid character classes instead of being treated as raw characters. In cases like `[]a]`, the expression is now correctly interpreted as a pattern that matches any single character from the list `]a`, rather than treating `[]` as an empty character class followed by `a]`.

**Trailing backslash in regular expressions**

Fixed a rare case where invalid regular expressions with a trailing backslash `\` were accepted.

## [](#firebolt-release-notes---version-48)Firebolt Release Notes - Version 4.8

### [](#new-features-9)New Features

**Introduced new bitwise shift functions `BIT_SHIFT_RIGHT` and `BIT_SHIFT_LEFT`**

The following bitwise shift functions are now supported:

- `BIT_SHIFT_RIGHT` shifts the bits of a number to the right by a specified number of positions, which effectively divides the number by `2` for each position shifted.
- `BIT_SHIFT_LEFT` shifts the bits of a number to the left by a specified number of positions, which effectively multiples the number by `2` for each position shifted.

**Introduced new trigonometric functions `ACOS`, `ATAN`, `ASIN`, `COS`, `COT`, `TAN`, `DEGREES`, and `PI`**

The following trigonometric functions are now supported:

- `ACOS` calculates the arccosine of a value in radians.
- `ATAN` calculates the arctangent of a value in radians.
- `ASIN` calculates the arcsine of a value in radians.
- `COS` calculates the cosine of a value in radians.
- `COT` calculates the cotangent of a value in radians.
- `TAN` calculates the tangent of a value in radians.
- `DEGREES` converts a value in radians to degrees.
- `PI` returns π as a value of type `DOUBLE PRECISION`.

**Introduced the `timezone` query-level setting with `time_zone` as an alias**

Added the `timezone` query-level setting. The previous `time_zone` query setting still works, and is now an alias for `timezone`.

**Introduced new `PERCENTILE_CONT` and `MEDIAN` aggregate functions**

Added the following aggregate functions:

- `PERCENTILE_CONT` calculates a specified percentile of values in an ordered dataset.
- `MEDIAN` returns the median of a given column. It is equivalent to `PERCENTILE_CONT(0.5)`: half the values in the column are smaller, and half are bigger than the returned value. If the number of values in the column is even, `MEDIAN` returns the arithmetic mean of the two middle values.

**Added support to meet HIPAA regulations for health information**

Added [support to meet federal HIPAA regulations](/Overview/Security/security.html#hipaa-compliance) to ensure the confidentiality, integrity, and availability of electronic protected health information within the Firebolt platform.

### [](#performance-improvements-6)Performance Improvements

**Improved expression comparison logic within queries**

Improved expression comparison logic to better recognize identical expressions within queries. This enhancement supports a broader range of queries and boosts the overall quality of query plans.

**Improving cold reads by reducing the amount of Amazon S3 requests needed to load data**

Improved the performance of cold reads by minimizing the number of Amazon S3 requests required to load data. In the case of tiny tablets, this improvement lead to a 50% improvement in performance.

### [](#bug-fixes-8)Bug Fixes

**Fixed a bug preventing view creation with type conversions to array types**

Fixed an issue that prevented users from creating database views that involve type conversion to array types.

## [](#db-version-47)DB version 4.7

### [](#new-features-10)New Features

**Added Snappy compression support to the COPY TO command for PARQUET output format**  
You can now apply Snappy compression, which is faster than GZIP, when using `COPY TO` with `TYPE=PARQUET`. Specify `COMPRESSION=SNAPPY` within `COPY TO` to enable this.

**Added `information_schema.engine_user_query_history` view to log only user-initiated queries**  
Added a new query history view, `information_schema.engine_user_query_history`, which shows all queries initiated by users. This view filters information from `information_schema.engine_query_history` view, which logs all engine queries including system-generated ones like UI updates and page-load requests.

**Added support for `information_schema.enabled_roles`**  
Added a new view `information_schema.enabled_roles` which lists the roles available in the account.

**Added a system setting `enable_subresult_cache` for controlling subresult reuse**  
A new system setting `enable_subresult_cache` allows users to enable or disable caching of query subresults for subsequent reuse. Caching remains enabled by default. This setting allows users to temporarily disabling caching, e.g. for benchmarking purposes.

**Added “FROM first” syntax allowing the `FROM` clause to precede the `SELECT` clause**  
Added support for the “FROM first” syntax, which allows placing the `FROM` clause before the `SELECT` clause, for example `FROM t SELECT a, SUM(b) GROUP BY a`. You can now also omit the `SELECT` clause, as in `FROM t`.

**Introduced a new function `GEN_RANDOM_UUID_TEXT` to generate a universally unique identifier (UUID)**  
The new function `GEN_RANDOM_UUID_TEXT` accepts no arguments and returns a version `4` UUID as defined by [RFC-4122](https://tools.ietf.org/html/rfc4122#section-4.4) as a `TEXT` value.

**Introduced `~` and `!~` operators as aliases for `REGEXP_LIKE` and `NOT REGEXP_LIKE`**  
Added the `~` operator as an alias for `REGEXP_LIKE`, and the `!~` operator, which serves as an alias for `NOT REGEXP_LIKE`.

**Introduced JSON functions `JSON_POINTER_EXTRACT_KEYS`, `JSON_POINTER_EXTRACT_VALUES`, `JSON_POINTER_EXTRACT_TEXT`**  
The following new JSON functions are now supported:

- `JSON_POINTER_EXTRACT_KEYS` extracts keys from a JSON object
- `JSON_POINTER_EXTRACT_VALUES` extracts values from a JSON object
- `JSON_POINTER_EXTRACT_TEXT` extracts the JSON string value as SQL TEXT

**Introduced trigonometric functions `RADIANS`, `SIN`, `ATAN2`**  
The following trigonometric functions are now supported:

- `RADIANS` to convert degrees into radians
- `SIN` to compute the sine in radians
- `ATAN2` to calculate the arctangent with two arguments. `ATAN2(y,x)` is the angle between the positive x-axis and the line from the origin to the point `(x,y)`, expressed in radians.

**Introduced new functions to calculate standard deviation and variance for both samples and populations**  
New functions that accept `REAL` and `DOUBLE` inputs and return standard deviations and variances:

- `STDDEV_SAMP` - Returns the sample standard deviation of all non-`NULL` numeric values produced by an expression, which measures how spread out values are in a sample.
- `STDDEV_POP` - Returns the population standard deviation of all non-`NULL` numeric values produced by an expression, which measures how spread out values are in an entire population.
- `VAR_SAMP` - Returns the sample variance of all non-`NULL` numeric values produced by an expression, which measures the average of the squared differences from the sample mean, indicating how spread out the values are within a sample.
- `VAR_POP` - Returns the population variance of all non-`NULL` numeric values produced by an expression. The population variance measures the average of the squared differences from the population mean, indicating how spread out the values are within the entire population.

**Introduced new array functions `ARRAY_ALL_MATCH` and `ARRAY_ANY_MATCH`**  
The new functions `ARRAY_ALL_MATCH` and `ARRAY_ANY_MATCH` accept an (optional) lambda function and an array and return `TRUE` if all elements (`ARRAY_ALL_MATCH`) or any element (`ARRAY_ANY_MATCH`) satisfy the lambda condition, and `FALSE` otherwise. When no lambda is passed, the array has to be of type `BOOLEAN`, and the identity lambda `x -> x` is used.

### [](#performance-improvements-7)Performance Improvements

**Improved performance of `JSON_EXTRACT`, `JSON_EXTRACT_ARRAY`, and `JSON_VALUE` functions**  
Enhanced the performance of the `JSON_EXTRACT`, `JSON_EXTRACT_ARRAY`, and `JSON_VALUE` functions.

### [](#behavior-changes-4)Behavior Changes

**Updated sorting method for array columns with `NULL` values to align with PostgreSQL behavior**

The sorting method for array columns containing `NULL` values has been updated to ensure that `ASC NULLS FIRST` places `NULL` values before arrays, and `DESC NULLS LAST` places `NULL` values after arrays, which aligns with PostgreSQL behavior.

The following code example creates a temporary table `tbl` which contains three rows: a `NULL` array, an array with the value `1`, and an array with a `NULL` element. Then, a `SELECT` statement sorts all rows in ascending order:

```
WITH tbl(i) AS (
  SELECT NULL::INT[]
  UNION ALL
  SELECT ARRAY[1]::INT[]
  UNION ALL
  SELECT ARRAY[NULL]::INT[]
)
SELECT * FROM tbl ORDER BY i ASC NULLS FIRST;
```

The query previously returned `{NULL}, {1}, NULL`, but now returns `NULL, {1}, {NULL}`.

`NULLS FIRST` and `NULLS LAST` apply to the array itself, not to its elements. By default, ascending order (`ASC`) assumes `NULLS LAST`, while descending order (`DESC`) assumes `NULLS FIRST` when sorting arrays.

**Allowed use of the SESSION\_USER function without parentheses**

The `SESSION_USER` function can now be used without parentheses, like this: `SELECT SESSION_USER`. As a result, any column named `session_user` now needs to be enclosed in double quotes as follows: `SELECT 1 AS "session_user"` or `SELECT "session_user" FROM table`.

### [](#bug-fixes-9)Bug Fixes

**Corrected JSON output format to display NaN values consistently as `nan`**  
The JSON output format previously showed some NaN values as `-nan`. This was corrected to consistently display NaN values as `nan` in the JSON output.

**Resolved an issue with `CHECKSUM` and `HASH_AGG` failing when combining literals and table columns**  
Fixed an issue where the `CHECKSUM` and `HASH_AGG` functions failed when used with a combination of literals and table columns.

**Fixed a rare inaccuracy that could cause incorrect results on multi-node engines when performing certain `UNION ALL` operations**  
Fixed a rare inaccuracy when performing certain `UNION ALL` operations on subqueries that are the result of aggregations or joins on overlapping but distinct keys, followed by an aggregation or join on the common keys of the subqueries’ aggregations or joins.

**Fixed a rare inaccuracy that could cause incorrect results with CTEs using `RANDOM()` in specific join scenarios**  
Fixed a rare inaccuracy that caused incorrect results when a common table expression using the `RANDOM()` function was used multiple times, and at least one of these uses was on the probe side of a join involving a primary index key of the underlying table.

## [](#db-version-46)DB version 4.6

**September 2024**

### [](#new-features-11)New Features

**`COPY TO` support for the `SNAPPY` compression type**

[COPY TO](/sql_reference/commands/data-management/copy-to.html) now supports `SNAPPY` as a new compression option for Parquet files. This enhancement offers greater flexibility for managing file size and performance, particularly for workloads requiring faster compression. Each file is written in Parquet format, with the specified compression applied to the data pages in the column chunks.

**`COPY FROM` support for filtering by source file metadata**

[COPY FROM](/sql_reference/commands/data-management/copy-from.html) now supports filtering by source file metadata using the `WHERE` clause.

**Added support for vector distance calculations with new functions**

Firebolt has added support for vector distance and similarity calculations with the following new functions: [VECTOR\_COSINE\_DISTANCE](/sql_reference/functions-reference/vector/vector-cosine-distance.html), [VECTOR\_MANHATTAN\_DISTANCE](/sql_reference/functions-reference/vector/vector-manhattan-distance.html), [VECTOR\_EUCLIDEAN\_DISTANCE](/sql_reference/functions-reference/vector/vector-euclidean-distance.html), [VECTOR\_SQUARED\_EUCLIDEAN\_DISTANCE](/sql_reference/functions-reference/vector/vector-squared-euclidean-distance.html), [VECTOR\_COSINE\_SIMILARITY](/sql_reference/functions-reference/vector/vector-cosine-similarity.html), and [VECTOR\_INNER\_PRODUCT](/sql_reference/functions-reference/vector/vector-inner-product.html).

### [](#behavior-changes-5)Behavior Changes

**Introduced `SHOW CATALOGS` statement and aliased `SHOW DATABASES` to it while deprecating `SHOW DATABASE X`**

A new statement `SHOW CATALOGS` now acts as an alias for `SHOW DATABASES`. The statement `SHOW DATABASE X` is no longer supported.

**`COPY FROM` now unzips Parquet files with gzip extensions**

Before version 4.6, the `COPY FROM` command did not apply file-level decompression to Parquet files with a `.gzip` or `.gz` extension. The command treated these files as standard Parquet files, assuming that any compression existed only within the internal Parquet format structure.

With the release of version 4.6, `COPY FROM` now processes Parquet files similarly to other formats. When a Parquet file has a `.gz` or `.gzip` extension, the command will first decompress the file before reading it as a Parquet format file. Hence, it will now fail while reading internally compressed Parquet files with gzip extensions. Users experiencing issues with loading files after this change should contact the support team at support@firebolt.io for assistance.

### [](#bug-fixes-10)Bug Fixes

**Fixed a rare bug that caused some query failures from incorrect computation of cacheable subresults**

Fixed a rare bug impacting the logic that determined which subresults could be cached and reused. This issue could have caused query failures in certain patterns, but it did not impact the accuracy of the query outcomes.

**Updated name of aggregatefunction2 to aggregatefunction in explain output**

The name `aggregatefunction2` has been updated to `aggregatefunction` in the [EXPLAIN](/sql_reference/commands/queries/explain.html) output.

**Fixed incorrect results in `ARRAY_AGG` expressions by excluding `NULL` values for false conditions in aggregating indexes**

Aggregate expressions like `ARRAY_AGG(CASE WHEN <cond> THEN <column> ELSE NULL END)` previously returned incorrect results by excluding `NULL` values for rows when the condition was `FALSE`.

## [](#db-version-45)DB version 4.5

**September 2024**

### [](#new-features-12)New Features

**Allowed casting from `TEXT` to `DATE` with truncation of timestamp-related fields** Casting from `TEXT` to `DATE` now supports text values containing fields related to timestamps. These fields are accepted, but truncated during conversion to `DATE`.

The following code example casts the `TEXT` representation of the timestamp `2024-08-07 12:34:56.789` to the `DATE` data type. The conversion truncates the time portion, leaving only the date, as follows:

Example:

```
SELECT '2024-08-07 12:34:56.789'::DATE`
```

Results in

```
DATE `2024-08-07`
```

**Added the `CONVERT_FROM` function**

Added the `CONVERT_FROM` function that converts a `BYTEA` value with a given encoding to a `TEXT` value encoded in UTF-8.

**Added the BITWISE aggregate functions**

Added support for the following functions: BIT\_OR (bitwise OR), BIT\_XOR (bitwise exclusive OR), and BIT\_AND (bitwise AND).

**Added the `REGEXP_LIKE_ANY` function**

Added the `REGEXP_LIKE_ANY` function that checks whether a given string matches any regular expression pattern from a specified list of patterns.

### [](#bug-fixes-11)Bug Fixes

**Updated `created` and `last_altered` column types in `information_schema.views` from `TIMESTAMP` to `TIMESTAMPTZ`** The data types of the `created` and `last_altered` columns in `information_schema.views` have been changed from `TIMESTAMP` to `TIMESTAMPTZ`.

**Fixed runtime constant handling in the sort operator** Fixed the handling of runtime constants in the sort operator. Now, the sort operator can be correctly combined with `GENERATE_SERIES`. For example, the query `SELECT x, GENERATE_SERIES(1,7,3) FROM GENERATE_SERIES(1,3) t(x)` now correctly displays values `1` to `3` in the first column, instead of just `1`.

## [](#db-version-44)DB version 4.4

**August 2024**

### [](#new-features-13)New Features

**Extended support for date arithmetic**

Now you can subtract two dates to get the number of elapsed days. For example, `DATE '2023-03-03' - DATE '1996-09-03'` produces `9677`.

**Role-based permissions for COPY FROM and external tables**

Added support for role-based permissions (ARNs) to the COPY FROM command and external table operations.

**Added `trust_policy_role` column to `information_schema.accounts` view for S3 access**

Added a new column `trust_policy_role` to the `information_schema.accounts` view. This column shows the role used by Firebolt to access customer S3 buckets.

**Enabled selection of external tables’ pseudo columns without adding data columns**

Users can now select an external table’s pseudo columns (source file name, timestamp, size, and etag) without adding any data columns. For example, `select $source_file_timestamp from t_external` returns the file timestamps for each row. The query `select count($source_file_timestamp) from t_external` returns the total number of rows in the external table, similar to `count(*)`. The query `select count(distinct $source_file_name) from t_external` returns the number of distinct objects containing at least one row in the source S3 location. Regarding `count(*)` performance, formats like CSV or JSON still require reading the data fully to determine an external file’s row count. However, Parquet files provide the row count as part of the file header, and this is now used instead of reading the full data.

**Extended support for arbitrary join conditions, including multiple inequality predicates**

We now support more join conditions. As long as there is one equality predicate comparing a left column to a right column of the join (not part of an OR expression), the remaining join condition can now be an arbitrary expression. The limitation on the number of inequality predicates was removed.

**New functions `URL_ENCODE` and `URL_DECODE`**

We added support for the `URL_ENCODE` and `URL_DECODE` functions.

**New logarithm functions `ln`, `log`**

We added support for calculating logarithms. The natural logarithm is available using `ln(val double precision)`. The base 10 logarithm is available using `log(val double precision)`. Logarithms with custom bases are available using `log(base double precision, val double precision)`.

**New function \`SQRT**

Added support for the `SQRT` function to compute the square root.

**New functions `JSON_VALUE`, `JSON_VALUE_ARRAY`, `JSON_EXTRACT_ARRAY`**

Added support for the functions `JSON_VALUE`, `JSON_VALUE_ARRAY`, and `JSON_EXTRACT_ARRAY`.

**New function `SESSION_USER`**

Support has been added for the `SESSION_USER` function, which retrieves the current user name.

**New columns in `information_schema.engine_query_history`**

Added two new columns to `information_schema.engine_query_history`: `query_text_normalized_hash` and `query_text_normalized`.

### [](#breaking-changes)Breaking Changes

**Reserved the keyword GEOGRAPHY, requiring double quotes for use as an identifier**

The word GEOGRAPHY is now a reserved keyword and must be quoted using double quotes for use as an identifier. For example, `create table geography(geography int);` will now fail, but `create table "geography" ("geography" int);` will succeed.

**Deprecated the legacy HTTP ClickHouse headers**

We no longer accept or return the legacy HTTP ClickHouse header format `X-ClickHouse-*`.

**Fixed `json_value` zero-byte handling**

The `json_value` function no longer returns null characters (0x00), as the TEXT datatype does not support them. For example, `select json_value('"\u0000"');` now results in an error.

**Change default values for NODES and TYPE during CREATE ENGINE**

When performing a CREATE ENGINE, the default values for NODES and TYPE parameters have changed. NODES defaults to `2` (previously `1`) and TYPE defaults to `M` (previously `S`). To create an engine with the previous default values, run the following command:

```
CREATE ENGINE my_engine WITH NODES=1 TYPE=S
```

### [](#bug-fixes-12)Bug Fixes

**Fixed directory structure duplication in the S3 path when using the COPY TO statement with SINGLE\_FILE set to FALSE**

Fixed an issue in `COPY TO` when `SINGLE_FILE=FALSE`. Previously, the specified directory structure in the location was repeated twice in the S3 path. For example, files were output to “s3://my-bucket/out/path/out/path/” instead of “s3://my-bucket/out/path/”.

**Fixed the file extension in the S3 path when using the COPY TO statement with GZIP-Parquet format**

Fixed an issue in `COPY TO` when `TYPE=PARQUET` and `COMPRESSION=GZIP`, which uses the Parquet file format with internal GZIP compression for the columns. Previously, the output files would have the extension “.parquet.gz”. Now, the extension is “.gz.parquet”.

## [](#db-version-43)DB version 4.3

**August 2024**

### [](#new-features-14)New Features

**Role-based permissions for COPY FROM and External Table processes**

Enabled role-based permissions for COPY FROM and External Table processes.

**HLL-based count distinct functions compatible with the Apache DataSketches library**

Firebolt now supports count-distinct functions using the HLL (HyperLogLog) algorithm, compatible with the Apache DataSketches library. For details and examples, see documentation on the functions [APACHE\_DATASKETCHES\_HLL\_BUILD](/sql_reference/functions-reference/datasketches/apache-datasketches-hll-build.html), [APACHE\_DATASKETCHES\_HLL\_MERGE](/sql_reference/functions-reference/datasketches/apache-datasketches-hll-merge.html), and [APACHE\_DATASKETCHES\_HLL\_ESTIMATE](/sql_reference/functions-reference/datasketches/apache-datasketches-hll-estimate.html).

**Supported additional join conditions and removed the restriction on the number of inequality predicates**

Firebolt has added enhanced support for more join conditions. As long as there is one equality predicate comparing a left column to a right column of the join, which is not part of a disjunctive (OR) expression, the remaining join condition can be arbitrary. The previous limitation on the number of inequality predicates has been removed.

### [](#performance-improvements-8)Performance Improvements

**Multi-node query performance**

Firebolt has improved the performance of data transfer between nodes, resulting in faster overall query execution times.

**Enhanced Interval Arithmetic Support**

Firebolt has enhanced support for interval arithmetic. You can now use expressions of the form `date_time + INTERVAL * d`, where `date_time` is a expression of type Date, Timestamp, TimestampTz, and `d` is an expression of type DOUBLE PRECISION. The interval is now scaled by `d` before being added to `date_time`. For example, writing `INTERVAL '1 day' * 3` is equivalent to writing `INTERVAL '3 days'`.

**Optimized selective inner and right joins on primary index and partition by columns to reduce rows scanned**

Selective inner and right joins on primary index and partition by columns now can now benefit from pruning. This reduces the number of rows scanned by filtering out rows that are not part of the join result early in the process. This optimization works best when joining on the first primary index column or a partition by column. The optimization is applied automatically when applicable, and no action is required. Queries that used this optimization will display “Prune:” labels on the table scan in the EXPLAIN (PHYSICAL) or EXPLAIN (ANALYZE) output.

### [](#bug-fixes-13)Bug Fixes

**Fixed a bug in the combination of cross join and the `index_of` function**

Resolved an issue where the `index_of` function would fail when applied to the result of a cross join that produced a single row.

### [](#breaking-changes-1)Breaking Changes

**Temporarily restricted column DEFAULT expressions in CREATE TABLE statements**

Column DEFAULT expressions in CREATE TABLE statements have been temporarily restricted, they can only consist of literals and the following functions: `CURRENT_DATE()`, `LOCALTIMESTAMP()`, `CURRENT_TIMESTAMP()`, `NOW()`. Existing tables with column DEFAULT expressions are not affected.

**Underflow detection while casting from TEXT to floating point data types**

Firebolt now detects underflow, a condition where a numeric value becomes smaller than the minimum limit that a data type can represent, when casting from TEXT to floating point data types. For example, the query `select '10e-70'::float4;` now returns an error, while it previously returned `0.0`.

**Returning query execution errors in JSON format through the HTTP API**

Firebolt’s HTTP API now returns query execution errors in JSON format, allowing for future enhancements like including metadata such as error codes, or the location of a failing expression within the SQL script.

**Changed default of case\_sensitive\_column\_mapping parameter in COPY FROM**

The default value for the `CASE_SENSITIVE_COLUMN_MAPPING` parameter in `COPY FROM` is now `FALSE`, meaning that if a target table contains column names in uppercase and the source file to ingest has the same columns in lowercase, the ingestion will consider them the same column and ingest the data.

**`extract` function returns Numeric(38,9) for Epoch, second, and millisecond extraction**

The result data type of the `extract` function for epoch, second, and millisecond was changed to return the type Numeric(38,9) instead of a narrower Numeric type. For example, `select extract(second from '2024-04-22 07:10:20'::timestamp);` now returns Numeric(38,9) instead of Numeric(8,6).

## [](#db-version-42)DB version 4.2

**July 2024**

### [](#new-features-15)New features

**New `ntile` window function**

Firebolt now supports the `ntile` window function. Refer to our [NTILE](/sql_reference/functions-reference/window/ntile.html) documentation for examples and usage.

### [](#enhancements-changes-and-new-integrations)Enhancements, changes and new integrations

**Improved query performance**

Queries with “`SELECT [project_list] FROM [table] LIMIT [limit]`” on large tables are now significantly faster.

**Updated table level RBAC**

Table level RBAC is now supported by Firebolt. This means that RBAC checks also cover schemas, tables, views and aggregating indexes. Refer to our [RBAC](/Guides/security/rbac.html) docs for a detailed overview of this new feature. The new Firebolt version inhibits the following change:

- System built-in roles are promoted to contain table level RBAC information. This means that new privileges are added to `account_admin`, `system_admin` and `public` roles. The effect is transparent— any user assigned with those roles will not be affected.

**Removal of Deprecated Columns from `INFORMATION_SCHEMA.ENGINES`**

We removed the following columns from `INFORMATION_SCHEMA.ENGINES` that were only for FB 1.0 compatibility: `region`, `spec`, `scale`, `warmup`, and `attached_to`. These columns were always empty. (These columns are hidden and do not appear in `SELECT *` queries, but they will still work if referenced explicitly.)

### [](#breaking-changes-2)Breaking Changes

**Improved rounding precision for floating point to integer casting**

Casting from floating point to integers now uses Banker’s Rounding, matching PostgreSQL’s behavior. This means that numbers that are equidistant from the two nearest integers are rounded to the nearest even integer:

Examples:

```
SELECT 0.5::real::int
```

This returns 0.

```
SELECT 1.5::real::int
```

This returns 2.

Rounding behavior has not changed for numbers that are strictly closer to one integer than to all others.

**JSON functions update**

Removed support for `json_extract_raw`, `json_extract_array_raw`, `json_extract_values`, and `json_extract_keys`. Updated `json_extract` function: the third argument is now `path_syntax`, which is a JSON pointer expression. See [JSON\_EXTRACT](/sql_reference/functions-reference/JSON/json-extract.html) for examples and usage.

**Cluster ordinal update**

Replaced `engine_cluster` with [cluster\_ordinal](/sql_reference/information-schema/engine-metrics-history.html) in `information_schema.engine_metrics_history`. The new column is an integer representing the cluster number.

**Configurable cancellation behavior on connection drop**

Introduced the `cancel_query_on_connection_drop` setting, allowing clients to control query cancellation on HTTP connection drop. Options include `NONE`, `ALL`, and `TYPE_DEPENDENT`. Refer to [system settings](/Reference/system-settings.html#query-cancellation-mode-on-connection-drop) for examples and usage.

**JSON format as default for error output**

The HTTP API now returns query execution errors in JSON format by default. This change allows for the inclusion of meta information such as error codes and the location of failing expressions in SQL scripts.

**STOP ENGINE will drain currently running queries first**

`STOP ENGINE` command now supports graceful drain, meaning any currently running queries will be run to completion. Once all the queries are completed, the engine will be fully stopped and terminated. If you want to stop the engine immediately, you can issue a STOP ENGINE command use the TERMINATE option. For example, to immediately stop an engine, my\_engine, you can use:

```
 STOP ENGINE myEngine WITH TERMINATE = TRUE
```

**Scaling engines will not terminate currently running queries**

`ALTER ENGINE` command now supports graceful drain, meaning when you scale an engine (vertically or horizontally), any currently running queries will not be terminated. New queries after the scaling operation will be directed to a new cluster, while queries running on the old cluster will be run to completion.

**Updated RBAC ownership management**

We have introduced several updates to role and privilege management:

- The `security_admin` role will be removed temporarily and re-introduced in a later release.
- `Information_object_privileges` includes more privileges. Switching to to a specific user database (e.g by executing `use database db`) will only show privileges relevant for that database. Account-level privileges no longer show up when attached to a specific database.
- Every newly created user is granted with a `public` role. This grant can be revoked.

## [](#db-version-41)DB version 4.1

**June 2024**

- [Resolved issues](#resolved-issues)

### [](#resolved-issues)Resolved issues

- Fixed an issue causing errors when using `WHERE column IN (...)` filters on external table scans.

## [](#db-version-40)DB version 4.0

**June 2024**

- [Enhancements, changes, and new integrations](#enhancements-changes-and-new-integrations)
- [Breaking Changes](#breaking-changes)

### [](#enhancements-changes-and-new-integrations-1)Enhancements, changes and new integrations

**Query Cancelation on HTTP Connection Drop**

Going forward, when the network connection between the client and Firebolt is dropped (for example because the Firebolt UI tab was closed or due to network issues), DML queries (INSERT, UPDATE, DELETE, etc) are no longer canceled automatically, but will keep running in the background. You can continue to monitor their progress in `information_schema.engine_running_queries` or cancel them manually using the `cancel query` statement if desired. DQL queries (SELECT) are still canceled automatically on connection drop.

**New Aggregate Functions: `CHECKSUM` and `hash_agg`**

`CHECKSUM` and `hash_agg` functions are now supported for aggregating indexes. Note that when the `hash_agg` function doesn’t receive rows, the result is 0.

### [](#breaking-changes-3)Breaking Changes

**Array Casting Nullability Update**

Cast to array will no longer support specifying nullability of the inner type. Example:

```
a::array(int null)
```

or

```
cast(a as array(int not null)) 
```

will now fail, and need to be rewritten as:

```
a::array(int) 
```

or

```
cast(a as array(int)). 
```

**Postgres-compliant Cast**

Casts now behave the same across the product and adhere to the list of supported casts. Some usages of casts (explicit, implicit, or assignment cast) that were previously allowed are no longer supported and now result in errors. For more details on list of supported casts, see the documentation [here](/sql_reference/data-types.html#type-conversion).

## [](#db-version-334)DB version 3.34

**May 2024**

- [Enhancements, changes, and new integrations](#enhancements-changes-and-new-integrations)
- [Resolved issues](#resolved-issues)

### [](#enhancements-changes-and-new-integrations-2)Enhancements, changes and new integrations

**Removed `MATCH` function**

The `match` function has been removed and replaced with [regexp\_like](/sql_reference/functions-reference/string/regexp-like.html).

**Producing an error for array function failure instead of NULL**

Array function queries that accept two or more array arguments now produce an error. If you call an array function such as `array_transform(..)` or `array_sort(..)` with multiple array arguments, the arrays must have the same size. For example:

```
array_transform(x, y -> x + y, arr1, arr2)
```

This raises an error if `array_length(arr1) != array_length(arr2)`. We now also perform this check for NULL literals. If you previously used `array_transform(x, y -> x + y, NULL::INT[], Array[5, 6])`, you got back `NULL`. Now, the query using that expression will raise an error.

**Added ARRAY\_FIRST function**

The [array\_first](../../sql_reference/functions-reference/Lambda/array-first.html) function has been added. It returns the first element in the given array for which the given function returns `true`.

**New name for `any_match`**

A new name for `any_match` has been added: [array\_any\_match](/sql_reference/functions-reference/Lambda/array-any-match.html). `any_match` will be kept as an alias.

**Updated ARRAY\_SUM return types**

The `array_sum` function of `bigint[]` now returns a numeric value and `array_sum` of `real[]` returns a real value.

**Precedence of operators**

Breaking change in operator precedence between comparison operators such as `=`, `<`, `>`, and `IS` operator. New behavior is compatible with Postgres.

Examples of query that changed behavior:

```
select 1 is null = 2 is null
```

This used to be `true`, because it was interpreted as `select (1 is null) = (2 is null)`. It now becomes an error of incompatible types in `=`

```
select false = false is not null
```

The result used to be `false` - `select false = (false is not null)`, but now is `true` - `select (false = false) is not null`.

**Dropping the role**

Role cannot be dropped if there are permissions granted to the role. The error message will be displayed if you need to manually drop permissions associated to the role.

**Coalesce Short-Circuiting**

`COALESCE` now supports short-circuiting in Firebolt. Queries such as `COALESCE(a, 1 / 0) FROM t` could fail before, even when there were no NULLs in t. Only `CASE WHEN` supported short circuiting. Firebolt is now aligned with PostgreSQL and supports short circuiting in `COALESCE` as well.

**Create table under I\_S schema**

You can now execute `CREATE TABLE`/`VIEW`/`AGGREGATING INDEX` only under the public schema.

**Improved error message for JSON `PARSE_AS_TEXT` format**

The error message for external tables created with JSON `PARSE_AS_TEXT` format has been revised. This format reads specifically into a *single* column of type either TEXT or `TEXT NOT NULL`. (Note there may be external table partition columns defined after the single TEXT column, and they are okay). Now, only the error message regarding the `CREATE EXTERNAL TABLE` statement on a user’s first attempt to use `SELECT` will be seen. Support for reading format JSON `PARSE_AS_TEXT=TRUE` into a `TEXT NOT NULL` column has been added.

**Implemented column\_mismatch**

Support for `ALLOW_COLUMN_MISMATCH` in `COPY INTO` has been added.

**Corrected NULL behavior of `STRING_TO_ARRAY`**

The behavior of `string_to_array` now matches its behavior in PostgreSQL. The change affects NULL delimiters where the string is split into individual characters, as well as empty strings and where the output is now an empty array.

**Changed city\_hash behavior for nullable inputs**

The behavior for `city_hash` has changed for nullable inputs. For example:

```
SELECT CITY_HASH([null]) = CITY_HASH([''])
```

This is now false.

**Function `ARRAY_AGG` now preserves NULLS**

The `array_agg` function has been changed to return PostgreSQL-compliant results:

- `array_agg` now preserves `NULL` values in its input, e.g. `select array_agg(x) from unnest(array [1,NULL,2] x)` returns `{1,NULL,2}`
- `array_agg` now returns `NULL` instead of an empty array if there are no input values

**Lambda parameters are no longer supported by `array_sum`**

Array aggregate functions no longer support lambda parameters. To get the old behavior for conditional lambda functions, use transform instead. For example:

```
array_sum(transform(...))
```

**Explicit Parquet conversion from DATE to INT is now needed**

A breaking change has been implemented in raising an error on reading a Parquet/ORC `DATE`/`TIMESTAMP` column if the `EXTERNAL TABLE` expects the column to have type `INT`/`BIGINT`. `DATE`/`TIMESTAMP` cannot be cast to `INT`/`BIGINT`, and external table scans will no longer allow this cast either. You need to explicitly transform the Parquet/ORC `DATE`/`TIMESTAMP` column with `EXTRACT`(`EPOCH FROM` col) to insert it into an `INT`/`BIGINT` column.

### [](#resolved-issues-1)Resolved issues

- Fixed a bug where negation did not check for overflows correctly.

## [](#db-version-333)DB version 3.33

**April 2024**

- [Enhancements, changes, and new integrations](#enhancements-changes-and-new-integrations)
- [Resolved issues](#resolved-issues)

### [](#enhancements-changes-and-new-integrations-3)Enhancements, changes and new integrations

**Removed ‘element\_at’ Function**

The `element_at` function for arrays has been removed and replaced with the `[]` operator.

**Change of return type from BIGINT to INTEGER**

The `index_of`/`array_position` function now returns INTEGER instead of BIGINT.

**Removed LIMIT DISTINCT syntax**

The `LIMIT_DISTINCT` syntax is no longer supported by Firebolt.

**Updated CAST function behavior**

All cast logic has been moved to runtime in Firebolt. The `castColumn` function is now replaced by `fbCastColumn`, ensuring consistent casting behavior and resolving issues with the `COPY FROM` operation and other cast calls. Uses of implicit/explicit `CAST` may result in errors due to this fix.

New breaking change.

### [](#resolved-issues-2)Resolved issues

- Fixed a bug in `array_position` where searching for `NULL` in an array with non-null elements incorrectly returned a match in some cases.

## [](#db-version-332)DB version 3.32

**April 2024**

- [New features](#new-features)
- [Enhancements, changes, and new integrations](#enhancements-changes-and-new-integrations)
- [Resolved issues](#resolved-issues)

### [](#new-features-16)New features

**Expose and document ‘typeof’ as a toTypeName function**

The `typeof` function has been added, which returns the data type of a SQL expression as a string.

### [](#enhancements-changes-and-new-integrations-4)Enhancements, changes and new integrations

**Spilling Aggregations**

Firebolt can now process most aggregations that exceed the available main memory of the engine by spilling to the SSD cache when needed. This happens transparently to the user. A query that made use of this capability will populate the `spilled_bytes` column in `information_schema.query_history`. Spilling does not support aggregations where a single group exceeds the available memory (e.g., `select count(distinct high_cardinality_column) from huge_table`) and may not yet work reliably for all aggregate functions or engine specs. We will continue improving the feature in upcoming releases.

**No overflow detected in cast from FLOAT to DECIMAL**

Fix results of casting from `float32` to decimals with precision &gt; 18. In addition to the correct results breaking change, there are certain queries that was working before that now will fail involving overflow.

Example query:

- `SELECT` 17014118346046923173168730371588410572::REAL::DECIMAL(37,0).

Previously, this was working and returned a wrong result, but now it will fail with an overflow error.

**ARRAY\_COUNT returns 0 instead of NULL**

`ARRAY_COUNT` on `NULL` array now returns `0` instead of `NULL`.

**No overflow check in arithmetic operations**

Arithmetic operators (i.e. multiplication, addition, subtraction, and division) now perform correct overflow checking. This means that queries that used to return wrong results in the past now throw runtime errors.

Example queries:

- `SELECT` 4294967296 * 4294967296 -&gt; now throws an error, before it would return 0
- `SELECT` 9223372036854775807 + 9223372036854775807 -&gt; now throws an error, before it would return -2
- `SELECT` (a + b) * c -&gt; this might throw runtime errors if there are large values in the column, but this is highly data dependent.

**Implement bool\_or/bool\_and aggregation functions**

New aggregate functions bool\_or and bool\_and have been added.

**Remove old deprecate REGENERATE AGGREGATING INDEX**

‘REGENERATE AGGREGATING INDEX’ syntax has now been removed.

**Align the syntax of our “escape” string literals with PostgreSQL**

Escape [string literals](/sql_reference/data-types.html) now support octal and Unicode escape sequences. As a result, escape string literals now behave exactly like PostgreSQL. Example: `SELECT E'\U0001F525b\x6F\154t';` returns `🔥bolt`. If the setting `standard_conforming_strings` is not enabled for you, regular string literals (e.g., `SELECT 'foo';`) will also recognize the new escape sequences. However, we recommend exclusively using escape string literals for using escape sequences. Please be aware that you will get different results if you previously used (escape) string literals containing the syntax we now use for Unicode and octal escape sequences.

**Change return value of length and octet\_length to INT**

Length and array\_length now return INTEGER instead of BIGINT.

**Subqueries in the GROUP BY/HAVING/ORDER BY clauses change**

Subqueries in `GROUP BY/HAVING/ORDER BY` can no longer references columns from the selection list of the outer query via their aliases as per PG compliance. `select 1 + 1 as a order by (select a);` used to work, but now fails with `unresolved name a` error.

**Bytea serialization to CSV fix**

Changed Bytea to CSV export: from escaped to non escaped.

Example:

- `COPY` (select ‘a’::bytea) to ‘s3…’; the results will now be “\\x61” instead of “\\x61”.

### [](#resolved-issues-3)Resolved issues

- Fixed results of casting literal float to numeric. In the past the float literal was casted to float first then to numeric, this caused us to lose precision.

Examples:

- `SELECT` 5000000000000000000000000000000000000.0::DECIMAL(38,1); -&gt; 5000000000000000000000000000000000000.0
- `SELECT` (5000000000000000000000000000000000000.0::DECIMAL(38,1)+5000000000000000000000000000000000000.0::DECIMAL(38 1)); -&gt; ERROR: overflow.

Note that before, it was not an error and resulted in: 9999999999999999775261218463046128332.8.

- Fixed a longstanding bug with &gt;= comparison on external table source\_file\_name. Whereas this would previously have scraped fewer files than expected off the remote S3 bucket, you will now get all files properly (lexicographically) compared against the input predicate.

<!--THE END-->

- Fixed a bug when `USAGE ANY ENGINE` (and similar) privileges were shown for * account. Now it is being show for current account.

<!--THE END-->

- Fixed a bug involving [‘btrim’](/sql_reference/functions-reference/string/btrim.html) string characters, where invoking `btrim`, `ltrim`, `rtrim`, or `trim` with a literal string but non-literal trim characters could result in an error.

## [](#db-version-331)DB version 3.31

**March 2024**

- [New features](#new-features)
- [Enhancements, changes, and new integrations](#enhancements-changes-and-new-integrations)
- [Resolved issues](#resolved-issues)

### [](#new-features-17)New features

**PG compliant division**

LQP2 has a new division operator that is PG compliant, by default.

**Prevents usage of new line delimeter for schema inference**

An error will now occur if schema inference is used with the option “delimiter” set to something other than the default.

### [](#enhancements-changes-and-new-integrations-5)Enhancements, changes and new integrations

**Simplified table protobuf representation**

Unique constraints in tables will be blocked for new accounts.

**Support for nullable arrays**

Support has been added to allow the [ARRAY\_ANY\_MATCH](/sql_reference/functions-reference/Lambda/array-any-match.html) lambda function to work with nullable arrays.

**Updated AWS billing error message**

The error message for an AWS billing issue on Engine Start was on Engine Start was changed to add more information and clarity.

**New requirements updated for EXPLAIN**

For `EXPLAIN` queries, we now allow only one of the following options at the same time: `ALL`, `LOGICAL`, `PHYSICAL`, `ANALYZE`.`EXPLAIN (ALL)` now returns the plans in multiple rows instead of multiple columns.

**Disabled Unix Time Functions**

The following functions are not supported anymore: ‘from\_unixtime’ ‘to\_unix\_timestamp’ ‘to\_unix\_time’

**Renamed spilled metrics columns**

The columns `spilled_bytes_uncompressed` and `spilled_bytes_compressed` of `information_schema.query_history` have been replaced by a single column `spilled_bytes`. It contains the amount of data that was spilled to disk temporarily while executing the query.

**Aggregating index placement**

Aggregating index is now placed in the same namespace as tables and views.

**Syntax and planner support for LATERAL scoping**

[LATERAL](/Reference/reserved-words.html) is now a reserved keyword. It must now be used within double-quotes when using it as an object identifier.

### [](#resolved-issues-4)Resolved issues

Changed return for division by 0 from null to fail.

Updated error log for upload failure for clarity.

Fixed a bug in ‘unnest’ table function that occurred when not all of the ‘unnest’ columns were projected.

Changed the behavior of [split\_part](/sql_reference/functions-reference/string/split-part.html) when an empty string is used as delimiter.

Fixed a bug where floating point values `-0.0` and `+0.0`, as well as `-nan` and `+nan` were not considered equal in distributed queries.

TRY\_CAST from TEXT to NUMERIC now works as expected: if the value cannot be parsed as NUMERIC it produces null.

## [](#db-version-330)DB version 3.30

**November 2023**

- [New features](#new-features)
- [Enhancements, changes, and new integrations](#enhancements-changes-and-new-integrations)
- [Resolved issues](#resolved-issues)

### [](#new-features-18)New features

**New comparison operators**

[New comparison operators](/sql_reference/operators.html) `IS DISTINCT FROM` and `IS NOT DISTINCT FROM` have been added.

### [](#enhancements-changes-and-new-integrations-6)Enhancements, changes and new integrations

**Support for nullable arrays**

Support has been added to allow the ANY\_MATCH lambda function to work with nullable arrays

### [](#resolved-issues-5)Resolved issues

- Indirectly granted privileges have been removed from the `information_schema.object_privileges` view.
- Fixed an issue where `ARRAY_FIRST` and `ARRAY_FIRST_INDEX` returned an error if the given input was nullable.

## [](#db-version-329)DB version 3.29

**October 2023**

- [New features](#new-features)
- [Enhancements, changes, and new integrations](#enhancements-changes-and-new-integrations)

### [](#new-features-19)New features

**EXPLAIN ANALYZE now available for detailed query metrics**

You can now use the [EXPLAIN command](/sql_reference/commands/queries/explain.html) to execute `EXPLAIN (ANALYZE) <select statement>` and get detailed metrics about how much time is spent on each operator in the query plan, and how much data is processed. The query plan shown there is the physical query plan, which you can inspect using `EXPLAIN (PHYSICAL) <select statement>` without executing the query. It shows how query processing is distributed over the nodes of an engine.

### [](#enhancements-changes-and-new-integrations-7)Enhancements, changes and new integrations

**Virtual column ‘source\_file\_timestamp’ uses new data type**

The virtual column `source_file_timestamp` has been migrated from the data type `TIMESTAMP` (legacy timestamp type without time zone) to the type `TIMESTAMPTZ` (new timestamp type with time zone).

Despite the increased resolution, the data is still in second precision as AWS S3 provides them only as unix seconds.

Use `source_file_timestamp - NOW()` instead of `DATE_DIFF('second', source_file_timestamp, NOW())`

**New function added**

A new alias [ARRAY\_TO\_STRING](/sql_reference/functions-reference/array/array-to-string.html) has been added to function `ARRAY_JOIN`.

## [](#db-version-328)DB version 3.28

**September 2023**

- [Resolved issues](#resolved-issues)

### [](#resolved-issues-6)Resolved issues

- `IN` expressions with scalar arguments now return Postgres-compliant results if there are `NULL`s in the `IN` list.
- information\_schema.running\_queries returns ID of a user that issued the running query, not the current user.
- Update error message to explain upper case behavior