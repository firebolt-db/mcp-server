# [](#load-semi-structured-parquet-data)Load semi-structured Parquet data

Apache Parquet is a binary file format that supports both structured columns and semi-structured data, including arrays, structs, and maps. If these nested structures do not align to a strictly relational schema, they are described as semi-structured. Firebolt’s external tables support extracting these semi-structured fields from Parquet files, treating them similarly as other semi-structured data such as JSON. This document shows how to load and query Parquet data that is stored as structs in arrays or as maps of key-value pairs.

- [Defining external table columns for Parquet arrays and maps](#defining-external-table-columns-for-parquet-arrays-and-maps)
- [Syntax for defining a Parquet nested structure](#syntax-for-defining-a-parquet-nested-structure)
- [Example–ingest and work with structs inside Parquet arrays](#exampleingest-and-work-with-structs-inside-parquet-arrays)
  
  - [Step 1–create an external table](#step-1create-an-external-table)
  - [Step 2–create a fact or dimension table](#step-2create-a-fact-or-dimension-table)
  - [Step 3–insert into the fact table from the external table](#step-3insert-into-the-fact-table-from-the-external-table)
  - [Step 4–query array values](#step-4query-array-values)
- [Example–ingest and work with maps](#exampleingest-and-work-with-maps)
  
  - [Step 1–create an external table](#step-1create-an-external-table-1)
  - [Step 2–create a fact or dimension table](#step-2create-a-fact-or-dimension-table-1)
  - [Step 3–insert into the fact table from the external table](#step-3insert-into-the-fact-table-from-the-external-table-1)
  - [Step 4–query map values](#step-4query-map-values)

## [](#defining-external-table-columns-for-parquet-arrays-and-maps)Defining external table columns for Parquet arrays and maps

When you set up an external table to ingest Parquet data files, you use a hierarchical dotted notation syntax to define table columns. Firebolt uses this notation to identify the field to ingest.

## [](#syntax-for-defining-a-parquet-nested-structure)Syntax for defining a Parquet nested structure

You specify the top grouping element of a nested structure in Parquet followed by the field in that structure that contains the data to ingest. You then declare the column type using the `ARRAY(<data_type>)` notation, where `<data type>` is the [Firebolt data type](/sql_reference/data-types.html) corresponding to the data type of the field in Parquet.

```
"<grouping1>.<datafield>" ARRAY(<data_type>)
```

Examples of this syntax in `CREATE EXTERNAL TABLE` queries are demonstrated below.

## [](#exampleingest-and-work-with-structs-inside-parquet-arrays)Example–ingest and work with structs inside Parquet arrays

Consider the Parquet schema example below. The following elements define an array of structs:

- A single, optional group field, `hashtags`, contains any number of another group, `bag`. This is the top grouping element.
- The `bag` groups each contain a single, optional group, `array_element`.
- The`array_element` group contains a single, optional field, `s`.
- The field `some_value` contains a value that is a `TEXT` type (in binary primitive format).

```
optional group hashtags (LIST) {
  repeated group bag {
    optional group array_element {
      optional binary some_value (TEXT);
    }
  }
}
```

The steps below demonstrate the process to ingest the array values into Firebolt. You create an external table, create a fact table, and insert data into the fact table from the external table, which is connected to the Parquet data store.

### [](#step-1create-an-external-table)Step 1–create an external table

The `CREATE EXTERNAL TABLE` example below creates a column in an external table from the Parquet schema shown in the example above. The column definition uses the top level grouping `hashtags` followed by the field `some_value`. Intermediate nesting levels are omitted.

```
CREATE EXTERNAL TABLE IF NOT EXISTS my_parquet_array_ext_tbl
(
  [...,] --additional columns possible, not shown
  "hashtags.some_value" ARRAY(TEXT)
  [,...]
)
CREDENTIALS = (AWS_KEY_ID = '****' AWS_SECRET_KEY = '*****')
URL = 's3://my_bucket_of_parquet_goodies/'
OBJECT_PATTERN = '*.parquet'
TYPE = (PARQUET);
```

### [](#step-2create-a-fact-or-dimension-table)Step 2–create a fact or dimension table

Create a fact or dimension table that defines a column of the same `ARRAY(TEXT)` type that you defined in the external table in step 1. The example below demonstrates this for a fact table.

```
CREATE FACT TABLE IF NOT EXISTS my_parquet_array_fact_tbl
(
  [...,] --additional columns possible, not shown
  some_value ARRAY(TEXT)
  [,...]
)
[...]
--required primary index for fact table not shown
--optional partitions not shown
;
```

### [](#step-3insert-into-the-fact-table-from-the-external-table)Step 3–insert into the fact table from the external table

The example below demonstrates an `INSERT` statement that selects the array values from Parquet data files using the external table column definition in step 1, and then inserts them into the specified fact table column, `some_value`.

```
INSERT INTO my_parquet_array_fact_tbl
  SELECT "hashtags.some_value" AS some_value
  FROM my_parquet_array_ext_tbl;
```

### [](#step-4query-array-values)Step 4–query array values

After you ingest array values into the fact table, you can query and manipulate the array using array functions and Lambda functions. For more information, see [Working with arrays](/Guides/loading-data/working-with-semi-structured-data/working-with-arrays.html).

Use multipart Parquet column names to extract data from nested structures. For simple `ARRAY(TEXT)`, use a single top-level field name.

## [](#exampleingest-and-work-with-maps)Example–ingest and work with maps

External tables connected to AWS Glue currently do not support reading maps from Parquet files.

Parquet stores maps as arrays of key-value pairs, where each key\_value group contains a key and its corresponding value. Consider the Parquet schema example below. The following define the key-value elements of the map:

- A single, optional group, `context`, is a group of mappings that contains any number of the group `key_value`.
- The `key_value` groups each contain a required field, `key`, which contains the key name as a `TEXT`. Each group also contains an optional field `value`, which contains the value as a `TEXT` corresponding to the key name in the same `key_value` group.

```
optional group context (MAP) {
    repeated group key_value {
      required binary key (TEXT);
      optional binary value (TEXT);
    }
  }
```

The steps below demonstrate the process of creating an external table, creating a fact table, and inserting data into the fact table from the Parquet file using the external table.

### [](#step-1create-an-external-table-1)Step 1–create an external table

When you create an external table for a Parquet map, you use the same syntax that you use in the example for arrays above. You create one column for keys and another column for values. The `CREATE EXTERNAL TABLE` example below demonstrates this.

```
CREATE EXTERNAL TABLE IF NOT EXISTS my_parquet_map_ext_tbl
(
  "context.keys" ARRAY(TEXT),
  "context.values" ARRAY(TEXT)
)
CREDENTIALS = (AWS_KEY_ID = '****' AWS_SECRET_KEY = '*****')
URL = 's3://my_bucket_of_parquet/'
OBJECT_PATTERN = '*.parquet'
TYPE = (PARQUET);
```

### [](#step-2create-a-fact-or-dimension-table-1)Step 2–create a fact or dimension table

Create a Firebolt fact or dimension table that defines columns of the same `ARRAY(TEXT)` types that you defined in the external table in step 1. The example below demonstrates this for a fact table.

```
CREATE FACT TABLE IF NOT EXISTS my_parquet_map_fact_tbl
(
  [...,] --additional columns possible, not shown
  my_parquet_array_keys ARRAY(TEXT),
  my_parquet_array_values ARRAY(TEXT)
  [,...]
)
[...] --required primary index for fact table not shown
      --optional partitions not shown
```

### [](#step-3insert-into-the-fact-table-from-the-external-table-1)Step 3–insert into the fact table from the external table

The example below demonstrates an `INSERT INTO` statement that selects the array values from Parquet data files using the external table column definition in step 1, and inserts them into the specified fact table columns, `my_parquet_array_keys` and `my_parquet_array_values`.

```
INSERT INTO my_parquet_map_fact_tbl
  SELECT "context.keys" AS my_parquet_array_keys,
         "context.values" AS my_parquet_array_values
  FROM my_parquet_map_ext_tbl;
```

### [](#step-4query-map-values)Step 4–query map values

After you ingest array values into the fact table, you can query and manipulate the array using array functions and Lambda functions. For more information, see [Working with arrays](/Guides/loading-data/working-with-semi-structured-data/working-with-arrays.html).

A query that uses a Lambda function to return a specific value by specifying the corresponding key value is shown below. For more information, see [Manipulating arrays using Lambda functions](/Guides/loading-data/working-with-semi-structured-data/working-with-arrays.html#manipulating-arrays-with-lambda-functions).

```
SELECT
  ARRAY_FIRST(v, k -> k = 'key_name_of_interest', my_parquet_array_keys, my_parquet_array_values) AS returned_corresponding_key_value
FROM
  my_parquet_map_ext_tbl;
```