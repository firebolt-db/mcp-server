# [](#information-schema-for-engine-metrics-history)Information schema for engine metrics history

You can use the `information_schema.engine_metrics_history` view to return information about the resource consumption of your engine. As engines can have 1 or more engine clusters that are transient, this provides visibilty into the utilization of each engine cluster. Each row represents the aggregate resource metrics for a single engine cluster at the given event time. Per default, we collect 2 metric snapshots per minute for the last 30 days. You can use a `SELECT` query to return the metrics for each engine cluster at a given point in time as shown in the example below.

```
SELECT
  *
FROM
  information_schema.engine_metrics_history
WHERE
  event_time > now() - INTERVAL '2 hours'
LIMIT 100;
```

## [](#columns-in-information_schemaengine_metrics_history)Columns in information\_schema.engine\_metrics\_history

Each row has the following columns with information about each engine cluster.

Column Name Data Type Description engine\_ordinal BIGINT the cluster number event\_time TIMESTAMPTZ timestamp at which the metrics where collected cpu\_used DECIMAL current CPU utilization (percentage) memory\_used DECIMAL current memory used (percentage) disk\_used DECIMAL currently used disk space which encompasses space used for cache and spilling (percentage) cache\_hit\_ratio DECIMAL current SSD cache hit ratio spilled\_bytes BIGINT amount of spilled data to disk in bytes running\_queries BIGINT number of currently running queries in the system suspended\_queries BIGINT number of queries that have been suspended (not-running and awaiting to be executed)