package doc_write

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	n = 100000
)

func writeDirect() time.Duration {
	f, _ := os.Create("UnbufferedFile.txt")
	defer f.Close()

	line := []byte("Buffered\n")

	start := time.Now()

	for range n {
		f.Write(line)
	}

	return time.Since(start)
}

func writeBuffer() time.Duration {
	f, _ := os.Create("Buffered_file.txt")
	defer f.Close()

	w := bufio.NewWriter(f)

	start := time.Now()

	for range n {
		w.WriteString("Unbuffered\n")
	}

	w.Flush() //write what's left in buffer
	return time.Since(start)

}

func printWords(path string, n int) {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	count := 0

	for scanner.Scan() && count < n {
		fmt.Print(scanner.Text(), " ")
		count++
	}
	fmt.Println()
}

func WriteTime() {
	// start := time.Now()
	// writeDirect()
	// t1 := time.Since(start)

	// start = time.Now()
	// writeBuffer()
	// t2 := time.Since(start)
	t1 := writeDirect()
	t2 := writeBuffer()

	fmt.Printf("Unbuffered: %v (~%d ms)\n", t1, t1.Milliseconds())
	fmt.Printf("Buffered:   %v (~%d ms)\n", t2, t2.Milliseconds())

	ub := "Buffered_file.txt"
	bf := "UnbufferedFile.txt"

	fmt.Print("First 5 words (unbuffered): ")
	printWords(ub, 5)
	fmt.Print("First 5 words (buffered):   ")
	printWords(bf, 5)

	// Delete file
	_ = os.Remove(ub)
	_ = os.Remove(bf)

}
