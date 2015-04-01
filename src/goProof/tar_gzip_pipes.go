//Reads a tar ball, and compress it into a gzip file using io.Pipe()
//Go example to show the power of unix like pipes.

package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	file, err := os.Open("/home/shishir/Selfhelp/sample_codes/untar_dir/img.tar")
	if err != nil {
		log.Fatalln(err)
	}

	pRdr, pWrtr := io.Pipe()
	gw := gzip.NewWriter(pWrtr)
	go func() {
		_, err = io.Copy(gw, file)
		if err != nil {
			log.Fatalln(err)
		}
		gw.Close()
		pWrtr.Close()
	}()

	gzipFile, err := os.Create("/home/shishir/Selfhelp/sample_codes/untar_dir/repo.tar.gz")
	if err != nil {
		log.Fatalln(err)
	}
	if _, err = io.Copy(gzipFile, pRdr); err != nil {
		log.Fatalln(err)
	}
	gzipFile.Close()
	pRdr.Close()

	fmt.Println("Tar ball gzipped successfully")

}
