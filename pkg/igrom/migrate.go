package igrom

import (
	"fmt"

	"gorm.io/gorm"
)

type ColumnType string

const (
	Int    ColumnType = "int"
	BigInt ColumnType = "bigint"

	Float  ColumnType = "float"
	Double ColumnType = "double"

	VARCHAR ColumnType = "varchar"
	Text    ColumnType = "text"
)

type TableMigrate struct {
	db   *gorm.DB
	name string
}

type ColumnMigrate struct {
	Name    string
	Type    ColumnType
	Len     int
	Decimal int
	NotNULL bool
	Comment string
}

type model struct {
	ID int64 `gorm:"primaryKey;autoIncrement;column:id"`
}

func NewTableMigrate(db *gorm.DB, tableName string) *TableMigrate {
	return &TableMigrate{
		db:   db,
		name: tableName,
	}
}

// CreateTable 创建表
func (t *TableMigrate) CreateTable(columns ...*ColumnMigrate) error {
	if err := t.db.Table(t.name).Migrator().CreateTable(&model{}); err != nil {
		return err
	}

	for _, c := range columns {
		if err := t.AddColumn(c); err != nil {
			return err
		}
	}

	return nil
}

// HasTable 表是否存在
func (t *TableMigrate) HasTable() bool {
	return t.db.Migrator().HasTable(t.name)
}

// DropTable 删除表
func (t *TableMigrate) DropTable() error {
	return t.db.Migrator().DropTable(t.name)
}

// RenameTable 重命名表
func (t *TableMigrate) RenameTable(name string) error {
	return t.db.Migrator().RenameTable(t.name, name)
}

// AddColumn 添加字段
func (t *TableMigrate) AddColumn(column *ColumnMigrate) error {
	if column == nil {
		return fmt.Errorf("column info  is not defined")
	}
	if !t.HasTable() {
		return fmt.Errorf("table is not exists")
	}

	return t.alterColumn(column, "add")
}

// AlterColumn 修改字段
func (t *TableMigrate) AlterColumn(column *ColumnMigrate) error {
	if column == nil {
		return fmt.Errorf("column info  is not defined")
	}
	if !t.HasTable() {
		return fmt.Errorf("table is not exists")
	}
	if !t.HasColumn(column.Name) {
		return fmt.Errorf("column is not exists")
	}

	return t.alterColumn(column, "modify")
}

func (t *TableMigrate) alterColumn(column *ColumnMigrate, op string) error {
	sql := fmt.Sprintf("ALTER TABLE `%v`", t.name)
	switch op {
	case "add":
		sql += fmt.Sprintf(" ADD COLUMN `%v`", column.Name)
	case "modify":
		sql += fmt.Sprintf(" MODIFY COLUMN `%v`", column.Name)
	}

	switch column.Type {
	case Float, Double:
		if column.Len > 0 && column.Decimal > 0 {
			sql += fmt.Sprintf(" %v(%v, %v)", column.Type, column.Len, column.Decimal)
		} else {
			sql += fmt.Sprintf(" %v", column.Type)
		}
	default:
		sql += fmt.Sprintf(" %v(%v)", column.Type, column.Len)
	}

	if column.NotNULL {
		sql += fmt.Sprintf(" %v", "NOT NULL")
	}
	sql += fmt.Sprintf(" COMMENT '%v';", column.Comment)

	if err := t.db.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

// HasColumn 是否有列
func (t *TableMigrate) HasColumn(column string) bool {
	return t.db.Table(t.name).Migrator().HasColumn(&model{}, column)
}

// DropColumn 删除列
func (t *TableMigrate) DropColumn(column string) error {
	return t.db.Table(t.name).Migrator().DropColumn(&model{}, column)
}

// RenameColumn 字段重命名
func (t *TableMigrate) RenameColumn(oldColumn, newColumn string) error {
	return t.db.Table(t.name).Migrator().RenameColumn(&model{}, oldColumn, newColumn)
}

//// CreateIndex 创建索引
//func (t *TableMigrate) CreateIndex(column string) error {
//	return t.db.Table(t.name).Migrator().CreateIndex(&model{}, column)
//}
//
//// DropIndex 删除索引
//func (t *TableMigrate) DropIndex(column string) error {
//	return t.db.Table(t.name).Migrator().DropIndex(&model{}, column)
//}
//
//// HasIndex 某个索引是否存在
//func (t *TableMigrate) HasIndex(column string) bool {
//	return t.db.Table(t.name).Migrator().HasIndex(&model{}, column)
//}
//
//// RenameIndex 重命名索引
//func (t *TableMigrate) RenameIndex(oldColumn, newColumn string) error {
//	return t.db.Table(t.name).Migrator().RenameIndex(&model{}, oldColumn, newColumn)
//}
