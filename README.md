# Golang web API
Twitter trending topic text analysis

## Requirements 
- Docker
- docker-compose

## Building
- Inside the project root run `docker-compose build`

## Running
- Inside the project root run `docker-compose up`

## Migrations
- The service must be up and running.
- In console run `docker-compose exec web bee migrate --conn "root:root@tcp(db:3306)/blog?charset=utf8"`

## Usage
### Creating an user
- Make a `POST` request to `localhost:8080/v1/users` with body like the following:
```json
{
  "name": "Jon",
  "email": "jon@snow.com",
  "password": "Test.94"
}
```

### Obtain authentication token
- Make a `POST` request to `localhost:8080/v1/users/login` with body like the following (existing user):
```json
{
  "email": "jon@snow.com",
  "password": "Test.94"
}
```
- The succesful response will be like:
```json
{
  "name": "Jon",
  "email": "jon@snow.com",
  "jwt": "some_example_jwt"
}
```
- The value corresponding to `jwt` field will be your auth token to use the sentiment analysis endpoints by putting it as a header (Authorization) value to the different requests.

## Endpoints
- [GET] `/v1/trends?id={WOEID}&exclude{info_to_exclude(e.g hashtags)}`
- [GET] `/v1/trends/anaysis?q={trend_to_analyze(e.g Bears)}`
