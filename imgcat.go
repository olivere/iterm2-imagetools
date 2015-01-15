// This tool implements the iterm2 image support as described here:
// http://iterm2.com/images.html
//
// Be sure to install iterm2 nightly.

package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	LEADIN = "\033]1337;"
)

var (
	width               = flag.String("width", "", "width (e.g. 100px, 10%, or auto)")
	height              = flag.String("height", "", "height (e.g. 100px, 10%, or auto)")
	size                = flag.String("size", "", "widthxheight in pixels (e.g. 1024x768)")
	preserveAspectRatio = flag.Bool("p", false, "preserve aspect ratio")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		// Read from stdin
		if err := display(os.Stdin); err != nil {
			log.Fatal(err)
		}
	} else {
		for _, filename := range flag.Args() {
			// Skip errors and directories
			if fi, err := os.Stat(filename); err != nil || fi.IsDir() {
				continue
			}

			f, err := os.Open(filename)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			if err := display(f); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func display(r io.Reader) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	width, height := widthAndHeight()

	fmt.Print(LEADIN)
	fmt.Printf("File=inline=1")
	if width != "" || height != "" {
		fmt.Printf(";")
		if width != "" {
			fmt.Printf("width=%s", width)
		}
		if height != "" {
			fmt.Printf(";height=%s", height)
		}
	}
	if *preserveAspectRatio {
		fmt.Print("preserveAspectRatio=1")
	}
	fmt.Print(":")
	fmt.Printf("%s", base64.StdEncoding.EncodeToString(data))
	fmt.Print("\a\n")

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
		sp := strings.SplitN(*size, "x", -1)
		if len(sp) == 2 {
			if i, err := strconv.ParseInt(sp[0], 10, 64); err == nil && i > 0 {
				w = fmt.Sprintf("%dpx", i)
			}
			if i, err := strconv.ParseInt(sp[1], 10, 64); err == nil && i > 0 {
				h = fmt.Sprintf("%dpx", i)
			}
		}
	}
	return
}

func usage() {
	fmt.Fprint(os.Stderr, "usage: imgcat [flags] filename\n")
	flag.PrintDefaults()
	os.Exit(2)
}
