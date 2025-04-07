# [](#numeric-data-type)NUMERIC data type

This topic describes the Firebolt implementation of the `NUMERIC` data type.

- [Overview](#overview)
  
  - [Storage size for precision](#storage-size-for-precision)
  - [Default values for precision and scale](#default-values-for-precision-and-scale)
  - [Precision vs. scale](#precision-vs-scale)
  - [Type conversion (type promotion)](#type-conversion-type-promotion)
  - [Supported operators](#supported-operators)

## [](#overview)Overview

The `NUMERIC` data type is an exact numeric data type defined by its precision (total number of digits) and scale (number of digits to the right of the decimal point).

`NUMERIC` has two optional input parameters: `NUMERIC(precision, scale)`

The maximum precision is 38. The scale value is bounded by the precision value `(scale<=precision)`.

The precision must be positive, while the scale can be zero or positive.

The `DECIMAL` data type is a synonym to the `NUMERIC` data type.

### [](#storage-size-for-precision)Storage size for precision

The storage size of the data type depends on the precision, as described in the following table:

precision bytes 1-9 4 10-18 8 18-38 16

For example, if the precision is between 1 and 9 digits, the data type will use 4 bytes of storage.

### [](#default-values-for-precision-and-scale)Default values for precision and scale

If the scale is not specified when declaring a column of `NUMERIC` data type, then it defaults to `NUMERIC(precision, min(9, precision))`

If both the precision and scale are not specified, then it defaults to `NUMERIC(38, 9)`

### [](#precision-vs-scale)Precision vs. scale

If the scale of a value to be stored is greater than the declared scale of the column, the system will round the value to the specified number of fractional digits. If the number of digits to the left of the decimal point exceeds the declared precision, minus the declared scale, an error results.

```
  SELECT CAST(100.76 AS NUMERIC(5,2)); -- 100.76
  SELECT CAST(100.76 AS NUMERIC(5,1)); -- 100.8
  SELECT(100.76 AS NUMERIC(3,1)); -- error
```

### [](#type-conversion-type-promotion)Type conversion (type promotion)

Note: `NUMERIC` precision and scale does not change with any operation, i.e., there is no automatic widening of the result type to prevent overflows. If a wider type is desired, cast the inputs to perform the operations on wider types.

1. **P1≠P2 or S1≠S2** (casting required)
   
   Any operation between two `NUMERIC`s with different precision and/or scale requires explicit casting of the input to the desired precision and scale.
   
   ```
    NUMERIC(P1, S1) <OP> NUMERIC(P2, S2) -> ERROR (P1≠P2 or S1≠S2) 
   ```
   
   where P1≠P2 or S1≠S2
2. **P1=P2 and S1=S2** (casting optional)
   
   If the two `NUMERIC`s have the same precision and scale, the result will implicitly cast to the same precision and scale. You can still explicitly cast to any other precision and scale.
   
   ```
    NUMERIC(P1, S1) <OP> NUMERIC(P2, S2) -> NUMERIC(P1, S1)
   ```

### [](#supported-operators)Supported operators

**Operators:**

\+  
\-  
\*  
/