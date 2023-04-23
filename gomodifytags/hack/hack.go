package main

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

func main() {
	f := genFile()
	fmt.Printf("%#v", f)
}

func genFuncRegisterFlags() jen.Code {
	fn := jen.Func().
		Params(jen.Id("f").Op("*").Id("UserForm")).
		Id("RegisterFlags").
		Params(jen.Id("flags").Op("*").Qual("github.com/spf13/pflag", "FlagSet"))

	calls := make([]jen.Code, 0)
	calls = append(calls, jen.Id("flags").Dot("StringVar").
		Call(jen.Op("&").Id("f").Dot("Name"), jen.Lit("name"), jen.Lit(""), jen.Lit("")))
	calls = append(calls, jen.Id("flags").Dot("StringVar").
		Call(jen.Op("&").Id("f").Dot("Email"), jen.Lit("email"), jen.Lit(""), jen.Lit("")))
	calls = append(calls, jen.Id("flags").Dot("IntVar").
		Call(jen.Op("&").Id("f").Dot("Age"), jen.Lit("age"), jen.Lit(0), jen.Lit("")))
	block := fn.Block(calls...)
	return block
}

func genFile() *jen.File {
	ret := jen.NewFile("main")
	ret.ImportName("github.com/spf13/pflag", "pflag")
	ret.Add(genFuncRegisterFlags())
	return ret
}
