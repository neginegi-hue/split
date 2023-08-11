package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var bytes int
var lines int
var suffixNumber bool

func init() {
	flag.BoolVar(&suffixNumber, "d", false, "接尾語を数値に変更する")
	flag.IntVar(&lines, "l", 1000, "行数の値を指定")
	flag.IntVar(&bytes, "b", 0, "バイト数単位の値を指定")
}

func main() {
	run()
}

func run() {
	flag.Parse()

	args := flag.Args()
	inputFile, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close()

	reader := bufio.NewReader(inputFile)
	var lineCount int
	var fileCount int
	var outputFile *os.File
	var writer *bufio.Writer

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("Error reading line:", err)
			return
		}

		if lineCount%lines == 0 {
			if outputFile != nil {
				writer.Flush()
				outputFile.Close()
			}
			fileCount++
			var newFilename string
			if suffixNumber {
				newFilename = createFilenameNumber(args[1], fileCount)
			} else {
				newFilename = createFilenameString(args[1], fileCount)
			}

			outputFile, err = os.Create(newFilename)
			if err != nil {
				fmt.Println("Error creating output file:", err)
				return
			}
			writer = bufio.NewWriter(outputFile)
		}

		writer.WriteString(line)
		lineCount++

		if err == io.EOF {
			break
		}
	}
	writer.Flush()
	if outputFile != nil {
		outputFile.Close()
	}

	fmt.Printf("%d files created.\n", fileCount)
}

func createFilenameString(filename string, fileCount int) string {
	// 接尾語の生成
	suffix := fmt.Sprintf("x%c%c", 'a'+(fileCount-1)/26, 'a'+(fileCount-1)%26)

	newFilename := filename + suffix
	return newFilename
}

func createFilenameNumber(filename string, fileCount int) string {
	// 接尾語の生成
	suffix := fmt.Sprintf("%03d", fileCount)

	newFilename := filename + suffix
	return newFilename
}
