# Flightify
A flight path tracker that helps express how a particular traveler's flight path may be queried.

### Set Up and Run

you can start Flightify in a terminal window or by building and running a docker container. 
```bash
make start
```
or

```bash
make docker-start
```

Flightify will listen on port 8080 by default. Optionally, you can override the default configuration via the [config.json](./config.json) file. The config file ([config.json](./config.json)) is automatically passed as an argument to the `make start` and `make docker-start` commands (see [setup and run](#set-up-and-run)) and must be present in the project root to build.  See [config.json](./config.json) for an example of what values can be overridden.

### Test

run tests (quick):

```bash
make testq
```

run tests (verbose):

```bash
make test
```

run code coverage (creates hidden .coverage directory in project root)

```bash
make covero
```

### API
API documentation can be found [here](./docs/markdown.md)