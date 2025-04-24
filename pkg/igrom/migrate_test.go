package igrom

import (
	"testing"
)

type ColumnMigrate1 struct {
	Name    string
	Type    ColumnType
	Len     int
	Decimal int
	NotNULL bool
	Comment string
}

func TestTable(t *testing.T) {
	//dsn := "root:root@tcp(localhost:13306)/testdb1?charset=utf8mb4&parseTime=True"
	//db, err := gorm.Open(mysql.Open(dsn))
	//if err != nil {
	//	t.Fatal(err)
	//}

	//name := "test_tbl"
	//columns := []*ColumnMigrate1{
	//	{Name: "name", Type: "varchar", Len: 255, NotNULL: true, Comment: "姓名", S: &S{Name: "1"}},
	//	{Name: "age", Type: "int", Comment: "年龄", S: &S{Name: "1"}},
	//	{Name: "score", Type: "float", Comment: "分数", S: &S{Name: "1"}},
	//}
	//if err = NewTableMigrate(db, name).CreateTable(columns...); err != nil {
	//	t.Fatal(err)
	//}

	//t.Log(NewTableMigrate(db, name).HasTable())
	//t.Log(NewTableMigrate(db, name).DropColumn("score"))
	//t.Log(NewTableMigrate(db, name).AddColumn(&ColumnMigrate{
	//	Name:    "score",
	//	Type:    "float",
	//	Len:     10,
	//	Decimal: 2,
	//	Comment: "分数",
	//}))
	//t.Log(NewTableMigrate(db, name).RenameColumn("score1", "score"))

	//t.Log(NewTableMigrate(db, name).AlterColumn(&ColumnMigrate{
	//	Name:    "name",
	//	Type:    VARCHAR,
	//	Len:     128,
	//	Decimal: 0,
	//	NotNULL: true,
	//	Comment: "姓名",
	//}))
}
