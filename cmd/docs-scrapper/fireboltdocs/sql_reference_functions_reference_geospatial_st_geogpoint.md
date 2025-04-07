# [](#st_geogpoint)ST\_GEOGPOINT

Constructs a Point in the `GEOGRAPHY` data type created from specified longitude and latitude coordinates.

## [](#syntax)Syntax

```
ST_GEOGPOINT(<longitude>, <latitude>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<langitude>` The longitude coordinate of the `GEOGRAPHY` Point to construct. `DOUBLE PRECISION` `<latitude>` The latitude of the `GEOGRAPHY` Point to construct. `DOUBLE PRECISION`

## [](#return-type)Return Type

`ST_GEOGPOINT` returns a value of type `GEOGRAPHY`.

## [](#example)Example

The following code example constructs a Point in the `GEOGRAPHY` data type from longitude and latitude coordinates and converts it to WKT format:

```
SELECT ST_ASTEXT(ST_GEOGPOINT(-73.98551041593687, 40.75793403395676)) AS result
```

**Returns**

result (TEXT) ‘POINT(-73.98551041593687 40.75793403395676)’