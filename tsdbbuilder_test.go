package tsbuilder_test

import (
	"fmt"
	"testing"

	"github.com/tkcrm/tsbuilder"
	"github.com/tkcrm/tsbuilder/tsfuncs"
)

func Test_Create(t *testing.T) {
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
		t.Fatal(err)
	}

	fmt.Println(sql)
}

func Test_Insert(t *testing.T) {
	b := tsbuilder.NewInsertBuilder()

	// add table 1
	b.AddTable("test_table_1").
		Using("s_table_name").
		Columns("column_1", "column_2", "column_3").
		Values(1, 2, 3).
		Values(1, 2, 3).
		Values(1, 2, 3)

	// replace table 1
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
		t.Fatal(err)
	}

	fmt.Println(sql)
}

func Test_Delete(t *testing.T) {
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

	fmt.Println(sql)
}

func Test_Select(t *testing.T) {
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

	fmt.Println(sql)
}

func Test_Database(t *testing.T) {
	b := tsbuilder.NewDatabaseBuilder().
		Name("db_name").
		Options(
			"adsasd 12",
			"fasasdas true",
		)

	sql, err := b.Build()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(sql)
}

func Test_STable(t *testing.T) {
	b := tsbuilder.NewSTableBuilder().
		Name("s_table_name").
		Definitions("vasdvasdv", "caqwqdw").
		Tags(map[string]any{
			"tag_1": 1,
			"tag_2": 2,
			"tag_3": 3,
		}).
		Options(
			"adsasd 12",
			"fasasdas true",
		)

	sql, err := b.Build()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(sql)
}
