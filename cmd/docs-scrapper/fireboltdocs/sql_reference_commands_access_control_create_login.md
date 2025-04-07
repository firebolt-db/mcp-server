# [](#create-login)CREATE LOGIN

Creates a new login.

For more information, see [Managing logins](/Guides/managing-your-organization/managing-logins.html).

## [](#syntax)Syntax

```
CREATE LOGIN [ IF NOT EXISTS ] <login_name> 
	   WITH [ IS_PASSWORD_ENABLED = { TRUE | FALSE }]
            [ IS_MFA_ENABLED = { TRUE | FALSE }]
            [ NETWORK_POLICY = <network_policy_name>]
            [ IS_ORGANIZATION_ADMIN = { TRUE | FALSE }]
	        FIRST_NAME = <first_name>,
	        LAST_NAME = <last_name> 
```

## [](#parameters)Parameters

Parameter Description `<login_name>` The name of the login in the form of an email address. The login must be unique within the organization. `IS_PASSWORD_ENABLED` A `BOOLEAN` value specifying if login with password is enabled. By default this is `TRUE` and passwords can be used to log in. `IS_MFA_ENABLED` A `BOOLEAN` value specifying if the login has multi-factor authentication (MFA) enabled. By default this value is `FALSE`. If set to `TRUE`, an enrollment email will be sent to the `<login_name>`. `<network_policy_name>` An optional parameter to define the network policy to link to the created login. `IS_ORGANIZATION_ADMIN` A `BOOLEAN` value specifying if the login is an organization admin. By default this value is `FALSE`. `<first_name>`, `<last_name>` The first and last name of the user that will use the created login.

## [](#example)Example

The following command will create an account in the US East (N. Virginia) region.

```
CREATE LOGIN "alexs@acme.com" WITH FIRST_NAME = 'Alex' LAST_NAME = 'Summers';
```