//Walking into a directory, its subdirectories and files and returning a map of the entries.
//NOTE: Please use your own input path to run the program.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
)

var entries []string
var dirMap map[string]bool

func InitMap() {
	dirMap = make(map[string]bool)
}

func WalkDir(pathin string) (map[string]bool, error) {
	names, err := ioutil.ReadDir(pathin)
	if err != nil {
		return nil, err
	}

	for _, fileName := range names {
		if fileName.IsDir() {
			dirMap[path.Join(pathin, fileName.Name())] = true
			_, err = WalkDir(path.Join(pathin, fileName.Name()))
			if err != nil {
				return nil, err
			}
		} else {
			dirMap[path.Join(pathin, fileName.Name())] = false
		}
	}
	return dirMap, nil
}

func main() {

	InitMap()
	m := make(map[string]bool)
	m, err := WalkDir("/home/shishir/Selfhelp/run_tmpfs_issue/input")
	if err != nil {
		log.Fatalln(err)
	}
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
}
