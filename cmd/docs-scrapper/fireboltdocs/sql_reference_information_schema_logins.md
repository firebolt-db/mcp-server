# [](#information-schema-for-logins)Information schema for logins

You can use the `information_schema.logins` view to return information about logins.

You can use a `SELECT` query to return information about each login, as shown in the example below.

```
SELECT
  *
FROM
  information_schema.logins;
```

For more information, see [Managing logins](/Guides/managing-your-organization/managing-logins.html).

## [](#columns-in-information_schemalogins)Columns in information\_schema.logins

Each row has the following columns with information about each login.

Column Name Data Type Description login\_name TEXT Name of the login (email address). first\_name TEXT First name of the user linked to the login. last\_name TEXT Last name of the user linked to the login. organization\_name TEXT Name of the organization. network\_policy\_name TEXT Name of the network policy associated with the login. is\_mfa\_enabled BOOLEAN Specifies if the login has multi-factor authentication enabled. is\_sso\_provisioned BOOLEAN Specifies if the login was provisioned with an identity provider defined in the organization’s SSO configuration. is\_password\_enabled BOOLEAN Specifies if log in with password is enabled. is\_organization\_admin BOOLEAN Specifies if the login is an organization admin. is\_enabled BOOLEAN Specifies if the login is allowed to authenticate. created TIMESTAMPTZ Time the login was created. login\_owner TEXT Name of the login who created the login. If the login was created by a service account, the service account name appears instead. last\_altered TIMESTAMPTZ Time the login was last altered. last\_altered\_by TEXT Name of the login who edited the login. If the login was edited by a service account, the service account name appears instead.