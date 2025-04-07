# [](#data-modeling)Data modeling

Firebolt optimizes data storage and retrieval by using indexing, partitioning, and compute scaling to enhance query efficiency. Design your data model and use databases, tables, and indexes to optimize query performance.

- **Minimizing data scans**  
  Firebolt indexes retrieve only the specific data ranges needed to satisfy a query, reducing the amount of data scanned.
- **Automatic index maintenance**  
  During data loading, Firebolt automatically sorts, compresses, and incrementally updates indexes to keep them optimized.
- **Tablet-based optimization**  
  Data and indexes are committed as tablets, which Firebolt automatically merges and optimizes as the data evolves.
- **Best practices for performance**  
  This guide provides recommendations for organizing tables, indexes, and data to achieve fast query results and peak performance.

## [](#how-it-works)How it works

Firebolt’s indexing and partitioning strategies are designed to take advantage of a cloud-based architecture that scales to handle very large data sets. Data is queried using multiple nodes for parallel processing. Data is also stored by columns, which allows for:

- **Optimized read operations**
- **Less disk space for storage**
- **Vectorized processing**

Firebolt also separates compute resources from storage resources so you can scale either up or down depending on your use case. Optimize your resources based on your changing workloads, and pay only for what you use.

Firebolt’s data modeling strategies work best with **Firebolt’s managed tables**, which leverage Firebolt’s performance optimizations. External tables, which allow users to access data without loading it into Firebolt, provide flexibility for querying external data sources like Amazon S3, but they do not offer the same performance benefits as managed tables.

The following sections show you how to use the previous data modeling strategies to decrease the number of bytes scanned to improve query performance, reduce storage costs, and optimize compute resources.

Topics:

- [Data modeling](#data-modeling)
  
  - [How it works](#how-it-works)
    
    - [Databases](#databases)
      
      - [Create a database](#create-a-database)
      - [Manage a database](#manage-a-database)
      - [Database best practices](#database-best-practices)
        
        - [Database best practice example](#database-best-practice-example)
      - [Evaluate your database for performance](#evaluate-your-database-for-performance)
    - [Schema](#schema)
      
      - [Schema best practices](#schema-best-practices)
    - [Tables](#tables)
      
      - [Firebolt-managed tables](#firebolt-managed-tables)
      - [External tables](#external-tables)
      - [Editing and deleting tables](#editing-and-deleting-tables)
      - [Primary indexes in tables](#primary-indexes-in-tables)
        
        - [Create a primary index](#create-a-primary-index)
        - [Primary index best practices](#primary-index-best-practices)
      - [Aggregating indexes in tables](#aggregating-indexes-in-tables)
        
        - [Create an aggregating index](#create-an-aggregating-index)
        - [Best practices for aggregate indexes](#best-practices-for-aggregate-indexes)
      - [Partitions in tables](#partitions-in-tables)
        
        - [Create a partition](#create-a-partition)
      - [Suggested indexes and partitions](#suggested-indexes-and-partitions)
    - [Additional resources](#additional-resources)

* * *

### [](#databases)Databases

**Logical structure**  
In Firebolt, a database is a logical structure that organizes and stores data for efficient querying and management. Databases are created using the `CREATE DATABASE` statement and can be modified or deleted using `ALTER` and `DROP`.

**Compute and storage separation**  
Firebolt separates compute and storage layers, ensuring databases are efficiently managed without impacting processing power.

**Data Security**  
Databases integrate with Firebolt’s data security layers, enforcing access controls and user permissions to safeguard sensitive information.

**Performance and integration**

- **Fast analytics**: Firebolt databases enable quick querying of large datasets, supporting fast analytics.
- **Third-Party integration**: They are easily integrated with third-party tools through APIs and drivers for programming languages like Python and .NET.

Database topics:

- [Create a database](#create-a-database) – Use a system or user engine to create a database.
- [Manage a database](#manage-a-database) – How to edit and delete databases.
- [Database best practices](#database-best-practices) – How to organize your databases for the best performance.
- [Evaluate your database for performance](#evaluate-your-database-for-performance).

#### [](#create-a-database)Create a database

Use [CREATE DATABASE](/sql_reference/commands/data-definition/create-database.html), which requires only the name of the database, and an **engine** to create a database.

Firebolt provides two types of engines:

- **System engines** handle administrative tasks like creating and managing databases and tables without any need for configuration. They are always available but cannot process user data. **User engines** perform both administrative tasks and handle queries that access and process user data. User engines provide the flexibility needed for data processing and analysis.

The following code example creates a `test_library` database with an optional description:

```
CREATE DATABASE IF NOT EXISTS test_library 
    WITH DESCRIPTION = 'database about library';
```

#### [](#manage-a-database)Manage a database

After you create a database, you can create additional objects and run queries within it. You can also modify the database parameters.

The following code example modifies the description of a database:

```
ALTER DATABASE library SET DESCRIPTION = 'database about library';
```

You can delete a database with [DROP DATABASE](/sql_reference/commands/data-definition/drop-database.html), which permanently deletes the database and all its contents. Since this action cannot be undone, use it carefully. Create a backup before dropping the database to prevent data loss.

The following code example uses the optional `IF EXISTS` clause to see if the `test_library` database exists, and then deletes it:

```
DROP DATABASE IF EXISTS test_library;
```

Active queries will continue using the latest version of the database, but new queries submitted after the database is dropped will fail. To minimize disruptions, monitor active queries and notify users in advance.

#### [](#database-best-practices)Database best practices

Efficient database design in Firebolt is key to optimizing query performance, managing scalability, and ensuring workload isolation. By using databases to logically separate data, workloads from each other. You can reduce the impact of schema changes, minimize query latency, and enhance access control through role-based policies. For example, when managing data for applications that operate across multiple regions, understanding when to use a single database versus separate databases can significantly affect performance and operational efficiency. The following best practices outline how to structure databases in Firebolt to achieve these goals:

- **Use databases to logically separate workloads** – Organize data into separate databases to isolate workloads and reduce the impact of schema changes on unrelated queries.
- **Group related tables in the same database** – Only include related tables within the same database to minimize the overhead caused by metadata synchronization during schema changes.
- **Separate data for different regions or use cases** – For multi-region applications, create separate databases for each region to avoid conflicts such as overlapping primary keys, and to prevent schema changes in one region from affecting query performance in another.
- **Minimize the number of tables in a single database** – Avoid adding too many unique tables to a single database, as schema changes require the metadata service to synchronize all associated schemas, potentially increasing query latency.
- **Leverage logical separation for access control** – Use separate databases to enforce role-based access control policies, ensuring users can only access data relevant to their roles and restricting queries across databases.
- **Consider metadata sync impacts on query latency** – Metadata synchronization can introduce slight query latencies after schema changes, especially as the number of tables in a database grows.

##### [](#database-best-practice-example)Database best practice example

The maker of the fictional [UltraFast game](https://help.firebolt.io/t/ultra-fast-gaming-firebolt-sample-dataset/250) plans to expand into a second region, where the game will generate data with the same structure and format as the first region but potentially overlapping primary keys.

The developers have two options:

1. Store the new region’s data in the existing database.
2. Create a separate database specifically for the new region.

**Recommended solution** Create a separate database for the new region for the following reasons:

- **Optimal Query Performance** – Firebolt’s decoupled architecture requires metadata synchronization for schema changes. As the number of tables in a database grows, this synchronization can introduce slight query latency. Storing data for each region in separate databases eliminates this risk and ensures efficient query performance.
- **Improved data isolation** – Queries for one region are not impacted by schema changes in another. Separate databases allow role-based access control, restricting access to data by region, which enhances security and operational control.

By creating a dedicated database for each region, the developers ensure optimal performance, scalability, and data management tailored to the needs of the UltraFast game expansion.

#### [](#evaluate-your-database-for-performance)Evaluate your database for performance

You can use the code in the [Firebolt DDL Performance GitHub repository](https://github.com/wagjamin/firebolt-ddl-performance/tree/main) to measure the efficiency of your database design. Use the code to generate a set of pdf plots for latency to evaluate your choices including the following:

- How increasing the number of tables impacts schema synchronization times.
- Analyze whether separating unrelated tables into different databases improves performance.
- Validate the benefits of isolating data for different use cases or regions in separate databases.
- Understand the trade-offs of schema changes on query latency and how to mitigate them through better database organization.

* * *

### [](#schema)Schema

A schema serves as a blueprint for how data is organized within a database, acting as a logical framework that groups tables, views, and indexes. Understanding how to structure and manage schemas is crucial for maintaining an organized database environment, making it easier to optimize queries and manage access. While schemas don’t directly improve query performance, the way you organize your data within them can significantly impact how efficiently queries run, especially as your data grows.

#### [](#schema-best-practices)Schema best practices

To enhance query performance, simplify access control, and ensure scalability as your data grows, use the following best practices:

- **Logical Organization** – Grouping related tables, views, and indexes within a schema, to create a logical structure that makes it easier to navigate and optimize your queries. Well-organized data allows for clearer queries, reducing complexity and improving readability.
- **Simplified Access Control** – Using schemas to manage access to certain tables or views ensures that only relevant data is queried by specific users or roles. This prevents unnecessary data scans or joins, which can improve query execution time.
- **Scalability** – As your database grows, a well-structured schema becomes even more important. Organizing data in a way that scales efficiently prevents bottlenecks, ensuring that queries can continue to run quickly even with larger datasets.

* * *

### [](#tables)Tables

In Firebolt, tables are the key components for organizing and storing data. They consist of rows (records) and columns (attributes), making them an integral tool in data modeling. Firebolt supports two main types of tables:

Table topics:

- [Firebolt-managed tables](#firebolt-managed-tables) – Fast and dimension tables make the best use of Firebolt’s optimization strategies.
- [External tables](#external-tables) – Users can access and query data without loading it into Firebolt. External tables generally have poorer performance compared to Firebolt-managed tables.
- [Editing and deleting tables](#editing-and-deleting-tables) – You can edit and delete an existing table.
- [Primary indexes in tables](#primary-indexes-in-tables) – Select the most efficient primary index for your tables based on your query patterns and data characteristics.
- [Aggregating indexes in tables](#aggregating-indexes-in-tables) – Precompute and update aggregation results for fast calculations and to save compute resources.
- [Partitions in tables](#partitions-in-tables) – Organize data based on date ranges, regions, customer types or other categories.
- [Suggested indexes and partitions](#suggested-indexes-and-partitions) – Use Firebolt’s [RECOMMEND\_DDL](/sql_reference/commands/queries/recommend_ddl.html) tool to suggest indexes and partitions based on your query history.

#### [](#firebolt-managed-tables)Firebolt-managed tables

Tables that are fully controlled by Firebolt make the best use of Firebolt’s optimization strategies. There are two types of managed tables, which serve different but complementary purposes:

- **Fact tables** – store large volumes of quantitative data, such as sales metrics or user interactions, and are used for analysis and reporting.
- **Dimension tables** – which hold descriptive data for enriching analyses, such as customer details or product categories. These tables are often duplicated across all nodes, and can dramatically improve query performance by providing fast access to frequently referenced information.

By designing tables with both fact and dimension roles, users can optimize data management for reporting and analytics.

**Fact tables**

The following code example creates a **fact table**:

```
CREATE TABLE borrowedbooks (
	transaction_id INT,
	book_id INT,
	member_id INT,
	checkout_date DATE,
	return_date DATE,
	late_fees NUMERIC
);
```

The internal structure of a fact table facilitates fast access to data even for very large datasets.

**Dimension tables**

Dimension tables are often replicated across all nodes, which enhances query performance and ensures quick access to frequently referenced information.

The following code example creates a **dimension table**:

```
CREATE DIMENSION TABLE books (
	book_id INT,
	title TEXT,
	author TEXT,
	genre TEXT,
	publication_year INT
);
```

In the previous example, the `books` table stores details about the publication and is ideal for joining with fact tables like `borrowedbooks` to analyze borrowing patterns or book popularity.

#### [](#external-tables)External tables

External tables allow users to access and query data stored outside the database, such as in Amazon S3, without loading it into Firebolt. This capability is particularly useful when working with large datasets stored externally that don’t require frequent access, enabling cost-efficient querying and avoiding data duplication. External tables generally have poorer performance compared to Firebolt managed tables because the data is not stored within Firebolt’s optimized infrastructure.

#### [](#editing-and-deleting-tables)Editing and deleting tables

After you create a table, you can modify your table ownership, delete it including all of its dependencies. Use the following SQL commands to manage your table:

**Change table ownership**

To change the ownership of a table, use the ALTER TABLE statement with the OWNER TO clause.

The following command assigns a new owner to the `borrowedbooks` table:

```
ALTER TABLE borrowedbooks OWNER TO 'new_owner';
```

**Delete a managed table**

Dropping a Firebolt managed table permanently deletes it along with all its data.

The following code checks to see if the `borrowedbooks` table exists and then deletes it:

```
DROP TABLE IF EXISTS borrowedbooks;
```

**Delete a managed table that has dependencies**

If the table that you want to delete has dependent objects such as views or aggregating indexes, you can use the `CASCADE` option to remove all dependencies together.

The following code deletes the `borrowedbooks` table and all related objects:

```
DROP TABLE borrowedbooks CASCADE;
```

**Drop an external table**

If you drop an external table, you only remove its definition from Firebolt. The actual data remains in the external source. The following code drops the `external_books` external table:

```
DROP TABLE external_books;
```

* * *

#### [](#primary-indexes-in-tables)Primary indexes in tables

Primary indexes optimize query performance by logically organizing data for efficient access based on key columns, using sparse indexing to minimize unnecessary scans during queries. This allows Firebolt’s engine to prune unnecessary data during queries, minimizing the amount of data read from disk, and accelerating query execution. By selecting columns frequently used in filters, the primary index enables fast, efficient data retrieval, which is crucial for large-scale, data-intensive applications. Properly configuring primary indexes ensures that Firebolt can maintain high performance, even with complex queries and large datasets.

**How Firebolt uses a primary index to optimize performance**

When new data is inserted into a table, Firebolt organizes it into tablets, which are logical partitions of the table. Each tablet holds a portion of the data, sorted according to the primary index. During query processing, Firebolt uses these indexes to eliminate blocks of rows not matching query predicate and scan only the necessary data, minimizing input and output operations, and optimizing performance.

For updates, Firebolt follows a “delete-and-insert” approach: the original row is marked for deletion, and the updated row is inserted into a new tablet. Deleted rows are not removed immediately but are flagged and later cleaned up during maintenance tasks.

##### [](#create-a-primary-index)Create a primary index

A primary index can only be defined when creating a new table, so if you need to modify the index, you’ll have to create a new table. To define a primary index, use the `PRIMARY INDEX` clause to specify your columns, as shown in the following example:

```
CREATE TABLE borrowedbooks_pi(
  transaction_id INT,
  book_id INT,
  member_id INT,
  checkout_date DATE,
  return_date DATE,
  late_fees NUMERIC)
PRIMARY INDEX (book_id, checkout_date);
```

In the previous example, [CREATE TABLE](/sql_reference/commands/data-definition/create-fact-dimension-table.html) creates a fact table by default, and the primary index acts as a composite index of both the `book_id` and `checkout_date` columns.

You can also create a primary index on a dimension table as shown in the following code example:

```
CREATE dimension TABLE books_pi(
  book_id INT,
  title TEXT,
  author TEXT,
  genre TEXT,
  publication_year INT
)
PRIMARY INDEX (book_id, title);
```

The primary index should include columns that are often used in queries that filter or aggregate data, including those that use `WHERE` and `GROUP BY` clauses, so that Firebolt can use these columns for efficient retrieval. Firebolt physically clusters the data in the table by the columns of the primary index. The primary index significantly improves query performance because it reduces the time spent scanning irrelevant rows and allows the query engine to retrieve only the necessary data.

##### [](#primary-index-best-practices)Primary index best practices

To optimize query performance and maximize indexing efficiency, use the following best practices for designing composite primary indexes:

- **Column selection** – Carefully select which columns to include in a composite primary index based on query patterns, cardinality, and the specific use cases for your table. Adding unnecessary columns can negatively impact performance.
- **Column order** – The order of columns in a composite primary index is critical. Place the most selective, or highest cardinality columns first in a composite primary index. Firebolt will search based on the order of columns in the primary index. High-cardinality columns, which have many distinct values, improve index efficiency by reducing the number of rows that need to be scanned during queries.

#### [](#aggregating-indexes-in-tables)Aggregating indexes in tables

Aggregating indexes precompute and store aggregation results from aggregation functions like `SUM`, `COUNT`, and `AVG`, as well as more complex calculations. When the underlying data is changed, Firebolt recalculates aggregate indexes automatically, so that they are always updated when you query them. Firebolt handles the following scenarios:

- Inserts – When new rows are added to a table included in an aggregating index, Firebolt recalculates the value to include the new data.
- Updates – If existing rows are updated, Firebolt adjusts the aggregation values to reflect the changes.
- Deletions – When rows are deleted, the aggregating index is recalculated to remove the row from the precomputed aggregated values.

When you utilize an aggregating index through a query, Firebolt utilizes the pre-calculated values instead of computing them in real-time. This reduces the computational burden at runtime and significantly speeds up query response times, especially for large datasets with high concurrency demands. Aggregating indexes are especially useful for frequently run queries that involve repeated aggregation operations, such as calculating totals, averages, or counts.

Aggregating indexes do require additional storage, because the precomputed data needs to be maintained. In write-heavy environments, frequent updates, inserts, or deletes can lead to increased computational overhead as the indexes must be recalculated and kept up to date. This can result in higher compute costs. It’s important to consider the performance benefits of faster queries against the additional storage and processing costs, especially for frequently changing datasets.

##### [](#create-an-aggregating-index)Create an aggregating index

Aggregating indexes can be created at the time a new table is made or afterward. You can define it as needed based on query patterns and performance.

The following code example shows how to create an aggregating index to precompute the number of transactions per borrower and their average late fee on the existing `borrowedbooks` table created in the previous **Fact tables** section in [Firebolt managed-tables](#firebolt-managed-tables):

```
CREATE AGGREGATING INDEX agg_borrower_statistics
   ON borrowedbooks (
  	borrower_id,
  	COUNT(transaction_id),
  	AVG(late_fee)
   );
```

When Firebolt runs a query that accesses either `transaction_id` or `late_fee`, it retrieves the precomputed results from the aggregating index, rather than computing them.

The following code example shows how to create an aggregating index to precompute the total late fees accumulated for each book:

```
CREATE AGGREGATING INDEX agg_total_late_fees
   ON borrowedbooks (
                   	book_id,
                   	SUM(late_fee)
               	);
```

Additionally, aggregating indexes integrate seamlessly with Firebolt’s partitioning strategies, further improving query performance by allowing the query engine to access only the relevant partitions. This reduces the amount of data scanned and processed, particularly when dealing with large, partitioned datasets. The combination of partition pruning and aggregate indexing helps achieve superior query performance in data-intensive environments, allowing for quicker insights and more efficient use of system resources.

##### [](#best-practices-for-aggregate-indexes)Best practices for aggregate indexes

To optimize query performance and manage resources effectively, follow these best practices for creating and maintaining aggregating indexes:

- **Focus on frequently run queries** – Create aggregating indexes for queries that use frequent aggregations including `SUM`, `COUNT`, and `AVG` that are run repeatedly.
- **Choose columns used in aggregations** – Select columns commonly used in GROUP BY and aggregation functions such as `borrower_id` or `book_id` from the examples in this section.
- **Choose columns strategically** – Build indexes on columns that are frequently queried for aggregations such as `customer_id` or `order_id` from the examples in this section. Avoid creating multiple indexes with overlapping aggregations on the same columns to minimize unnecessary overhead and costs.
- **Consider Data Freshness** – Ensure the performance gains of precomputed values outweigh the index maintenance costs for frequently changing data.

#### [](#partitions-in-tables)Partitions in tables

Firebolt’s table partitioning supports efficient data lifecycle management by organizing data based on specific criteria: date ranges, regions, product categories, or customer types. This helps streamline data retention, archiving, and access, ensuring optimal performance. Partitioning reduces the amount of data scanned during queries, improving speed, particularly when query patterns consistently filter data by columns like date or region. It also enables data pruning, skipping partitions that don’t match the query criteria, making tasks like vacuuming and archiving more efficient.

##### [](#create-a-partition)Create a partition

You can create a partition for a new table; however, it’s recommended to define the partition when creating a new table to ensure the data is organized from the start.

The following code example **creates a new table** with a primary index, and partitions the table by month:

```
CREATE TABLE librarybooks (
	transaction_id INT,
	book_id INT,
	borrower_id INT,
	checkout_date DATE,
	due_date DATE,
	return_date DATE,
	late_fee DECIMAL(10, 2)
)
PRIMARY INDEX transaction_id
PARTITION BY DATE_TRUNC('month', checkout_date);
```

In the previous example, the new `librarybooks` table is partitioned by the month of `checkout_date`, so that Firebolt can prune irrelevant partitions and speed up queries for specific date ranges.

You can **drop** an existing partition as shown in the following code example, which removes a partition for January 1, 2023:

```
ALTER TABLE borrowedbooks
DROP PARTITION '2023-01-01';
```

#### [](#suggested-indexes-and-partitions)Suggested indexes and partitions

If you understand your data and query patterns, you should select high-cardinality columns frequently used in `WHERE`, `JOIN`, or `GROUP BY` clauses to minimize the amount of data scanned and improve query efficiency.

If you don’t know how to effectively select a primary index or partition your data, you can use Firebolt’s [RECOMMEND\_DDL](/sql_reference/commands/queries/recommend_ddl.html) tool to provide automated insights. `RECOMMEND_DDL` will make recommendations to optimize database performance by analyzing your query patterns and suggesting the most efficient configurations for primary indexes and partitioning. By examining historical query data, the tool identifies columns frequently used in filtering, grouping, or aggregating operations and recommends appropriate primary indexes and partition keys. These suggestions help reduce the amount of data scanned during queries, enabling faster execution and improved resource utilization.

`RECOMMEND_DDL` is particularly useful in complex environments where query patterns evolve over time. By reviewing historical query data, Firebolt identifies columns that are frequently used in filtering or aggregation and recommends appropriate primary index and partitioning strategies.

The following code example uses `RECOMMEND_DDL` to analyze query patterns on the books table, created in the \*\*Dimension tables \** section under [Firebolt managed-tables](#firebolt-managed-tables), based on queries run in the past week:

```
CALL recommend_ddl(
  books,
  (SELECT query_text FROM information_schema.engine_query_history WHERE query_start_ts > NOW() - INTERVAL '1 week')
);
```

If the `books` table is frequently queried based on `genre` and `book_id`, Firebolt’s `RECOMMEND_DDL` command might provide the following example output:

recommended\_partition\_key recommended\_primary\_index average\_pruning\_improvement analyzed\_queries DATE\_TRUNC(‘month’, borrow\_date) book\_id, borrower\_id 0.35 200

The example output under `recommended_partition_key` suggests partitioning the `borrowedbooks` table by month based on the `borrow_date` column. The `recommended_primary_index` suggests creating a primary index on the `book_id` and `borrower_id` columns. An average pruning improvement of 0.35 indicates 35% less data will be scanned on average by applying these recommendations. The analyzed queries column shows that 200 queries were analyzed to generate these suggestions.

### [](#additional-resources)Additional resources

- [Working with tables](/Overview/indexes/using-indexes.html#firebolt-managed-tables) – An overview of how to create, manage, and optimize tables.
- [Partitions in tables](/Overview/indexes/using-indexes.html#partitions-in-tables) – How to divide large tables into smaller partitions for optimal query efficiency and performance.
- [RECOMMEND\_DDL](/sql_reference/commands/queries/recommend_ddl.html) – Information on syntax, parameters and examples of using Firebolt’s tool to automatically recommend optimal primary index and partition strategies.

* * *

- [Aggregating index](/Overview/indexes/aggregating-index.html)
- [Primary index](/Overview/indexes/primary-index.html)