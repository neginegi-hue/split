package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	filename := "example.txt"

	// 拡張子を取得
	extension := filepath.Ext(filename)

	// 拡張子を削除
	filenameWithoutExtension := filename[:len(filename)-len(extension)]

	// 拡張子を再び結合
	filenameWithExtension := filenameWithoutExtension + extension

	fmt.Println(filenameWithoutExtension) // 出力は "example"
	fmt.Println(filenameWithExtension)    // 出力は "example.txt"
}
