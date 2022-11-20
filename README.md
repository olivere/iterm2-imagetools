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

# Installation

To install `imgcat`, type:

```go
GO111MODULE=on go install github.com/olivere/iterm2-imagetools/cmd/imgcat@latest
```

To install `imgls`, type:

```go
GO111MODULE=on go install github.com/olivere/iterm2-imagetools/cmd/imgls@latest
```

# License

MIT-LICENSE. See [LICENSE](http://olivere.mit-license.org/)
or the LICENSE file provided in the repository for details.
