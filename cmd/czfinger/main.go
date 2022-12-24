package main

import (
	cli "github.com/jawher/mow.cli"
	"os"
)

var (
	app *cli.Cli
)

func main() {

	app = cli.App("czfinger", "Simple website fingerprinting tool")
	var (
		target = app.StringOpt("t target", "", "Tager url")
		//targetfile = app.StringOpt("tf targetfile", "", "Tager url file")
		//threads    = app.IntOpt("t threads", 10, "threads nubmer ")
		//timeout    = app.IntOpt("t timeout", 20, "timeout ")
	)
	app.Version("v version", "czfinger 1.0")
	app.Spec = "-v | (-t=<target> )"
	app.Action = func() {
		targetsSlice := make([]string, 0) //创建一个目标地址切片
		if len(*target) != 0 {
			targetsSlice = append(targetsSlice, *target)
		}
		for _, v := range targetsSlice {
		}

	}
	err := app.Run(os.Args)
	if err != nil {
		return
	}
}
