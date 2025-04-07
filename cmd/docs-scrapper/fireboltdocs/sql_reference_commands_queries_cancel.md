# [](#cancel-query)CANCEL QUERY

The `CANCEL QUERY` statement is used to cancel a running query within the Firebolt engine. This statement provides a mechanism to terminate the execution of a specific query identified by its unique `query_id`. Using [query labels](/Reference/system-settings.html#query-labelingtagging) makes it easier to retrieve the query id of your query.

## [](#syntax)Syntax

```
CANCEL QUERY WHERE query_id = '<query_id>';
```

## [](#parameters)Parameters

Parameter Description `<query_id>` A unique identifier assigned to the query that needs to be canceled.

## [](#example)Example

```
CANCEL QUERY WHERE query_id = '12345';
```

## [](#notes)Notes

- The query\_id is typically obtained from the `information_schema.engine_running_queries` view.
- When a query is canceled, it terminates the processes of that specific query. This action both frees up resources and prevents any further execution of the query.
- This statement is designed for administrative use and may require appropriate privileges to execute. However, all users can cancel their own queries by default.

## [](#example-use-case)Example Use Case

Consider a scenario where there is a long-running query that needs to be canceled to free up system resources. The following statement can be used to obtain the query\_id:

```
SELECT status, duration_us, query_text, query_id FROM information_schema.engine_running_queries;
```

Suppose the long-running query is identified with query\_id ‘12345’:

```
CANCEL QUERY WHERE query_id = '12345';
```

This will initiate the cancellation process for the specified query\_id. While the query is not immediately removed from running\_queries, its status will be observed as `CANCELING` there. After the cancellation process is completed, the query will have the `CANCELED` status in the `information_schema.engine_query_history` view:

```
SELECT query_id, query_text FROM information_schema.engine_query_history WHERE status = 'CANCELED';
```