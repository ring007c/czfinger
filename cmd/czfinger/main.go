package main

import (
	"czfinger/core"
	"czfinger/utils"
	cli "github.com/jawher/mow.cli"
	"github.com/remeh/sizedwaitgroup"
	"os"
)

var (
	app *cli.Cli
	swg sizedwaitgroup.SizedWaitGroup
)

func main() {

	app = cli.App("czfinger", "Simple website fingerprinting tool")
	var (
		target      = app.StringOpt("u target", "", "Tager url")
		targetFiles = app.StringsOpt("tf", make([]string, 0), "Tager url file")
		threads     = app.IntOpt("t threads", 10, "threads nubmer ")
		//timeout    = app.IntOpt("t timeout", 20, "timeout ")
	)
	app.Version("v version", "czfinger 1.0")
	app.Spec = "-v | (-u=<target> ) | (--tf=<targetfile> )| (-t=<threads> )"
	app.Action = func() {
		targetsSlice := make([]string, 0) //创建一个目标地址切片
		if len(*target) != 0 {
			targetsSlice = append(targetsSlice, *target)
			text := core.Reqdata(*target)
			core.Detect(text)
			//fmt.Println(strings.ToLower(string(text.Content)))
			//fmt.Println(text.Url)
		}
		if *threads > len(targetsSlice) {
			*threads = len(targetsSlice)
		}
		swg = sizedwaitgroup.New(*threads)
		for _, targetFile := range *targetFiles {

			if utils.Exists(targetFile) {
				err, lineSlice := utils.ReadLineFile(targetFile)

				if err != nil {
					utils.OptionsError(err.Error(), 2)
				}

				targetsSlice = append(targetsSlice, lineSlice...)
				//fmt.Println(targetsSlice)
				//fmt.Println("文件存在")

			}

		}

		if len(targetsSlice) == 0 {
			utils.OptionsError("No target targetsSlice", 2)
		}
		targetsSlice = utils.RemoveRepeatedElement(targetsSlice)
		//fmt.Println(targetsSlice)
		for i := 0; i < len(targetsSlice); i++ {
			swg.Add()
			// 开启一个并发
			go func(url string) {
				// 使用defer, 表示函数完成时将等待组值减1
				defer swg.Done()
				//fmt.Println(url)
				req := core.Reqdata(url)
				//fmt.Println(strings.ToLower(string(req.Content)))
				core.Detect(req)

			}(targetsSlice[i])

		}
		swg.Wait()

	}
	app.Run(os.Args)
}
