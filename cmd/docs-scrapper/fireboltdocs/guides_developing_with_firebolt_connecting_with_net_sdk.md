# [](#firebolt-net-sdk)Firebolt .NET SDK

## [](#overview)Overview

The Firebolt .NET SDK is a software development kit designed to facilitate the integration of Firebolt’s high-performance database capabilities into .NET applications. This SDK provides developers with the tools and interfaces needed to interact with Firebolt databases efficiently, enabling effective data manipulation and query execution.

## [](#installation)Installation

Install the Firebolt .NET SDK by adding the NuGet package to your project. You can do this in several ways:

### [](#via-package-manager-console)Via Package Manager Console

```
Install-Package FireboltNetSdk
```

### [](#via-net-cli)Via .NET CLI

```
dotnet add package FireboltNetSdk
```

### [](#via-packagereference)Via PackageReference

Add the following line to your project file:

```
<PackageReference Include="FireboltNetSdk" Version="x.x.x" />
```

Make sure to replace `x.x.x` with the specific version you want to use.

### [](#via-visual-studio-ui)Via **Visual Studio UI**

`Tools` &gt; `NuGet Package Manager` &gt; `Manage NuGet Packages for Solution` and search for `Firebolt`

For more details and versioning information, please visit the [NuGet Gallery](https://www.nuget.org/packages/FireboltNetSdk/).

## [](#quick-start)Quick Start

Here’s a simple example to get started with the Firebolt .NET SDK:

```
using System.Data.Common;
using FireboltDotNetSdk.Client;

public class Program
{
    public static async Task Main(string[] args)
    {
        // Name of your Firebolt account
        string account = "my_firebolt_account";
        // Client credentials, that you want to use to connect
        string clientId = "my_client_id";
        string clientSecret = "my_client_secret";
        // Name of database and engine to connect to (Optional)
        string database = "my_database_name";
        string engine = "my_engine_name";
        
        // Construct a connection string using defined parameter
        string conn_string = $"account={account};clientid={clientId};clientsecret={clientSecret};database={database};engine={engine}";
        
        // Create a new connection using generated connection string
        using var conn = new FireboltConnection(conn_string);
        // Open a connection
        conn.Open();

        // First you would need to create a command
        var command = conn.CreateCommand();
        
        // ... and set the SQL query
        command.CommandText = "SELECT * FROM my_table";
        
        // Execute a SQL query and get a DB reader
        DbDataReader reader = command.ExecuteReader();
        
        // Optionally you can check whether the result set has rows
        Console.WriteLine($"Has rows: {reader.HasRows}");
        
        // Close the connection after all operations are done
        conn.Close();
    }
}
```

## [](#documentation)Documentation

For more detailed documentation, including API references and advanced usage, please refer to the [README](https://github.com/firebolt-db/firebolt-net-sdk/blob/main/README.md) file in the repository.

## [](#support)Support

For support, issues, or contributions, please refer to the repository’s issue tracker and contributing guidelines.

## [](#license)License

This SDK is released under **Apache License 2.0**. Please see the [LICENSE](https://github.com/firebolt-db/firebolt-net-sdk/blob/main/LICENSE) file for more details.