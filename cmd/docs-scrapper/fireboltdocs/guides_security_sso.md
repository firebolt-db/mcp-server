# [](#configure-your-idp)Configure your IdP

An **Identity Provider (IdP)** is a service that handles user authentication and manages user identities. When you set up Single Sign-On (SSO), the IdP verifies your users’ credentials and allows them to access multiple applications, including Firebolt, without needing to login repeatedly.

For your organization, an IdP simplifies user management and strengthens security. You can enforce centralized security policies, like multi-factor authentication (MFA), and quickly revoke access when someone leaves the team.

For your users, using single sign-on gives them access to all the tools they need, including Firebolt.

Single-sign on (SSO) is an authentication process that allows access to multiple applications or services with a single set of credentials. It provides a centralized authentication mechanism for your organization so that it’s easier to manage user access, enforce security policies, and revoke access when necessary.

## [](#pre-requisites)Pre-requisites

Before you can use SSO with Firebolt, you must complete specific configuration steps in your Identity Provider (IdP) system, which is responsible for authenticating users and managing their credentials. Part of these steps include defining an **Audience URI**, which specifies the intended recipient of a SAML assertion about a user’s authentication. The configuration of an Audience URI depends on your IdP. See the following list of supported IdPs for specific instructions.

If your Audience URI is not configured correctly, Security Assertion Markup Language (SAML) assertions used for authentication will fail, preventing users from signing in using SSO.

## [](#supported-idps)Supported IdPs

Firebolt allows you to sign in using federated identities. The SSO implementation supports the following IdPs:

- [Auth0](/Guides/security/sso/auth0.html)
- [Okta](/Guides/security/sso/okta.html)
- [OneLogin](/Guides/security/sso/onelogin.html)
- [Salesforce](/Guides/security/sso/salesforce.html)
- [PingFederate (Ping Identity)](/Guides/security/sso/pingfederate.html)
- [Custom Identity provider](/Guides/security/sso/custom-sso.html)

If your IdP is not listed but supports SAML2.0, contact the [Firebolt support team](mailto:support@firebolt.io).