# TDEngine sql builder

## Overview

High performance SQL builder with zero dependency for [TDEngine](https://tdengine.com)

## Features

- Databases builder
- Super table builder
- Table builder
- Select builder
- Insert builder
- Delete builder
- Native TDEngine query funcs
- Zero dependency

## Install

```bash
go get github.com/tkcrm/tsbuilder
```

## Examples

### Database

```go
b := tsbuilder.NewDatabaseBuilder().
    Name("db_name").
    Options(
        "PRECISION ms",
        "CACHEMODEL last_row",
    )

sql, err := b.Build()
if err != nil {
    log.Fatal(err)
}
```

### Supertable

```go
b := tsbuilder.NewSTableBuilder().
    Name("s_table_name").
    Definitions(
        "ts TIMESTAMP",
        "lat FLOAT",
        "lng FLOAT",
        "speed FLOAT",
    ).
    Tags(map[string]any{
        "deviceID": tsfuncs.Binary("36"),
    }).
    Options(
        "option_1 value_1",
        "option_2 value_2",
    )

sql, err := b.Build()
if err != nil {
    log.Fatal(err)
}
```

### Table

```go
b := tsbuilder.NewCreateTableBuilder().
    TableName("test_table").
    STable("s_table_name").
    Tags(map[string]any{
        "test":  1,
        "test2": 2,
        "test3": 3,
    })

sql, err := b.Build()
if err != nil {
    log.Fatal(err)
}
```

### Insert

```go
b := tsbuilder.NewInsertBuilder()

b.AddTable("test_table_1").
    Using("s_table_name_2").
    Columns("column_1", "column_2", "column_3").
    Values(1, 2, tsfuncs.Now()).
    Values(1, 2, tsfuncs.Abs("4321")).
    Values(1, 2, 3)

// add table 2
b.AddTable("test_table_2").
    Using("s_table_name").
    Tags(map[string]any{
        "tag_1": 1,
        "tag_2": 2,
        "tag_3": 3,
    }).
    Columns("column_1", "column_2", "column_3").
    Values(1, 2, 3).
    Values(1, 2, 3).
    Values(1, 2, 3)

sql, err := b.Build()
if err != nil {
    log.Fatal(err)
}
```

### Select

```go
b := tsbuilder.NewSelectBuilder().
    Columns("col_1", "col_2", "col_3").
    From("dbName.test_table").
    Where(
        "asasd > asd",
        "asdfasdf <= 1212",
    )

sql, err := b.Build()
if err != nil {
    t.Fatal(err)
}
```

### Delete

```go
b := tsbuilder.NewDeleteBuilder().
    From("dbName.test_table").
    Where(
        "asasd > asd",
        "asdfasdf <= 1212",
    )

sql, err := b.Build()
if err != nil {
    t.Fatal(err)
}
```

## Available funcs

- `tsfuncs.Abs(expr)`
- `tsfuncs.Acos(expr)`
- `tsfuncs.Asin(expr)`
- `tsfuncs.Atan(expr)`
- `tsfuncs.Avg(expr)`
- `tsfuncs.Binary(expr)`
- `tsfuncs.Ceil(expr)`
- `tsfuncs.Cos(expr)`
- `tsfuncs.Cos(expr)`
- `tsfuncs.Count(columns ...string)`
- `tsfuncs.Floor(expr)`
- `tsfuncs.Now()`
- `tsfuncs.Round(expr)`
- `tsfuncs.Sin(expr)`
- `tsfuncs.Sum(expr)`

## TODO

- Add more tdengine funcs
- Add `Drop`, `Alter` methods for database, super table and table
- Add more tests

## How to contribute

If you have an idea or a question, just post a pull request or an issue. Every feedback is appreciated.

## License

This project is licensed under the terms of the MIT license.
