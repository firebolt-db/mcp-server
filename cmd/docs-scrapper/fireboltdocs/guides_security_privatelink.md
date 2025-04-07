# [](#aws-privatelink)AWS PrivateLink

Firebolt supports AWS PrivateLink to help you securely access the Firebolt API without exposing traffic to the public internet. AWS PrivateLink enhances security, minimizes data exposure, and improves network reliability by keeping traffic within AWS. This guide shows you how to:

- [Request AWS PrivateLink access](#request-aws-privatelink-access)
- [Configure your VPC endpoint](#configure-your-vpc-endpoint)
- [Configure your service account](#configure-your-service-account)
- [Test your AWS PrivateLink connection](#test-your-aws-privatelink-connection)

AWS PrivateLink for Firebolt is in public preview and available in all Firebolt regions.

## [](#prerequisites)Prerequisites

Before setting up AWS PrivateLink, ensure you have the following:

- An **AWS account** with permissions to create a [VPC interface endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/create-interface-endpoint.html). If you don’t have access, you can [sign up](https://signin.aws.amazon.com/signup) for an AWS account.
- A **Firebolt Account**: Ensure that you have access to an active Firebolt account. If you don’t have access, you can [sign up for an account](https://www.firebolt.io/sign-up). For more information about how to register with Firebolt, see [Get started with Firebolt](/Guides/getting-started/).
- **Service Account**: You must have access to an active Firebolt [service account](/Guides/managing-your-organization/service-accounts.html), which facilitates programmatic access to Firebolt. Your service account must be associated with a [user](/Overview/organizations-accounts.html#users) with privileges associated with an [account administrative role](/Overview/organizations-accounts.html#account-administrative-role) or an [organizational administrative role](/Overview/organizations-accounts.html#organizational-administrative-role) to request AWS PrivateLink access.

## [](#request-aws-privatelink-access)Request AWS PrivateLink access

After meeting the previous prerequisites, request AWS PrivateLink access using Firebolt’s user interface as follows:

1. [Log in](https://go.firebolt.io/signup) to the **Firebolt Workspace**.
2. Select the **Help** icon (?) at the bottom of the left navigation pane.
3. Select **Request PrivateLink**.
4. In the pop-up window, enter the following information:
   
   1. **Organization Name** – The name of your [organization](/Overview/organizations-accounts.html#organizations) in Firebolt.
   2. **Account** – The Firebolt [account](/Overview/organizations-accounts.html#accounts) associated with a [role](/Overview/organizations-accounts.html#roles) with sufficient permission to request AWS PrivateLink. These include the [account administrative](/Overview/organizations-accounts.html#account-administrative-role) or [organizational administrative roles](/Overview/organizations-accounts.html#organizational-administrative-role).
   3. **AWS Account IDs** – The AWS account IDs for which you want to create a PrivateLink integration. For additional information about permissions, see [Manage permissions](https://docs.aws.amazon.com/vpc/latest/privatelink/configure-endpoint-service.html#add-remove-permissions).
5. After you submit the AWS PrivateLink request, Firebolt’s support team will review it, provision a dedicated VPC endpoint in your Account’s AWS Region, and send an email to the requestor containing the `Endpoint URL` and `endpointServiceId`. Save this information for configuration.

### [](#configure-your-vpc-endpoint)Configure your VPC endpoint

After you have requested AWS PrivateLink on Firebolt’s user interface, login to AWS and configure a VPC endpoint as follows:

01. Sign in to the [AWS Management Console](https://aws.amazon.com/console/).
02. In the search bar at the top, enter **VPC**.
03. Select **VPC (Virtual Private Cloud)** from the dropdown list.
04. In the left navigation pane under **VPC Dashboard**, expand **PrivateLink and Lattice**.
05. Select **Endpoints**.
06. In the upper right corner, select **Create endpoint**.
07. In the **Create endpoint** pane, enter an optional **Name tag** to identify your endpoint.
08. Select the radio button next to **Endpoint services that use NLBs and GWLBs**.
09. In the **Service settings** pane, in the text box under **Service name**, enter the `endpointServiceId` provided in the email from Firebolt’s support team in the previous step to [request AWS PrivateLink access](#request-aws-privatelink-access).
10. Select **Verify service** to confirm that your AWS PrivateLink access is configured correctly. and select Other endpoint services.
11. In the **Network settings** pane, select the down arrow to select your autopopulated **VPC** from the dropdown list.
12. Select the checkbox **Enable DNS NAME**.
13. In the **Subnets** pane select the checkbox next to the subnets that match the **Availability Zone** where your resources reside in your AWS Region.
14. Select the down arrow under **Subnet ID** and hoose the appropriate Subnet ID for your VPC.
15. In the **Security groups** pane, select the checkbox next to the **Group ID** of your security group. Your security group should allow inbound traffic on **port 443** in order to interact with the Firebolt API.
16. Select the **Create endpoint** button in the bottom-right corner of the main workspace.
17. After the endpoint is created, ensure that your security groups and route tables are correctly configured to allow traffic to the endpoint, so that your intended workloads can access Firebolt over the Private API endpoint. Use the following code example to validate your connection to Firebolt by sending it from an EC2 instance in your VPC:

```
curl -v https://api.app.firebolt.io --resolve api.app.firebolt.io:443:<PRIVATE_IP_OF_VPC_ENDPOINT>
```

In the previous code example, replace with the private IP address of your newly created VPC endpoint network interface.

### [](#configure-your-service-account)Configure your service account

Configure your Firebolt [service account](/Guides/managing-your-organization/service-accounts.html) with the `PRIVATE_ONLY` connection preference to ensure it accesses Firebolt only through AWS PrivateLink and uses private APIs. A claim is a setting that defines how a connection behaves. The `PRIVATE_ONLY` claim enforces private networking by restricting access to public endpoints.

Use the following code example inside the **Develop Workspace** in the **Firebolt Workspace**:

```
CREATE SERVICE ACCOUNT IF NOT EXISTS "test_sa"
WITH CONNECTION_PREFERENCE = PRIVATE_ONLY;
```

### [](#test-your-aws-privatelink-connection)Test your AWS PrivateLink connection

After configuring your VPC to use the Firebolt AWS PrivateLink endpoint, test connectivity using the endpoint URL provided by the Firebolt support team.

Use the following curl command to retrieve the private endpoint from your account:

```
curl https://api.go.firebolt.io/web/v3/account/developer/engineUrl \
-H 'Accept: application/json' \
-H "Authorization: Bearer $TOKEN"
{
    "engineUrl": "01hnj9r1xrx3a4t3kb1ec7qs2b.api-private.us-east-1.app.firebolt.io"
}
```

If your service account has the `PRIVATE_ONLY` claim, requests from that service account to any Firebolt public endpoint will fail.

When using the `PRIVATE_ONLY` claim, requests to the private endpoint complete successfully if the traffic originates from an authorized AWS VPC endpoint with the necessary route tables, security group rules, and network access control lists to enable communication.

The following code example sends a `SELECT 42` query to a Firebolt private API endpoint using `curl`, authenticates with a bearer token, and returns a JSON response containing the query result:

```
curl --location 'https://01hnj3r1xrx3a4t3kb1ec7qs2b.api-private.us-east-1.app.firebolt.io' \
--header "Authorization: Bearer $TOKEN" \
--data 'SELECT 42'

{
    "meta": [
        {
            "name": "?column?",
            "type": "int"
        }
    ],
    "data": [
        {
            "?column?": 42
        }
    ],
    "rows": 1,
    "statistics": {
        "elapsed": 0.014256,
        "rows_read": 1,
        "bytes_read": 1
}
```