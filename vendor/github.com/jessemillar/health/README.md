# health-check
A small Go package for quickly adding a `/health` endpoint to RESTful APIs

## Dependencies
- [Echo](https://labstack.com/echo)

## Installation
```
go get github.com/jessemillar/health
```

## Usage
```
import "github.com/jessemillar/health"
```

```
func main() {
  port := ":8000"
  router := echo.New()

  router.Get("/health", health.Check)

  router.Run(fasthttp.New(port))
}
```
