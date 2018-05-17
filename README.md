<h1 align="center">
  <br>
  <br>
  Distance-Calculator
  <br>
  <br>
</h1>

<h4 align="center">Go package that calculates the distance between two lat/long points</h4>

<p align="center">
  <a href="https://travis-ci.org/pedrolopesme/distance-calculator"> <img src="https://api.travis-ci.org/pedrolopesme/distance-calculator.svg?branch=master" /></a>
  <a href="https://goreportcard.com/report/github.com/pedrolopesme/distance-calculator"> <img src="https://goreportcard.com/badge/github.com/pedrolopesme/distance-calculator" /></a>
  <a href="https://codeclimate.com/github/pedrolopesme/distance-calculator/maintainability"> <img src="https://api.codeclimate.com/v1/badges/2623b16f41d3a69fba1c/maintainability" /></a>
  <a href="https://godoc.org/github.com/pedrolopesme/distance-calculator"> <img src="https://img.shields.io/badge/Check%20the-GoDocs-1f425f.svg" /></a>
</p>
<br>

### Supported Conversion Units

| Meters | Kilometers | Miles | Nautical Miles |
|---|---|---|---|  
| CalcMeters() | CalcKilometers() | CalcMiles() | CalcNauticalMiles() |   
 
### Makefile

This project provides a Makefile with all common operations need to develop and test

* setup: installs govendor package via go get
* vendoring: vendorizing your application
* test: runs all tests
* fmt: runs gofmt for all go files


### Running tests

Tests were write using [Testify](https://github.com/stretchr/testify). In order to run them, just type:

```shell
$ make test
```

### Credits

These are the main external packages that make up distance-calculator:

| packages | description |
|---|---|
| **[testify](https://github.com/stretchr/testify)** | **A toolkit with common assertions and mocks that plays nicely with the standard library** |


### License

[MIT](LICENSE.md)