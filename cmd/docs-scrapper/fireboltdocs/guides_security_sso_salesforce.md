# [](#salesforce)Salesforce

[Salesforce](https://www.salesforce.com/) is a cloud-based customer relationship management platform that offers applications for sales, marketing, commerce and IT. Salesforce supports implementing secure login systems with authentication methods including single-sign on (SS0).

To integrate Salesforce with Firebolt’s platform, you must [configure your Salesforce organization](https://help.salesforce.com/s/articleView?id=sf.sso_sfdc_idp_saml_parent.htm&type=5) as a SAML identity provider (IdP) for an external service. Then, you must [configure Firebolt’s SSO to work with Salesforce](#configure-firebolt-for-salesforce). Detailed instructions can be found in the following sections:

#### [](#configure-salesforce-application)Configure Salesforce application

1. Login to [Salesforce](https://login.salesforce.com/). If you don’t yet have an account with Salesforce, you can [sign-up](https://www.salesforce.com/form/signup/freetrial-salesforce-starter/) to access their services.
2. In the Salesforce Lightning user interface (UI), enter “Identity Provider” into the **Quick Find** text box at the top of the left navigation bar. For more information about the Lightning user interface, see [Find Your Way Around Setup in Lightning Experience](https://help.salesforce.com/s/articleView?id=sf.lex_find_your_way_around_setup.htm&type=5).
3. From the search results, select **Identity Provider**.
4. Select **Enable Identity Provider**.
5. Select a certificate from the dropdown list. You can use the default self-signed certificate generated with the [SHA-256](https://en.wikipedia.org/wiki/SHA-2) signature algorithm or use your own. For more information, see Salesforce’s documentation for [Certificates and Keys](https://help.salesforce.com/s/articleView?id=sf.security_keys_about.htm&type=5).
6. Select **Download Certificate**. This should download a certificate ending in `.crt` to a local drive.
7. Navigate to the directory containing your downloaded certificate.
8. Convert the certificate to [PEM format](https://en.wikipedia.org/wiki/Privacy-Enhanced_Mail) by entering the following code into a terminal or command prompt on a system where [OpenSSL](https://openssl-library.org/) is installed: `openssl x509 -in original.crt -out sfcert.pem -outform PEM` In the previous code example, replace `original.crt` with the name of your downloaded certificate.
9. From the **Identity Provider** page in Saleforce’s Lightning UI, select **Download Metadata**.

Open the downloaded xml file and find the SingleSignOnService binding, Location attribute, which ends with `../HttpPost`. It will look like: `https://<your-salesforce-account>.my.salesforce.com/idp/endpoint/HttpPost`. Save this value to be used as the SignOnURL in Firebolt SSO configuration.

1. Select **Download Certificate**, and convert the downloaded .crt file to PEM format. You could do this using the following command: `openssl x509 -in original.crt -out sfcert.pem -outform PEM` where `original.crt` is the name of the downloaded .crt file.
2. Select on the provided link to create a new connected app in Salesforce.
3. You will be redirected to the **Manage Connected Apps / New Connected App** view. Fill in required fields **Connected App Name, API Name** (for instance, type ‘Firebolt’) and **Contact email**.
4. Move to **Web App Settings,** and check the **Enable SAML** box.
5. Fill in the Entity Id field with value: `urn:auth0:app-firebolt-v2:<organization_name>-<provider>`, where
   
   - `<organization_name>` is the name of the organization in Firebolt, and
   - `<provider>` is the IdP name, ‘salesforce’ in this case For example: `urn:auth0:app-firebolt-v2:acmeorg-salesforce`
6. Fill the ACS URL field with a URL in the following format (contact Firebolt to get your organization\_identifier) `https://id.app.firebolt.io/login/callback?connection=<organization_name>-<provider>&organization=<organization_identifier>` For example: `https://id.app.firebolt.io/login/callback?connection=acmeorg-salesforce&organization=org_82u3nzTNQPA8RyoM`
   
   > **`<org_name>`** represents the Organizational name used to create your Firebolt Account. The org name is referenced in your vanity URL.  
   > **`<provider>`** represents the provider we’re configuring as our IdP. **`<organization_identifier>`** is the unique identifier for your Organization. To retrieve your **`<organization_identifier>`** , you can navigate to **Configure &gt; SSO** in the Firebolt UI, and **Click Copy organization SSO identifier**.
7. Keep **Subject Type** as Username, and **Name ID Format** as unspecified. Select **Save**.

#### [](#configure-firebolt-for-salesforce)Configure Firebolt for SalesForce

Once your Identity Provider(IdP) is configured, you can now configure Firebolt to integrate with your IdP. This can be done using either the Firebolt UI, or using SQL.

##### [](#ui)UI

1. To configure the Firebolt SSO integration with Salesforce using the UI, Navigate to **Configure &gt; SSO** in Firebolt.
2. Once there, enter your Sign-on URL, Issuer, Provider, Label, Certificate, and field-mappings, where

<!--THE END-->

- `signOnUrl`: The sign-on URL, provided by the SAML identity provider, to which Firebolt sends the SAML requests. The URL is IdP-specific and is determined by the identity provider during configuration.
- `signoutUrl(optional)`: The sign-out URL, provided by the application owner, to be used when the user signs out of the application.\`\`\`
- `issuer`: A unique value generated by the SAML identity provider specifying the issuer value.
- `provider`: The provider’s name - for example: `Salesforce`.
- `label`: The label to use for the SSO login button. If not provided, the Provider field value is used.
- `certificate`: The certificate to verify the communication between the identity provider and Firebolt. The certificate needs to be in PEM or CER format, and can be uploaded from your computer by choosing **Import certificate** or entered in the text box.
- `field mapping`: Mapping to your identity provider’s first and last name in key-value pairs. If additional fields are required, choose **Add another key-value pair**. Mapping is required for Firebolt to fill in the login’s given and last names the first time the user logs in using SSO. Here’s an example of how to set up field mapping:
  
  ```
  {
      "given_name": "name",
      "family_name": "surname"
  }
  ```
  
  In the previous code example, `given_name` is your first name, and is mapped to the “name” field from the IDP. The `family_name` is your last name, and is mapped from the “surname” field. 3. Choose **Update changes**

##### [](#sql)SQL

Values for SQL to create the SSO connection are as follows:

```
ALTER ORGANIZATION acmeorg SET SSO = '{
  "signOnUrl": "https://firebolttest-dev-ed.my.salesforce.com/idp/endpoint/HttpPost",
  "issuer": "salesforce",
  "provider": "salesforce",
  "label": "Salesforce Company App",
  "certificate": "<certificate>"
}';
```

where

- `signOnURL` is the SAML 2.0 endpoint value copied during Salesforce setup,
- `issuer` is the name of the issuer, ‘salesforce’ in this case,
- `provider` is the IdP name, ‘salesforce’ in this case,
- `label` is text that will appear on the **Sign in** form (this defaults to ‘-’ if a value is not provided, for instance ‘acme-salesforce\`), and
- `certificate` is the X.509 certificate in PEM format downloaded during setup.