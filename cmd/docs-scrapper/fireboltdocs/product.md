# [](#what-is-firebolt)What is Firebolt?

Firebolt is a next-generation cloud data warehouse designed for organizations that require lightning-fast analytics at scale. Whether you’re dealing with complex data applications, high-concurrency workloads, or low-latency queries, Firebolt empowers you to unlock the full potential of your data without compromising on performance or cost.

# [](#why-firebolt)Why Firebolt?

Firebolt is designed specifically to address the demands of modern data workloads, prioritizing **exceptional efficiency** and **low cost** while delivering industry-leading performance for data-intensive applications.

Firebolt addresses the following challenges:

- High latency and costly performance during data processing and retrieval.
- Poor query performance under heavy loads.
- Scaling difficulties when managing large datasets.
- Complexity in managing and optimizing data workflows while working with unfamiliar languages.

The following sections highlight Firebolt’s key benefits, as well as its compatibility with various frameworks and workloads for seamless data integration and processing:

- [Key benefits](#key-benefits)
- [Eco-system and integrations](#eco-system-and-integrations)

# [](#key-benefits)Key benefits

Firebolt is inherently **scalable** helping you to adapt rapidly to changes in your data and workloads. Key benefits include:

- [High Efficiency](#-high-efficiency) — Achieve exceptional price-to-performance ratios, delivering fast analytics without high costs.
- [Concurrency at Scale](#-concurrency-at-scale) — Run thousands of queries concurrently, maintaining sub-second performance even under heavy loads.
- [Elasticity](#-elasticity) — Seamlessly scale to handle hundreds of terabytes of data without sacrificing speed or efficiency.
- [SQL Simplicity](#-sql-simplicity) — Use a subset of PostgreSQL-compatible SQL, allowing teams to adapt easily with minimal training.

### [](#-high-efficiency)![Icon for efficiency.](../../assets/images/icon-efficiency.png) High efficiency

Firebolt delivers low-latency, high performance analytics with one of the best price-to-performance ratios in the industry. Its architecture is optimized for fast query execution through features like vectorized processing and sparse indexing, which minimize data scans and optimize CPU usage. These features enable fast responses to data-intensive queries, even at petabyte scale, without overloading compute resources. Whether analyzing structured or semi-structured data, Firebolt delivers millisecond-level query responses, backed by ACID compliance to ensure data consistency, integrity, and reliability.

### [](#-concurrency-at-scale)![Icon for concurrency at scale.](../../assets/images/icon-concurrency.png) Concurrency at scale

Firebolt enables thousands of concurrent queries, ensuring your applications can manage heavy query loads with consistent, reliable performance. Its fine-grained scaling capabilities allow for high query throughput efficiently, even during peak workloads, ensuring optimal resource allocation and minimizing query latency.

### [](#-elasticity)![Icon for elasticity or vertical, horizontal or concurrent scaling.](../../assets/images/icon-elasticity.png) Elasticity

Firebolt’s fully decoupled architecture and multi-dimensional elasticity allow compute, storage, and management resources to scale independently, optimizing both performance and cost efficiency. This architecture enables fine-grained control over resources as your workloads evolve, including scaling out to accommodate massive datasets. You can access any database from any engine, giving you flexibility to access any data while offering workload isolation to achieve predictable performance. Furthermore, Firebolt’s system allows fine-grained control over provisioned resources for achieving needed price-performance characteristics and minimizing cost. Firebolt’s architecture supports:

- **Vertical scaling** - Scale up to increase the capacity of your engine to process complex queries for data-intensive workloads.
- **Horizontal scaling** - Scale out by adding more compute nodes to handle higher data processing demands efficiently.
- **Concurrent scaling** - Run multiple clusters within a single engine, which can scale up to ten clusters simultaneously. Firebolt manages concurrency scaling transparently to user applications without requiring any endpoint changes.

The following diagram includes code examples of how to scale vertically, horizontally or scale for concurrency using SQL in the **Firebolt Workspace**:

![You can scale vertically, horizontally, or concurrently in the Firebolt Workspace.](../../assets/images/product-scaling-engines.png)

Firebolt’s multi-dimensional approach to elasticity allows it to dynamically adapt to any workload, ensuring optimal system performance while keeping costs under control.

### [](#-sql-simplicity)![Icon for SQL simplicity.](../../assets/images/icon-simplicity.png) SQL simplicity

Firebolt supports a PostgreSQL-compliant SQL dialect, allowing your teams to leverage Firebolt’s capabilities without needing to learn a new query language. This simplifies integration with existing workflows that contain tasks that include data provisioning, processing, and management. With support for both structured and [semi-structured data](/Guides/loading-data/working-with-semi-structured-data/working-with-semi-structured-data.html), Firebolt allows you to analyze diverse datasets within a single platform.

# [](#eco-system-and-integrations)Eco-system and integrations

Firebolt’s platform is optimized for integration within modern data workflows. It supports ingesting data efficiently through ELT tools, making it easy to move data from data lakes, relational databases, and other source systems into Firebolt. With support for Amazon S3 cloud storage and popular file formats like Avro, Parquet, and ORC, Firebolt allows you to centralize and query your data with ultra-fast analytics, seamlessly fitting into your existing data architecture.

You can leverage industry-standard tools like Apache Airflow, dbt, and Superset for orchestration and visualization. Firebolt’s SDKs offer wide support for language clients like Python, Node.js, Java, and .NET. This flexibility empowers your team to build and query data using their preferred environments, ensuring smooth data workflows from ingestion to advanced analytics.

![Firebolt supports popular SDKs and connectors to integrate with many workflows.](../../assets/images/firebolt-framework.png)

### [](#-workload-isolation-for-smooth-operations)![Icon for workload isolation.](../../assets/images/icon-isolation.png) Workload isolation for smooth operations

Firebolt optimizes workloads by considering configuration, resource utilization, and history-based statistics to balance both latency and throughput. Any Firebolt engine can handle both read and write operations on any database, ensuring strong consistency across all engines. Workloads are managed independently with dedicated compute resources, allowing you to run complex ELT processes, fast queries, BI reports, among others, without interference. Workload isolation ensures that resource-heavy tasks do not impact your most critical applications and dashboards, enabling smooth operations across diverse use cases.

# [](#next-steps)Next steps

Read about Firebolt’s [platform capabilities](/product/product-platform.html).

# [](#more-firebolt-resources)More Firebolt resources

Learn how to load and query your data with our [Get Started](/Guides/getting-started/) and [Load data](/Guides/loading-data/loading-data.html) guides or use one of Firebolt’s [test data sets](https://www.firebolt.io/free-sample-datasets).

[Sign up](https://go.firebolt.io/signup) to access your $200 Firebolt credit for your first 30 days. No credit card needed.

- See news, blog posts, whitepapers, events, and videos at [Firebolt’s Knowledge center](https://www.firebolt.io/knowledge-center).
- Explore our [competitive pricing options](https://www.firebolt.io/pricing).

Need help? Firebolt’s support team at [support@firebolt.io](mailto:support@firebolt.io) is here to assist you with:

- Onboarding support.
- Troubleshooting query performance.
- Optimizing database configurations.
- Addressing data loading issues.
- Getting best practices for data modeling.