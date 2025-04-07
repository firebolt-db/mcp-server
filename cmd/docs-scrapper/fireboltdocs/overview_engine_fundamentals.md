# [](#firebolt-engines)Firebolt Engines

Engines are compute resources that process data and serve queries in Firebolt. Use engines to load data into Firebolt and run queries on the ingested data.

Firebolt engines provide **full workload isolation**, so that multiple workloads run independently while sharing the same data. Engines are **decoupled from databases**, meaning:

- A single engine can run queries on multiple databases.
- A database can be queried using multiple engines.

Start, stop and modify engines at any time using the SQL API. You can also **dynamically scale engines** based on workload needs without stopping them.

This document explains engine configuration, scaling, monitoring, security and connection options.

**Topics**

- [Key engine concepts](#key-engine-concepts) – Learn about the `TYPE`, `NODES`, and `CLUSTERS` attributes that define a Firebolt engine’s configuration and scaling options.
- [Multi-dimensional elasticity](#multi-dimensional-elasticity) – Scale engines dynamically by adjusting engine attributes without stopping workloads.
- [Connecting to engines](#connecting-to-engines) – How to connect to a Firebolt engine using the UI, Engine URL, or third-party tools.
- [Monitoring engine usage](#monitoring-engine-usage) – Track engine performance using observability views in `information_schema` to optimize resource allocation.
- [Engine governance and security](#engine-governance-and-security) – Control engine access using [Role-Based Access Control (RBAC)](/Overview/Security/Role-Based%20Access%20Control/) and account-level isolation to enforce security policies.
- [Viewing and understanding engine status](#viewing-and-understanding-engine-status) – Learn how to use `SHOW ENGINES` to check the status of all engines, including running, resizing, and stopped states.

## [](#key-engine-concepts)Key engine concepts

Engines in Firebolt are defined by three attributes: **Type**, **Nodes**, and **Clusters**. These attributes determine the engine’s configuration and scaling options.

**Family**  
Compute nodes can also be storage-optimized with larger cache sizes or compute-optimized which have smaller caches. The default is storage optimized.

**Nodes**  
This attribute represents the number (1 - 128) of compute nodes, allowing granular horizontal scaling to fine-tune query performance characteristics while avoiding overprovisioning and unnecessary cost. Both scaling in and out are supported.

**Clusters**  
A cluster is a collection of compute resources, described by “Type” and “Nodes” attributes. A given Firebolt engine can contain one or more clusters. The maximum number of clusters is specified by the Clusters attribute. Only homogeneous cluster configurations (clusters with the same number of Nodes and Type) are supported within a single engine. Users can leverage the “Clusters” attribute to support query concurrency scaling.

![An engine cluster in Firebolt](/assets/images/engine_cluster_type_M.png)  
**An engine cluster with four nodes of type ‘M’**

![A Firebolt engine with two clusters, each cluster containing four nodes of type 'M'](/assets/images/Engine_two_clusters_type_M.png)  
**A Firebolt engine with two clusters, each cluster containing four nodes of type ‘M’**

The three attributes: `TYPE`, `NODES` and `CLUSTERS` – form the configuration of an engine.

To create an engine, use the [CREATE ENGINE command](/sql_reference/commands/engines/create-engine.html), specifying the node type to be used for the engine, number of clusters and number of nodes per cluster.

The following code example creates two clusters, each containing four nodes of type `M`:

```
CREATE ENGINE IF NOT EXISTS MyEngine 
WITH TYPE = M 
NODES = 4 
CLUSTERS = 2;
```

For a full list of engine attributes, see [CREATE ENGINE](/sql_reference/commands/engines/create-engine.html)

## [](#multi-dimensional-elasticity)Multi-dimensional elasticity

Firebolt engines enable dynamic and fully online scaling operations, meaning you do not need to stop your engines to scale them. You can scale an engine along three dimensions:

Scaling Type Action Example SQL Command **Vertical Scaling** Change the node type `ALTER ENGINE MyEngine SET TYPE = L;` **Horizontal Scaling** Change the number of nodes `ALTER ENGINE MyEngine SET NODES = 3;` **Concurrency Scaling** Change the number of clusters `ALTER ENGINE MyEngine SET CLUSTERS = 2;`

You can scale up or down using the engine type, scaling out or in with number of nodes and add or remove clusters for concurrency scaling. This multidimensional scaling allows you to fine-tune the price-performance characteristics of engines and dynamically scale your compute resources based on your workload requirements.

Use the [ALTER ENGINE](/sql_reference/commands/engines/alter-engine.html) to modify the configuration of an engine to dynamically scale the engine even while it is running, without impacting the workload.

**Best practices**

- Use a **larger node type** to improve performance for both single queries and multiple concurrent queries, especially as data size grows
- Increase the number of nodes for finer control over scaling, such as distributing workloads across multiple smaller nodes or when further vertical scaling is not possible.
- Increase the **number of clusters** to support higher query concurrency.

The following code example uses `ALTER ENGINE` to horizontally scale an engine from two to three nodes:

```
ALTER ENGINE MyEngine SET NODES = 3;
```

The following code example changes the type of node used in an engine from ‘M’ to ‘L’:

```
ALTER ENGINE MyEngine SET TYPE = L;
```

The following code example changes more than one attribute at the same time:

```
ALTER ENGINE MyEngine SET NODES = 3 TYPE = L;
```

For more information on modifying engines, see [ALTER ENGINE](/sql_reference/commands/engines/alter-engine.html).

## [](#connecting-to-engines)Connecting to engines

You can connect to an engine using the following methods:

- Firebolt’s [user interface](https://go.firebolt.io/login).
- An engine URL.
- Third-party [connectors](/Guides/integrations/integrations.html) such as Airflow and DBT.

The engine URL is based on your account name and org name, with the following format:

`<account-name>.<org-name>.region.firebolt.io`

The combined length of `account-name` and `org-name` must not exceed 62 characters.

## [](#monitoring-engine-usage)Monitoring engine usage

You can use the observability views in `information_schema` to track engine performance and usage.

View Description `engine_metrics_history` Captures CPU and RAM usage every **30 seconds** and retains data for **30 days**. `engine_running_queries` Lists active queries and queries waiting to be processed.

You can use the information in the previous [information\_schema views](/sql_reference/information-schema/) to decide whether you need to change the engine configuration type, number of nodes or clusters based on your workload needs.

The [engine\_metrics\_history](/sql_reference/information-schema/engine-metrics-history.html) view gathers engine resource utilization metrics such as CPU and RAM consumption at a given time snapshot. Utilization snapshots are captured every 30 seconds and retained for 30 days, allowing users to understand engine utilization and consumption trends.

The following code example retrieves CPU and RAM usage for `MyEngine`:

```
SELECT * 
FROM information_schema.engine_metrics_history 
WHERE engine_name = 'MyEngine' 
ORDER BY event_time DESC;
```

The [engine\_running\_queries](/sql_reference/information-schema/engine-running-queries.html) view exposes information about queries currently running or waiting to be run in the system. Based on the number of queries that are queued and waiting to run, you can modify the engine configuration to best suit your performance requirements.

The following code example retrieves currently running queries:

```
SELECT * 
FROM information_schema.engine_running_queries;
```

If the previous query shows that queries remain in the queue for too long, increase the number of nodes or clusters.

To understand how information views can help with engine resizing, see [Working with Engines](/Guides/operate-engines/sizing-engines.html).

## [](#engine-governance-and-security)Engine governance and security

Firebolt provides **account-level isolation** and **Role Based Access Control (RBAC)** to provide strict governance over data access and infrastructure costs.

**Account isolation**

You can create multiple accounts within a given organization, where each account can represent a fully isolated environment such as development, test, or production. This enables engines across different environments to be fully isolated from each other.

**RBAC for engine management**

The Firebolt [RBAC model](/Guides/security/rbac.html) allows administrators to control user actions on resources that are created within a given account. For example, administrators can control which users are allowed to modify the configuration of engines and control which users can create new engines.

The following code example creates an engine administrator role and grants it full permissions on `MyEngine`:

```
CREATE ROLE engine_admin;
GRANT ALL PRIVILEGES ON ENGINE MyEngine TO engine_admin;
```

For more information on using RBAC for engines, see [Governing Engines](/Guides/operate-engines/rbac-for-engines.html).

## [](#viewing-and-understanding-engine-status)Viewing and understanding engine status

Use [SHOW ENGINES](/sql_reference/commands/metadata/show-engines.html) to list all engines and their statuses in your Firebolt account as follows:

`SHOW ENGINES` and UI Description STARTING The engine is provisioning resources, and will be ready to use soon. RUNNING The engine is running queries or available to run queries. The engine can be modified while it is running. RESIZING The engine is currently being resized after an `ALTER ENGINE` command. The engine will be in this state when the user has issued a request to change the engine `TYPE`, number of `NODES` or number of `CLUSTERS`. DRAINING The engine is completing running queries before shutting down. STOPPING The engine is shutting down and cannot accept new queries. STOPPED The engine is fully stopped and not available to run queries.