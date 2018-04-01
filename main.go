package main

import (
    "bufio"
    "os"
    "fmt"
    
    "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()

  app.Name = "sampleApp"
  app.Usage = "This app echo input arguments"
  app.Version = "0.0.1"

  app.Action = func (context *cli.Context) error {

    // 引数で与えられたファイルをオープン
	  file, err := os.Open(context.Args().Get(0))
	  if err != nil {
		  panic(err)
	  }
    defer file.Close()
    
    // 一行ずつ読み出し
	  scanner := bufio.NewScanner(file)
	  for scanner.Scan() {
	  	line := scanner.Text()
	  	fmt.Println(line)
	  }

    return nil
  }

  app.Run(os.Args) 
}