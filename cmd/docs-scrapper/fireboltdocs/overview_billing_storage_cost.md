# [](#explore-storage-cost)Explore storage cost

You can use data from the [`information_schema.storage_billing`](/sql_reference/information-schema/storage-billing.html) and [`information_schema.storage_history`](/sql_reference/information-schema/storage-history.html) views to analyze and optimize storage costs across different regions, accounts, and data usage trends. The following queries will guide you in tracking storage expenses, understanding regional cost variations, and identifying storage inefficiencies.

## [](#examples)Examples

**Track and attribute storage costs**

- [Track total storage cost over time](#track-total-storage-cost-over-time) – Learn how to monitor daily and monthly storage costs.
- [Identify storage cost by region](#identify-storage-cost-by-region) – Learn how to determine which regions contribute most to storage expenses.
- [Attribute storage costs to accounts](#attribute-storage-costs-to-accounts) – Learn how to attribute storage costs to specific customers or business units.
- [Track storage credits usage](#track-storage-credits-usage) – Learn how to monitor what percentage of your storage cost is paid using Firebolt credits.

**Monitor data efficiency**

- [Monitor active vs inactive data trends](#monitor-active-vs-inactive-data-trends) – Learn how to track the proportion of active and inactive data in your storage.
- [Identify databases with high inactive storage](#identify-databases-with-high-inactive-storage) – Learn how to find databases with significant amounts of inactive data.
- [Track storage growth by catalog](#track-storage-growth-by-catalog) – Learn how to monitor storage growth trends across catalogs, which include tables, views, and schemas.

**Evaluate cost efficiency and inactive data**

- [Evaluate cost efficiency per unit of storage](#evaluate-cost-efficiency-per-unit-of-storage) – Learn how to calculate the cost per unit storage to assess the efficiency of your storage expenses relative to the amount of data stored.
- [Assess inactive data as a percentage of total storage](#assess-inactive-data-as-a-percentage-of-total-storage) – Learn how to understand what percentage of your total storage is used by data that is inactive.
- [Attribute storage billing to catalogs](#attribute-storage-billing-to-catalogs) – Learn how to attribute storage billing to specific tables, views, and schemas.

### [](#track-total-storage-cost-over-time)Track total storage cost over time

You can track total storage cost over time to monitor spending trends, identify potential cost increases, and ensure that storage usage aligns with budget expectations. Query `information_schema.storage_billing` to help you proactively manage costs, so you can make adjustments to optimize storage efficiency and prevent unexpected charges.

The following code example calculates the daily storage cost and the total monthly storage cost by summing the billed costs for each day and using a window function to compute the monthly total for each day within the same month:

```
SELECT usage_date, 
       SUM(billed_cost) AS total_cost,
       DATE_TRUNC('month', usage_date) AS month, 
       SUM(SUM(billed_cost)) OVER (PARTITION BY DATE_TRUNC('month', usage_date)) AS monthly_cost
FROM information_schema.storage_billing
GROUP BY usage_date
ORDER BY usage_date;
```

### [](#identify-storage-cost-by-region)Identify storage cost by region

Use [`information_schema.storage_billing`](/sql_reference/information-schema/storage-billing.html) inside a query to understand which regions incur the highest storage costs, helping you optimize resource allocation based on geographical pricing differences.

The following code example calculates the total storage cost for each region and sorts the results in descending order of total cost:

```
SELECT region, SUM(billed_cost) AS total_cost
FROM information_schema.storage_billing
GROUP BY region
ORDER BY total_cost DESC;
```

### [](#attribute-storage-costs-to-accounts)Attribute storage costs to accounts

You can use `information_schema.storage_billing` to attribute storage costs to specific accounts or business units and gain visibility into where your storage expenses originate. This can help you track usage, optimize costs, and allocate resources more effectively.

The following code example calculates the total storage cost for each account and sorts the results in descending order of total cost:

```
SELECT account_name, SUM(billed_cost) AS total_cost
FROM information_schema.storage_billing
GROUP BY account_name
ORDER BY total_cost DESC;
```

### [](#track-storage-credits-usage)Track storage credits usage

Analyzing how much of your storage costs is paid with Firebolt credits versus billed costs helps optimize expenses, track credit usage, identify cost-saving opportunities, and improve budget management. Use `information_schema.storage_billing` to analyze how much of your storage cost is paid using Firebolt credits and how much is billed.

The following code example uses the `is_credit` field to calculate the total storage cost, grouped by whether the cost was paid through credits or billed:

```
SELECT is_credit, SUM(billed_cost) AS total_cost
FROM information_schema.storage_billing
GROUP BY is_credit;
```

### [](#monitor-active-vs-inactive-data-trends)Monitor active vs inactive data trends

Inactive data refers to data that is no longer actively used or is infrequently accessed but still consumes storage resources. Retaining large volumes of inactive data can result in higher costs and wasted storage. Use [`information_schema.storage_history`](/sql_reference/information-schema/storage-history.html) to identify this data so that you can make decisions about optimizing storage to reduce unnecessary expenses.

The following code example shows daily trends for active and inactive data in Gibibyte (GiB), which is equal to 1,073,741,824 bytes:

```
SELECT usage_date,
       SUM(active_data_size_bytes)/1024/1024/1024 AS active_gib,
       SUM(inactive_data_size_bytes)/1024/1024/1024 AS inactive_gib
FROM information_schema.storage_history
GROUP BY usage_date
ORDER BY usage_date;
```

### [](#identify-databases-with-high-inactive-storage)Identify databases with high inactive storage

You can identify databases that contain large amounts of inactive data, which may be a candidate for cleanup or archiving.

The following code example calculates the total inactive data size for each database:

```
SELECT catalog_name,
       SUM(inactive_data_size_bytes)/1024/1024/1024 AS inactive_gib
FROM information_schema.storage_history
GROUP BY catalog_name
ORDER BY inactive_gib DESC;
```

### [](#track-storage-growth-by-catalog)Track storage growth by catalog

You can use `information_schema.storage_history` to track storage growth by catalog, or objects that hold metadata and database objects such as tables, views, databases and schemas, to monitor how storage usage is changing across different databases or projects. Identify trends, optimize storage allocation, and ensure that resources are being used efficiently, while also providing insight into which catalogs may require further optimization or cost management.

The following code example calculates the total storage size in GiB for each catalog by summing active and inactive data sizes and ordering the results by catalog and usage date:

```
SELECT catalog_name, usage_date,
       (active_data_size_bytes + inactive_data_size_bytes)/1024/1024/1024 AS total_gib
FROM information_schema.storage_history
ORDER BY catalog_name, usage_date;
```

### [](#evaluate-cost-efficiency-per-unit-of-storage)Evaluate cost efficiency per unit of storage

You can use `information_schema.storage_billing` to evaluate cost efficiency per unit of storage by assessing whether storage costs are proportional to the amount of data being stored. This can help you identify discrepancies, optimize spending, and ensure that storage resources are being utilized efficiently. Determine if your storage costs align with the volume of data stored by calculating the cost per GiB.

The following code example calculates the cost per GiB of storage for each day by dividing the total billed cost by the total consumed GiB, and orders the results by usage data:

```
SELECT usage_date,
       SUM(billed_cost) / NULLIF(SUM(consumed_gib_per_month), 0) AS cost_per_gib
FROM information_schema.storage_billing
GROUP BY usage_date
ORDER BY usage_date;
```

### [](#assess-inactive-data-as-a-percentage-of-total-storage)Assess inactive data as a percentage of total storage

You can use `information_schema.storage_history` to measure the amount of inactive data as a percentage of total storage to identify how much of storage is being underutilized. Assessing inactive data as a percentage of total storage helps you understand how much of your storage is underutilized, enabling you to optimize costs by cleaning up or archiving inactive data.

The following code example calculates the percentage of inactive data relative to the total storage, composed of both active and inactive storage, for each day and orders the results by usage date:

```
SELECT usage_date,
       100 * SUM(inactive_data_size_bytes)::FLOAT / 
       NULLIF(SUM(active_data_size_bytes + inactive_data_size_bytes), 0) AS inactive_pct
FROM information_schema.storage_history
GROUP BY usage_date
ORDER BY usage_date;
```

### [](#attribute-storage-billing-to-catalogs)Attribute storage billing to catalogs

You can use `information_schema.storage_billing` to attribute storage billing to catalogs, or objects that hold metadata and database objects such as tables, databases, views, and schemas. Use this data to better understand and allocate storage costs to specific departments, projects, or business units, so that you can enable more accurate cost tracking, budgeting, and optimization across different catalogs.

The following code example calculates the total storage billing cost for each catalog by joining storage billing data with catalog account mappings and storage history, and then sorts the results by total cost in descending order:

```
SELECT storage_history.catalog_name, SUM(storage_billing.billed_cost) AS total_cost
FROM information_schema.storage_billing storage_billing
JOIN some_catalog_account_mapping m ON storage_billing.account_name = m.account_name
JOIN information_schema.storage_history sh ON m.catalog_name = storage_history.catalog_name AND storage_billing.usage_date = storage_history.usage_date
GROUP BY storage_history.catalog_name
ORDER BY total_cost DESC;
```