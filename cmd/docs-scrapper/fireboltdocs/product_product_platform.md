# [](#firebolt-platform-capabilities)Firebolt platform capabilities

## [](#-data-management)![Icon for data management.](../../assets/images/icon-data_management.png) Data management

Firebolt is built for performance and efficiency, with built-in features that reduce manual effort. Its robust architecture supports parallel data ingestion, ACID-compliant DML, and optimized storage, allowing it to manage large and complex datasets. Automated indexing and backend optimization ensure consistent query performance, even during data modifications like inserts, updates, and deletes. Firebolt also supports both structured and semi-structured data, making it easy to integrate into existing workflows for tasks such as table creation, data ingestion, and transaction management.

For more information, see [Data management](/Overview/data-management.html).

## [](#-query-processing)![Icon for query processing.](../../assets/images/icon-processing.png) Query processing

Firebolt’s query processing is optimized for low latency, high concurrency, and dynamic scaling to meet workload demands. Its optimizer creates efficient execution plans by evaluating data distribution, indexing, and previous query patterns. The distributed runtime system leverages multithreading and tiered caching for faster processing, while resource-aware scheduling and efficient memory management enhance cluster performance. This setup ensures fast, consistent query execution, even for complex queries and large datasets.

You can also use the Firebolt API for programmatic access to performing tasks including querying, managing resources, and automating workflows.

For more information, see the [Firebolt API](/Guides/query-data/using-the-api.html), and the Firebolt [functions glossary](/sql_reference/functions-reference/functions-glossary.html).

## [](#-security)![Icon for security.](../../assets/images/icon-security.png) Security

Firebolt employs a multi-layered approach to data protection, utilizing industry-standard encryption, secure key management, and granular access control to meet the needs of organizations that build [large, data-intensive applications and products](https://www.firebolt.io/knowledge-center/case-studies). With HIPAA and SOC 2 compliance, Firebolt ensures the confidentiality, integrity, and availability of your data. Identity management combines username/password authentication, Single Sign-On (SSO), and Multi-Factor Authentication (MFA) to secure access to resources. Role-Based Access Control (RBAC) assigns permissions through hierarchical and composable roles, allowing only authorized users to access critical data and system resources.

![Firebolt's security layers include access control, identity management, infrastructure and network security.](../../assets/images/firebolt-security-layers.png)

For more information, see [Security](/Overview/Security/security.html).

## [](#-observability)![Icon for observability.](../../assets/images/icon-observability.png) Observability

Firebolt’s observability features provide detailed metrics on CPU, RAM, disk usage, and cache efficiency through the `engine_metrics_history` and `engine_running_queries` views. These metrics help optimize engine configurations and resource allocation. Firebolt also integrates with OpenTelemetry, enabling users to track telemetry data for deeper insights into performance across distributed systems. This integration enhances observability and supports data-driven adjustments. Access to these tools is managed through Role-Based Access Control (RBAC), ensuring secure management of system resources.

## [](#-collaborative-workspace)![Icon for collaboration.](../../assets/images/icon-collaboration.png) Collaborative workspace

Delivering insights and data products requires collaboration among multiple roles, including data architects, engineers, developers, and IT operations. The **Firebolt Workspace** facilitates team coordination by providing visibility across the entire data lifecycle. It includes dedicated areas for security and governance, data modeling, exploration, SQL development, and performance management, ensuring each role can contribute effectively.

In the Firebolt **Develop Space**, you can edit and run SQL scripts, manage databases, and view query results. This space features a document editor with auto-complete and script templates. You can also rename, copy, or export scripts, and execute SQL snippets or entire scripts with real-time result display. The **Develop Space** interface supports light and dark modes and allows for exporting query results for further use.