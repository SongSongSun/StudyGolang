package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}

func createMyQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q)
		tableName := t.Name()
		fieldQuery := fmt.Sprintf("insert into %s(", tableName)
		valueQuery := fmt.Sprintf("values(")
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			fieldName := t.Field(i).Name
			if i == 0 {
				fieldQuery = fmt.Sprintf("%s%s", fieldQuery, fieldName)
			} else {
				fieldQuery = fmt.Sprintf("%s, %s", fieldQuery, fieldName)
			}
			switch v.FieldByName(fieldName).Kind() {
			case reflect.Int:
				if i == 0 {
					valueQuery = fmt.Sprintf("%s%d", valueQuery, v.FieldByName(fieldName).Int())
				} else {
					valueQuery = fmt.Sprintf("%s, %d", valueQuery, v.FieldByName(fieldName).Int())
				}
			case reflect.String:
				if i == 0 {
					valueQuery = fmt.Sprintf("%s\"%s\"", valueQuery, v.FieldByName(fieldName).String())
				} else {
					valueQuery = fmt.Sprintf("%s, \"%s\"", valueQuery, v.FieldByName(fieldName).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		fieldQuery = fmt.Sprintf("%s)", fieldQuery)
		valueQuery = fmt.Sprintf("%s)", valueQuery)
		fmt.Println(fieldQuery, valueQuery)
		return

	}
	fmt.Println("unsupported type")
}

func sqlCreater() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createMyQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createMyQuery(e)
	i := 90
	createMyQuery(i)

}
