# Important note
You are on Firebolt Core! This means that you should pay attention to the differences between Firebolt Core and Managed Firebolt, as several functions, for instance information_schema objects are not available or behave differently.

# Firebolt Core Foundational Knowledge Layer

This document provides a foundational overview of Firebolt Core, designed to help an LLM learn its specific architecture, concepts, and capabilities. Firebolt Core is a version of Firebolt that can be rolled out locally or in self-managed environments, with some differences and limitations compared to the managed Firebolt SaaS solution.

**Model Behavior Guidelines:**

* If you lack sufficient information to answer a user's question with high confidence, state: *"Please consult the official Firebolt documentation at https://docs.firebolt.io for more details."*
* **Never** hallucinate function names, table structures, SQL syntax, or other technical details. When unsure, refer to the official documentation or indicate the need for it.
* **Distinguish between Firebolt Core and Managed Firebolt:** Always be clear about which environment the information applies to, as many SaaS features (like RBAC, Accounts, Engines) are not available or behave differently in Firebolt Core.

---

## 1. Firebolt Core Architecture Overview

Firebolt Core is a high-performance, low-latency data warehouse designed for local or self-managed deployments. Unlike the managed SaaS version, Firebolt Core does not decouple storage and compute.

* **Key Characteristics:** Local deployment, high efficiency, and SQL simplicity.
* **Nodes:** A cluster consists of one or more Firebolt Core nodes. Each node runs in a Docker container.
* **State Management:** Persistent database state is stored on the local filesystem of individual nodes. No data is uploaded to cloud storage (S3/GCS) for persistence of managed objects, though they can be used for ingestion/export.
* **Catalog and Sharding:** 
  * The **database catalog** (metadata) is stored on **node 0**.
  * **Table data** is sharded across all nodes in the cluster.
* **Scalability Limitations:** Because data is sharded across nodes on local disk, Firebolt Core clusters **cannot be easily resized** (adding or removing nodes) without data loss or skew. Note: This restriction doesn't apply to workloads exclusively using External or Iceberg tables.

## 2. No Storage-Compute Isolation

In Firebolt Core, storage and compute are coupled.

* **Local Storage:** Data managed by Firebolt Core is tied to the specific cluster and stored on the nodes' local filesystems (typically mounted via Docker volumes).
* **No Shared Data:** Data managed by one Firebolt Core cluster cannot be shared with another Firebolt Core cluster.
* **Compute (Nodes):** Queries are submitted to any node, which acts as the "leader" for that query, coordinating with "follower" nodes.

## 3. Organizations, Accounts, and Users (Limitations)

Many of the hierarchical management entities found in Managed Firebolt are unavailable or simplified in Firebolt Core.

* **No Organizations/Accounts:** Firebolt Core operates as a single-account environment. SQL commands like `CREATE ACCOUNT` or `CREATE ORGANIZATION` are not available.
* **Placeholder Information Schema:** `information_schema.accounts`, `information_schema.logins`, etc., may exist but will return placeholder results.
* **No Authentication:** Firebolt Core has **no built-in authentication**. Anyone with network access to the HTTP endpoint can submit queries.
* **No RBAC:** There is **no role-based access control**. Commands like `CREATE USER`, `CREATE ROLE`, `GRANT`, or `REVOKE` are unavailable and will return errors.

## 4. Security and Network

Firebolt Core is designed for secure environments and lacks built-in security features found in the SaaS version.

* **Network Encryption:** There is **no encryption** for inter-node traffic or client-to-node traffic (uses unencrypted HTTP).
* **Access Control:** As there is no authentication or RBAC, security must be managed at the network level (e.g., via firewalls or VPCs).
* **Secrets:** Secrets used for accessing external storage (AWS keys) are transmitted in plain text over the network if not secured by external means.

## 5. Data Modeling and Performance

Firebolt Core stores data in a compressed columnar format within internal objects called tablets. Effective data modeling is crucial for performance.

* **Table Types:**
  * **Fact Tables:** Store large volumes of quantitative/event data. Distributed/sharded across nodes. (Default type for `CREATE TABLE`).
  * **Dimension Tables:** **Not supported** in Firebolt Core. All tables are sharded.
  * **External Tables:** Reference data in S3/GCS without loading it. Used primarily for ingestion. Use `CREATE EXTERNAL TABLE`.
* **Partitioning:** Tables can be partitioned (e.g., by date) to improve query performance by allowing the engine to scan only relevant partitions (partition pruning). Use `PARTITION BY` in `CREATE TABLE`.
  * Common partition keys: `DATE_TRUNC`, `TO_YYYYMM`, `TO_YYYYMMDD`.
* **Optimization:**
  * **Manual VACUUM:** Reclaims storage from deleted/updated rows and optimizes tablets. **No auto-vacuum** support; it must be run manually.
  * **Data Pruning:** Firebolt uses index metadata to skip scanning irrelevant data ranges (data pruning) and tablets (tablet pruning).
  * **No Auto-scaling:** Cluster size is fixed upon deployment.
  * **No Admission Controller:** Clients are responsible for load balancing and retrying failed queries.

## 6. Indexes and Best Practices

Indexes are critical for performance in Firebolt Core.

* **Primary Index:**
  * Sorts data within tablets based on specified columns. Essential for efficient data pruning.
  * Defined at table creation using `PRIMARY INDEX <col1>, [<col2>...]`. Cannot be altered later.
  * **Best Practices:** Order columns by selectivity (highest cardinality first); include columns frequently used in `WHERE`, `JOIN`, `GROUP BY` clauses.
* **Aggregating Index:**
  * Precomputes and stores results of aggregate functions (`SUM`, `COUNT`, `AVG`, etc.).
  * Defined using `CREATE AGGREGATING INDEX ... ON ... (<grouping_cols...>, <aggregations...>)`.
  * Automatically maintained upon data changes (`INSERT`, `UPDATE`, `DELETE`).
  * **Best Practices:** Define for frequently run aggregation queries; include all necessary grouping columns and measures.
* **No RECOMMEND DDL:** The `recommend_ddl` tool is **not available** in Firebolt Core.

## 7. Information Schema in Firebolt Core

While `information_schema` is available, many views related to managed features return placeholder data or are restricted.

* **Available (Local Context):** `databases`, `tables`, `columns`, `indexes`, `schemata`, `views`, `routines`.
* **Monitoring:** `engine_query_history`, `engine_running_queries` (local to the cluster).
* **Placeholders/Unavailable:** `accounts`, `logins`, `service_accounts`, `network_policies`, `billing` related views.

## 8. SQL Syntax Overview

Firebolt Core uses a SQL dialect largely compatible with PostgreSQL.

* **Data Definition Language (DDL):**
  * `CREATE`, `ALTER`, `DROP` for `DATABASE`, `TABLE`, `VIEW`, `INDEX`.
  * `USE DATABASE <db_name>`: Switches the current database context (requires client-side header handling).
* **Data Manipulation Language (DML):**
  * `INSERT INTO ... VALUES ...`: For inserting individual rows.
  * `INSERT INTO ... SELECT ...`: For inserting results of a query.
  * `UPDATE ... SET ... WHERE ...`: For modifying existing rows.
  * `DELETE FROM ... WHERE ...`: For removing rows.
  * `TRUNCATE TABLE ...`: For removing all rows quickly.
  * `COPY FROM ...`: Bulk load data from S3/GCS into a Firebolt table. Supports schema inference and various file types (CSV, Parquet, JSON, etc.).
  * `COPY (...) TO ...`: Export query results to S3/GCS. Supports CSV, TSV, JSON, Parquet.
  * `VACUUM ...`: Optimize table storage and remove deleted rows (manual only).
* **Data Query Language (DQL):**
  * `SELECT ... FROM ... WHERE ... GROUP BY ... HAVING ... ORDER BY ... LIMIT ... OFFSET ...`: Standard query structure.
  * Supports `JOIN` types (`INNER`, `LEFT`, `RIGHT`, `FULL`, `CROSS`).
  * Supports `UNION [ALL]`, CTEs (`WITH ... AS ...`), and Window Functions.
  * **Pipe Syntax:** Alternative syntax for chaining transformations (`FROM table |> WHERE ... |> SELECT ...`).

## 9. Transactions and Protocol

* **Transactions:**
  * **Fully Transactional:** Statements are auto-committed.
  * **Write Limit:** At most **one write transaction** can be active at a time across the entire cluster.
  * **Explicit Transactions:** Supported via `BEGIN TRANSACTION`, `COMMIT`, `ROLLBACK`.
  * **Multi-node Health:** In multi-node setups, all nodes must be up for the cluster to be healthy.
* **Stateless Protocol & Sessions:**
  * Communication is via HTTP POST.
  * The protocol is stateless, but the cluster uses the `Firebolt-Update-Parameters` HTTP header to instruct clients to maintain session state (e.g., for `USE DATABASE` or `transaction_id`).

## 10. Supported Data Types

* **Numeric:** `INTEGER` (INT), `BIGINT` (LONG), `NUMERIC` (DECIMAL), `REAL` (FLOAT4), `DOUBLE PRECISION` (FLOAT, DOUBLE).
* **Boolean:** `BOOLEAN` (BOOL).
* **String:** `TEXT` (UTF-8).
* **Binary:** `BYTEA`.
* **Date/Time:** `DATE`, `TIMESTAMP`, `TIMESTAMPTZ`.
* **Composite:** `ARRAY(<type>)`.
* **Spatial:** `GEOGRAPHY`.

## 11. Supported Functions (Categories)

Firebolt Core supports a wide range of SQL functions. Query `information_schema.routines` for a full list.

* **Aggregation:** `AVG`, `COUNT`, `SUM`, `MIN`, `MAX`, `ARRAY_AGG`, `MEDIAN`, `PERCENTILE_CONT`, `HLL_COUNT_BUILD`, `MIN_BY`, `MAX_BY`.
* **Array & Lambda:** `ARRAY_CONCAT`, `ARRAY_CONTAINS`, `ARRAY_LENGTH`, `ARRAY_FILTER`, `ARRAY_TRANSFORM`, `FLATTEN`, `UNNEST`.
* **JSON:** `JSON_EXTRACT`, `JSON_VALUE`, `JSON_POINTER_EXTRACT_TEXT`.
* **Date & Time:** `CURRENT_TIMESTAMP`, `DATE_ADD`, `DATE_DIFF`, `DATE_TRUNC`, `EXTRACT`, `TO_CHAR`.
* **Numeric & String:** `ABS`, `CEIL`, `FLOOR`, `ROUND`, `CONCAT`, `LOWER`, `UPPER`, `SUBSTRING`, `REPLACE`, `REGEXP_LIKE`.
* **Table-Valued:** `GENERATE_SERIES`, `READ_CSV`, `READ_PARQUET`, `READ_ICEBERG`.

## 12. Data Ingestion and Export

* **Inbound:** Supports `COPY FROM`, `CREATE EXTERNAL TABLE`, and table-valued functions like `read_parquet`, `read_csv`, and `read_iceberg`.
* **Object Storage:** Supports S3, GCS (via S3 interoperability), and S3-compatible storage (like MinIO via `default_s3_endpoint_override`).
* **Outbound:** Supports `COPY TO` for exporting to S3/GCS.

## 13. Connection and SDKs

* **HTTP Endpoint:** Default port is `3473`.
* **Output Formats:** `TabSeparatedWithNamesAndTypes` (default), `JSON_Compact`, `JSONLines_Compact`.
* **SDK and Tool Support:**
  * **CLI:** Use `fb --core` to connect. Port 3473 by default.
  * **Go SDK:** Supported from version 1.10.0+. Use `firebolt:///<database>?url=<http_endpoint_url>`.
  * **JDBC Driver:** Supported from version 3.6.3+.
  * **Python SDK:** Supported from version 1.13.0+.
  * **Web UI:** Available on port `9100` on each node.
  * **Integrations:** Supports Apache Superset and dbt.

## 14. System Requirements (Local Run)

* **OS:** Linux kernel >= 6.1 (requires `io_uring`).
* **Docker:** Requires `seccomp=unconfined` and increased `memlock` ulimit.
* **Hardware:** amd64 (with SSE 4.2) or arm64; 16GB RAM and 25GB SSD recommended.

## 15. Unavailable Statements (Summary)

The following statements commonly used in Managed Firebolt are **unavailable** in Firebolt Core:
* **Access Control:** `GRANT`, `REVOKE`, `CREATE ROLE`, `CREATE USER` (Core has no built-in auth/RBAC).
* **Account Management:** `CREATE ACCOUNT`, `CREATE ORGANIZATION`, `CREATE LOGIN`, `CREATE SERVICE ACCOUNT`.
* **Engine Management:** `CREATE ENGINE`, `START ENGINE`, `STOP ENGINE`, `ALTER ENGINE`.
* **Optimization Tools:** `RECOMMEND DDL`.
* **Other:** `DIMENSION` table type, `asynchronous query execution`, `LLM token budget accounting`.

---

