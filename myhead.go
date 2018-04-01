package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

var myHeadFlags = []cli.Flag{
	cli.StringFlag{
		Name:        "line, n",
		Value:       "10",
		Usage:       "ファイルの先頭から指定した行数を出力します",
		Destination: &argsNum,
	},
	cli.BoolFlag{
		Name:        "verbose, ve",
		Usage:       "ファイル名を表示します",
		Destination: &shouldShowFileName,
	},
}

var argsNum string
var shouldShowFileName bool
var hasMultiplefiles bool

func doMyhead(c *cli.Context) error {

	// 引数で渡されたファイル数を取得
	numOfFiles := len(c.Args())
	if 1 < numOfFiles {
		hasMultiplefiles = true
	}

	// オプションで入力された行数を型変換（string -> int）
	numberOfLines, _ := strconv.Atoi(argsNum)

	for i := 0; i < numOfFiles; i++ {
		// 引数で与えられたファイルをオープン
		file, err := os.Open(c.Args().Get(i))
		if err != nil {
			panic(err)
		}
		defer file.Close()

		if shouldShowFileName || hasMultiplefiles {
			fmt.Printf("==> %s <==\n", c.Args().Get(i))
		}

		// 一行ずつ読み出し
		scanner := bufio.NewScanner(file)
		j := 0
		for j < numberOfLines && scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
			j++
		}

	}

	return nil
}
