package main

// Exercise 1.7: The function call io.Copy(dst, src) reads from src and writes to dst. Use it
// instead of ioutil.ReadAll to copy the response body to os.Stdout without requiring a
// buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.
// Exercise 1.8: Modify fetch to add the prefix http:// to each argument URL if it is missing.
// You might want to use strings.HasPrefix.
// Exercise 1.9: Modify fetch to also print the HTTP status code, found in resp.Status.

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		var dst io.Writer = os.Stdout
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		writtenBytes, err := io.Copy(dst, resp.Body)
		statusLine := fmt.Sprintf("Status code: %d\n", resp.StatusCode)
		dst.Write([]byte(statusLine))
		fmt.Printf("%d bytes\n", writtenBytes)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
