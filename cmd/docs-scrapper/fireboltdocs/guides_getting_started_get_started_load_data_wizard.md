# [](#get-started-using-a-wizard)Get started using a wizard

The **Load data** wizard guides you through creating a database and engine, and loading data from an Amazon S3 bucket. You can specify basic configurations, including what character to use as a file delimiter, which columns to import and their schema. After loading your data, continue working in the **Develop Space** to run and optimize a query, and export to an external table, as shown in the following diagram:

![A simple workflow using the load data wizard starts with registering, using the wizard, running a query, optimizing your workflow, and cleaning up. ](../../assets/images/get_started_wizard_workflow.png)

## [](#register-with-firebolt)Register with Firebolt

![The first step in getting started is to register with Firebolt.](../../assets/images/get_started_wizard_register.png)

Use the following steps to register with Firebolt:

1. [Sign up](https://go.firebolt.io/signup) on Firebolt’s registration page. Fill in your email, name, choose a password, and select **Get Started**.
2. Firebolt will send a confirmation to the address that you provided. To complete your registration, select **Verify** in the email to take you to Firebolt’s [login page](https://go.firebolt.io/login).
3. Type in your email and password and select **Log In**.

New accounts receive credits ($200) to get started exploring Firebolt’s capabilities. Credits must be used within 30 days of account creation.

Firebolt’s billing is based on engine runtime, measured in seconds. AWS S3 storage costs are passed through at the rate of $23 per TB per month. Your cost depends primarily on which engines you use and how long those engines are running.

You can view your total cost in FBU up to the latest second and in $USD up to the latest day. For more information, see the following **Create a Database** section. For more information about costs, see [Data Warehouse Pricing](https://www.firebolt.io/pricing). If you need to buy additional credits, connect Firebolt with your AWS Marketplace account. For more information about AWS Marketplace, see the following section: [Registering through AWS Marketplace section](/Guides/getting-started/get-started-next.html#register-through-the-aws-marketplace).

## [](#use-the-load-data-wizard)Use the Load data wizard

![After registering, use the load data wizard to create a database, engine, and load data.](../../assets/images/get_started_wizard_wizard.png)

You can use the **Load data** wizard to load data in either CSV or Parquet form.

To start the **Load data** wizard, select the plus (+) icon in the **Develop Space** next to **Databases** in the left navigation pane and select **Load data**. The wizard will guide you through creating a database, an engine, and loading data. See [Load data using a wizard](/Guides/loading-data/loading-data-wizard.html#load-data-using-a-wizard) for detailed information about the workflow and the available options in the wizard.

Even though the **Load data** wizard creates a database and engine for you, the [**Create a Database**](/Guides/getting-started/get-started-sql.html#create-a-database) and [**Create an Engine**](/Guides/getting-started/get-started-sql.html#create-an-engine) sections in the [Use SQL to load data](/Guides/getting-started/get-started-sql.html) guide contain useful information about billing for engine runtime and schema.

To use the **Load data** wizard, select the plus (+) icon. For detailed information about how to use the **Load data** wizard, see the [Load data](/Guides/loading-data/loading-data.html) guide.

## [](#run-query-optimize-clean-up-and-export)Run query, optimize, clean up, and export

![After using the load data wizard, a simple workflow continues with running a query, optimization, cleaning up, and optionally exporting a dataset.](../../assets/images/get_started_wizard_next.png)

After you have loaded your data in the wizard, the rest of the steps in getting started are the same as if you ran your workflow in SQL. You can use either the **Develop Space** in the **Firebolt Workspace** to enter SQL, or use the [Firebolt API](/Guides/query-data/using-the-api.html).

- For information about how to get started running a query, see [Run query](/Guides/getting-started/get-started-sql.html#run-query).
- For information about how to get started optimizing your workflow, see [Optimize your workflow](get-started-sql#optimize-your-workflow).
- For information about how to get started cleaning up resources and data, see [Clean up resources](./get-started-sql#clean-up).
- For information on how to export your data, see [Export data](/Guides/getting-started/get-started-sql.html#export-data).

## [](#next-steps)Next steps

To continue learning about Firebolt’s architecture, capabilities, using Firebolt after your trial period, and setting up your organization, see [Resources beyond getting started](/Guides/getting-started/get-started-next.html).