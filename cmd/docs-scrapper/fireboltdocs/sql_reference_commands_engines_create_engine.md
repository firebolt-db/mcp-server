# [](#create-engine)CREATE ENGINE

Creates an engine, which serves as a compute cluster for running queries and processing data.

## [](#syntax)Syntax

```
CREATE ENGINE [IF NOT EXISTS] <engine_name>
[WITH 
    [AUTO_START = <true/false>]
    [AUTO_STOP = <minutes>]
    [DEFAULT_DATABASE = <database_name>]
    [INITIALLY_STOPPED = <true/false>]
    [START_IMMEDIATELY = <true/false>]
    [CLUSTERS = <clusters>]
    [MIN_CLUSTERS = <clusters>]
    [MAX_CLUSTERS = <clusters>]
    [NODES = <nodes>]
    [TYPE = <type>]
    [FAMILY = <family>]]
```

## [](#options)Options

Parameter Description `<engine_name>` The name of the engine to be created. `AUTO_START = <true/false>` When `true`, sending a query to a stopped engine will start the engine before processing the query.

If not specified, `true` will be used as default. `AUTO_STOP = <minutes>` The amount of idle time (in minutes) after which the engine automatically stops.  
Setting the minutes to `0` indicates that `AUTO_STOP` is disabled.

If not specified, `20` is used as default. `DEFAULT_DATABASE = <database_name>` The database an engine will attempt to use by default when dealing with queries that require a database.

If not specified, `NULL` is used as default. `INITIALLY_STOPPED = <true/false>` When `false`, the newly created engine will be started as part of the `CREATE ENGINE` command.  
Cannot be used with `START_IMMEDIATELY`.

If not specified, `false` is used as default. `START_IMMEDIATELY = <true/false>` When `true`, the newly created engine will be started as part of the `CREATE ENGINE` command.  
Cannot be used with `INITIALLY_STOPPED`.

If not specified, `true` is used as default. `CLUSTERS = <clusters>` Specifies the number of clusters in an engine. Each cluster is a group of nodes, and all clusters within an engine are identical in terms of node type and number of nodes.

If not specified, `1` is used as default. `MIN_CLUSTERS = <clusters>` Specifies a minimum number of clusters in an engine. If `MIN_CLUSTERS` is different from `MAX_CLUSTERS`, the engine will automatically change the number of clusters depending on load.

If not specified, `1` is used as default. `MAX_CLUSTERS = <clusters>` Specifies a maximum number of clusters in an engine. If `MIN_CLUSTERS` is different from `MAX_CLUSTERS`, the engine will automatically change the number of clusters depending on load.

If not specified, `1` is used as default. `NODES = <nodes>` Indicates the number of nodes in each cluster within an engine. This number can range from `1` to `128`.

If not specified, `2` is used as default. `TYPE =<type>` Defines the type of node used in the engine. Options include `S`, `M`, `L`, or `XL`

If not specified, `M` is used as default. `FAMILY =<family>` Defines the family of node used in the engine. Options include `STORAGE_OPTIMIZED` (alias `SO`) or `COMPUTE_OPTIMIZED` (alias `CO`)

If not specified, `STORAGE_OPTIMIZED` is used as default.

**Limitations:**

- The number of clusters per engine is limited to two.
- The number of nodes per cluster is limited to ten.
- The total number of nodes x clusters cannot exceed 15.
- Only small and medium engines are available for use right away.

If you would like to remove the above limitations or use a large or extra-large engine, reach out to Firebolt Support at [support@firebolt.io](mailto:support@firebolt.io).

## [](#examples)Examples

### [](#create-a-basic-engine)Create a basic engine

The following example creates an engine with one cluster, using node type ‘S’ and 1 nodes per cluster :

```
CREATE ENGINE my_engine;
```

### [](#create-an-engine-with-multiple-nodes)Create an engine with multiple nodes

The following example creates an engine with one cluster, using node type ‘S’ and 5 nodes per cluster :

```
CREATE ENGINE my_engine
WITH TYPE="S" NODES = 5 CLUSTERS = 1;
```

### [](#create-an-engine-with-delayed-start)Create an engine with delayed start

The following example creates an engine with one cluster, using node type ‘M’ and 3 nodes per cluster. The engine will not be automatically started after creation because `INITIALLY_STOPPED` is set to true.

```
CREATE ENGINE my_engine
WITH TYPE="M" NODES=3 INITIALLY_STOPPED=true;
```

### [](#create-an-engine-with-auto-stop)Create an engine with auto-stop

The following example creates an engine with one cluster, using node type ‘L’ in the compute-optimized family with 2 nodes per cluster. The engine will be automatically stopped after 10 minutes of idle time because `AUTO_STOP` is set to 10.

```
CREATE ENGINE my_engine
WITH TYPE="L" FAMILY="COMPUTE_OPTIMIZED" NODES=2 AUTO_STOP=10;
```

If you need to use a large or extra-large engine, reach out to [support@firebolt.io](mailto:support@firebolt.io).

### [](#create-an-engine-with-concurrency-auto-scaling)Create an engine with concurrency auto-scaling

The following example creates an engine that automatic concurrency scaling. The engine adjusts the number of clusters based on workload demand, maintaining between one and two clusters.

```
CREATE ENGINE my_engine
WITH MIN_CLUSTERS=1 MAX_CLUSTERS=2;
```