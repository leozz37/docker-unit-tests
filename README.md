# Running Unit Tests inside a Docker Container

![Docker](https://miro.medium.com/max/700/0*gbfua3eCsElT77Eb.png)

Check my article [here](https://medium.com/@leonardoaugusto287/running-unit-tests-inside-a-docker-container-ec68c2274522).

## Running the examples

Go to your language of choice directory, and run the following commands:

```bash
$ docker build . -t example

$ docker run example sh -c "TEST COMMAND"
```

On "TEST COMMAND", swap for your language test command.

### Golang

```bash
$ docker run example sh -c "go test"
```

### Python

```bash
$ docker run example sh -c "python main_test.py"
```

## Contribuiting

Feel free to add any other language example.
