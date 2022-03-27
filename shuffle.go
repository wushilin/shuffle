package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
)

var seed = flag.String("s", "ABCD1234", "Seed for shuffling")
var reverse = flag.Bool("r", false, "Do reverse convert")

func main() {
	flag.Parse()
	var args = flag.Args()
	if len(args) == 0 {
		fmt.Println("Please specify list of files to process")
		return
	}
	dict := make([]byte, 256)
	makeSeed(*seed)
	makeDict(dict)
	revDict := makeRevDict(dict)
	for _, value := range args {
		if *reverse {
			DoEncoding(value, fmt.Sprintf("%s.res", value), revDict)
		} else {
			DoEncoding(value, fmt.Sprintf("%s.rev", value), dict)
		}
	}
}

func makeRevDict(dict []byte) []byte {
	result := make([]byte, len(dict))
	for idx, value := range dict {
		result[value] = byte(idx)
	}
	return result
}

func DoEncoding(what string, out string, dict []byte) {
	fmt.Printf("Encoding %s => %s, seed %s\n", what, out, *seed)
	ofh, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer ofh.Close()

	ifh, err := os.Open(what)
	if err != nil {
		panic(err)
	}

	defer ifh.Close()
	buff := make([]byte, 4096)
	for {
		nread, err := ifh.Read(buff)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		transform(buff, dict)
		nwritten, err := ofh.Write(buff[:nread])
		if err != nil {
			panic(err)
		}
		if nread != nwritten {
			panic("Short Write")
		}
	}
}

func transform(buff []byte, dict []byte) {
	for idx, value := range buff {
		buff[idx] = dict[value]
	}
}
func makeSeed(rseed string) {
	var theSeed int64 = 0
	var br = []byte(rseed)
	for index, value := range br {
		theSeed += int64(value) * int64(math.Pow(31, float64(len(rseed)-index-1)))
	}
	rand.Seed(theSeed)
}

func makeDict(dict []byte) {
	var i int = 0
	for i = 0; i < len(dict); i++ {
		dict[i] = byte(i)
	}

	shuffle(dict)
}

func shuffle(dict []byte) {
	var i int = 0
	for i = 0; i < len(dict); i++ {
		curr := dict[i]
		index := randInt(i, len(dict))
		if index != i {
			toSwap := dict[index]
			dict[i] = toSwap
			dict[index] = curr
		}
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
