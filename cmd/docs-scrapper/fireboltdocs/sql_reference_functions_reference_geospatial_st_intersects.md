# [](#st_intersects)ST\_INTERSECTS

The `ST_INTERSECTS` function determines whether two input `GEOGRAPHY` objects intersect each other.

If either input is empty, `ST_INTERSECTS` will return `FALSE`.

Before performing the intersection check, the two inputs are aligned through a snapping process, ensuring precise calculation. For more details on snapping, refer to the [snapping documentation](/sql_reference/geography-data-type.html#snapping).

## [](#syntax)Syntax

```
ST_INTERSECTS(<geo1>, <geo2>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<geo1>` The first `GEOGRAPHY` object to check for intersection with the second object. `GEOGRAPHY` `<geo2>` The second `GEOGRAPHY` object to check for intersection with the first object. `GEOGRAPHY`

## [](#return-type)Return Type

`ST_INTERSECTS` returns a value of type `BOOLEAN`.

## [](#example)Example

The following code example constructs two LineStrings near Times Square in New York City from their WKT representations as `GEOGRAPHY` objects and determines if the LineStrings intersect:

```
SELECT ST_INTERSECTS(
    ST_GEOGFROMTEXT('LINESTRING(-73.98507474330196 40.75858394491826, -73.98601039902333 40.7573327842733)'),
    ST_GEOGFROMTEXT('LINESTRING(-73.98625022413634 40.758256445020976, -73.98473358363454 40.757637261118134)')
    ) AS result
```

**Returns**

result (BOOLEAN) t