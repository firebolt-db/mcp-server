# [](#st_y)ST\_Y

Extracts the latitude coordinate of a `GEOGRAPHY` Point. Returns `NULL` for empty geography objects. Returns an error if the input is not a single Point (and not an empty `GEOGRAPHY` object).

## [](#syntax)Syntax

```
ST_Y(<object>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<object>` The `GEOGRAPHY` Point to extract the latitude of. `GEOGRAPHY`

## [](#return-type)Return Type

`ST_Y` returns a value of type `DOUBLE PRECISION`.

## [](#example)Example

The following code example constructs a Point in the `GEOGRAPHY` data type from longitude and latitude coordinates and extracts its latitude coordinate:

```
SELECT ST_Y(ST_GEOGPOINT(-73.98551041593687, 40.75793403395676)) AS result
```

**Returns**

result (DOUBLE PRECISION) 40.75793403395676