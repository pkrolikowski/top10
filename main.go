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

	ntop10 := getTOP10(file)
	for _, j := range ntop10.getURLS() {
		fmt.Println(j)
	}
}

func getTOP10(f *os.File) loglines {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	eList := loglines{}

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

		if eList.Len() == 0 {
			eList = append(eList, line)
		} else {

			smallest := eList[eList.Len()-1]
			if line.size > smallest.size && eList.Len() == 10 {
				eList[eList.Len()-1] = line
			} else if line.size == smallest.size && eList.Len() == 10 {
				continue
			} else {
				eList = append(eList, line)
			}
			sort.Sort(eList)
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

func (l loglines) getURLS() []string {
	var urls []string
	for _, u := range l {
		urls = append(urls, u.url)
	}
	return urls
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
