# Firebolt Foundational Knowledge Layer

This document provides a foundational overview of Firebolt, designed to help an LLM learn its core concepts and capabilities effectively and responsibly. The focus is on *how to learn* Firebolt, leveraging its introspection capabilities and official documentation, rather than memorizing static details.

**Model Behavior Guidelines:**

* If you lack sufficient information to answer a user's question with high confidence, state: *"Please consult the official Firebolt documentation at https://docs.firebolt.io for more details."*
* **Never** hallucinate function names, table structures, SQL syntax, or other technical details. When unsure, refer to the official documentation or indicate the need for it.

---

## 1. Firebolt Architecture Overview

Firebolt is a cloud data warehouse optimized for high performance, low latency, and scalability, particularly for data-intensive applications.

* **Key Differentiator:** Efficiency and low cost at scale.
* **Designed For:** Addressing high latency, poor concurrency performance, and scaling challenges in data processing and retrieval.
* **Core Principles:** High efficiency (price-to-performance), concurrency at scale, elasticity, and SQL simplicity.

*(To understand the detailed architecture, consult the Firebolt documentation sections on Architecture Overview.)*

## 2. Decoupled Storage and Compute

Firebolt's architecture separates storage from compute resources.

* **Storage:** Managed via cloud storage (Amazon S3 supported), optimized for large datasets using techniques like compression and sparse indexing. Data is organized within **Databases** (also referred to as Catalogs), which logically contain schemas, tables, views, and indexes.
* **Compute (Engines):** Independent compute clusters (Engines) process queries. Engines can be started, stopped, and resized dynamically.
  * **Database Access:** Any active engine can access data from any database within the same account, highlighting the decoupling. This allows different engines (perhaps sized differently for specific workloads) to operate on the same underlying data.
  * **Multi-dimensional Elasticity:** Engines can scale independently:
    * **Vertically:** Increase node capacity (e.g., S, M, L, XL node types).
    * **Horizontally:** Increase the number of nodes within a cluster.
    * **Concurrently:** Increase the number of clusters within an engine (auto-scaling is supported in Enterprise editions).
  * **Workload Isolation:** Compute resources are isolated per engine, preventing interference between different workloads (e.g., ELT vs. BI) even when accessing the same database.

*(To learn more about engine configuration and operation, query `information_schema.engines` or refer to the Engine Fundamentals and Operate Engines guides in the docs. To list databases, query `information_schema.databases` or `information_schema.catalogs`.)*

## 3. Organizations, Accounts, Logins & Service Accounts, Users

Firebolt uses a hierarchical model for managing access and resources.

* **Organization:** The top-level entity, typically representing a company. Organization names are globally unique and based on the domain name used during registration. Manages billing and global settings like Logins and Service Accounts.
* **Account:** A container within an Organization, often used to isolate environments (dev, staging, prod) or departments. Each account exists in a specific AWS region. Holds regional resources like Databases, Engines, Users, and Roles.
  *(To list accounts, query `information_schema.accounts`. To manage accounts, use `CREATE ACCOUNT`, `ALTER ACCOUNT`, `DROP ACCOUNT`.)*
* **Login:** Represents a human person, identified by an email address. Used for authentication via credentials or SSO. Managed at the Organization level. Must be associated with a User at the Account level.
  *(To list logins, query `information_schema.logins`. To manage logins, use `CREATE LOGIN`, `ALTER LOGIN`, `DROP LOGIN`.)*
* **Service Account:** Represents a machine or application for programmatic access. Uses an ID and Secret for authentication. Managed at the Organization level. Must be associated with a User at the Account level.
  *(To list service accounts, query `information_schema.service_accounts`. To manage service accounts, use `CREATE SERVICE ACCOUNT`, `ALTER SERVICE ACCOUNT`, `DROP SERVICE ACCOUNT`. Use `CALL fb_GENERATESERVICEACCOUNTKEY` to manage secrets.)*
* **User:** An identity within a specific Account that interacts with Firebolt. Must be linked to **either** a Login or a Service Account. Roles are assigned to Users to grant permissions. Managed at the Account level.
  *(To list users, query `information_schema.users`. To manage users, use `CREATE USER`, `ALTER USER`, `DROP USER`.)*

## 4. Security and Access Control

Firebolt employs a multi-layered security approach.

* **Network Security:**
  * End-to-end encryption using TLS 1.2.
  * Network Policies allow defining allowed/blocked IP ranges at the organization, login, or service account level.
    *(To manage network policies, use `CREATE NETWORK POLICY`, `ALTER NETWORK POLICY`, `DROP NETWORK POLICY`. Query `information_schema.network_policies` to view.)*
  * Supports AWS PrivateLink for secure connections without exposing traffic to the public internet.
* **Identity Management:**
  * Uses Auth0 for identity management.
  * Supports username/password authentication, Single Sign-On (SSO), and Multi-Factor Authentication (MFA).
* **Access Control (RBAC):**
  * Permissions are managed via Roles assigned to Users.
  * Follows the principle of least privilege.
  * Includes system roles (`public`, `account_admin`, `organization_admin`) and allows creation of custom roles.
  * Permissions are hierarchical (e.g., database USAGE grants access to objects within).
  * Supports object Ownership, where the creator gains full control.
    *(Use `GRANT` and `REVOKE` to manage permissions. Query `information_schema.object_privileges`, `information_schema.enabled_roles`, `information_schema.applicable_roles` to view.)*
* **Data Protection:**
  * Data at rest is encrypted using Amazon S3-managed keys or AWS KMS keys.
  * Data in motion is encrypted end-to-end (TLS 1.2).
  * Supports HIPAA compliance.

## 5. Data Modeling and Performance Optimization

Firebolt stores data in a compressed columnar format within internal objects called tablets. Effective data modeling is crucial for performance.

* **Table Types:**
  * **Fact Tables:** Store large volumes of quantitative/event data. Distributed across engine nodes. (Default type for `CREATE TABLE`)
  * **Dimension Tables:** Store descriptive attributes. Often smaller and replicated across engine nodes for faster joins. Use `CREATE DIMENSION TABLE`.
  * **External Tables:** Reference data in S3 without loading it into Firebolt. Used primarily for ingestion. Use `CREATE EXTERNAL TABLE`.
* **Partitioning:** Tables can be partitioned (e.g., by date) to improve query performance by allowing the engine to scan only relevant partitions (partition pruning). Use `PARTITION BY` in `CREATE TABLE`.
  *(Use functions like `DATE_TRUNC`, `TO_YYYYMM`, `TO_YYYYMMDD`, `EXTRACT` for partition keys.)*
* **Optimization:**
  * **VACUUM:** Reclaims storage from deleted/updated rows and optimizes tablet structure. Can be run manually or occurs automatically in the background.
  * **Data Pruning:** Firebolt uses index metadata to skip scanning irrelevant data ranges (data pruning) and tablets (tablet pruning). Partitioning adds partition pruning.

*(To learn how to design tables and indexes, refer to the Data Modeling and Using Indexes sections in the documentation.)*

## 6. Indexes and Best Practices

Indexes are critical for query performance in Firebolt.

* **Primary Index:**
  * Sorts data within tablets based on specified columns. Essential for efficient data pruning.
  * Defined at table creation using `PRIMARY INDEX <col1>, [<col2>...]`. Cannot be altered later.
  * **Best Practices:** Order columns by selectivity (highest cardinality first); include columns frequently used in `WHERE`, `JOIN`, `GROUP BY` clauses; use columns directly in filters without transformations.
  * *(Refer to the Primary Index documentation for details.)*
* **Aggregating Index:**
  * Precomputes and stores results of aggregate functions (`SUM`, `COUNT`, `AVG`, etc.).
  * Defined using `CREATE AGGREGATING INDEX ... ON ... (<grouping_cols...>, <aggregations...>)`. Can be added/dropped after table creation.
  * Automatically maintained upon data changes (`INSERT`, `UPDATE`, `DELETE`).
  * Queries automatically use the index if the `GROUP BY` clause and aggregations match the index definition.
  * **Best Practices:** Define for frequently run aggregation queries; include all necessary grouping columns and measures; place low-cardinality columns first in the grouping key definition.
  * *(Use `EXPLAIN` to verify index usage. Refer to the Aggregating Index documentation for details.)*
* **Index Granularity:** Advanced setting (`index_granularity` in `CREATE TABLE WITH (...)`) controls the number of rows per index range (default 8192). Smaller values may improve highly selective queries but increase memory usage.
* **Recommendation Engine:** Use `CALL recommend_ddl(<table_name>, (<select_query_workload>))` to get suggestions for primary index and partitioning based on query history.

## 7. Available Views in `information_schema`

The `information_schema` database provides read-only views containing metadata about your Firebolt organization and accounts. Query these views to understand your environment.

* **Organization Level:**
  * `information_schema.accounts`: Lists accounts within the organization.
  * `information_schema.logins`: Lists logins (human users) in the organization.
  * `information_schema.service_accounts`: Lists service accounts (programmatic access).
  * `information_schema.network_policies`: Lists defined network policies.
* **Account Level (Context Dependent):**
  * `information_schema.databases` / `information_schema.catalogs`: Lists databases in the current account.
  * `information_schema.engines`: Lists engines in the current account.
  * `information_schema.users`: Lists users defined within the current account.
  * `information_schema.applicable_roles`: Lists roles and their direct grantees within the account.
  * `information_schema.transitive_applicable_roles`: Lists roles and all direct/indirect grantees.
  * `information_schema.enabled_roles`: Lists roles accessible by the current user session.
  * `information_schema.object_privileges`: Shows permissions granted on objects. Context-dependent (account or database level).
* **Database Level (Requires specific database context):**
  * `information_schema.schemata`: Lists schemas within the current database.
  * `information_schema.tables`: Lists tables (Fact, Dimension, External) and views in the current database.
  * `information_schema.views`: Lists views specifically in the current database.
  * `information_schema.columns`: Lists columns for tables/views in the current database.
  * `information_schema.indexes`: Lists primary and aggregating indexes in the current database.
* **Monitoring & History:**
  * `information_schema.engine_query_history`: Contains history of queries executed on engines (up to 10k per cluster).
  * `information_schema.engine_user_query_history`: Filtered query history showing only user-initiated queries.
  * `information_schema.engine_running_queries`: Shows queries currently executing or queued on engines (up to 10k per cluster).
  * `information_schema.engine_history`: Shows engine lifecycle events (create, start, stop, alter).
  * `information_schema.engine_metering_history`: Hourly engine FBU consumption.
  * `information_schema.engine_metrics_history`: Engine resource utilization snapshots (CPU, RAM, cache, disk) taken every 30s, retained for 30 days.
  * `information_schema.storage_billing`: Daily storage costs per account/region.
  * `information_schema.engines_billing`: Daily engine compute costs per engine/account/region.
  * `information_schema.storage_history`: Detailed daily storage consumption per database.
  * `information_schema.storage_metering_history`: Aggregated daily storage consumption per account.
  * `information_schema.routines`: Lists available SQL functions.

*(To explore the specific columns available in each view, query the view directly or consult the `information_schema` documentation section.)*

## 8. SQL Syntax Overview

Firebolt uses a SQL dialect largely compatible with PostgreSQL. Key operations include:

* **Data Definition Language (DDL):**
  * `CREATE`, `ALTER`, `DROP` for `DATABASE`, `TABLE`, `VIEW`, `INDEX`, `ENGINE`, `ROLE`, `USER`, `LOGIN`, `SERVICE ACCOUNT`, `NETWORK POLICY`, `ACCOUNT`, `ORGANIZATION`.
* **Data Manipulation Language (DML):**
  * `INSERT INTO ... VALUES ...`: For inserting individual rows.
  * `INSERT INTO ... SELECT ...`: For inserting results of a query.
  * `UPDATE ... SET ... WHERE ...`: For modifying existing rows.
  * `DELETE FROM ... WHERE ...`: For removing rows.
  * `TRUNCATE TABLE ...`: For removing all rows quickly.
  * `COPY FROM ...`: Bulk load data from S3 into a Firebolt table. Supports schema inference and various file types (CSV, Parquet, JSON, etc.).
  * `COPY (...) TO ...`: Export query results to S3. Supports CSV, TSV, JSON, Parquet.
  * `VACUUM ...`: Optimize table storage and remove deleted rows.
* **Data Query Language (DQL):**
  * `SELECT ... FROM ... WHERE ... GROUP BY ... HAVING ... ORDER BY ... LIMIT ... OFFSET ...`: Standard query structure.
  * Supports `JOIN` types (`INNER`, `LEFT`, `RIGHT`, `FULL`, `CROSS`).
  * Supports `UNION [ALL]`.
  * Supports Common Table Expressions (`WITH ... AS ...`). Materialized CTEs available via `WITH ... AS MATERIALIZED ...`.
  * Supports Window Functions (`... OVER (PARTITION BY ... ORDER BY ...)`).
  * Supports Table-Valued Functions like `UNNEST` and `GENERATE_SERIES`.
* **Pipe Syntax:** Alternative syntax for chaining transformations (`FROM table |> WHERE ... |> SELECT ...`).

*(For detailed syntax of any command or function, query `information_schema.routines` or consult the SQL Reference section of the documentation.)*

## 9. Supported Functions (List by Category)

Firebolt provides a wide range of SQL functions. Query `information_schema.routines` for a full list or refer to the documentation. Categories include:

* **Aggregation:** `AVG`, `COUNT`, `SUM`, `MIN`, `MAX`, `BOOL_AND`, `BOOL_OR`, `BIT_AND`, `BIT_OR`, `BIT_XOR`, `ARRAY_AGG`, `HASH_AGG` (CHECKSUM), `MEDIAN`, `PERCENTILE_CONT`, `HLL_COUNT_BUILD`, `HLL_COUNT_MERGE`, `MIN_BY`, `MAX_BY`, `STDDEV_POP`, `STDDEV_SAMP`, `VAR_POP`, `VAR_SAMP`, `APACHE_DATASKETCHES_HLL_BUILD`, `APACHE_DATASKETCHES_HLL_MERGE`.
* **Array:** `ARRAY_CONCAT`, `ARRAY_CONTAINS`, `ARRAY_COUNT_DISTINCT`, `ARRAY_DISTINCT`, `ARRAY_ENUMERATE`, `ARRAY_LENGTH`, `ARRAY_MAX`, `ARRAY_MIN`, `ARRAY_REVERSE`, `ARRAY_TO_STRING`, `FLATTEN`, `INDEX_OF`.
* **Lambda (for Arrays):** `ARRAY_ALL_MATCH`, `ARRAY_ANY_MATCH`, `ARRAY_COUNT`, `ARRAY_FILTER`, `ARRAY_FIRST`, `ARRAY_REVERSE_SORT`, `ARRAY_SORT`, `ARRAY_SUM`, `ARRAY_TRANSFORM`.
* **Binary:** `CONVERT_FROM`, `DECODE`, `ENCODE`, `OCTET_LENGTH`.
* **Conditional/Misc:** `CASE`, `CAST`, `TRY_CAST`, `COALESCE`, `IF`, `IFNULL`, `NULLIF`, `GREATEST`, `LEAST`, `CITY_HASH`, `HASH`, `TYPEOF`, `VERSION`.
* **Date & Time:** `CURRENT_DATE`, `CURRENT_TIMESTAMP` (NOW), `LOCALTIMESTAMP`, `DATE`, `DATE_ADD`, `DATE_DIFF`, `DATE_TRUNC`, `EXTRACT`, `TO_DATE`, `TO_TIMESTAMP` (FROM_UNIXTIME), `TO_YYYYMM`, `TO_YYYYMMDD`, `TO_CHAR`.
* **JSON:** `JSON_EXTRACT` (JSON_POINTER_EXTRACT), `JSON_EXTRACT_ARRAY` (JSON_POINTER_EXTRACT_ARRAY), `JSON_VALUE`, `JSON_VALUE_ARRAY`, `JSON_POINTER_EXTRACT_KEYS`, `JSON_POINTER_EXTRACT_VALUES`, `JSON_POINTER_EXTRACT_TEXT`.
* **Numeric:** `ABS`, `ACOS`, `ASIN`, `ATAN`, `ATAN2`, `CEIL` (CEILING), `COS`, `COT`, `DEGREES`, `FLOOR`, `IS_FINITE`, `IS_INFINITE`, `LN`, `LOG` (LOG10), `MOD`, `PI`, `POW` (POWER), `RADIANS`, `RANDOM`, `ROUND`, `SIGN`, `SIN`, `SQRT`, `TAN`, `BIT_SHIFT_LEFT`, `BIT_SHIFT_RIGHT`, `HLL_COUNT_ESTIMATE`, `APACHE_DATASKETCHES_HLL_ESTIMATE`.
* **String:** `CONCAT` (||), `LOWER`, `UPPER`, `LENGTH`, `LPAD`, `RPAD`, `LTRIM`, `RTRIM`, `BTRIM` (TRIM), `POSITION`, `STRPOS`, `REPLACE`, `SUBSTRING` (SUBSTR), `SPLIT_PART`, `STRING_TO_ARRAY`, `LIKE`, `ILIKE`, `REGEXP_LIKE`, `REGEXP_LIKE_ANY`, `REGEXP_EXTRACT`, `REGEXP_EXTRACT_ALL`, `REGEXP_REPLACE`, `REGEXP_REPLACE_ALL`, `URL_ENCODE`, `URL_DECODE`.
* **Vector:** `VECTOR_ADD`, `VECTOR_SUBTRACT`, `VECTOR_COSINE_DISTANCE`, `VECTOR_COSINE_SIMILARITY`, `VECTOR_EUCLIDEAN_DISTANCE`, `VECTOR_SQUARED_EUCLIDEAN_DISTANCE`, `VECTOR_INNER_PRODUCT`, `VECTOR_MANHATTAN_DISTANCE`.
* **Window:** `AVG`, `COUNT`, `SUM`, `MIN`, `MAX`, `ROW_NUMBER`, `RANK`, `DENSE_RANK`, `PERCENT_RANK`, `LEAD`, `LAG`, `FIRST_VALUE`.
* **Table-Valued:** `GENERATE_SERIES`, `LIST_OBJECTS`, `READ_CSV`, `READ_PARQUET`, `UNNEST`.
* **Geospatial:** `ST_ASBINARY`, `ST_ASEWKB`, `ST_ASGEOJSON`, `ST_ASTEXT`, `ST_CONTAINS`, `ST_COVERS`, `ST_DISTANCE`, `ST_GEOGFROMGEOJSON`, `ST_GEOGFROMTEXT`, `ST_GEOGFROMWKB`, `ST_GEOGPOINT`, `ST_INTERSECTS`, `ST_S2CELLIDFROMPOINT`, `ST_X`, `ST_Y`.

*(To learn about a specific function's syntax, parameters, and behavior, search for its name in the Firebolt documentation.)*

## 10. Supported Data Types

Firebolt supports standard SQL data types.

* **Numeric:** `INTEGER` (INT), `BIGINT` (LONG), `NUMERIC` (DECIMAL), `REAL` (FLOAT4), `DOUBLE PRECISION` (FLOAT, DOUBLE).
* **Boolean:** `BOOLEAN` (BOOL).
* **String:** `TEXT`. Uses UTF-8 encoding.
* **Binary:** `BYTEA`.
* **Date/Time:** `DATE`, `TIMESTAMP` (Timestamp without time zone), `TIMESTAMPTZ` (Timestamp with time zone).
* **Composite:** `ARRAY(<type>)`. Can be nested.
* **Spatial:** `GEOGRAPHY`. Represents points, linestrings, polygons on the WGS84 spheroid.

*(For details on range, precision, storage size, literal formats, and casting rules, consult the Data Types section in the documentation.)*

---