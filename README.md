Welcome to the Email-Deduper!
----------------------------
_Your emails de-duped--for less!_

The goal of this application is to take an input of 100,000 emails with 50% of duplicates and remove the duplicates in under 1 second. Order also must be preserved for the emails.

In addition, this application allows configuring outside of the default 100,000 emails and 50% duplicates by setting the values mentioned below.

This application can be ran both using Go or Docker, depending on which you have installed on your machine.

## Configuration

The following variables can be changed in the `config.yml` file:

| Variable             | Default      | Type    | Description                                             |
| -------------------- | ------------ | ------- | ------------------------------------------------------- |
| GENERATE_EMAIL_COUNT | 100000       | `int`   | Determines the number of emails to generate             |
| DUPLICATE_PERCENTAGE | 0.5          | `float` | Determines the percentage of emails that are duplicates |

## Run the application

### Using Go Directly

In the terminal, just run:

`go run main.go emailgenerator.go config.go`

This will output a look with the amount of time in took to run the application, for example:

```
2020/10/28 15:16:51 Took 15.374993ms to remove 50000 duplicate emails from 100000 emails
```

### Using Docker

If you have Docker installed, you can also run the application by running:

`docker-compose up`

Doing so will output the logs once the script is ran, as seen below:

```
Successfully built 09a7d67c7cc5
Successfully tagged email-deduper_email-deduper:latest
Recreating email-deduper_email-deduper_1 ... done
Attaching to email-deduper_email-deduper_1
email-deduper_1  | 2020/10/29 01:36:17 Took 484.8986ms to remove 50000 duplicate emails from 100000 emails
email-deduper_email-deduper_1 exited with code 0
```

If you need to rebuild the application (for example if there are any code changes) you can run the command below:

`docker-compose up --build`

## Benchmarks

To run a benchmark against the dedupe function with the default 100,000 emails and 50% duplicates, run the following command:

`go test -bench=.`

You should get output like the following:

```
goos: darwin
goarch: amd64
pkg: github.com/zhongruige/email-deduper
BenchmarkDeDupe-12             3         394592342 ns/op
PASS
ok      github.com/zhongruige/email-deduper     3.267s
```

If you'd like to adjust the benchmarks to test performance for different inputs, simply adjust the values in `deduperbenchmark_test.go`.

## Unit Tests

To execute all the unit tests, run:

`go test`

For verbose output while the tests run, add in `-v`:

`go test -v`