package main

import (
	"bufio"
	"fmt"
	"log"
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
	cli.StringFlag{
		Name:        "bytes, c",
		Value:       "0",
		Usage:       "ファイルの先頭から指定した文字数を出力します",
		Destination: &argsChar,
	},
	cli.BoolFlag{
		Name:        "verbose, ve",
		Usage:       "ファイル名を表示します",
		Destination: &shouldShowFileName,
	},
	cli.BoolFlag{
		Name:        "quiet, q",
		Usage:       "ファイル名を表示しません",
		Destination: &shouldHiddenFileName,
	},
}

var argsNum, argsChar string
var shouldShowFileName, shouldHiddenFileName bool

func doMyhead(c *cli.Context) error {

	// 引数で渡されたファイル数を取得
	var hasMultiplefiles bool
	numOfFiles := len(c.Args())
	if 1 < numOfFiles {
		hasMultiplefiles = true
	}

	numberOfLines, _ := strconv.Atoi(argsNum)
	for i := 0; i < numOfFiles; i++ {
		// 引数で与えられたファイルをオープン
		file, err := os.Open(c.Args().Get(i))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// オプションに応じてファイル名を表示制御
		if (shouldShowFileName || hasMultiplefiles) && !shouldHiddenFileName {
			fmt.Printf("==> %s <==\n", c.Args().Get(i))
		}

		numberOfChar, _ := strconv.Atoi(argsChar)
		if numberOfChar == 0 {
			// バイト数が未指定の場合、行ごとに読み出し
			scanner := bufio.NewScanner(file)
			j := 0
			for j < numberOfLines && scanner.Scan() {
				line := scanner.Text()
				fmt.Println(line)
				j++
			}

		} else {
			// バイト数が指定されていた場合、該当バイト数分読み出し
			buf := make([]byte, numberOfChar)
			file.Read(buf)

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(buf))
		}

		fmt.Print("\n")
	}

	return nil
}
