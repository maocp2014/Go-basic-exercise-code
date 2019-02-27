package main

import (
	"go_pratice_code/learn_ccmouse_code/crawler/engine"
	"go_pratice_code/learn_ccmouse_code/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
