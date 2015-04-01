// Reads an input tar ball into memory, compress it into a gzip stream. Write that compressed gzip stream onto a buffer.
//Takes a gzip reader on that buffer and uncompress the gzip stream. Write the original tar ball back onto the disk.

//NOTE: To run this program please use your own tar ball.

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// Get a file descriptor on the tar ball.
	file, err := os.Open("/home/shishir/docker/src/github.com/docker/docker/img.tar")
	if err != nil {
		log.Fatalln(err)
	}

	//Create a buffer for writing gzip content.
	buf := new(bytes.Buffer)

	//Create a gzip writer.
	gw := gzip.NewWriter(buf)

	//Read the file into byte slice.
	inputBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	//Write the byte slice of the entry into the buffer (in gzip format)
	gw.Write(inputBytes)

	gw.Close()

	fmt.Printf("The size of the buffer is %d\n", buf.Len())
	err = ioutil.WriteFile("/home/shishir/Selfhelp/sample_codes/untar_dir/repo.tar.gz", buf.Bytes(), 0666)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("GZIP written successfully")

	gr, err := gzip.NewReader(buf)

	if err != nil {
		log.Fatalln(err)
	}
	target_file, err := os.Create("/home/shishir/Selfhelp/sample_codes/untar_dir/repo.tar")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = io.Copy(target_file, gr)
	if err != nil {
		log.Fatalln(err)
	}
}
