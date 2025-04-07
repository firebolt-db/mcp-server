# [](#apache_datasketches_hll_build)APACHE\_DATASKETCHES\_HLL\_BUILD

An aggregate function that creates a new [HyperLogLog sketch](https://datasketches.apache.org/docs/HLL/HLL.html). This sketch is compatible **only** with the [HyperLogLog sketch](https://datasketches.apache.org/docs/HLL/HLL.html) of the [Apache DataSketches](https://datasketches.apache.org/) library. HyperLogLog (HLL) sketches are a highly efficient probabilistic data structure used for cardinality estimation.

The `APACHE_DATASKETCHES_HLL_BUILD` function is used to create a new HLL (HyperLogLog) sketch from a dataset. This is particularly useful for large datasets where exact counting is computationally expensive.

Multiple sketches can be merged to a single sketch using the aggregate function [APACHE\_DATASKETCHES\_HLL\_MERGE](/sql_reference/functions-reference/datasketches/apache-datasketches-hll-merge.html). To estimate the final distinct count value from a sketch, the scalar function [APACHE\_DATASKETCHES\_HLL\_ESTIMATE](/sql_reference/functions-reference/datasketches/apache-datasketches-hll-estimate.html) can be used.

`APACHE_DATASKETCHES_HLL_BUILD` requires less memory than exact count distinct aggregation, but also introduces statistical error. For more information see [Apache HyperLogLog sketch docs](https://datasketches.apache.org/docs/HLL/HLL.html).

## [](#syntax)Syntax

```
APACHE_DATASKETCHES_HLL_BUILD
(<expression> [,hll_precision => <hll_precision>] [, hll_type => <hll_type>] [, text_utf16_little_endian => <text_utf16_little_endian>])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` Any column name or function that return a column name. `INTEGER`, `BIGINT`, `REAL`, `DOUBLE`, `TEXT` `<hll_precision>` Optional. Literal integer value for precision. If not set explicitly, a precision of 12 is used. Range: 12-20. This value represents the log (2 based) of the size of the sketch. For more information see [Apache HyperLogLog sketch](https://datasketches.apache.org/docs/HLL/HLL.html). `INTEGER` `<hll_type>` Optional. Literal text value to set [hll type](https://datasketches.apache.org/docs/HLL/HLL.html). If not included, the default is ‘HLL\_4’. Valid values are: ‘HLL\_4’, ‘HLL\_6’, ‘HLL\_8’. These values represent different levels of compression of the final HLL array where 4, 6 and 8 refer to the number of bits each bucket of the HLL array is compressed down to. HLL\_4 is the most compressed variant but generally slightly slower than the other two, especially during merge operations. `TEXT` `<text_utf16_little_endian>` Optional. Literal boolean value to specify encoding of text if the `<expression>` is of `TEXT` datatype. By default set to `false` means UTF-8 encoding, if set to `true` it means UTF-16 little endian encoding. `BOOLEAN`

## [](#return-type)Return Type

`BYTEA`

## [](#performance-considerations)Performance Considerations

Using HLL sketches is efficient for large datasets, but the accuracy of the estimate can vary. It is recommended to understand the trade-offs between accuracy and performance when using this function. For more information, see [Apache DataSketches HLL Performance](https://datasketches.apache.org/docs/HLL/HllPerformance.html).

## [](#compatibility-with-other-systems)Compatibility With Other Systems

The created sketch is compatible with all the other systems that support the [HyperLogLog sketch](https://datasketches.apache.org/docs/HLL/HLL.html) of the [Apache DataSketches](https://datasketches.apache.org/) library. Few systems applies different encoding before inserting data to the sketch and after creating the sketch, so the sketches created by this function may be incompatible with them as is, but with encoding modification it should be compatible.

#### [](#utf-16-little-endian-text-encoding)UTF-16 Little Endian Text Encoding

Few vendors ([Imply](https://imply.io/)/[Druid](https://druid.apache.org/)) use UTF-16 Little Endian encoding for text data before inserting it into the sketch. In order to be compatible with them, we added this optional parameter. This transition to UTF-16 little endian encoding is done before hashing (sending the string to the DataSketches lib). For example the string ‘ab’ in UTF-8 encoded as ‘0x3132’. in UTF-16: ‘0x00310032’. in UTF-16LE: ‘0x31003200’.

#### [](#sketch-blob-encoding)Sketch Blob Encoding

Few vendors ([Imply](https://imply.io/)/[Druid](https://druid.apache.org/)) apply different encoding to the sketch blob after creating the sketch. In the Druid Database they apply base64 encoding to the sketch blob, in order to make this function sketch compatible with theirs, you can use the following querys:

In order to export a sketch that is compatible with Druid, you can use the following query:

```
SELECT REPLACE(ENCODE(APACHE_DATASKETCHES_HLL_BUILD(...), 'base64'), e'\n', '+');
```

In order to import a sketch that is created with Druid, you can use the following query:

```
SELECT DECODE(<druid_sketch>, 'base64');
```

## [](#example)Example

```
CREATE TABLE data_to_count AS
SELECT *
FROM generate_series(0, 10000000, 3) a;

SELECT count(distinct a) AS accurate_count
FROM data_to_count;
```

accurate\_count (BIGINT) 3333334

```
CREATE TABLE data_to_count2 AS
SELECT *
FROM generate_series(0, 10000000, 2) a;

SELECT count(distinct a) AS accurate_count
FROM data_to_count2;
```

accurate\_count (BIGINT) 5000001

```
CREATE TABLE sketch_of_data_to_count AS
SELECT APACHE_DATASKETCHES_HLL_BUILD(a) a
FROM data_to_count;

INSERT INTO sketch_of_data_to_count
SELECT APACHE_DATASKETCHES_HLL_BUILD(a)
FROM data_to_count2;

SELECT APACHE_DATASKETCHES_HLL_ESTIMATE(a) AS estimate, a AS sketch
FROM sketch_of_data_to_count
ORDER BY 1;
```

estimate (BIGINT) sketch (BYTEA) 3333526 \\x0a0107120a180102…. 5001149 \\x0a0107120a180102….

```
SELECT APACHE_DATASKETCHES_HLL_ESTIMATE(APACHE_DATASKETCHES_HLL_MERGE(a)) AS estimate
FROM sketch_of_data_to_count;
```

estimate (BIGINT) 6673219