# utils-go

![GitHub last commit](https://img.shields.io/github/last-commit/azam-akram/utils-go)
![GitHub license](https://img.shields.io/github/license/azam-akram/utils-go)

`utils-go` is a collection of utility functions and packages written in Go (Golang) to simplify and enhance common tasks in Go programming. This repository aims to provide reusable code and tools that can be easily integrated into your Go projects, saving you development time and effort.

### [Json utils](https://github.com/azam-akram/utils-go/tree/main/json_utils):

Details about HTTP json-utils can be found [here](https://solutiontoolkit.com/2023/01/the-common-json-utility-functions-in-go-language/)

```
type JsonHandler_Interface interface {
	ConvertGenericInterfaceToMap()
	ConvertStringToMap(s string) (map[string]interface{}, error)
	ConvertMapToString(m map[string]interface{}) (string, error)
	ConvertStringToStruct(s string) (*Employee, error)
	ConvertStructToString(e *Employee) (string, error)
	ConvertStringToByte(s string) []byte
	ConvertByteToString([]byte) (string, error)
	ConvertByteToStruct(jsonBytes []byte) (*Employee, error)
	ConvertStructToByte(emp *Employee) (jsonBytes []byte, err error)
	ModifyInputJson(s string) (map[string]interface{}, error)
	DisplayAllJsonHandlers()
}
```
#### How to use:
Have a look at [DisplayAllJsonHandlers()](https://github.com/azam-akram/utils-go/blob/85de9b1f6804834765c9b0320d00ad944cac7b75/json_utils/json_handler.go#L54) to know how to use other functions exposed by `JsonHandler_Interface` interface. If you want to call all of these functions, simply call `DisplayAllJsonHandlers()`
```
jsonHandler := json_utils.JsonHandler{}
jsonHandler.DisplayAllJsonHandlers()
```

### [HTTP utils](https://github.com/azam-akram/utils-go/tree/main/http_utils):

Details about HTTP client-server can be found [here]([https://solutiontoolkit.com/2023/01/a-retryable-http-client-server-communication-in-go-language/])

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

### [Customised Logger](https://github.com/azam-akram/utils-go/blob/main/logger/logger.go):
A customised logger, which can log different data structures like, string, key-value map etc.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute the code in this repository as long as you retain the original license terms. Refer to the [LICENSE](LICENSE) file for more information.

---

Thanks for using `utils-go`! If you have any questions or need further assistance, feel free to contact the project maintainer or open an issue on the repository. Happy coding!
