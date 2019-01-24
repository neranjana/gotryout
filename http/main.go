package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// fmt.Println(resp)

	// 1st approach
	// bs := make([]byte, 99999) //creates a byte slice with 99999 elements initialized
	// // Read fuction cannot resize the slice if the slice is full

	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// 2nd approach

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
