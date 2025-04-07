# [](#st_distance)ST\_DISTANCE

The `ST_DISTANCE` function calculates the shortest distance, measured as a geodesic arc between two `GEOGRAPHY` objects, measured in meters. It models the earth as a perfect sphere with a fixed radius of 6,371,008 meters.

If either input is empty, `ST_DISTANCE` will return `NULL`.

## [](#syntax)Syntax

```
ST_DISTANCE(<geo1>, <geo2>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<geo1>` The first `GEOGRAPHY` object to calculate the distance between. `GEOGRAPHY` `<geo2>` The second `GEOGRAPHY` object to calculate the distance between. `GEOGRAPHY`

## [](#return-type)Return Type

`ST_DISTANCE` returns a value of type `DOUBLE PRECISION`.

## [](#example)Example

The following code example constructs two Points from their WKT representations: One at the Statue of Liberty in New York City, and one at the Big Ben in London. It then returns the shortest distance between them, as measured as a geodesic arc, in meters:

```
SELECT ST_DISTANCE(
ST_GEOGFROMTEXT('POINT(-74.04447010745835 40.68924450077543)'),
ST_GEOGFROMTEXT('POINT(-0.12418551935155021 51.50086274661804)')
) AS result
```

**Returns**

result (DOUBLE PRECISION) 5574863.932096738