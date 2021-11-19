
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
- [License](#license)

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

Result of `/set`
```json
{"key":"foo","value":"bar","CreatedAt":"2021-11-17T21:31:21.6085805Z","UpdatedAt":"2021-11-17T21:31:21.6085806Z"}
```

### Get Key
Gets data from parameter key with `GET` request, endpoint is `/get/{key-id}`

Result of `/get/foo`
```json
{"key":"foo","value":"bar","CreatedAt":"2021-11-17T21:31:21.6085805Z","UpdatedAt":"2021-11-17T21:31:21.6085806Z"}
```
### Flush Keys

Flush all data with `DELETE` request, endpoint is `/flush`

Result of `/flush`
```sh
Succesfully flushed all data
```

# License

MIT License

Copyright (c) 2021 Ulaş Kasım

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
