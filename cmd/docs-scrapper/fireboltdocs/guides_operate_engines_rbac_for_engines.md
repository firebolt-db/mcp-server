# [](#governing-engines)Governing Engines

Use [Role Based Access Control](/Guides/security/rbac.html) (RBAC) to granularly control which users within an account can create new engines, use, operate, monitor and modify existing engines. Accordingly, Firebolt provides CREATE, USAGE, OPERATE, MONITOR and MODIFY permissions to control these actions. You can use RBAC to control whether a user has permissions to perform these actions for specific engines or for all engines in a given account. Note that permissions for CREATE ENGINE can only be granted at the account level.

Follow the below steps to control what permissions a user has for a given engine or for any engine within an account:

- Create a new role
- Grant permissions to the role
- Assign role to a user

**Example 1:** We want to provide a user kate with permissions to create and operate engines

```
CREATE ROLE prodAdminRole;

GRANT CREATE ENGINE ON ACCOUNT myAccount IN ORGANIZATION myOrg TO prodAminRole; 

GRANT OPERATE ON ENGINE myEngine IN ACCOUNT myAccount TO prodAdminRole; 

GRANT ROLE prodAdminRole TO USER kate;  
```

**Example 2:** We want to provide a user kate with permissions to only use and operate engines

```
CREATE ROLE prodAdminRole;

GRANT USAGE ENGINE ON myEngine IN ACCOUNT myAccount TO prodAminRole; 

GRANT OPERATE ON ENGINE myEngine IN ACCOUNT myAccount TO prodAdminRole; 

GRANT ROLE prodAdminRole TO USER kate;  
```

**Example 3:** We want to provide a user kate with permissions to use, operate and monitor engine metrics

```
CREATE ROLE prodAdminRole;

GRANT USAGE ENGINE ON myEngine IN ACCOUNT myAccount TO prodAminRole;

GRANT MONITOR USAGE ON ENGINE myEngine IN ACCOUNT myAccount TO prodAminRole; 

GRANT OPERATE ON ENGINE myEngine IN ACCOUNT myAccount TO prodAdminRole; 

GRANT ROLE prodAdminRole TO USER kate;  
```

**Example 4:** We want to provide a user kate with permissions to create and modify engines

```
CREATE ROLE prodAdminRole;

GRANT CREATE ENGINE ON ACCOUNT myAccount IN ORGANIZATION myOrg TO prodAminRole; 

GRANT MODIFY ON ENGINE myEngine IN ACCOUNT myAccount TO prodAdminRole; 

GRANT ROLE prodAdminRole TO USER kate;  
```