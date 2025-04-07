# [](#load-data-using-sql)Load data using SQL

If the **Load data** wizard does not meet your needs, or you prefer to write directly in SQL, you can enter SQL and run it in the **Firebolt Workspace**, or use an API.

Before you can load data using a SQL script, you must register with Firebolt, and create a database and an engine.

A general workflow to load data using SQL is shown in the following diagram, with the highlighted path representing the SQL workflow and the muted path representing using the **Load data** wizard:

![You can use either the load data wizard or SQL to create a database, engine, and then load data.](../../assets/images/load_data_sql_workflow.png)

For more information on how to register, create a database and engine using the **Firebolt Workspace**, see the [Get Started](/Guides/getting-started/) guide. To create an engine using SQL, use [CREATE ENGINE](/sql_reference/commands/engines/create-engine.html). You can check how many engines are defined in your current account using [SHOW ENGINES](/sql_reference/commands/metadata/show-engines.html). For more information and examples of how to create engines, see [Work with engines using DDL](/Guides/operate-engines/working-with-engines-using-ddl.html). To create a database, use [CREATE DATABASE](/sql_reference/commands/data-definition/create-database.html). You can check how many databases (i.e., catalogs) are defined in your current account using [SHOW CATALOGS](/sql_reference/commands/metadata/show-catalogs.html). Next, log into the **Firebolt Workspace** and enter SQL into the script tab in the **SQL Editor**.

The following code examples show different workflows based on need and complexity:

- [The simplest COPY FROM workflow](#the-simplest-copy-from-workflow)
- [Define a schema, create a table, and load data](#define-a-schema-create-a-table-and-load-data)
- [Load multiple files into a table](#load-multiple-files-into-a-table)
- [Filter data before loading using OFFSET and LIMIT](#filter-data-before-loading-using-offset-and-limit)
- [Aggregating data during data load](#aggregating-data-during-data-load)
- [Update an existing table from an external table](#update-an-existing-table-from-an-external-table)
- [Load source file metadata into a table](#load-source-file-metadata-into-a-table)
- [Continue loading even with errors](#continue-loading-even-with-errors)
- [Log errors during data load](#log-errors-during-data-load)

## [](#the-simplest-copy-from-workflow)The simplest COPY FROM workflow

Although there are many options to handle different data loading workflows, `COPY FROM` requires only two parameters:

1. The name of the table that you are loading data into.
2. A location to load the data from.

An example of the **simplest** way to invoke `COPY FROM` is:

```
COPY INTO tutorial FROM 
's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/levels.csv' WITH HEADER=TRUE;
```

The previous code creates a table named `tutorial`, reads a CSV file with headers from a public Amazon S3 bucket, automatically generates a schema, and loads the data.

If the data is contained in an Amazon S3 bucket with restricted access, you will need to provide credentials. The following example shows how to provide credentials and read a file with headers, and automatically generate a schema:

```
COPY INTO tutorial 
FROM 's3://your_s3_bucket/your_file.csv'
WITH
CREDENTIALS = (
    AWS_ROLE_ARN='arn:aws:iam::123456789012:role/my-firebolt-role'
)
HEADER=TRUE AUTO_CREATE=TRUE;
```

Firebolt supports authentication using both permanent AWS access keys and temporary security credentials obtained through Amazon’s [AssumeRole](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html) feature. To provide your credentials for the previous example, follow these steps:

#### [](#static-credentials)Static Credentials

Replace `<aws_access_key_id>` with an AWS access key ID associated with an AWS user or IAM role. The access key ID is a 20-character string (e.g., AKIAIOSFODNN7EXAMPLE). Replace `<aws_secret_access_key>` with the AWS secret access key associated with the AWS user or IAM role. The secret access key is a 40-character string (e.g., wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY). You can also specify an `AWS_SESSION_TOKEN`.

**Example:**

```
COPY INTO tutorial 
FROM 's3://test-bucket/data.csv'
WITH
CREDENTIALS = (
    AWS_ACCESS_KEY_ID = 'AKIAIOSFODNN7EXAMPLE' AWS_SECRET_ACCESS_KEY = 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY'
)
```

#### [](#assume-role-authentication)Assume Role Authentication

Replace with your role's Amazon Resource Name (ARN) of the IAM role that you want Firebolt to assume. This method gives Firebolt temporary credentials to authenticate and access your Amazon S3 bucket.

**Example:**

```
COPY INTO tutorial 
FROM 's3://test-bucket/data.csv'
WITH
CREDENTIALS = (
    AWS_ROLE_ARN='arn:aws:iam::746669185839:role/example'
)
```

## [](#define-a-schema-create-a-table-and-load-data)Define a schema, create a table, and load data

You can also load data into an existing table using your own schema definition. Manually defining your own schema, can give you finer control over data ingestion. This example contains the following two steps:

1. Create the target table.
   
   Create a table to load the data into, as shown in the following code example:
   
   ```
    CREATE TABLE IF NOT EXISTS levels (
    LevelID INT,
    Name TEXT,
    GameID INT,
    LevelType TEXT,
    MaxPoints INT,
    PointsPerLap DOUBLE,
    SceneDetails TEXT
   );
   ```
   
   The previous code example creates a table named `levels`, and defines each of the columns with a name and data type. For more information about the data types that Firebolt supports, see [Data types](/sql_reference/data-types.html).
2. Run COPY FROM.
   
   Use COPY FROM to load the data from an Amazon S3 bucket into the levels table, as shown in the following code example:
   
   ```
    COPY INTO levels
    FROM 's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/levels.csv'
    WITH TYPE = CSV
    HEADER = TRUE;
   ```
   
   The previous code example reads data from a Firebolt test data set from the fictional [Ultra Fast Gaming Inc.](https://help.firebolt.io/t/ultra-fast-gaming-firebolt-sample-dataset/250) company. The `levels` data set is in CSV format, but you can also use `COPY FROM` to read files in `Parquet` format. If you are reading in a `CSV` file and specify `HEADER = TRUE`, then Firebolt expects the first line of your file to contain column names.

## [](#load-multiple-files-into-a-table)Load multiple files into a table

You can use the `PATTERN` option in `COPY FROM` to load several files at the same time from an Amazon S3 bucket. The `PATTERN` option uses standard regular expressions. For more information about regular expressions, see the Wikipedia [glob programming](https://en.wikipedia.org/wiki/Glob_%60%28programming%60%29) article.

```
COPY INTO nyc_restaurant_inspections FROM 
's3://firebolt-sample-datasets-public-us-east-1/nyc_sample_datasets/nyc_restaurant_inspections/parquet/'
WITH PATTERN="*.parquet" AUTO_CREATE=TRUE TYPE=PARQUET;
```

In the previous code example, the following apply:

- **COPY INTO**: Specifies the target table to load the data into.
- **FROM**: Specifies the S3 bucket location of the data.
- **WITH PATTERN**= “\*.parquet”: Uses a regular expressions pattern with wildcards (\*) to include all Parquet files in the directory.
- **AUTO\_CREATE=TRUE**: Automatically creates the table and the schema if the table does not already exist. Parquet files include rich data, and typically have schema information for simple and high-fidelity schema creation. Specifying AUTO\_CREATE to TRUE ensures the schema in the Parquet file is preserved after loading.
- **TYPE = PARQUET**: Specifies the data format as Parquet.

## [](#filter-data-before-loading-using-offset-and-limit)Filter data before loading using OFFSET and LIMIT

You can use `COPY FROM` with the `LIMIT` and `OFFSET` clauses to filter out data before you load it into Firebolt. The following example demonstrates how to filter the source data by skipping the first five rows of data and inserting only the next three rows.

```
COPY offset_limit
FROM 's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/levels.csv'
OFFSET 5 LIMIT 3
WITH TYPE = CSV HEADER = TRUE;
```

In the previous code example, the following apply:

- **OFFSET**: Specifies a non-negative number of rows that are skipped before returning results from the query.
- **LIMIT**: Restricts the number of rows that are included in the result set.
- **TYPE = CSV**: Specifies the data format as CSV.
- **HEADER**: Specifies that the first row of the source file contains column headers.

For more information about `OFFSET` and `LIMIT`, see [SELECT Query Syntax](/sql_reference/commands/queries/select.html).

## [](#aggregating-data-during-data-load)Aggregating data during data load

If you frequently use [aggregation functions](/sql_reference/functions-reference/aggregation/) such as `COUNT`, `MAX`, or `SUM`, you can perform these aggregations on top of an external table without loading the raw data into Firebolt. This approach allows you to avoid costs associated with importing and storing the dataset, particularly if you don’t need to store the originating data set.

The following example shows how to aggregate data using an [external table](/Guides/loading-data/working-with-external-tables.html). Then, define a table in the Firebolt database with the desired aggregations. Finally, insert data from the external table into the internal table. This example contains the following three steps:

1. Create an external table linked to files in an Amazon S3 bucket.
   
   The following code creates an external table that links to files in Amazon S3 bucket. The table has a defined schema that matches the type and names of the originating data:
   
   ```
    CREATE EXTERNAL TABLE IF NOT EXISTS ex_playstats (
      GameID INTEGER,
      PlayerID INTEGER,
      StatTime TIMESTAMPNTZ,
      SelectedCar TEXT,
      CurrentLevel INTEGER,
      CurrentSpeed REAL,
      CurrentPlayTime BIGINT,
      CurrentScore BIGINT,
      Event TEXT,
      ErrorCode TEXT,
      TournamentID INTEGER) 
    URL = 's3://firebolt-sample-datasets-public-us-east-1/gaming/parquet/playstats/'
    OBJECT_PATTERN = '*'
    TYPE = (PARQUET);
   ```
   
   The previous code uses `OBJECT_PATTERN` to link all (\*) files inside the specified directory contained in `URL`, and `TYPE` to specify the file format.
2. Define a table in the Firebolt database with the desired aggregations, as shown in the following code example:
   
   ```
    CREATE TABLE IF NOT EXISTS playstats_max_scores (
      PlayerID INTEGER,
      TournamentID INTEGER,
      MaxCurrentLevel INTEGER,
      MaxCurrentSpeed REAL,
      MaxCurrentScore BIGINT
    ) PRIMARY INDEX TournamentID, PlayerID;
   ```
   
   The previous code creates a table with the aggregate values `MaxCurrentLevel`, `MaxCurrentSpeed`, and `MaxCurrentScore`.
3. Insert data from the external table into the internal table using the aggregate functions, as shown in the following code example:

```
    INSERT INTO playstats_max_scores 
    SELECT PlayerID,  
      TournamentID,
      MAX(CurrentLevel),
      MAX(CurrentSpeed),
      MAX(CurrentScore)
    FROM ex_playstats
    GROUP BY ALL
```

The previous code calculates the aggregate function `MAX` before loading the data into the `playstats_max_scores` table.

## [](#update-an-existing-table-from-an-external-table)Update an existing table from an external table

Firebolt saves metadata including virtual columns, and the source file’s name, size and timestamp when mapping data from an Amazon S3 bucket to a Firebolt database. You can query this metadata directly for troubleshooting and analysis, or use it to find new data, as shown in this example.

To load only new and updated data from an Amazon S3 bucket into an existing table, use an external table and two temporary tables. This section guides you through creating a new table, which will serve as the existing table in a complete example. If you already have an existing table, its schema definition must include the file timestamp and file name metadata. For more information about these metadata columns, see **Using metadata virtual columns** in [Work with external tables](/Guides/loading-data/working-with-external-tables.html).

The full workflow involves creating an internal source data table, an external table linked to the source data, and two temporary tables for the latest timestamp and updated data. The `updates_table` selects new data and uses an inner join to insert these records into your existing table, as illustrated in the diagram below:

![Use an external table and two temporary tables to update a main internal table by timestamp.](../../assets/images/workflow_update_from_external_table.png)

This example contains the following nine steps:

1. Create a table
   
   The following code example shows you how to create a `players` table from a sample players dataset, and then copy data from a parquet file in an Amazon S3 bucket into it:
   
   ```
    CREATE TABLE IF NOT EXISTS
    players (
      PlayerID INTEGER,
      Nickname TEXT,
      Email TEXT,
      AgeCategory TEXT,
      Platforms ARRAY (TEXT NULL),
      RegisteredOn DATE,
      IsSubscribedToNewsletter BOOLEAN,
      InternalProbabilityToWin DOUBLE PRECISION,
      SOURCE_FILE_NAME TEXT,
      SOURCE_FILE_TIMESTAMP TIMESTAMPNTZ)
      PRIMARY INDEX agecategory, registeredon;
   ```
   
   The previous code example defines the schema for the players table, which includes the [metadata columns](/Guides/loading-data/working-with-external-tables.html) `SOURCE_FILE_NAME` and `SOURCE_FILE_TIMESTAMP`.
2. Create an external table
   
   Use an external table to query the source data directly to compare it to data in your existing table. The advantages of using an external table to check for new data are as follows:
   
   - An external table links to the data source without loading it into a database, which avoids costs associated with importing and storing it.
   - Using an external table isolates data operations to reduce the risk of corrupting data contained in the main `players` table.
   
   The following code example creates an external players\_ext table linked to the source data:
   
   ```
    CREATE EXTERNAL TABLE IF NOT EXISTS
    players_ext (
      PlayerID INTEGER,
      Nickname TEXT,
      Email TEXT,
      AgeCategory TEXT,
      Platforms ARRAY (TEXT NULL),
      RegisteredOn PGDATE,
      IsSubscribedToNewsletter BOOLEAN,
      InternalProbabilityToWin DOUBLE PRECISION)
    URL = 's3://firebolt-sample-datasets-public-us-east-1/gaming/parquet/players/'
    OBJECT_PATTERN = '*'
    TYPE = (PARQUET);
   ```
   
   The previous code example defines the schema for parquet data and links to all parquet files in the Amazon S3 bucket that contains the source data.
   
   If you are using an external table to link to data in parquet format, the order of the columns in the external table does not have to match the order of the columns in the source data. If you are reading data in csv format, the order must match the order in the source data.
3. Copy data from the `players_ext` external table into an internal `players` table, as shown in the following code example:
   
   ```
    INSERT INTO players
    SELECT *,
      $SOURCE_FILE_NAME,
      $SOURCE_FILE_TIMESTAMP
    FROM players_ext
   ```
4. Create a temporary table that contains the most recent timestamp from your existing table, as shown in the following code example:
   
   ```
    CREATE TABLE IF NOT EXISTS control_maxdate AS (
    SELECT MAX(source_file_timestamp) AS max_time
    FROM players
    );
   ```
   
   The previous code example uses an aggregate function `MAX` to select the most recent timestamp from the existing `players` table.
5. Create a temporary table to select and store data that has a newer timestamp than that contained in the control\_maxdate table, as shown in the following code example:
   
   ```
    CREATE TABLE IF NOT EXISTS updates_table AS (
    WITH external_table AS (
    SELECT *,
      $SOURCE_FILE_NAME AS source_file_name_new,
      $SOURCE_FILE_TIMESTAMP AS source_file_timestamp_new,
    FROM players_ext
    WHERE $source_file_timestamp > (SELECT max_time FROM control_maxdate)
    AND playerid IN (SELECT DISTINCT playerid FROM players)
    )
    SELECT
    e.*
    FROM players f
    INNER JOIN external_table e
    ON f.playerid = e.playerid
    );
   ```
   
   The previous code example creates an `updates_table` using a `SELECT` statement to filter out data that is older than the previously recorded timestamp. The code includes a table alias `e`, which refers to `external_table`, and the table alias `f`, which refers to the `players` table. The `INNER JOIN` uses `playerid` to match rows in the external table to those in the `player` table, and then updates the `players` table.
6. Delete records from the original players table that have been updated, based on matching player IDs in the `updates_table`, as shown in the following code example:
   
   ```
    DELETE FROM players
    WHERE playerid IN (SELECT playerid FROM updates_table);
   ```
7. Insert updated records from the `updates_table`, including a new timestamp, into the `players` table to replace the deleted records from the previous step, as shown in the following example:
   
   ```
    INSERT INTO players
    SELECT
      playerid,
      nickname,
      email,
      agecategory,
      platforms,
      registeredon,
      issubscribedtonewsletter,
      internalprobabilitytowin,
      source_file_name_new,
      source_file_timestamp_new
    FROM updates_table;
   ```
8. Insert any entirely new, rather than updated, records into the `players` table, as shown in the following code example:
   
   ```
    INSERT INTO players
    SELECT *,
      $SOURCE_FILE_NAME,
      $SOURCE_FILE_TIMESTAMP
    FROM players_ext
    WHERE $SOURCE_FILE_TIMESTAMP > (SELECT max_time FROM control_maxdate)
    AND playerid NOT IN (SELECT playerid FROM players);
   ```
9. Clean up resources. Remove the temporary tables used in the update process as shown in the following code example:
   
   ```
    DROP TABLE IF EXISTS control_maxdate;
    DROP TABLE IF EXISTS updates_table;
   ```

## [](#load-source-file-metadata-into-a-table)Load source file metadata into a table

When you load data from an Amazon S3 bucket, Firebolt uses an external table which holds metadata about your source file to map into a Firebolt database. You can load metadata from the virtual columns contained in the external file into a table. You can use the name, timestamp and file size to determine the source of a row of data in a table. When adding data to an existing table, you can use this information to check whether new data is available, or to determine the vintage of the data. The external table associated with your source file contains the following fields:

- `source_file_name` - the name of your source file.
- `source_file_timestamp` - the date that your source file was modified in the Amazon S3 bucket that it was read from.
- `source_file_size` - the size of your source file in bytes.

The following code example shows you how to create and load metadata into a `levels_meta` table, which contains only the metadata:

```
    CREATE TABLE levels_meta (
      day_of_creation date, 
      name_of_file text, 
      size_of_file int);
    COPY levels_meta(
      day_of_creation $source_file_timestamp, 
      name_of_file $source_file_name, 
      size_of_file $source_file_size) 
    FROM  's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/levels.csv'
    WITH AUTO_CREATE = TRUE
    HEADER = TRUE
    TYPE = CSV;
```

The following code shows you how to read in the source data from the Amazon S3 bucket and add the metadata as new columns in that table:

```
CREATE TABLE IF NOT EXISTS levels_meta_plus (
	"LevelID" INT,
	"Name" TEXT,
	"GameID" INT,
	"LevelType" TEXT,
	"MaxPoints" INT,
	"PointsPerLap" DOUBLE,
	"SceneDetails" TEXT,
	day_of_creation date,
	name_of_file text,
	size_of_file int
);

COPY INTO levels_meta_plus (
  "LevelID",    
   "GameID",
   "Name",
  "LevelType",
  "MaxPoints",
  "PointsPerLap",
  "SceneDetails",
  day_of_creation $source_file_timestamp,  
  name_of_file $source_file_name,      	
  size_of_file $source_file_size
)
FROM  's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/levels.csv'
WITH
HEADER = TRUE
TYPE = CSV;
```

For more information about metadata, see **Using metadata virtual columns** in [Work with external tables](/Guides/loading-data/working-with-external-tables.html).

## [](#continue-loading-even-with-errors)Continue loading even with errors

By default, if Firebolt runs into an error when loading your data, the job will stop loading and end in error. If you want to continue loading your data even in the presence of errors, set `MAX_ERRORS_PER_FILE` to a percentage or integer larger than `0`. `COPY FROM` will then continue to load data until it exceeds the specified percent based on the total number of rows in your data. If you enter an integer between `0` and `100`, `COPY FROM` will interpret the integer as a percentage of rows. You can specify only `0%` or `100%`.

For example, if `MAX_ERRORS_PER_FILE` is set to `0` or `0%`, `COPY FROM` will load data until one row has an error, and then return an error. Setting `MAX_ERRORS_PER_FILE` to either `100` or `100%` allows the loading process to continue even if every row has an error. If all rows have errors, no data will load into the target table.

The following code example loads a sample CSV data set with headers, and will finish the loading job even if every row contains an error.

```
COPY INTO new_levels_auto
FROM 's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/levels.csv'
WITH AUTO_CREATE = TRUE
HEADER = TRUE
TYPE = CSV
MAX_ERRORS_PER_FILE = '100%';
```

In the previous code example, the following apply:

- `COPY INTO new_levels_auto`: Creates a new table named `new_levels_auto`. The `INTO` clause is optional. If the table already exists, `COPY FROM` will add the rows to the existing table.
- `FROM`: Specifies the S3 bucket location of the data. In this example, the dataset is located in a publicly accessible bucket, so you do not need to provide credentials.
- `AUTO_CREATE=TRUE`: Creates a target table and automatically infers the schema.
- `HEADER=TRUE`: Specifies that the first row of the source file contains column headers.
- `TYPE`: Specifies the data format of the incoming data.
- `MAX_ERRORS_PER_FILE`: Specified as an integer or literal text. In the previous example, `MAX_ERRORS_PER_FILE` uses text.

## [](#log-errors-during-data-load)Log errors during data load

`COPY FROM` supports an option to generate error files that describe the errors encountered and note the rows with errors. To store these files in an Amazon S3 bucket, you must provide credentials to allow Firebolt to write to the bucket.

The following example sets an error handling threshold and specifies an Amazon S3 bucket as the source data and another to write the error file:

```
COPY INTO my_table
FROM 's3://my-bucket/data.csv'
WITH
CREDENTIALS = (
  AWS_ROLE_ARN='arn:aws:iam::123456789012:role/my-firebolt-role'
)
MAX_ERRORS_PER_FILE = '100%'
ERROR_FILE = 's3://my-bucket/error_logs/'
ERROR_FILE_CREDENTIALS = (
  AWS_ROLE_ARN='arn:aws:iam::123456789012:role/my-firebolt-role'
)
HEADER = TRUE
```

In the previous code example, the following apply:

- `COPY INTO`: Specifies the target table to load the data into.
- `FROM`: Specifies the S3 bucket location of the data.
- `CREDENTIALS`: Specifies AWS credentials to access information in the Amazon S3 bucket that contains the source data. AWS [AssumeRole](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html) authentication is used for dynamic, temporary credentials. For more Information about credentials and how to set them up, see [The simplest COPY FROM workflow](/Guides/loading-data/loading-data-sql.html#the-simplest-copy-from-workflow).
- Error Handling:
  
  - `MAX_ERRORS_PER_FILE = ‘100%’`: Allows errors in up to `100%` of the rows per file before the load data job fails.
  - `ERROR_FILE`: Specifies the Amazon S3 bucket location to write the error file.
- `HEADER = TRUE`: Indicates that the first row of the CSV file contains column headers.

**How to examine the error files**

If you specify an S3 path with the necessary permissions for an error file and the `COPY FROM` process encounters errors, two different files will be generated in your bucket. The following queries show you how to load these error files into new tables so that you can query and examine the error details and the corresponding rows.

The following query loads the `error_reasons` csv file, which contains a header with column names:

```
COPY error_reasons FROM 's3://my-bucket/error_logs/' 
WITH PATTERN='*error_reasons*.csv' HEADER=TRUE;

SELECT * from error_reasons;
```

The following query loads a file containing all rows that encountered errors. Although this file has no header, the table schema should match that of the source file where the errors occurred.

```
COPY rejected_rows FROM 's3://my-bucket/error_logs/'
WITH PATTERN='*rejected_rows*.csv' HEADER=FALSE;

SELECT * FROM rejected_rows;
```

Configure error handling parameters such as `MAX_ERRORS_PER_FILE`, `ERROR_FILE`, and `ERROR_FILE_CREDENTIALS` to manage how errors are handled, ensure data integrity, and record errors for future review. For more information about `ERROR_FILE` or `ERROR_FILE_CREDENTIALS`, see the **Parameters** section of [COPY FROM](/sql_reference/commands/data-management/copy-from.html).