# [](#st_covers)ST\_COVERS

The `ST_COVERS` function determines if one `GEOGRAPHY` object fully encompasses another. Specifically, it checks whether every point in the second `GEOGRAPHY` object (`geo2`) lies within or on the boundary of the first `GEOGRAPHY` object (`geo1`). If `geo1` covers `geo2`, the function returns `TRUE`; otherwise, it returns `FALSE`. This function is commonly used to assess spatial relationships, such as whether a larger geographic area completely includes a smaller one.

If either `geo1` or `geo2` is empty, `ST_COVERS` will return `FALSE`.

Before performing the coverage check, `geo1` and `geo2` are aligned through a snapping process, ensuring precise calculation. For more details on snapping, refer to the [snapping documentation](/sql_reference/geography-data-type.html#snapping).

## [](#comparison-with-st_contains)Comparison with [ST\_CONTAINS](/sql_reference/functions-reference/geospatial/st_contains.html)

`ST_COVERS` is similar to `ST_CONTAINS`, but there is a key distinction: `ST_CONTAINS` returns `FALSE` if all points in `geo2` lie exactly on the boundary of `geo1`, while `ST_COVERS` will return `TRUE` in this case. If this distinction is not important, use `ST_COVERS`, which is more efficient to compute in Firebolt.

## [](#syntax)Syntax

```
ST_COVERS(<geo1>, <geo2>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<geo1>` The object being checked to see if it fully covers the second object. `GEOGRAPHY` `<geo2>` The object being checked to see if it fully covers the first object. `GEOGRAPHY`

## [](#return-type)Return Type

`ST_COVERS` returns a value of type `BOOLEAN`.

## [](#example)Example

The following codes example creates a Polygon and a Point around Times Square in New York City from their WKT representations and checks if the Polygon covers the Point:

```
SELECT ST_COVERS(
ST_GEOGFROMTEXT('POLYGON((-73.98519512134791 40.75939223091697, -73.98566488210841 40.75807135585606, -73.9856079414103 40.75804080469413, -73.98470163529633 40.75918017383259, -73.98519512134791 40.75939223091697))'),
ST_GEOGFROMTEXT('POINT(-73.98504378555772 40.75894662495352)')
) AS result
```

**Returns**

result (BOOLEAN) t

## [](#example-1)Example

The following example illustrates the difference between `ST_CONTAINS` and `ST_COVERS`. The Polygon `POLYGON((0 0, 0 1, 1 1, 1 0, 0 0))` covers, but does not contain the LineString `LINESTRING(0 0, 0 0.5)` because it only lies on the Polygon’s boundary and does not intersect the interior of the Polygon. The LineString `LINESTRING(0 0, 0 0.5, 0.5 0.5)` only partially lies on the Polygon’s boundary but also intersects its interior, so it is covered and contained.

```
SELECT 
    polygon, 
    linestring, 
    ST_CONTAINS(
        ST_GEOGFROMTEXT(polygon),
        ST_GEOGFROMTEXT(linestring)
    ) AS contains,
    ST_COVERS(
        ST_GEOGFROMTEXT(polygon),
        ST_GEOGFROMTEXT(linestring)
    ) AS covers
    FROM UNNEST (
        ['POLYGON((0 0, 0 1, 1 1, 1 0, 0 0))','POLYGON((0 0, 0 1, 1 1, 1 0, 0 0))'], 
        ['LINESTRING(0 0, 0 0.5)','LINESTRING(0 0, 0 0.5, 0.5 0.5)']
    ) AS shapes(polygon, linestring)
```

**Returns**

polygon (TEXT) linestring (TEXT) contains (BOOLEAN) covers (BOOLEAN) ‘POLYGON((0 0, 0 1, 1 1, 1 0, 0 0))’ ‘LINESTRING(0 0, 0 0.5)’ f t ‘POLYGON((0 0, 0 1, 1 1, 1 0, 0 0))’ ‘LINESTRING(0 0, 0 0.5, 0.5 0.5)’ t t