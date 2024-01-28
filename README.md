# New Story Microservice Starter #

![Image](ahoy.png?raw=true)

Ahoy! This is a starter for microservices developed by New Story. It will always contain the latest code / improvements to make development as easy as possible.

## What is this repository for? ##

* This repository can be used a base for a new microservice.
* This repository can be used for testing new tools / libraries ( just make sure you don't accidentally merge ).
* The idea is that this repository will always be the latest ( most updated ) version of our microservices.
* New versions of Go can be tested here first.

## How do I get set up? ##

### Setting up your environment ###
To begin using this repository first rename your .env.example to .env by running the following command ( or via finder / your editor ):
```
cp .env.example .env
```

If this is the first time you're running the application make sure to set an APP key as well. You can create one here https://keygen.io/.

### Start your service ###
This services uses a `Makefile` that contains all commands that can be run on this repository. If you've never worked with a `Makefile` cd to the root of this project and run:
```
make
```

### Watching and debugging
Run `make watch` to start developing! Depending on which file your change a watch script will determine the action(s) that run.
* All go file changes will be gofmt'ed.
* `_mock.go` and `interface` files will be ignored.
* `golangci-lint` runs on the folder the changed file is in.
* `go test` will be triggered if you edit a test file.
* `go generate` will be called to generate your interfaces or do other magic.
* And of course the application will rebuild!

It is recommended to run the following commands before committing:
* `make test` will run tests on the entire microservice and fail immediately if a tests fails.
* `make lint` will run linting on the entire microservice and output lint errors in your console.
* `make format` will run gofmt on all code.


### Setup error

### Setting up your local database ###
When running the postgres container for the first time, the database will automatically be initialized. 

Your database will be created and all users will be setup. The only thing that will be missing are your tables.

This can be solved by running the migrations.

### Running migrations ###
By default after copying the `.env.example` to `.env` your environment variables should be good for development. If you're coming from an older version, make sure the `DB_MIGRATION_USER` and `DB_MIGRATION_PASSWORD` ENV variables are set in your .ENV file. If not: set them and restart your service.

Now run:
```
make db:migrate
```

Now you're all setup to start coding!

## Codeship, Sonarqube, Rancher and deployments ##
For deployments with Rancher, Sonarqube and Codeship there is an existing guide:
https://aanzee.atlassian.net/wiki/spaces/VDVP/pages/1394737158/Nieuwe+Microservice+opzetten

## FAQ's ##

#### I changed some .ENV values but nothing changes? ####
`.env` changes only get loaded when you restart your Docker.

## Todo's for the future ##


## Contribution guidelines ##
* Make sure new code is 100% tested. There are always exceptions, but this is the benchmark.
* All new code should be merged via a Pull Request on Bitbucket and shared in the #pull-requests channel.

## Who do I talk to? ##
For questions about this repo check the #go_aan_zee or the #pull-requests groups on Slack.
