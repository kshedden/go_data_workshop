Data processing in Go
---------------------

These are the materials for a CSCAR workshop on using
[Go](http://golang.org) for basic data processing, focusing on
manipulating text files.  No prior exposure to Go is expected, however
note that Go may not be the easiest language to start with if you have
never programmed before.  People familiar with scripting languages
such as Python should find most aspects of Go to be fairly
straightforward.

Go is a new language that is being adopted by organizations that
process large volumes of data.  It is also used in data center
management and as a back-end language for web applications.  We will
not try to characterize Go as a language here, or say much about its
possible use cases.  If you are interested in this topic, you may want
to watch these presentations by one of the primary designers of Go:
[Another Go at language
design](https://www.youtube.com/watch?v=7VcArS4Wpqk), [Simplicity is
complicated]( https://www.youtube.com/watch?v=rFejpH_tAHM).

## Installation and setup

The Go project and tools are both open source and community-driven.
To compile and run Go code on your computer, you will need to download
and install the "Go tool", which can be found
[here](https://golang.org/dl).  More detailed installation
instructions are [here](https://golang.org/doc/install).
Alternatively, if you are at UM and are using Flux, type "module load
go" at the shell prompt to make the Go tool available.

Go source files can be written in a text editor, or using various
IDE's (integrated development environments).  There is no official or
predominant IDE or editor for Go.  [This
screencast](https://www.youtube.com/watch?v=XCsL89YtqCs) walks through
the standard uses of the Go tool.

## Compiling and running a simple Go program

Go scripts can be placed anywhere in your file system.  Go packages
that are able to be used by other Go programs should be placed in your
GOHOME directory, which defaults to the go directory in the top level
of your home directory, i.e. ~/go.

Go source files are text files with suffix ".go".  A very simple Go
program is:

```
package main

import "fmt"

func main() {
    fmt.Printf("A Go program...\n")
}
```

Suppose you save this in a file called `simple.go` in your current
working directory.  You should be able to run it using the command `go
run simple.go`.  Alternatively, you can compile it to an executable
using `go build simple.go`, then run the executable using `simple`.

Difficulties with any of the steps above are likely due to a missing
or incomplete installation of the Go tools, or misconfigured
environment variables such as PATH.

## Resources

The [effective go](https://golang.org/doc/effective_go.html) document
is a complete overview of the Go language.  A somewhat friendlier
introduction to the language is the [Go
tour](https://tour.golang.org/welcome/1).
