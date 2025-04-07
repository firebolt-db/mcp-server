# [](#first_value)FIRST\_VALUE

Returns the first value evaluated in the specified window frame. If there are no rows in the window frame, returns NULL.

## [](#syntax)Syntax

```
FIRST_VALUE( <expression> ) OVER ( [ PARTITION BY <partition_by> ] ORDER BY <order_by> [ASC|DESC] )
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` A SQL expression of any type to evaluate. Any `<partition_by>` An expression used for the `PARTITION BY` clause. Any `<order_by>` An expression used for the order by clause. Any

## [](#return-types)Return Types

Same as the input type of `<expression>`.

This function respects `NULL` values, and the results will be ordered with the default `NULL` ordering `NULLS LAST` unless

otherwise specified in the `ORDER BY` clause. If no `ORDER BY` clause is applied, the order will be undefined.

## [](#examples)Examples

The following code example selects the `nickname`, `level`, `current_score`, and `highest_score` for each level, using the `NTH_VALUE` function to retrieve the top score within each level, ordered by `current_score` in descending order.

**Example**

The following query uses `UNNEST` to convert an array `[1,2,5]` into a column, order the values, and returns the first value as `result`:

```
SELECT FIRST_VALUE(a) OVER (ORDER BY a) as result FROM UNNEST(array[1,2,5]) as a;
```

**Returns**

result (INTEGER) 1 1 1

The previous code example returns the first element after ordering.

**Example**

The following query uses `UNNEST` to convert an array `[100, NULL, 1]` into a column, order the values, and returns the first value as `result`:

```
SELECT FIRST_VALUE(a) OVER (ORDER BY a desc nulls first) as result FROM UNNEST(array[100,NULL,1]) as a;
```

**Returns**

result (INTEGER) NULL NULL NULL

The previous code example returns `NULL` values because `ORDER BY` specifies ‘nulls first’.