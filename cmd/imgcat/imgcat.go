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
	"net/http"
	"os"
	"strings"
)

var (
	width               = flag.String("width", "", "width (e.g. 100px, 10%, or auto)")
	height              = flag.String("height", "", "height (e.g. 100px, 10%, or auto)")
	size                = flag.String("size", "", "width,height in pixels (e.g. 1024px,768px or 3,3)")
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
			if strings.HasPrefix(filename, "http://") || strings.HasPrefix(filename, "https://") {
				displayNetPicture(filename)
			} else {
				// Skip errors and directories
				if fi, err := os.Stat(filename); err != nil || fi.IsDir() {
					continue
				}

				displayLocalPicture(filename)
			}
		}
	}
}

func displayNetPicture(filename string) {
	res, err := http.Get(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if err := display(res.Body); err != nil {
		log.Fatal(err)
	}
}

func displayLocalPicture(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := display(f); err != nil {
		log.Fatal(err)
	}
}

func display(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	width, height := widthAndHeight()

	fmt.Print("\033]1337;")
	fmt.Printf("File=inline=1")
	if width != "" || height != "" {
		if width != "" {
			fmt.Printf(";width=%s", width)
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
		sp := strings.SplitN(*size, ",", -1)
		if len(sp) == 2 {
			w = sp[0]
			h = sp[1]
		}
	}
	return
}

func usage() {
	fmt.Fprint(os.Stderr, "usage: imgcat [flags] filename\n")
	flag.PrintDefaults()
	os.Exit(2)
}
