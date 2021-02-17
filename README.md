# Skillz.io tests Techniques
Repository to hold code for technical tests.

# API

Small API that fetch datas from a [Prometheus](https://prometheus.io/) Instance with 3 exporters using [Prometheus HTTP API](https://prometheus.io/docs/prometheus/latest/querying/api/).

## Usage:
```shell
  docker-compose -d up
```
Use [Postman files](API/tests), CLI or curl to test API

## Allows a user to:
  - Login
  - Fetch Datas from :
    * All user's Exporter
    * A Specific Exporter

## It expose 3 routes:
| Routes                                | Query Parameters                                                    | Description                                                                 | Response Type                         |
|:-------------------------------------:|:-------------------------------------------------------------------:|:---------------------------------------------------------------------------:|:-------------------------------------:|
| "/login"                              | "login" {string} **MANDATORY** <br> "pass" {string} **MANDATORY**   | Log a user, claim a JWT Token for future validation and return this Token.  | string                                |
| "/query"                              | "datas" {string} **OPTIONAL**                                       | Fetch all Datas from user exporters.<br>If query parameter "datas" specified, fetch only the specified data.                                                            | JSON Object                           |
| "/query/{exporterUUID}"               | "datas" {string} **OPTIONAL**                                       | Fetch all Datas from exporter precised by {exporterUUID}.<br>If query parameter "datas" specified, fetch only the specified data | JSON Object                           |

## Mandatory
 - Environment Variables
   * **DB_DRIVER** choose between:
      - postgres
      - sqlite
   * For DB_DRIVER == "sqlite"
      * DB_NAME
   * For DB_DRIVER == "postgres"
      * DB_NAME
      * DB_HOST
      * DB_PORT
      * DB_USER
      * DB_PASSWORD
   * **API_SECRET** for JWT auth.

## Architecture
| Package            | Description                                                          |
|:------------------:|:--------------------------------------------------------------------:|
| auth               | JWT utility functions                                                |
| controllers        | API Controller (GetUser, ...)                                        |
| database           | Database With Gorm: support postgresql or sqlite                     |
| middlewares        | API Middlewares (JWTAuth)                                            |
| models             | Database Models (User, Exporter)                                     |
| router             | API Route management                                                 |
| server             | Server Type with Gin Server                                          |
| tests              | Postman test collection and environment                              |


# CLI

## Description
Small Cli to quickly test api outside Postman.

## Usage:
  - Enter cli/ directory
  - go run main.go

| Command                                           | Description                                                          |
|:-------------------------------------------------:|:--------------------------------------------------------------------:|
| "login" {login} {pass}                            | log an user                                                          |
| "query" ["exporter" {exporterUUID}] {dataToFetch} | Fetch a Data                                                         |
| "query"                                           | Fetch all Data from user                                             |



## Sprints

### Sprint 1:
- 06/02/2021 (3 hours):
    * Setup Repository
    * Setup Airtable
    * Research On subject
- 07/02/2021 (1 hour)
    * Research on Promotheus API

### Sprint 2:
- 10/02/2021 (4 hours)
    * Small poc with MVC architecture
    * Incomprehension about user management, asking more details from Client.
    * Testing Prometheus Api calls with curl
    * Setup Postman Tests.

- 13/02/2021 (7 hours)
    * Finishing routes for Poc.
    * Setup Jwt token auth.
    * Adding Postman tests.

### Sprint 3:
- 14/02/2021 (4 hours)
    - Finishing Poc with docker / Docker compose
    - Dev small cli to quickly test backend
    - Research on Cloud Management.

### Sprint 4:
- 16/02/2021 (4 hours)
  - Research on Cloud Management.
  - Schema on Cloud Architecture
- 17/02/2021 (1 hour)
  - Gitlab Repository, README and Pull request
  - "query/{exporterUUID}": Checking if exported uuid is an among user's exporters uuid
