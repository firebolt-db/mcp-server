# [](#organizations-and-accounts)Organizations and accounts

A governance model can help manage cloud data warehouse resources by addressing challenges such as data security, cost management, resource isolation, and observability. For example, development, staging, and production environments often require isolation to prevent unintentional changes in development from affecting production or to restrict developers’ access to only their code and data. Similarly, departments may need isolated access to their resources while limiting access to others. Additionally, governance models can support consolidated billing while providing visibility into consumption by department or development environment.

To address these requirements, Firebolt supports concepts of organizations and accounts. You can have different accounts within your organization and additionally benefit from consolidated billing, unified authentication, and efficient account management across all accounts.

**Topics**

- [Organizations and accounts](#organizations-and-accounts)
  
  - [Organizations](#organizations)
  - [Accounts](#accounts)
    
    - [Users](#users)
    - [Roles](#roles)
      
      - [Firebolt built-in roles](#firebolt-built-in-roles)
        
        - [Public role](#public-role)
        - [System administrative role](#system-administrative-role)
        - [Account administrative role](#account-administrative-role)
        - [Organizational administrative role](#organizational-administrative-role)
  - [Additional resources](#additional-resources)
    
    - [Billing](#billing)
    - [Security](#security)
    - [Available regions](#available-regions)

The Firebolt object model is hierarchical and comes with strong containment properties in that parent objects can contain one or more child objects. Child objects are sole children of their parent objects and cannot be shared. Furthermore, there are two classes of objects: global and regional. Global objects are managed globally and can contain objects that are deployed and grouped regionally.

The following Firebolt object model depicts an **organization** at the highest level, with **four** layers directly underneath it:

![Firebolt's organizational structure has four layers](../../assets/images/manage-organization-org-structure.png)

## [](#organizations)Organizations

An organization is a fundamental object in Firebolt, providing a logical structure for managing accounts, billing, and authentication. When registering to Firebolt, the organization name you’ll provide is the same as the domain name you use in your email. Organization names are globally unique. No two organizations can have the same name, but organizations can contain multiple accounts. Each account can contain multiple objects including users, roles, databases, tables, views, and engines.

In the Firebolt object model, an organization has the following levels:

- **Login** - an email associated with a user that is used for identification. A single login can be associated with users across multiple accounts, but can only be associated with one user per account. For more information about logins, see [Manage logins](/Guides/managing-your-organization/managing-logins.html).
- **Service account** - Use a service account to access Firebolt programmatically through an API. For more information about how to set up and manage a service account, see [Manage programmatic access](/Guides/managing-your-organization/service-accounts.html).
- **Network policy** - Set a network security configuration that controls and restricts network access to specific resources within Firebolt based on IP addresses or IP ranges. See [Manage network policies](/Guides/security/network-policies.html).
- **Account** - A group of resources which can include a database, an engine and several users with associated logins or service accounts and roles. See the following [Accounts](#accounts) section for high-level information about accounts and [Manage accounts](/Guides/managing-your-organization/managing-accounts.html) for more detailed information. Under account are the following levels:
  
  - **User** - An individual with specific permissions and roles that allow access to and interaction with the database and engine within an account. A user must be associated with a login (or service account), and with an account. See the following [Users](#users) section for more high-level information and [Manage users and roles](/Guides/managing-your-organization/managing-users.html) for more detailed information.
  - **Role** - A set of permissions that defines a user’s access and capabilities, which can be assigned to one or more users to manage their privileges. See the following Roles section for more high-level information and Manage users and roles for more detailed information.
  - **Database** - A logical collection of schemas and data objects, such as tables and views, that organizes and manages user data and metadata for querying and data processing. For more information about databases see [Create a Database](/Guides/getting-started/get-started-sql.html#create-a-database) in the [Get started using SQL](/Guides/getting-started/get-started-sql.html) guide. Under database are the following levels:
    
    - **Schema** - A collection of system views containing metadata about objects in the current database, including tables, columns, indexes, and other database components. For more information, see [Information Schema Views](/sql_reference/information-schema/). Under schema are the following levels:
      
      - **External table** - Tables that store metadata objects that reference files stored in an Amazon S3 bucket, rather than actual data. For more information, see [Work with external tables](/Guides/loading-data/working-with-external-tables.html).
      - **Managed table** - A Firebolt-managed internal structured data object within a database that stores rows and columns of data. Firebolt’s managed tables have built-in optimizations for fast query processing times. For more information, see [Tables](/Overview/indexes/using-indexes.html#tables).
      - **View** - A virtual table that represents the result of a stored query, including both user-defined views and [information schema views](/sql_reference/information-schema/), which provide metadata about database objects like tables, columns, and indexes.
      - **Index** - A database structure that optimizes data retrieval by organizing specific columns, improving query performance and enabling efficient filtering, sorting, and joining of datasets. For more information about indexes, see [Data modeling](/Overview/indexes/using-indexes.html).
  - **Engine** - A compute resource that processes queries and manages data operations and can be scaled independently and turned on or off as needed. For more information, see [Firebolt Engines](/Overview/engine-fundamentals.html).

When you [register for the first time](/Guides/managing-your-organization/creating-an-organization.html), Firebolt sets up an organization for you. During registration, you’ll set up your first account, with one user. The first user that is added is the account administrator, as shown in the following diagram:

![When you first register, Firebolt sets up an organization with one account and user that has account administrator privileges.](../../assets/images/manage-organization-first-registration.png)

Then, you can add resources and users to this account. The following apply:

- You can have multiple users within an account.
- A user should be associated with either a login email for personal access or to a service account, for programmatic access.
- You can have multiple accounts in an organization.
- Each account name within an organization must be unique.
- You can’t have one account in multiple AWS regions.
- You can add resources such as databases and engines to the account.

**Example**

In the following example structure, an organization has an account set up for their marketing department and for managers in two different AWS regions:

![An account can span only one AWS region.](../../assets/images/manage-organization-accounts-regions.png)

In the previous diagram, the organization has three separate accounts:

- A `marketing_account`, which has access to resources associated with marketing tasks, and two different users. The `user_1` user is associated with a `login_1` linked to an email account.
- A `manager_account_region_1`, which has access to resources associated with a manager account in one AWS region, and one `user_3` that is associated with login\_1 linked to the same email account as `user_1` in marketing\_account.
- A `manager_account_region_2`, which has access to resources in a different region than `manager_account_region_1`, and one user also associated with `login_1`.

The manager of the marketing department, `user_1`, is associated with `login_1`, which is associated with both the marketing account and both manager accounts. These accounts have access to a different set of resources and permissions. The users `user_1`, `user_3`, and `user_4` are all the same person because they have the same login and email. The manager also manages projects across AWS regions and must access those resources in a **different** account. Another employee, `user_2`, works in the marketing department and has access to only the marketing resources designated to `marketing_account` using permissions defined by his role.

## [](#accounts)Accounts

An account in Firebolt is an object within an organization that encapsulates resources for storing, querying, and managing data. Accounts provide:

- **Access control:** Firebolt implements role-based access control (RBAC). Every object in the Firebolt object model is a securable and it comes with a set of privileges. Privileges allow administrators to control functionality Firebolt users can exercise when logged in.
- **Data modeling:** Using objects including databases, tables, views, and indexes, developers and architects can design their data warehouses and describe various business entities without compromising to deliver on ever-demanding performance needs.
- **Cost control:** With engines, system administrators can deploy engines that fit the need while achieving desired price-performance characteristics. Engines can scale vertically up and down, and horizontally out and in to meet business needs while allowing granular cost control.
- **Workload management:** Firebolt offers full workload isolation for computations, data and metadata. Firebolt users can deploy separate engines to support heterogeneous workloads, while having access to the same data. Firebolt supports a variety of workloads, including data-intensive applications requiring instant data access, complex business-critical dashboards needing timely updates, and intricate Extract-Load-Transform (ELT) processes for data ingestion.

Each account in Firebolt exists in a single AWS region, and can have engines and databases associated with it. Initially after registration, an account contains no resources, and only one user that has an account administrator role. An account can contain many users, as shown in the following diagram:

![In Firebolt, an account can contain many users, an engine and database.](../../assets/images/manage-organization-accounts-contents.png)

### [](#users)Users

A user must be associated with a role, which grants them permission to access resources. These users can be associated with different roles within a single account. Each user must be associated with either a login for personal access or a service account for programmatic access, as shown in the following diagram:

![A user must be associated with either a login or a service account.](../../assets/images/user_login_service-account.png)

A login consists of an **email address**. This login uniquely identifies the user.

**Example**

In the following example account structure, `user_1` has a manager role that grants access to engines and databases associated with human resources tasks, as well as a marketing role that grants them access to everything that their employee has access to. A marketing employee, `user_2`, has read-only access to the tables in the database in `marketing_account`, but they cannot insert new entries or delete entries from a table.

![A user can have multiple roles in an account.](../../assets/images/manage-organization-marketing-accounts-example.png)

### [](#roles)Roles

In Firebolt, each user is associated with either a **login**, which is an email address, or a **service account**. Each user must also have a role, as shown in the following diagram:

![A user must be associated with either a login or a service account.](../../assets/images/user_login_service-account.png)

The role grants the user permission to access resources inside the account that they are associated with. A user can have several roles associated with them at the same time. Firebolt has built-in roles with defined permissions. You can also define a [custom role](/Overview/Security/Role-Based%20Access%20Control/role-management/custom-roles.html) that grants permissions specific to your use case.

#### [](#firebolt-built-in-roles)Firebolt built-in roles

Firebolt has the following built-in roles with associated permissions for objects including databases, engines, users, network policies, and accounts:

- [Organizations and accounts](#organizations-and-accounts)
  
  - [Organizations](#organizations)
  - [Accounts](#accounts)
    
    - [Users](#users)
    - [Roles](#roles)
      
      - [Firebolt built-in roles](#firebolt-built-in-roles)
        
        - [Public role](#public-role)
        - [System administrative role](#system-administrative-role)
        - [Account administrative role](#account-administrative-role)
        - [Organizational administrative role](#organizational-administrative-role)
  - [Additional resources](#additional-resources)
    
    - [Billing](#billing)
    - [Security](#security)
    - [Available regions](#available-regions)

##### [](#public-role)Public role

A public role is associated with a user that can:

- Use a database.
- Use a public schema.
- Create a public table.
- Create a public view.
- Create a public index.
- Create a public external table.

A public role has the **lowest** access privileges of all roles in Firebolt, as shown in the following diagram:

![The public role has permission for schema and its underlying components.](../../assets/images/manage-organization-public-role.png)

##### [](#system-administrative-role)System administrative role

A system administrative role has privileges to manage databases, engines, schemas, and objects within those schemas. A system administrator can:

- Create a database in an account.
- Create an engine in an account.
- Monitor engine use.
- Has all privileges for:
  
  - Any database and its properties.
  - Any engine and its properties.
  - Any schema.
  - Any view.
  - Any external table.

The previous system administrative privileges are shown in the following diagram:

![The system admin role has privileges in database, engine, and schema plus children.](../../assets/images/manage-organization-sys-admin-role.png)

##### [](#account-administrative-role)Account administrative role

An account administrative role includes all privileges associated with system administrators and can also manage accounts and users. An account administrator has:

- All system administrator privileges.
- All privileges for an account.
- The ability to meter and monitor account use.
- The ability to cancel a query on any engine in an account.

The previous account administrative privileges are shown in the following diagram:

![An account admin has privileges over the account and all its children.](../../assets/images/manage-organization-account-admin-role.png)

##### [](#organizational-administrative-role)Organizational administrative role

An organizational administrative role has all privileges associated with system administrators and can also manage accounts and users. An organizational administrator has:

- All privileges for an organization.
- All privileges for any account in an organization.
- All privileges for any login in an organization.
- All privileges for any service account in an organization.
- All privileges for any network policies in an organization.
- The ability to monitor any usage in the organization.
- The ability to set any organization-related property.

An organizational administrative role has the **highest** access privileges of all roles in Firebolt, as shown in the following diagram:

![An org admin has privileges over the entire organization and all of its children.](../../assets/images/manage-organization-org-admin-role.png)

- **Global authentication method:** Firebolt handles user authentication and access control at the organization level. A login (represented by an email) is created for each user accessing Firebolt.
- **Programmatic access:** [Service accounts](/Guides/managing-your-organization/service-accounts.html) enable programmatic access to Firebolt.
- **Network policy enforcement:** [Network policies](/Guides/security/network-policies.html) provide fine-grain control of IP ranges that are allowed or blocked from accessing an organization.

## [](#additional-resources)Additional resources

### [](#billing)Billing

Firebolt provides billing at the organization level, but gives you billing observability at both organization and account levels. This allows:

- **Organization-level governance:** Monitor and analyze the overall billing for all accounts to gain insights into the organization’s cost distribution and resource utilization at the organization level.
- **Account-level observability:** Delve into detailed billing information specific to each account, allowing you to track individual accounts’ usage, costs, storage, and compute consumption patterns.

Firebolt bills are based on the consumption of resources within each account in your organization. This includes the total amount of data stored and engine usage. Learn how to [manage billing](/Guides/managing-your-organization/billing.html).

### [](#security)Security

Learn about authentication methods, role-based access control, network policies, and object ownership in [Configure security](/Guides/security/).

### [](#available-regions)Available regions

View the [AWS regions](/Reference/available-regions.html) where you can use Firebolt.