// Program urlencode escapes stdin so it can be safely used in a HTTP URL parameter.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
)

var flagReverse = flag.Bool("u", false, "unescape instead")

func main() {
	flag.Parse()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *flagReverse {
		u, err := url.QueryUnescape(string(b))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Print(u)
	} else {
		fmt.Print(url.QueryEscape(string(b)))
	}
}
