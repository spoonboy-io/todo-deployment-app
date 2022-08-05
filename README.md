# Todo Deployment App

This is a simple Todo list application which connects to a Postgres database or cluster, intended to assist
in a Morpheus Deployments training session.

Based on the example found in [here](https://blog.logrocket.com/building-simple-app-go-postgresql/).

## Setup

Grab a release from the [releases page](https://github.com/spoonboy-io/todo-deployment-app/releases/latest), download and extract the binary

Create a Morpheus deployment with the binary 

Create a Morpheus script `postgres.env` in the same deployment folder as the binary which contains the single environment variable `PG_HOST`. All other connecion
info is hardcoded in the app for simplicity. 

```bash
# POSTGRES CONFIG
PG_HOST=54.237.172.196

# OPTIONAL TO OVERRIDE DEFAULTS
#PG_USER=               # default is postgres
#PG_PASSWORD=           # default is Password123?
#PG_DATABASE=           # default is todos
#APP_SERVER_PORT=       # default is 8090

```

## Run
Run the application like this

```bash
./todo
```

Optionally `nohup` the call or set up as a service so tasks can be created to start and stop the app in Morpheus.

View the app in your browser at: 

```
http://deploymentserver_ip:8090
```
