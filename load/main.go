package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getFile ( index int ) [] byte {

	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", index)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error downloading file: ", err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Skipping %d: got %d\n ",index, resp.StatusCode)
		return nil
	}

	data := new(bytes.Buffer)
	_, err = io.Copy(data, resp.Body)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading data: ", err)
		os.Exit(-1)
	}

	return data.Bytes()

}

func main(){

	var (
		output io.WriteCloser = os.Stdout
		err error
		cnt int 
		fails int 
		data [] byte
	)


	if len(os.Args) > 1{
		output , err = os.Create(os.Args[1])

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating file: ", err)
			os.Exit(-1)
		}

		defer output.Close()
	}

	fmt.Println("[")
	defer fmt.Println("]")

	for i:=1 ; fails < 2 ; i++ {
		if data = getFile(i); data == nil {
			fails++
			continue
		}

		if cnt > 0 {
			fmt.Fprintln(output, ",")
		}

		_,err = io.Copy(output, bytes.NewReader(data))

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error writing data: ", err)
			os.Exit(-1)
		}

		fails = 0
		cnt++

	}
	fmt.Fprintf(os.Stderr, "Read %d comics\n", cnt +1)

}