//Reads a tar ball by browsing its entries, and compress it into a gzip format.

//NOTE: Please use your own tar ball to run this program.

package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/home/shishir/docker/src/github.com/docker/docker/img.tar")
	buf := new(bytes.Buffer)
	gw := gzip.NewWriter(buf)

	if err != nil {
		log.Fatalln(err)
	}

	tr := tar.NewReader(file)

	for {

		_, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		inputBytes, err := ioutil.ReadAll(tr)
		if err != nil {
			log.Fatalln(err)
		}
		gw.Write(inputBytes)
	}

	gw.Close()
	fmt.Printf("The size of the buffer is %d\n", buf.Len())
	err = ioutil.WriteFile("/home/shishir/Selfhelp/sample_codes/untar_dir/repo.tar.gz", buf.Bytes(), 0666)
	if err != nil {
		log.Fatalln(err)
	}

	b := make([]byte, buf.Len())
	gr, err := gzip.NewReader(buf)
	if err != nil {
		log.Fatalln(err)
	}
	gr.Read(b)
	fmt.Printf("The length of the byte slice is %d\n", len(b))

	fmt.Println("GZIP written successfully")
}
