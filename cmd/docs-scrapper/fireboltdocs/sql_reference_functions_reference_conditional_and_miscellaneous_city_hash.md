# [](#city_hash)CITY\_HASH

Takes one or more input parameters of any data type and returns a 64-bit non-cryptographic hash value. `CITY_HASH` uses the CityHash algorithm for string data types, implementation-specific algorithms for other data types, and the CityHash combinator to produce the resulting hash value. If any of the inputs is `NULL`, the result will be `NULL`. See [HASH](/sql_reference/functions-reference/conditional-and-miscellaneous/hash.html) if `NULL` values should not produce `NULL`.

## [](#syntax)Syntax

```
CITY_HASH(<expression>, [, expression [,...]])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` An expression that returns any data type that Firebolt supports. Any

## [](#return-type)Return type

`BIGINT`

## [](#example)Example

```
SELECT CITY_HASH('esimpson', '08-25-2016')
```

**Returns:** `-6,509,667,128,195,191,394`

```
SELECT CITY_HASH(NULL, '08-25-2016')
```

**Returns:** `NULL`