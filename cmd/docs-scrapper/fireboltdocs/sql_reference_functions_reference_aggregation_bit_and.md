# [](#bit_and)BIT\_AND

Performs a bitwise `AND` operation on an integer expression, ignoring null input values. Bitwise `AND` compares two bits and returns `1` if both are `1`, and `0` otherwise.

Numbers are represented in two’s complement, a binary method for signed integers as follows:

- Positive numbers are represented in standard binary form, while negative numbers are derived by inverting the bits of their positive counterpart and adding `1`.
- A leftmost bit of `0` indicates a positive number, while `1` indicates a negative number.

## [](#syntax)Syntax

```
BIT_AND
([ DISTINCT ] <expression>)
```

Note: `DISTINCT` has no effect on the function’s result.

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The expression used to compute the result. `INT`, `BIGINT`

## [](#return-types)Return Types

The `BIT_AND` function returns a result of either type `INT` or `BIGINT`, depending on the type of the input.

## [](#examples)Examples

**Example**

The following code example performs a bitwise `AND` operation across all integers ranging from `1` to `3`:

```
SELECT BIT_AND(a)
FROM GENERATE_SERIES(1, 3) as a;
```

**Returns**

The previous code example returns `0`. In a 4-bit system, the binary representation of integers from `1` to `3` is:

- `1` `->` `0001`
- `2` `->` `0010`
- `3` `->` `0011`

The bitwise `AND` function returns `0` because there is no bit position where all three integers have a `1`. The bitwise `AND` of `0001` and `0010` is `0000`, which equals `0`. The bitwise `AND` between `0000` and `0011` is `0000`, or `0`.

**Example**

The following code performs a bitwise `AND` operation across two rows in column `a` that contain the values `-1` and `3`, respectively:

```
SELECT BIT_AND(a)
FROM UNNEST([-1,3]) as a;
```

**Returns**

The previous code example returns `3`. In a 4-bit system, the binary representation of `-1` and `3` is:

- `-1` `->` `1111`
- `3` `->` `0011`

The bitwise `AND` of `1111` and `0011` is `0011` because the last two positions have `1` in both values.