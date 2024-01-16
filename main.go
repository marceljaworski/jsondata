package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// This app use struct types to generate JSON data with json.Marshal
type myJSON struct {
	IntValue        int       `json:"intValue"`
	BoolValue       bool      `json:"boolValue"`
	StringValue     string    `json:"stringValue"`
	DateValue       time.Time `json:"dateValue"`
	ObjectValue     *myObject `json:"objectValue"`
	NullStringValue *string   `json:"nullStringValue,omitempty"`
	NullIntValue    *int      `json:"nullIntValue,omitempty"`
}

type myObject struct {
	ArrayValue []int `json:"arrayValue"`
}

func main() {
	nullInt := 0
	emptyString := ""
	data := &myJSON{
		IntValue:    1234,
		BoolValue:   true,
		StringValue: "hello!",
		DateValue:   time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		ObjectValue: &myObject{
			ArrayValue: []int{1, 2, 3, 4},
		},
		NullStringValue: &emptyString,
		NullIntValue:    &nullInt,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	fmt.Printf("json data: %s\n", jsonData)

	// Parsing JSON using a map
	// We use the json.Unmarshal function with a map[string]interface{} variable to unmarshal JSON data into Go data.

	jsonDataU := `
		{
			"intValue":1234,
			"boolValue":true,
			"stringValue":"hello!",
			"dateValue":"2022-03-02T09:10:00Z",
			"objectValue":{
				"arrayValue":[1,2,3,4]
			},
			"nullStringValue":null,
			"nullIntValue":null,
	
		}
	`

	var dataMap map[string]interface{}
	err = json.Unmarshal([]byte(jsonDataU), &dataMap)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	fmt.Printf("json map: %v\n", dataMap)

	rawDateValue, ok := dataMap["dateValue"] //extract the value of dateValue from the Go data
	if !ok {
		fmt.Printf("dateValue does not exist\n")
		return
	}
	dateValue, ok := rawDateValue.(string) // type assertion
	if !ok {
		fmt.Printf("dateValue is not a string\n")
		return
	}
	fmt.Printf("date value: %s\n", dateValue)

	//Parsing JSON using a Struct

	var dataStruct *myJSON
	err = json.Unmarshal([]byte(jsonDataU), &dataStruct)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	fmt.Printf("json struct: %#v\n", dataStruct)
	fmt.Printf("dateValue: %#v\n", dataStruct.DateValue)
	fmt.Printf("objectValue: %#v\n", dataStruct.ObjectValue)
}
