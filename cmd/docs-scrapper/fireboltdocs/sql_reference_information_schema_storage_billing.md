# [](#information-schema-for-storage-billing)Information schema for storage billing

You can use the `information_schema.storage_billing` view to see daily billing information for storage in all the accounts across all the regions in organization.

```
SELECT
  *
FROM
  information_schema.storage_billing;
```

## [](#columns-in-information_schemastorage_billing)Columns in information\_schema.storage\_billing

Each row has the following columns.

Column Name Data Type Description account\_name TEXT Account for which the storage is billed region TEXT Region where the data is stored usage\_date DATE Date for which the usage is reported consumed\_gib\_per\_month NUMERIC Amount of data billed for the given date in GiB billed\_cost NUMERIC The cost for the storage consumed for the given date is\_credit BOOLEAN Indicates whether costs were used as credit