# Farmer

Farmer is a simple PaaS wrapped around `docker-compose` to create, deploy and manage small projects.

## Installation
As simple as running docker compose and passing a root password for your MySql server.
```sh
export DATABASE_ROOT_PASSWORD=yourRandomRootPassword
docker-compose up -d
```

If you want to run a Farmer instance on a remote docker engine you need to configure docker compose using environment variables.

## Usage
To talk to a farmer instance you need to run `docker-compose`. (If your docker engine is on a remote server you need to set appropriate environment variables)

## Features
* HAProxy load balancer to handle http/https traffic
* Easy domain management
* MySql service broker
