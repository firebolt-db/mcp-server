# [](#if)IF

Evaluates a condition and returns different results based on whether the condition is true or false. The `IF` function is a simplified alternative to the `CASE` expression for handling conditional logic.

## [](#syntax)Syntax

```
IF(<condition>, <then>, <else>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<condition>` Condition that the function evaluates. `BOOLEAN`

`<then>` Value returned when evaluates to \`true\`. Any

`<else>` Value returned when evaluates to \`false\` or \`NULL\`. Must match the data type of . Any

## [](#return-type)Return type

The `IF` function returns the same data type as the and parameters.

## [](#example)Example

The following example uses the `IF` function to determine if the current day is a weekend or weekday:

```
SELECT IF(EXTRACT(DOW FROM CURRENT_DATE()) % 6 = 0, 'Weekend', 'Weekday')
```

The previous query returns:

- `Weekend` when the current day is Saturday or Sunday.
- `Weekday` for any other day.