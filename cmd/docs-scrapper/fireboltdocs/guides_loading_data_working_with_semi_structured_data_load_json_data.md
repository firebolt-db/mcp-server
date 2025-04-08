# [](#load-semi-structured-json-data)Load semi-structured JSON data

Semi-structured data does not follow a strict table format but contains structured tags or key-value pairs. JSON is an example of semi-structured data. Firebolt supports the following three ways to ingest JSON based on how your data changes and how you query it:

- [Load JSON into a fixed schema](#load-json-into-a-fixed-schema) if your JSON data has a stable set of fields with shallow nesting.
- [Transform the input during load](#transform-the-input-during-load) if your table must always contain certain fields.
- [Store JSON as text](#store-json-as-text) if you need only specific fields on demand or if the table structure changes frequently.

This document shows you how to load data using each of the previous methods and the sample JSON dataset in the following section.

### [](#sample-json-dataset)Sample JSON dataset

The following JSON data shows two session records for a website, where each line represents a single JSON object. This sample data is used in each of the examples in this document.

```
[
  {
    "id": 1,
    "StartTime": "2020-01-06 17:00:00",
    "Duration": 450,
    "tags": ["summer-sale", "sports"],
    "user_agent": {
      "agent": "Mozilla/5.0",
      "platform": "Windows NT 6.1",
      "resolution": "1024x4069"
    }
  },
  {
    "id": 2,
    "StartTime": "2020-01-05 12:00:00",
    "Duration": 959,
    "tags": ["gadgets", "audio"],
    "user_agent": {
      "agent": "Safari",
      "platform": "iOS 14"
    }
  }
]
```

The following code example creates a staging table that stores the raw JSON data, allowing you to run the subsequent examples:

```
-- Create a staging table for raw JSON data with one JSON object per row
DROP TABLE IF EXISTS doc_visits_source;
CREATE TABLE doc_visits_source (
  raw_json TEXT
);

-- Insert raw JSON data as individual rows
INSERT INTO doc_visits_source (raw_json)
VALUES
('{"id": 1, "StartTime": "2020-01-06 17:00:00", "Duration": 450, "tags": ["summer-sale", "sports"], "user_agent": {"agent": "Mozilla/5.0", "platform": "Windows NT 6.1", "resolution": "1024x4069"}}'),
('{"id": 2, "StartTime": "2020-01-05 12:00:00", "Duration": 959, "tags": ["gadgets", "audio"], "user_agent": {"agent": "Safari", "platform": "iOS 14"}}');
```

If you want to load JSON data from an Amazon S3 bucket, you can create an external table that references the file as follows:

```
CREATE EXTERNAL TABLE visits_external (
  raw_json TEXT
)
LOCATION = 's3://your-bucket-name/path/to/json-file/'
FILE_FORMAT = (TYPE = 'JSON');
```

## [](#load-json-into-a-fixed-schema)Load JSON into a fixed schema

If your JSON data has a stable set of fields with shallow nesting, you can load it into a table with a fixed schema to simplify queries. Missing keys are assigned default values. This method allows you to query columns directly without additional parsing, making queries faster and easier to write. Extra keys that are not explicitly mapped are excluded from structured tables, making this approach less flexible for changing data. If stored separately in a `TEXT` column, they remain accessible for later extraction.

The following code example uses the previously created `doc_visits_source` table to define columns that map directly to known keys:

```
-- Create the target table 'visits_fixed' with a fixed schema
DROP TABLE IF EXISTS visits_fixed;
CREATE FACT TABLE visits_fixed (
  id INT DEFAULT 0,
  start_time TIMESTAMP DEFAULT '1970-01-01 00:00:00',
  duration INT DEFAULT 0,
  tags ARRAY(TEXT) DEFAULT []
)
PRIMARY INDEX start_time;

-- Insert data into 'visits_fixed' by extracting values from the raw JSON
INSERT INTO visits_fixed
SELECT
  JSON_POINTER_EXTRACT(raw_json, '/id')::INT AS id,
  TO_TIMESTAMP(TRIM(BOTH '"' FROM JSON_POINTER_EXTRACT(raw_json, '/StartTime')), 'YYYY-MM-DD HH24:MI:SS') AS start_time,
  JSON_POINTER_EXTRACT(raw_json, '/Duration')::INT AS duration,
  JSON_POINTER_EXTRACT(raw_json, '/tags')::ARRAY(TEXT) AS tags
FROM doc_visits_source;
```

The following table shows the expected results:

id start\_time duration tags 1 1/6/2020 17:00 450 \[“summer-sale”, “sports”] 2 1/5/2020 12:00 959 \[“gadgets”, “audio”]

Important characteristics of the table:

- The mandatory scalar fields, `id`, `start_time`, and `duration`, are stored in separate columns, which makes it easier to filter, sort, or join by these fields.
- Each column maps directly to a known JSON key, allowing for simpler queries without the need for JSON functions.
- Default values ensure that the table loads even if some fields are missing or additional keys appear. Extra JSON fields such as `user_agent` with `agent`, `platform`, and `resolution` are ignored and not stored in the table.
- Array columns are used to store `tags`, which supports arbitrary numbers of values without schema changes.

## [](#transform-the-input-during-load)Transform the input during load

Parsing JSON data during ingestion eliminates the need for subsequent query-time parsing, simplifying and accelerating queries. However, transforming data during load also requires well-defined JSON paths that remain consistent. If the JSON paths change, the load might fail.

The following code example uses the previously created `doc_visits_source` table to parse JSON data as it loads and inserts extracted fields into a Firebolt table named `visits_transformed`. It shows how to use [JSON\_POINTER\_EXTRACT\_KEYS](/sql_reference/functions-reference/JSON/json-pointer-extract-keys.html) and [JSON\_POINTER\_EXTRACT\_VALUES](/sql_reference/functions-reference/JSON/json-pointer-extract-values.html) to store a dynamic key-value pair – `agent_props_keys` and `agent_props_vals` – from a nested object:

```
DROP TABLE IF EXISTS visits_transformed;
CREATE FACT TABLE visits_transformed (
  id INT,
  start_time TIMESTAMP,
  duration INT,
  tags ARRAY(TEXT),
  agent_props_keys ARRAY(TEXT),
  agent_props_vals ARRAY(TEXT)
)
PRIMARY INDEX start_time;

INSERT INTO visits_transformed
SELECT
  JSON_POINTER_EXTRACT(raw_json, '/id')::INT,
  TO_TIMESTAMP(TRIM(BOTH '"' FROM JSON_POINTER_EXTRACT(raw_json, '/StartTime')), 'YYYY-MM-DD HH24:MI:SS'),
  JSON_POINTER_EXTRACT(raw_json, '/Duration')::INT,
  JSON_POINTER_EXTRACT(raw_json, '/tags')::ARRAY(TEXT),
  JSON_POINTER_EXTRACT_KEYS(raw_json, '/user_agent')::ARRAY(TEXT),
  JSON_POINTER_EXTRACT_VALUES(raw_json, '/user_agent')::ARRAY(TEXT)
FROM doc_visits_source;
```

The following table shows the expected results:

id start\_time duration tags agent\_props\_keys agent\_props\_vals 1 1/6/2020 17:00 450 \[“summer-sale”,”sports”] \[“agent”, “platform”, “resolution”] \[“Mozilla/5.0”, “Windows NT 6.1”, “1024x4069”] 2 1/5/2020 12:00 959 \[“gadgets”,”audio”] \[“agent”, “platform”] \[“Safari”, “iOS 14”]

Important characteristics of the previous table:

- The `user_agent` object is stored in two arrays: `agent_props_keys` and `agent_props_vals`. The [`JSON_POINTER_EXTRACT_KEYS`](/sql_reference/functions-reference/JSON/json-pointer-extract-keys.html) function extracts the keys from the `user_agent` object into the `agent_props_keys` array. The [`JSON_POINTER_EXTRACT_VALUES`](/sql_reference/functions-reference/JSON/json-pointer-extract-values.html) function extracts the corresponding values into the `agent_props_vals` array. Storing keys and values in parallel arrays offers flexibility when the `user_agent` map changes and avoids schema updates for new or removed fields.

A common error may occur if a field path does not exist in the JSON document. Firebolt returns an error because `NULL` values cannot be cast to `INT`. For example, the following query attempts to extract a non-existent field `/unknown_field` and cast it to `INT`, which results in an error:

```
SELECT JSON_POINTER_EXTRACT(raw_json, '/unknown_field')::INT 
FROM doc_visits_source;
```

To avoid this error, use a default value or conditional expression as shown in the following code example:

```
INSERT INTO visits_transformed
SELECT
  CASE 
    WHEN JSON_POINTER_EXTRACT(raw_json, '/unknown_field') IS NOT NULL 
    THEN JSON_POINTER_EXTRACT(raw_json, '/unknown_field')::INT 
    ELSE NULL
  END AS id
FROM doc_visits_source;
```

The following table shows the expected results:

id start\_time duration tags 0 NULL NULL NULL 0 NULL NULL NULL 1 1/6/2020 17:00 450 \[“summer-sale”, “sports”] 2 1/5/2020 12:00 959 \[“gadgets”, “audio”]

## [](#store-json-as-text)Store JSON as text

You can store JSON as a single text column if the data structure changes frequently or if you only need certain fields in some queries. This approach simplifies ingestion since no parsing occurs during loading, but it requires parsing fields at query time, which can make queries more complex if you need to extract many fields regularly.

The following code example uses the previously created intermediary `doc_visits_source` table to create a permanent table that stores raw JSON, allowing you to parse only what you need on demand:

```
DROP TABLE IF EXISTS visits_raw;
CREATE FACT TABLE visits_raw (
  raw_json TEXT
);

-- Insert data into the 'visits_raw' table from the staging table
INSERT INTO visits_raw
SELECT raw_json
FROM doc_visits_source;
```

The following table shows the expected results:

raw\_json {“id”: 1, “StartTime”: “2020-01-06 17:00:00”, “Duration”: 450, “tags”: \[“summer-sale”, “sports”], “user\_agent”: {“agent”: “Mozilla/5.0”, “platform”: “Windows NT 6.1”, “resolution”: “1024x4069”}} {“id”: 2, “StartTime”: “2020-01-05 12:00:00”, “Duration”: 959, “tags”: \[“gadgets”, “audio”], “user\_agent”: {“agent”: “Safari”, “platform”: “iOS 14”}}

Important characteristics of the table:

- The `id`, `start_time`, `durations`, and `tags` columns follow the same purpose as in the [previous table example](#transform-the-input-during-load).
- Each row in the previous table contains a complete JSON object stored in a single `TEXT` column, rather than being parsed into separate fields. This approach is beneficial when the required fields are unknown at ingestion or the JSON structure changes frequently, allowing for flexible data storage without modifying the schema. Fields can be extracted dynamically at query time using Firebolt’s JSON functions, though frequent parsing may increase query complexity and cost.
- Parsing occurs at query time, which can save upfront processing when data is loaded, but it might increase query complexity and cost if you need to parse many fields frequently.
- Subsequent queries need to extract fields manually with JSON functions as needed.