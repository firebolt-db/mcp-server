# [](#information-schema-routines)Information schema routines

You can use the `information_schema.routines` view to return information about SQL functions including their database, schema, name, type, return data type, parameter data types, and whether they are deterministic.

**Example**

The following code example lists all aggregate functions in the database and displays their name, parameters, and return data type, sorted by function name:

```
SELECT
  routine_name || '(' || array_to_string(routine_parameters) || ') => ' || data_type
FROM
  information_schema.routines
WHERE
  routine_type = 'AGGREGATE FUNCTION'
ORDER BY routine_name
```

**Example**

The following code example retrieves the names of all routines from the database that return a data type of `GEOGRAPHY`:

```
SELECT
  routine_name
FROM
  information_schema.routines
WHERE
  data_type = 'geography'
```

## [](#columns-in-information_schemaroutines)Columns in `information_schema.routines`

Each row in `information_schema.routines` contains the following information about the SQL function:

Column Name Data Type Description routine\_catalog TEXT The database containing the function. routine\_schema TEXT The schema containing the function. routine\_name TEXT The name of the function. routine\_type TEXT The type of the function, which can be either `FUNCTION`, `AGGREGATE FUNCTION`, `WINDOW FUNCTION` or `TABLE FUNCTION`. data\_type TEXT The return type name of the function. These type names are simplified, for example, returning `ARRAY` instead of `ARRAY(INT)`. If the function accepts and returns data any type, its `data_type` is `ANY`. is\_deterministic TEXT A flag that indicates if the function is deterministic or not. A deterministic function returns the same output when called with the same inputs *within the same query*, even if it can produce a different output when called in a different query. An example of a deterministic function is `CURRENT_TIMESTAMP`. An example of a non-deterministic function is `RANDOM`. routine\_parameters ARRAY(TEXT) An array of data types representing the function’s parameters.