package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Age    int
	Weight float64
}

func GetValuesFromInterface(object interface{}) {
	v := reflect.ValueOf(object)
	for i := 0; i < v.NumField(); i++ {
		//get field name
		name := v.Type().Field(i).Name
		fmt.Println(name)
		//check type of value
		valType := v.Field(i).Type()
		fmt.Println(valType)
		//value from direct index
		val := v.Field(i).Interface()
		fmt.Println(val)
		//get value with specific types
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		}
	}
}
func main() {
	var data interface{}
	// default value of empty interface is nil
	fmt.Println(data)
	data = "Aman"
	//fmt.Printf("Type of interface is %T and value is %v", data, data)
	object := Person{
		Name:   "Aman Jain",
		Age:    27,
		Weight: 70.52,
	}
	GetValuesFromInterface(object)
}
