# [](#manage-logins)Manage logins

Logins are managed at the organization level and are used for authentication. Logins are a combination of a login name (email), first name, last name, and password, unless you’ve configured [Single Sign-On (SSO)](../security/sso/). Moreover, logins can be configured with advanced authentication properties such as [MFA](/Guides/security/enabling-mfa.html) and [network policies](/Guides/security/network-policies.html). Logins are linked to users at the account level, so that roles may be managed separately per account. A user must be linked to either a login or a service account for programmatic use to gain access to Firebolt. You can add, edit or delete logins using SQL or in the UI.

To view all logins, click **Configure** to open the configure space, then choose **Logins** from the menu, or query the [information\_schema.logins](/sql_reference/information-schema/logins.html) view.

Managing logins requires the org\_admin role.

## [](#create-a-new-login)Create a new login

### [](#sql)SQL

To create a login using SQL, use the [CREATE LOGIN](/sql_reference/commands/access-control/create-login.html) statement. For example:

```
CREATE LOGIN "alexs@acme.com" WITH FIRST_NAME = 'Alex' LAST_NAME = 'Summers';
```

### [](#ui)UI

To create a login via the UI:

1. Click **Configure** to open the configure space, then choose **Logins** from the menu:

<!--THE END-->

![Configure > Logins](/assets/images/loginspage.png)

1. From the Logins management page, choose **Create Login**.
2. Enter the following details:
   
   - First name: specifies the first name of the user for the login.
   - Last name: specifies the last name of the user for the login.
   - Login name: specifies the login in the form of an email address. This must be unique within your organization.
3. Optionally, you can:
   
   - Associate a [network policy](/Guides/security/network-policies.html) with the login by choosing a network policy name under the **Network policy attached** field.
   - Enable password login, which specifies if the login can authenticate Firebolt using a password.
   - Enable multi-factor authentication (MFA). Read more about how to configure MFA [here](/Guides/security/enabling-mfa.html).
   - Set the login as **organisation admin**, which enables fully managing the organization.

## [](#edit-an-existing-login)Edit an existing login

### [](#sql-1)SQL

To edit an existing login using SQL, use the [ALTER LOGIN](/sql_reference/commands/access-control/alter-login.html) statement. For example:

```
ALTER LOGIN "alexs@acme.com" SET NETWORK_POLICY = my_network_policy
```

### [](#ui-1)UI

To edit a login via the UI:

1. Click **Configure** to open the configure space, then choose **Logins** from the menu.
2. Search for the relevant login using the top search filters, or by scrolling through the list of logins. Hover over the right-most column to make the login menu appear, then choose **Edit login details**. Edit the desired fields and choose **Save**.

Login name can not be changed for logins that were provisioned via SSO.

![Edit login](../../assets/images/editlogin.png)

## [](#deleting-an-existing-login)Deleting an existing login

### [](#sql-2)SQL

To delete an existing login using SQL, use the [DROP LOGIN](/sql_reference/commands/access-control/drop-login.html) statement. For example:

```
DROP LOGIN "alexs@acme.com";
```

### [](#ui-2)UI

To delete a login via the UI:

1. Click **Configure** to open the configure space, then choose **Logins** from the menu.
2. Search for the relevant login using the top search filters, or by scrolling through the logins list. Hover over the right-most column to make the login menu appear, then choose **Delete login**.

If the login is linked to any users, deletion will not be permitted. The login must be unlinked from all users before deletion.