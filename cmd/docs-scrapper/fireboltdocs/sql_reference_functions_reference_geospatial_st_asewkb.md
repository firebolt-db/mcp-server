# [](#st_asewkb)ST\_ASEWKB

Converts shapes of the `GEOGRAPHY` data type to the [extended Well-Known Binary (EWKB)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Format_variations) format using Spatial Reference Identifier (SRID) 4326, which corresponds to the [WGS84](https://en.wikipedia.org/wiki/World_Geodetic_System#WGS_84) coordinate system.

## [](#syntax)Syntax

```
ST_ASEWKB(<object>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<object>` The `GEOGRAPHY` object to convert to EWKB format. `GEOGRAPHY`

## [](#return-type)Return Type

`ST_ASEWKB` returns a value of type `BYTEA`.

## [](#example)Example

The following code example creates a `GEOGRAPHY` object from the WKT representation of a Point at specified longitude and latitude coordinates and converts it to EWKB format:

```
SELECT ST_ASEWKB(ST_GEOGFROMTEXT('POINT(-73.98551041593687 40.75793403395676)')) AS result
```

**Returns**

result (BYTEA) ‘\\x0101000020e61000003d94479a127f52c0502f80fb03614440’