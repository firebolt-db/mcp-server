# [](#st_geogfromwkb)ST\_GEOGFROMWKB

Constructs a `GEOGRAPHY` object from a [Well-Known Binary](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Well-known_binary) (WKB) byte string. The [extended WKB](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Format_variations) format is supported only for Spatial Reference Identifier (SRID) 4326, which corresponds to the [WGS84](https://en.wikipedia.org/wiki/World_Geodetic_System#WGS_84) coordinate system.

## [](#syntax)Syntax

```
ST_GEOGFROMWKB(<WKB>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<WKB>` WKB representation of the `GEOGRAPHY` object. `BYTEA`

## [](#return-type)Return Type

`ST_GEOGFROMWKB` returns a value of type `GEOGRAPHY`.

## [](#example)Example

The following code example constructs a Point in the `GEOGRAPHY` data type from a WKB byte string and converts it to WKT format:

```
SELECT ST_ASTEXT(ST_GEOGFROMWKB('\x01010000003d94479a127f52c0502f80fb03614440'::BYTEA)) AS result
```

**Returns**

result (TEXT) ‘POINT(-73.98551041593687 40.75793403395676)’