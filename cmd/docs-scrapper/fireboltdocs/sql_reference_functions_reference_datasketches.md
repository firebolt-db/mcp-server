## [](#apache-datasketches-functions)Apache DataSketches functions

[Apache DataSketches](https://datasketches.apache.org/) is a robust library that provides a collection of advanced algorithms for data analysis. These algorithms are designed to work efficiently with large datasets and provide accurate results for various data analysis tasks. The following functions are part of the Apache Datasketches library and can be used in SQL queries to perform advanced data analysis tasks.

### [](#distinct-counting-with-hyperloglog-sketches)Distinct Counting with HyperLogLog Sketches

HyperLogLog (HLL) Sketches are a highly efficient probabilistic data structure used for cardinality estimation, which means they can quickly and accurately count the number of distinct elements in large datasets. The following functions are used to work with HLL sketches in SQL queries: [APACHE\_DATASKETCHES\_HLL\_BUILD](/sql_reference/functions-reference/datasketches/apache-datasketches-hll-build.html), [APACHE\_DATASKETCHES\_HLL\_MERGE](/sql_reference/functions-reference/datasketches/apache-datasketches-hll-merge.html), [APACHE\_DATASKETCHES\_HLL\_ESTIMATE](/sql_reference/functions-reference/datasketches/apache-datasketches-hll-estimate.html).

* * *

- [APACHE\_DATASKETCHES\_HLL\_BUILD](/sql_reference/functions-reference/datasketches/apache-datasketches-hll-build.html)
- [APACHE\_DATASKETCHES\_HLL\_ESTIMATE](/sql_reference/functions-reference/datasketches/apache-datasketches-hll-estimate.html)
- [APACHE\_DATASKETCHES\_HLL\_MERGE](/sql_reference/functions-reference/datasketches/apache-datasketches-hll-merge.html)