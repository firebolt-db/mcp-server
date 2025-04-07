# [](#information-schema-for-storage-metering-history)Information schema for storage metering history

You can use the `information_schema.storage_metering_history` view to see daily information about aggregated storage consumption within an account.

```
SELECT
  *
FROM
  information_schema.storage_metering_history;
```

## [](#columns-in-information_schemastorage_metering_history)Columns in information\_schema.storage\_metering\_history

Each row has the following columns.

Column Name Data Type Description usage\_date DATE Date for which the usage is reported consumed\_gib\_per\_month NUMERIC Number of bytes consumed by the account for the given date