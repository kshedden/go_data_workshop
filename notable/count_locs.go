/*
This script demonstrates reading and calculating summary statistics
from a compressed CSV file.  The file contains birth and death
locations and dates for notable people.  We calculate and print the
number of distinct locations that appear as birth locations, as death
locations, and the number of distinct birth / death locations.

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

	// Read the header
	_, err = crd.Read()
	if err != nil {
		panic(err)
	}

	blocs := make(map[string]int)
	dlocs := make(map[string]int)
	bdlocs := make(map[string]int)

	for {
		row, err := crd.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		blocs[row[3]]++
		dlocs[row[8]]++
		bdlocs[row[3]+"::"+row[8]]++
	}

	fmt.Printf("Number of distinct birth locations:          %d\n", len(blocs))
	fmt.Printf("Number of distinct death locations:          %d\n", len(dlocs))
	fmt.Printf("Number of distinct birth x death locations:  %d\n", len(bdlocs))
}
