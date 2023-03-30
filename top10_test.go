package main

import (
	"log"
	"os"
	"testing"
)

func TestSameValueMultipleTimes(t *testing.T) {

	file, err := os.Open("test_files/urls_same_values.txt")
	if err != nil {
		log.Fatalf("failed to open file")
	}
	defer file.Close()

	want := loglines{
		logline{url: "http://problem.com/albo/i/nie", size: 9999998958502},
		logline{url: "http://problem.com/albo/i/nie", size: 9999998958502},
		logline{url: "http://mam-1-problem.com/albo/i/nie", size: 8958502},
		logline{url: "http://mam-2-problem.com/albo/i/nie", size: 8958502},
		logline{url: "http://mam-3-problem.com/albo/i/nie", size: 8958502},
		logline{url: "http://mam-4-problem.com/albo/i/nie", size: 8958502},
		logline{url: "http://mam-5-problem.com/albo/i/nie", size: 8958502},
		logline{url: "http://mam-6-problem.com/albo/i/nie", size: 8958502},
		logline{url: "http://mam-7-problem.com/albo/i/nie", size: 8958502},
		logline{url: "http://mam-8-problem.com/albo/i/nie", size: 8958502},
	}

	output := getTOP10(file)

	if want.Len() != output.Len() {
		t.Fatalf("Wrong number of elements in output list!")
	}
	for i := range output.getURLS() {
		if output[i] != want[i] {
			t.Error("Wrong order of elements in output list")
		}
	}
}
