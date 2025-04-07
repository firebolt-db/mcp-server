# [](#date)DATE

Converts a `TIMESTAMP`, `TIMESTAMPTZ`, `DATE`, or `TEXT` value to a `DATE` value. If the conversion cannot be performed, the `DATE` function raises an error.

## [](#syntax)Syntax

```
DATE(<expression>)
```

### [](#aliases)Aliases

```
CAST(<expression> AS DATE)
<expression>::DATE
```

## [](#parameters)Parameters

Parameter Description `<expression>` The expression that should be converted to a `DATE`.

## [](#return-type)Return Type

The `DATE` function returns a result of type `DATE`. If the conversion cannot be performed, the `DATE` function raises an error. For more information on conversion rules, check out the documentation of the [DATE data type](/sql_reference/date-data-type.html).

## [](#example)Example

The following code example converts the `TEXT` value `1990-01-01` to a `DATE` value representing that date:

```
SELECT DATE('1990-01-01');
```

**Returns**

```
1990-01-01
```