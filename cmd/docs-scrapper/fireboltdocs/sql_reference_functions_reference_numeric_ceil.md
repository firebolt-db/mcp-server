# [](#ceil)CEIL

Rounds an input `<value>` up to the nearest multiple based on the specified precision or decimal place. Optionally, you can specify a second parameter to determine which decimal place the value should be rounded up to.

**Alias:** `CEILING`

## [](#syntax)Syntax

```
CEIL(<value>);
CEIL(<value>, <digit>);
```

OR

```
CEILING(<value>);
CEILING(<value>, <digit>);
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` The number to be rounded up to the nearest specified place. `NUMERIC`, `DOUBLE PRECISION` `<digit>` (Optional) You can specify a second parameter to define the position of the digit to round up to, based on its distance from the decimal point. Positive numbers indicate digits after the decimal, while negative numbers refer to digits before. For example, `1` rounds up to the first digit after the decimal, and `-1` rounds up to the nearest ten. A value of `0` rounds up to the nearest whole number. The default is `0`, which means that rounding up happens at the whole number. `INTEGER`

## [](#return-type)Return Type

Input Type Output Type `NUMERIC` `NUMERIC` with same `precision` and `scale`. `DOUBLE PRECISION` `DOUBLE PRECISION`

## [](#remarks)Remarks

When the input is of type `NUMERIC`, `CEIL` throws an overflow error if the result of `CEIL` exceeds the defined precision and scale limits of the return data type.

The following code example calculates the nearest whole number larger than `99.99` and specifies that the output should contain a total of `4` digits, with only `2` digits reserved for the decimal part:

```
SELECT CEIL('99.99'::NUMERIC(4,2));
```

**Returns**

The previous code returns an `OVERFLOW ERROR` because `CEIL` returns `100.00`, which exceeds the `NUMERIC(4,2)` data type’s limit of `2` digits before the decimal point, and `100` requires `3` digits.

## [](#examples)Examples

The following example returns the nearest whole number larger than `2.5549900`:

The following example calculates the nearest whole number larger than `213.1549`, and returns a result of type `NUMERIC(20,4)`, which allows for a total of 20 digits, with 4 values allowed after the decimal point:

The following example rounds the number `2.5549900` up to the second decimal place. It returns `2.56` because the second parameter 2 specifies rounding to the second digit after the decimal, which corresponds to the hundredths place.

The following example calculates the nearest whole number greater than `1998` that is a multiple of `1000`. It returns `2000` because the second parameter -3 specifies rounding to the third digit before the decimal point, which corresponds to the thousands place.