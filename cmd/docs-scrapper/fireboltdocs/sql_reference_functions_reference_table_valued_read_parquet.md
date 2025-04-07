# [](#read_parquet)READ\_PARQUET

A table-valued function (TVF) that reads data from Parquet files stored in Amazon S3. The function can use either a location object (recommended) or direct credentials to access the data. `READ_PARQUET` returns a table with data from the specified Parquet file.

## [](#syntax)Syntax

```
-- Using location object (recommended)
READ_PARQUET (
  LOCATION => location_name
)
|
-- Using static credentials
READ_PARQUET (
  URL => <file_url>
  [, AWS_ACCESS_KEY_ID => <aws_access_key_id>]
  [, AWS_SECRET_ACCESS_KEY => <aws_secret_access_key>]
  [, AWS_SESSION_TOKEN => <aws_session_token>]
  [, AWS_ROLE_ARN => <aws_role_arn>]
  [, AWS_ROLE_EXTERNAL_ID => <aws_role_external_id>]
)
```

## [](#parameters)Parameters

Parameter Description Supported input types `LOCATION` The name of a location object that contains the Amazon S3 URL and credentials. Firebolt recommends using `LOCATION` to store credentials for authentication. See [CREATE LOCATION](/sql_reference/commands/data-definition/create-location.html) for details. `IDENTIFIER` `URL` The location of the Amazon S3 bucket containing your files. The expected format is `s3://{bucket_name}/{full_file_path_glob_pattern}`. `TEXT` `AWS_ACCESS_KEY_ID` The AWS access key ID. `TEXT` `AWS_SECRET_ACCESS_KEY` The AWS secret access key. `TEXT` `AWS_SESSION_TOKEN` The AWS session token. `TEXT` `AWS_ROLE_ARN` The AWS role ARN. `TEXT` `AWS_ROLE_EXTERNAL_ID` The AWS role external ID. `TEXT`

- When using static credentials:
  
  - The `URL` can be passed as either the first positional parameter or a named parameter
  - If you provide either `AWS_ACCESS_KEY_ID` or `AWS_SECRET_ACCESS_KEY`, you must provide both
  - Providing an AWS session token is optional
  - Credentials are not required for accessing public buckets

## [](#return-type)Return Type

The result is a table with data from the Parquet files. Columns are read and parsed using their inferred data types.

## [](#best-practice)Best practice

Firebolt recommends using a `LOCATION` object to store credentials for authentication.

## [](#examples)Examples

**Example**

The following code example reads the first 5 rows from a Parquet file using a `LOCATION` object to store credentials for authentication:

```
SELECT * 
FROM READ_PARQUET(
    LOCATION => my_location
) 
LIMIT 5;
```

**Example**

The following code example reads the first 5 rows from a Parquet file using static credentials for authentication:

```
SELECT * 
FROM READ_PARQUET(
    URL => 's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/playstats/TournamentID=92/cc2a2a0b4e8b4fb39abf20a956e7cc3e-0.parquet'
) 
LIMIT 5;
```

**Returns**

GameID PlayerID Timestamp SelectedCar CurrentLevel CurrentSpeed CurrentPlayTime CurrentScore Event ErrorCode 1 845 2022-10-27 13:36:33 Solara 1 0 0 0 Brake NoError 1 845 2022-10-27 13:36:33 Solara 1 339 0.9872 2 RightTurn GraphicsFreeze 1 845 2022-10-27 13:36:34 Solara 1 288 1.9744 20 Tilt NoError 1 845 2022-10-27 13:36:35 Solara 1 260 2.9616 53 Block TextNotFound 1 845 2022-10-27 13:36:36 Solara 1 196 3.9488 81 FullSpeed NoError

### [](#using-url)Using URL

- The `URL` can be passed as either the first positional parameter or a named parameter. For example, the following two queries will both read the same file:
- Credentials are optional.

```
SELECT * FROM READ_PARQUET('s3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/playstats/TournamentID=92/cc2a2a0b4e8b4fb39abf20a956e7cc3e-0.parquet');
SELECT * FROM READ_PARQUET(URL => 's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/playstats/TournamentID=92/cc2a2a0b4e8b4fb39abf20a956e7cc3e-0.parquet');
```

- The `url` can represent a single file or a [glob](https://en.wikipedia.org/wiki/Glob_%28programming%29) pattern. If a glob pattern is used, all files matching the pattern will be read:

```
SELECT * FROM READ_PARQUET('s3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/playstats/*.parquet')
```