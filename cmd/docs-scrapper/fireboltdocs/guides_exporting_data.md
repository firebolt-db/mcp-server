# [](#export-data)Export data

You can export data from a `SELECT` query directly to an Amazon S3 location using [COPY TO](/sql_reference/commands/data-management/copy-to.html). This method is more flexible and efficient than downloading query results manually from the **Firebolt Workspace**, making it ideal for data sharing, integration, and archival.

## [](#how-to-export-data)How to export data

The following code example uses `COPY TO` to export the result of a `SELECT` query from `my_table` to a specified Amazon S3 bucket in CSV format using the provided [AWS credentials](/sql_reference/commands/data-management/copy-to.html#credentials):

```
COPY (
    SELECT column1, column2 FROM my_table WHERE condition
) 
TO 's3://your-bucket/path/'
WITH (FORMAT = 'CSV')
CREDENTIALS = ('aws_key_id'='your-key' 'aws_secret_key'='your-secret');
```

## [](#choose-the-right-export-format)Choose the right export format

Format Best For Characteristics Recommended Use **CSV (Comma-Separated)** General data exchange, spreadsheets, SQL. Simple, widely supported, and easy to read. Best for spreadsheets, databases, or general data exchange. **TSV (Tab-Separated)** Structured text data. Like CSV, but uses tabs instead of commas. Best for Excel, databases, or general data exchange. **JSON** APIs, web applications, NoSQL databases. Flexible, human-readable, and supports nested data. Best for web apps, APIs, or NoSQL integrations. **PARQUET** Big data processing, analytics workloads. Compressed, columnar, and optimized for querying. Ideal for analytics, performance-sensitive workloads, and large datasets.

## [](#examples)Examples

**Export data in CSV format**

Use CSV when you need a simple, widely supported format for spreadsheets, relational databases, or data exchange.

The following code example exports `user_id`, `event_type`, and `timestamp` data and headers from the `user_events` table to a CSV file in an Amazon S3 bucket:

```
COPY (SELECT user_id, event_type, timestamp FROM user_events) 
TO 's3://my-export-bucket/user_events.csv'
WITH (FORMAT = 'CSV', HEADER = TRUE)
CREDENTIALS = ('aws_key_id'='your-key' 'aws_secret_key'='your-secret');
```

**Export data in Parquet format**

Parquet is best for big data workloads, as it offers compressed, columnar storage optimized for analytics and query performance.

The following code example exports all data from the `sales_data` table to an Amazon S3 bucket in Parquet format using the provided AWS credentials:

```
COPY (SELECT * FROM sales_data) 
TO 's3://my-export-bucket/sales_data.parquet'
WITH (FORMAT = 'PARQUET')
CREDENTIALS = ('aws_key_id'='your-key' 'aws_secret_key'='your-secret');
```

**Export data in JSON format**

JSON is ideal for APIs, web applications, and NoSQL databases, as it supports nested and flexible data structures.

The following code example exports `order_id` and `order_details` from the `orders` table to an Amazon S3 bucket in JSON format using the provided AWS credentials:

```
COPY (SELECT order_id, order_details FROM orders) 
TO 's3://my-export-bucket/orders.json'
WITH (FORMAT = 'JSON')
CREDENTIALS = ('aws_key_id'='your-key' 'aws_secret_key'='your-secret');
```

**Export data in TSV format**

TSV is similar to CSV but uses tab delimiters, making it useful for structured text data that may contain commas.

The following code example exports `name`, `age`, and `city` from the `customers` table to an Amazon S3 bucket in TSV format using the provided AWS credentials:

```
COPY (SELECT name, age, city FROM customers) 
TO 's3://my-export-bucket/customers.tsv'
WITH (FORMAT = 'TSV')
CREDENTIALS = ('aws_key_id'='your-key' 'aws_secret_key'='your-secret');
```

## [](#additional-considerations)Additional Considerations

**Performance tips**

- Export only required columns and use filters to reduce data volume
- Ensure proper permissions are set on your S3 bucket

**Security and credentials**

- Always use **secure AWS credentials**.
- Use **IAM roles** instead setting credentials directly in the code for better security.

## [](#next-steps)Next Steps

For more information about advanced options including **compression**, **partitioning**, and **null handling**, see [COPY TO](/sql_reference/commands/data-management/copy-to.html).