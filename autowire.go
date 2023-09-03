package main

import (
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

func createInsertQueryAuto(q interface{}) string {
	return createInsertQuery(reflect.TypeOf(q).Name(), q)
}

func createInsertQuery(tableName string, q interface{}) string {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	if v.Kind() != reflect.Struct {
		panic("unsupported argument type!")
	}
	sql := fmt.Sprintf("insert into %s ", tableName)
	columns := "("
	values := "values ("
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Int:
			if i == 0 {
				columns += t.Field(i).Name
				values += fmt.Sprintf("%d", v.Field(i).Int())
			} else {
				columns += fmt.Sprintf(", %s", t.Field(i).Name)
				values += fmt.Sprintf(", %d", v.Field(i).Int())
			}
		case reflect.String:
			if i == 0 {
				columns += t.Field(i).Name
				values += fmt.Sprintf("'%s'", v.Field(i).String())
			} else {
				columns += fmt.Sprintf(", %s", t.Field(i).Name)
				values += fmt.Sprintf(", '%s'", v.Field(i).String())
			}
		}
	}
	columns += ") "
	values += "); "
	sql += columns + values
	return sql
}

func getSQLType(goType string) string {
	switch goType {
	case "int", "int8", "int16", "int32", "int64":
		return "INT"
	case "string":
		return "VARCHAR(255)"
	// 可根据需要添加其他数据类型的映射
	default:
		return "VARCHAR(255)"
	}
}

func createTable(tableName string, sampleData interface{}) string {
	// 获取结构体类型信息
	t := reflect.TypeOf(sampleData)
	if t.Kind() != reflect.Struct {
		return "data is not a struct"
	}

	// 准备创建表的SQL语句
	createTableSQL := "CREATE TABLE IF NOT EXISTS " + tableName + " ("
	numFields := t.NumField()

	for i := 0; i < numFields; i++ {
		field := t.Field(i)
		fieldName := field.Name
		fieldType := field.Type.Name()
		createTableSQL += fieldName + " " + getSQLType(fieldType)

		// 添加逗号分隔符（如果不是最后一个字段）
		if i < numFields-1 {
			createTableSQL += ", "
		}
	}

	createTableSQL += ")"
	return createTableSQL
}
