# Suave (U - Username Av - Availability Se - Service)

### Microservice that checks whether a username is available for use or not using redis-bloom in-memory database.

### Pre-reqs

* Golang 1.16.5
* Docker
* make

### Steps

1. Clone the repo
2. Run `make test` so that redis-bloom image is pulled and a local dev image is built. This will also run test-cases for redis-bloom DB client.
3. Run `make seed` which will result in generation of 10,000 usernames & will be inserted into redis-bloomfilter namely `username`.
4. Run `make build` to generate application binary. Binary will be placed in `projectFolder/bin/suave`.
5. Run this command to execute binary.
> ./bin/suave
6. A `seed.txt` file will be created as a result of **Step 3** in your project folder holding usernames generated. Pick one and you can test it via
> curl --location --request GET 'http://localhost:8080/v1/username/availability?username=johndoe123'
7. To check health of services running, use this command:
> curl --location --request GET 'http://localhost:8080/v1/health'

