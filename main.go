package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/k0kubun/pp"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tzmfreedom/alt-apex/parser"
	"gopkg.in/urfave/cli.v1"
)

var Version string

func main() {
	node, err := ParseFile(os.Args[1])
	visitor := &StringVisitor{}
	str, err := node.Accept(visitor)
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}

func _cli() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(Version)
	}
	app := cli.NewApp()
	app.Name = "alt-apex"
	app.Usage = "Apex Alternative Language"
	app.Version = Version
	app.Commands = []cli.Command{
		{},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func ParseFile(f string) (parser.Node, error) {
	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	src := string(bytes)
	input := antlr.NewInputStream(src)
	return parse(input, f), nil
}

func ParseString(src string) (parser.Node, error) {
	input := antlr.NewInputStream(src)
	return parse(input, "<string>"), nil
}

func parse(input antlr.CharStream, src string) parser.Node {
	lexer := parser.NewKotlinLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewKotlinParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.KotlinFile()
	return tree.Accept(&parser.AstBuilder{}).(parser.Node)
}

func debug(args ...interface{}) {
	pp.Println(args...)
}
