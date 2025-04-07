# [](#floor)FLOOR

Rounds an input `<value>` down to the nearest multiple based on the specified precision or decimal place. Optionally, you can specify a second parameter to determine which decimal place the value should be rounded down to.

## [](#syntax)Syntax

```
FLOOR(<value>);
FLOOR(<value>, <digit>);
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` The number to be rounded down to the nearest specified place. `NUMERIC`, `DOUBLE PRECISION` `<digit>` (Optional) You can specify a second parameter to define the position of the digit to round down to, based on its distance from the decimal point. Positive numbers indicate digits after the decimal, while negative numbers refer to digits before the decimal. For example, `1` rounds down to the first digit after the decimal, and `-1` rounds down to the nearest ten. A value of `0` rounds down to the nearest whole number. The default is `0`, which means that rounding down happens at the whole number. `INTEGER`

## [](#return-type)Return Type

Input Type Output Type `NUMERIC` `NUMERIC` with same `precision` and `scale`. `DOUBLE PRECISION` `DOUBLE PRECISION`

## [](#remarks)Remarks

When the input is of type `NUMERIC`, `FLOOR` throws an overflow error if the result of `FLOOR` exceeds the defined precision and scale limits of the return data type.

The following code example calculates the nearest whole number smaller than `-99.99` and specifies that the output should contain a total of `4` digits, with only `2` digits reserved for the decimal part:

```
SELECT FLOOR(-'99.99'::NUMERIC(4,2));
```

**Returns**

The previous code returns an `OVERFLOW ERROR` because `FLOOR` returns `-100.00`, which exceeds the `NUMERIC(4,2)` data type’s limit of `2` digits before the decimal point, and `-100` requires `3` digits.

## [](#examples)Examples

The following example returns the nearest whole number smaller than `2.5549900`:

The following example calculates the nearest whole number smaller than `213.1549`, and returns a result of type `NUMERIC(20,4)`, which allows for a total of 20 digits, with 4 values allowed after the decimal point:

The following code example rounds the number `2.5549900` down to the second decimal place. It returns `2.55` because the second parameter 2 specifies rounding to the second digit after the decimal, which corresponds to the hundredths place.

The following example calculates the nearest whole number smaller than `1998` that is a multiple of `1000`. It returns `1000` because the second parameter -3 specifies rounding to the third digit before the decimal point, which corresponds to the thousands place.