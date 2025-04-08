# [](#integrate-estuary-flow-with-firebolt)Integrate Estuary Flow with Firebolt

![Estuary logo](../../assets/images/estuary.png)

Estuary Flow is a real-time data integration platform designed to streamline the movement and transformation of data between diverse sources and destinations. It provides an event-driven architecture and a user-friendly interface for building pipelines with minimal effort. You can use Flow to set up pipelines to load data from various sources, such as cloud storage and databases, into Firebolt’s cloud data warehouse for low-latency analytics.

This guide shows you how to set up a Flow pipeline that automatically moves data from your Amazon S3 bucket to your Firebolt database using the Estuary Flow user interface (UI). You must have access to an Estuary Flow account, an Amazon S3 bucket, and a Firebolt service account.

Topics:

- [Integrate Estuary Flow with Firebolt](#integrate-estuary-flow-with-firebolt)
  
  - [Prerequisites](#prerequisites)
  - [Configure your Estuary Flow source](#configure-your-estuary-flow-source)
  - [Configure your Estuary Flow destination](#configure-your-estuary-flow-destination)
  - [Monitor your materialization](#monitor-your-materialization)
  - [Validate your materialization](#validate-your-materialization)
  - [Additional resources](#additional-resources)

## [](#prerequisites)Prerequisites

1. **Estuary Flow account** – You must have access to an active Estuary Flow account. If you do not have access, you can [sign up](https://www.estuary.dev) with Estuary.
2. **Amazon S3 bucket** – you must have access to the following:
   
   - An [AWS Access Key ID and AWS Secret Access Key](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html) for an Amazon S3 bucket.
   - The name and path to an [Amazon S3 bucket](https://aws.amazon.com/s3/) that contains your data.
3. **Firebolt service account** –
   
   - Access to an organization in Firebolt. If you don’t have access, you can [create an organization](/Guides/managing-your-organization/creating-an-organization.html).
   - Access to a Firebolt database and engine. If you don’t have access, you can [create a database](/Guides/getting-started/get-started-sql.html#create-a-database) and [create an engine](/Guides/getting-started/get-started-sql.html#create-an-engine).
   - Access to a Firebolt service account, which is used for programmatic access, its [service account ID](/Guides/managing-your-organization/service-accounts.html#get-a-service-account-id) and [secret](/Guides/managing-your-organization/service-accounts.html#generate-a-secret-using-the-ui). If you don’t have access, you can [create a service account](/Guides/managing-your-organization/service-accounts.html#create-a-service-account).

## [](#configure-your-estuary-flow-source)Configure your Estuary Flow source

To set up an Estuary Flow pipeline that automatically moves data from your Amazon S3 bucket, you must create a capture that defines how and where data should be collected. Create a capture for the Estuary Flow source as follows:

1. Sign in to your [Estuary Flow Dashboard](https://dashboard.estuary.dev).
2. Select **Sources** from the left navigation pane.
3. In the **Sources** window, select **+ NEW CAPTURE**.
4. From the list of available connectors, navigate to **Amazon S3**, and select **Capture**.
5. Under **Capture Details**, enter a descriptive name for your capture in the text box under **Name**.
6. Under **Endpoint Config**, enter the following:
   
   1. **AWS Access Key ID** – The AWS account ID associated with the Amazon S3 bucket containing your data.
   2. **AWS Secret Access Key** – The AWS secret access key associated with the Amazon S3 bucket containing your data.
   3. **AWS Region** – The [AWS region](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/) that contains your Amazon S3 bucket. For example: `us-east-1`.
   4. **Bucket** – The name of your Amazon S3 bucket. For example, `firebolt-publishing-public`.
   5. **Prefix** (Optional) – A folder or key prefix to restrict the data to a specific path within the bucket. An example prefix structure follows: `/help_center_assets/firebolt_sample_dataset/levels.csv`.
   6. **Match Keys** (Optional) – Use a filter to include only specific object keys under the prefix, narrowing the capture’s scope.
7. Select the **NEXT** button in the upper-right corner of the page.
8. Test and save your connection as follows:
   
   1. Select **TEST** in the upper-right corner of the page. Estuary will run a test for your capture and display **Success** if it completes successfully.
   2. Select **CLOSE** in the bottom-right corner of the page.
   3. Select the **SAVE AND PUBLISH** button in the upper-right corner of the page. Estuary will test, save, and publish your capture and display **Success** if it completes successfully.
   4. Select **CLOSE** in the bottom-right corner of the page.

## [](#configure-your-estuary-flow-destination)Configure your Estuary Flow destination

To set up an Estuary Flow pipeline that automatically moves data from your Amazon S3 bucket, you must create a materialization that defines how the data should appear in the destination system, including any schema or transformation logic. Create a materialization for the Estuary Flow destination as follows:

1. Select **Destinations** from the left navigation pane.
2. Select the **+ NEW MATERIALIZATION** button in the upper-left corner of the page.
3. Navigate to the **Firebolt** connector and select **Materialization**.
4. Under **Materialization Details**, enter a descriptive name for your materialization in the text box under **Name**.
5. Under **Endpoint Config**, enter the following:
   
   01. **Client ID** – The service account ID for your Firebolt service account.
   02. **Client Secret** – The secret for your Firebolt service account.
   03. **Account Name** – The name of your service account.
   04. **Database** – The name of the Firebolt database where you want to put your data. For example, `my-database`.
   05. **Engine Name** – The name of the Firebolt engine to run the queries. For example: `my-engine-name`.
   06. **S3 Bucket** – The name of the Amazon S3 bucket to store temporary intermediate files related to the operation of the external table. For example, `my-bucket`.
   07. **S3 Prefix** – (Optional) A folder or key prefix to restrict the data to a specific path within the bucket. An example prefix structure follows the format in: `temp_files/`.
   08. **AWS Key ID** – The access key ID for the AWS account linked to the Amazon S3 bucket for temporary file storage.
   09. **AWS Secret Key** – The AWS secret key associated with the Amazon S3 bucket to store temporary files.
   10. **AWS Region** – The [AWS region](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/) of your Amazon S3 bucket. For example: `us-east-1`.
6. Select the **NEXT** button in the upper-right corner of the page.
7. Under **Source Collections**, do the following:
   
   1. Select **Source From Capture**.
   2. In the **Captures** window, select the checkbox next to the Amazon S3 source you specified when you configured your Estuary Flow source.
   3. Select the **CONTINUE** button in the bottom-right corner of the page.
   4. Verify that the **Table** name and type in the **CONFIG** tab under **Resource Configuration** are correct, and update if necessary.
   5. (Optional) Choose **Refresh** next to **Field Selection** to preview the fields, their types, and actions that will be written to Firebolt.
8. Test and save your materialization as follows:
   
   1. Select the **TEST** button in the upper-right corner of the page. Estuary will run a test for your materialization and display **Success** if it completes successfully.
   2. Select **CLOSE** in the bottom-right corner of the page.
   3. Select the **SAVE AND PUBLISH** button in the upper-right corner of the page. Estuary will test, save, and publish your materialization and display **Success** if it completes successfully.
   4. Select **CLOSE** in the bottom-right corner of the page.

## [](#monitor-your-materialization)Monitor your materialization

You can monitor your new data pipeline in Estuary Flow’s dashboard as follows:

1. Select **Destinations** from the left navigation pane.
2. Select your newly created materialization to view a dashboard with the following tabs:
   
   1. **OVERVIEW** – Provides a high-level summary of the materialization that includes throughput over time.
   2. **SPEC** – Displays the configurations and specifications of the materialization that includes schema mapping from the source to destination, the configuration of the destination, and any filters or constrains on the materialized data.
   3. **LOGS** – Provides records of materialization activity including success and failure events, messages, and errors.

Ensure that your data is being ingested and transferred as expected.

## [](#validate-your-materialization)Validate your materialization

You can validate that your data has arrived at Firebolt as follows:

1. Log in to the [Firebolt Workspace](https://firebolt.go.firebolt.io/signup).
2. Select the **Develop** icon (![The Firebolt Develop Space icon.](../../assets/images/develop-icon.png)) from the left navigation pane.
3. In the **Script Editor**, run a query on the table that you specified as an Estuary Flow destination to confirm the transfer of data as follows:
   
   1. Select the name of the database that you specified as your Estuary Flow destination from the drop-down list next to **Databases**.
   2. Enter a script in the script editor to query the table that you specified as an Estuary Flow destination. The following code example returns the contents of all rows and all columns from the `games` table:
   
   ```
    SELECT * FROM games
   ```
   
   You’ve successfully set up an Estuary Flow pipeline to move data from an Amazon S3 source to a Firebolt destination. Next, explore the following resources to continue expanding your knowledge base.

## [](#additional-resources)Additional resources

- Explore the [core concepts](https://docs.estuary.dev/concepts/) of Estuary Flow.
- Access [tutorials](https://docs.estuary.dev/getting-started/tutorials/) for Estuary Flow including a tutorial on [data transformation](https://docs.estuary.dev/guides/derivation_tutorial_sql/).
- Learn more about Estuary Flow’s [command line interface](https://docs.estuary.dev/concepts/flowctl/).