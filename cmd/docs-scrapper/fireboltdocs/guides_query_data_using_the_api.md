# [](#firebolt-api)Firebolt API

Use the Firebolt REST API to execute queries on engines programmatically. Learn how to use the API, including authentication, working with engines and executing queries. A service account is required to access the API. Learn about [managing programmatic access to Firebolt](/Guides/managing-your-organization/service-accounts.html).

- [Firebolt API](#firebolt-api)
  
  - [Create a service account and associate it with a user](#create-a-service-account-and-associate-it-with-a-user)
  - [Use tokens for authentication](#use-tokens-for-authentication)
  - [Get the system engine URL](#get-the-system-engine-url)
  - [Execute a query on the system engine](#execute-a-query-on-the-system-engine)
  - [Get a user engine URL](#get-a-user-engine-url)
  - [Execute a query on a user engine](#execute-a-query-on-a-user-engine)

## [](#create-a-service-account-and-associate-it-with-a-user)Create a service account and associate it with a user

Create a service account with organization administrator privilege, i.e., the service account property\_is\_organization\_admin_ must be *true*. Next, create a user with role privileges you would like to have the service account and associate the service account with the user.

## [](#use-tokens-for-authentication)Use tokens for authentication

To authenticate Firebolt using the service accounts with the properties as described above via Firebolt’s REST API, send the following request to receive an authentication token:

```
curl -X POST --location 'https://id.app.firebolt.io/oauth/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=client_credentials' \
--data-urlencode 'audience=https://api.firebolt.io' \
--data-urlencode "client_id=${service_account_id}" \
--data-urlencode "client_secret=${service_account_secret}"
```

where:

Property Data type Description client\_id TEXT The service [account ID](/Guides/managing-your-organization/service-accounts.html#get-a-service-account-id). client\_secret TEXT The service [account secret](/Guides/managing-your-organization/service-accounts.html#generate-a-secret).

**Response**

```
{
  "access_token":"access_token_value",
  "token_type":"Bearer",
  "expires_in":86400
}
```

In the previous example response, the following apply:

- The `access_token` is a unique token that authorizes your API requests that acts as a temporary key to access resources or perform actions. You can use this token to authenticate with Firebolt’s platform until it expires.
- The `token_type` is `Bearer`, which means that the access token must be included in an authorization header of your API requests using the format: `Authorization: Bearer <access_token>`.
- The token `expires_in` indicates the number of seconds until the token expires.

Use the returned access\_token to authenticate with Firebolt.

To run a query using the API, you must first obtain the url of the engine you want to run on. Queries can be run against any engine in the account, including the system engine.

## [](#get-the-system-engine-url)Get the system engine URL

Use the following endpoint to return the system engine URL for `<account name>`.

```
curl https://api.app.firebolt.io/web/v3/account/<account name>/engineUrl \
-H 'Accept: application/json' \
-H 'Authorization: Bearer <access token>'
```

**Example:** `https://api.app.firebolt.io/web/v3/account/my-account/engineUrl`

**Response**

```
{
  "engineUrl":"<prefix>.api.us-east-1.app.firebolt.io"
}
```

## [](#execute-a-query-on-the-system-engine)Execute a query on the system engine

Use the following endpoint to run a query on the system engine:

```
curl --location 'https://<system engine URL>' \
--header 'Authorization: Bearer <access token>' \
--data '<SQL query>'
```

where:

Property Data type Description system engine URL TEXT The system engine URL ([retrieved here](#get-the-system-engine-url)) SQL query TEXT Any valid SQL query (optional) database name TEXT The database name

## [](#get-a-user-engine-url)Get a user engine URL

Get a user engine url by running the following query against the `information_schema.engines` table:

```
SELECT url 
FROM information_schema.engines 
WHERE engine_name='<engine_name>'
```

You can run the query on the system engine using the API with the following request:

```
curl --location 'https://<system engine URL>/query' \
--header 'Authorization: Bearer <access token>' \
--data 'SELECT * FROM information_schema.engines WHERE engine_name='\''my_engine'\'''
```

## [](#execute-a-query-on-a-user-engine)Execute a query on a user engine

Use the following endpoint to run a query on a user engine:

```
curl --location 'https://<user engine URL>&database=<database name>' \
--header 'Authorization: Bearer <access token>' \
--data '<SQL query>'
```

where:

Property Data type Description user engine URL TEXT The user engine URL ([retrieved here](#get-a-user-engine-url)) database name TEXT The database to run the query SQL query TEXT Any valid SQL query

Queries are per request. To run multiple statement queries, separate queries each into one request.