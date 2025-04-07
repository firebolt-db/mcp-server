# [](#st_astext)ST\_ASTEXT

Converts shapes of the `GEOGRAPHY` data type to the [Well-Known Text (WKT)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry) format.

## [](#syntax)Syntax

```
ST_ASTEXT(<object>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<object>` The `GEOGRAPHY` object to convert to WKT format. `GEOGRAPHY`

## [](#return-type)Return Type

`ST_ASTEXT` returns a value of type `TEXT`.

## [](#example)Example

The following code example creates a `GEOGRAPHY` object from the WKT representation of a Point at specified longitude and latitude coordinates and converts it to WKT format:

```
SELECT ST_ASTEXT(ST_GEOGFROMTEXT('POINT(-73.98551041593687 40.75793403395676)')) AS result
```

**Returns**

result (TEXT) ‘POINT(-73.98551041593687 40.75793403395676)’