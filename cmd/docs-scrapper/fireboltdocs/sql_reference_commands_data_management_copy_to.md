# [](#copy-to)COPY TO

Copies exports or unloads the results of a `SELECT` query to an Amazon S3 source in the file format that you specify. You can specify the output file format, compression type, and other options to control how the data is written, such as file name conventions and whether to overwrite existing files. Use `COPY TO` for backing up query results or integrating Firebolt with other data processing tools.

- [Syntax](#syntax)
- [Parameters](#parameters)
- [Credentials](#credentials)
  
  - [Credentials for Amazon S3](#credentials-for-amazon-s3)
  - [Least privileged permissions](#least-privileged-permissions)
- [Using Location Objects](#using-location-objects)
  
  - [Example with Location Object](#example-with-location-object)
- [Examples](#examples)
  
  - [COPY TO with location object](#copy-to-with-location-object)
  - [COPY TO with defaults and role ARN](#copy-to-with-defaults-and-role-arn)
  - [COPY TO with single file and AWS access keys](#copy-to-with-single-file-and-aws-access-keys)
  - [COPY TO with custom file name prefix and ARN with external ID](#copy-to-with-custom-file-name-prefix-and-arn-with-external-id)
  - [COPY TO with custom file name and overwrite set to TRUE](#copy-to-with-custom-file-name-and-overwrite-set-to-true)

## [](#syntax)Syntax

```
COPY (<select_query>)
  TO <location_name>
[ TYPE = CSV | TSV | JSON | PARQUET ]
[ COMPRESSION = GZIP | SNAPPY | NONE ]
[ INCLUDE_QUERY_ID_IN_FILE_NAME = TRUE | FALSE ]
[ FILE_NAME_PREFIX = <file_string> ]
[ SINGLE_FILE = FALSE | TRUE ]
[ MAX_FILE_SIZE = <bytes> ]
[ OVERWRITE_EXISTING_FILES = FALSE | TRUE ]
[ HEADER = TRUE | FALSE ]
|
COPY (<select_query>)
  TO '<url>'
  CREDENTIALS = { AWS_ACCESS_KEY_ID = '<aws_access_key_id>' AWS_SECRET_ACCESS_KEY = '<aws_secret_access_key>' [ AWS_SESSION_TOKEN = '<aws_session_token>' ] | AWS_ROLE_ARN = '<aws_role_arn>' [ AWS_ROLE_EXTERNAL_ID = '<aws_role_external_id>' ] }
[ TYPE = CSV | TSV | JSON | PARQUET ]
[ COMPRESSION = GZIP | SNAPPY | NONE ]
[ INCLUDE_QUERY_ID_IN_FILE_NAME = TRUE | FALSE ]
[ FILE_NAME_PREFIX = <file_string> ]
[ SINGLE_FILE = FALSE | TRUE ]
[ MAX_FILE_SIZE = <bytes> ]
[ OVERWRITE_EXISTING_FILES = FALSE | TRUE ]
[ HEADER = TRUE | FALSE ]
```

## [](#parameters)Parameters

Parameter Description `<select_query>` Any valid `SELECT` statement. `<location_name>` The name of a location object that contains the Amazon S3 URL and credentials. This is the recommended approach for specifying the destination. See [CREATE LOCATION](/sql_reference/commands/data-definition/create-location.html) for more information. `<URL>` The path to an Amazon S3 URL where the query result file or files are saved. For example, `s3://my_bucket/my_folder`. `CREDENTIALS` The Amazon S3 credentials for accessing the specified `<s3_url>`. For more information, see [CREDENTIALS](#credentials). `TYPE` Specifies the file type to save to Amazon S3. If omitted, `CSV` is the default. `COMPRESSION` Specifies the compression algorithm to use for the output files. If omitted, defaults to GZIP. If `NONE` is specified, exported files are not compressed. For `TYPE=PARQUET`, `COMPRESSION` can be set to `NONE`, the default `GZIP`, or `SNAPPY`. Each output file is written in Parquet format, using the specified algorithm to compress each column according to the Parquet specification. This kind of compression is internal to the Parquet file. For `TYPE=CSV`, `TSV`, or `JSON`, `COMPRESSION` can be set to `NONE` or the default `GZIP`. With `GZIP`, the file is first created in the specified format and then compressed into a GZIP file. `INCLUDE_QUERY_ID_IN_FILE_NAME` Specifies whether a query ID is included in the file name. Each time the statement runs, Firebolt generates a new query ID. If omitted, defaults to `TRUE`, and Firebolt saves file names using the pattern `<query_id>.[type].gz`, for example, `123ABCXY2.csv.gz`. This allows Firebolt to generate unique file names by default. If `FALSE`, and `FILE_NAME_PREFIX` is not specified, files are exported with a generic `output` file name, for example, `output.csv.gz`. `FILE_NAME_PREFIX` Specifies an optional string to use in the file name. If `FILE_NAME_PREFIX` is omitted and `INCLUDE_QUERY_ID_IN_FILE_NAME` is `TRUE`, the exported file name is in the pattern `<query_id><file_string>.<type>.gz`. If `FILE_NAME_PREFIX` is specified and `INCLUDE_QUERY_ID_IN_FILE_NAME` is set to `FALSE`, the specified string replaces `output` and the file is in the pattern `<file_string>.<type>.gz`. `SINGLE_FILE` Specifies whether the export should be a single file or multiple files. If omitted, the default is `FALSE`, and the export is split based on the `MAX_FILE_SIZE` value. Exported files are appended with `_<n>` incrementally to indicate the position in series, starting with `0`. For example, the first file in a series might be named `123ABCXY2_0.parquet.gz`. If `TRUE`, only a single file is written. If set to `TRUE` and the file exceeds `MAX_FILE_SIZE`, an error occurs. `MAX_FILE_SIZE` Specifies the max file size in bytes. If omitted, the default value is 16,000,000 bytes (about 16 MB). The maximum file size that can be specified is 5,368,709,120 bytes (5 GB). `OVERWRITE_EXISTING_FILES` Specifies whether exported files should overwrite existing files of the same name in the specified Amazon S3 URL. If omitted, defaults to `FALSE`, and files of the same name are not overwritten. `HEADER` Specifies whether CSV files should contain a header line containing the column names. If omitted, defaults to `TRUE`. Must be omitted for TSV, JSON, or Parquet files.

## [](#credentials)Credentials

Firebolt needs permissions to write query results to the specified Amazon S3 URL. You can specify IAM credentials using the AWS access keys, and the specified credentials must be associated with a user with permissions to write objects to the bucket.

### [](#credentials-for-amazon-s3)Credentials for Amazon S3

Specify access key credentials using the syntax shown below.

```
  CREDENTIALS = { AWS_ACCESS_KEY_ID = '<aws_access_key_id>' AWS_SECRET_ACCESS_KEY = '<aws_secret_access_key>' [ AWS_SESSION_TOKEN = '<aws_session_token>' ] | AWS_ROLE_ARN = '<aws_role_arn>' [ AWS_ROLE_EXTERNAL_ID = '<aws_role_external_id>' ] }
```

For Amazon S3 locations, you can use either access key-based authentication or role-based authentication:

Parameter Description `AWS_ACCESS_KEY_ID` Your AWS access key ID `AWS_SECRET_ACCESS_KEY` Your AWS secret access key `AWS_SESSION_TOKEN` Optional temporary session token for temporary credentials `AWS_ROLE_ARN` The ARN of the IAM role to assume `AWS_ROLE_EXTERNAL_ID` Optional external ID for role assumption

For more information on how to create access keys, see [Creating Access Key and Secret ID](/Guides/loading-data/creating-access-keys-aws.html).

### [](#least-privileged-permissions)Least privileged permissions

The example AWS IAM policy statement below demonstrates the minimum actions that must be allowed for Firebolt to write query files to an example Amazon S3 URL. A permissions policy that allows at least these actions for the `<s3_url>` that you specify in the `COPY TO` statement must be attached to the user or role specified in the `CREDENTIALS` clause.

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

## [](#using-location-objects)Using Location Objects

Location objects provide a secure and reusable way to specify Amazon S3 destinations for your data exports. This is the recommended approach as it:

- Centralizes credential management
- Eliminates the need to specify credentials in each COPY statement
- Provides better security through role-based access control
- Simplifies maintenance and updates

For more information about location objects, see [CREATE LOCATION](/sql_reference/commands/data-definition/create-location.html).

### [](#example-with-location-object)Example with Location Object

```
COPY (SELECT * FROM test_table)
  TO my_export_location
  TYPE = JSON
  COMPRESSION = NONE
  SINGLE_FILE = TRUE
  FILE_NAME_PREFIX = 'latest-export'
  MAX_FILE_SIZE = 5000000000
  OVERWRITE_EXISTING_FILES = TRUE;
```

The location object `my_export_location` contains the Amazon S3 URL and credentials, making the statement cleaner and more secure.

## [](#examples)Examples

- [COPY TO with location object](#copy-to-with-location-object)
- [COPY TO with defaults and role ARN](#copy-to-with-defaults-and-role-arn)
- [COPY TO with single file and AWS access keys](#copy-to-with-single-file-and-aws-access-keys)
- [COPY TO with custom file name prefix and ARN with external ID](#copy-to-with-custom-file-name-prefix-and-arn-with-external-id)
- [COPY TO with custom file name and overwrite set to TRUE](#copy-to-with-custom-file-name-and-overwrite-set-to-true)

### [](#copy-to-with-location-object)COPY TO with location object

This example demonstrates the recommended approach using a location object. The location object contains the Amazon S3 URL and credentials, making the statement more secure and maintainable.

```
COPY (SELECT * FROM test_table)
  TO my_export_location
  TYPE = JSON
  COMPRESSION = NONE
  SINGLE_FILE = TRUE
  FILE_NAME_PREFIX = 'latest-export'
  MAX_FILE_SIZE = 5000000000
  OVERWRITE_EXISTING_FILES = TRUE;
```

Firebolt writes a single file as shown below.

```
s3://my_bucket/exports/
  latest-export.json
```

### [](#copy-to-with-defaults-and-role-arn)COPY TO with defaults and role ARN

The example below shows a `COPY TO` statement with minimal parameters that specifies an `AWS_ROLE_ARN`. Because `TYPE` is omitted, the file or files will be written in CSV format, and because `COMPRESSION` is omitted, they are compressed using GZIP (`*.csv.gz`).

```
COPY (SELECT * FROM test_table)
  TO 's3://my-bucket/path/to/data'
  CREDENTIALS = (AWS_ROLE_ARN='arn:aws:iam::123456789012:role/my-firebolt-role');
```

Firebolt assigns the query ID `16B903C4206098FD` to the query at runtime. The compressed output is 40 MB, exceeding the default of 16 MB, so Firebolt writes 4 files as shown below.

```
s3://my_bucket/my_fb_queries/
  16B903C4206098FD_0.csv.gz
  16B903C4206098FD_1.csv.gz
  16B903C4206098FD_2.csv.gz
  16B903C4206098FD_3.csv.gz
```

### [](#copy-to-with-single-file-and-aws-access-keys)COPY TO with single file and AWS access keys

The example below exports a single, uncompressed JSON file with the same name as the query ID. If the file to be written exceeds the specified `MAX_FILE_SIZE` (5 GB), an error occurs. AWS access keys are specified for credentials.

```
COPY (SELECT * FROM test_table)
  TO 's3://my-bucket/path/to/data'
  TYPE=JSON
  COMPRESSION=NONE
  CREDENTIALS=(AWS_ACCESS_KEY_ID='AKIAIOSFODNN7EXAMPLE' AWS_SECRET_ACCESS_KEY='wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY')
  INCLUDE_QUERY_ID_IN_FILE_NAME=TRUE
  SINGLE_FILE=TRUE
  MAX_FILE_SIZE=5000000000;
```

Firebolt writes a single file as shown below.

```
s3://my_bucket/my_fb_queries/
  16B90E96716236D0.json
```

### [](#copy-to-with-custom-file-name-prefix-and-arn-with-external-id)COPY TO with custom file name prefix and ARN with external ID

In the example below, because `FILE_NAME_PREFIX` parameter is specified, Firebolt adds a string to the query ID to form file names. The IAM role specified for `CREDENTIALS` also has an external ID configured in AWS, so the ID is specified.

```
COPY (SELECT * FROM test_table)
  TO 's3://my-bucket/path/to/data'
  TYPE=JSON
  COMPRESSION=NONE
  CREDENTIALS = (AWS_ROLE_ARN='arn:aws:iam::123456789012:role/my-firebolt-role' AWS_ROLE_EXTERNAL_ID='99291')
  FILE_NAME_PREFIX='_query_result';
```

Firebolt assigns the query an id of `16B90E96716236D6` at runtime and writes files as shown below.

```
s3://my_bucket/my_fb_queries/
  16B90E96716236D6_query_result_0.json
  16B90E96716236D6_query_result_1.json
  16B90E96716236D6_query_result_2.json
  16B90E96716236D6_query_result_3.json
  16B90E96716236D6_query_result_4.json
```

### [](#copy-to-with-custom-file-name-and-overwrite-set-to-true)COPY TO with custom file name and overwrite set to TRUE

In the example below, `INCLUDE_QUERY_ID_IN_FILE_NAME` is set to `FALSE` and a `FILE_NAME_PREFIX` is specified. In this case, Firebolt writes files using only the specified `FILE_NAME_PREFIX` as the file name. In addition, `SINGLE_FILE` and `OVERWRITE_EXISTING_FILES` are set to `TRUE` so that the Amazon S3 URL always contains a single file with the latest query results.

```
COPY (SELECT * FROM test_table)
  TO 's3://my_bucket/my_fb_query'
  TYPE=JSON
  COMPRESSION=NONE
  CREDENTIALS=(AWS_ROLE_ARN = 'arn:aws:iam::123456789012:role/my-firebolt-role')
  INCLUDE_QUERY_ID_IN_FILE_NAME=FALSE
  SINGLE_FILE=TRUE
  FILE_NAME_PREFIX='latest-fb-query-result'
  MAX_FILE_SIZE=5000000000
  OVERWRITE_EXISTING_FILES=TRUE;
```

Firebolt writes a single file as shown below.

```
s3://my_bucket/my_fb-query/
  latest-fb-query-result.json
```