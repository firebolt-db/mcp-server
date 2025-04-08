# [](#create-view)CREATE VIEW

Creates a view, which is useful to filter, focus, and simplify a data set for querying. Views provide a level of abstraction that can make subqueries easier to write, especially for commonly referenced subsets of data.

View results are not stored for future usage. Each time a query references a view, the view runs its `SELECT` query. For this reason, views do not provide a performance advantage. Consider using a materialized common table expression (CTE) as an alternative. For more information, see [Materialized common table expressions](/sql_reference/commands/queries/select.html#materialized-common-table-expressions).

The optional `IF NOT EXISTS` and `OR REPLACE` clauses are mutually exclusive. They specify behavior if a view with the same name already exists. If neither clause is specified, an error occurs if a view with the same `<view_name>` already exists.

## [](#syntax)Syntax

```
CREATE VIEW [IF NOT EXISTS] <view_name> [(<column_list>)]
AS SELECT <select_statement>
```

**—OR—**

```
CREATE [OR REPLACE] VIEW <view_name> [(<column_list>)]
AS SELECT <select_statement>
```

## [](#parameters)Parameters

Parameter Description `IF NOT EXISTS` Specifies that an existing view with `<view_name>` will not be replaced with this view definition. This is the default behavior, but this clause suppresses the error message that would otherwise occur, which is useful for scripted and programmatic implementations. `OR REPLACE` Specifies that an existing view with `<view_name>` will be replaced with this view definition. `<view_name>` An identifier that specifies the name of the view. This name must be unique within the database. `<column_list>` An optional list of column names to be used for columns of the view. If not specified, the column names are deduced from the query. `<select_statement>` The select statement for creating the view.

## [](#example)Example

```
CREATE VIEW fob_shipments
AS SELECT   l_shipmode,
            l_shipdate,
            l_linestatus,
            l_orderkey,
FROM    lineitem
WHERE   l_shipdate > 1990-01-01
AND     l_shipmode = 'FOB';
```