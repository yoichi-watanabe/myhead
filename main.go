package main

import (
    "bufio"
    "os"
    "fmt"
    "strconv"
    
    "github.com/urfave/cli"
)

func main() {
  var (
    argsNum string
  )

  app := cli.NewApp()

  app.Name = "sampleApp"
  app.Usage = "This app echo input arguments"
  app.Version = "0.0.1"

  // オプション
  app.Flags = []cli.Flag{
    cli.StringFlag{
        Name:        "line, n",
        Value:       "10",
        Usage:       "ファイルの先頭から指定した行数を出力します",
        Destination: &argsNum,
    },
}

  app.Action = func (context *cli.Context) error {

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

  app.Run(os.Args) 
}
