# [](#list_objects)LIST\_OBJECTS

A table-valued function (TVF) that lists objects and their metadata from Amazon S3. The function can use either a location object (recommended) or direct credentials to access the data. Each row in the output contains four columns: `object_name` (`TEXT`), `object_type` (`TEXT`), `object_bytes` (`BIGINT`), and `file_timestamp` (`TIMESTAMPTZ`).

## [](#syntax)Syntax

```
-- Using location object (recommended)
LIST_OBJECTS (
  LOCATION => location_name
)
|
-- Using static credentials
LIST_OBJECTS (
  URL => <url>
  [, AWS_ACCESS_KEY_ID => <aws_access_key_id>]
  [, AWS_SECRET_ACCESS_KEY => <aws_secret_access_key>]
  [, AWS_SESSION_TOKEN => <aws_session_token>]
  [, AWS_ROLE_ARN => <aws_role_arn>]
  [, AWS_ROLE_EXTERNAL_ID => <aws_role_externakl_id>]
)
```

## [](#parameters)Parameters

Parameter Description Supported input types `LOCATION` The name of a location object that contains the S3 URL and credentials. This is the recommended approach. See [CREATE LOCATION](/sql_reference/commands/data-definition/create-location.html) for details. `IDENTIFIER` `URL` The location of the Amazon S3 bucket containing your files. The expected format is `s3://{bucket_name}/{full_file_path}`. `TEXT` `AWS_ACCESS_KEY_ID` The AWS access key ID. `TEXT` `AWS_SECRET_ACCESS_KEY` The AWS secret access key. `TEXT` `AWS_SESSION_TOKEN` The AWS session token. `TEXT` `AWS_ROLE_ARN` The AWS role ARN. `TEXT` `AWS_ROLE_EXTERNAL_ID` The AWS role external ID. `TEXT`

## [](#return-type)Return Type

The output is a table with four columns:

- `object_name` (`TEXT`)
- `object_type` (`TEXT`)
- `object_bytes` (`BIGINT`)
- `last_modified` (`TIMESTAMPTZ`)

**Column Descriptions**

- **`object_name`** : Contains both the full path and the file extension.
- **`object_type`** : Can be either “file” or “folder”.
  
  - If `object_type` = “folder”, the `object_bytes` and `last_modified` columns will contain `NULL` values, as folders do not have associated sizes or timestamps.
  - If `object_type` = “file”, the following apply:
    
    1. The `last_modified` column is populated from the `LastModified` attribute in Amazon S3. Note that AWS does not expose the creation timestamp, so the values in this column only differ from the creation time if an immutable object has been overwritten.
    2. The `object_bytes` column contains the size of the file in bytes.

Amazon S3 is not a traditional filesystem. In AWS, what is commonly referred to as “folders” is called “common\_prefixes,” and what are typically considered “files” are referred to as “objects.”

## [](#examples)Examples

### [](#using-location-object-to-store-credentials)Using LOCATION object to store credentials

**Best practice**

Firebolt recommends using a `LOCATION` object to store credentials for authentication.

The following code example retrieves all objects from the specified `LOCATION` using the `LIST_OBJECTS` function:

```
SELECT * FROM LIST_OBJECTS(
    LOCATION => my_location
);
```

**Returns**:

object\_name object\_type object\_bytes last\_modified data/file1.csv file 1421277 2023-02-27 10:49:13+01 data/folder1/ folder NULL NULL

### [](#using-static-credentials)Using static credentials

The following code examples show how to list all objects (folders) that start with a specified prefix within an Amazon S3 bucket. You can specify either a part of the prefix or the full prefix. For example, the URLs ending with the prefixes `fire` or `firebolt_sample_dataset` both return identical results because both are valid matches for the `firebolt_sample_dataset` folder as follows:

```
SELECT * FROM LIST_OBJECTS(url => 's3://firebolt-publishing-public/help_center_assets/fire')
```

and

```
SELECT * FROM LIST_OBJECTS(url => 's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset')
```

**Both return**:

object\_name (TEXT) object\_type (TEXT) object\_bytes (BIGINT) last\_modified (TIMESTAMPTZ) help\_center\_assets/firebolt\_sample\_dataset/ folder NULL NULL

The following code example shows how to list all objects such as files, folders, and associated metadata, that start with a specified prefix in an Amazon S3 bucket:

```
SELECT * FROM LIST_OBJECTS(url => 's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/')
```

**Returns**:

object\_name (TEXT) object\_type (TEXT) object\_bytes (BIGINT) last\_modified (TIMESTAMPTZ) help\_center\_assets/firebolt\_sample\_dataset/players.json file 1,421,277 2023-02-27 10:49:13+01 help\_center\_assets/firebolt\_sample\_dataset/levels.csv file 83,596 2023-02-27 11:06:52+01 help\_center\_assets/firebolt\_sample\_dataset/tournaments.csv file 83,351 2022-12-15 15:34:14+01 help\_center\_assets/firebolt\_sample\_dataset/games.json file 872 2023-02-27 13:18:54+01 help\_center\_assets/firebolt\_sample\_dataset/rankings/ folder NULL NULL help\_center\_assets/firebolt\_sample\_dataset/playstats/ folder NULL NULL

The following code examples show how to list all objects such as files and associated metadata, that start with a specified prefix in an Amazon S3 bucket. The URLs ending with the prefixes `lev` or `levels.csv` return identical results because both are valid matches for the `levels.csv` file as follows:

```
SELECT * FROM LIST_OBJECTS(url => 's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/lev')
```

and

```
SELECT * FROM LIST_OBJECTS(url => 's3://firebolt-publishing-public/help_center_assets/firebolt_sample_dataset/levels.csv')
```

**Both Return**:

object\_name (TEXT) object\_type (TEXT) object\_bytes (BIGINT) last\_modified (TIMESTAMPTZ) help\_center\_assets/firebolt\_sample\_dataset/levels.csv file 83,596 2023-02-27 11:06:52+01

The following code example shows how to use your Amazon credentials to list objects in an Amazon S3 bucket that is not publicly accessible:

```
SELECT * FROM LIST_OBJECTS(url => 's3://example_bucket/foo.csv', AWS_ACCESS_KEY_ID => 'my_key_id', AWS_SECRET_ACCESS_KEY => 'my_secret_key')
```

**Returns**:

object\_name (TEXT) object\_type (TEXT) object\_bytes (BIGINT) last\_modified (TIMESTAMPTZ) foo.csv file 10 2024-06-12 9:30:00 UTC