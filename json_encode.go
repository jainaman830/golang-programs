// create a json encoding system for objects
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type Employee struct {
	FullName       string `json:"full_name"`
	Position       string `json:"position"`
	Experience     int32  `json:"experience"`
	Age            int32  `json:"age"`
	ActiveEmployee bool   `json:"active_employee"`
}

// JsonEncode : JsonEncode will return io.Reader and error
func JsonEncode(employee *Employee) (io.Reader, error) {
	encodedData, err := json.Marshal(employee)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(string(encodedData)), nil
}

// ObjectWithoutZeroOrNilValues:ObjectWithoutZeroOrNilValues will return io.reader without nil or 0 values in random order
func ObjectWithoutZeroOrNilValues(employee *Employee) (io.Reader, error) {
	finalFields := make(map[string]interface{})
	v := reflect.ValueOf(*employee)
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i).Interface()
		if !reflect.DeepEqual(val, reflect.Zero(v.Field(i).Type()).Interface()) {
			finalFields[v.Type().Field(i).Tag.Get("json")] = val
		}
	}
	encodedData, err := json.Marshal(finalFields)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(string(encodedData)), nil
}

// StringInSameOrderAsStruct:StringInSameOrderAsStruct will return io.reader without nil or 0 values in same order
func StringInSameOrderAsStruct(employee *Employee) (io.Reader, error) {
	finalStr := "{"
	v := reflect.ValueOf(*employee)
	isBlank := true
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i).Interface()
		if !reflect.DeepEqual(val, reflect.Zero(v.Field(i).Type()).Interface()) {
			comma := ""
			if !isBlank {
				comma = ","
			}
			finalStr += comma + "\"" + v.Type().Field(i).Tag.Get("json") + "\":"
			switch val.(type) {
			case string:
				finalStr += val.(string)
			case int32:
				finalStr += strconv.Itoa(int(val.(int32)))
			case bool:
				if val.(bool) {
					finalStr += "true"
				} else {
					finalStr += "false"
				}
			}
			isBlank = false
		}
	}
	finalStr += "}"
	return strings.NewReader(string(finalStr)), nil
}
func main() {
	Employee := &Employee{
		FullName:       "Aman Jain",
		Position:       "Software Engineer",
		Experience:     0,
		Age:            27,
		ActiveEmployee: false,
	}
	reader, err := JsonEncode(Employee)
	if err != nil {
		log.Fatalf("Error in Json Encode : %f", err)
	}
	outputData, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("Error in reading encoded data : %f", err)
	}
	fmt.Println(string(outputData))

	reader, err = ObjectWithoutZeroOrNilValues(Employee)
	if err != nil {
		log.Fatalf("Error in object encode : %f", err)
	}
	outputData, err = io.ReadAll(reader)
	if err != nil {
		log.Fatalf("Error in reading encoded object : %f", err)
	}
	fmt.Println(string(outputData))

	reader, err = StringInSameOrderAsStruct(Employee)
	if err != nil {
		log.Fatalf("Error in object encode in same order : %f", err)
	}
	outputData, err = io.ReadAll(reader)
	if err != nil {
		log.Fatalf("Error in reading encoded object : %f", err)
	}
	fmt.Println(string(outputData))
}
