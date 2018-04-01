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
}

var argsNum string

func doMyhead(context *cli.Context) error {

	// 引数で与えられたファイルをオープン
	file, err := os.Open(context.Args().Get(0))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// オプションで入力された行数を型変換（string -> int）
	numberOfLines, _ := strconv.Atoi(argsNum)

	// 一行ずつ読み出し
	scanner := bufio.NewScanner(file)
	i := 0
	for i < numberOfLines && scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		i++
	}

	return nil
}
