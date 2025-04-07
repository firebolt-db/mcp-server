# [](#engine-consumption)Engine Consumption

This page explains the engine consumption metric, Firebolt Units (FBUs), and provides detailed examples showing how to calculate the following:

- The FBUs available to an engine.
- The actual FBUs consumed by a running engine.

For more information about Firebolt engines, see [Understanding Engine Fundamentals](/Overview/engine-fundamentals.html).

Firebolt uses a metric called Firebolt Units (FBU) to track engine consumption. The number of FBUs available to an engine is determined by the following engine attributes:

- The node type.
- The number of nodes.
- The number of clusters.

Each node type provides a minimum number of FBUs as shown in the following table:

Node type Compute family Firebolt Units (FBUs) per hour Small (S) Storage-optimized 8 Medium (M) Storage-optimized 16 Large (L) Storage-optimized 32 Extra Large (XL) Storage-optimized 64 Small (S) Compute-optimized 4 Medium (M) Compute-optimized 8 Large (L) Compute-optimized 16 Extra Large (XL) Compute-optimized 32

Small and medium engines are available for use right away. If you want to use a large or extra-large engine, reach out to support@firebolt.io.

As you add more nodes and more clusters to your engine, the number of FBUs available to the engine increases linearly. The available FBUs for a given engine configuration are calculated as follows:

## [](#fbu-per-hour-for-a-given-engine--fbu-of-node-type-x-nodes-x-clusters)FBU-per-hour for a given Engine = (FBU of node Type x Nodes x Clusters)

### [](#calculating-fbus-per-hour-for-a-given-engine---example-1)Calculating FBUs-per-hour for a given Engine - Example 1:

- If you create an engine with the following configuration: TYPE = “S”, NODES = 2, CLUSTERS=1, it will have 16 FBUs available per hour (8 x 2 x 1).
- If you scale out the engine configuration to use 3 nodes instead of 2, the available FBUs per hour increases to 24 FBUs (8 x 3 x 1).

### [](#calculating-fbus-per-hour-for-a-given-engine---example-2)Calculating FBUs-per-hour for a given Engine - Example 2:

- If you create an engine with the following configuration: TYPE = “L”, NODES = 3, CLUSTERS=2, it will have 192 FBUs available per hour (32 x 3 x 2).
- If you scale down the engine to use a “M” type node, the available FBUs per hour will decrease to 96 FBUs (16 x 3 x 2).

FBUs are consumed only when an engine is in a running state. The FBUs consumed by an engine depends on both the current configuration of the engine (node type, number of nodes and number of clusters) and the duration for which the engine has been running in that configuration. Since Firebolt provides per-second billing, the consumed FBU is calculated at per-second granularity, as shown in the following examples:

## [](#fbus-consumed---fbus-per-hour--3600-x-engine-runtime-in-seconds)FBUs Consumed = (FBUs per hour / 3600) x (engine runtime in seconds)

### [](#consumed-fbu---example-1)Consumed FBU - Example 1:

If you create an engine with the following configuration: TYPE = “S”, NODES = 2, CLUSTERS=1, it has 16 available FBUs per hour. If the engine runs for only 15 minutes of that hour, then the engine consumes 4 FBUs.

### [](#consumed-fbu---example-2)Consumed FBU - Example 2:

If you create an engine with the following configuration: TYPE = “L”, NODES = 3, CLUSTERS=2, it has 192 FBUs available per hour. If the engine was only running for 40 seconds to quickly ingest a small amount of data. The consumed FBU is calculated as follows:

FBU Consumed = (192/3600) x 40 = 2.13

## [](#the-impact-of-resize-operations-on-engine-consumption)The impact of resize operations on engine consumption

If you perform scaling operations on an engine, you add additional compute resources (clusters) to the engine. These operations include scaling up or down, and scaling out or in.

Any new clusters resulting from a resizing operation will have the new **node type** for a scale up or down operation. Any new clusters will have the desired **number of nodes** for a scale in or out operation. Existing clusters will be removed after they finish running any currently running queries. New queries will run on the new resized clusters. This can result in a temporary overlap when both old and new clusters are running concurrently, which will be reflected in the engine’s consumption.

### [](#example-1)Example 1:

If you create an engine with the following configuration: TYPE = “S”, NODES = 2, CLUSTERS=1, there are 16 available FBUs per hour.

- **First 15 minutes during the hour:** The engine is running in the previous configuration.
- **After first 15 minutes:** The engine is scaled out to use five nodes, keeping the same node type and same number of clusters. The old cluster with two nodes continues to run for 5 minutes to finish executing the queries. The new cluster with five nodes runs for the next fifteen minutes.

The consumed FBUs are calculated as follows:

- **Minutes 1-15:** Only one cluster with two nodes is running, consuming 4 FBUs.
- **Minutes 16-20:** Two clusters are running. the cluster with nodes consumes 1.3 FBUs, and the cluster with five nodes consumes 3.3 FBUs.
- **Minutes 20-30:** Only the cluster with five nodes is running and consumes 6.7 FBUs.

Total number of FBUs consumed for the hour = 4 + 1.3 + 3.3 + 6.7 = 15.3 FBUs.

### [](#example-2)Example 2:

If you create an engine with the following configuration: TYPE = “L”, NODES = 2, CLUSTERS=1, there are 64 available FBUs per hour.

- **First 15 minutes during the hour:** The engine is running in the above configuration.
- **After first 15 minutes:** The engine is scaled down to use “M” type nodes, keeping the same number of nodes and same number of clusters. The old cluster with “L” type nodes continues to run for 5 minutes to finish running the queries. The new cluster with “M” type nodes runs for the next fifteen minutes.

The consumed FBUs are calculated as follows:

- **Minutes 1-15:** Only one cluster with two “L” type nodes is running, consuming 16 FBUs.
- **Minutes 16-20:** Two clusters running. Cluster with two “L” type nodes consumes 5.3 FBU and Cluster with two “M” type nodes consumes 2.7 FBU.
- **Minutes 20-30:** Only the cluster with two “M” type nodes is running, consuming 5.3 FBUs.

Total number of FBUs consumed for the hour = 16 + 5.3 + 2.7 + 5.3 = 29.3 FBUs.

Firebolt Engines are priced based on the amount of FBUs consumed by a given engine. For more details on engine pricing, including examples, visit the [Pricing page](https://www.firebolt.io/pricing).