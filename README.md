# Theater Reservation Kata

## Credits

This kata is inspired from Emily Bache's ["Theater kata"](https://github.com/emilybache/Theater-Kata). It is a Golang rewrite of the sample code from [Arnaud Thiefaine and Dorra Bartaguiz](https://github.com/athiefaine/theater-reservation-kata/tree/100-start).

## Objective

The purpose of this kata is to perform a refactoring of the [theater service](internal/service/theater.go) to illustrate DDD concepts. The refactoring is secured by a set of [approval tests](internal/service/theater_test.go) which guarantee that the original behavior is not broken.

The successive refactoring steps are described in this [Devoxx France 2023 video](https://www.youtube.com/watch?v=3jI6TpzLU1c).

## Execution

Run the sample program in [`main.go`](main.go) with

```sh
go run .
```

Run the approval tests with

```sh
go test ./...
```

If for some reason the reference test files in [`testdata` directory](internal/service/testdata) need to be updated, run tests with `UPDATE_APPROVALS=1` environment variable.
