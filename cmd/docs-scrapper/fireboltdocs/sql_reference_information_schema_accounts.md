# [](#information-schema-for-accounts)Information schema for accounts

You can use the `information_schema.accounts` view to return information about accounts.

You can use a `SELECT` query to return information about each account, as shown in the example below.

```
SELECT
  *
FROM
  information_schema.accounts;
```

Read more about managing accounts [here](/Guides/managing-your-organization/managing-accounts.html).

## [](#columns-in-information_schemaaccounts)Columns in information\_schema.accounts

Each row has the following columns with information about the account.

Column Name Data Type Description account\_name TEXT The name of the account. organization\_name TEXT The name of the organization to which the account belongs. region TEXT The region in which the account can be used. url TEXT The account login page URL. account\_id TEXT The unique account ID. trust\_policy\_role TEXT Role provided by Firebolt to enable access to customer S3 buckets created TIMESTAMP Time (UTC) that the account was created. account\_owner TEXT The name of the login that created the account. last\_altered TIMESTAMP Time (UTC) that the user was last edited. last\_altered\_by TEXT Name of the last user to edit the role.