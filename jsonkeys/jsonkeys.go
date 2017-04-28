package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strings"
)

func iterate(data interface{}) interface{} {

	if reflect.ValueOf(data).Kind() == reflect.Slice {
		d := reflect.ValueOf(data)
		tmpData := make([]interface{}, d.Len())
		returnSlice := make([]interface{}, d.Len())
		for i := 0; i < d.Len(); i++ {
			tmpData[i] = d.Index(i).Interface()
		}
		for i, v := range tmpData {
			returnSlice[i] = iterate(v)
		}
		return returnSlice
	} else if reflect.ValueOf(data).Kind() == reflect.Map {
		d := reflect.ValueOf(data)
		tmpData := make(map[string]interface{})
		for _, k := range d.MapKeys() {
			match, _ := regexp.MatchString("$", k.String())
			typeOfValue := reflect.TypeOf(d.MapIndex(k).Interface()).Kind()
			if match {
				new_key := strings.Replace(k.String(), "$", "", -1)
				if typeOfValue == reflect.Map || typeOfValue == reflect.Slice {
					tmpData[new_key] = iterate(d.MapIndex(k).Interface())
				} else {
					tmpData[new_key] = d.MapIndex(k).Interface()
				}
			} else {
				fmt.Println("debug")
				if typeOfValue == reflect.Map || typeOfValue == reflect.Slice {
					tmpData[k.String()] = iterate(d.MapIndex(k).Interface())
				} else {
					tmpData[k.String()] = d.MapIndex(k).Interface()
				}
			}
		}
		return tmpData
	}
	return data
}

func main() {

	file, e := ioutil.ReadFile("./nested.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	myJson := string(file)
	m, ok := gjson.Parse(myJson).Value().(map[string]interface{})
	if !ok {
		fmt.Println("Error")
	}

	newM := iterate(m)
	jsonBytes, err := json.Marshal(newM)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}
