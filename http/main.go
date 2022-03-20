package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//custom writer
type logWriter struct{}

func main() {
	resp, err := http.Get("http://shriharsha.me/")
	if err != nil {
		fmt.Println("Error occured:", err)
		os.Exit(1)
	}
	// fmt.Println(resp.Header)

	// bs := []byte{}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Printed out ", len(bs), "bytes!")

	return len(bs), nil
}
