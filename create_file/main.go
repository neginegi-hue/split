package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 作成するファイルの名前と行数を設定
	fileName := "input.txt"
	numberOfLines := 10000 // ここに行数を設定します

	// ファイルを開く
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("ファイルの作成に失敗しました:", err)
		return
	}
	defer file.Close()

	// バッファ付きライターを使用して効率的に書き込む
	writer := bufio.NewWriter(file)

	// 指定した行数だけループしてファイルに書き込む
	for i := 1; i <= numberOfLines; i++ {
		var line string
		if i != numberOfLines {
			line = fmt.Sprintf("%d\n", i)

		} else {
			line = fmt.Sprintf("%d", i)
		}
		writer.WriteString(line)
	}

	// バッファをフラッシュして全てのデータをファイルに書き込む
	writer.Flush()

	fmt.Printf("%d行のダミーファイル%sが作成されました。\n", numberOfLines, fileName)
}
