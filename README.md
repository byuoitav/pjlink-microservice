### pjlink-microservice

![Circle CI Build Status](https://circleci.com/gh/byuoitav/pjlink-microservice/tree/master.svg?style=shield)

### Usage
Send a `POST` request to the `/command` endpoint with a body similar to the following:
```
{
    "Address": "10.6.36.54",
    "Port": "23",
    "Class": 1,
    "Password": "stuff",
    "Command": "stuff",
    "Parameter": "stuff"
}
```
