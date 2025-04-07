# [](#information-schema-for-engine-metering-history)Information schema for engine metering history

You can use the `information_schema.engine_metering_history` view to return information about each engine’s FBU consumed over time in an account. See [Engine Consumption](/Overview/engine-consumption.html) for more details on consumption, nodes, and FBU. It is often useful to filter this view to a particular engine. In the example below, a filter is applied to look at the history of engines starting with `capacity_test_`. By default, shows hourly data from the last 30 days.

```
SELECT
  *
FROM
  information_schema.engine_metering_history
WHERE
  engine_name LIKE 'capacity_test_%'
```

## [](#columns-in-information_schemaengine_metering_history)Columns in information\_schema.engine\_metering\_history

Each row has the following columns with information about each engine.

Column Name Data Type Description engine\_name TEXT The last known name of the engine. Either reflects the current name or the name at time of deletion. start\_hour TIMESTAMPTZ Start time for the hour. end\_hour TIMESTAMPTZ End time for the hour. consumed\_FBU TEXT Firebolt units consumed within the hour based on the the duration the engine has been running.