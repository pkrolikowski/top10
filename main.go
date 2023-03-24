package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// Supported URL Schemas
var urlSchemas = []string{"http", "https"}

func main() {
	var f string

	flag.StringVar(&f, "f", "", "(mandatory) absolute path to logfile")
	flag.Parse()

	if f == "" {
		flag.Usage()
		os.Exit(1)
	}

	if !filepath.IsAbs(f) {
		log.Fatal("Provided filepath is not absolute")
	}

	file, err := os.Open(f)
	if err != nil {
		log.Fatalf("failed to open file")
	}
	defer file.Close()

	top10 := getTOP10(file)

	for _, i := range top10 {
		fmt.Println(i.url)
	}
}

func getTOP10(f *os.File) []logline {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	eList := []logline{}

	// Iterate through file line by line
	for scanner.Scan() {
		scanned := scanner.Text()
		fields := strings.Fields(scanned)

		if len(fields) != 2 {
			fmt.Printf("ERROR: wrong number of fields in line %s\n", scanned)
			continue
		}
		size, _ := strconv.Atoi(fields[1])

		if !validateURL(fields[0]) {
			fmt.Printf("ERROR: wrong URL in line %s\n", scanned)
			continue
		}

		// Map parsed line fields to logline struct
		line := logline{
			url:  fields[0],
			size: size,
		}

		if len(eList) == 0 {
			eList = append(eList, line)
		} else {
			smallest := eList[len(eList)-1]
			if line.size > smallest.size && len(eList) == 10 {
				eList[len(eList)-1] = line
			} else {
				eList = append(eList, line)
			}
			sort.Sort(loglines(eList))
		}
	}
	return eList
}

func validateURL(r string) bool {
	_, err := url.ParseRequestURI(r)
	if err != nil {
		return false
	}
	u, err := url.Parse(r)
	if err != nil || !contains(urlSchemas, u.Scheme) || u.Host == "" {
		return false
	}
	return true
}

func contains(list []string, element string) bool {
	for _, v := range list {
		if element == v {
			return true
		}
	}
	return false
}

type loglines []logline

// Struct to fetch logline
type logline struct {
	url  string
	size int
}

// Create Len, Less, Swat functions to implement sort.Interface
func (e loglines) Len() int {
	return len(e)
}

func (e loglines) Less(i, j int) bool {
	return e[i].size > e[j].size
}

func (e loglines) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
