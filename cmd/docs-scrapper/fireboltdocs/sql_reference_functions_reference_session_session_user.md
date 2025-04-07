# [](#session_user)SESSION\_USER

Returns the name of the user running the current query.

## [](#syntax)Syntax

```
SESSION_USER()
```

## [](#return-types)Return Types

`TEXT`

## [](#examples)Examples

The following code example shows the effective privileges of the roles directly assigned to the user running the query:

**Example**

Dynamic security through a view which uses `session_user()`.

```
-- user bob created view:
create view my_employee_data as select * from employees where user_name = session_user();

-- user alice queries it:
select * from my_employee_data; -- session_user() will be evaluated to 'alice' for this query
```

**Returns**

user\_name … alice …