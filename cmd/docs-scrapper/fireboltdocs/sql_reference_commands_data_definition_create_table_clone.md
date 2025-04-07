# [](#create-table-clone)CREATE TABLE CLONE

Creates a table that is a copy of an existing table in the database.

The `CREATE TABLE CLONE` operation is extremely fast and highly cost-efficient because it copies the table’s structure and references, without duplicating the actual data.

The column names and data types of the cloned table are automatically inferred from the source table.

Once the clone is created, the cloned table and the source table become independent. Any future changes to the data or schema of either table will not affect the other.

## [](#source-table-indexes)Source Table Indexes

The primary index and partition expressions are copied from the source table.

By default, all aggregating indexes are implicitly cloned from the source table. This behavior can be modified using the `INCLUDING INDEXES` or `EXCLUDING INDEXES` options as follows:

- **INCLUDING INDEXES:** Clones all aggregating indexes from the source table, or clone a partial list of indexes if one is provided.
- **EXCLUDING INDEXES:** Clones the table without any of the aggregating indexes from the source table.

The newly created aggregating indexes will either retain the original aggregating index’s name or include an enumerated suffix if that name already exists.

## [](#syntax)Syntax

```
CREATE TABLE [IF NOT EXISTS] <target_table_name> CLONE <source_table_name> [{INCLUDING INDEXES [<index_name>,...] | EXCLUDING INDEXES}] 
```

**— OR —**

```
CREATE [OR REPLACE] TABLE <target_table_name> CLONE <source_table_name> [{INCLUDING INDEXES [<index_name>,...] | EXCLUDING INDEXES}] 
```

## [](#parameters)Parameters

Parameter Description `IF NOT EXISTS` By default, specifies that an existing table `<target_table_name>` will not be replaced with the new table schema definition. This clause prevents an error message that would otherwise occur, making it useful for scripted and programmatic implementations. `OR REPLACE` Specifies that an existing `<target_table_name>` will be replaced with the new table definition. `<target_table_name>` An identifier that specifies the name of the table. This name should be unique within the target database. `<source_table_name>` An identifier that specifies the name of the source table to clone. The source table must exist within the source database. `<index_name>` An aggregating index to clone from the source table, the index must exist and be associated with the source table.

### [](#examples)Examples

The following code creates a table with two aggregating indexes, which will serve as the base table for the upcoming example use cases:

```
CREATE TABLE sales_data (
    sale_id INTEGER,
    sale_date DATE,
    product_quantity INTEGER,
    customer_quantity INTEGER
);
CREATE AGGREGATING INDEX product_idx ON sales_data (COUNT(product_quantity));
CREATE AGGREGATING INDEX customer_idx ON sales_data (COUNT(customer_quantity));
```

**Example**

The following code example creates an exact copy of the existing `sales_data` table, including its structure and data:

```
CREATE TABLE cloned_sales_data CLONE sales_data;
```

**Example**

The following code example creates an exact copy of the `sales_data` table excluding any of its aggregating indexes:

```
CREATE TABLE cloned_sales_data CLONE sales_data EXCLUDING INDEXES;
```

**Example**

The following code example creates a copy of the `sales_data` table and includes only the specified only index `product_idx`:

```
CREATE TABLE cloned_sales_data CLONE sales_data INCLUDING INDEXES product_idx;
```

**Example**

The following code example creates a copy of the `sales_data` table and includes all associated indexes:

```
CREATE TABLE cloned_sales_data CLONE sales_data INCLUDING INDEXES;
```

# [](#clone-a-table-across-databases)CLONE A TABLE ACROSS DATABASES

You can clone a table from one database to another using a three-part identifier for either the source or target table.

When cloning across databases, you can keep the original table name as long as it does not already exist in the target database.

**Example**

The following code example creates an exact copy of the existing `sales_data` table from the `staging` database in the `prod` database:

```
CREATE TABLE prod.public.sales_data CLONE staging.public.sales_data;
```

### [](#limitations)Limitations

- You can only use `CREATE TABLE CLONE` to clone Firebolt-managed tables. External tables and views are not supported.
- `CREATE TABLE CLONE` is only supported for cloning tables created on Firebolt version 4.10 or higher.