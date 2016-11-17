# json-response

JSON wants returned values to be in key-value format. This package removes a couple lines of code I got tired of copypasta-ing between projects.

## Dependencies
- [Echo](https://labstack.com/echo)

## Installation
```
go get github.com/jessemillar/jsonresp
```

## Usage
```
import "github.com/jessemillar/jsonresp"
```
```
func SampleFunction(context echo.Context) error {
	err := functionThatReturnsAnError()
	if err != nil {
		return jsonresp.New(context, http.StatusBadRequest, "Could not read request body: "+err.Error())
	}
}
```
