# pjlink-microservice
[![CircleCI](https://img.shields.io/circleci/project/byuoitav/pjlink-microservice.svg)](https://circleci.com/gh/byuoitav/pjlink-microservice) [![Apache 2 License](https://img.shields.io/hexpm/l/plug.svg)](https://raw.githubusercontent.com/byuoitav/pjlink-microservice/master/LICENSE)


[![View in Swagger](http://jessemillar.github.io/view-in-swagger-button/button.svg)](http://byuoitav.github.io/swagger-ui/?url=https://raw.githubusercontent.com/byuoitav/pjlink-microservice/master/swagger.json)

Provides a REST API to interact with PJLink capable devices. Commands
are sent in JSON format. Responses are parsed from the initial response string and returned in JSON format.

This service does not interpret PJLink responses; a separate micro-service should probably be written to provide more user-friendly mappings to PJLink commands and response codes. The complete PJLink specification can be found [here](http://pjlink.jbmia.or.jp/english/data/5-1_PJLink_eng_20131210.pdf)

## Setup
You need to set the `PJLINK_PORT` and `PJLINK_PASS` environment variables before running.

## Usage
Send a `POST` request to the `/command` endpoint with a body similar to the following:
```
{
    "address": "10.66.9.14",
    "port": "4352",
    "password": "sekret",
    "class": "1",
    "command": "INST",
    "parameter": "?"
}
```
All fields specified in the request as shown in the above example are mandatory. The corresponding response for the request above will be something like:
```
{
  "class": "1",
  "command": "INST",
  "response": [
    "11",
    "12",
    "21",
    "31",
    "32",
    "33",
    "34"
  ]
}
```
As shown above, all responses will follow the form of 'class' (string), 'command' (string), and 'response' [] (string), an array with one or more elements depending on the request command.
