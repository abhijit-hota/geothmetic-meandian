package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	debug = flag.Bool("d", false, "Debug mode. Prints all iterations.")
)

func parseFloatArr(sts []string) ([]float64, error) {
	var nums []float64
	for _, str := range sts {
		num, err := strconv.ParseFloat(str, 64)

		if err != nil {
			return nil, fmt.Errorf("error: non-numeric argument %v", str)
		}

		nums = append(nums, num)
	}
	return nums, nil
}

const helpStr = `
Usage:   %s [OPTIONS] <numeric args>
Example: 
	./gmdn 1 1 2 3 5
	./gmdn -d 1 1 2 3 5

Options:
`

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), helpStr, os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		flag.Usage()
		return
	}

	nums, err := parseFloatArr(args)
	if err != nil {
		log.Fatal(err)
	}

	gmdn := GeothmeticMeandian(nums)

	fmt.Println(gmdn)
}
