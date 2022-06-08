package json_utils

import (
	"azam-akram/go/utils/json_utils/logger"
	"encoding/json"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

var log *zap.Logger

var jHandler JsonHandler_Interface

type JsonHandler_Interface interface {
	ConvertStringToMap(s string) (map[string]interface{}, error)
	ConvertMapToString(m map[string]interface{}) (string, error)
	ConvertStringToStruct(s string) (*Employee, error)
	ConvertStructToString(e *Employee) (string, error)
	ConvertStringToByte(s string) []byte
	ConvertByteToString([]byte) (string, error)
	ConvertByteToStruct(jsonBytes []byte) (*Employee, error)
	ConvertStructToByte(emp *Employee) (jsonBytes []byte, err error)
	ConvertGenericInterfaceToMap()
	DisplayAllJsonHandlers()
}

type JsonHandler struct {
}

var employeeStr = string(`{
    "id": "The ID",
    "name": "The User",
    "designation": "CEO",
    "address":
    [
        {
            "doorNumber": 1,
            "street": "The office street",
            "city": "The office city",
            "country": "The office country"
        },
        {
            "doorNumber": 1,
            "street": "The home street",
            "city": "The home city",
            "country": "The home country"
        }
    ]
}`)

func (jHandler JsonHandler) DisplayAllJsonHandlers() {

	modifiedEmpMap, err := jHandler.ModifyInputJson(employeeStr)
	if err != nil {
		logger.GetLogger().Print(err)
	}

	logger.GetLogger().PrintKeyValue("DisplayAllJsonHandlers::ModifyInputJson", "modifiedEmpMap", modifiedEmpMap)

	jMap, err := jHandler.ConvertStringToMap(employeeStr)
	if err != nil {
		logger.GetLogger().Print(err)
	}
	logger.GetLogger().PrintKeyValue("DisplayAllJsonHandlers::ConvertStringToMap", "jMap", jMap)

	mapData := map[string]interface{}{
		"id":   "The ID",
		"user": "The User",
	}

	str, err := jHandler.ConvertMapToString(mapData)
	if err != nil {
		logger.GetLogger().Print(err)
	}

	logger.GetLogger().PrintKeyValue("DisplayAllJsonHandlers::ConvertMapToString", "str", str)

	emp, err := jHandler.ConvertStringToStruct(employeeStr)
	if err != nil {
		logger.GetLogger().Print(err)
	}
	logger.GetLogger().PrintKeyValue("DisplayAllJsonHandlers::ConvertStringToStruct", "emp", emp)

	str, err = jHandler.ConvertStructToString(emp)
	if err != nil {
		logger.GetLogger().Print(err)
	}
	logger.GetLogger().PrintKeyValue("DisplayAllJsonHandlers::ConvertStructToString", "str", str)

	jsonBytes := jHandler.ConvertStringToByte(employeeStr)
	logger.GetLogger().PrintKeyValue("DisplayAllJsonHandlers::ConvertStringToByte", "jsonBytes", jsonBytes)

	bytesStr := jHandler.ConvertByteToString(jsonBytes)
	logger.GetLogger().PrintKeyValue("DisplayAllJsonHandlers::ConvertByteToString", "bytesStr", bytesStr)

	emp, err = jHandler.ConvertByteToStruct(jsonBytes)
	if err != nil {
		logger.GetLogger().Print(err)
	}
	logger.GetLogger().PrintKeyValue("DisplayAllJsonHandlers::ConvertByteToStruct", "emp", emp)

	jsonBytes, err = jHandler.ConvertStructToByte(emp)
	if err != nil {
		logger.GetLogger().Print(err)
	}
	logger.GetLogger().PrintKeyValue("DisplayAllJsonHandlers::ConvertStructToByte", "jsonBytes", jsonBytes)

	jHandler.ConvertGenericInterfaceToMap()
}

func (jHandler JsonHandler) ConvertStringToMap(s string) (map[string]interface{}, error) {
	var emp = make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &emp); err != nil {
		return nil, errors.New("cannot convert string to map")
	}

	return emp, nil
}

func (jHandler JsonHandler) ConvertMapToString(m map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	return string(jsonBytes), nil
}

func (jHandler JsonHandler) ConvertStringToStruct(s string) (*Employee, error) {
	var emp *Employee
	err := json.Unmarshal([]byte(employeeStr), &emp)
	if err != nil {
		return nil, errors.New("cannot convert string to struct")
	}

	return emp, nil
}

func (jHandler JsonHandler) ConvertStructToString(e *Employee) (string, error) {
	var jsonBytes []byte

	jsonBytes, err := json.Marshal(e)
	if err != nil {
		return "", errors.New("cannot convert struct to string")
	}

	return string(jsonBytes), nil
}

func (jHandler JsonHandler) ConvertStringToByte(s string) []byte {
	return []byte(s)
}

func (jHandler JsonHandler) ConvertByteToString(jsonBytes []byte) string {
	return string(jsonBytes)
}

func (jHandler JsonHandler) ConvertByteToStruct(jsonBytes []byte) (emp *Employee, err error) {

	err = json.Unmarshal([]byte(employeeStr), &emp)
	if err != nil {
		return nil, errors.New("cannot convert bytes to struct")
	}

	return emp, nil
}

func (jHandler JsonHandler) ConvertStructToByte(emp *Employee) (jsonBytes []byte, err error) {

	jsonBytes, err = json.Marshal(emp)
	if err != nil {
		return nil, errors.New("cannot convert struct to string")
	}

	return jsonBytes, nil
}

func (jHandler JsonHandler) ConvertGenericInterfaceToMap() {
	b := []byte(`{"k1":"v1","k2":6,"k3":["v3","v4"]}`)
	//fmt.Println(b)
	var i interface{}
	_ = json.Unmarshal(b, &i)
	fmt.Println(i)

	d := i.(map[string]interface{})

	for k, v := range d {
		switch vv := v.(type) {
		case string:
			fmt.Printf("key = %s, value = %s, value type = string\n", k, vv)
		case float64:
			fmt.Printf("key = %s, value = %f, value type = float64\n", k, vv)
		case []interface{}:
			fmt.Println(k, "'s value is a array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "unknown type")
		}
	}
}

func (jHandler JsonHandler) ModifyInputJson(s string) (map[string]interface{}, error) {
	var inputEmpMap = make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &inputEmpMap); err != nil {
		return nil, errors.New("cannot convert string to map")
	}

	logger.GetLogger().PrintKeyValue("ModifyInputJson", "inputEmpMap", inputEmpMap)

	inputEmpMap["degree"] = "phd"
	return inputEmpMap, nil
}

/*
func reverseSlice() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Print(nums)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	Print(nums)
}

func printMap() {
	myMap := make(map[string]int)
	myMap["k1"] = 1
	myMap["k2"] = 2
	myMap["k3"] = 3
	Print(myMap)
	Print(myMap["k3"])
}*/
