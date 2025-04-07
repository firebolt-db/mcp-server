# [](#jdbc-driver)JDBC driver

Firebolt’s [type 4](https://en.wikipedia.org/wiki/JDBC_driver#Type_4_driver_%E2%80%93_Database-Protocol_driver/Thin_Driver%28Pure_Java_driver%29) JDBC driver lets Java applications connect to Firebolt. The JDBC driver is open-source software released under an Apache 2 license. You can browse, fork, download, and contribute to its development on [GitHub](https://github.com/firebolt-db/jdbc).

- [Download the JAR file](#download-the-jar-file)
- [Adding the Firebolt JDBC driver as a Maven dependency](#adding-the-firebolt-jdbc-driver-as-a-maven-dependency)
- [Adding the Firebolt JDBC driver as a Gradle dependency](#adding-the-firebolt-jdbc-driver-as-a-gradle-dependency)
- [Connecting to Firebolt with the JDBC driver](#connecting-to-firebolt-with-the-jdbc-driver)
- [Authentication](#authentication)
  
  - [Available connection parameters](#available-connection-parameters)
  - [System settings as connection parameters](#system-settings-as-connection-parameters)
- [Applying system settings using SET](#applying-system-settings-using-set)
- [Connection validation](#connection-validation)
- [Full reference documentation](#full-reference-documentation)

## [](#download-the-jar-file)Download the JAR file

The Firebolt JDBC driver is provided as a JAR file and requires [Java 11](https://java.com/en/download/manual.jsp) or later.

Download the driver from [GitHub JDBC releases](https://github.com/firebolt-db/jdbc/releases).

## [](#adding-the-firebolt-jdbc-driver-as-a-maven-dependency)Adding the Firebolt JDBC driver as a Maven dependency

To connect your project to Firebolt using [Apache Maven](https://maven.apache.org/), add the Firebolt JDBC driver as a dependency in your **pom.xml** configuration file. Link to the [Firebolt Maven repository](https://central.sonatype.com/artifact/io.firebolt/firebolt-jdbc), so that Maven can download and include the JDBC driver in your project, as shown in the following code example:

```
<!-- pom.xml  -->

<project ...>
       <dependency>
            	<groupId>io.firebolt</groupId>
            	<artifactId>firebolt-jdbc</artifactId>
            	<version>3.3.0</version>
       </dependency>
</project>
```

In the previous code example, replace `<version>3.3.0</version>` with the latest version available in the [Firebolt Maven Central repository](https://central.sonatype.com/artifact/io.firebolt/firebolt-jdbc).

## [](#adding-the-firebolt-jdbc-driver-as-a-gradle-dependency)Adding the Firebolt JDBC driver as a Gradle dependency

If you are using the [Gradle Build Tool](https://gradle.org/), you can configure your Gradle project to use the Firebolt JDBC driver by specifying Apache’s [Maven Central](https://maven.apache.org/repository/index.html) as a repository and adding the Firebolt JDBC driver as a dependency as follows:

```
/* build.gradle */

repositories {
    mavenCentral()
}

dependencies {
    implementation 'io.firebolt:firebolt-jdbc:3.3.0'
}
```

In the previous code example, replace `3.3.0` with the latest version available in the [Firebolt Maven Central repository](https://central.sonatype.com/artifact/io.firebolt/firebolt-jdbc).

## [](#connecting-to-firebolt-with-the-jdbc-driver)Connecting to Firebolt with the JDBC driver

Provide connection details to the Firebolt JDBC driver using a connection string in the following format:

```
jdbc:firebolt:<database>?<connection_params>
```

In the previous connection example, the following apply:

- `<database>` - Specifies the name of the Firebolt database to connect to.
- `<connection_params>` - A list of connection parameters formatted as a standard [URL query string](https://en.wikipedia.org/wiki/Query_string#Structure).

## [](#authentication)Authentication

To authenticate, use a [service account ID and secret](/Guides/managing-your-organization/service-accounts.html). A service account, which is used for programmatic access to Firebolt, uses a `client_id` and a `client_secret` for identification. To ensure compatibility with tools external to Firebolt, you can specify the service account’s `client_id` as `user` and `client_secret` as `password`.

The following are examples of how to specify connection strings for authentication and configuration:

**Example**

The following example connection string configures the Firebolt JDBC driver to connect to `my_database` using a specified `client_id` and `secret_id` for authentication:

```
  jdbc:firebolt:my_database?client_id=<client_id>&client_secret=<client_secret>&account=my_account&engine=my_engine&buffer_size=1000000&connection_timeout_millis=10000
```

The previous example string also specifies an account name `my_account`, an engine name `my_engine`, a buffer size of `1000000` bytes, and a connection timeout of `10000` milliseconds, or `10` seconds.

**Example**

The following example provides `client_id` and `client_secret` as separate properties, rather than embedding them directly in the connection string, as shown in the previous example.

Connection string:

```
  jdbc:firebolt:my_database?account=my_account&engine=my_engine&buffer_size=1000000&connection_timeout_millis=10000`
```

Connection properties:

```
  client_id=<client_id>
  client_secret=<client_secret>
```

**Example**

The following example connects to `my_database` using only connection properties for authentication and parameters, without including any parameters directly in the string.

Connection string:

```
  jdbc:firebolt:my_database
```

Connection properties:

```
  client_id=<client_id>
  client_secret=<client_secret>
  account=my_account
  engine=my_engine
  buffer_size=1000000
  connection_timeout_millis=10000
```

**Example**

The following example is a minimal URL that connects to `my_database` using `client_id` and `client_secret` as connection properties for authentication, omitting the engine name and therefore connects to default engine and relying on default values for all other parameters:

Connection string:

```
  jdbc:firebolt:my_database
```

Connection properties:

```
  client_id=<client_id>
  client_secret=<client_secret>
  account=my_account
```

Because the previous configuration example omits specifying the engine name, `my_database` connects to the default engine.

Since the connection string is a URI, make sure to [percent-encode](https://en.wikipedia.org/wiki/Percent-encoding) any reserved characters or special characters used in parameter keys or parameter values.

### [](#available-connection-parameters)Available connection parameters

The following table lists the available parameters that can be added to a Firebolt JDBC connection string. All parameter keys are case-sensitive.

Parameter key Data type Default value Range Description client\_id TEXT No default value.   (**Required**) The Firebolt service account ID. client\_secret TEXT No default value.   (**Required**) The secret generated for the Firebolt service account. account TEXT No default value.   (**Required**) Your Firebolt account name. database TEXT No default value.   The name of the database to connect to. Takes precedence over the database name provided as a path parameter. engine TEXT The default engine attached to the specified database.   The name of the engine to connect to. buffer\_size INTEGER `65536` `1` to `2147483647` The buffer size, in bytes, that the driver uses to read the responses from the Firebolt API. connection\_timeout\_millis INTEGER `60000` `0` to `2147483647` The wait time in milliseconds before a connection to the server is considered failed. A timeout value of zero means that the connection will wait indefinitely. max\_connections\_total INTEGER `300` `1` to `2147483647` The maximum total number of connections. socket\_timeout\_millis INTEGER `0` `0` to `2147483647` The socket timeout, in milliseconds, which specifies the maximum wait time for data, defining the longest allowed inactivity between consecutive data packets. A value of zero means that there is no timeout limit. connection\_keep\_alive\_timeout\_millis INTEGER `300000` `1` to `2147483647` Defines the duration to keep a server connection open in the connection pool before it is closed. ssl\_mode TEXT `strict` `strict` or `none` When set to `strict`, the SSL or TLS certificate is validated for accuracy and authenticity. If set to `none`, certificate verification is omitted. ssl\_certificate\_path TEXT No default value.   The absolute file path for the SSL root certificate. validate\_on\_system\_engine BOOLEAN `FALSE` `TRUE` or `FALSE` When set to `TRUE`, the connection is always validated against a system engine, even if it’s connected to a regular engine. For more information, see [Connection validation](#connection-validation).

### [](#system-settings-as-connection-parameters)System settings as connection parameters

In addition to the parameters specified in the previous table, any [system setting](/Reference/system-settings.html) can be passed as a connection string parameter. For example, to set a custom time zone, use the following format:

```
jdbc:firebolt:my_database?time_zone=UTC&<other_connection_params>
```

## [](#applying-system-settings-using-set)Applying system settings using SET

In addition to passing system settings as connection string parameters, any [system setting](/Reference/system-settings.html) can be passed using the SQL `SET` command. Multiple `SET` statements can be run consecutively, separated by semicolons, as shown below:

```
SET time_zone = 'UTC';
SET standard_conforming_strings = false;
```

## [](#connection-validation)Connection validation

The Firebolt JDBC driver validates the connection by sending a `SELECT 1` query to the system engine. If this query fails, the driver throws an exception. You can use the `validate_on_system_engine` parameter to customize validation. When it is set to `true`, the validation query is sent to the system engine, even if the connection is established with a regular engine. This feature can be useful if you want to stop the regular engine but still need to validate the connection.

The following example configures the Firebolt JDBC driver to connect to `my_database` and validate the connection using the system engine with additional connection parameters specified in `other_connection_parameters`:

```
jdbc:firebolt:my_database?validate_on_system_engine=true&<other_connection_params>
```

## [](#full-reference-documentation)Full reference documentation

The complete documentation for classes and methods in the Firebolt JDBC driver is available in the [Firebolt JDBC API reference guide](https://jdbc.docs.firebolt.io/javadoc/).