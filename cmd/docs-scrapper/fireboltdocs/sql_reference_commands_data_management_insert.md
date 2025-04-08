# [](#insert)INSERT

Inserts one or more values into a specified table. Specifying column names is optional.

- [Syntax](#syntax)
- [Example](#example)

## [](#syntax)Syntax

```
INSERT INTO <table> [(<column1>[, <column2>][, ...])]
{ <expression> | VALUES ([<value1>[, <value2>][, ...]) }
```

## [](#parameters)Parameters

Parameter Description `<table>` The target table where values are to be inserted. `(<column1>[, <column2>][, ...])]` A list of column names from `<table_name>` for the insertion. If not defined, the columns are deduced from the `<select_statement>`. `<expression>`  
–OR–  
`VALUES ([<value1>[, <value2>][, ...])]` You can specify either a [`SELECT` query](/sql_reference/commands/queries/select.html) that determines values to or an explicit list of `VALUES` to insert.

## [](#example)Example

First, create a table populated with student information as follows:

```
CREATE TABLE students (
    id INT,
    full_name TEXT,
    dob DATE,
    gender TEXT)
```

Next, use `INSERT` to add two rows into the `students` table as follows:

```
INSERT INTO students VALUES
    (1, 'Harry Potter', DATE '1980-07-31', 'M'),
    (2, 'Hermione Granger', DATE '1979-09-19', 'F')
```

You can also add another row with only some of the columns populated. The missing `dob` column for date of birth does not have default value, so Firebolt sets it to `NULL`.

```
INSERT INTO students (id, full_name, gender) VALUES (3, 'Ron Weasley', 'M')
```