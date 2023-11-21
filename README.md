# Display images in iTerm2

This tool implements the [iTerm2](https://www.iterm2.com/)
image support as described
[here](https://www.iterm2.com/documentation-images.html).
It supports both local files as well as images loaded via http(s).

Be sure to
[install the latest version of iTerm2](https://www.iterm2.com/downloads.html).

![imgcat](/img/imgcat.png?raw=true "imgcat")

![imgcat](/img/imgcat-via-http.png?raw=true "imgcat via http")

![gnuplot+imgcat](/img/gnuplot1.png?raw=true "gnuplot")

![imgls](/img/imgls.png?raw=true "imgls")

## Installation

To install `imgcat`, you should have a recent version of Go (1.21+ at the time of writing this), and type:

```go
go install github.com/olivere/iterm2-imagetools/cmd/imgcat@latest
```

To install `imgls`, type:

```go
go install github.com/olivere/iterm2-imagetools/cmd/imgls@latest
```

The binaries then get installed into the `$GOPATH/bin` directory (or `$HOME/go/bin` if `GOPATH` is not set; see `go help install`), which you can add to your `$PATH`.

## License

MIT-LICENSE. See [LICENSE](http://olivere.mit-license.org/)
or the LICENSE file provided in the repository for details.
