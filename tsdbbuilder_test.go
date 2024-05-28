package tsbuilder_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tkcrm/tsbuilder"
	"github.com/tkcrm/tsbuilder/tsfuncs"
)

func Test_Create(t *testing.T) {
	tests := []struct {
		name   string
		stable string
		tags   map[string]any
		expect string
	}{
		{
			name:   "table_name",
			stable: "s_table_name",
			tags: map[string]any{
				"test":  tsfuncs.Binary("16"),
				"test2": tsfuncs.Binary("24"),
				"test3": 3,
			},
			expect: "CREATE TABLE IF NOT EXISTS table_name USING s_table_name (test, test2, test3) TAGS (BINARY(16), BINARY(24), 3);",
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			b := tsbuilder.NewCreateTableBuilder().
				TableName(tc.name).
				STable(tc.stable).
				Tags(tc.tags)

			sql, err := b.Build()
			require.NoError(t, err)

			require.Equal(t, tc.expect, sql)
		})
	}
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
		Values(1, 2, nil)

	// add table 2
	b.AddTable("test_table_2").
		Using("s_table_name").
		Tags(map[string]any{
			"tag_1": 1.1,
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
	tests := []struct {
		from   string
		wheres []string
		expect string
	}{
		{
			from:   "dbName.test_table",
			wheres: []string{"asasd > asd", "asdfasdf <= 1212"},
			expect: "DELETE FROM dbName.test_table WHERE asasd > asd AND asdfasdf <= 1212;",
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			b := tsbuilder.NewDeleteBuilder().
				From(tc.from).
				Where(tc.wheres...)

			sql, err := b.Build()
			require.NoError(t, err)

			require.Equal(t, tc.expect, sql)
		})
	}
}

func Test_Select(t *testing.T) {
	limit := uint64(10)
	offset := uint64(0)

	tests := []struct {
		columns []string
		from    string
		wheres  []string
		orderBy string
		limit   uint64
		offset  uint64
		expect  string
	}{
		{
			columns: []string{"col_1", "col_2", "col_3"},
			from:    "dbName.test_table",
			wheres:  []string{"asasd > asd", "asdfasdf <= 1212"},
			orderBy: "ts desc",
			limit:   10,
			offset:  0,
			expect:  "SELECT col_1, col_2, col_3 FROM dbName.test_table WHERE asasd > asd AND asdfasdf <= 1212 ORDER BY ts desc LIMIT 10 OFFSET 0;",
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			b := tsbuilder.NewSelectBuilder().
				Columns(tc.columns...).
				From(tc.from).
				Where(tc.wheres...).
				OrderBy(tc.orderBy).
				Limit(&limit).
				Offset(&offset)

			sql, err := b.Build()
			require.NoError(t, err)

			require.Equal(t, tc.expect, sql)
		})
	}
}

func Test_Database(t *testing.T) {
	tests := []struct {
		name    string
		options []string
		expect  string
	}{
		{
			name:    "db_name",
			options: []string{"adsasd 12", "fasasdas true"},
			expect:  "CREATE DATABASE IF NOT EXISTS db_name adsasd 12 fasasdas true;",
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			b := tsbuilder.NewDatabaseBuilder().
				Name(tc.name).
				Options(tc.options...)

			sql, err := b.Build()
			require.NoError(t, err)

			require.Equal(t, tc.expect, sql)
		})
	}
}

func Test_STable(t *testing.T) {
	tests := []struct {
		name        string
		definitions []string
		tags        map[string]any
		options     []string
		expect      string
	}{
		{
			name:        "s_table_name",
			definitions: []string{"vasdvasdv", "caqwqdw"},
			tags: map[string]any{
				"tag_1": 1,
				"tag_2": 2,
				"tag_3": 3,
			},
			options: []string{"adsasd 12", "fasasdas true"},
			expect:  "CREATE STABLE IF NOT EXISTS s_table_name (vasdvasdv, caqwqdw) TAGS (tag_1 1, tag_2 2, tag_3 3) adsasd 12 fasasdas true;",
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			b := tsbuilder.NewSTableBuilder().
				Name(tc.name).
				Definitions(tc.definitions...).
				Tags(tc.tags).
				Options(tc.options...)

			sql, err := b.Build()
			require.NoError(t, err)

			require.Equal(t, tc.expect, sql)
		})
	}
}

func Test_DropTable(t *testing.T) {
	tests := []struct {
		tables []string
		expect string
	}{
		{
			tables: []string{"db_name.test_table", "db_name.test_table_2"},
			expect: "DROP TABLE IF EXISTS db_name.test_table, IF EXISTS db_name.test_table_2;",
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			b := tsbuilder.NewDropTableBuilder().
				Tables(tc.tables...)

			sql, err := b.Build()
			require.NoError(t, err)

			require.Equal(t, tc.expect, sql)
		})
	}
}
