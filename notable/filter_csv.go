/*
This script demonstrates reading a compressed CSV file, and
printing a subset of the lines to standard output (the screen).

The data can be obtained as an Excel sheet from this site:

http://science.sciencemag.org/content/suppl/2014/07/30/345.6196.558.DC1
*/

package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	fname := "SchichDataS1_FB.csv.gz"

	// Open a reader for the file
	fid, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer fid.Close()

	// Decompress the stream on-the-fly
	gid, err := gzip.NewReader(fid)
	if err != nil {
		panic(err)
	}
	defer gid.Close()

	// Parse the CSV data
	crd := csv.NewReader(gid)
	if err != nil {
		panic(err)
	}

	// Read the header
	_, err = crd.Read()
	if err != nil {
		panic(err)
	}

	for {
		row, err := crd.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if row[3] == "Detroit" || row[8] == "Detroit" {
			fmt.Printf(strings.Join(row, ",") + "\n")
		}
	}
}
