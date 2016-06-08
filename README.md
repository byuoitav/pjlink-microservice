# pjlink-microservice [![CircleCI](https://circleci.com/gh/byuoitav/pjlink-microservice.svg?style=svg)](https://circleci.com/gh/byuoitav/pjlink-microservice)

Provides a RESTful micro-service to interact with PJLink capable devices. Commands
are sent in JSON format. Responses are parsed from the initial response string and returned in JSON format. 

This service does not interpret PJLink responses; a separate micro-service should probably be written to provide more user-friendly mappings to PJLink commands and response codes. The complet PJLink specification can be found [here](http://pjlink.jbmia.or.jp/english/data/5-1_PJLink_eng_20131210.pdf)

## Usage
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
