
- [HomeWork](#homework)
  - [Requirements](#requirements)
  - [Definitions](#definitions)
  - [Bonuses](#bonuses)
- [Yemeksepeti Golang Rest](#yemeksepeti-golang-rest)
  - [Quick Start](#quick-start)
  - [API](#api)
    - [Set Key](#set-key)
    - [Get Key](#get-key)
    - [Flush Keys](#flush-keys)

# HomeWork

REST-API service with store data as key-value in memory

## Requirements

- Endpoint to 'set' key
- Endpoint to 'get' key
- Endpoint to flush all data
- Writing all data to file at a specified interval
- If the application stops and restarts and there is data saved, it should load the data into memory
## Definitions

- Develop app with using the standart library
- All values can be string, there is no type difference
- Value can be a single data, there is no need to be list or array
## Bonuses

- Is there any design pattern? (ddd, repository)
- Is there a README file describing the application?
- Has the application been paid attention to 'go doc'?
- Is there an 'api doc'?
- Are the tests written?
- Is there a server log that shows all HTTP requests?
- Has the application contains golang coding standards?
- Is there any style guide?
- Is the application containerized?
- Is the application deployed anywhere? (heroku? gcloud? aws?)
- Does the application have requirements as if it were an open-source project?


# Yemeksepeti Golang Rest

Yemeksepeti Golang Rest is written in Go (Golang) it features store, get and flush data in standart API

## Quick Start

Yemeksepeti Golang Rest requires [Docker](https://www.docker.com/) to run.

Install docker, open it and simple start the server

```sh
  docker-compose up
```

## API
Contains all the endpoints with port `8080`
- Set Key
- Get Key
- Flush Keys

### Set Key
Sets data to key with `POST` request, endpoint is `/set`

| Value | Type   |
| ----- | ------ |
| key   | string |
| value | string |

Example json data
```json
{
    "key":"foo",
    "value":"bar"
}
```

### Get Key
Gets data from parameter key with `GET` request, endpoint is `/get/{key-id}`

Result of `/get/foo`
```json
{"key":"foo","value":"bar","CreatedAt":"2021-11-17T21:31:21.6085805Z","UpdatedAt":"2021-11-17T21:31:21.6085806Z"}
```
### Flush Keys

Flush all data with `DELETE` request, endpoint is `/flush`

