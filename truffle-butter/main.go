package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"sync"

	"github.com/MrGossett/go-play/truffle-butter/splitter"
)

func main() {
	hexReader := &bytes.Buffer{}
	b64Reader := io.TeeReader(os.Stdin, hexReader)

	hexScanner := bufio.NewScanner(hexReader)
	hexScanner.Split(splitter.HexSplitter.ScanCharset)

	b64Scanner := bufio.NewScanner(b64Reader)
	b64Scanner.Split(splitter.B64Splitter.ScanCharset)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		sniff(b64Scanner, 4.5)
		wg.Done()
	}()
	go func() {
		sniff(hexScanner, 3)
		wg.Done()
	}()
	wg.Wait()
}

func sniff(scanner *bufio.Scanner, threshold float64) {
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) < 15 {
			continue
		}
		if e := shannonEntropy(t); e > threshold {
			fmt.Printf("%5.2f : %s\n", e, t)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func shannonEntropy(s string) float64 {
	var entropy float64

	len := float64(len(s))
	seen := map[rune]struct{}{}

	for _, r := range s {
		if _, ok := seen[r]; ok {
			continue
		}
		pr := float64(strings.Count(s, string(r))) / len
		entropy += pr * math.Log2(pr) * -1
	}
	return entropy
}
