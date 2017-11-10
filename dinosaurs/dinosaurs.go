package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		dinosaurs map[string]map[string]string
	)

	dinosaurs = make(map[string]map[string]string)
	// Process file1 and add to data structure
	file1, e := ioutil.ReadFile("./dinosaurs1.csv")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	dino1 := strings.Split(strings.TrimSpace(string(file1)), "\n")
	for _, l := range dino1 {
		i := strings.Split(l, ",")
		dinosaurs[i[0]] = make(map[string]string)
		dinosaurs[i[0]]["length"] = i[1]
		dinosaurs[i[0]]["stance"] = i[2]
	}
	// Process file2 and add to data structure
	file2, e := ioutil.ReadFile("./dinosaurs2.csv")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	dino2 := strings.Split(strings.TrimSpace(string(file2)), "\n")
	for _, l := range dino2 {
		i := strings.Split(l, ",")
		dinosaurs[i[0]]["stride"] = i[1]
		dinosaurs[i[0]]["food"] = i[2]
	}
	for k := range dinosaurs {
		stride, _ := strconv.ParseFloat(dinosaurs[k]["stride"], 64)
		length, _ := strconv.ParseFloat(dinosaurs[k]["length"], 64)
		dinosaurs[k]["speed"] = strconv.FormatFloat(stride*length, 'E', -1, 64)
	}
	// Create sorted slice [[Name, Speed], [Name, Speed]]
	sorted := [][]string{}
	for k := range dinosaurs {
		if dinosaurs[k]["stance"] == "bipedal" {
			sorted = append(sorted, []string{k, dinosaurs[k]["speed"]})
		}
	}
	sort.Slice(sorted, func(i, j int) bool {
		a, _ := strconv.ParseFloat(sorted[j][1], 64)
		b, _ := strconv.ParseFloat(sorted[i][1], 64)
		return a < b
	})
	// Print the sorted array
	for _, v := range sorted {
		f, _ := strconv.ParseFloat(v[1], 64)
		fmt.Printf("Name: %v, Speed: %0.2f\n", v[0], f)
	}
}
