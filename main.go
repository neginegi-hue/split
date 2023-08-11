package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var bytes int
var LineCount int
var suffixNumber bool

func init() {
	flag.BoolVar(&suffixNumber, "d", false, "接尾語を数値に変更する")
	flag.IntVar(&LineCount, "l", 1000, "行数の値を指定")
	flag.IntVar(&bytes, "b", 0, "バイト数単位の値を指定")
}

func main() {
	flag.Parse()

	if bytes == 0 {
		linesSplit()
	} else {
		//bytesSplit()
	}
}

func linesSplit() {

	args := flag.Args()

	if len(args) < 1 {
		return
	}

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

		if lineCount%LineCount == 0 {

			if outputFile != nil {
				writer.Flush()
				outputFile.Close()
			}

			fileCount++
			var newFilename string
			var curFilename string

			if len(args) < 2 {
				curFilename = ""
			} else {
				curFilename = args[1]
			}

			if suffixNumber {
				newFilename = createFilenameNumber(curFilename, fileCount)
			} else {
				newFilename = createFilenameString(curFilename, fileCount)
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

func bytesSplit() {

	args := flag.Args()

	if len(args) < 1 {
		return
	}

	inputFile, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close()

	content, err := ioutil.ReadFile(filePath)
	var fileCount int
	var outputFile *os.File
	var writer *bufio.Writer

	for i, j := 0, bytes; i < len(content); i, j = i+bytes, j+bytes {
		chunk := content[i:j]

		// バッファのサイズが足りない場合
		if j > len(content) {
			chunk = content[i:len(content)]
		}

		// チャンクをファイルに保存
		outputPath := fmt.Sprintf("%s_part_%d", filePath, i/bytes)
		err := ioutil.WriteFile(outputPath, chunk, os.ModePerm)
		if err != nil {
			fmt.Println("ファイルの書き込みエラー:", err)
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
