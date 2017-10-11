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
	"sort"
	"strconv"
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

	year := make(map[string]float64)
	num := make(map[string]int)

	for {
		row, err := crd.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		y, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			continue
		}

		year[row[3]] += y
		num[row[3]]++
	}

	// Divide the total by the number of values to get the mean
	for k, _ := range year {
		year[k] /= float64(num[k])
	}

	var a []string
	for k, _ := range year {
		a = append(a, k)
	}
	sort.StringSlice(a).Sort()

	out, err := os.Create("mean_by_year.csv")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	for _, k := range a {
		if num[k] > 10 {
			out.WriteString(fmt.Sprintf("%s,%.0f,%d\n", k, year[k], num[k]))
		}
	}
}
