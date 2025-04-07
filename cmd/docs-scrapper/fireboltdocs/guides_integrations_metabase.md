![Metabase](/assets/images/metabase.png)

# [](#connecting-to-metabase)Connecting to Metabase

[Metabase](https://www.metabase.com/) is an open-source business intelligence platform. You can use Metabase’s user interface to explore, analyze, and visualize data, query databases, generate reports, and create dashboards.

This guide shows you how to [set up a Firebolt connector](#set-up-a-connector-to-metabase) for a self-hosted Metabase instance and how to [create a connection](#create-a-connection-to-metabase). If you are using either the managed or cloud-hosted version of [**Metabase Cloud**](https://www.metabase.com/docs/latest/cloud/start), you can skip directly to the [Create a Connection](#create-a-connection-to-metabase).

You can also watch a short video on how to connect Metabase to Firebolt:

**Topics:**

1. [Set up a connector to metabase](#set-up-a-connector-to-metabase)
2. [Create a connection to metabase](#create-a-connection-to-metabase)
3. [Additional Resources](#additional-resources)

### [](#set-up-a-connector-to-metabase)Set up a connector to metabase

Metabase can be deployed as a **self-hosted instance**, which is a version that you install and manage on your own server infrastructure. If you are using either the managed or cloud-hosted version of [**Metabase Cloud**](https://www.metabase.com/docs/latest/cloud/start), you can skip directly to the [Create a Connection](#create-a-connection-to-metabase).

For self-hosted deployments on-premises, the Firebolt connector must be installed manually using the following steps:

1. **Download the Firebolt Metabase driver**
   
   - Go to the [GitHub Releases page for Firebolt](https://github.com/firebolt-db/metabase-firebolt-driver/releases).
   - Locate the most recent version of the Firebolt driver, and download it.
2. **Move the driver file to the plugins directory**
   
   - Save the downloaded driver file in the `/plugins` directory on your Metabase host system.
   - By default, the `/plugins` directory is located in the same folder where the `metabase.jar` file runs. After completing these steps, the Firebolt connector will be available for configuration within Metabase.

### [](#create-a-connection-to-metabase)Create a connection to metabase

After setting up the Firebolt connector, use the following steps to create a connection between Metabase and your Firebolt database:

1. Open your Metabase instance’s home page in a web browser.
2. Select **Settings** from the top-right menu of the Metabase interface.
3. Select **Admin** from the dropdown menu.
4. On the **Admin** page, select **Databases** in the top navigation bar.
5. Select the **Add Database** button.
6. From the **Database Type** dropdown list, select **Firebolt**.
   
   Fill out the required connection details using the descriptions provided in the following table:
   
   Field Description **Display Name** A name to identify your database in Metabase. Use the same name as your Firebolt database for simplicity. **Client ID** The [service account ID](/Guides/managing-your-organization/service-accounts.html#get-a-service-account-id) associated with your Firebolt database. **Client Secret** The [secret for the service account](/Guides/managing-your-organization/service-accounts.html#generate-a-secret) associated with your Firebolt database. **Database name** Specify the name of the Firebolt database you want to connect to. **Account name** The name of your Firebolt account, which is required to log in and authenticate your database connection. **Engine name** Provide the name of the Firebolt engine that will be used to run queries against the database. **Additional JDBC options** Add any extra parameters needed for the connection, such as `connection_timeout_millis=10000`. For more options, access the [JDBC connection parameters guide](/Guides/developing-with-firebolt/connecting-with-jdbc.html#available-connection-parameters).
7. Select **Save** to store your database configuration.

Verify the connection by confirming that Metabase displays a success message indicating that your Firebolt database has been added successfully. If the connection fails, double-check your settings and ensure all required fields are correct.

### [](#additional-resources)Additional Resources

For more information about Metabase configuration and troubleshooting, refer to the following resources:

- [**Adding and Managing Databases**](https://www.metabase.com/docs/latest/databases/connecting) — Official Metabase documentation on connecting to data sources and managing database connections.
- [**Troubleshooting Database Connections**](https://www.metabase.com/docs/latest/troubleshooting-guide/db-connection) — Guidance on resolving issues when connecting [Metabase](https://www.metabase.com/docs/latest/databases/connecting) to your databases.
- [**Troubleshooting Database Performance**](https://www.metabase.com/docs/latest/troubleshooting-guide/db-performance) — Tips for identifying and addressing performance issues with connected databases.