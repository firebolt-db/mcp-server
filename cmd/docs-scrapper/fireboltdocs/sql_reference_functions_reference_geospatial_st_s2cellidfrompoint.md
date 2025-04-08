# [](#st_s2cellidfrompoint)ST\_S2CELLIDFROMPOINT

Returns the [S2 cell ID](http://s2geometry.io/devguide/s2cell_hierarchy), which uniquely identifies the region on Earth that fully contains, or covers, a single Point `GEOGRAPHY` object.

`ST_S2CELLIDFROMPOINT` exclusively supports single Point `GEOGRAPHY` objects, and returns `NULL` for all other values including MultiPoint, LineString, Polygon, and an empty `GEOGRAPHY` object.

## [](#syntax)Syntax

```
ST_S2CELLIDFROMPOINT(<geo> [, <cell_level> ])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<geo>` A single valid `GEOGRAPHY` Point for which `ST_S2CELLIDFROMPOINT` calculates and returns the S2 cell that covers it. `GEOGRAPHY` `<cell_level>` The resolution of the S2 cell to return. Valid levels range from 0, the coarsest level, to 30, the finest level. The default level is 30 if unspecified. `BIGINT`

## [](#return-type)Return Type

`ST_S2CELLIDFROMPOINT` returns a value of type `BIGINT`, which may include negative values.

## [](#example)Example

The following code example constructs a dataset containing three `GEOGRAPHY` objects: a valid Point, an empty Point, and a LineString. The code example computes their S2 cell IDs at cell levels 30 and 10, returning results only for the valid Point:

```
WITH data AS (
  SELECT 1 AS id, ST_GEOGFROMTEXT('POINT(-122 47)') AS geo
  UNION ALL
  SELECT 2 AS id, ST_GEOGFROMTEXT('POINT EMPTY') AS geo
  UNION ALL
  SELECT 3 AS id, ST_GEOGFROMTEXT('LINESTRING(1 2, 3 4)') AS geo
)
SELECT id,
       ST_S2CELLIDFROMPOINT(geo) cell30,
       ST_S2CELLIDFROMPOINT(geo, 10) cell10
FROM data;
```

**Returns**

id (INTEGER) cell30 (BIGINT) cell10 (BIGINT) 1 6093613931972369317 6093613287902019584 2 NULL NULL 3 NULL NULL

The previous code example returns the S2 cell ID for the first input at the default, finest resolution of 30. The function returns `NULL` for the second input because it is empty and contains no geographic data, and for the third input because LineString is not a supported `GEOGRAPHY` type.