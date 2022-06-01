package jsonhandling

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var empStr = string(`{
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

func Test_ConvertStringToMap_Success(t *testing.T) {
	assertThat := assert.New(t)

	jh := JsonHandler{}
	jMap, _ := jh.ConvertStringToMap(empStr)

	id := jMap["id"].(string)
	user := jMap["name"].(string)

	assertThat.Equal(id, "The ID")
	assertThat.Equal(user, "The User")
}

func Test_ConvertMapToString_Success(t *testing.T) {
	assertThat := assert.New(t)

	expectedRes := "{\"id\":\"The ID\",\"user\":\"The User\"}"

	mapData := map[string]interface{}{
		"id":   "The ID",
		"user": "The User",
	}

	jh := JsonHandler{}
	jsonStr, _ := jh.ConvertMapToString(mapData)

	assertThat.Equal(jsonStr, expectedRes)
}

func Test_ConvertStringToStruct_Success(t *testing.T) {
	assertThat := assert.New(t)

	jh := JsonHandler{}

	employee, _ := jh.ConvertStringToStruct(empStr)

	assertThat.Equal(employee.ID, "The ID")
	assertThat.Equal(employee.Name, "The User")
}

func Test_ConvertStructToString_Success(t *testing.T) {
	assertThat := assert.New(t)

	employee := &Employee{
		ID:   "The ID",
		Name: "The User",
	}

	jh := JsonHandler{}

	jh.ConvertStructToString(employee)

	assertThat.Equal(employee.ID, "The ID")
	assertThat.Equal(employee.Name, "The User")
}

func Test_ConvertStringToByte_Success(t *testing.T) {
	jh := JsonHandler{}

	jh.ConvertStringToByte(empStr)

	assert.NotNil(t, empStr)
}

func Test_ConvertByteToString_Success(t *testing.T) {
	assertThat := assert.New(t)
	jh := JsonHandler{}

	inputBytes := []byte(`{"id": "The ID", "name": "The User"}`)

	outputString := jh.ConvertByteToString(inputBytes)
	outputBytes := []byte(outputString)

	assertThat.Equal(inputBytes, outputBytes)
}

func Test_ConvertByteToStruct_Success(t *testing.T) {
	assertThat := assert.New(t)
	jh := JsonHandler{}

	byteValue, _ := ioutil.ReadFile("testdata/employee.json")

	outputPref, _ := jh.ConvertByteToStruct(byteValue)

	assertThat.Equal(outputPref.ID, "The ID")
}

func Test_ConvertStructToByte_Success(t *testing.T) {
	jh := JsonHandler{}

	employee := &Employee{
		ID:   "The ID",
		Name: "The User",
	}

	jsonBytes, _ := jh.ConvertStructToByte(employee)

	assert.NotNil(t, jsonBytes)
}
