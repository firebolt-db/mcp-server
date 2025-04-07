# [](#information-schema-for-users)Information schema for users

You can use the `information_schema.users` view to return information about users. You can use a `SELECT` query to return information about each user, as shown in the example below.

To view information about users, you must have [user privilege](/Overview/Security/Role-Based%20Access%20Control/user-permissions.html#user-permissions) or ownership of the user object.

```
SELECT
  *
FROM
  information_schema.users;
```

For more information, see [Managing users](/Guides/managing-your-organization/managing-users.html).

## [](#columns-in-information_schemausers)Columns in information\_schema.users

Each row has the following columns with information about each user.

Column Name Data Type Description user\_name TEXT The name of the user. login\_name TEXT The name of the login linked to the user. Empty if the login is linked to a service account. service\_account\_name TEXT The name of the service account linked to the user. Empty if the login is linked to a service account. account\_name TEXT The name of the account. organization\_name TEXT The name of the organization. default\_database TEXT The default database set for the user. default\_engine TEXT The default engine set for the user. created TIMESTAMPTZ Time the user was created. user\_owner TEXT The name of the user who created the user. last\_altered TIMESTAMPTZ Time the user was last altered. last\_altered\_by TEXT The name of the last user to edit the user.