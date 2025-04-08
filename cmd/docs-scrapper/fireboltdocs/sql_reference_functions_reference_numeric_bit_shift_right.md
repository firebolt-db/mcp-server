# [](#bit_shift_right)BIT\_SHIFT\_RIGHT

Shifts the bits in the first argument to the right by `n` bits, where `n` is the second argument. Shifting right by `n` positions is equivalent to dividing the number by `2^n`.

## [](#syntax)Syntax

```
BIT_SHIFT_RIGHT(<value>, <bits>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` Specifies the value to shift. `INT`, `BIGINT` `<bits>` The number of bits to shift. `INT`

## [](#return-types)Return Types

The `BIT_SHIFT_RIGHT` function returns a result of either type `INT` or `BIGINT`, depending on the type of the input `<expression>`.

## [](#examples)Examples

**Example**

The following code example shifts `0001`, the binary representation of `1`, to the right by two bits, which yields `0000`, the binary representation for `0`:

**Example**

The following code example shifts `00101`, the binary representation of `5`, to the right by two bits, which yields `00001`, the binary representation for `1`:

**Example**

The following code example shifts the binary representation of `-3`, which is `1111111111111101` in signed two’s complement, one bit to the right, resulting in `1111111111111110`, the signed two’s complement of `-2`: