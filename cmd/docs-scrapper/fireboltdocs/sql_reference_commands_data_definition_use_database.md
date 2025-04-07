# [](#use-database)USE DATABASE

The `USE DATABASE` statement in allows users to specify the database context for their queries within a session. This statement is crucial for operations within Firebolt where users interact with multiple databases, ensuring queries are executed against the intended database without the need to fully qualify object names in their SQL statements.

## [](#syntax)Syntax

```
USE [DATABASE] <database_name>;
```

## [](#parameters)Parameters

Parameter Description `DATABASE` This is an optional keyword. `<database_name>` The target database for session context. Essential for directing queries to the correct database without specifying the database name in each query. The database name must adhere to identifier rules supported by Firebolt, allowing for quoted or unquoted identifiers. An error is thrown if the database does not exist or the user lacks the necessary permissions.

While `USE DATABASE` is designed to set the context at a session level, it plays a crucial role in the resolution of unqualified object names. All subsequent queries in the session resolve unqualified single-part names (like table or function names) based on the set database context, unless explicitly overridden by using fully qualified names or other USE statements is encountered.

## [](#example)Example

The following demonstrates setting the current database context to ‘sales’ database.

```
USE sales; -- Sets the current database context to 'sales'
SELECT * FROM public.revenue; -- Executes against 'sales.public.revenue'
```

The following example demonstrat switching between databases within a single session

```
USE database sales; -- Context set to 'sales'
USE marketing; -- Context switched to 'marketing'
SELECT * FROM campaigns; -- Accesses 'campaigns' tables in  'marketing' database using one part name.
SELECT * FROM sales.public.revenue; -- Accesses 'sales.public.revenue' using fully qualified name, despite current context being 'marketing'
```

## [](#faqs)FAQs

1. What is the purpose of the `USE DATABASE` command? The `USE DATABASE` command sets the current database context for the session, ensuring that all subsequent queries operate within the specified database unless another `USE` command changes the context.
2. How does the `USE DATABASE` command affect the resolution of unqualified object names in queries? Once the `USE DATABASE` command is executed, all unqualified single-part names or two-part names, such as table or function names, in subsequent queries are resolved based on the set database context, simplifying queries by not requiring full names.
3. Can I access objects in different databases without changing the session’s database context after using `USE DATABASE`? Yes, you can access objects in different databases by using fully qualified names (including the database name and schema name) in your queries, which allows you to bypass the current session’s database context.
4. What happens if I use another `USE DATABASE` command within the same session? Executing another `USE DATABASE` command changes the session’s database context to the new database specified, affecting the resolution of unqualified object names in all subsequent queries.
5. How do I ensure my query targets the correct database and object if I have used the `USE DATABASE` command? To target a specific database and object, you can use fully qualified names in your queries regardless of the current session’s database context, or ensure the `USE DATABASE` command has set the desired database before executing unqualified queries.
6. Can I use `USE DATABASE` within a transaction? Yes, but it will only affect the database context for the duration of the transaction.

## [](#common-pitfalls)Common Pitfalls:

• Forgetting to switch back to the original database context after temporary changes can lead to unintended query executions against the wrong database.