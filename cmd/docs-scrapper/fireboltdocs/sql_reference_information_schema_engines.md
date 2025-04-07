# [](#information-schema-for-engines)Information schema for engines

You can use the `information_schema.engines` view to return information about each engine in an account. The view is available for each database and contains one row for each engine in the account. You can use a `SELECT` query to return information about each engine as shown in the example below, which uses a `WHERE` clause to return all engines attached to databases that begin with `deng`.

To view engine information, the user must have ownership of the engine or access to the necessary [engine](/Overview/Security/Role-Based%20Access%20Control/engine-permissions.html#engine-permissions) privileges.

```
SELECT
  *
FROM
  information_schema.engines
WHERE
  attached_to ILIKE 'deng%'
```

## [](#columns-in-information_schemaengines)Columns in information\_schema.engines

Each row has the following columns with information about each engine.

Column Name Data Type Description engine\_name TEXT The name of the engine. region TEXT \[DEPRECATED] The region in which the engine was created. spec TEXT \[DEPRECATED] The specification of nodes comprising the engine. scale INTEGER \[DEPRECATED] The number of nodes in the engine. type TEXT Determines the capability of the nodes in the engine. family TEXT The family of a given engine. Choose from `STORAGE_OPTIMIZED` or `COMPUTE_OPTIMIZED`. nodes INTEGER The number of nodes in a cluster. clusters INTEGER The number of node groupings in an engine. status TEXT The engine status. For more information, see [Viewing and understanding engine status](/Overview/engine-fundamentals.html#viewing-and-understanding-engine-status). attached\_to TEXT \[DEPRECATED] The name of the database to which the engine is attached. auto\_start BOOLEAN When true, queries issued to a stopped engine will attempt to start the engine first. auto\_stop INTEGER Indicates the amount of time (in minutes) after which the engine automatically stops. engine\_type TEXT \[DEPRECATED] The type of the engine. initially\_stopped BOOLEAN When true, the engine will have attempted to start after creation. url TEXT A url which can be used to issue queries to this engine. warmup TEXT \[DEPRECATED] The warmup method of the engine. default\_database TEXT The database an engine will attempt to use by default when dealing with queries that require a database. version TEXT The engine version. last\_started TIMESTAMPTZ The last time this engine was started (UTC). last\_stopped TIMESTAMPTZ The last time this engine was stopped (UTC). description TEXT A user defined description for the engine. created TIMESTAMPTZ The time when this engine was created (UTC). engine\_owner TEXT The name of the user who created the engine. last\_altered\_by TEXT The user who last altered this engine. last\_altered TIMESTAMPTZ The time when this engine was last altered (UTC). fbu\_rate NUMERIC Hourly FBU consumption rate of engines based on engine topology at the time the view is invoked by the user.