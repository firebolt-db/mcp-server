# [](#bit_or)BIT\_OR

Performs a bitwise `OR` operation on an integer expression, ignoring null input values. Bitwise `OR` compares two bits and returns `1` if either is `1`, and `0` otherwise.

Numbers are represented in two’s complement, a binary method for signed integers, as follows:

- Positive numbers are represented in standard binary form, while negative numbers are derived by inverting the bits of their positive counterpart and adding `1`.
- A leftmost bit of `0` indicates a positive number, while `1` indicates a negative number.

## [](#syntax)Syntax

```
BIT_OR
([ DISTINCT ] <expression>)
```

Note: `DISTINCT` has no effect on the function’s result.

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The expression used to compute the result. `INT`, `BIGINT`

## [](#return-types)Return Types

The `BIT_OR` function returns a result of either type `INT` or `BIGINT`, depending on the type of the input.

## [](#examples)Examples

**Example**

The following code example performs a bitwise `OR` operation across a series of integers ranging from `1` to `3`:

```
SELECT BIT_OR(a)
FROM GENERATE_SERIES(1, 3) as a;
```

**Returns**

The previous code example returns `3`. In a 4-bit system, the binary representation of integers from `1` to `3` is:

- `1` `->` `0001`
- `2` `->` `0010`
- `3` `->` `0011`

The bitwise `OR` of `0001` and `0010` is `0011`, which equals `3`. The bitwise `OR` between `0011` and itself is `0011`, or `3`.

**Example**

The following code example performs a bitwise `OR` operation across a series of integers ranging from `-1` to `1`:

```
SELECT BIT_OR(a)
FROM generate_series(-1, 1) as a;
```

**Returns**

The previous code example returns `-1`. In a 4-bit system, the binary representation of integers from `-1` to `1` is:

- `-1` `->` `1111`
- `0` `->` `0000`
- `1` `->` `0001`

The bitwise `OR` of `1111` and any integer is `1111`, or `-1`.