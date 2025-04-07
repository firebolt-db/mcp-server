# [](#information-schema-for-engine-query-history)Information schema for engine query history

You can use the `information_schema.engine_query_history` view to return information about queries saved to query history. The view is available in each database and contains two rows, the starting and ending row for each historical query in the database. The table includes the last ten thousand queries per engine cluster. You can run a `SELECT` query to retrieve details about recent queries, as shown in the following example:

```
SELECT
  *
FROM
  information_schema.engine_query_history
LIMIT
  100;
```

The `information_schema.engine_query_history` view retains only the most recent 10,000 queries. Queries exceeding this limit are excluded and will not appear in the query history. This limitation is important for high-volume workloads and tools like the [OTEL exporter](/Guides/integrations/otel-exporter.html), which can rapidly fill the query history. To retain critical query data, regularly export or archive query history.

## [](#columns-in-information_schemaengine_query_history)Columns in information\_schema.engine\_query\_history

Each row has the following columns with information about each query in query history.

Column Name Data Type Description   account\_name TEXT The name of the account that ran the query.   user\_name TEXT The user name that was used to run the query. The `user_name` is present for account-level operations, and `NULL` for organization-level operations.   login\_name TEXT The login name that was used to run the query. The `login_name` is present for organization-level statements, and otherwise `NULL`.   service\_account\_name TEXT The service account name that was used to run the query. The `service_account_name` is present for organization-level statements, and otherwise `NULL`.   submitted\_time TIMESTAMPTZ The time that the user submitted the query.   start\_time TIMESTAMPTZ The time that the query started running in Coordinated Universal Time (UTC).   end\_time TIMESTAMPTZ The time that the query stopped running in UTC.   duration\_us BIGINT The duration of query run time in microseconds.   e2e\_duration\_us BIGINT The end-to-end duration of query run time. Starting from the time the query was submitted and ending when the result was fully returned in microseconds.   status TEXT Can be one of the following values:  
`STARTED_EXECUTION`–Successful start of query execution.  
`ENDED_SUCCESSFULLY`–Successful end of query execution.  
`CANCELED_EXECUTION`–Query was canceled.  
`PARSE_ERROR`–Query could not be parsed.  
`EXECUTION_ERROR`–Query could not be executed successfully.   request\_id TEXT The ID of the request from which the query originates.   query\_id TEXT The unique identifier of the SQL query.   query\_label TEXT A user-provided query label.   query\_text TEXT The text of the SQL statement.   query\_text\_normalized TEXT The normalized text of the SQL statement.   query\_text\_normalized\_hash TEXT The hash of the normalized text of the SQL statement.   error\_message TEXT The returned error message.   scanned\_rows BIGINT The total number of rows scanned.   scanned\_bytes BIGINT The total number of uncompressed bytes scanned.   scanned\_cache\_bytes BIGINT The total number of compressed bytes scanned from disk-based cache.   scanned\_storage\_bytes BIGINT The total number of compressed bytes scanned from Firebolt-managed storage. Does not apply to [EXTERNAL tables](/Overview/indexes/using-indexes.html#external-tables).   inserted\_rows BIGINT The total number of rows written.   inserted\_bytes BIGINT The total number of bytes written to both cache and storage.   spilled\_bytes BIGINT The total number of uncompressed bytes spilled.   returned\_rows BIGINT The total number of rows returned from the query.   returned\_bytes BIGINT The total number of bytes returned from the query.   time\_in\_queue\_us BIGINT The number of microseconds the query spent in queue.   retries BIGINT The number of retried attempts in case of query failure. Defaults to 0.