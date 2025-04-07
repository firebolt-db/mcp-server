# [](#load-data)Load data

You can load data into Firebolt from an Amazon S3 bucket using two different workflows.

If you want to get started quickly, load data using a **wizard** in the **Firebolt Workspace**. If you want a more customized experience, you can write **SQL scripts** to handle each part of your workflow. This guide shows you how to load data using both the wizard and SQL, and some common data loading workflows and errors.

![You can use either the load data wizard or SQL to create a database, engine, and then load data.](../../assets/images/load_data_workflow.png)

Before you can load data, you must first register with Firebolt, then create a database and an engine. For information about how to register, see [Get Started](../../Guides/getting-started/). See the following sections for information about how to create a database and engine.

## [](#load-data-using-a-wizard)Load data using a wizard

You can use the **Load data** wizard in the **Firebolt Workspace** to load data in either CSV or Parquet format, and choose from a variety of different loading parameters which include the following:

- Specifying a custom delimiter, quote character, escape character, and other options.
- How to handle errors during data load.
- Specifying a primary index.

The **Load data** wizard guides you through the process of creating an engine and database as part of the loading process.

See [Load data using a wizard](/Guides/loading-data/loading-data-wizard.html) for information about the options available in the **Load data** wizard.

## [](#load-data-using-sql)Load data using SQL

You can use SQL to load data in CSV, Parquet, TSV, AVRO, JSON Lines or ORC formats. Prior to loading data, you must also create a database and engine using either of the following options:

- Use buttons in the **Firebolt Workspace** to create a database and engine. For more information, see the [Create a Database](/Guides/getting-started/get-started-sql.html#create-a-database) and [Create an Engine](/Guides/getting-started/get-started-sql.html#create-an-engine) sections in the [Get Started using SQL](/Guides/getting-started/get-started-sql.html) guide.
- Use the SQL commands [CREATE DATABASE](/sql_reference/commands/data-definition/create-database.html) and [CREATE ENGINE](/sql_reference/commands/engines/create-engine.html).

See [SQL to load data](/Guides/loading-data/loading-data-sql.html) for information and code examples to load data using SQL.

## [](#optimizing-during-data-loading)Optimizing during data loading

Optimizing your workflow for Firebolt starts when you load your data. Use the following guidance:

1. A primary index is a sparse index that uniquely identifies rows in a table. Having a primary index is critical to query performance at Firebolt because it allows a query to locate data without scanning an entire dataset. If you are familiar with your data and query history well enough to select an optimal primary index, you can define it when creating a table. If you don’t, you can still load your data without a primary index. Then, once you know your query history patterns, you must create a new table in order to define a primary index.
   
   You can specify primary indexes in either the **Load data** wizard or inside SQL commands. The [Load data using a wizard](/Guides/loading-data/loading-data-wizard.html) guide discusses considerations for selecting and how to select primary indexes. The [Load data using SQL](/Guides/loading-data/loading-data-sql.html) discusses considerations for selecting and shows code examples that select primary indexes. For more advanced information, see [Primary indexes](/Overview/indexes/primary-index.html).
2. If you intend to use [aggregate functions](/sql_reference/functions-reference/aggregation/) in queries, you can calculate an aggregating index when loading your data. Then queries use these pre-calculated values to access information quickly. For an example of calculating an aggregating index during load, see [Load data using SQL](/Guides/loading-data/loading-data-sql.html). For an introduction to aggregating indexes, see the [Aggregating indexes](/Guides/getting-started/get-started-sql.html#aggregating-indexes) section of the **Get Started** guide. For more information, see [Aggregating indexes](/Overview/indexes/aggregating-index.html).

## [](#next-steps)Next steps

After you load your data, you can start running and optimizing your queries. A typical workflow has the previous steps followed by data and resource cleanup as shown in the following diagram:

![The load data workflow includes using the load data wizard or SQL to create a database, engine, and then load data.](../../assets/images/get_started_workflow.png)

- [Load data using a wizard](/Guides/loading-data/loading-data-wizard.html)
- [Load data using SQL](/Guides/loading-data/loading-data-sql.html)