package main

import (
	"czfinger/utils"
	"fmt"
	cli "github.com/jawher/mow.cli"

	"os"
)

var (
	app *cli.Cli
)

func main() {

	app = cli.App("czfinger", "Simple website fingerprinting tool")
	var (
		target      = app.StringsOpt("t target", make([]string, 0), "Tager url")
		targetFiles = app.StringsOpt("tf", make([]string, 0), "Tager url file")
		threads     = app.IntOpt("t threads", 10, "threads nubmer ")
		//timeout    = app.IntOpt("t timeout", 20, "timeout ")
	)
	app.Version("v version", "czfinger 1.0")
	app.Spec = "-v | (-t=<target> ) | (--tf=<targetfile> )| (-t=<threads> )"
	app.Action = func() {
		targetsSlice := make([]string, 0) //创建一个目标地址切片
		if len(*target) != 0 {
			targetsSlice = append(targetsSlice, *target...)
		}
		for _, targetFile := range *targetFiles {
			if utils.Exists(targetFile) {
				err, lineSlice := utils.ReadLineFile(targetFile)
				if err != nil {
					utils.OptionsError(err.Error(), 2)
				}
				targetsSlice = append(targetsSlice, lineSlice...)
				fmt.Println(targetsSlice)
				fmt.Println("文件存在")

			}

		}
		if len(targetsSlice) == 0 {
			utils.OptionsError("No target targetsSlice", 2)
		}
		targetsSlice = utils.RemoveRepeatedElement(targetsSlice)
		fmt.Println(targetsSlice)
		//设置进程数 根据目标数量进行设置
		if *threads > len(targetsSlice) {
			*threads = len(targetsSlice)
		}

	}
	err := app.Run(os.Args)
	if err != nil {
		return
	}
}
