# [](#primary-index)Primary index

The Firebolt primary index optimizes data retrieval by organizing it based on column values. This enables efficient data pruning and high-performance queries for large-scale analytics. The primary index ensures that queries target only the most relevant portions of the data, significantly reducing the volume of unnecessary scans. This selectivity is especially powerful when the indexed columns align closely with query patterns, allowing the database to quickly locate and retrieve the required data. As a result, query performance is not only optimized but also remains consistent even as data volumes grow.

Topics:

- [Key features](#key-features)
- [Syntax](#syntax)
- [Parameters](#parameters)
- [Example](#example)
- [Considerations](#considerations)
- [Advanced option: index granularity](#advanced-option-index-granularity)

## [](#key-features)Key features

- **Customizable indexing**:
  
  - Unlike traditional databases, Firebolt allows primary indexes on any column, so that you have the flexibility to align with query patterns.
  - Primary key constraints are not enforced, allowing for greater customization.
- **Tablet-based data organization**:
  
  - Data is divided into tablets, which are self-contained, fixed-size data chunks approximately 3GB each, sorted by primary index keys for efficient querying.
  - Within each tablet, your data is further segmented into intra-tablet ranges, approximately 8,000 rows each, for faster access.
- **Support for compound indexes**:
  
  - Multiple columns can be included in the primary index, optimizing complex query patterns by allowing efficient pruning and retrieval for queries that filter or group data based on combinations of these columns, such as filtering by both a date range and a category or joining fact and dimension tables on multiple keys.
- **Sparse indexing**:
  
  - Firebolt employs a sparse indexing approach, storing only the first row of each tablet range. This approach allows target reads and parallel processing, significantly reducing index size while increasing query performance.
- **Automatic metadata updates**:
  
  - The primary index is automatically maintained during data inserts, deletions, and updates ensuring optimal performance.
- **Handling low-cardinality clauses**:
  
  - Starting the primary index with low-cardinality columns can enhance pruning efficiency by creating long ordered runs of data. Low-cardinality columns are those with a limited number of unique values such as months, regions, or statuses, which group similar rows together. This reduces the number of tablet ranges Firebolt needs to scan, enabling faster query execution and better data pruning.
- **Inclusion of join key columns**:
  
  - In star schema designs, including join key columns, or foreign keys, in the primary index of fact tables can accelerate queries by facilitating efficient data retrieval.
- **Leverage indexed columns directly**:
  
  - Design queries to use indexed columns directly in WHERE clauses without transformations, ensuring that the primary index can be utilized to its full potential for faster data pruning and query execution.

## [](#syntax)Syntax

To define a primary index, use the following syntax within a `CREATE TABLE` statement:

```
CREATE TABLE <table_name> (
   <column1> <data_type>
   [, <column2> <data_type>,
   ...]
)
PRIMARY INDEX <column_name1>[, <column_name2>, ...]
[WITH ( index_granularity = <index_granularity> ) ];
```

## [](#parameters)Parameters

Parameter Description `table_name` The name of the table where the primary index is applied. `column_name1, ...` The columns chosen to be included in the primary index. `index_granularity` The maximum number of rows in each granule. See [Index granularity](#advanced-option-index-granularity).

## [](#example)Example

The following example creates a table with a primary index optimized for query performance by filtering for `SubmitDate` and `EngineName`:

```
CREATE [FACT|DIMENSION] TABLE QueryHistory (
  QueryID TEXT,
  QueryText TEXT,
  SubmitDate DATE,
  EngineName TEXT,
  SubmitTime DATE,
  Latency INT
)
PRIMARY INDEX SubmitDate, EngineName;
```

## [](#considerations)Considerations

- **Non-enforced primary key constraint**:  
  Firebolt does not enforce primary key constraints, so users must manage data integrity externally.
- **Managing fragmentation**:  
  Fragmentation can occur as you insert, delete, or update data in a table, which impacts storage efficiency and potentially affects your query performance. Firebolt provides tools to help mitigate this effect:
  
  - **Efficient deletion management**:  
    Instead of immediately removing rows from the table, Firebolt uses a deletion mask vector to flag rows as deleted. This vector marks rows for exclusion during queries while keeping the underlying data intact until cleanup is performed.  
    This approach ensures consistency and avoids disrupting the primary index during updates or deletions.
  - **fragmentation metric**:  
    Use the `information_schema.tables` to access the fragmentation metric to assess fragmentation levels and determine whether maintenance actions are needed.
  - **[VACUUM](/sql_reference/commands/data-management/vacuum.html) command**:  
    You can use the `VACUUM` command to clean up rows flagged for deletion and reorganize fragmented data. It is particularly useful when large numbers of rows have been deleted or updates have introduced significant fragmentation.
- **Query Performance Overhead**:  
  While sparse indexes enable targeted reads and parallel processing to improve query performance, they may still require scanning one tablet range from multiple tablets, even for highly selective filters. This can result in more data being scanned compared to a globally sorted index, potentially affecting performance in certain scenarios.
- **Column Selection**:  
  Choose columns with high selectivity and relevance to query patterns for optimal performance. **Selectivity** refers to the ability of a column to significantly narrow down the dataset when filtered, typically measured by the proportion of unique values in the column. Columns with higher selectivity, such as IDs or timestamps, help reduce the number of rows scanned, leading to faster query execution and better resource efficiency.

Using Firebolt’s primary indexes can help you enhance your query performance, optimize data management, and scale efficiently for modern analytics workloads.

## [](#advanced-option-index-granularity)Advanced option: index granularity

The `index_granularity` [storage parameter](/sql_reference/commands/data-definition/create-fact-dimension-table.html#storage-parameters), specified in the `WITH` clause, is an advanced setting that may be useful for improving performance in very specific query patterns. It defines the maximum number of rows per granule, which directly impacts how data is indexed and queried.

### [](#how-index-granularity-works)How index granularity works

A granule is the smallest block of rows that Firebolt can skip or read during query filtering. Index granularity defines the number of rows in each granule. In other words, it sets the smallest group of rows the engine can access independently.

- **Lower index granularity** creates smaller granules, allowing more precise filtering and reducing unnecessary row scans in selective queries. However, lower index granularity also increases memory usage and overhead from managing more granules.
- **Higher index granularity values** creates larger granules, lowering memory usage and management overhead but increasing the chance of scanning irrelevant rows, especially in selective queries.

For more information about the fundamentals of Firebolt’s primary indexes and granules, see Firebolt’s blog post on [primary indexes](https://www.firebolt.io/blog/primary-indexes-in-firebolt-a-comprehensive-guide-to-understanding-managing-and-selecting).

### [](#accepted-values)Accepted values

`<index_granularity>` must be a power of 2, ranging from 128 to 8192. The default value is 8192. We recommend using the default value, but lower values can decrease query latency by 10x or more in some query patterns.

### [](#best-practices)Best practices

Use the default value of `index_granularity`, which should translate to good performance for most queries. The following workload patterns may benefit from higher or lower values for `index_granularity`:

- If your queries access only a few rows per granule, such as single-row queries or individual rows spread throughout a table, setting a **lower** `index_granularity` value can reduce unnecessary row scans and improve efficiency. However, this increases static memory usage for storing the index.
- If most of your queries scan large portions of the table, such as a large bounded range of primary index columns, a **higher** `index_granularity` value is more efficient, as it reduces index memory usage and overhead introduced by each granule boundary.

If you want to adjust `index_granularity`, start with the default value, then create duplicate tables with different settings to compare both the query latency and memory usage.