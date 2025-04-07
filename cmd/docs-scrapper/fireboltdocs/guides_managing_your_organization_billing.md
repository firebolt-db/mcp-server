# [](#billing)Billing

Firebolt bills are based on the consumption of resources within each account in your organization. This includes the total amount of data stored and engine usage.

- **Data storage** usage is calculated on the daily average amount of data (in bytes) stored under your Firebolt account name for indexes and raw compressed data.
- **Engine resources** usage is calculated with **one-second granularity** between the time Firebolt makes the engine available for queries and when the the engine moves to the stopped state.

## [](#set-up-account-billing-through-aws-marketplace)Set-up account billing through AWS Marketplace

To continue using Firebolt’s engines for query execution after your initial $200 credit, valid for 30 days, you’ll need to set-up a billing account by connecting your account to the [AWS Marketplace](https://aws.amazon.com/marketplace).

**Steps for registration:**

1. On the [Firebolt Workspace page](https://go.firebolt.io/), select the **Configure**(![AggIndex](../../assets/images/configure-icon.png)) icon from the left navigation pane.
2. Under **Organization settings**, select **Billing**.
3. Click **Connect to AWS Marketplace** to take you to the Firebolt page on AWS Marketplace.
4. On the AWS Marketplace page, click the **View Purchase Options** in the top right hand corner of the screen.
5. Click **Setup Your Account**.

Your account should now be associated with AWS Marketplace.

## [](#invoices)Invoices

Invoices for Firebolt engines and data are submitted through the AWS Marketplace. The final monthly invoice is available on the third day of each month through the AWS Marketplace. A billing cycle starts on the first day of the month and finishes on the last day of the same month.

## [](#viewing-billing-information)Viewing billing information

Users with the **Org Admin** role can monitor the cost history of each account in the organization.

**To view cost information for your organization** Organization cost details are captured in two information\_schema tables. Query those two tables and retrieve any information about the organization’s cost  
1\) [Engines billing](/sql_reference/information-schema/engines-billing.html) 2) [Storage billing](/sql_reference/information-schema/storage-billing.html)

Firebolt billing is reported to the AWS Marketplace at the beginning of the next day. By default, the **Accounts &amp; Billing** page displays the engine usage breakdown based on billing time. If you prefer to see the engine usage by actual usage day, you can click the **Engines breakdown** selector under the **Usage cost by engine** table and click **Actual running time**.