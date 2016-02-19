//Read and Write a Comma separated values (CSV) file using encoding/csv

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	lvmCountConfig = "/tmp/lvmCountConfig.csv"
)

type vol struct {
	Name       string
	Mountpoint string
}

var count map[*vol]int

func main() {
	count = make(map[*vol]int)
	populateMap()

	// Write contents of the map (count) into a csv file
	fh, err := os.Create(lvmCountConfig)
	if err != nil {
		log.Fatal("Cannot create csv file: ", err)
	}
	defer fh.Close()

	csvWriter := csv.NewWriter(fh)
	for v, c := range count {
		if err := csvWriter.Write([]string{v.Name, v.Mountpoint, strconv.Itoa(c)}); err != nil {
			log.Fatal(err)
		}
	}

	csvWriter.Flush()
	fmt.Println("csv write successful")

	//Flush the contents of the count map.
	deleteMap()

	// Load contents into the map (count) from the csv file

	fhRead, err := os.Open(lvmCountConfig)
	if err != nil {
		log.Fatal(err)
	}

	defer fhRead.Close()

	csvReader := csv.NewReader(fhRead)
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		v := &vol{Name: record[0], Mountpoint: record[1]}
		c, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatal(err)
		}
		count[v] = c
	}

	fmt.Println("csv read successful")

	// Print the contents of the map after csv read.

	fmt.Println("Contents of the count map are:")
	for key, value := range count {
		fmt.Println(key.Name + " " + key.Mountpoint)
		fmt.Println(value)
	}

}

func populateMap() {
	v := &vol{Name: "yahoo", Mountpoint: "/var/run/docker-lvm/yahoo"}
	count[v] = 0
	v = &vol{Name: "redhat", Mountpoint: "/var/run/docker-lvm/redhat"}
	count[v] = 0
	v = &vol{Name: "google", Mountpoint: "/var/run/docker-lvm/google"}
	count[v] = 0

}

func deleteMap() {
	for key := range count {
		delete(count, key)
	}
}
