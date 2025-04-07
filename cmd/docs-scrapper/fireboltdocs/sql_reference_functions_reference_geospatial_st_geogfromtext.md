# [](#st_geogfromtext)ST\_GEOGFROMTEXT

Constructs a `GEOGRAPHY` object from a [Well-Known Text (WKT)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry) string. The [extended WKT format](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Format_variations) is supported only for Spatial Reference Identifier (SRID) 4326, which corresponds to the [WGS84](https://en.wikipedia.org/wiki/World_Geodetic_System#WGS_84) coordinate system.

## [](#syntax)Syntax

```
ST_GEOGFROMTEXT(<WKT>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<WKT>` A WKT string to convert to a `GEOGRAPHY` object. `TEXT`

## [](#return-type)Return Type

`ST_GEOGFROMTEXT` returns a value of type `GEOGRAPHY`.

## [](#example)Example

The following code example constructs a Point from a WKT string describing a Point at specified longitude and latitude coordinates and converts it to WKT format:

```
SELECT ST_ASTEXT(ST_GEOGFROMTEXT('POINT(-73.98551041593687 40.75793403395676)')) AS result
```

**Returns**

result (TEXT) ‘POINT(-73.98551041593687 40.75793403395676)’