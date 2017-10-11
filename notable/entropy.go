/*
This script demonstrates reading and calculating summary statistics
from a compressed CSV file.  The file contains birth and death
locations and dates for notable people.

We calculate and print the entropy of the frequency distributions of
birth and death locations.

The data can be obtained as an Excel sheet from this site:

http://science.sciencemag.org/content/suppl/2014/07/30/345.6196.558.DC1

To run this script, the data should be extracted from Excel and
converted to gziped text/csv.
*/

package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
)

// entropy returns the entropy of the frequency distribution of the
// values of the map m.  The argument n is the total of all the
// values.
func entropy(m map[string]float64, n int) float64 {

	e := float64(0)

	for _, v := range m {
		p := v / float64(n)
		e -= p * math.Log(p)
	}

	return e
}

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

	// The number of people per birth location
	bloc := make(map[string]float64)

	// The number of people per death location
	dloc := make(map[string]float64)

	// The overall number of people
	n := 0

	for {
		row, err := crd.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// Update statistics
		bloc[row[3]]++
		dloc[row[8]]++
		n++
	}

	fmt.Printf("Birth location entropy: %f\n", entropy(bloc, n))
	fmt.Printf("Death location entropy: %f\n", entropy(dloc, n))
}
