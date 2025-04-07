# [](#security)Security

Firebolt employs a layered security strategy to deliver a secure and trusted cloud data warehouse tailored to modern enterprise needs. By integrating advanced security features and industry best practices, Firebolt ensures that its security practices are compliant with security standards. Your data is protected, remains secure, and is accessible only to authorized individuals. Access to objects and resources is managed through accounts, establishing clear boundaries to isolate and secure data.

Firebolt’s layered security model has the following key areas:

- [Security](#security)
  
  - [Network security](#network-security)
  - [Identity management](#identity-management)
    
    - [Single sign-On (SSO)](#single-sign-on-sso)
    - [Multi-factor authentication (MFA)](#multi-factor-authentication-mfa)
  - [Access control](#access-control)
    
    - [Roles](#roles)
    - [Users](#users)
    - [Objects and permissions](#objects-and-permissions)
  - [Data protection](#data-protection)
    
    - [Data at rest](#data-at-rest)
    - [Data in motion](#data-in-motion)
    - [Secure communication protocols](#secure-communication-protocols)
    - [AWS PrivateLink](#aws-privatelink)
    - [HIPAA compliance](#hipaa-compliance)

## [](#network-security)Network security

Firebolt ensures secure data transmission by implementing end-to-end encryption with Transport Layer Security (TLS) version 1.2, protecting data as it moves between end users and the cloud service.

Firebolt supports the creation of custom [network policies](/Guides/security/network-policies.html), adding an extra layer of security to your applications. This functionality allows administrators to exercise fine-grained control over which IP ranges can access Firebolt.

In Firebolt, network policies contain `allowed_IP_list` and `blocked_IP_list` properties that capture definition of IP address ranges. Each property is a list that can contain one or more IP ranges.

Firebolt supports individual and programmatic access through the following:

- [Login](/Guides/managing-your-organization/managing-logins.html) - A login object represents an individual user, identified by an email address, who will authenticate by verifying their identity to access Firebolt.
- [Service account](/Guides/managing-your-organization/service-accounts.html) - A service account object is used to represent a machine or application that will authenticate and interact with Firebolt without human intervention.

**Example**

The following code example creates a network policy `my_network_policy` with a description that allows only two IP addresses:

```
CREATE NETWORK POLICY IF NOT EXISTS my_network_policy WITH ALLOWED_IP_LIST = (‘4.5.6.1’, ‘2.4.5.1’) 
DESCRIPTION = 'my new network policy'
```

A network policy can be attached to [an organization](/Guides/managing-your-organization/creating-an-organization.html), individual logins, and service accounts.

**Example**

The following code example applies `my_network_policy` to `my_organization` and to the login associated with `kate@acme.com`, which restricts access according to the rules of that policy:

```
ALTER ORGANIZATION my_organization SET NETWORK_POLICY = my_network_policy;
ALTER LOGIN 'kate@acme.com' SET NETWORK_POLICY = my_network_policy;
```

For more information, see [network policies](/Guides/security/network-policies.html).

## [](#identity-management)Identity management

Identity management is a multi-step verification process designed to ensure that only authorized individuals, services, and applications can access organizational resources. Identity management involves both identifying users and verifying their identity through authentication.

Firebolt uses [Auth0](https://auth0.com/) as its identity provider for managing customer registration. All authentication data stored in Auth0 is protected with industry-standard cryptography. Authentication information is securely exchanged using the [SAML 2.0 protocol](https://auth0.com/intro-to-iam/what-is-saml).

Firebolt provides the SSO and MFA authentication methods.

### [](#single-sign-on-sso)Single sign-On (SSO)

[Single Sign-On (SSO)](/Guides/security/sso/) is an authentication method that allows users to access multiple applications or services using a single set of login credentials, simplifying the authentication process and improving security through centralized identity management. Firebolt uses SSO to simplify and streamline implementation of secure access to its platform, enhancing the overall security posture and protecting against unauthorized access and data breaches. SSO configuration is accessible to users with the `org_account` built-in role.

### [](#multi-factor-authentication-mfa)Multi-factor authentication (MFA)

[MFA](/Guides/security/enabling-mfa.html) strengthens security by requiring users to provide multiple forms of authentication to access their accounts. Many industries have compliance and regulatory standards that require the use of MFA for securing certain types of data and systems. Firebolt fully supports these standards by offering MFA configuration and implementation directly linked to the login object, ensuring secure and compliant access control.

**Example**

The following code example creates a login for a user, enables password-based authentication, and requires MFA to authenticate:

```
CREATE LOGIN "kate@acme.com" WITH
FIRST_NAME = 'Kate'
LAST_NAME = 'Peterson'
IS_PASSWORD_ENABLED = TRUE
IS_MFA_ENABLED = TRUE;
```

## [](#access-control)Access control

Access control ensures that users have the necessary and appropriate permissions to engage with Firebolt’s system or resources. Firebolt implements [role-based access control](/Guides/security/rbac.html) (RBAC) to manage permissions.

The RBAC model is centered around the following principles:

- All objects can be secured.
- Every supported statement requires explicit permission, preventing unauthorized actions.
- The RBAC model is composable, which means that a user’s total permissions are the result of combining all the roles assigned to them, giving them the collective access granted by each role.
- Roles are hierarchical and allow permissions to be inherited through role relationships.

The RBAC model contains the following:

- [Roles](#roles)
- [Users](#users)
- [Objects and permissions](#objects-and-permissions).

### [](#roles)Roles

A role is a set of permissions assigned to a user or group that defines what actions they are authorized to perform and what resources they can access within Firebolt. Firebolt has the following types of roles:

1\) **Built-in roles** have a set of pre-defined permissions and custom user-defined roles that can allow for more specific use cases. You can use [GRANT](/sql_reference/commands/access-control/grant.html) and [REVOKE](/sql_reference/commands/access-control/revoke.html) statements to modify permissions for custom roles. Built-in roles become available as soon as a new organization is created, and the first account is set up.

2\) **User-defined roles** are custom roles that administrators can create to grant a specific set of permissions.

3\) **System-defined roles** align with common user personas and responsibilities including `public`, which is granted to each new user by default, a `system_admin` role, and an `account_admin` role. For more information about these roles, see [System-defined roles](/Overview/Security/Role-Based%20Access%20Control/role-management/system-roles.html).

You can create a role by using either the **Firebolt Workspace** or using the [CREATE ROLE](/sql_reference/commands/access-control/create-role.html) SQL statement.

**Example**

The following example creates a new role `sales`, which can later be assigned specific permissions and granted to users to inherit those permissions:

```
CREATE ROLE sales;
```

### [](#users)Users

Users are linked to either a login or service account in order to gain access to Firebolt. They can be created using the [CREATE USER](/sql_reference/commands/access-control/create-user.html) statement in SQL or through the **Firebolt Workspace**.

**Example**

The following code creates a new users `kate` and `bob`:

```
CREATE USER kate;
CREATE USER bob;
```

**Example**

The following code example grants the permissions associated with the `sales` role to `kate` and revokes it from `bob`:

```
GRANT ROLE sales TO USER kate;
REVOKE ROLE sales FROM USER bob; 
```

### [](#objects-and-permissions)Objects and permissions

Permissions in Firebolt define the actions or operations that can be performed, such as managing databases and engines, running queries, or accessing data. Each instance of a securable object, or an object that can be protected by access controls, has specific permissions that are associated with it, controlling what users can do with it. Examples of securable objects include databases, tables, and engines. If there are multiple instances of an engine object, each instance has its own set of predefined permissions. For a full list of available permissions, see [role-based access control](/Overview/Security/Role-Based%20Access%20Control/)

Any permission that Firebolt supports can be [granted](/sql_reference/commands/access-control/grant.html) or [revoked](/sql_reference/commands/access-control/revoke.html) to or from roles.

Privileges can be granted or revoked only for roles, not directly for users. Once a role has the necessary permissions, it can then be assigned to users, allowing them to inherit those privileges.

**Example**

The following code example grants the `sales` role permission to use the `sales_db` database, allows the role to access any database within `dev_account`, and revokes the ability of the `sales` role to start or stop the `sales_eng` engine:

```
GRANT USAGE ON DATABASE sales_db TO sales;    -- grants the ability to use sales_db database to the sales role
GRANT USAGE ANY DATABASE ON ACCOUNT dev_account TO sales;    -- grants the ability to use any database in dev_account to the sales role
REVOKE OPERATE ON ENGINE sales_eng FROM sales;   -- revokes the ability to START and STOP the sales_eng engine from sales role
```

## [](#data-protection)Data protection

Firebolt is firmly committed to data security, privacy, and compliance by ensuring that all data it manages is properly safeguarded and protected through strict encryption standards for data both in motion and at rest. The following security functionality is automatically available to customers:

### [](#data-at-rest)Data at rest

By default, all data at rest is encrypted and stored using Amazon Simple Storage Service (S3). All new objects are automatically encrypted using either Amazon S3-managed keys or AWS Key Management Service (KMS) keys, which are securely managed through AWS KMS.

### [](#data-in-motion)Data in motion

Firebolt automatically encrypts sensitive data being transmitted between service components, ensuring that it remains secure as it moves across networks and cannot be intercepted by unauthorized parties.

### [](#secure-communication-protocols)Secure communication protocols

Firebolt uses secure communication protocols, such as Transport Layer Security (TLS), to provide an additional layer of protection against man-in-the-middle attacks. These attacks occur when an unauthorized party intercepts or alters data as it is transmitted between two points. By encrypting data and ensuring secure connections, Firebolt prevents unauthorized access or tampering during data transmission, safeguarding sensitive information as it moves between systems.

### [](#aws-privatelink)AWS PrivateLink

AWS PrivateLink enables secure connectivity between your VPC and Firebolt without routing traffic over the public internet. Firebolt supports AWS PrivateLink to provide a private, one-way connection from your VPC to the Firebolt private API, ensuring data confidentiality, integrity, and availability during transit.

Users with an [account administrator](/Overview/organizations-accounts.html#account-administrative-role) or [organizational administrative](/Overview/organizations-accounts.html#organizational-administrative-role) role can request access to PrivateLink. For more information, see [Request PrivateLink access](/Guides/security/privatelink.html#request-aws-privatelink-access).

### [](#hipaa-compliance)HIPAA compliance

HIPAA compliance consists of federal regulations designed to safeguard the privacy and security of patient health information. Firebolt supports HIPAA compliance to ensure the confidentiality, integrity, and availability of electronic protected health information (ePHI) stored within its platform.

To modify the state of HIPAA compliance for your account, contact the [Firebolt support team](mailto:support@firebolt.io).

* * *

- [Role-Based Access Control](/Overview/Security/Role-Based%20Access%20Control/)