# Display images in iterm2

This tool implements the iterm2 image support as described
[here](http://iterm2.com/images.html). It supports both local
files as well as images loaded via http(s).

Be sure to install iterm2 nightly.

![imgcat](/img/imgcat.png?raw=true "imgcat")

![imgcat](/img/imgcat-via-http.png?raw=true "imgcat via http")

![gnuplot+imgcat](/img/gnuplot1.png?raw=true "gnuplot")

![imgls](/img/imgls.png?raw=true "imgls")

# Installation

To install `imgcat`, type:

```go
go get -u github.com/olivere/iterm2-imagetools/cmd/imgcat
```

To install `imgls`, type:

```go
go get -u github.com/olivere/iterm2-imagetools/cmd/imgls
```

# License

MIT-LICENSE. See [LICENSE](http://olivere.mit-license.org/)
or the LICENSE file provided in the repository for details.
