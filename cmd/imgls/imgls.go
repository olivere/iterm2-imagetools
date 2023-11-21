// This tool implements the iTerm2 image support as described here:
// https://www.iterm2.com/documentation-images.html
//
// Be sure to install a latest version of iTerm2 (e.g. 3.2.0 or later).

package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	width  = flag.String("width", "", "width (e.g. 100px, 10%, or auto)")
	height = flag.String("height", "", "height (e.g. 100px, 10%, or auto)")
	size   = flag.String("size", "", "width,height in pixels (e.g. 1024px,768px or 3,3)")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	for _, pattern := range flag.Args() {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			log.Fatal(err)
		}

		for _, filename := range matches {
			// Skip errors and directories
			if fi, err := os.Stat(filename); err != nil || fi.IsDir() {
				continue
			}

			f, err := os.Open(filename)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			if err := display(filename, f); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func display(filename string, r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	b64filename := base64.StdEncoding.EncodeToString([]byte(filename))

	width, height := widthAndHeight()
	if width == "" && height == "" {
		width = "3"
		height = "3"
	}

	fmt.Print("\033]1337;")
	fmt.Printf("File=inline=1;preserveAspectRatio=1;name='%s'", b64filename)
	if width != "" || height != "" {
		if width != "" {
			fmt.Printf(";width=%s", width)
		}
		if height != "" {
			fmt.Printf(";height=%s", height)
		}
	}
	fmt.Print(":")
	fmt.Printf("%s", base64.StdEncoding.EncodeToString(data))
	fmt.Print("\a")
	fmt.Printf("\033[A%s\n", filename)

	return nil
}

func widthAndHeight() (w, h string) {
	if *width != "" {
		w = *width
	}
	if *height != "" {
		h = *height
	}
	if *size != "" {
		sp := strings.SplitN(*size, ",", -1)
		if len(sp) == 2 {
			w = sp[0]
			h = sp[1]
		}
	}
	return
}

func usage() {
	fmt.Fprint(os.Stderr, "usage: imgls [flags] filename\n")
	flag.PrintDefaults()
	os.Exit(2)
}
