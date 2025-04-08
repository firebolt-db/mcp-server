# [](#get-started-using-sql)Get started using SQL

You can also use SQL to create a database and engine, and load data. If you use the **Develop Space** inside the **Firebolt Workspace**, you can customize your workflow to handle more unique workflows than with the **Load data** wizard, including loading data in TSV, Avro, JSON Lines, or ORC formats.

The following sections will guide you through a simple workflow to register, create a database and engine, load and query data, learn how to optimize your workflow, and clean up resources as shown in the following diagram:  
![A simple workflow using SQL includes registering, creating a database and engine, loading data, querying data, optimizing your workflow, cleaning up, and optionally exporting data.](../../assets/images/get_started_sql_workflow.png)

## [](#register-with-firebolt)Register with Firebolt

![The first step in the workflow is to register with Firebolt.](../../assets/images/get_started_sql_register.png)  
To get started using Firebolt, begin by registering using the following steps:

1. [Register](https://go.firebolt.io/signup) with Firebolt. Fill in your email, name, choose a password, and select ‘Get Started’.
2. Firebolt will send a confirmation to the address that you provided. To complete your registration, select **Verify** in the email to take you to Firebolt’s [login page](https://go.firebolt.io/login).
3. Type in your email and password and select **Log In**.

New accounts receive credits ($200) to get started exploring Firebolt’s capabilities. These credits must be used within 30 days of account creation.

Firebolt’s billing is based on engine runtime, measured in seconds. We also pass through AWS S3 storage costs at the rate of $23 per TB per month. The amount that you spend is dependent primarily on which engines you use and how long those engines are running.

You can view your total cost in FBU up to the latest second and in $USD up to the latest day. For more information, see the following **Create a Database** section. For more information about costs, see [Data Warehouse Pricing](https://www.firebolt.io/pricing). If you need to buy additional credits, connect Firebolt with your AWS Marketplace account. For more information about AWS Marketplace, see the following section: [Registering through AWS Marketplace section](/Guides/getting-started/get-started-next.html#register-through-the-aws-marketplace).

## [](#create-a-database)Create a Database

![After registering, create a database.](../../assets/images/get_started_sql_database.png)  
Firebolt decouples storage and compute resources so that multiple engines can run computations on the same database. You can also configure different engine sizes for different workloads. These workloads can run in parallel or separately. Because storage is decoupled from compute, you must first create both a database and an engine before you can run your first query.

Firebolt’s structure is organized as follows:

- A database holds the elements that you need to run queries such as tables, views and information schema.
- An [engine](/Overview/engine-fundamentals.html) provides the compute resources for ingesting data and running queries. For more information on using Firebolt engines and how to select the correct size for your workload, see [Operate engines](/Guides/operate-engines/operate-engines.html).

If you used the **Load data** wizard, Firebolt has already created a database for you, and you can skip creating a database.

The following instructions show you how to create a database and then an engine. Note that you can also create the engine first.

1. In the left navigation pane, select the **+** to the right **Databases**.
2. Select **Create new database**.
3. Enter the name for your database in the **Database Name** field. For this example, use “tutorial\_database” as your database name. In Firebolt, the names of engines and databases are **case-sensitive**. If you are using uppercase characters in their names, enclose their name inside double quotes (“) when you refer to them in SQL.

Firebolt creates a new database with the following two default schemas:

- **Public** - A namespace where you can create and manage your database objects including tables, engines and queries. The default schema includes **tables**, **external tables**, and **views**.
- **Information\_schema** - A standardized set of read-only views that provide metadata about database objects including tables, engines, cost information, and queries.

You can find these schema by selecting your database under **Databases** in the left navigation pane. Next to the name of your database, select the drop-down arrow to expand and view the schemas and their contents. You can view your total cost in FBU up to the latest second and in $USD up to the latest day in **Information\_schema**.

If you’re using the **Develop Space**, expand **Information\_schema**, and then **Views** to show the following:

- **engine\_metering\_history** - contains information about billing cost in FBU up to the latest second in **consumed\_fbu**.
- **engine\_billing** - contains information about billing cost in US dollars up to the latest day in **billed\_cost**.

To see values for the previous costs, select the **More options** icon (![More options icon](../../assets/images/more_options_icon.png)) next to either **consumed\_fbu** or **billed\_cost**, Then select **Preview data**. You can also run a query in the script tab as shown in the following code example:

```
SELECT * 
FROM information_schema.engine_metering_history 
```

## [](#create-an-engine)Create an Engine

![After creating a database, create an engine.](../../assets/images/get_started_sql_engine.png)  
To process a query, you must use an engine. You can either create an engine based on the following recommendations, or use the system engine. You can only use the system engine to run metadata-related queries, but it is always running, so you don’t have to wait for it to start. You can use the system engine to process data in any database. If you create your own engine, there is a small start up time associated with it.

Firebolt recommends the following initial engine configurations based on where you are in your exploration of Firebolt’s capabilities. An FBU stands for a Firebolt Unit, and is equivalent to 35 US cents.

Each FBU is related to the amount of time as follows:

Task Expected Usage Ingest initial data 4-16 FBU Run test queries 8-32 FBU Find optimal query performance 32-240 FBU Find optimal test integrations 32-240 FBU

Each engine node can cache data locally to improve performance.

Small and medium engines are available for use right away. If you want to use a large or extra-large engine, reach out to support@firebolt.io. The default engine configuration uses a small node, which is sufficient for this tutorial. To learn more about how to select the correct engine size for your workload, see [Sizing Engines](/Guides/operate-engines/sizing-engines.html).

By default, when you login to **Firebolt’s Workspace** for the first time, Firebolt creates a tab in the **Develop Space** called **Script 1**. The following apply:

- The database that **Script 1** will run using is located directly below the tab name. If you want to change the database, select another database from the drop-down list.
- An engine must be running to process the script in a selected tab. The name and status of the engine that **Script 1** uses for computation is located to the right of the current selected database. To change either the engine or the status, select the drop-down arrow next to the engine name. You can select a new engine and change its status from **Stopped** to **Running** by selecting **Start engine**. If you select **Run** at the bottom of the workspace, the selected engine starts automatically. Select **Stop engine** to change the status to **Stopped**. Firebolt automatically stops your engine if it is inactive for 20 minutes.

Because an engine is a dedicated compute node that nobody else can use, you are charged for each second that your engine is **Running**, even if it’s not processing a query.

If you used the **Load data** wizard, Firebolt has already created an engine for you, and you can skip the following step.

1. Select the **(+)** icon next to **Databases**.
2. Select **Create new engine** from the drop-down list.
3. Enter the name of your engine in the **New engine name** text box. For this example, enter “tutorial\_engine” as your engine name.

## [](#load-data)Load Data

![After creating an engine, you can load your data.](../../assets/images/get_started_sql_load.png)

After creating an engine, you can load your data. This tutorial uses Firebolt’s publicly available Firebolt’s sample dataset, from the fictional [“Ultra Fast Gaming Inc.”](https://help.firebolt.io/t/ultra-fast-gaming-firebolt-sample-dataset/250) company. This dataset does not require access credentials. If your personal dataset requires access credentials, you will need to provide them. For examples of how to provide access credentials and more complex loading workflows, see [Loading data](/Guides/loading-data/loading-data.html). For more information about AWS access credentials, see [Creating Access key and Secret ID](/Guides/loading-data/creating-access-keys-aws.html)

If you used the **Load data** wizard, skip ahead to the following **Run query** section.

Use [COPY FROM](/sql_reference/commands/data-management/copy-from.html) in the **Develop Space** to copy data directly from a source into a Firebolt managed table.

1. Enter the following into the **Script 1** tab to load data using the following steps:
   
   ```
     COPY INTO tutorial FROM 's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/levels.csv' WITH HEADER=TRUE;
   ```
   
   For examples of more complex loading workflows, see [Load data](/Guides/loading-data/loading-data.html).
2. Select **Run**.
3. In the left navigation pane under the **Tutorial\_Database**, **Tables** now contains the **tutorial** table.
4. Expand the drop down menu next to **Columns** to view the name and data format of each column.
5. Select the **More options** icon (![More options icon](../../assets/images/more_options_icon.png)) next to the data type of each column name to open a pop-up that allows you to insert the name of the column into your SQL script. You can also select **Preview data**.
6. To view the contents of the **tutorial** table, run a SELECT query as shown in the following code example. To run this in a new tab, select the (**+**) icon next to the **Script 1** tab.
   
   ```
     SELECT
    *
     FROM
    tutorial
   ```
7. Select **Run**. The bottom of your workspace includes information about your processing job in the following tabs:
   
   - The **Results** tab at the bottom of your **Develop Space** shows the contents returned by your query. After running the previous SELECT statement, the **Results** tab should display column names and values for the data in the tutorial.
   
   <!--THE END-->
   
   - Select the filter icon (![Filter icon](../../assets/images/filter-icon.png)) to change which columns are shown.
   - Select the **More options** icon (![More options icon](../../assets/images/more_options_icon.png)) to export the contents of the **Results** tab to a JSON or CSV file.
   - The Statistics tab shows information about running your query including how long it took to run and its status. After running the previous SELECT statement, the **Statistics** tab shows the status of the statement, its STATUS as having succeeded or failed, how long it took to run the query, the number of rows processed, and the amount of data scanned.
   - Select the **More options** icon (![More options icon](../../assets/images/more_options_icon.png)) to export the contents of the **Statistics** tab to a JSON or CSV file.
   - The **Query Profile** tab contains metrics for each operator used in your query and a **Query id**. Select an operation to view its metrics. These metrics include the following:
     
     - The output cardinality - the number of rows that each operator produced.
     - The thread time - the sum of the wall clock time that threads spent to run the selected operation across all nodes.
     - The CPU time - the sum of the time that threads that ran the operator were scheduled on a CPU core.
     - The output types - the data types of the result of the operator.

You can use these metrics to analyze and measure the efficiency and performance of your query. For example, If the CPU time is much smaller than thread time, the input-output (IO) latency may be high or the engine that you are using may be running multiple queries at the same time. For more information, see [Example with ANALYZE](/sql_reference/commands/queries/explain.html).

- The **Engine monitoring** tab shows monitoring information including the percent CPU, memory, disk use and cache read. Information is shown from the last 5 minutes by default. Select a different time interval from the drop-down menu next to **Last 5 minutes**. You can also select the **Refresh** icon next to the drop-down menu to update the graphical information.
- The **Query history** tab shows detailed information associated with each query, listed by its **Query id**. This information includes the query status, start time, number of rows and bytes scanned during the load, user and account information. You can choose the following options at the top of the bottom panel:
  
  - Select the **Refresh** icon to update the query history and ID.
  - Select the filter icon (![Filter data icon](../../assets/images/filter-icon.png)) to remove or add columns to display.
  - Select the **More options** icon (![More options icon](../../assets/images/more_options_icon.png)) to export the contents of the **Query history** tab to a JSON or CSV file.

For more information about Firebolt’s **Develop Space**, see [Using the develop workspace](/Guides/query-data/using-the-develop-workspace.html).

## [](#run-query)Run Query

![After loading your data, you can run a query.](../../assets/images/get_started_sql_query.png)

To run a query on your data, do the following:

1. Select the (**+**) icon next to the **Script 2** tab to open a new tab.
2. Enter the following simple query, which fetches a list of databases associated with your account:
   
   ```
     SHOW CATALOGS;
   ```
3. Select **Run** to process the query. Firebolt uses the engine listed to the right of your database to run your query and its status of **Running** or **Stopped**. You can select a different engine from the dropdown menu next to the engine (![Engine icon](../../assets/images/engine-icon.png)) icon.
   
   If your engine is **Stopped**, Firebolt may prompt you to start your engine. Select **Start Engine**. Engine startup typically requires a few moments to complete, as Firebolt prepares your environment for data analysis.

For more information about Firebolt’s **Develop Space**, see [Use the Develop Space](/Guides/query-data/using-the-develop-workspace.html).

## [](#optimize-your-workflow)Optimize your workflow

![After running a baseline query, you can optimize your workflow for better performance.](../../assets/images/get_started_sql_optimize.png)

Firebolt uses a number of optimization strategies to reduce query times. Over small datasets like those specified in this guide, the uplift may not be noticeable. However, these strategies can **dramatically improve** query performance for larger datasets. The following sections discuss how [primary indexes](#primary-indexes) and [aggregating indexes](#aggregating-indexes) to do the following:

- Reduce the amount of data that the query scans.
- Pre-calculate values that are used repeatedly during computations.

### [](#primary-indexes)Primary Indexes

One of Firebolt’s key optimization strategies is to select a primary index for columns that are used frequently in `WHERE`, `JOIN`, `GROUP_BY`, and clauses used for sorting. In Firebolt, a primary index is a type of **sparse index**. Thus, selecting the best primary index can reduce query run times significantly by reducing the data set that the query searches over. Selecting primary indexes also allows Firebolt to manage updates, deletions and insertions to tables and provide optimal query performance.

If you have a composite primary index, the order that the columns are listed is important. Specify the column that has a large number of unique values, or high cardinality, first, followed by columns with lower cardinality. A sort order with the previous characteristics allows Firebolt to prune, or eliminate irrelevant data, so that it doesn’t have to scan it in query processing. Pruning significantly enhances query performance.

You can create a primary index **only** when you create a table. If you want to change the primary index, you must create a new table. The following example shows how to use [CREATE TABLE](/sql_reference/commands/data-definition/create-fact-dimension-table.html) to create a new `levels` table, define the schema, and set two primary indexes:

```
CREATE TABLE IF NOT EXISTS levels (
    "LevelID" INT,
    "Name" TEXT,
    "GameID" INT,
    "LevelType" TEXT,
    "MaxPoints" INT,
    "PointsPerLap" DOUBLE,
    "SceneDetails" TEXT
)
PRIMARY INDEX "LevelID", "Name";
```

In the previous code example, the primary index contains two values. The first value, `LevelID`, is required in order to create a primary index. The second value, `Name`, and any following values are optional. Firebolt will use all listed primary indexes to optimize query scans. If Name has lower cardinality than `LevelID`, then Firebolt can optimize these indexes to eliminate scanning over irrelevant data. For more information about primary indexes and sort order, see [Primary index](/Overview/indexes/primary-index.html).

To read data into the `levels` table, enter the following into a new script tab:

```
COPY INTO levels
FROM 's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/levels.csv'
WITH TYPE = CSV
HEADER = TRUE;
```

### [](#aggregating-indexes)Aggregating Indexes

Another key optimization strategy includes pre-calculating aggregate values for columns that are frequently used in functions that combine data such as `COUNT`, `SUM`, `MAX`, `MIN`, `AVG`, `JOIN`, and `GROUP BY`. Rather than computing aggregate values each time they are used in a calculation, the results are accessed from storage, which helps run queries quickly and saves compute resources.

An aggregating index combines columns into a statistical result. You can calculate an aggregate index on an entire table, or more efficiently, calculate them over a subset of table columns. You can also use your knowledge of which dimensions and aggregate functions are used most often for your use case to predefine what table dimensions and which aggregate functions to use.

Once you create aggregate indexes, Firebolt maintains them automatically for you. If you load new data into your table or alter it, your aggregate indexes are automatically updated. You can also have multiple aggregate indexes for a single table. When you query a table with multiple aggregate indexes, Firebolt will automatically select the best index to use to optimize performance.

From the **tutorial** table that you created in the previous step, assume you want to run a query to look at the AVG(NumberOfLaps), grouped by LevelType. The following example code shows you how to create an aggregating index **levels\_agg\_idx** on the **LevelType** column to pre-calculate the average number of laps for each level.

```
CREATE AGGREGATING INDEX
  levels_agg_idx
ON tutorial (
  "LevelType" 
  , AVG("NumberOfLaps")
  );
```

After you run the script, the `levels_agg_idx` aggregate index listed in the left navigation pane under **Indexes** in the **tutorial** table. Any queries that run over the tutorial table that use an average of the **NumberOfLaps** column grouped by **LevelType** will now use the `levels_agg_idx` index instead of reading the entire table to calculate it.

For more information, see [Aggregating index](/Overview/indexes/aggregating-index.html).

### [](#warm-data-and-cache-eviction)Warm data and cache eviction

Another key optimization strategy is to read warm data, or data accessed from cache, rather than reading in “cold” data from an Amazon S3 bucket. Querying cold data can be significantly slower than querying warm data, particularly for large datasets that contain millions of rows or more. If you reach about 80% of your available cache, the least recently used data will be moved out of cache into an Amazon S3 bucket.

#### [](#warm-data)Warm data

When data is warm, Firebolt transfers data from remote storage in Amazon (S3) to a local (cache). Data is automatically warmed when you access it during a query, and stored in a solid state drive (SSD) cache. However, when you query data to warm it, you use an engine, and incur [engine consumption](/Overview/engine-consumption.html) costs. Therefore, you should use filters to warm only the data that you need to access frequently in your queries.

The following guidance applies:

- If you need access to all the data in a table, use `CHECKSUM` to warm the entire table as follows:
  
  ```
  SELECT CHECKSUM(*) FROM levels;
  ```
- If you only need a few columns that meet a certain criteria, filter them before warming the data as shown in the following code example:
  
  ```
  SELECT CHECKSUM("Name", "MaxPoints") FROM levels WHERE "MaxPoints" BETWEEN 100 AND 250;
  ```
- If you have a large dataset, you can divide the data into smaller segments, and execute the queries in parallel, as shown in the following code example:
  
  ```
  SELECT CHECKSUM("Name", "MaxPoints") FROM levels WHERE "MaxPoints" BETWEEN 0 AND 100;
  SELECT CHECKSUM("Name", "MaxPoints") FROM levels WHERE "MaxPoints" BETWEEN 101 AND 200;
  SELECT CHECKSUM("Name", "MaxPoints") FROM levels WHERE "MaxPoints" > 200;
  ```

#### [](#cache-eviction)Cache eviction

When your cache usage exceeds approximately 80% of its capacity, Firebolt automatically evicts the some data from the cache. If you query this data later, Firebolt reloads it into the cache before processing the query. The total available cache size depends on your engine’s size and family. Larger engine sizes provide more cache space, and the storage-optimized family offers more cache than the compute-optimized family. Small and medium sized engines are available for use right away. If you want to use a large or extra-large engine, reach out to [support@firebolt.io](mailto:support@firebolt.io).

You can check the size of your cache using the following example code:

```
SHOW CACHE;
```

The previous code example shows your cache usage in GB per total cache available.

When data is loaded into Firebolt, it is stored in units of data storage called tablets. A tablet contains a subset of a table’s rows and columns. If you reach your cache’s 80% capacity, the entire tablet that contains the least recently used data, is evicted.

The following code example shows how to view information about the tablets that are used to store your table including the number of uncompressed and compressed bytes on disk:

```
SELECT * FROM information_schema.engine_tablets  where table_name = 'levels';
```

## [](#clean-up)Clean up

![To avoid incurring costs, clean up data and resources when you're finished with your workflow.](../../assets/images/get_started_sql_clean.png)

After you’ve completed the steps in this guide, avoid incurring costs associated with the getting started exercises by doing the following:

- Stop any running engines.
- Remove data from storage.

### [](#stop-any-running-engines)Stop any running engines

Firebolt shows you the status of your current engine next to the engines icon (![Engine icon](../../assets/images/engine-icon.png)) under your script tab as either **Stopped** or **Running**. To shut down your engine, select your engine from the drop-down list next to the name of the engine, and then select one of the following:

- Stop engine - Allow all of the currently running queries to finish running and then shut down the engine. Selecting this option will allow the engine to run for as long as it takes to complete all queries running on the selected engine.
- Terminate all queries and stop - Stop the engine and stop running any queries. Selecting this option stops the engine in about 20-30 seconds.

### [](#remove-data-from-storage)Remove data from storage

To remove a table and all of its data, enter [DROP TABLE](/sql_reference/commands/data-definition/drop-table.html) into a script tab, as shown in the following code example:

```
DROP TABLE levels
```

To remove a database and all of its associated data, do the following in the Firebolt **Develop Space**:

- Select the database from the left navigation bar.
- Select the **More options** (![More options icon](../../assets/images/more_options_icon.png)) icon.
- Select **Delete database**. Deleting your database will permanently remove your database from Firebolt. You cannot undo this action. Select **Delete**.

## [](#export-data)Export data

![The following section shows how to optionally export data after cleaning up.](../../assets/images/get_started_sql_export.png)

If you want to save your data outside of Firebolt, you can use [COPY TO](/sql_reference/commands/data-management/copy-to.html) to export data to an external table. This section shows how to set the minimal AWS permissions and use `COPY TO` to export data to an [AWS S3 bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/Welcome.html). You may have to reach out to your administrator to obtain or change AWS permissions.

### [](#set-permissions-to-write-to-an-aws-bucket)Set permissions to write to an AWS bucket

Firebolt must have the following permissions to write to an AWS S3 bucket:

1. AWS access key credentials. The credentials must be associated with a user with permissions to write objects to the bucket. Specify access key credentials using the following syntax:

```
    CREDENTIALS = (AWS_ACCESS_KEY_ID = '<aws_access_key_id>' AWS_SECRET_ACCESS_KEY = '<aws_secret_access_key>')
```

In the previous credentials example, `<aws_access_key_id>` is the AWS access key ID associated with an AWS user or an IAM role. An access key ID has the following form: `AKIAIOSFODNN7EXAMPLE`. The value `<aws_secret_access_key>` is the AWS secret access key. A secret access key has the following form: `wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`.

1. An AWS IAM policy statement attached to a user role. Firebolt requires the following minimum permissions in the IAM policy:
   
   ```
    {
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "s3:Get*",
                "s3:List*",
                "s3:PutObject",
                "s3:DeleteObject"
            ],									
            "Resource": [
                "arn:aws:s3:::my_s3_bucket",
                "arn:aws:s3:::my_s3_bucket/*"
            ]
        }
     ]
    }
   ```
   
   For more information about AWS access keys and roles, see [Creating Access Key and Secret ID in AWS](/Guides/loading-data/creating-access-keys-aws.html).

### [](#export-to-an-aws-bucket)Export to an AWS bucket

Use [COPY TO](/sql_reference/commands/data-management/copy-to.html) select all the columns from a table and export to an AWS S3 bucket as shown in the following code example:

```
COPY (SELECT * FROM test_table)
  TO 's3://my-bucket/path/to/data'
  CREDENTIALS = 
  (AWS_ROLE_ARN= 'arn:aws:iam::123456789012:role/my-firebolt-role');
```

In the previous code example, the role ARN ([Amazon Resource Name](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html)) identifies the AWS IAM role that specifies the access for users or services. An ARN follows the following structure: arn:aws:iam::account-id:role/role-name. Because TYPE is omitted from `COPY TO`, the file or files will be written in the default CSV format. Because `COMPRESSION` is also omitted, the output data is compressed using GZIP (\*.csv.gz) format.

Firebolt assigns a query ID, that has the following example format `16B903C4206098FD`, to the query at runtime. If the size of the compressed output exceeds the default of `16` MB, Firebolt writes multiple GZIP files. In the following example, the size of the output is `40` MB, so Firebolt writes `4` files.

```
s3://my_bucket/my_fb_queries/
16B903C4206098FD_0.csv.gz
16B903C4206098FD_1.csv.gz
16B903C4206098FD_2.csv.gz
16B903C4206098FD_3.csv.gz
```

## [](#next-steps)Next steps

To continue learning about Firebolt’s architecture, capabilities, using Firebolt after your trial period, and setting up your organization, see [Resources beyond getting started](/Guides/getting-started/get-started-next.html).