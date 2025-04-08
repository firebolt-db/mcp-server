# [](#sizing-engines)Sizing Engines

Selecting an appropriate engine size for your workload depends on multiple factors such as the size of your active dataset, latency and throughput requirements of your workload, your considerations for price-performance and the number of users and queries your workload is expected to handle concurrently. Our guidance is to start small with an engine size that fits your active dataset and monitor the workload using the engine observability metrics (see below). Based on these metrics, you can then dynamically resize your engine to meet the needs of your workload.

If your workload requires high processing power relative to data size, use a compute-optimized node type. These nodes have the approximately same processing power as storage-optimized nodes but have less memory, cache space, and cost.

## [](#dimensions-of-engine-sizing)Dimensions of engine sizing

Firebolt allows you to change:

- The type of nodes in your engine.
- The compute family of nodes in your engine.
- The number of nodes in a cluster of your engine.
- The number of copies of that cluster in your engine.

See the [engine fundamentals](/Overview/engine-fundamentals.html) page for details.

## [](#using-observability-metrics-to-resize-an-engine)Using Observability Metrics to Resize an Engine

Firebolt provides engine observability metrics that give visibility into how the engine resources are being utilized by your workloads. Use the [Information\_Schema.engine\_metrics\_history](/sql_reference/information-schema/engine-metrics-history.html) view to understand how much CPU, RAM, and disk are utilized by your workloads. In addition, this view also provides details on how often your queries hit the local cache and how much of your query data is spilling onto the disk. These metrics can help you decide whether your engine needs a different node type and whether you need to add more nodes to improve the query performance. Use the [Information\_Schema.engine\_running\_queries](/sql_reference/information-schema/engine-running-queries.html) view to understand how many queries are waiting in the queue to be run. If there are a number of queries still waiting to be run, adding another cluster to your engine may help improve the query throughput.

## [](#initial-sizing)Initial Sizing

**ELT Workloads**

For the ELT workloads, the engine size would depend on the number of files and the size of the files used to ingest the data. You can parallelize the ingest process with additional nodes, which can provide improved performance.

**Queries**

To correctly size an engine for querying data, there are several factors to consider:

- The size of frequently accessed data under your query pattern. More data will require a engine with a larger cache size.
- The relative amount of processing performed within the queries in your query pattern. More complex queries will generally require more CPU cores.
- The Queries Per Second (QPS) of your workload. At higher QPS you may need to enable auto-scaling or multiple clusters on your engine.
- The number of requests outstanding or the time submitted queries run. Longer running queries may raise the requirements of the engine to need instance types with more memory or more clusters in the engine.

For query processing, our recommendation is to start with a S or M storage-optimized instance type. Then, run a checksum over the dataset you expect to be queried frequently. Firebolt Engines cache the data locally, which helps serve queries at low latencies. The cache size provided by the engines varies depending on the type of node used in your engines, with each size having twice the cache of the next smallest size. Compute-optimized instances have approximately one quarter of the cache size of storage-optimized instances. After the checksum, you can use [Information\_Schema.engine\_metrics\_history](/sql_reference/information-schema/engine-metrics-history.html) to see the cache utilization percentage. If an acceptable percentage of your active dataset fits, you can then run queries at your expected QPS on the engine.

**TIP:** You can use [Multi-Cluster Engine Warmup](/Reference/system-settings.html#multi-cluster-engine-warmup) to submit your checksum queries to all clusters in a multi-cluster engine.

Small and medium storage-optimized engines are available for use right away. Compute-optimized instance types are available, but may see longer engine start times. If you want to use a large or extra-large engine, reach out to [support@firebolt.io](mailto:support@firebolt.io).

**TIP:** You also have the option to run your workload simultaneously on engines with different configurations and use these metrics to identify which configuration best fits your needs.

You will need to have the appropriate [RBAC](/Guides/operate-engines/rbac-for-engines.html) permissions to use the engine observability metrics.