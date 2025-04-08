# [](#explain)EXPLAIN

The `EXPLAIN` feature analyzes and displays how Firebolt’s query processing system runs your query. You can use this information to understand resource utilization, and identify opportunities to improve and optimize your query performance.

If you specify the `ANALYZE` option for `EXPLAIN`, Firebolt also collects detailed metrics about each operator during query runtime, including the amount of time spent on the operator, and how much data it processes.

## [](#syntax)Syntax

```
EXPLAIN [( <option_name> [<option_value>] [, ...] )] <select_statement>
```

Parameter Description `option_name` The name of an option. See below for a list of all available options. `option_value` The value of the option. If no value is specified, the default is `TRUE`. `<select_statement>` Any select statement.

## [](#explain-options)Explain Options

You can augment the output of `EXPLAIN` by specifying options. The following table lists all available options and their functionalities:

Option Name Option Values Description `LOGICAL` `TRUE`, `FALSE` Returns the optimized logical query plan by default, unless otherwise specified. `PHYSICAL` `TRUE`, `FALSE` Returns an optimized physical query plan containing shuffle operators for queries on distributed engines, showing how work is distributed between nodes of an engine. `ANALYZE` `TRUE`, `FALSE` Returns an optimized physical query plan annotated with metrics from query execution. For more information about these metrics, see [Example with ANALYZE](#example-with-analyze). `ALL` `TRUE`, `FALSE` Returns all of the previous `LOGICAL`, `PHYSICAL`, and `ANALYZE` plans. Use the following sample syntax: `EXPLAIN (ALL) <select statement>`. `STATISTICS` `TRUE`, `FALSE` Returns an annotated query plan that includes estimates from Firebolt’s query optimizer. This option works with `LOGICAL`, `PHYSICAL`, `ANALYZE`, and `ALL`. Use the following sample syntax: `EXPLAIN (STATISTICS) <select statement>`.

You may specify only one of the following options: `LOGICAL`, `PHYSICAL`, `ANALYZE`, and `ALL`. If you need to view several plans at once, use the `ALL` option.

## [](#example)Example

The following example shows how to generate an `EXPLAIN` statement for a `SELECT` query on a table named `lineitem`.

```
EXPLAIN
SELECT
	l_shipdate,
	l_linestatus,
	l_orderkey,
	AVG(l_discount)
FROM
	lineitem
WHERE
	l_returnflag = 'N'
	AND l_shipdate > '1996-01-01'
GROUP BY
	ALL
ORDER BY
	1,2,3,4;
```

**Returns:**

```
[0] [Projection] lineitem.l_shipdate, lineitem.l_linestatus, lineitem.l_orderkey, avgOrNull(lineitem.l_discount)
 \_[1] [Sort] OrderBy: [lineitem.l_shipdate Ascending Last, lineitem.l_linestatus Ascending Last, lineitem.l_orderkey Ascending Last, avgOrNull(lineitem.l_discount) Ascending Last]
    \_[2] [Aggregate] GroupBy: [lineitem.l_orderkey, lineitem.l_linestatus, lineitem.l_shipdate] Aggregates: [avgOrNull(lineitem.l_discount)]
       \_[3] [Projection] lineitem.l_orderkey, lineitem.l_discount, lineitem.l_linestatus, lineitem.l_shipdate
          \_[4] [Predicate] equals(lineitem.l_returnflag, 'N'), greater(lineitem.l_shipdate, toPGDate('1996-01-01'))
             \_[5] [StoredTable] Name: 'lineitem', used 5/16 column(s) FACT
```

The `EXPLAIN` output shows the sequence of operations that Firebolt’s engine will use to run the query:

1. Read the required columns from the `lineitem` table.
2. Filter the `lineitem` table by the `WHERE` conditions.
3. Remove the columns that are no longer required.
4. Perform the aggregation as specified by the `GROUP BY` clause.
5. Sort resulting rows in ascending order by all columns from the `ORDER BY` clause.
6. Bring the columns into the order specified in the `SELECT` clause.

## [](#example-with-analyze)Example with ANALYZE

The following code example runs the same query as the previous example on a multi-node engine using the `(ALL)` option, which is the same as specifying `(ALL TRUE)` or `(LOGICAL, PHYSICAL, ANALYZE)`.

```
EXPLAIN (ALL)
SELECT
	l_shipdate,
	l_linestatus,
	l_orderkey,
	AVG(l_discount)
FROM
	lineitem
WHERE
	l_returnflag = 'N'
	AND l_shipdate > '1996-01-01'
GROUP BY
	ALL
ORDER BY
	1,2,3,4;
```

**Returns:**

Disclaimer

The format of this result is subject to change.

The query in the previous example returns three columns `explain_logical`, `explain_physical`, and `explain_analyze`, as shown in the following sections.

Each column can be toggled on or off using the corresponding options. For example, to display only `explain_logical` and `explain_analyze`, you can specify either `(LOGICAL, ANALYZE)` or `(ALL, PHYSICAL FALSE)`. With the `ANALYZE` option enabled, this query took 4.7 seconds to run. Without the `ANALYZE` option, the query does not run and should return the result almost immediately.

### [](#explain-logical-output)`EXPLAIN (LOGICAL)` output

```
[0] [Projection] ref_2, ref_1, ref_0, ref_3
|   [RowType]: date not null, text not null, bigint not null, double precision null
 \_[1] [Sort] OrderBy: [ref_2 Ascending Last, ref_1 Ascending Last, ref_0 Ascending Last, ref_3 Ascending Last]
   |   [RowType]: bigint not null, text not null, date not null, double precision null
    \_[2] [Aggregate] GroupBy: [ref_0, ref_2, ref_3] Aggregates: [avg2(ref_1)]
      |   [RowType]: bigint not null, text not null, date not null, double precision null
       \_[3] [Projection] ref_0, ref_1, ref_3, ref_4
         |   [RowType]: bigint not null, double precision not null, text not null, date not null
          \_[4] [Filter] (ref_2 = 'N'), (ref_4 > DATE '1996-01-01')
            |   [RowType]: bigint not null, double precision not null, text not null, text not null, date not null
             \_[5] [StoredTable] Name: "lineitem", used 5/16 column(s) FACT, ref_0: "l_orderkey" ref_1: "l_discount" ref_2: "l_returnflag" ref_3: "l_linestatus" ref_4: "l_shipdate"
                   [RowType]: bigint not null, double precision not null, text not null, text not null, date not null
```

The `explain_logical` column shows the optimized logical query plan of the `SELECT` statement, identical to the plan in [Example](#example). Additionally, it shows the output row type for each operator. The variable `ref_2` refers to the second input column from the operator below. For operators with multiple input operators such as the `Join` operator, input columns are concatenated. For example, if the first input produces three columns, and the second input produces four columns, the first input’s columns are labeled `ref_0` to `ref_2`, while the second input’s columns are labeled `ref_3` to `ref-6`.

### [](#explain-physical-output)`EXPLAIN (PHYSICAL)` output

```
[0] [MaybeCache]
|   [RowType]: date not null, text not null, bigint not null, double precision null
 \_[1] [Projection] ref_2, ref_1, ref_0, ref_3
   |   [RowType]: date not null, text not null, bigint not null, double precision null
    \_[2] [SortMerge] OrderBy: [ref_2 Ascending Last, ref_1 Ascending Last, ref_0 Ascending Last, ref_3 Ascending Last]
      |   [RowType]: bigint not null, text not null, date not null, double precision null
       \_[3] [Shuffle] Gather
         |   [RowType]: bigint not null, text not null, date not null, double precision null
         |   [Affinity]: many nodes
          \_[4] [Sort] OrderBy: [ref_2 Ascending Last, ref_1 Ascending Last, ref_0 Ascending Last, ref_3 Ascending Last]
            |   [RowType]: bigint not null, text not null, date not null, double precision null
             \_[5] [AggregateMerge] GroupBy: [ref_0, ref_1, ref_2] Aggregates: [avg2merge(ref_3)]
               |   [RowType]: bigint not null, text not null, date not null, double precision null
                \_[6] [Shuffle] Hash by [ref_0, ref_1, ref_2]
                  |   [RowType]: bigint not null, text not null, date not null, aggregatefunction(avg2ornull, double precision not null) not null
                  |   [Affinity]: many nodes
                   \_[7] [AggregateState partial] GroupBy: [ref_0, ref_2, ref_3] Aggregates: [avg2(ref_1)]
                     |   [RowType]: bigint not null, text not null, date not null, aggregatefunction(avg2ornull, double precision not null) not null
                      \_[8] [Projection] ref_0, ref_1, ref_3, ref_4
                        |   [RowType]: bigint not null, double precision not null, text not null, date not null
                         \_[9] [Filter] (ref_2 = 'N'), (ref_4 > DATE '1996-01-01')
                           |   [RowType]: bigint not null, double precision not null, text not null, text not null, date not null
                            \_[10] [StoredTable] Name: "lineitem", used 5/16 column(s) FACT, ref_0: "l_orderkey" ref_1: "l_discount" ref_2: "l_returnflag" ref_3: "l_linestatus" ref_4: "l_shipdate"
                                  [RowType]: bigint not null, double precision not null, text not null, text not null, date not null
```

The `explain_physical` column displays a detailed, optimized physical plan that includes Shuffle operators for distributed query execution, offering insights into how tasks are distributed across the engine’s nodes. A `Shuffle` operator redistributes data across engine nodes. For example, scans of `FACT` tables, like operator `[9] [StoredTable]` in the previous example, and the operators following it are automatically distributed across all nodes of an engine. A `Shuffle` operator of type `Hash` indicates that the workload of the following operators is distributed across all nodes, while a `Shuffle` operator of type `Gather` consolidates data onto a single node of the engine. In the previous example, only the operator `[1] [SortMerge]` runs on a single node, merging the sorted partial query results from all other nodes.

The `MaybeCache` operator, at the top of the plan, caches its input in main memory on a best-effort basis for future runs of the same (sub-) query. It considers the entire plan leading to its input, as well as the state of the scanned tables. If the data in a table changes, the `MaybeCache` operator will know not to read an outdated cached entry. It may also skip caching large results or prioritize caching results that offer greater time savings. The `MaybeCache` operator may appear in different places of a plan, but in this query it caches the full query result. For more information about the `MaybeCache` operator, see [Subresult Reuse in Firebolt](/Overview/queries/understand-query-performance-subresult.html).

### [](#explain-analyze-output)`EXPLAIN (ANALYZE)` output

```
[0] [MaybeCache]
|   [RowType]: date not null, text not null, bigint not null, double precision null
|   [Execution Metrics]: output cardinality = 270958684, thread time = 7ms, cpu time = 4ms
 \_[1] [Projection] ref_2, ref_1, ref_0, ref_3
   |   [RowType]: date not null, text not null, bigint not null, double precision null
   |   [Execution Metrics]: Optimized out
    \_[2] [SortMerge] OrderBy: [ref_2 Ascending Last, ref_1 Ascending Last, ref_0 Ascending Last, ref_3 Ascending Last]
      |   [RowType]: bigint not null, text not null, date not null, double precision null
      |   [Execution Metrics]: output cardinality = 270958684, thread time = 27932ms, cpu time = 27692ms
       \_[3] [Shuffle] Gather
         |   [RowType]: bigint not null, text not null, date not null, double precision null
         |   [Affinity]: many nodes
         |   [Execution Metrics]: output cardinality = 270958684, thread time = 1846ms, cpu time = 1779ms
          \_[4] [Sort] OrderBy: [ref_2 Ascending Last, ref_1 Ascending Last, ref_0 Ascending Last, ref_3 Ascending Last]
            |   [RowType]: bigint not null, text not null, date not null, double precision null
            |   [Execution Metrics]: output cardinality = 270958684, thread time = 286199ms, cpu time = 285590ms
             \_[5] [AggregateMerge] GroupBy: [ref_0, ref_1, ref_2] Aggregates: [avg2merge(ref_3)]
               |   [RowType]: bigint not null, text not null, date not null, double precision null
               |   [Execution Metrics]: output cardinality = 270958684, thread time = 90694ms, cpu time = 77308ms
                \_[6] [Shuffle] Hash by [ref_0, ref_1, ref_2]
                  |   [RowType]: bigint not null, text not null, date not null, aggregatefunction(avg2ornull, double precision not null) not null
                  |   [Affinity]: many nodes
                  |   [Execution Metrics]: output cardinality = 270958691, thread time = 27649ms, cpu time = 24247ms
                   \_[7] [AggregateState partial] GroupBy: [ref_0, ref_2, ref_3] Aggregates: [avg2(ref_1)]
                     |   [RowType]: bigint not null, text not null, date not null, aggregatefunction(avg2ornull, double precision not null) not null
                     |   [Execution Metrics]: output cardinality = 270958691, thread time = 191561ms, cpu time = 184761ms
                      \_[8] [Projection] ref_0, ref_1, ref_3, ref_4
                        |   [RowType]: bigint not null, double precision not null, text not null, date not null
                        |   [Execution Metrics]: Optimized out
                         \_[9] [Filter] (ref_2 = 'N'), (ref_4 > DATE '1996-01-01')
                           |   [RowType]: bigint not null, double precision not null, text not null, text not null, date not null
                           |   [Execution Metrics]: output cardinality = 275468571, thread time = 3578ms, cpu time = 3566ms
                            \_[10] [StoredTable] Name: "lineitem", used 5/16 column(s) FACT, ref_0: "l_orderkey" ref_1: "l_discount" ref_2: "l_returnflag" ref_3: "l_linestatus" ref_4: "l_shipdate"
                                  [RowType]: bigint not null, double precision not null, text not null, text not null, date not null
                                  [Execution Metrics]: output cardinality = 334006486, thread time = 29642ms, cpu time = 29608ms
```

The `explain_analyze` column contains the same query plan as the `explain_physical` column, but annotated with metrics collected during query execution. For each operator, it shows the number of rows it produced in `output_cardinality`, and how much time was spent on that operator in `thread_time` and `cpu_time`. The `thread_time` is the sum of the wall-clock time that threads spent working on the operator across all nodes, while `cpu_time` is the total time those threads were scheduled on a CPU core. A significantly smaller `cpu_time` compared to `thread_time` suggests that the operator is either waiting on input or output operations or that the engine is handling multiple queries simultaneously.

#### [](#analyzing-the-metrics)Analyzing the Metrics

In the previous example, the `cpu_time` is almost as high as the `thread_time` on the `StoredTable` node. This indicates that the data of the `lineitem` table was in the SSD cache. On a cold query run, where data must be retrieved from an Amazon S3 bucket, `thread_time` is considerably higher than `cpu_time` for the same operator:

```
[10] [StoredTable] Name: "lineitem", used 5/16 column(s) FACT, ref_0: "l_orderkey" ref_1: "l_discount" ref_2: "l_returnflag" ref_3: "l_linestatus" ref_4: "l_shipdate"
    [RowType]: bigint not null, double precision not null, text not null, text not null, date not null
    [Execution Metrics]: output cardinality = 334006486, thread time = 29642ms, cpu time = 29608ms
```

Additionally, a significant amount of time is devoted to sorting the query results:

```
[2] [SortMerge] OrderBy: [ref_2 Ascending Last, ref_1 Ascending Last, ref_0 Ascending Last, ref_3 Ascending Last]
|   [RowType]: bigint not null, text not null, date not null, double precision null
|   [Execution Metrics]: output cardinality = 270958684, thread time = 27932ms, cpu time = 27692ms
 \_[3] [Shuffle] Gather
   |   [RowType]: bigint not null, text not null, date not null, double precision null
   |   [Affinity]: many nodes
   |   [Execution Metrics]: output cardinality = 270958684, thread time = 1846ms, cpu time = 1779ms
    \_[4] [Sort] OrderBy: [ref_2 Ascending Last, ref_1 Ascending Last, ref_0 Ascending Last, ref_3 Ascending Last]
      |   [RowType]: bigint not null, text not null, date not null, double precision null
      |   [Execution Metrics]: output cardinality = 270958684, thread time = 286199ms, cpu time = 285590ms
```

Removing the `ORDER BY` can considerably speed up query runtimes. In this case, the total query time is nearly halved, dropping from 1 minute and 19 seconds to 47.3 seconds. If only the initial results in the specified sorting order are needed, introducing a `LIMIT` operator can further enhance performance. Adding `LIMIT 10000` to the query produces the following output:

```
[2] [SortMerge] OrderBy: [ref_2 Ascending Last, ref_1 Ascending Last, ref_0 Ascending Last, ref_3 Ascending Last] Limit: [10000]
|   [RowType]: bigint not null, text not null, date not null, double precision null
|   [Execution Metrics]: output cardinality = 10000, thread time = 1ms, cpu time = 1ms
 \_[3] [Shuffle] Gather
   |   [RowType]: bigint not null, text not null, date not null, double precision null
   |   [Affinity]: many nodes
   |   [Execution Metrics]: output cardinality = 20000, thread time = 0ms, cpu time = 0ms
    \_[4] [Sort] OrderBy: [ref_2 Ascending Last, ref_1 Ascending Last, ref_0 Ascending Last, ref_3 Ascending Last] Limit: [10000]
      |   [RowType]: bigint not null, text not null, date not null, double precision null
      |   [Execution Metrics]: output cardinality = 20000, thread time = 6403ms, cpu time = 6372ms
       \_[5] [AggregateMerge] GroupBy: [ref_0, ref_1, ref_2] Aggregates: [avg2merge(ref_3)]
         |   [RowType]: bigint not null, text not null, date not null, double precision null
         |   [Execution Metrics]: output cardinality = 270958684, thread time = 83213ms, cpu time = 74525ms
```

The time spent in the `Sort` operator is significantly reduced, nearly by an order of magnitude. The `Sort` operator takes the `LIMIT` clause directly into account and emits only the required minimum number of rows. This reduces the overall query time from 1 minute and 19 seconds to 40.3 seconds, which is an even greater improvement than removing the `ORDER BY` clause, as less result data needs to be serialized for client return.

A secondary benefit is that the result is now small enough to be cached by the `MaybeCache` operator. Running the query with `LIMIT 10000` again produces the following output:

```
[0] [MaybeCache]
|   [RowType]: date not null, text not null, bigint not null, double precision null
|   [Execution Metrics]: output cardinality = 10000, thread time = 0ms, cpu time = 0ms
 \_[1] [Projection] ref_2, ref_1, ref_0, ref_3
   |   [RowType]: date not null, text not null, bigint not null, double precision null
   |   [Execution Metrics]: Optimized out
    \_[2] [SortMerge] OrderBy: [ref_2 Ascending Last, ref_1 Ascending Last, ref_0 Ascending Last, ref_3 Ascending Last] Limit: [10000]
      |   [RowType]: bigint not null, text not null, date not null, double precision null
      |   [Execution Metrics]: Nothing was executed
```

## [](#example-with-statistics)Example with STATISTICS

The following code example uses the `STATISTICS` option to display the logical and physical query plan for the `SELECT` statement, annotated with estimated statistics from Firebolt’s query optimizer:

```
EXPLAIN (STATISTICS)
SELECT
	l_shipdate,
	l_linestatus,
	l_orderkey,
	AVG(l_discount)
FROM
	lineitem
WHERE
	l_returnflag = 'N'
	AND l_shipdate > '1996-01-01'
GROUP BY
	ALL
ORDER BY
	1,2,3,4;
```

In the following output, each plan node includes a `[Logical Profile]`, which specifies the `source` of the estimation for that node, along with estimated row counts and, if available, distinct column counts:

```
explain TEXT
[0] [Projection] ref_2, ref_1, ref_0, ref_3
|   [RowType]: date null, text null, bigint null, numeric(38, 9) null
|   [Logical Profile]: [est. #rows=6676, column profiles={}, source: estimated]
 \_[1] [Sort] OrderBy: [ref_2 Ascending Last, ref_1 Ascending Last, ref_0 Ascending Last, ref_3 Ascending Last]
   |   [RowType]: bigint null, text null, date null, numeric(38, 9) null
   |   [Logical Profile]: [est. #rows=6676, column profiles={}, source: estimated]
    \_[2] [Aggregate] GroupBy: [ref_0, ref_2, ref_3] Aggregates: [avg2(ref_1)]
      |   [RowType]: bigint null, text null, date null, numeric(38, 9) null
      |   [Logical Profile]: [est. #rows=6676, column profiles={[ref_3: #distinct=6676]}, source: estimated]
       \_[3] [Projection] ref_0, ref_1, ref_3, ref_4
         |   [RowType]: bigint null, numeric(38, 9) null, text null, date null
         |   [Logical Profile]: [est. #rows=23489, column profiles={}, source: estimated]
          \_[4] [Filter] (ref_2 = 'N'), (ref_4 > DATE '1996-01-01')
            |   [RowType]: bigint null, numeric(38, 9) null, text null, text null, date null
            |   [Logical Profile]: [est. #rows=23489, column profiles={[ref_2: #distinct=1]}, source: estimated]
             \_[5] [StoredTable] Name: "lineitem", used 5/16 column(s) FACT, ref_0: "l_orderkey" ref_1: "l_discount" ref_2: "l_returnflag" ref_3: "l_linestatus" ref_4: "l_shipdate"
                   [RowType]: bigint null, numeric(38, 9) null, text null, text null, date null
                   [Logical Profile]: [est. #rows=6.59431e+08, column profiles={}, source: metadata]
```

The possible sources are as follows:

**Source** **Description** `hardcoded` This node received hard-coded defaults because no additional information is available, a scenario that typically occurs only with external tables. `estimated` This node’s profile was computed from its child nodes’ profiles, based on assumptions about the data that may not always hold true. These assumptions can introduce inaccuracies, which may be further amplified. `metadata` This node’s profile was constructed from available metadata, ensuring that the information is either precise or highly accurate. `history` This node’s profile was obtained from a previously recorded run, and is accurate unless data has changed since the recording. `learned` This node’s profile was predicted using machine learning.