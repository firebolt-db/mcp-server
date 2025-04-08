# [](#role-based-access-control-rbac)Role-Based Access Control (RBAC)

Firebolt uses Role-Based Access Control (RBAC) to manage permissions and ensure that users and roles have only the necessary access to perform operations within the system. RBAC follows the [principle of least privilege](https://en.wikipedia.org/wiki/Principle_of_least_privilege), where access is restricted to the minimum required for tasks.

Permissions in Firebolt are managed through roles, which control access to databases, schemas, tables, engines, and other objects. Permissions propagate from higher-level objects to their related objects, simplifying access management.

With RBAC you can:

- Assign roles to users and other roles to streamline permissions.
- Grant or revoke access at **global**, **regional**, or **object-specific** levels.
- Control operations across your Firebolt environment, such as managing data, creating resources, or executing queries.

## [](#firebolts-hierarchical-object-model-and-rbac-structure)Firebolt’s hierarchical object model and RBAC structure

Firebolt uses an object model to organize resources in a way that complements how organizations manage their data warehouses. This model enforces a one-to-many structure where:

- An object can encompass multiple related objects beneath it.
- Each related object is associated with a single higher-level object and cannot be shared across multiple higher-level objects.
- Permissions flow from higher-level objects to related objects. For example, granting a role usage on a database also provides access to all schemas and tables within that database.

Objects in the Firebolt object model are securable and come with a set of permissions, enabling administrators to control what identities they have access to when accessing their Firebolt cluster.

For more information about the organizational and account structure in Firebolt’s object model , see [Organization and accounts](/Overview/organizations-accounts.html).

## [](#key-object-types)Key object types

Firebolt divides objects into **global** and **regional** types, depending on their scope and management level.

- **Global objects**: Managed globally at the [organization](/Overview/organizations-accounts.html#organizations) level, they can contain objects that are deployed and grouped regionally, including the following:
  
  - Network Policies
  - Logins
  - Service Accounts
  - Accounts
- **Regional objects**: Tied to specific regions grouped under an [account](/Overview/organizations-accounts.html#accounts), they can include the following:
  
  - Users
  - Roles
  - Databases
    
    - Schemas
      
      - External Tables
      - Managed Tables
      - Views
      - Indexes
  - Engines

Firebolt provides the [organization\_admin](/Overview/organizations-accounts.html#organizational-administrative-role) role to manage organizational resources. While granular RBAC is currently only available at the account level, Firebolt plans to include making RBAC available at the organizational level in a future release.

For more information about Firebolt’s RBAC model &amp; how to administer your Firebolt cluster, access the sections below:

* * *

- [Database permissions](/Overview/Security/Role-Based%20Access%20Control/database-permissions/)
- [Account Permissions](/Overview/Security/Role-Based%20Access%20Control/account-permissions.html)
- [Engine Permissions](/Overview/Security/Role-Based%20Access%20Control/engine-permissions.html)
- [Ownership](/Overview/Security/Role-Based%20Access%20Control/ownership.html)
- [Check Assigned Privileges](/Overview/Security/Role-Based%20Access%20Control/check-privileges.html)
- [Role Management](/Overview/Security/Role-Based%20Access%20Control/role-management/)