# [](#manage-accounts)Manage accounts

Your organization comes prepared with one account for your convenience. You can add more accounts, edit existing accounts, or delete accounts using SQL or in the UI.

To view all accounts, click **Configure** to open the configure space, then choose **Accounts** from the menu, or query the [information\_schema.accounts](/sql_reference/information-schema/accounts.html) view.

## [](#create-a-new-account)Create a new account

Creating an account requires the org\_admin role.

### [](#sql)SQL

To create an account using SQL, use the [CREATE ACCOUNT](/sql_reference/commands/data-definition/create-account.html) statement. For example:

```
CREATE ACCOUNT dev WITH REGION = 'us-east-1';
```

### [](#ui)UI

To create an account via the UI:

![Configure > Accounts](/assets/images/accountspage.png)

1. Click **Configure** to open the configure space, then choose **Accounts** from the menu.
2. From the Accounts management page, choose **Create Account**. Type a name for the account and choose a region. You won’t be able to change the region for this account later, so choose carefully.
3. Choose **Create**.

![Create account](../../assets/images/createaccount.png)

Then you will see your new account on the **Accounts management** page.

There can be up to 20 accounts per organization and you can use `CREATE ACCOUNT` 25 times. If you have a need for additional account creations beyond this limit, contact [Firebolt Support](https://docs.firebolt.io/godocs/Reference/help-menu.html) for assistance. Our team can provide guidance and, if appropriate, adjust your account settings to accommodate your needs.

## [](#edit-an-existing-account)Edit an existing account

Editing an account requires the account\_admin or org\_admin role.

### [](#sql-1)SQL

To edit an existing account using SQL, use the [ALTER ACCOUNT](/sql_reference/commands/data-definition/alter-account.html) statement. For example:

```
ALTER ACCOUNT dev RENAME TO staging;
```

### [](#ui-1)UI

To edit an account via the UI:

1. Click **Configure** to open the configure space, then choose **Accounts** from the menu.
2. Search for the relevant account using the top search filters or by scrolling through the accounts list. Hover over the right-most column to make the account menu appear then choose **Edit account**. Edit the name of the account.
3. Choose **Save**.

![Edit account](../../assets/images/editaccount.png)

## [](#delete-an-existing-account)Delete an existing account

Deleting an account requires the account\_admin or org\_admin role.

### [](#sql-2)SQL

To delete an existing account using SQL, use the [DROP ACCOUNT](/sql_reference/commands/data-definition/drop-account.html) statement. For example:

```
DROP ACCOUNT dev;
```

### [](#ui-2)UI

To delete an account via the UI:

1. Click **Configure** to open the configure space, then choose **Accounts** from the menu.
2. Search for the relevant account using the top search filters or by scrolling through the accounts list. Hover over the right-most column to make the account menu appear then choose **Delete account**. If your account is not empty (for example, if it contains other objects such as users/databases/engines/etc.), you will need to confirm that you will also delete the sub-objects by selecting **Delete account sub-objects permanently**.
3. Choose **Confirm**.

![Delete account](../../assets/images/deleteaccount.png)

The account will be removed from the **Accounts management** page.

## [](#switch-accounts)Switch accounts

To switch the account you are using:

### [](#ui-3)UI

Click on your login button - the current account will be marked. Choose an account you would like to switch to.

![Switch account](../../assets/images/switch_account.png)