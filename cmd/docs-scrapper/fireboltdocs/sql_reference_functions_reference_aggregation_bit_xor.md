# [](#bit_xor)BIT\_XOR

Performs a bitwise `XOR` operation on an integer expression, ignoring null input values. Bitwise `XOR`, or exclusive `OR`, compares two bits and returns `1` if they are different, and `0` if they are the same.

Numbers are represented in two’s complement, a binary method for signed integers as follows:

- Positive numbers are represented in standard binary form, while negative numbers are derived by inverting the bits of their positive counterpart and adding `1`.
- A leftmost bit of `0` indicates a positive number, while `1` indicates a negative number.

## [](#syntax)Syntax

```
BIT_XOR
([ DISTINCT ] <expression>)
```

Note: `DISTINCT` has no effect on the function’s result.

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The expression used to compute the result. `INT`, `BIGINT`

## [](#return-types)Return Types

The `BIT_XOR` function returns a result of either type `INT` or `BIGINT`, depending on the type of the input.

## [](#examples)Examples

**Example**

The following code example performs a bitwise `XOR` operation across three rows in column `a` that contain the values `1`, `2`, and `1`, respectively:

```
SELECT BIT_XOR(a)
FROM UNNEST([1,2,1]) as a;
```

**Returns**

The previous code example returns `2`. In a 4-bit system, the binary representation of `1`, and `2` are:

- `1` `->` `0001`
- `2` `->` `0010`

The bitwise `XOR` of `0001` and `0010` from the first and second rows is `0011`, or `3`. The bitwise `XOR` between this result and the `1` in the third row is `0010`, or `2`, because the `XOR` of two identical numbers is `0`.

**Example**

The following code example uses `DISTINCT` to eliminate duplicate values from column `a`, and performs a bitwise `XOR` operation across two rows that contain the values `1` and `2`:

```
SELECT BIT_XOR(DISTINCT a)
FROM UNNEST([1,2,1]) as a;
```

**Returns**

The previous code example returns `3`. In a 4-bit system, the binary representation of `1` and `2` are:

- `1` `->` `0001`
- `2` `->` `0010`

The bitwise `XOR` of `0001` and `0010` is `0011`, or `3`.

**Example**

The following code example performs a bitwise `XOR` operation across a series of integers ranging from `1` to `3`:

```
SELECT BIT_XOR(a)
FROM generate_series(1, 3) as a;
```

**Returns**

The previous code example returns `0`. In a 4-bit system, the binary representation of integers from `1` to `3` are:

- `1` `->` `0001`
- `2` `->` `0010`
- `3` `->` `0011`

The bitwise `XOR` of `0001` and `0010` is `0011`, which equals `3`. The bitwise `XOR` between `0011` and itself is `0000`, or `0`.

**Example**

The following code example performs a bitwise `XOR` operation across a series of integers ranging from `-1` to `2`:

```
SELECT BIT_XOR(a)
FROM generate_series(-1, 2) as a;
```

**Returns**

The previous code example returns `-4`. In a 4-bit system, the binary representation of integers from `-1` to `1` are:

- `-1` `->` `1111`
- `0` `->` `0000`
- `1` `->` `0001`
- `2` `->` `0010`

The bitwise `XOR` of `1111` and `0000` is `1111`, which equals `-1`. The bitwise `XOR` between `1111` and `0001` is `1110`, or `-2`. The bitwise `XOR` between `1110` and `0010` is `1100`, which equals `-4`.