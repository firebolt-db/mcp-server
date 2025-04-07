# [](#work-with-semi-structured-data)Work with semi-structured data

Semi-structured data is any data that does not follow a strict tabular schema and often includes fields that are not standard SQL data types. This data typically has a nested structure and supports complex types such as arrays, maps, and structs.

Common formats of semi-structured data include:

- **JSON**— A widely used format for semi-structured data. For information on loading JSON data with Firebolt, see [Load semi-structured JSON data](/Guides/loading-data/working-with-semi-structured-data/load-json-data.html).
- **Parquet and ORC**— Serialization formats that support nested structures and complex data types. For information on loading Parquet data with Firebolt, see [Load semi-structured Parquet data](/Guides/loading-data/working-with-semi-structured-data/load-parquet-data.html).

## [](#firebolts-approach-to-semi-structured-data)Firebolt’s approach to semi-structured data

Firebolt transforms semi-structured data using arrays, enabling efficient querying. Arrays in Firebolt represent the following data constructs:

- **Variable-length arrays**— Arrays with unpredictable lengths in the source data are supported by Firebolt. These arrays can have arbitrary nesting levels, provided the nesting level is consistent within a column and known during table creation.
- **Maps**— Maps, also known as dictionaries, are represented using two coordinated arrays—one for keys and one for values. This approach is particularly useful for JSON-like data where objects have varying keys.

* * *

- [Load semi-structured JSON data](/Guides/loading-data/working-with-semi-structured-data/load-json-data.html)
- [Load semi-structured Parquet data](/Guides/loading-data/working-with-semi-structured-data/load-parquet-data.html)
- [Work with arrays](/Guides/loading-data/working-with-semi-structured-data/working-with-arrays.html)