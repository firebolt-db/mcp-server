# [](#integrate-paradime-with-firebolt)Integrate Paradime with Firebolt

![Paradime logo](../../assets/images/paradime_logo.png)

[Paradime](https://www.paradime.io/) is a unified platform for data science and analytics that streamlines workflows for data teams. It offers a collaborative workspace where data scientists and analysts can explore, analyze, and visualize data across multiple tools and environments. Paradime integrates with tools including Jupyter notebooks, SQL editors, and Tableau. You can use the Paradime connector to link the Paradime platform directly to Firebolt’s cloud data warehouse. This connection allows you to run SQL queries, visualize results, and collaborate with team members all within the Paradime workspace.

This guide shows you how to connect Paradime to Firebolt using the Paradime user interface (UI). You must have a Firebolt account, a Firebolt service account, access to a Firebolt database, and an account with Paradime. These instructions build on the steps in Paradime’s [Getting Started with your Paradime Workspace](https://docs.paradime.io/app-help/guides/paradime-101/getting-started-with-your-paradime-workspace) guide, providing Firebolt-specific configuration details.

Topics:

- [Prerequisites](#prerequisites)
- [Create a Paradime workspace](#create-a-paradime-workspace).
- (Optional) [Create a schedule](#create-a-schedule-optional).

## [](#prerequisites)Prerequisites

Before you can connect Paradime to Firebolt, you must have the following:

1. **Firebolt Account**: Ensure that you have access to an active Firebolt account. If you don’t have access, you can [sign up for an account](https://www.firebolt.io/sign-up). For more information about how to register with Firebolt, see [Get started with Firebolt](/Guides/getting-started/).
2. **Service Account**: You must have access to an active Firebolt [service account](/Guides/managing-your-organization/service-accounts.html), which facilitates programmatic access to Firebolt.
3. **Firebolt Database**: You must have access to a Firebolt database. If you don’t have access, you can [create a database](/Guides/getting-started/get-started-sql.html#create-a-database).
4. **Paradime Account**: You must have access to an active Paradime account. If you don’t have access, you can [sign up](https://app.paradime.io) for one.

## [](#create-a-paradime-workspace)Create a Paradime workspace

Create a Paradime workspace to connect to Firebolt as follows:

01. In the Paradime UI, navigate to your account profile in the upper-right corner of the page.
02. Select **Profile Settings**.
03. In the **Workspaces** window, select the **New Workspace** button.
04. Enter a descriptive name for your workspace in the text box under **Name**.
05. Select **Create Workspace**.
06. Select **Continue**.
07. Select the most recent dbt-core version from the drop-down list.
08. Select **Continue**.
09. Select a dbt repository. You can either use an existing data build tool ([dbt](https://www.getdbt.com/blog/what-exactly-is-dbt)) repository or fork Firebolt’s sample [Jaffle Shop](https://github.com/firebolt-db/jaffle_shop_firebolt) repository from GitHub. Paradime supports the following providers: Azure DevOps, Bitbucket, GitHub, and GitLab.
10. Select **Next**.
11. Enter the SSH URI for your repository in the text box under **Repository URI**. Copy the key that appears under the **Deploy Key**.
12. Add the new deploy key to your dbt repository and allow write access. The following are resources for providers supported by Paradime:
    
    - [Add the deploy key](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/managing-deploy-keys#set-up-deploy-keys) in **github**.
    - [Add a deployment key](https://www.atlassian.com/blog/bitbucket/deployment-keys) in **Bitbucket**.
    - Use [deploy keys](https://docs.gitlab.com/ee/user/project/deploy_keys/) in **gitlab**.
    - [Use SSH key authentication](https://learn.microsoft.com/en-us/azure/devops/repos/git/use-ssh-keys-to-authenticate?view=azure-devops) to connect with **Azure DevOps**.
13. Select **Continue**.
14. If your repository connected successfully, select **Continue**.
15. Select **Firebolt** from the choices under **Warehouse connection**.
16. Under **Connection Settings**, enter the following:
    
    1. **Profile Name** – The name of a [connection profile](https://docs.getdbt.com/docs/core/connect-data-platform/connection-profiles) that is defined in `dbt_project.yaml` by a workspace administrator, and contains configurations including credentials to connect to a data warehouse. For more information, see Paradime’s [Setting up your profile](https://docs.getdbt.com/docs/core/connect-data-platform/connection-profiles#setting-up-your-profile) guide.
    2. **Target** – Specify the [target variable](https://docs.getdbt.com/reference/dbt-jinja-functions/target) that contains information about your data warehouse connection including its name, schema, and type.
    3. **Host Name** – Enter `api.app.firebolt.io`.
17. Under **Development Credentials**, enter the following:
    
    1. **Client Id** – Enter your Firebolt [service account ID](/Guides/managing-your-organization/service-accounts.html#get-a-service-account-id). Do not enter your Firebolt login email.
    2. **Client Secret** – Enter your Firebolt [service account secret](/Guides/managing-your-organization/service-accounts.html#generate-a-secret). Do not enter your Firebolt password.
    3. **Account Name** – Enter your Firebolt [account name](/Guides/managing-your-organization/managing-accounts.html).
    4. **Engine Name** – Enter the name of the engine where you want to run your queries.
    5. **Database Name** – Specify the Firebolt database name.
    6. Select **Test Connection** to connect to Firebolt and authenticate.
    7. Select **Next**.

For more information about the previous connection settings, see Paradime’s documentation to [add a development connection](https://docs.paradime.io/app-help/documentation/settings/connections/development-environment/firebolt).

## [](#create-a-schedule-optional)Create a schedule (Optional)

Paradime offers a scheduling feature using a [Bolt user interface](https://docs.paradime.io/app-help/documentation/bolt) to automatically run dbt commands on a specified interval or event. You can use Bolt to run a dbt job in a production environment, in a test environment prior to merging changes to production, or in an environment that runs jobs only on changed models.

To create a new schedule:

1. Login to your [Paradime account](https://app.paradime.io/?target=main-app).
2. Select **Bolt** from the left navigation bar.
3. Select **+ New Schedule**.
4. Select a pre-configured template from a list of popular Bolt templates or create a new schedule using a blank template. For information about how to configure settings in a Paradime schedule, see [Schedule Fields](https://docs.paradime.io/app-help/guides/paradime-101/running-dbt-in-production-with-bolt/creating-bolt-schedules#ui-based-schedule-fields).
5. Select **Publish**.
6. To view the new schedule, select **Bolt** from the left navigation pane.

# [](#additional-resources)Additional resources

- Learn about the [Paradime integrated development Environment](https://docs.paradime.io/app-help/guides/paradime-101/getting-started-with-the-paradime-ide).
- Learn to use the [Bolt scheduler](https://docs.paradime.io/app-help/bolt-scheduler/running-dbt-tm-in-production/creating-bolt-schedules) to run your dbt jobs.
- Learn how to [manage your Bolt schedule](https://docs.paradime.io/app-help/documentation/bolt/managing-schedules).