# Dayatani Farmer API

## Getting started

Service of farmer

## Prerequisities
1. Docker
2. Golang >=1.20
3. PostgreSQL >=15

## How to run
1. Create `.env` file
   ```sh
    NAME="Dayatani Farmer API" # The service name. The default value is Dayatani Farmer API. The value is required.
    PORT=PORT # The service running port. The default value is 8080

    DB_HOST=DB_HOST # The host of PostgreSQL database. The value is required
    DB_PORT=DB_PORT # The port of PostgreSQL database. The value is required
    DB_USERNAME=DB_USERNAME # The auth username of the PosgreSQL database. The value is required
    DB_PASSWORD=DB_PASSWORD # The auth password of the PosgreSQL database. The value is required
    DB_NAME=DB_NAME # The name of the PosgreSQL database. The value is required

    DB_SSL_MODE=DB_SSL_MODE # The SSL mode of PostgreSQL connection. Please refer to https://www.postgresql.org/docs/current/libpq-ssl.html.
    DB_STATEMENT_TIMEOUT=DB_STATEMENT_TIMEOUT # The query statement timeout
    DB_TRANSACTION_SESSION_TIMEOUT=DB_TRANSACTION_SESSION_TIMEOUT # The transaction session timeout
    DB_SESSION_TIMEOUT=DB_SESSION_TIMEOUT # The session timeout

    DB_LOG_MODE=DB_LOG_MODE # The log mode of the database. If the value is true, the query is logged. If the otherwise, the query isn't logged. The default value is false
    DB_SLOW_THRESHOLD=DB_SLOW_THRESHOLD # The threshold of slow query. If the running query took the time greater than the threshold, the query is logged as "SLOW QUERY".

    DB_RETRY=DB_RETRY # The retry number of db connection
    DB_WAIT_SLEEP=DB_WAIT_SLEEP # The sleep duration between connection retrying.

    LOG_MODE=LOG_MODE # The log mode of the service

    HASHED_COST=HASHED_COST # The bcrypt cost

    HASHED_AUTH_USERNAME=HASHED_AUTH_USERNAME # The bcrypt hashed username of basic authentication header
    HASHED_AUTH_PASSWORD=HASHED_AUTH_PASSWORD # The bcrypt hashed password of basic authentication header

    BASE64_ENCODING_TYPE=BASE64_ENCODING_TYPE # The base64 encoding type. The default value is std. The base64 is used for decode the requested basic authentication.
   ```
2. Migrate the database by
   ```sh
   make db DOCKER= SQLPATH= ENVFILENAME=

   # Note:
   # DOCKER= // If the migration process use docker, please set as true.
   # SQLPATH= // Set the value if the DOCKER parameter is true
   # ENVFILENAME= // Set the value if the DOCKER parameter is true. The default value is .env.
   ```
   Note:
   The command `make db` automatically create the docker image (`make db_init`) and run the migration (`make db_migrate`) if the `DOCKER` parameter is `true`

3. Run the service by
   ```sh
   make service DOCKER= PORT= ENVFILENAME= TEST_RUNNING= LINT_RUNNING=

   # Note:
   # DOCKER= // If the migration process use docker, please set as true.
   # SQLPATH= // Set the value if the DOCKER parameter is true
   # ENVFILENAME= // Set the value if the DOCKER parameter is true. The default value is .env.
   # TEST_RUNNING= // The value is used if the DOCKER parameter is true. If the value true, the docker building run the test files. The default value is false.
   # LINT_RUNNING= // The value is used if the DOCKER parameter is true. If the value true, the docker building check the linter. The default value is false.
   ```
   Note:
   1. The command `make app` automatically create the docker image (`make app_init`) and run the migration (`make app_run`) if the `DOCKER` parameter is `true`.
   2. If the `TEST_RUNNING` value is true and the testing is failed, the build process will be aborted.
   3. If the `LINT_RUNNING` value is true and the linter checking process is failed, the build process will be aborted.

## API Docs

### Authorization
The authorization use the Basic Authentication type. The values are username and password. The username and password put in the header `Authorization` must be the raw version of `HASHED_AUTH_USERNAME` and `HASHED_AUTH_PASSWORD`. The username and password are joined by ':' and encoded to base64 by the `BASE64_ENCODING_TYPE`.

The format of the `Authorization` header is `Basic <Base64>`.

If the authentication is invalid, the endpoint will return

```json
// Response (401)
{
    "ok": false,
    "message": "invalid token",
    "data": {
        "error": "invalid token"
    }
}
```

### Get Farmers List (GET /farmers)
#### Request
##### Headers
1. Authorization: `Basic <Base64>`

##### Query Params
1. `limit`. The number of returned items. The default value is 10
2. `offset`. The number of skipped items in the database. The default value is 0.
3. `sorts`. The list of sort columns separated by comma. The values are
    - `name`. The returned items are sorted by name ASC
    - `-name`. The returned items are sort by name DESC
   example: `?sorts=name,-name`

#### Response
##### 200
```json
{
    "ok": true,
    "message": "Success",
    "data": {
        "count": 0,
        "limit": 0,
        "offset": 0,
        "sorts": [
            ""
        ],
        "items": [
            {
                "id": 0,
                "name": "",
                "birth_date": "YYYY-MM-DD"
            }
        ]
    }
}
```

##### 400
```json
{
    "ok": false,
    "message": "",
    "data": {
        "error": ""
    }
}
```

##### 500
```json
{
    "ok": false,
    "message": "",
    "data": {
        "error": ""
    }
}
```

### Get Farmer Detail (GET /farmers/:id)
#### Request
##### URI Params
1. `id`. The ID of retrieving farmer detail
##### Headers
1. Authorization: `Basic <Base64>`

#### Response
##### 200
```json
{
    "ok": true,
    "message": "Success",
    "data": {
        "id": 0,
        "name": "",
        "birth_date": "YYYY-MM-DD"
    }
}
```

##### 400
```json
{
    "ok": false,
    "message": "",
    "data": {
        "error": ""
    }
}
```

##### 500
```json
{
    "ok": false,
    "message": "",
    "data": {
        "error": ""
    }
}
```

### Create Farmer (POST /farmers)
#### Request
##### Headers
1. Authorization: `Basic <Base64>`
2. Content-Type: 'application/json'
##### Body
```json
{
    "name": "", // required
    "birth_date": "YYYY-MM-DD" // required. The format must be YYYY-MM-DD
}
```
#### Response
##### 200 (No Content)

##### 500
```json
{
    "ok": false,
    "message": "",
    "data": {
        "error": ""
    }
}
```

### Update Farmer Detail (PUT /farmers/:id)
#### Request
##### URI Params
1. `id`. The ID of retrieving farmer detail
##### Headers
1. Authorization: `Basic <Base64>`
2. Content-Type: 'application/json'
##### Body
```json
{
    "name": "", // required
    "birth_date": "YYYY-MM-DD" // required. The format must be YYYY-MM-DD
}
```
#### Response
##### 200 (No Content)

##### 400
```json
{
    "ok": false,
    "message": "",
    "data": {
        "error": ""
    }
}
```

##### 500
```json
{
    "ok": false,
    "message": "",
    "data": {
        "error": ""
    }
}
```

### Delete Farmer Detail (DELETE /farmers/:id)
#### Request
##### URI Params
1. `id`. The ID of retrieving farmer detail
##### Headers
1. Authorization: `Basic <Base64>`

#### Response
##### 200 (No Content)

##### 400
```json
{
    "ok": false,
    "message": "",
    "data": {
        "error": ""
    }
}
```

##### 500
```json
{
    "ok": false,
    "message": "",
    "data": {
        "error": ""
    }
}
```

### Show Swagger Documentation (GET /docs/index.html)
#### Response
##### 200
Return the swagger page