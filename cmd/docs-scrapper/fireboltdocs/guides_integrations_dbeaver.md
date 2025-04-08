# [](#integrate-with-dbeaver)Integrate with DBeaver

![DBeaver logo](../../assets/images/DBeaver-logo.png)

DBeaver is a free, open-source database administration tool that supports multiple database types. It provides a graphical interface for managing databases, running queries, and analyzing data. DBeaver is widely used for database development, troubleshooting, and administration, making it a versatile choice for both developers and database administrators. You can connect DBeaver to Firebolt using the [Firebolt JDBC driver](/Guides/developing-with-firebolt/connecting-with-jdbc.html).

- [Prerequisites](#prerequisites)
- [Add the Firebolt JDBC Driver in DBeaver](#add-the-firebolt-jdbc-driver-in-dbeaver)
- [Connect to Firebolt in DBeaver](#connect-to-firebolt-in-dbeaver)
- [Query Firebolt in DBeaver](#query-firebolt-in-dbeaver)
- [Additional Resources](#additional-resources)

## [](#prerequisites)Prerequisites

You must have the following prerequisites before you can connect your Firebolt account to DBeaver:

- **Firebolt account** – You need an active Firebolt account. If you do not have one, you can [sign up](https://go.firebolt.io/signup) for one.
- **Firebolt database and engine** – You must have access to a Firebolt database. If you do not have access, you can [create a database](/Guides/getting-started/get-started-sql.html#create-a-database) and then [create an engine](/Guides/getting-started/get-started-sql.html#create-an-engine).
- **Firebolt service account** – You must have an active Firebolt [service account](/Guides/managing-your-organization/service-accounts.html) for programmatic access, along with its ID and secret.
- **Sufficient permissions** – Your service account must be [associated](/Guides/managing-your-organization/service-accounts.html#create-a-user) with a user. The user should have [USAGE](/Overview/Security/Role-Based%20Access%20Control/database-permissions/) permission to query your database, and [OPERATE](/Overview/Security/Role-Based%20Access%20Control/engine-permissions.html) permission to start and stop an engine if it is not already started. It should also have at least USAGE and SELECT [permissions](/Overview/Security/Role-Based%20Access%20Control/database-permissions/schema-permissions.html) on the schema you are planning to query.
- **DBeaver installed** – You must have downloaded and installed [DBeaver](https://dbeaver.io/download/).

## [](#add-the-firebolt-jdbc-driver-in-dbeaver)Add the Firebolt JDBC Driver in DBeaver

To connect to Firebolt, you must add the Firebolt JDBC driver to DBeaver as follows:

1. Download the [Firebolt JDBC driver](/Guides/developing-with-firebolt/connecting-with-jdbc.html#download-the-jar-file).
2. In the DBeaver user interface (UI), under **Database**, select **Driver Manager**.
3. In **Driver Manager**, select **New** and enter the following parameters:
   
   - **Driver Name**: `Firebolt`
   - **Class Name**: `com.firebolt.FireboltDriver`
4. Select the **Libraries** tab.
5. Select **Add File**, and then select the JDBC driver you downloaded in the first step.
6. Select **Close**.

## [](#connect-to-firebolt-in-dbeaver)Connect to Firebolt in DBeaver

To connect to Firebolt, you must configure a new database connection in DBeaver as follows:

1. In DBeaver, select **Database**, then **New Database Connection**.
2. Enter `Firebolt` in the search box, then select it from the list.
3. Select **Next&gt;**.
4. Enter the connection parameters in the **Main** tab as follows:
   
   Parameter Description **JDBC URL** Use `jdbc:firebolt:<db_name>?engine=<engine_name>&account=<account_name>` replacing `<db_name>` with your Firebolt [database name](/Overview/indexes/using-indexes.html#databases), `<engine_name>` with your [engine name](/Guides/getting-started/get-started-sql.html#create-an-engine) and `<account_name>` with your [account name](/Guides/managing-your-organization/managing-accounts.html). **Username** Your Firebolt [service account](/Guides/managing-your-organization/service-accounts.html#get-a-service-account-id) ID. **Password** Your Firebolt [service account](/Guides/managing-your-organization/service-accounts.html#generate-a-secret) secret.
5. Select **Test Connection** to verify the connection. Ensure your Firebolt database is running before testing.
6. If the connection is successful, select **Finish**.

## [](#query-firebolt-in-dbeaver)Query Firebolt in DBeaver

1. In the database navigator, right-click or open the context menu of your Firebolt connection, select **SQL Editor**, then select **New SQL Script**.
2. Enter SQL queries into the SQL editor to interact with your Firebolt database.

## [](#additional-resources)Additional Resources

- Learn more about the [Firebolt JDBC driver](/Guides/developing-with-firebolt/connecting-with-jdbc.html).
- Explore [DBeaver’s documentation](https://dbeaver.io/documentation/) for details on its UI, integrations, tools, and features.
- Discover other tools that [Firebolt integrates](/Guides/integrations/integrations.html) with.