# utils-go
Common utilitiy code

### Json utilities:
[/json_handling/json_handler.go](https://github.com/azam-akram/utils-go/blob/main/json_handling/json_handler.go)
```
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
```

### HTTP utilities:
Simple HTTP client-server communication but with retry mechanism. HTTP server deliberately returns HTTP 500 (internal server error) for first 3 HTTP request, to "pretend" server is facing some internal errors. The HTTP client continues to retry HTTP request (with 5 maximum number of retries attempt) until it gets 200 OK response.