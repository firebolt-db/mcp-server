# [](#show-engines)SHOW ENGINES

Returns a table with a row for each Firebolt engine defined in the current Firebolt account, with columns containing information about each engine as listed below.

## [](#syntax)Syntax

```
SHOW ENGINES;
```

## [](#returns)Returns

The returned table has the following columns.

Column name Data Type Description engine\_name TEXT The name of the engine. engine\_owner TEXT (5) Name of the user who created the engine. type TEXT The specification of nodes comprising the engine. clusters INT Collection of nodes, where each node is of a certain type. All the clusters in an engine have the same type and same number of nodes. nodes INT The number of nodes for each cluster in an engine. Can be an integer ranging from `1` to `128`. status TEXT The engine status. For more information, see [Viewing and understanding engine status](/Overview/engine-fundamentals.html#viewing-and-understanding-engine-status) auto\_stop INT The amount of idle time (in minutes) after which the engine automatically stops. url TEXT Engine endpoint. version TEXT The engine version. initially\_stopped BOOLEAN If `false`, engine was started as part of the `CREATE ENGINE` command. default\_database TEXT The database an engine will attempt to use by default when dealing with queries that require a database.

If not specified, `NULL` is used as default. created TIMESTAMPTZ Creation time of the engine. last\_altered\_by TEXT Name of the last user who edited the engine. last\_altered TIMESTAMPTZ Last modified time of the engine. fbu\_rate INT Hourly FBU consumption rate of running engines based on engine topology at the time the view is invoked by the user.