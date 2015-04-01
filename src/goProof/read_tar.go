//Untar the tar ball to a specified destination.

//NOTE: Please use your own tar ball to run the program.

package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func main() {

	file, err := os.Open("/home/shishir/docker/src/github.com/docker/docker/img.tar")

	if err != nil {
		log.Fatalln(err)
	}

	tr := tar.NewReader(file)
	for {
		hdr, err := tr.Next()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		f1 := hdr.FileInfo()

		fmt.Printf("Header for entry is %s\n", strings.Trim(hdr.Name, "/"))
		if f1.IsDir() {
			if err = os.MkdirAll(path.Join("/home/shishir/Selfhelp/sample_codes/untar_dir", hdr.Name), f1.Mode()); err != nil {
				log.Fatalln(err)
			}
		} else {

			tarFile, err := os.Create(path.Join("/home/shishir/Selfhelp/sample_codes/untar_dir", hdr.Name))
			if err != nil {
				log.Fatalln(err)
			}
			if _, err = io.Copy(tarFile, tr); err != nil {
				log.Fatalln(err)
			}
			tarFile.Close()
		}

	}
	fmt.Println("Tar ball untarred successfully")
}
