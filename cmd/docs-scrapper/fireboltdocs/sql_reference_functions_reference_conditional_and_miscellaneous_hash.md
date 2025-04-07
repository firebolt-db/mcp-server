# [](#hash)HASH

Takes one or more input parameters of any data type and returns a 64-bit non-cryptographic hash value. `HASH` uses the CityHash algorithm for string data types, implementation-specific algorithms for other data types, and the CityHash combinator to produce the resulting hash value. `NULL` values of any type get the same fixed value. See [CITY\_HASH](/sql_reference/functions-reference/conditional-and-miscellaneous/city-hash.html) if `NULL` values should produce `NULL`.

## [](#syntax)Syntax

```
HASH(<expression>, [, expression [,...]])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` An expression that returns any data type that Firebolt supports. Any

## [](#return-type)Return type

`BIGINT`

## [](#example)Example

```
SELECT HASH('esimpson', '08-25-2016')
```

**Returns:** `-6,509,667,128,195,191,394`

```
SELECT HASH(NULL, '08-25-2016')
```

**Returns:** `7610523868633494549`