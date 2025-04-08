# [](#sql-pipe-syntax)SQL Pipe syntax

Firebolt supports SQL Pipe syntax, an alternative SQL syntax that uses the `|>` operator to chain query transformations step by step. SQL Pipe syntax supports the same functionality as standard SQL, but can improve readability by allowing queries to flow in a linear, top-to-bottom structure, which makes it easier to express, compose and understand queries. This syntax was first presented by Google in the research paper [SQL Has Problems. We Can Fix Them: Pipe Syntax In SQL](https://research.google/pubs/sql-has-problems-we-can-fix-them-pipe-syntax-in-sql/).

SQL Pipe syntax has the following structure:

- Each pipe operator consists of the pipe symbol, `|>`, an operator name, and arguments: `|> operator_name arguments`.
- You can add pipe operators to the end of any valid query.
- You can apply pipe operators in any order, and apply them as many times as needed.
- Pipe syntax can be used anywhere that standard syntax is used including in queries, subqueries, and views.
- You can combine pipe syntax with standard SQL syntax in the same query. For example, a parent query can use standard SQL while the subquery uses pipe syntax, or the other way around.

For more information, see Google’s original research paper, [SQL Has Problems. We Can Fix Them: Pipe Syntax In SQL](https://research.google/pubs/sql-has-problems-we-can-fix-them-pipe-syntax-in-sql/).

### [](#example)Example

The following code example is taken from [Query 13 in the TPC-H benchmark suite](https://github.com/apache/impala/blob/master/testdata/workloads/tpch/queries/tpch-q13.test), a standard set of queries used to measure database performance. Query 13 specifically analyzes customer order patterns to identify how many customers fall into different order-count categories. In **standard SQL**, the following query counts how many customers have placed a certain number of orders, excluding special requests, and sorts the results by customer count and order count:

```
SELECT c_count, COUNT(*) AS custdist
  FROM
  (SELECT c_custkey, COUNT(o_orderkey) c_count
    FROM customer
    LEFT OUTER JOIN orders ON c_custkey = o_custkey
                           AND o_comment NOT LIKE '%special%requests%'
   GROUP BY c_custkey
) AS c_orders
GROUP BY c_count
ORDER BY custdist DESC, c_count DESC;
```

The following query shows the previous query rewritten using pipe syntax:

```
FROM customer
|> LEFT OUTER JOIN orders ON c_custkey = o_custkey
|> WHERE o_comment NOT LIKE '%special%requests%'
|> AGGREGATE COUNT(o_orderkey) AS c_count  GROUP BY c_custkey
|> AGGREGATE COUNT(*)          AS custdist GROUP BY c_count
|> ORDER BY custdist DESC, c_count DESC;
```

## [](#from-first)`FROM` first

Firebolt supports [`FROM` first queries](/sql_reference/commands/queries/select.html#from-first), which are especially useful with pipe syntax, which builds queries in a step-by-step flow, with each step transforming the result of the previous one. Starting with the `FROM` clause defines the initial data source, allowing you to clearly chain joins, filters, and transformations. A pipe query typically starts with a `FROM` clause that references a table, view, subquery, join, or table-valued function.

### [](#examples)Examples

**Example**

The following code example selects all data from the `levels` table:

```
FROM levels
```

**Example**

The following code example generates a series of numbers from 1 to 5:

```
FROM generate_series(1,5);
```

## [](#pipe-operators)Pipe operators

- [SELECT](#select-pipe-operator)
- [EXTEND](#extend-pipe-operator)
- [AS](#as-pipe-operator)
- [WHERE](#where-pipe-operator)
- [LIMIT](#limit-pipe-operator)
- [AGGREGATE](#aggregate-pipe-operator)
- [ORDER BY](#order-by-pipe-operator)
- [JOIN](#join-pipe-operator)

### [](#select-pipe-operator)`SELECT` pipe operator

The `SELECT` pipe operator works similarly to the standard SQL `SELECT` statement, and is used to specify the columns or expressions you want in the result, including column references, scalar expressions, `*` wildcards, and clauses like `DISTINCT` or `EXCLUDE`. The `SELECT` pipe operator enables flexible data selection within the pipe query flow.

#### [](#syntax)Syntax

```
|> SELECT [ ALL | DISTINCT ] <select_list>
```

The `<select_list>` in the `SELECT` pipe operator follows the same syntax as the [`SELECT` statement](/sql_reference/commands/queries/select.html#select-1) in standard SQL. Specifically, it can contain components including column references, scalar expressions, `*` wildcards, and the `EXCLUDE` clause.

**Example**

The following code example selects the `level`, `name`, `maxpoints`, and `pointsperlap` columns from the `levels` table and computes the number of laps by dividing `maxpoints` by `pointsperlap`:

```
FROM levels
|> SELECT level, name, maxpoints, pointsperlap, maxpoints / pointsperlap AS number_of_laps
```

### [](#extend-pipe-operator)`EXTEND` pipe operator

The `EXTEND` pipe operator adds new computed columns to query results while keeping the existing columns so that you can perform calculations or transformations on your data without losing any original columns. This is useful when you want to enrich your results with additional insights.

#### [](#syntax-1)Syntax

```
|> EXTEND <expression> [AS <alias>] [, ...]
```

The `EXTEND` operator propagates all the columns from the input and adds computed columns, similar to the `SELECT *, <expression> [AS <alias>]` syntax in standard SQL. These computed columns can include scalar expressions or window functions defined using the format `|> EXTEND <expression> [AS <alias>] [, ...]`.

**Example**

The following code example adds two computed columns to the `levels` table: `number_of_laps`, calculated by dividing `maxpoint`s by `pointsperlap`, and `total_max_points`, which is the cumulative sum of `maxpoints` ordered by `level`:

```
FROM levels
|> EXTEND 
     maxpoints / pointsperlap AS number_of_laps,
     SUM(maxpoints) OVER (ORDER BY level ASC) AS total_max_points
```

### [](#as-pipe-operator)`AS` pipe operator

The `AS` pipe operator is used to rename tables or columns within a query, similar to the `AS` keyword in standard SQL. You can use it to assign aliases to tables or columns, and make your query more readable and simplify complex operations. The `AS` pipe operator can help clarify the intent of your query by giving meaningful names to columns or tables, improving its structure and understanding within the pipe query flow.

#### [](#syntax-2)Syntax

```
|> AS <table_alias> | <table_alias>(<column1_alias>, ...)
```

**Example**

The following code example generates a series of numbers from 1 to 5, aliases the result as `table(i)`, and selects the `i` column from the aliased table:

```
FROM generate_series(1,5)
|> AS table(i)
|> SELECT table.i
```

### [](#where-pipe-operator)`WHERE` pipe operator

The `WHERE` pipe operator filters the input data based on a specified condition, similar to the [WHERE](/sql_reference/commands/queries/select.html#where) clause in standard SQL, and also replaces the `HAVING` clause in pipe syntax, allowing you to exclude rows that don’t meet the criteria.

#### [](#syntax-3)Syntax

```
|> WHERE <condition>
```

**Example**

The following code example filters the `levels` table to include only rows where the `leveltype` is `FastestLap`:

```
FROM levels
|> WHERE leveltype = 'FastestLap'
```

### [](#limit-pipe-operator)`LIMIT` pipe operator

The `LIMIT` pipe operator restricts the number of rows in the result set, similar to the [LIMIT clause](/sql_reference/commands/queries/select.html#limit) in standard SQL, and can optionally use `OFFSET` to skip a specified number of rows, useful for limiting large result sets, pagination, or testing smaller data subsets.

#### [](#syntax-4)Syntax

```
|> LIMIT <count> [OFFSET <start>]
```

**Example**

The following code example generates a series of numbers from 1 to 1000 and limits the result to the first 3 rows:

```
FROM generate_series(1,1000)
|> LIMIT 3
```

### [](#aggregate-pipe-operator)`AGGREGATE` pipe operator

You can use the `AGGREGATE` pipe operator to perform either **full table aggregation** or **aggregation across groups**, similar to using the [`GROUP BY`](/sql_reference/commands/queries/select.html#group-by) clause in standard SQL. Use this operator to apply aggregate functions like `SUM`, `AVG`, or `COUNT` to grouped data.

Unlike SQL, where grouping expressions need to be repeated in both the `SELECT` and `GROUP BY` clauses, in pipe syntax, grouping expressions are listed only once in the `GROUP BY` clause and are automatically included in the output columns. The `AGGREGATE` operator’s output first includes the grouping expressions, followed by the aggregated expressions, using their assigned aliases as column names.

#### [](#syntax-for-full-table-aggregation)Syntax for full table aggregation

```
|> AGGREGATE <aggregate_expression> [AS <alias>] [, ...]
```

#### [](#syntax-for-aggregation-with-grouping)Syntax for aggregation with grouping

```
|> AGGREGATE [<aggregate_expression> [AS <alias>] [, ...]] GROUP BY <grouping_expression> [, ...]
```

**Examples**

The following code example calculates the total sum of `maxpoints` from the `levels` table without grouping:

```
FROM levels
|> AGGREGATE SUM(maxpoints)
```

The following code example calculates the maximum `maxpoints` for each `leveltype` and returns the corresponding `level` for each maximum value, grouping the results by `leveltype`:

```
FROM levels
|> AGGREGATE MAX_BY(level, maxpoints) level, MAX(maxpoints) maxpoints GROUP BY leveltype
```

### [](#order-by-pipe-operator)`ORDER BY` pipe operator

Use the `ORDER BY` pipe operator to sort the input data based on one or more expressions, similar to the [`ORDER BY`](/sql_reference/commands/queries/select.html#order-by) clause in standard SQL to organize query results in a meaningful order. You can use `ORDER BY` to specify the sorting order with options for ascending (`ASC`) or descending (`DESC`) order, as well as the handling of `NULL` values with `NULLS FIRST` or `NULLS LAST`.

#### [](#syntax-5)Syntax

```
|> ORDER BY <expression> [ ASC | DESC ] [ NULLS FIRST | NULLS LAST] [, ...]
```

**Example**

The following code example sorts the `levels` table by `numberoflaps` in descending order and then selects the `level` and `numberoflaps` columns from the sorted result:

```
FROM levels
|> ORDER BY numberoflaps DESC
|> SELECT level, numberoflaps
```

### [](#join-pipe-operator)`JOIN` pipe operator

Use the `JOIN` pipe operator to combine rows from two or more tables based on a related column and merge datasets in a query similar to the standard SQL [`JOIN`](/sql_reference/commands/queries/select.html#join) clause. You can perform different types of joins, such as `INNER`, `LEFT`, `RIGHT`, and `FULL`, to retrieve and merge data based on specific conditions.

#### [](#syntax-6)Syntax

```
|> [ INNER | LEFT | RIGHT | FULL ] JOIN <join_table> ON <join_condition>
```