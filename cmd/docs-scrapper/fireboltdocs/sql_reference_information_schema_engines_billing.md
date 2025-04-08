# [](#information-schema-for-engines-billing)Information schema for engines billing

You can use the `information_schema.engines_billing` view to see daily billing information for all the engines in all the accounts across all the regions in organization.

```
SELECT
  *
FROM
  information_schema.engines_billing;
```

## [](#columns-in-information_schemaengines_billing)Columns in information\_schema.engines\_billing

Each row has the following columns.

Column Name Data Type Description engine\_name TEXT Name of the engine account\_name TEXT Account to which the engine belongs engine\_description TEXT Description of engine as entered by user when the engine is created region TEXT Region where the engine was created usage\_date DATE Date for which the usage is reported consumed\_fbu NUMERIC Number of FBUs consumed by the engine for the given date billed\_cost NUMERIC The cost for the FBUs consumed for the given date is\_credit BOOLEAN Indicates whether costs were used as credit