/*
This script demonstrates creation of a json-formatted file containing
the contents of a CSV file.  The source data file contains various
information about notable people.  Each record is placed into a
struct, then the struct is serialized to disk in json format.

The data can be obtained as an Excel sheet from this site:

http://science.sciencemag.org/content/suppl/2014/07/30/345.6196.558.DC1

To run this script, the data should be extracted from Excel and
converted to gziped text/csv.
*/

package main

import (
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"strconv"
)

// A struct holding information about a notable person
type notable struct {

	// The person's name
	PrsLabel string

	// The person's year of birth
	BYear int

	// The person's birth location
	BLocLabel string

	// The latitude of the person's birth location
	BLocLat float64

	// The longitude of the person's birth location
	BLocLong float64

	// The year of the person's birth
	DYear int

	// The location where the person died
	DLocLabel string

	// The latitude of the location where the person died
	DLocLat float64

	// The longitude of the location where the person died
	DLocLong float64

	// The person's gender
	Gender string
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

	// Open a file and create a json encoder to write to it
	out, err := os.Create("schich.json")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	enc := json.NewEncoder(out)

	for {
		row, err := crd.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// Create a struct holding the data for one record.
		// Note that unparseable values will be represented by
		// 0.
		r := notable{
			PrsLabel:  row[1],
			BLocLabel: row[3],
			DLocLabel: row[8],
			Gender:    row[12],
		}
		r.BYear, _ = strconv.Atoi(row[2])
		r.DYear, _ = strconv.Atoi(row[7])
		r.BLocLat, _ = strconv.ParseFloat(row[5], 64)
		r.BLocLong, _ = strconv.ParseFloat(row[6], 64)
		r.DLocLat, _ = strconv.ParseFloat(row[10], 64)
		r.DLocLong, _ = strconv.ParseFloat(row[11], 64)

		err = enc.Encode(r)
		if err != nil {
			panic(err)
		}
	}
}
