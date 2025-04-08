# [](#read_csv)READ\_CSV

A table-valued function (TVF) that reads CSV files from Amazon S3. The function can use either a location object (recommended) or direct credentials to access the data. `READ_CSV` returns a table with data from the specified CSV file, where each cell is read as `TEXT`.

## [](#syntax)Syntax

```
-- Using LOCATION object (recommended)
READ_CSV ( 
  LOCATION => location_name
  [, COMPRESSION => <file_compression>]
  [, HEADER => { TRUE | FALSE }]
  [, DELIMITER => <field_delimiter>]
  [, QUOTE => { "'" | '"' | SINGLE_QUOTE | DOUBLE_QUOTE}]
  [, NULL_STRING => <null_string>]
  [, ESCAPE => <escape_character>]
  [, SKIP_BLANK_LINES => { TRUE | FALSE }]
  [, EMPTY_FIELD_AS_NULL => { TRUE | FALSE }]
)
|
-- Using static credentials
READ_CSV ( 
  URL => <url>
  [, COMPRESSION => <file_compression>]
  [, AWS_ACCESS_KEY_ID => <aws_access_key_id>]
  [, AWS_SECRET_ACCESS_KEY => <aws_secret_access_key>]
  [, AWS_SESSION_TOKEN => <aws_session_token>]
  [, AWS_ROLE_ARN => <aws_role_arn>]
  [, AWS_ROLE_EXTERNAL_ID => <aws_role_external_id>]
  [, HEADER => { TRUE | FALSE }]
  [, DELIMITER => <field_delimiter>]
  [, QUOTE => { "'" | '"' | SINGLE_QUOTE | DOUBLE_QUOTE}]
  [, NULL_STRING => <null_string>]
  [, ESCAPE => <escape_character>]
  [, SKIP_BLANK_LINES => { TRUE | FALSE }]
  [, EMPTY_FIELD_AS_NULL => { TRUE | FALSE }]
)
```

## [](#parameters)Parameters

Parameter Description Supported input types LOCATION The name of a location object that contains the Amazon S3 URL and credentials. Firebolt recommends using `LOCATION` to store credentials for authentication. See [CREATE LOCATION](/sql_reference/commands/data-definition/create-location.html) for details. `IDENTIFIER` `URL` The location containing your files in an Amazon S3 bucket. The expected format is `s3://{bucket_name}/{full_file_path_glob_pattern}`. `TEXT` `COMPRESSION` The [compression type](/sql_reference/commands/data-definition/create-external-table.html#compression) of the input file. If `compression` is not set, `compression` is inferred from the file extension. `TEXT` `AWS_ACCESS_KEY_ID` The AWS access key ID. `TEXT` `AWS_SECRET_ACCESS_KEY` The AWS secret access key. `TEXT` `AWS_SESSION_TOKEN` The AWS session token. `TEXT` `AWS_ROLE_ARN` The AWS role arn. `TEXT` `AWS_ROLE_EXTERNAL_ID` The AWS role external ID. `TEXT` `HEADER` Set to `TRUE` if the first row of the CSV file contains a header row containing the column names. `BOOL` `DELIMITER` Specify the character used to separate fields. The default delimiter is a comma (`,`). `TEXT` `QUOTE` Specify the character used for quoting fields. The default is double quote (`"`). If a single quote is specified, the quote character will be set to (`'`). Accepts only `DOUBLE_QUOTE`, `SINGLE_QUOTE`, `'`, or `"`. `TEXT` `NULL_STRING` Specify the string used to represent `NULL` values. The default is an empty string, which means that empty strings are interpreted as `NULL` values. `TEXT` `ESCAPE` Specify the character used to escape special characters. The default character is the quote (`'`) character. `TEXT` `SKIP_BLANK_LINES` Set to `TRUE` to ignore blank lines in the file. `BOOL` `EMPTY_FIELD_AS_NULL` Specify whether empty fields should be interpreted as `NULL` values. The default is `TRUE`. If set to `FALSE`, empty fields are interpreted as empty strings. `BOOL`

## [](#return-type)Return Type

The result is a table with the data from the CSV file. Each cell is read as a `TEXT`.

## [](#examples)Examples

### [](#using-location-object)Using LOCATION object

**Best practice**

Firebolt recommends using a `LOCATION` object to store credentials for authentication.

The following code example reads a CSV file from the location specified by `my_location`, treating the first row as a header containing column names:

```
SELECT * FROM READ_CSV(
    LOCATION => my_location,
    HEADER => true
);
```

### [](#using-static-credentials)Using static credentials

```
SELECT * FROM READ_CSV(
    URL => 's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/levels.csv',
    HEADER => true
);
```

**Example**

In the following example, the `URL` is set as the first positional parameter and reads a CSV file:

```
SELECT * FROM READ_CSV('s3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/levels.csv');
```

**Returns**:

f0 f1 f2 f3 f4 f5 f6 f7 f8 … LevelID GameID Level Name LevelType NextLevel MinPointsToPass MaxPoints NumberOfLaps … 1 1 1 Thunderbolt Circuit FastestLap 2 5 20 5 … 2 1 2 Velocity Vale FirstToComplete 3 15 30 10 … 3 1 3 Raceway Ridge FastestLap 4 25 40 20 … 4 1 4 Nitro Narrows FirstToComplete 5 60 100 10 … 5 1 5 Thunder Road FirstToComplete 6 80 150 15 … 6 1 6 Burnout Boulevard Drift 7 50 80 8 … 7 1 7 Speed Street FastestLap 8 40 70 7 … 8 1 8 Racing Ravine FastestLap 9 60 100 20 … 9 1 9 Drift District Drift 10 100 250 25 … 10 1 10 Acceleration Alley FirstToComplete null 200 500 50 …

**Example**

The following example accepts `URL` as a named parameter and reads a CSV file with column names in the first row:

```
SELECT * FROM READ_CSV(URL => 's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/levels.csv', 
        HEADER => true);
```

**Returns**

LevelID GameID Level Name LevelType NextLevel MinPointsToPass MaxPoints NumberOfLaps … 1 1 Thunderbolt Circuit FastestLap 2 5 20 5 20 … 2 1 Velocity Vale FirstToComplete 3 15 30 10 10 … 3 1 Raceway Ridge FastestLap 4 25 40 20 20 … 4 1 Nitro Narrows FirstToComplete 5 60 100 10 10 … 5 1 Thunder Road FirstToComplete 6 80 150 15 15 … 6 1 Burnout Boulevard Drift 7 50 80 8 8 … 7 1 Speed Street FastestLap 8 40 70 7 7 … 8 1 Racing Ravine FastestLap 9 60 100 20 20 … 9 1 Drift District Drift 10 100 250 25 25 … 10 1 Acceleration Alley FirstToComplete null 200 500 50 50 …

**Example**

The following example reads a CSV with headers and reads empty values as empty strings, rather than `NULL` values:

```
SELECT * FROM READ_CSV(URL => 's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/levels.csv',
        HEADER => true, EMPTY_FIELD_AS_NULL => false);
```

**Returns**:

LevelID GameID Level Name LevelType … 1 1 Thunderbolt Circuit FastestLap 2 … 2 1 Velocity Vale FirstToComplete 3 … 3 1 Raceway Ridge FastestLap 4 … 4 1 Nitro Narrows FirstToComplete 5 … 5 1 Thunder Road FirstToComplete 6 … 6 1 Burnout Boulevard Drift 7 … 7 1 Speed Street FastestLap 8 … 8 1 Racing Ravine FastestLap 9 … 9 1 Drift District Drift 10 … 10 1 Acceleration Alley FirstToComplete   …