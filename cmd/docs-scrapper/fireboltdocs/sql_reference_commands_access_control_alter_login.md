# [](#alter-login)ALTER LOGIN

Updates the configuration of the specified login.

For more information, see [Managing logins](/Guides/managing-your-organization/managing-logins.html).

## [](#syntax)Syntax

```
ALTER LOGIN <login_name> SET 
      [ IS_PASSWORD_ENABLED = { TRUE | FALSE } ]
      [ IS_MFA_ENABLED = { TRUE | FALSE } ]
      [ NETWORK_POLICY = <network_policy_name> | DEFAULT ]
      [ IS_ORGANIZATION_ADMIN = { TRUE | FALSE } ]
      [ IS_ENABLED = { TRUE | FALSE } ]
      [ FIRST_NAME = <first_name> ]
      [ LAST_NAME = <last_name> ] 
```

or

```
ALTER LOGIN <login_name> RENAME TO <new_login_name>;
```

A login that was provisioned via SSO cannot be renamed.

## [](#parameters)Parameters

Parameter Description `<login_name>` The name of the login in the form of an email address. The login must be unique within the organization. `IS_PASSWORD_ENABLED` A `BOOLEAN` value specifying if login with password is enabled. By default this is `TRUE` and passwords can be used to log in. `IS_MFA_ENABLED` A `BOOLEAN` value specifying if the login has multi-factor authentication (MFA) enabled. By default this value is `FALSE`. If set to `TRUE`, an enrollment email will be sent to the `<login_name>`. `<network_policy_name>` An optional parameter to define the network policy to link to the created login. Specifying `DEFAULT` will detach any linked network policy. `IS_ORGANIZATION_ADMIN` A `BOOLEAN` value specifying if the login is an organization admin. By default this value is `FALSE`. `IS_ENABLED` A `BOOLEAN` value specifying whether authentication with this login should be possible. Disable login if you want to prevent access to the system without dropping it. `<first_name>`, `<last_name>` The first and last name of the user to use the login. If the parameter is included, these values cannot be empty. `<new_login_name>` The new name of the login in the form of an email address. The login must be unique within the organization. Note: if the login was provisioned via SSO, it cannot be renamed.

## [](#example)Example

The following command will rename the “alexs@acme.com” login to “alexspotter@acme.com”.

```
ALTER LOGIN "alexs@acme.com" RENAME TO "alexspotter@acme.com";
```

This command will link the network policy “my\_network\_policy” to the “alexs@acme.com” login.

```
ALTER LOGIN "alexs@acme.com" SET NETWORK_POLICY = "my_network_policy";
```