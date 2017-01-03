#Hologram-Go-Library

[![Go Report Card](https://goreportcard.com/badge/github.com/hologram-io/hologram-go)](https://goreportcard.com/report/github.com/hologram-io/hologram-go)

## Introduction
This is a library built for using our REST APIs.

## Installation

One nice thing about Go is how easy it is to work with packages.
First, run:

```go
go get github.com/hologram-io/hologram-go
```

All you need to do after that is add this snippet of code in your source:

```
import (
	<your other imports go here>
	"github.com/hologram-io/hologram-go"
	<your other imports go here>
)
```

in your main folder.

## Unit Tests
Yay! you're ready for some good testing. All of the test cases are located in the
top `test` directory. Mocked data (in JSON) are in the json directory. To run a unit test,
just type:

```
go test <whichever test here>
```

## Support
Please feel free to [reach out to us](https://hologram.io/) if you have any questions/concerns.

## Credits
Special thanks to the amazing work that @kurrik put in his
[twittergo](https://github.com/kurrik/roomanna/) library. The Errors and Parse
interface/functions were taken from there and modified for our needs.
