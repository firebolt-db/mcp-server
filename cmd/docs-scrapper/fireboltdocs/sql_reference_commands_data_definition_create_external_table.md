# [](#create-external-table)CREATE EXTERNAL TABLE

Creates an external table. External tables serve as connectors to your external data sources. External tables contain no data within Firebolt other than metadata virtual columns that are automatically populated with metadata. For more information, see [Working with external tables](/Guides/loading-data/working-with-external-tables.html). Data that you ingest must be in an Amazon S3 bucket in the same AWS Region as the Firebolt database.

- [Using location objects](#using-location-objects)
- [Common `LOCATION` errors](#common-location-errors)
- [PARTITION](#partition)
- [Credentials](#credentials)
  
  - [Syntax–authenticating using an access key and secret](#syntaxauthenticating-using-an-access-key-and-secret)
- [OBJECT\_PATTERN](#object_pattern)
- [TYPE](#type)
  
  - [CSV Types](#csv-types)
  - [JSON Types](#json-types)
  - [Other Types](#other-types)
- [COMPRESSION](#compression)

## [](#syntax)Syntax

```
-- Using location object (recommended)
CREATE EXTERNAL TABLE [IF NOT EXISTS] <table>
(
    <column_name> <column_type>[ PARTITION('<regex>')]
    [, <column_name2> <column_type2> [PARTITION('<regex>')]]
    [,...<column_name2> <column_type2> [PARTITION('<regex>')]]
)
LOCATION = <location_name>
OBJECT_PATTERN = '<object_pattern>'
TYPE = ( <type> )
[ <type_option> ]
[ COMPRESSION = <compression_type> ]
|
-- Using static credentials
CREATE EXTERNAL TABLE [IF NOT EXISTS] <table>
(
    <column_name> <column_type>[ PARTITION('<regex>')]
    [, <column_name2> <column_type2> [PARTITION('<regex>')]]
    [,...<column_name2> <column_type2> [PARTITION('<regex>')]]
)
[ CREDENTIALS = { AWS_ACCESS_KEY_ID = '<aws_access_key_id>' AWS_SECRET_ACCESS_KEY = '<aws_secret_access_key>' [ AWS_SESSION_TOKEN = '<aws_session_token>' ] | AWS_ROLE_ARN = '<aws_role_arn>' [ AWS_ROLE_EXTERNAL_ID = '<aws_role_external_id>' ] } ]
URL = 's3://<bucket_name>[/<folder>][/...]/'
OBJECT_PATTERN = '<object_pattern>'
TYPE = ( <type> )
[ <type_option> ]
[ COMPRESSION = <compression_type> ]
```

## [](#parameters)Parameters

Parameter Description `<table>` An identifier that specifies the name of the external table. This name should be unique within the database. For identifier usage and syntax, see [Object identifiers](/Reference/object-identifiers.html). `<column_name>` An identifier that specifies the name of the column. This name should be unique within the table.  
**Note:** If column names are using mixed case, wrap your column name definitions in double quotes (`"`); otherwise they will be translated to lower case and will not match the mixed case Parquet schema. `<column_type>` Specifies the data type for the column. `PARTITION` An optional keyword. When specified, allows you to use a regular expression `<regex>` to extract a value from the file prefix to be stored as the column value. For more information, see [PARTITION](#partition). `LOCATION` The name of a location object that contains the Amazon S3 URL and credentials. This is the recommended approach for specifying the source. See ([CREATE LOCATION](/sql_reference/commands/data-definition/create-location.html)) for details. The location must exist and you must have appropriate permissions to use it. `CREDENTIALS` Specifies the AWS credentials with permission to access the Amazon S3 location specified using `URL`. For more information, see [CREDENTIALS](#credentials). `URL` The path to an Amazon S3 URL where the source files are located. For example, `s3://my_bucket/my_folder/`. `OBJECT_PATTERN` Specifies the file naming pattern that Firebolt ingests when using this table. For more information, see [OBJECT\_PATTERN](#object_pattern). `TYPE` Specifies the file type Firebolt expects to ingest given the `OBJECT_PATTERN`. If a file referenced using `OBJECT_PATTERN` does not conform to the specified `TYPE`, an error occurs. For more information, see [TYPE](#type). `<type_option>` Allows configuration for ingesting different CSV file formats. Type option can be set at this top level, or as an option in the `TYPE` parameter. `COMPRESSION` See [COMPRESSION](#compression).

### [](#using-location-objects)Using location objects

Firebolt recommends using location objects as a secure, centralized way to manage Amazon S3 credentials and URLs.

Use location objects for:

- Centralized credential management.
- Eliminating the exposure of credentials in queries.
- Enabling role-based access control.
- Simplified maintenance and updates.

**Example**

The following code example creates `my_external_table` in that loads data from Parquet files matching the \*.parquet pattern in `my_data_location`:

```
CREATE EXTERNAL TABLE my_external_table
(
    c_id INTEGER,
    c_name TEXT
)
LOCATION = my_data_location
OBJECT_PATTERN = '*.parquet'
TYPE = (PARQUET);
```

For detailed information, see [CREATE LOCATION](/sql_reference/commands/data-definition/create-location.html).

### [](#common-location-errors)Common `LOCATION` errors

- Providing the location name as a string literal instead of an identifier results in an error, as shown in the following code example:
  
  ```
  CREATE EXTERNAL TABLE et_1 (x INT NOT NULL, y TEXT NOT NULL) 
  LOCATION = 'location_1'  -- Error: Parameter LOCATION must be of type IDENTIFIER
  OBJECT_PATTERN = 'simple.csv' 
  TYPE = CSV;
  ```
- If the specified location doesn’t exist or you lack permission to use it, an error occurs, as shown in the following code example:
  
  ```
  CREATE EXTERNAL TABLE et_2 (x INT NOT NULL, y TEXT NOT NULL) 
  LOCATION = nonexistent_location  -- Error: location 'nonexistent_location' does not exist or not authorized
  OBJECT_PATTERN = 'simple.csv' 
  TYPE = CSV;
  ```

### [](#partition)PARTITION

In some applications, such as Hive partitioning, files use a folder naming convention to identify which partition their data belongs to. The `PARTITION` keyword allows you to specify a regular expression, `<regex>`, to extract a portion of the file path and store it in the specified column when Firebolt uses the external table to ingest partitioned data.

Using `PARTITION` in this way is one method of extracting partition data from file paths. Another method is to use the table metadata column, `$source_file_name`, during the `INSERT` operation.

#### [](#guidelines-for-creating-the-regex)Guidelines for creating the regex

- The regular expression is matched against the object prefix, not including the `s3://<bucket_name>/` portion of the prefix.
- You can specify as many `PARTITION` columns as you need, each extracting a different portion of the object prefix.
- For each `PARTITION` column, you must specify a [re2 regular expression](https://github.com/google/re2/wiki/Syntax) that contains a capturing group, which determines the column value.
- When `<column_type>` is `DATE`, Firebolt requires up to three capturing groups that must be in the order of year, month, and day.
- When `<column_type>` is `TIMESTAMP`, Firebolt requires up to six capturing groups that must be in the order of year, month, day, hour, minute, second.
- Firebolt tries to convert the captured string to the specified `<column_type>`. If the type conversion fails, the ingest will error out.

In most cases, the easiest way to build a regular expression is as follows:

1. Count the number of folders in the path, not including the bucket name.
2. Concatenate the string `[^/]+/` according to the number of folders.
3. Prefix the regex with an additional `[^/]+` for the file name.
4. Wrap the `[^/]+` in the right folder with a capturing group parenthesis, such as `([^/]+).`

For more information, see [Match groups](https://regexone.com/lesson/capturing_groups) on the RegexOne website. To test your regular expressions, online tools such as [regex101](https://regex101.com) are available.

#### [](#exampleextract-hive-compatible-partitions)Example–extract Hive-compatible partitions

The example below demonstrates a `CREATE EXTERNAL TABLE` statement that creates the table `my_ext_table`. This table is used to ingest all files with a `*.parquet` file extension in any sub-folder of the Amazon S3 bucket `s3://my_bucket`.

Consider an example where folders and files in the bucket have the following consistent pattern, which is common for Hive partitions:

```
s3://my_bucket/c_type=xyz/year=2018/month=01/part-00001.parquet
s3://my_bucket/c_type=xyz/year=2018/month=01/part-00002.parquet
s3://my_bucket/c_type=abc/year=2018/month=01/part-00001.parquet
s3://my_bucket/c_type=abc/year=2018/month=01/part-00002.parquet
[...]
```

In the example `CREATE EXTERNAL TABLE` statement below, the `PARTITION` keyword in the column definition for `c_type` specifies a regular expression. This expression extracts the portion of the Amazon S3 path name that correspond to the `xyz` or `abc` within `c_type=xyz` or `c_type=abc`.

```
CREATE EXTERNAL TABLE my_ext_table (
  c_id    INTEGER,
  c_name  TEXT,
  c_type  TEXT PARTITION('[^/]+/c_type=([^/]+)/[^/]+/[^/]+')
)
CREDENTIALS = (AWS_ACCESS_KEY_ID = 'AKIAIOSFODNN7EXAMPLE' AWS_SECRET_ACCESS_KEY = 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY')
URL = 's3://my_bucket/'
OBJECT_PATTERN= '*.parquet'
TYPE = (PARQUET)
```

When Firebolt ingests the data from a Parquet file stored in that path, the `c_type` column for each row contains the extracted portion of the path. For the files listed above, the extraction results in the following values. `c_id` and `c_name` are values stored within the respective Parquet files, while `c_type` are values extracted from the file path.

c\_id c\_name c\_type 1ef4302294 Njimba xyz 8b98470659 Yuang xyz 98734hkk89 Cole xyz 38cjodjlo8 Blanda xyz 448dfgkl12 Harris abc j987rr3233 Espinoza abc

### [](#credentials)Credentials

The credentials for accessing your data on AWS Amazon S3 using access key &amp; secret.

#### [](#syntaxauthenticating-using-an-access-key-and-secret)Syntax–authenticating using an access key and secret

```
CREDENTIALS = (AWS_ACCESS_KEY_ID = '<aws_access_key_id>' AWS_SECRET_ACCESS_KEY = '<aws_secret_access_key>' [ AWS_SESSION_TOKEN = '<aws_session_token>' ] )
```

## [](#parameters-1)Parameters

Parameters Description Data type `AWS_ACCESS_KEY_ID` The AWS access key ID. `TEXT` `AWS_SECRET_ACCESS_KEY` The AWS secret access key. `TEXT` `AWS_SESSION_TOKEN` The AWS session token. `TEXT` `AWS_ROLE_ARN` The AWS role ARN (Amazon Resource Name). `TEXT` `AWS_ROLE_EXTERNAL_ID` The AWS role external id. `TEXT`

In case you don’t have the access key and secret to access your Amazon S3 bucket, read more [here](https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys) on how to obtain them.

### [](#object_pattern)OBJECT\_PATTERN

An external table enables reading some (or all) files from an Amazon S3 bucket that you have read access to. The Amazon S3 bucket that you reference must be in the same AWS Region as the Firebolt database.

The`OBJECT_PATTERN` parameter identifies which files represent the data for the external table.`OBJECT_PATTERN` is a glob that selects files within the `URL` or `LOCATION`.

#### [](#syntax-1)Syntax

```
OBJECT_PATTERN = '<object_pattern>'
```

## [](#parameters-2)Parameters

Parameters Description Data type `OBJECT_PATTERN` Specify the data pattern to be found in your data source. For example, \*.parquet indicates that all parquet files should be found. `TEXT`

The following wildcards are supported:

- `'*'` matches any sequence of characters
- `'?'` matches any single character
- `[SET]` matches any single character in the specified set
- `[!SET]` matches any character, not in the specified set.

#### [](#example)Example

In the following layout of objects in a bucket, the data is partitioned according to client type, year, and month, with multiple parquet files in each partition. The examples demonstrate how choosing both URL and OBJECT\_PATTERN impacts the objects that are retrieved from Amazon S3.

```
s3://bucket/c_type=xyz/year=2018/month=01/part-00001.parquet
s3://bucket/c_type=xyz/year=2018/month=01/part-00002.parquet
...
s3://bucket/c_type=xyz/year=2018/month=12/part-00001.parquet
s3://bucket/c_type=xyz/year=2018/month=12/part-00002.parquet
...
s3://bucket/c_type=xyz/year=2019/month=01/part-00001.parquet
s3://bucket/c_type=xyz/year=2019/month=01/part-00002.parquet
...
s3://bucket/c_type=xyz/year=2020/month=01/part-00001.parquet
s3://bucket/c_type=xyz/year=2020/month=01/part-00002.parquet
...
s3://bucket/c_type=abc/year=2018/month=01/part-00001.parquet
s3://bucket/c_type=abc/year=2018/month=01/part-00002.parquet
...
```

Following are some common use cases for URL and object pattern combinations:

Use cases Syntax Get all files for file type xyz *URL = ‘s3://bucket/c\_type=xyz/’*  
*OBJECT\_PATTERN = ‘\*’*

*URL = ‘s3://bucket/’*  
*OBJECT\_PATTERN = ‘c\_type=xyz/\*’* Get one specific file: `c_type=xyz/year=2018/month=01/part-00001.parquet` \*URL = ‘s3://bucket/c\_type=xyz/year=2018/month=01/’  
OBJECT\_PATTERN = ‘part-00001.parquet’ Get all parquet files for type xyz *URL = ‘s3://bucket/c\_type=xyz/’*  
*OBJECT\_PATTERN = ‘\*.parquet’*

### [](#type)TYPE

Specifies the type of the files in Amazon S3. The following types and type options are supported.

#### [](#csv-types)CSV Types

```
TYPE = (CSV [ <type_option> ])
```

or

```
TYPE = (CSV)
[ <type_option> ]
```

The following type options allow configuration for ingesting different CSV file formats.

- `[ALLOW_DOUBLE_QUOTES = {TRUE|FALSE}]`  
  `[ALLOW_SINGLE_QUOTES = {TRUE|FALSE}]`  
  With `ALLOW_DOUBLE_QUOTES = TRUE` or `ALLOW_SINGLE_QUOTES = TRUE` you define that unescaped double or single quotes in CSV input file will not cause an error to be generated on ingest. By default `ALLOW_DOUBLE_QUOTES` and `ALLOW_SINGLE_QUOTES` are set to `TRUE`.
- `[ALLOW_COLUMN_MISMATCH = {TRUE|FALSE}]`  
  With `ALLOW_COLUMN_MISMATCH = TRUE` the number of delimited columns in a CSV input file can be fewer than the number of columns in the corresponding table. By default, `ALLOW_COLUMN_MISMATCH` is set to `FALSE`, and an error is generated if the number of columns is fewer than the number of columns defined in the external table. If set to `TRUE`, and an input file record contains fewer columns than defined in the external table, the non-matching columns in the table are loaded with `NULL` values.
- `[ALLOW_UNKNOWN_FIELDS = {TRUE|FALSE}]`  
  With `ALLOW_UNKNOWN_FIELDS = TRUE` the number of delimited columns in a CSV input file can be more than the number of columns in the corresponding table. By default, `ALLOW_UNKNOWN_FIELDS` is set to `FALSE`, and an error is generated if the number of columns is more than the number of columns defined in the external table. If set to `TRUE`, and an input file record contains more columns than defined in the external table, the non-matching columns in the table are ignored.
- `[ESCAPE_CHARACTER = {‘<character>’|NONE}`  
  With `ESCAPE_CHARACTER = '<character>'` you can define which character is used to escape, to change interpretations from the original. By default, the `ESCAPE_CHARACTER` value is set to `\`. If, for example, you want to use `"` as a value and not as delimiter for string, you can escape like `\"`, with the default escape character.
- `[FIELD_DELIMITER = '<field_delimeter>']`  
  With `FIELD_DELIMITER = '<field_delimeter>'`, you can define a custom field delimiter to separate fields for ingest. By default, the `FIELD_DELIMITER` is set as `,`.
- `[NEW_LINE_CHARACTER = '<new_line_character>']`  
  With `NEW_LINE_CHARACTER = '<new_line_character>'`, you can define a custom new line delimiter to separate entries for ingest. By default, the `NEW_LINE_CHARACTER` is set as the end of line character `\n`, but also supports other end of line conventions, such as `\r\n`, `\n\r`, and `\r`, as well as multi-character delimiters, such as `#*~`.
- `[NULL_STRING = '<null_string>']`  
  With `NULL_STRING = '<null_string>'` you can define which set of characters is interpreted as `NULL`. By default, the `NULL_STRING` value is set to `\\N`.
- `[SKIP_BLANK_LINES {TRUE|FALSE}]`  
  With `SKIP_BLANK_LINES = TRUE` any blank lines encountered in the CSV input file will be skipped. By default, `SKIP_BLANK_LINES` is set to `FALSE`, and an error is generated if blank lines are enountered on ingest.
- `[SKIP_HEADER_ROWS = {TRUE|FALSE}]`  
  With `SKIP_HEADER_ROWS = TRUE`, Firebolt assumes that the first row in each file read from Amazon S3 is a header row and skips it when ingesting data. When set to `FALSE`, which is the default if not specified, Firebolt ingests the first row as data.

#### [](#json-types)JSON Types

- `TYPE = (JSON [PARSE_AS_TEXT = {TRUE|FALSE}])`  
  With `TYPE = (JSON PARSE_AS_TEXT = TRUE)`, Firebolt ingests each JSON object literal in its entirety into a single column of type `TEXT`. With `TYPE = (JSON PARSE_AS_TEXT = FALSE)`, Firebolt expects each key in a JSON object literal to map to a column in the table definition. During ingestion, Firebolt inserts the key’s value into the corresponding column.

#### [](#other-types)Other Types

- `TYPE = (ORC)`
- `TYPE = (PARQUET)`
- `TYPE = (AVRO)`
- `TYPE = (TSV)`

All type options for CSV above, except for `FIELD_DELIMITER`, are also supported for the TSV file type.

#### [](#example-1)Example

Creating an external table that reads parquet files from Amazon S3 is being done with the following statement:

```
CREATE EXTERNAL TABLE my_external_table
(
    c_id INTEGER,
    c_name TEXT
)
CREDENTIALS = (AWS_ACCESS_KEY_ID = '****' AWS_SECRET_ACCESS_KEY = '****')
URL = 's3://bucket/'
OBJECT_PATTERN= '*.parquet'
TYPE = (PARQUET)
```

### [](#compression)COMPRESSION

Specifies the compression type of the files matching the specified `OBJECT_PATTERN` in Amazon S3.

#### [](#syntax-2)Syntax

```
[COMPRESSION = <compression_type>]
```

#### [](#parameters-3)Parameters

Parameters Description `<compression_type>` Specifies the compression type of files. `GZIP` is supported.

#### [](#example-2)Example

The example below creates an external table to ingest CSV files from Amazon S3 that are compressed using gzip. The credentials for an IAM user with access to the bucket are provided.

```
CREATE EXTERNAL TABLE my_external_table
(
    c_id INTEGER,
    c_name TEXT
)
CREDENTIALS = (AWS_ACCESS_KEY_ID = 'AKIAIOSFODNN7EXAMPLE' AWS_SECRET_ACCESS_KEY = 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY')
URL = 's3://mybucket/'
OBJECT_PATTERN= '*.csv.gz'
TYPE = (CSV)
COMPRESSION = GZIP
```