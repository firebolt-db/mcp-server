# [](#alter-engine)ALTER ENGINE

Modify the configuration of an existing Firebolt engine.

## [](#alter-engine-set)ALTER ENGINE SET

Updates the configuration of the specified engine.

### [](#syntax)Syntax

```
ALTER ENGINE <engine_name> SET
    [AUTO_START = <true/false>]
    [AUTO_STOP = <minutes>]
    [DEFAULT_DATABASE = <database_name>]
    [CLUSTERS = <clusters>]
    [MIN_CLUSTERS = <clusters>]
    [MAX_CLUSTERS = <clusters>]
    [NODES = <nodes>]
    [TYPE = <type>]
    [FAMILY = <family>]
```

### [](#options)Options

Parameter Description `<engine_name>` The name of the engine to be altered. `AUTO_START = <true/false>` When set to `TRUE`, sending a query to a stopped engine will automatically start the engine before processing the query. `AUTO_STOP = <minutes>` Specifies the number of minutes after which an engine automatically stops. Setting `minutes` to `0` disables `AUTO_STOP`. `DEFAULT_DATABASE = <database_name>` Specifies the default database that an engine will attempt to use when processing queries that require a database. To remove the default database, set `DEFAULT_DATABASE=default`. `CLUSTERS = <clusters>` Specifies the number of clusters in an engine. Each cluster is a group of nodes, and all clusters within an engine are identical in terms of node type and number of nodes. `MIN_CLUSTERS = <clusters>` Specifies a minimum number of clusters in an engine. If `MIN_CLUSTERS` is different from `MAX_CLUSTERS`, the engine will automatically change the number of clusters depending on load. If not specified, `1` is used as default. `MAX_CLUSTERS = <clusters>` Specifies a maximum number of clusters in an engine. If `MIN_CLUSTERS` is different from `MAX_CLUSTERS`, the engine will automatically change the number of clusters depending on load. If not specified, `1` is used as default. `NODES = <nodes>` Specifies the number of nodes in each cluster within an engine. You can specify any integer between `1` to `128`, inclusive. `TYPE =<type>` Specifies the node type for the engine. You can choose `S`, `M`, `L`, or `XL`. `FAMILY =<family>` Defines the family of node used in the engine. Options include `STORAGE_OPTIMIZED` (alias `SO`) or `COMPUTE_OPTIMIZED` (alias `CO`).

**Limitations:**

- Each engine is limited to a maximum of two clusters.
- Each cluster can have up to ten nodes.
- The total number of nodes across all clusters cannot exceed 15.
- When scaling a running engine either vertically or horizontally, new queries will be directed to the new cluster. Queries running on the original clusters will continue until completion. The clusters will wait up to 24 hours for these queries to finish, after which any unfinished queries may be stopped.
- Only small and medium engines are available for use right away.

If you would like to remove the following limitations, reach out to Firebolt support at [support@firebolt.io](mailto:support@firebolt.io) to do the following:

- Use a large or extra-large engine.
- Use more than ten nodes per cluster.
- Use more than 15 nodes across all clusters.

## [](#examples)Examples

### [](#scale-out-an-engine)Scale out an engine

The following code example scales out, or increases the number of nodes, in an engine by setting the engine’s `NODES` to `3`:

```
ALTER ENGINE my_engine SET NODES = 3;
```

### [](#scale-up-an-engine)Scale up an engine

The following code example scales up an engine by increasing its capacity from small (`S`) to large (`L`) by setting the engine’s `TYPE` parameter to `L`:

```
ALTER ENGINE my_engine SET TYPE = "L";
```

### [](#scale-up-and-out-an-engine)Scale up and out an engine

The following code example both scales up and scales out an engine by increasing node capacity and the number of nodes:

```
ALTER ENGINE my_engine SET TYPE = "L" NODES = 5;
```

If you need to use a large or extra-large engine, reach out to [support@firebolt.io](mailto:support@firebolt.io).

### [](#enable-concurrency-auto-scaling)Enable concurrency auto-scaling

The following code example adjusts the maximum number of engine clusters to enable concurrency auto-scaling:

```
ALTER ENGINE my_engine SET MAX_CLUSTERS = 2;
```

## [](#alter-engine-rename-to)ALTER ENGINE RENAME TO

Renames an engine.

### [](#syntax-1)Syntax

```
ALTER ENGINE <engine_name> RENAME TO <new_engine_name>
```

### [](#parameters)Parameters

Parameter Description `<engine_name>` The name of the engine to rename. `<new_engine_name>` The new name of the engine.

## [](#alter-engine-owner-to)ALTER ENGINE OWNER TO

Changes the owner of an engine. You can view the current owner in the the `engine_owner` column of the `information_schema.engines` view.

For more information, see [ownership](/Guides/security/ownership.html).

### [](#syntax-2)Syntax

```
ALTER ENGINE <engine_name> OWNER TO <user>
```

### [](#parameters-1)Parameters

Parameter Description `<engine_name>` The name of the engine to change the owner of. `<user>` The new owner of the engine.