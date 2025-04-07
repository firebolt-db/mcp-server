# [](#information-schema-for-engine-history)Information schema for engine history

You can use the `information_schema.engine_history` view to return information about each engine’s history in an account. The view shows operations performed on each engine, including creation, deletion, starts, stops, resizing, and logical operations like renames. It is often useful to filter this view to a particular engine. In the example below, a filter is applied to look at the history of engines starting with `capacity_test_`. By default, shows the events from the last 30 days.

```
SELECT
  *
FROM
  information_schema.engine_history
WHERE
  engine_name LIKE 'capacity_test_%'
```

## [](#columns-in-information_schemaengines_history)Columns in information\_schema.engines\_history

Each row has the following columns with information about each engine.

Column Name Data Type Description engine\_name TEXT The last known name of the engine. Either reflects the current name or the name at time of deletion. engine\_owner TEXT The name of the user who owns the engine. cluster\_id INT Ordinal numbers to identify engine clusters. type TEXT(5) Node type used in a given engine (S, M, L or XL). family TEXT The family of a given engine. Choose from `STORAGE_OPTIMIZED` or `COMPUTE_OPTIMIZED`. nodes INT Number of nodes in each of the cluster of the engine. clusters INT The number of clusters used in the engine. auto\_start BOOLEAN If True, automatically start the engine if in stopped state when a query comes in. auto\_stop INT Automatically stop the engine after specified number of minutes. initially\_stopped BOOLEAN If True, the engine will not be automatically started after creation. url TEXT Engine URL used by the users to issue queries to the engine. default\_database TEXT Default database for the engine as specified by the user. version TEXT Engine version. event\_type TEXT Name of the event. event\_reason TEXT Reason why the engine event was triggered. event\_status TEXT Status of the event, indicating whether the action corresponding to the event has succeeded, failed or is in process. Can have one of the following values: SUCCEEDED, FAILED or IN\_PROGRESS. event\_start\_time TIMESTAMPTZ Time when the event was initiated. event\_finish\_time TIMESTAMPTZ Time when the event was completed. user\_name TEXT User who triggered the event. description TEXT Description of the engine as specified by the user. query\_id TEXT Unique identifier for the SQL query used for engine operations.