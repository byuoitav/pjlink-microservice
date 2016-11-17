# wso2jwt
[![CircleCI](https://img.shields.io/circleci/project/byuoitav/wso2jwt.svg)](https://circleci.com/gh/byuoitav/wso2jwt) [![Codecov](https://img.shields.io/codecov/c/github/byuoitav/wso2jwt.svg)](https://codecov.io/gh/byuoitav/wso2jwt) [![Apache 2 License](https://img.shields.io/hexpm/l/plug.svg)](https://raw.githubusercontent.com/byuoitav/wso2jwt/master/LICENSE)

Validating [JWT tokens](https://jwt.io/) with BYU's implementation of [WSO2](http://wso2.com/products/api-manager/) isn't the easiest thing on the planet. We wanted to make sure that we protected every route with a JWT check and, as part of that check, pull in the latest signing key.

## Dependencies
- [Echo](https://labstack.com/echo)

## Installation
```
go get github.com/byuoitav/wso2jwt
```

## Usage
```
import "github.com/byuoitav/wso2jwt"
```
```
func main() {
  port := ":8000"
  router := echo.New()
  router.Use(wso2jwt.ValidateJWT())

  router.Run(fasthttp.New(port))
}
```

## Local Development
The WSO2 JWT checking "pipeline" will not work from your local box. To enable local development without headaches, be sure to set the `LOCAL_ENVIRONMENT` variable in you shell environment.

Example (using `/etc/environment`):
```
export LOCAL_ENVIRONMENT="true"
```

## Notes
If you're not interested in protecting every single route (putting [`/health`](https://github.com/jessemillar/health) behind authentication just seems silly), you can use [Echo Groups](https://echo.labstack.com/middleware/overview) instead of `router.Use()`:
```
router.Group("/command", wso2jwt.ValidateJWT())
```
