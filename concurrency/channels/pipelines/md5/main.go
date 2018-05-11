package main

import (
	"fmt"
	"os"
	"projects/golang/channels/pipelines/md5/bounded"
	"projects/golang/channels/pipelines/md5/parallel"
	"projects/golang/channels/pipelines/md5/sequential"
	"sort"
	"time"
)

func callSequential(root string) {
	// Calc md5 of all files under specified directory
	// print results sorted by path name

	m, err := sequential.MD5All(root)
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x, %s\n", m[path], path)
	}
	return
}

func callParallel(root string) {
	// Calc md5 of all files under specified directory
	// print results sorted by path name

	m, err := parallel.MD5All(root)
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x, %s\n", m[path], path)
	}
}

func callBounded(root string) {
	// Calc md5 of all files under specified directory
	// print results sorted by path name

	m, err := bounded.MD5All(root)
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x, %s\n", m[path], path)
	}
}
func main() {
	//fmt.Println("#### Sequential ####")
	//start1 := time.Now()
	//callSequential(os.Args[1])
	//elapsed1 := time.Since(start1)
	//fmt.Printf("Time for sequential: %v\n", elapsed1)
	//fmt.Println("#### Parallel ####")
	//start2 := time.Now()
	//callParallel(os.Args[1])
	//elapsed2 := time.Since(start2)
	//fmt.Printf("Time for parallel: %v\n", elapsed2)
	fmt.Println("#### Parallel Bounded ####")
	start3 := time.Now()
	callBounded(os.Args[1])
	elapsed3 := time.Since(start3)
	fmt.Printf("Time for parallel bounded: %v\n", elapsed3)
}
