## [](#information-schema-views)Information Schema Views

The Firebolt information schema consists of a set of system views that contain metadata information on the objects defined in the current account, as well as database usage views. Firebolt supports `information_schema` schema based on ANSI SQL standard, but extended with Firebolt specific information, both as additional columns in standard views, and as additional views. For better compatibility with applications that rely on Postgres specific system views, Firebolt also supports `pg_catalog` schema and subset of Postgres compatible views in it.

Not all information schema views are available on Firebolt’s system engine or within the **Firebolt Workspace**. The following views are unavailable:

- `engine_metrics_history`
- `engine_running_queries`
- `engine_query_history`
- `engine_user_query_history`

* * *

- [Accounts](/sql_reference/information-schema/accounts.html)
- [Applicable roles](/sql_reference/information-schema/applicable-roles.html)
- [Catalogs](/sql_reference/information-schema/catalogs.html)
- [Columns](/sql_reference/information-schema/columns.html)
- [Enabled roles](/sql_reference/information-schema/enabled-roles.html)
- [Engine history](/sql_reference/information-schema/engine-history.html)
- [Engine metering history](/sql_reference/information-schema/engine-metering-history.html)
- [Engine metrics history](/sql_reference/information-schema/engine-metrics-history.html)
- [Engine query history](/sql_reference/information-schema/engine-query-history.html)
- [Engine running queries](/sql_reference/information-schema/engine-running-queries.html)
- [Engine user query history](/sql_reference/information-schema/engine-user-query-history.html)
- [Engines](/sql_reference/information-schema/engines.html)
- [Engines Billing](/sql_reference/information-schema/engines-billing.html)
- [Indexes](/sql_reference/information-schema/indexes.html)
- [Locations](/sql_reference/information-schema/locations.html)
- [Logins](/sql_reference/information-schema/logins.html)
- [Network policies](/sql_reference/information-schema/network_policies.html)
- [Object privileges](/sql_reference/information-schema/object-privileges.html)
- [Routines](/sql_reference/information-schema/routines.html)
- [Schemata](/sql_reference/information-schema/schemata.html)
- [Service accounts](/sql_reference/information-schema/service-accounts.html)
- [Storage Billing](/sql_reference/information-schema/storage-billing.html)
- [Storage Metering History](/sql_reference/information-schema/storage-metering-history.html)
- [Storage history](/sql_reference/information-schema/storage-history.html)
- [Tables](/sql_reference/information-schema/tables.html)
- [Transitive applicable roles](/sql_reference/information-schema/transitive-applicable-roles.html)
- [Users](/sql_reference/information-schema/users.html)
- [Views](/sql_reference/information-schema/views.html)