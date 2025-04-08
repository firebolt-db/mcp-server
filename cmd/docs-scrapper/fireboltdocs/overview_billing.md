# [](#firebolt-pricing-and-deployment-models)Firebolt pricing and deployment models

Firebolt offers flexible deployment models that provide scalability and cost efficiency to meet diverse business needs. Choose between fully-managed and self-managed solutions tailored to your infrastructure and workload requirements.

![Firebolt offers fully-managed and self-managed deployment models](../../assets/images/deployment-models.png)

## [](#choose-your-deployment-model)Choose your deployment model

Firebolt offers two deployment models: **fully-managed** and **self-managed**. Each model includes specific editions and pricing options.

- The [fully-managed](#fully-managed-editions-and-pricing) deployment option includes **Standard**, **Enterprise**, and **Dedicated** editions with Pay-As-You-Go or Committed-Use discount pricing
- The [self-managed](#self-managed-editions-and-pricing) deployment model includes Firebolt’s **Firebolt Core** and **Private Cloud**, with pricing based on your infrastructure setup.

All fully-managed and **Private Cloud** editions have associated [support plans and service-level agreements](#support-plans-and-service-level-agreements) that define a time to first response.

The following sections outline the available deployment models, their editions, and associated pricing options.

### [](#fully-managed-editions-and-pricing)Fully managed editions and pricing

Firebolt manages compute infrastructure, software maintenance, and upgrades across all fully-managed editions. Firebolt handles multi-dimensional elasticity, high performance, scaling, maintenance, and performance optimizations.

- [Fully-managed editions](#fully-managed-editions) – Firebolt offers the **Standard**, **Enterprise** and **Dedicated** editions.
- [Fully-managed pricing model](#fully-managed-pricing-model) – Your total price consists of a storage and a compute cost.
  
  - [Data storage pricing](#data-storage-pricing) – The storage portion of your total cost.
  - [Compute usage pricing](#compute-usage-pricing) – The compute portion of your total cost.
- [Fully-managed pricing plans](#fully-managed-pricing-plans) – How to pay for resource usage.
  
  - [Billing setup and monitoring](#billing-setup-and-monitoring) – How to use the Billing dashboard to view resource consumption and set up billing and a plan.
  - [Set up billing for fully-managed plans](#set-up-billing-for-fully-managed-plans) – How to set up pay-as-you-go billing through AWS Marketplace.
  - [Sign up or change your fully-managed edition](#sign-up-or-change-your-fully-managed-edition) – How to set up or change your fully-managed edition type.

#### [](#fully-managed-editions)Fully-managed editions

There are three editions available: **Standard**, **Enterprise**, and **Dedicated**.

![Firebolt offers the Standard, Enterprise, and Dedicated editions with different features.](../../assets/images/firebolt-fully-managed.png)

**Standard edition**  
The **Standard** edition is best for teams seeking a low-cost entry point without the need to manage compute infrastructure. It offers high performance with sub-second query latency, flexible compute scaling within a single cluster. Security features include Role-Based Access Control (RBAC), single sign-on (SSO), multi-factor authentication (MFA), and audit logging. **Standard** includes a flexible compute option that is optimized for either storage or for compute. Firebolt manages your compute infrastructure, software maintenance, and upgrades.

**Enterprise edition**  
The **Enterprise** edition is designed for organizations that need advanced security, compliance, and automatic compute scaling. It includes all the features in the **Standard** edition, plus multi-cluster scaling, auto-scaling for concurrency, [AWS PrivateLink](/Guides/security/privatelink.html), and HIPAA compliance support.

**Dedicated edition**  
The **Dedicated** edition is ideal for organizations that require high levels of security and isolation, such as those operating in government clouds. It includes all the features in the **Enterprise** edition, plus complete data isolation on a single-tenant infrastructure, offering the highest level of security.

### [](#fully-managed-pricing-model)Fully-managed pricing model

The total cost for Firebolt’s fully-managed editions consists of a **cost for data storage** plus a **cost for compute usage**. The compute usage cost depends on the type of engine that you select: compute-optimized or storage-optimized.

![Total cost consists of data storage and compute usage with two types of compute options.](../../assets/images/compute-usage-cost.png)

#### [](#data-storage-pricing)Data storage pricing

Data storage costs are based on the amount of compressed data stored, including indexes and raw data. Pricing is based on Amazon S3 costs in your selected AWS region. There are no additional storage charges beyond S3 costs.

Available AWS Regions Price per TB/month ($USD) US East (N. Virginia) $23.00 US West (Oregon) $23.00 Europe (Frankfurt) $24.50 Europe (Ireland) $23.00 Asia Pacific (Singapore) $25.00

#### [](#compute-usage-pricing)Compute usage pricing

Compute costs are measured in Firebolt Units (FBUs) and vary based on [engine](/Overview/engine-fundamentals.html) node type, number of nodes or cluster size, the number of clusters and usage duration. Costs are only billed for the time Firebolt engines are running. Firebolt offers two compute family options:

- **Storage-optimized** (default): High SSD capacity for caching and production workloads.
- **Compute-optimized**: About 2x cheaper; ideal for development and test environments or workloads with smaller active datasets.

The following table outlines the available node types, their compute family, and the corresponding FBU sizing:

Node type Compute family Sizing in FBU Small (S) Storage-optimized 8 Medium (M) Storage-optimized 16 Large (L) Storage-optimized 32 Extra Large (XL) Storage-optimized 64 Small (S) Compute-optimized 4 Medium (M) Compute-optimized 8 Large (L) Compute-optimized 16 Extra Large (XL) Compute-optimized 32

Firebolt’s compute usage pricing is based on **FBUs**, which vary by node type and region. The following table provides hourly pricing for **storage-optimized** node types across **Standard** and **Enterprise** plans, distinguishing between **US and Non-US regions**. Compute costs are only incurred while Firebolt engines are running, with **per-second billing**.

**Storage-optimized pricing for compute usage**

Node Type Sizing in FBU Standard: US Region Pricing ($0.23/FBU/hr) Enterprise: US Region Pricing ($0.35/FBU/hr) Standard: Non-US Region Pricing ($0.28/FBU/hr) Enterprise: Non-US Region Pricing ($0.42/FBU/hr) **Small (S)** 8 $1.84 $2.80 $2.24 $3.36 **Medium (M)** 16 $3.68 $5.60 $4.48 $6.72 **Large (L)** 32 $7.36 $11.20 $8.96 $13.44 **Extra Large (XL)** 64 $14.72 $22.40 $17.92 $26.88

The following table provides hourly pricing for **compute-optimized** node types across **Standard** and **Enterprise** plans, distinguishing between **US and Non-US regions**. Compute costs are only incurred while Firebolt engines are running, with **per-second billing**.

**Compute-optimized pricing for compute usage**

Node Type Sizing in FBU Standard: US Region Pricing ($0.23/FBU/hr) Enterprise: US Region Pricing ($0.35/FBU/hr) Standard: Non-US Region Pricing ($0.28/FBU/hr) Enterprise: Non-US Region Pricing ($0.42/FBU/hr) **Small (S)** 4 $0.92 $1.40 $1.12 $1.68 **Medium (M)** 8 $1.84 $2.80 $2.24 $3.36 **Large (L)** 16 $3.68 $5.60 $4.48 $6.72 **Extra Large (XL)** 32 $7.36 $11.20 $8.96 $13.44

### [](#fully-managed-pricing-plans)Fully-managed pricing plans

Firebolt offers two pricing options for the **Standard** and **Enterprise** editions:

- **Pay-as-you-go**: A flexible plan that provides on-demand pricing with no upfront cost or commitment. This plan is ideal for startups or teams with unpredictable workloads. Customers get billed monthly based on actual usage and only pay for what they use with per-second billing.
- **Committed-use discounts**: A consumption model that provides discounted rates against a prepaid usage commitment. This plan is ideal for organizations with consistent, high-volume workloads and results in lower total costs compared to the pay-as-you-go plan. Once all prepaid FBU credits are consumed, your plan switches to the pay-as-you-go pricing model. Customers can always continue using Firebolt, with consumption either drawing from prepaid credits or transitioning to the pay-as-you-go model when credits run out.

Contact [support@firebolt.io](mailto:support@firebolt.io) to discuss a committed-use plan, annual pricing commitments, or **Dedicated** edition pricing.

#### [](#billing-setup-and-monitoring)Billing setup and monitoring

You can use Firebolt’s billing dashboard to monitor resource consumption, track expenses, monitor payments, and analyze billing trends efficiently. To view the billing dashboard, follow steps 1-3 in the following section to [set up billing for fully-managed plans](#set-up-billing-for-fully-managed-plans).

Billing invoices are generated on a monthly basis, and provide a detailed breakdown of resource consumption and associated costs.

#### [](#set-up-billing-for-fully-managed-plans)Set up billing for fully-managed plans

**Pay-As-You-Go setup via AWS Marketplace:**

1. Log in to [Firebolt’s Workspace](https://go.firebolt.io/login). If you haven’t yet registered with Firebolt, see the [Get Started](/Guides/getting-started/) guide.
2. In the Firebolt Workspace, select the Configure(![The Firebolt Configure Space icon](../../assets/images/configure-icon.png)) icon from the left navigation pane.
3. Under **Configure**, select **Billing**. This page allows you to view invoices and consumption details.
4. Select **Connect to AWS Marketplace**.
5. On AWS Marketplace, click **View Purchase Options** &gt; **Setup Your Account**.

Firebolt will bill you monthly through **AWS Marketplace** based on usage.

#### [](#sign-up-or-change-your-fully-managed-edition)Sign up or change your fully-managed edition

If you want to sign up or upgrade your fully-managed edition type, you can choose the **Standard**, **Enterprise** or **Dedicated** plan. Select your new choice in the **Firebolt Workspace** as follows:

1. Log in to [Firebolt’s Workspace](https://go.firebolt.io/login).
2. In the Firebolt Workspace, select the Configure(![The Firebolt Configure Space icon](../../assets/images/configure-icon.png)) icon from the left navigation pane.
3. Under **Configure**, expand the drop-down list next to **Billing**.
4. Under **Billing**, select **Plan** to open the list of available plans. Your active plan is labeled as **Current Plan**. For information about each plan, select **Learn more** to be directed to Firebolt’s [Pricing](https://www.firebolt.io/pricing) page.
5. Select your desired plan.
   
   1. To select the **Standard** or **Enterprise** plan, choose **Select plan** and confirm your selection.
   2. To select the **Dedicated** plan, do the following:
      
      1. Choose **Talk to Sales**.
      2. In the pop-up window, **Your email** is automatically populated with the email associated with your login.
      3. Enter a **Subject** or accept the default **Pricing plan** entry.
      4. Enter a **Description**.
      5. Select **Send** to notify Firebolt’s support team.

Changing your plan is **not immediate** and may take **up to 24 hours** to process. You will receive updates about the status of your request through email.

### [](#self-managed-editions-and-pricing)Self-managed editions and pricing

Firebolt offers two self-managed options, where you run Firebolt on your own infrastructure: **Firebolt Core** and **Private Cloud**.

![Firebolt offers two self-managed editions. One that is free that you manage, and one for a private cloud.](../../assets/images/self-managed-editions.png)

**Private Cloud (BYOC)**  
The **Private Cloud** edition is a BYOC (bring your own cloud) offering for organizations that want Firebolt’s software but prefer to use their own cloud infrastructure. Customers manage their own infrastructure for both compute and storage, whereas Firebolt manages hosting, Firebolt upgrades and maintenance. For BYOC pricing, contact [support@firebolt.io](mailto:support@firebolt.io).

**Firebolt Core**  
The **Firebolt Core** edition is a free downloadable version that can be deployed on cloud, on-premises, or on a local machine. This option is best for teams needing full control over deployment with a lightweight Firebolt engine. Customers manage compute and storage infrastructure, hosting, all software upgrades, and maintenance.

## [](#support-plans-and-service-level-agreements)Support plans and service level agreements

Firebolt offers support options based on your selected edition for fully-managed and **Private Cloud** editions.

Response Time Commitments (TFR = Time to First Response)

Severity Level Issue Type Business support: TFR for Standard Edition Premium support: TFR for Enterprise, Dedicated or Private Cloud Critical (Sev1) Service outage or major disruption Response within 4 hours Response within 30 minutes High (Sev2) Significant performance degradation Response within 8 business hours Response within 2 hours Medium (Sev3) Minor impact or feature issue Response within 24 business hours Response within 6 business hours Low (Sev4) General inquiries or documentation questions Response within 48 business hours Response within 24 business hours

**Premium support features**

**Enterprise**, **Private Cloud**, and **Dedicated** edition customers receive the following additional support benefits beyond response time commitments:

- **Proactive monitoring**: Alerts for issues and potential optimizations available to **Enterprise** edition customers.
- **Enhanced support channels**: Support through Slack, email and the [help menu](/Reference/help-menu.html) in the **Firebolt Workspace**.
- **Dedicated support engineer**: Customers are assigned a designated support engineer with deep knowledge of their environment, providing personalized support. This premium support level differs from **Standard** customers, who receive assistance from the general support pool.

Contact [support@firebolt.io](mailto:support@firebolt.io) to learn more about **Enterprise** edition Support offerings.

* * *

- [Explore compute cost](/Overview/billing/compute-cost.html)
- [Explore storage cost](/Overview/billing/storage-cost.html)