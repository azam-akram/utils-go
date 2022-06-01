# utils-go
Common utilitiy code

### [Json utils](https://github.com/azam-akram/utils-go/tree/main/json_utils):
```
type JsonHandler_Interface interface {
	DisplayAllJsonHandlers()
	ConvertStringToMap(s string) (map[string]interface{}, error)
	ConvertMapToString(m map[string]interface{}) (string, error)
	ConvertStringToStruct(s string) (*Employee, error)
	ConvertStructToString(e *Employee) (string, error)
	ConvertStringToByte(s string) []byte
	ConvertByteToString([]byte) (string, error)
	ConvertByteToStruct(jsonBytes []byte) (*Employee, error)
	ConvertStructToByte(emp *Employee) (jsonBytes []byte, err error)
	ConvertGenericInterfaceToMap()
}
```
How to use
Look at `DisplayAllJsonHandlers()` to know how to use other functions exposed by `JsonHandler_Interface` interface. If you want to call all of these functions, simply call `DisplayAllJsonHandlers()`
```
jsonHandler := json_utils.JsonHandler{}
jsonHandler.DisplayAllJsonHandlers()
```

### [HTTP utils](https://github.com/azam-akram/utils-go/tree/main/http_utils):
A simple HTTP client-server communication but with retry in place. 
[HTTP Server](https://github.com/azam-akram/utils-go/tree/main/http_utils/server) deliberately returns HTTP 500 (internal server error) for first 3 HTTP request, to "mock" some internal server errors. 
[HTTP Client](https://github.com/azam-akram/utils-go/tree/main/http_utils/client) continues to retry HTTP request (with 5 maximum number of retries attempt) until it gets 200 OK response.
Retry config
```
const (
	retryCount              = 5
	retryMinWaitTimeSeconds = 5
	retryMaxWaitTimeSeconds = 15
)
```
