# Flightify
A flight path tracker that helps express how a particular traveler's flight path may be queried.

### Set Up and Run

you can start Flightify in a terminal window or by building and running a docker container. 
```bash
make start args=<list args here. seperate by space.>
```
or

```bash
make docker-start
```

Flightify will listen on port 8080 by default. Optionally, you can provide a configuration file (.yaml or .json) to override the default configuration. see [config.json](./config.json) for an example of what values can be overridden.

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