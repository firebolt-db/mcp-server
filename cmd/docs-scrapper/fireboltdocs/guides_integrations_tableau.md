# [](#integrate-with-tableau)Integrate with Tableau

![Tableau icon](../../assets/images/Tableau.png)

[Tableau](https://www.tableau.com/) is a visual analytics platform that empowers users to explore, analyze, and present data through interactive visualizations. It supports diverse use cases such as data exploration, reporting, and collaboration, and helps users gain insights and make informed decisions. This guide shows you how to set up your Firebolt account to integrate with [Tableau Desktop](https://www.tableau.com/products/desktop) and [Tableau Exchange](https://exchange.tableau.com).

The latest Firebolt version is not compatible with Tableau Online, and you will not be able to connect it to your Firebolt account. You can only use the connector from Tableau Exchange with an older version of Firebolt. If you want to use the latest version, use Tableau Desktop or Tableau Server and follow the instructions below.

## [](#prerequisites)Prerequisites

You must have the following prerequisites before you can connect your Firebolt account to Tableau:

- **Tableau account** – You must have access to an active Tableau account. If you do not have access, you can [sign up](https://www.tableau.com/products/trial) for one.
- **Firebolt account** – You need an active Firebolt account. If you do not have one, you can [sign up](https://go.firebolt.io/signup) for one.
- **Firebolt database and table** – You must have access to a Firebolt database that contains a table with data ready for visualization. If you don’t have access, you can [create a database](/Guides/getting-started/get-started-sql.html#create-a-database) and then [load data](/Guides/loading-data/loading-data.html) into it.
- **Firebolt service account** – You must have access to an active Firebolt [service account](/Guides/managing-your-organization/service-accounts.html), which facilitates programmatic access to Firebolt, its ID and secret.
- **Firebolt user** – You must have a user that is [associated](/Guides/managing-your-organization/service-accounts.html#create-a-user) with your service account. The user should have [USAGE](/Overview/Security/Role-Based%20Access%20Control/database-permissions/) permission to query your database, and [OPERATE](/Overview/Security/Role-Based%20Access%20Control/engine-permissions.html) permission to start and stop an engine if it is not already started.

## [](#connect-to-tableau)Connect to Tableau

To connect to Tableau, you must download a Firebolt connector, a [JDBC driver](/Guides/developing-with-firebolt/connecting-with-jdbc.html#jdbc-driver), connect to Firebolt, and select a database and schema to query. You can either install [Tableau Desktop](https://www.tableau.com/products/desktop) for individual use or [Tableau Server](https://www.tableau.com/products/server) for centralized access to dashboards on a shared server.

1. **Download and install Tableau**
   
   1. To download Tableau’s Desktop, navigate to Tableau’s Desktop [download page](https://www.tableau.com/en-gb/products/desktop/download), and follow the prompts to install the program. To use Tableau Server, follow Tableau’s instructions for [installation and configuration](https://help.tableau.com/current/server/en-us/install_config_top.htm).
   2. Follow the prompts to install Tableau.
2. **Download the latest Firebolt connector**
   
   Download the latest version of Firebolt’s Tableau connector from Firebolt’s GitHub [repository](https://github.com/firebolt-db/tableau-connector/releases). The earliest version of the driver that is compatible with the latest version of Firebolt is [v1.1.0](https://github.com/firebolt-db/tableau-connector/releases/tag/v1.1.0). The name of the file has the following format: `firebolt_connector-<version>.taco`, and should be saved in a specific directory that depends on the operating system used as follows:
   
   For Tableau Desktop, save the file connector to:
   
   - Windows - `C:\Users\[Windows User]\Documents\My Tableau Repository\Connectors`
   - MacOS - `/Users/[user]/Documents/My Tableau Repository/Connectors`
   
   For any other installations including Tableau Server and older versions of Tableau, follow the steps in the Tableau [guide](https://help.tableau.com/current/pro/desktop/en-us/examples_connector_sdk.htm#use-a-connector-built-with-tableau-connector-sdk).
3. **Download the latest JDBC driver**
   
   Download a JDBC driver, which will allow Tableau to interact with a Firebolt databases using Java, from Firebolt’s GitHub [repository](https://github.com/firebolt-db/jdbc/releases). The name of the file has the following format: `firebolt-jdbc-<version>.jar`, and should be saved in a specific directory that depends on the operating system as follows:
   
   - Windows: `C:\Program Files\Tableau\Drivers`
   - Mac: `/Users/<username>/Library/Tableau/Drivers`
   - Linux: `/opt/tableau/tableau_driver/jdbc`
4. **Start Tableau and verify Firebolt connector availability**
   
   1. Start your Tableau Desktop or Server. If you already started Tableau prior to downloading the drivers, restart Tableau.
   2. In the left navigation panel, under **To a Server**, select the `>` to the right of **More…**.
   3. Search for and select the **Firebolt by Firebolt Analytics Inc** connector in the search bar.
   4. In the left navigation panel, under **To a Server**, select the `>` to the right of **More…**.
   5. Select **Firebolt Connector by Firebolt**.
   6. Enter the following parameters in the **General** tab:
      
      **Field** **Required** **Description** **Host** No Most users should not enter a value in the text box under `Host`. **Account** Yes The name of your Firebolt account within your organization. **Engine Name** Yes The name of the [engine](/Overview/engine-fundamentals.html) to run queries. **Database** Yes The name of the Firebolt [database](/Overview/indexes/using-indexes.html#databases) to connect to. **Client ID** Yes The [ID of your service account](/Guides/managing-your-organization/service-accounts.html#get-a-service-account-id). **Client Secret** Yes The [secret](/Guides/managing-your-organization/service-accounts.html#generate-a-secret) for your service account authentication.
   7. Select **Sign in**.
5. **Choose the database and the schema to query**
   
   After successful authentication, **Database** and **Schema** drop-down lists appear in the left navigation pane under **Connections**. The database name from the previous step appears in the database drop-down list. To change the database, you must repeat the previous step and set up a new connector.
   
   Choose the schema and tables as follows:
   
   1. Select the drop-down list under **Schema** to select a [schema](/Overview/indexes/using-indexes.html#schema). Most users should choose `public`. For more information about schema permissions and privileges, see [Schema permissions](/Overview/Security/Role-Based%20Access%20Control/database-permissions/schema-permissions.html).
   2. Drag and drop tables from the list of available tables in your schema to use them in Tableau.
6. **Visualize your data**
   
   Once your data source is selected you can begin visualizing the data by creating graphs and charts as follows:
   
   1. Select `Sheet 1` tab from the bottom-left corner of your Tableau window next to **Data Source**.
   2. In the left navigation panel under **Sheets**, drag and drop any available columns or pre-defined aggregation from your table into the Tableau workspace to start building charts. See Tableau’s [Build a view from scratch](https://help.tableau.com/current/pro/desktop/en-us/getstarted_buildmanual_ex1basic.htm) documentation for more information.

## [](#limitations)Limitations

- Firebolt does not support [Tableau Cloud](https://www.tableau.com/products/cloud-bi).
- Once you have set up a connection to Firebolt, you cannot change the database that you specified during setup. In order to change the database, you must repeat step 4 to **Start Tableau and verify Firebolt connector availability** in [Connect to Tableau](#connect-to-tableau) to set up a new connection.

## [](#additional-resources)Additional resources

- Watch Tableau’s [free training videos](https://www.tableau.com/en-gb/learn/training) on getting started, preparing data, and geographical analysis.
- Read Tableau’s data visualization [articles](https://www.tableau.com/en-gb/learn/articles) about creating effective, engaging, and interactive examples.
- Follow Tableau’s [blog](https://www.tableau.com/en-gb/blog) for new features and tips.