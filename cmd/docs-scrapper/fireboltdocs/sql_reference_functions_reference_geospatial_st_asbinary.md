# [](#st_asbinary)ST\_ASBINARY

Converts shapes of the `GEOGRAPHY` data type to the [Well-Known Binary (WKB)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Well-known_binary) format for geographic objects.

## [](#syntax)Syntax

```
ST_ASBINARY(<object>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<object>` The `GEOGRAPHY` object to convert to WKB format. `GEOGRAPHY`

## [](#return-type)Return Type

`ST_ASBINARY` returns a value of type `BYTEA`.

## [](#example)Example

The following query uses `ST_GEOGFROMTEXT` to create a `GEOGRAPHY` object from the WKT representation of a Point at specified longitude and latitude coordinates, and then uses `ST_ASBINARY` to convert it to WKB representation:

```
SELECT ST_ASBINARY(ST_GEOGFROMTEXT('POINT(-73.98551041593687 40.75793403395676)')) AS result
```

**Returns**

result (BYTEA) ‘\\x01010000003d94479a127f52c0502f80fb03614440’