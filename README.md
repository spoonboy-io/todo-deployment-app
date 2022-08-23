# Todo Deployment App

This is a simple Todo list application which connects to a Postgres database or cluster, intended to assist
in a Morpheus Deployments training session.

Based on the example found in [here](https://blog.logrocket.com/building-simple-app-go-postgresql/).

Modified so that the external assets are part of the binary to simplify deployment.

## Setup

Grab a release from the [releases page](https://github.com/spoonboy-io/todo-deployment-app/releases/latest), download and extract the binary.

There are two releases (v0.5.0 & v0.7.0) which have visual differences in the UI to demonstrate deployment versioning.

Create a Morpheus deployment with the binary. 

Create a Morpheus script file `config.env` which contains the single required environment variable `PG_HOST`. 
All other connection info is hardcoded in the app for simplicity, but may be overridden here. The location of this file 
is passed as a command line flag.

```bash
# POSTGRES CONFIG
PG_HOST=54.237.172.XXX

# OPTIONAL TO OVERRIDE DEFAULTS
#PG_PORT=               # default is 5432
#PG_USER=               # default is postgres
#PG_PASSWORD=           # default is Password123?
#PG_DATABASE=           # default is todos
#APP_SERVER_PORT=       # default is 8090
```

## Run
Run the application like this:

```bash
./todo -config={absolute path to config.env}
```

Optionally `nohup` the call or set up as a service so tasks can be created to start and stop the app in Morpheus.

View the app in your browser at: 

```
http://deploymentserver_ip:8090
```
