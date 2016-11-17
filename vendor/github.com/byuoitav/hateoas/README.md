# hateoas
[![CircleCI](https://img.shields.io/circleci/project/byuoitav/hateoas.svg)](https://circleci.com/gh/byuoitav/hateoas) [![Codecov](https://img.shields.io/codecov/c/github/byuoitav/hateoas.svg)](https://codecov.io/gh/byuoitav/hateoas) [![Apache 2 License](https://img.shields.io/hexpm/l/plug.svg)](https://raw.githubusercontent.com/byuoitav/hateoas/master/LICENSE)

[HATEOAS](http://timelessrepo.com/haters-gonna-hateoas) as a concept is awesome. `hateoas` is Go package to populate HATEOAS responses for RESTful APIs with proper Swagger documentation.

### Dependencies
- [YAML](https://github.com/go-yaml/yaml/tree/v2)
- [Echo](https://labstack.com/echo) (for `RootResponse`)

### Quickstart
```
import "github.com/byuoitav/hateoas"
```

In your main Go file (probably something like `server.go`), load your Swagger YAML file into the HATEOAS library:
```
err := hateoas.Load("https://raw.githubusercontent.com/byuoitav/av-api/master/swagger.yaml")
if err != nil {
	log.Fatalln("Could not load swagger.json file. Error: " + err.Error())
}
```

```
func main() {
  port := ":8000"
  e := echo.New()

  e.Get("/", hateoas.RootResponse)

  e.Run(fasthttp.New(port))
}
```

In your to-be-returned structs, be sure to include a `Links` attribute similar to the following:
```
type Building struct {
	Links    []hateoas.Link `json:"links,omitempty"`
	Building string         `json:"building"`
}
```

In your code, before returning your structs, add the HATEOAS links:
```
links, err := hateoas.AddLinks(c, []string{allBuildings.Buildings[i].Building})
if err != nil {
	return c.JSON(http.StatusBadRequest, helpers.ReturnError(err))
}

allBuildings.Buildings[i].Links = links
```

### Syntax
`hateoas.AddLinks` takes two arguments. The first is a string representing the current HTTP path (EG: `/endpoint/:variable`). The second is an array of strings specifying values for URL parameters used in the HTTP path. The strings will be used by `hateoas` to populate URLs for downstream endpoints.
```
func AddLinks(path string, parameters []string) ([]Link, error) {
  ...
}
```
