### pjlink-microservice

![Circle CI Build Status](https://circleci.com/gh/byuoitav/pjlink-microservice/tree/master.svg?style=shield)

### Usage
Send a `POST` request to the `/command` endpoint with a body similar to the following:
```
{
    "Address": "10.66.9.14",
    "Port": "4352",
    "Password": "sekret",
    "Class": "1",
    "Command": "INST",
    "Parameter": "?"
}
```
