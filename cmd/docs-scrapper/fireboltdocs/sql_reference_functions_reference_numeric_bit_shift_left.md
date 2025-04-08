# [](#bit_shift_left)BIT\_SHIFT\_LEFT

Shifts the bits in the first argument to the left by `n` bits, where `n` is the second argument. Shifting left by `n` positions is equivalent to multiplying the number by `2^n`.

## [](#syntax)Syntax

```
BIT_SHIFT_LEFT(<value>, <bits>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` Specifies the value to shift. `INT`, `BIGINT` `<bits>` The number of bits to shift. `INT`

## [](#return-types)Return Types

The `BIT_SHIFT_LEFT` function returns a result of either type `INT` or `BIGINT`, depending on the type of the input `<expression>`.

## [](#examples)Examples

**Example**

The following code example shifts `0001`, the binary representation of `1`, to the left by two bits, which yields `0100`, the binary representation for `4`:

**Example**

The following code example shifts `00101`, the binary representation of `5`, to the left by two bits, which yields `10100`, the binary representation of `20`:

**Example**

The following code example shifts the binary representation of `-3`, which is `1111111111111101` in signed two’s complement, one bit to the left, resulting in `1111111111111010`, the binary signed two’s complement of `-6`: