# [](#create-aggregating-index)CREATE AGGREGATING INDEX

## [](#overview)Overview

Aggregating indexes in Firebolt help speed up query performance by precomputing and storing frequent aggregations, reducing the need for expensive runtime calculations. Use aggregating indexes for queries that involve [aggregate functions](/sql_reference/functions-reference/aggregation/) including `GROUP BY`, `SUM`, and `AVG`. Instead of scanning a full table, queries use aggregating indexes to retrieve pre-aggregated results, leading to faster response times. These indexes update automatically when the underlying table changes, ensuring that queries using them run efficiently without manual intervention, and are most beneficial when querying large tables with precomputable aggregations.

## [](#syntax)Syntax

```
CREATE AGGREGATING INDEX <index_name> ON <table_name> (
  <key_expression>[,...<key_expressionN>],
  <aggregation>[,...<aggregationN>]
);
```

### [](#parameters)Parameters

Parameter Description `<index_name>` A unique name for the aggregating index. `<table_name>` The name of the table on which the index is created. `<key_expressions>` One or more expressions that will be used for filtering or as aggregation dimensions in queries. `<aggregations>` [Aggregation functions](/sql_reference/functions-reference/aggregation/) including `SUM`, `COUNT`, `AVG`, applied to table column expressions.

Firebolt automatically maintains the aggregating index when the table is updated, ensuring it stays optimized for queries.

## [](#examples)Examples

The following examples show use cases for the following:

- [Create a simple aggregating index](#create-a-simple-aggregating-index)
- [Create a complex aggregating index](#create-a-complex-aggregating-index)
- [Compute AVG with multiple filters](#compute-avg-with-multiple-filters)
- [Filter by a calculated column](#filter-by-a-calculated-column)

### [](#create-a-simple-aggregating-index)Create a simple aggregating index

You can optimize a frequently run query that counts distinct values, sums the `amount` for products, and groups by `product_name`, as shown in the following code example:

```
SELECT
  product_name,
  COUNT(DISTINCT source),
  SUM(amount)
FROM
  my_table
WHERE
  origin_country = 'USA'
GROUP BY
  product_name;
```

To optimize the previous query, create an aggregating index on `my_table` to precompute aggregations for `COUNT(DISTINCT source)` and `SUM(AMOUNT)`, grouped by `origin_country` and `product_name` as follows:

```
CREATE AGGREGATING INDEX my_table_agg_idx1 ON my_table (
  origin_country,
  product_name,
  COUNT(DISTINCT source),
  SUM(amount)
);
```

The aggregating index in the previous code example precomputes and stores aggregated values for `COUNT(DISTINCT source)` and `SUM(amount)`, reducing computation at query time.

In order to use the aggregating index, **a query must group by all indexed columns**. In the previous code example, you must include `origin_country` as a filter expression because aggregating indexes store grouped data hierarchically based on the first column listed. If you don’t include the first grouping column, the entire table is scanned, reducing performance benefits. Filtering on `product_name` is optional because all `product_name` values are already grouped under each `origin_country`. If filtering on `origin_country` is not always required, you can define an index on another column. For example, the following code example creates an aggregating index on `product_name`:

```
CREATE AGGREGATING INDEX my_table_agg_idx2 ON my_table (
  product_name,
  COUNT(DISTINCT source),
  SUM(amount)
);
```

Then, you can use the aggregating index in a query as follows:

```
SELECT product_name, COUNT(DISTINCT source), SUM(amount)
FROM my_table
GROUP BY product_name;
```

### [](#create-a-complex-aggregating-index)Create a complex aggregating index

The following code example groups products by name and replaces names longer than 100 characters with ‘Other.’ Then, it counts distinct brands approximately and filters records where `origin_country` is ‘USA’ and `amount` exceeds `100`:

```
SELECT
  CASE WHEN LENGTH(product_name) < 100 THEN product_name ELSE 'Other' END AS product_name,
  APPROX_COUNT_DISTINCT(brand)
FROM
  my_table
WHERE
  origin_country = 'USA' AND amount > 100
GROUP BY
  product_name;
```

To optimize the previous query, create an aggregating index on `my_table` that precomputes `APPROX_COUNT_DISTINCT(brand)`, grouped by `origin_country`, `amount`, and `product_name` as follows:

```
CREATE AGGREGATING INDEX my_table_agg_idx2 ON my_table (
  origin_country,
  amount,
  product_name,
  APPROX_COUNT_DISTINCT(brand)
);
```

In order to use the aggregating index, **a query must group by all indexed columns in the index**. In the previous code example, you must include `origin_country` and `amount` as filter expressions because aggregating indexes store grouped data hierarchically based on the first columns listed. If you don’t include the first grouping column, the entire table is scanned, reducing performance benefits. The second grouping column `amount` is necessary to efficiently filter down results. Without both groupings, Firebolt cannot sufficiently narrow the dataset efficiently.

## [](#compute-avg-with-multiple-filters)Compute `AVG` with multiple filters

The following aggregating index precomputes `AVG(price)` and `COUNT(product_id)`, grouped by `category`, `region` and `price`, allowing efficient filtering and retrieval of aggregates:

```
CREATE AGGREGATING INDEX avg_sales_per_category ON sales_data (
  category,
  region,
  price,
  AVG(price),
  COUNT(product_id)
);
```

The following query retrieves the precomputed `AVG(price)` and `COUNT(product_id)` while filtering on `region` and `price`:

```
SELECT category, AVG(price), COUNT(product_id)
FROM sales_data
WHERE region = 'North America' AND price > 100
GROUP BY category, region, price;
```

The previous code example utilizes the `avg_sales_per_category` aggregating index without requiring an explicit reference to it because Firebolt’s query optimizer automatically applies an aggregating index when a query’s `GROUP BY` structure matches the index’s grouping structure. If the query optimizer finds a matching structure, it retrieves the precomputed values instead of scanning and computing the raw data.

### [](#filter-by-a-calculated-column)Filter by a calculated column

The following aggregating index precomputes `SUM(order_value)` and `COUNT(order_id)`, grouped by `customer_segment` and `ROUND(order_value, 2)`, allowing efficient filtering on the rounded `order_value`:

```
CREATE AGGREGATING INDEX revenue_index ON orders (
  customer_segment,
  ROUND(order_value, 2),
  SUM(order_value),
  COUNT(order_id)
);
```

The following code example retrieves precomputed aggregates while filtering on the rounded `order_value`:

```
SELECT customer_segment, SUM(order_value), COUNT(order_id)
FROM orders
WHERE ROUND(order_value, 2) > 100
GROUP BY customer_segment, ROUND(order_value, 2);
```

The previous code example fully utilizes the aggregating index because it groups by all indexed columns: `customer_segment`, `ROUND(order_value, 2)`, allowing the use of the precomputed index instead of computing rounding at query time.

## [](#how-to-check-if-a-query-is-applying-an-aggregating-index)How to check if a query is applying an aggregating index

You can use [EXPLAIN](/sql_reference/commands/queries/explain.html) to check if Firebolt is applying an aggregating index as shown in the following code example:

```
EXPLAIN
SELECT product_name, COUNT(DISTINCT source), SUM(amount)
FROM my_table
WHERE origin_country = 'USA'
GROUP BY product_name;
```

If the aggregating index is being used, the query execution plan will show output as follows: `Aggregating Index: my_table_agg_idx1`

If the aggregating index is not used, ensure that the query filters on the first grouping column.

## [](#best-practices)Best practices

- **Always include query filter expressions** – Queries should filter on indexed columns to ensure the aggregating indexes are applied efficiently.
- **Prioritize low-cardinality columns** – Place columns with fewer unique values first to improve pruning efficiency.
- **Match aggregations to queries** – Ensure the index includes all measures used in `GROUP BY` queries for optimal performance.

By following these guidelines, you can use aggregating indexes to significantly speed up queries.

## [](#limitations)Limitations

- Because the index structure is precomputed, aggregating indexes are not suitable for queries that frequently change `GROUP BY` structures.
- Aggregating indexes do not optimize queries that filter on columns that are not part of the grouping structure unless additional indexing strategies are applied.
- Inserts and updates may be slower because of index maintenance, so avoid over-indexing to save storage space. Additionally, large updates can temporarily impact query performance.
- Aggregating indexes are not beneficial for transactional workloads that require frequent row-level updates.