#Hologram-Go-Tests

## Introduction
Here are the unit tests for our Hologram Go library. We wanted to split this up
from the main library since it further encapsulates the package to the core
functionalities without having any testing related stuff within that package.

All unit tests can be done here without interfacing with the main package.

## Installation

1. If you haven't already installed the Hologram Go library, go on your favorite terminal
and type:

```
go get github.com/hologram-io/hologram-go
```

2. Once you are done with that, type this:

```
go get github.com/hologram-io/hologram-go-tests
```

3. Yay! you're ready for some good testing. All of the test cases are located in the
top directory. Mocked data (in JSON) are in the json directory. To run a unit test,
just type:

```
go test <whichever test here>
```

For more info about the Hologram Go library, please go [here](https://github.com/hologram-io/hologram-go) to learn more about it.

## Support
Please feel free to reach out to us!
