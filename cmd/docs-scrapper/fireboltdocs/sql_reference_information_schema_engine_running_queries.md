# [](#information-schema-for-engine-running-queries)Information schema for engine running queries

You can use the `information_schema.engine_running_queries` view to return information about queries currently running in a database. This view is available in each database and includes one row for each running query. The table contains a maximum of ten thousand queries per engine cluster. You can use a `SELECT` query to return information about each running query as shown in the example below.

```
SELECT
  *
FROM
   information_schema.engine_running_queries
LIMIT
  100;
```

## [](#columns-in-information_schemaengine_running_queries)Columns in information\_schema.engine\_running\_queries

Each row has the following columns with information about each running query.

Column Name Data Type Description account\_name TEXT The name of the account in which the query was executed. user\_name TEXT The user that executed the query submitted\_time TIMESTAMPTZ The time when the query was submitted by a user (UTC) start\_time TIMESTAMPTZ The query execution start time (UTC). duration\_us BIGINT The time elapsed in microseconds from `<START_TIME>` to when the query on `information_schema.engine_running_queries` returns results. e2e\_duration\_us BIGINT The time elapsed in microseconds from the time of query submission to when the query on `information_schema.engine_running_queries` returns results. status TEXT The status of the query, eiteher `RUNNING`, `SUSPENDED`, or `CANCELING`. request\_id TEXT The ID of the request from which the query originates. query\_id TEXT The query id of this query. query\_label TEXT User provided query label (query\_label parameter) query\_text TEXT Text of the SQL statement. scanned\_rows BIGINT The number of rows scanned to return query results. scanned\_bytes BIGINT The number of bytes scanned from cache and storage. inserted\_rows BIGINT The number of rows written inserted\_bytes BIGINT The number of bytes written. async\_token TEXT If the query is an async query, this handle can be used to check the status via the built in stored procedure fb\_GetAsyncStatus(). retries BIGINT The total number of retries to execute a given query after a failure (by default, the number of retries is 0 and the number increases with each retry)