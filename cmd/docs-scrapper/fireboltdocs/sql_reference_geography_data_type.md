Firebolt’s geospatial capabilities are currently in public preview. We are gathering feedback and further refining this feature.

# [](#geography-data-type)GEOGRAPHY data type

This topic describes the Firebolt implementation of the `GEOGRAPHY` data type.

- [Overview](#overview)
  
  - [GEOGRAPHY types in Firebolt](#geography-types-in-firebolt)
- [Input and output](#input-and-output)
  
  - [Literal string interpretation](#literal-string-interpretation)
  - [Cast from BYTEA](#cast-from-bytea)
  - [Cast to BYTEA](#cast-to-bytea)
  - [Cast from TEXT](#cast-from-text)
  - [Cast to TEXT](#cast-to-text)
  - [Normalization](#normalization)
  - [Invalid inputs](#invalid-inputs)
- [Snapping](#snapping)
- [Comparison, sorting, and grouping](#comparison-sorting-and-grouping)
- [Driver version requirements](#driver-version-requirements)

## [](#overview)Overview

The `GEOGRAPHY` data type represents geospatial data on the earth’s surface. It represents Points, LineStrings, Polygons, or combinations of these. Points in Firebolt’s `GEOGRAPHY` data type represent points on the [WGS84 reference spheroid](https://en.wikipedia.org/wiki/World_Geodetic_System) (SRID 4326). All edges between points in LineStrings and Polygons are geodesics, meaning they represent the shortest path between points along the Earth’s curvature. Functions that use the `GEOGRAPHY` data type model the Earth as a sphere with a radius of 6,371,008 meters.

### [](#geography-types-in-firebolt)GEOGRAPHY types in Firebolt

Firebolt supports the following `GEOGRAPHY` types, with examples provided in the [Well-Known Text](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry) (WKT) format for spatial data representation:

- Point: A 0-dimensional object representing a single location in coordinate space.
  
  - Example: `POINT (1 2)`
- LineString: A 1-dimensional line made up of a contiguous sequence of line segments, where each segment connects two points. The end of one segment forms the start of the next.
  
  - Example: `LINESTRING (1 2, 3 4, 5 6)`
- Polygon: A 2-dimensional area, defined by an outer boundary (shell) and optionally one or more inner boundaries (holes).
  
  - Example: `POLYGON ((0 0 ,4 0 ,4 4 ,0 4 ,0 0 ), (1 1 ,2 1 ,2 2 ,1 2 ,1 1 )`
- MultiPoint: A collection of multiple Points.
  
  - Example: `MULTIPOINT ((0 0), (1 2))`
- MultiLineString: A collection of multiple LineStrings.
  
  - Example: `MULTILINESTRING ((0 0,1 1,1 2), (2 3,3 2,5 4))`
- MultiPolygon: A collection of Polygons that do not overlap or share adjacent boundaries, although they may touch at finite Points.
  
  - Example: `MULTIPOLYGON (((1 5, 5 5, 5 1, 1 1, 1 5)), ((6 5, 9 1, 6 1, 6 5)))`
- GeometryCollection: A mixed collection of Geographies.
  
  - Example: `GEOMETRYCOLLECTION (POINT(2 3), LINESTRING(2 3, 3 4))`

## [](#input-and-output)Input and output

Firebolt supports creating `GEOGRAPHY` from the industry standard GeoJSON, Well-Known Text (WKT), and Well-Known Binary (WKB) representations, as well as the extended formats EWKT and EWKB introduced by the PostGIS extension. See the corresponding [function documentations](/sql_reference/functions-reference/geospatial/) for further detail.

### [](#literal-string-interpretation)Literal string interpretation

You can create a `GEOGRAPHY` object from a string by using the `GEOGRAPHY` prefix in a query, as shown in the following code example:

```
SELECT GEOGRAPHY '<string in supported format>';
```

`GEOGRAPHY` literals are automatically decoded based on the encoding format used. Supported formats include:

- [(Extended) Well-Known Text (EWKT)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry).
- [GeoJSON](https://datatracker.ietf.org/doc/html/rfc7946).
- Hex-encoded [(Extended) Well-Known Binary (EWKB)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Well-known_binary).

For more information, refer to [Cast from TEXT](#cast-from-text).

**Examples**

You can use any supported format to create a `GEOGRAPHY` object. The following examples demonstrate how to create the same `GEOGRAPHY` object representing a Point at longitude -73.98551041593687 and latitude 40.75793403395676 using WKT, GeoJSON, and EWKB formats:

```
-- Example: WKT format
SELECT GEOGRAPHY 'POINT(-73.98551041593687 40.75793403395676)';
-- Example: GeoJSON format
SELECT GEOGRAPHY '{"type":"Point","coordinates":[-73.98551041593687,40.75793403395676]}';
-- Example: Hex-encoded EWKB format
SELECT GEOGRAPHY '0101000020E61000003D94479A127F52C0502F80FB03614440';
```

### [](#cast-from-bytea)Cast from BYTEA

The cast from `BYTEA` to `GEOGRAPHY` behaves exactly like the [ST\_GEOGFROMWKB](/sql_reference/functions-reference/geospatial/st_geogfromwkb.html) function.

### [](#cast-to-bytea)Cast to BYTEA

The cast from `GEOGRAPHY` to `BYTEA` behaves exactly like the [ST\_ASEWKB](/sql_reference/functions-reference/geospatial/st_asewkb.html) function.

### [](#cast-from-text)Cast from TEXT

The cast from `TEXT` to `GEOGRAPHY` automatically detects the decoding used and supports the [Extended Well-Known Text (EWKT)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry), [GeoJSON](https://datatracker.ietf.org/doc/html/rfc7946), and hex-encoded [Extended Well-Known Binary (EWKB)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Well-known_binary) formats. It behaves exactly like [ST\_GEOGFROMTEXT](/sql_reference/functions-reference/geospatial/st_geogfromtext.html) for [Well-Known Text (WKT)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry) inputs, like [ST\_GEOGFROMGEOJSON](/sql_reference/functions-reference/geospatial/st_geogfromgeojson.html) for [GeoJSON](https://datatracker.ietf.org/doc/html/rfc7946) inputs, and like [ST\_GEOGFROMWKB](/sql_reference/functions-reference/geospatial/st_geogfromwkb.html)[(DECODE(input, ‘HEX’))](/sql_reference/functions-reference/bytea/decode.html) for hex-encoded [Well-Known Binary (WKB)](https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Well-known_binary) inputs.

### [](#cast-to-text)Cast to TEXT

The cast from `GEOGRAPHY` to `TEXT` returns the hexadecimal representation of the output of the [ST\_ASEWKB](/sql_reference/functions-reference/geospatial/st_asewkb.html) function in upper case letters.

### [](#normalization)Normalization

For all input formats, Firebolt applies normalization steps to the input. These steps ensure that Firebolt can perform operations efficiently and correctly.

- All shapes that overlap a Polygon will be removed. This means that Polygons may be merged if they overlap and LineStrings may be cut to only include the parts outside of any Polygons.

![Overlapping Polygons with holes and a LineString crossing through them.](../assets/images/geography/normalization_before.png)

Input with overlapping Polygons and a LineString.

![A single Polygon that combines the two input Polygons and three LineStrings covering the parts of the input LineString not covered by Polygons.](../assets/images/geography/normalization_after.png)

In the output, the Polygons have been merged and the LineString was cut at the Polygon boundaries.

- Duplicate Points in MultiPoints and GeometryCollections are removed.
- Duplicate vertices in LineStrings are removed.
- Any empty shapes will be removed.
- Polygon orientations are adjusted such that the interior of the Polygon is the smaller possible option. This means that any Polygon can cover at most half of the earth.

### [](#invalid-inputs)Invalid inputs

Additionally, Firebolt fixes some otherwise invalid inputs when reading:

- Polygons are split at self-intersections.

![A Polygon with a hole that intersects itself.](../assets/images/geography/intersection_in_hole_before.png)

Input with a self-intersecting hole. The purple vertices are the input vertices of the hole.

![The same Polygon but the hole has been split into two holes with no intersections.](../assets/images/geography/intersection_in_hole_after.png)

In the output, the hole has been split at the intersection. The resulting Polygon has two holes. One consisting of the purple vertices and one consisting of the yellow vertices.

- Degenerate parts of Polygons are removed, potentially splitting the Polygon. Parts of Polygons are considered degenerate if they collapse to a single line or point.

![A Polygon with a part that has no area.](../assets/images/geography/degenerate_part_removal_before.png)

Input Polygon with a degenerate part.

![The same Polygon split into two with the degenerate part removed.](../assets/images/geography/degenerate_part_removal_after.png)

In the output, the degenerate part has been removed, turning the Polygon into a MultiPolygon.

- A special case of removing these degeneracies occurs when Polygon holes share a boundary edge with the Polygon. In this case, the hole is removed from the Polygon and the outer boundary is reshaped accordingly, but only if the edge vertices of the hole are also vertices of the Polygon.

![A Polygon with a hole that shares two edges with the outer shell of the Polygon.](../assets/images/geography/outer_and_inner_ring_share_edge_before.png)

Input Polygon with a hole that shares two edges with the outer shell, making these edges degenerate.

![The same Polygon but with the degenerate parts removed.](../assets/images/geography/outer_and_inner_ring_share_edge_after.png)

In the output, the degenerate part has been removed.

## [](#snapping)Snapping

Functions that determine relations between `GEOGRAPHY` objects perform a process called snapping on their inputs. During snapping, any points of the inputs objects that are less than 1 micrometer apart from each other are moved to be exactly the same. For example, this means that a Point that is less than 1 micrometer away from a Polygon is treated as being on the Polygon’s boundary and two Polygons that are separated by less than 1 micrometer are treated as intersecting.

## [](#comparison-sorting-and-grouping)Comparison, sorting, and grouping

Objects with a `GEOGRAPHY` data type cannot be compared. This means that `ORDER BY` or comparison operators are not supported on `GEOGRAPHY` columns.

This also means that `GEOGRAPHY` columns cannot be used in `GROUP BY` clauses because they cannot be compared for equality.

## [](#driver-version-requirements)Driver version requirements

The following minimum driver versions are required to use the `GEOGRAPHY` data type:

- [Python](/Guides/developing-with-firebolt/connecting-with-Python.html) - version 1.6.0.
- [JDBC](/Guides/developing-with-firebolt/connecting-with-jdbc.html) - version 3.4.0.
- [Node.js](/Guides/developing-with-firebolt/connecting-with-nodejs.html) - any version.
- [Go](/Guides/developing-with-firebolt/connecting-with-go.html) - version 1.3.1.
- [.NET](/Guides/developing-with-firebolt/connecting-with-net-sdk.html) - version 1.4.0.