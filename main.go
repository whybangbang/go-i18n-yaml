package main

import (
	"flag"
	"fmt"
	"i18n-yaml/internal"
)

var (
	tplFile     = flag.String("tpl_file", "/Users/why/go/src/i18n-yaml/example/tpl.yaml", "source yaml")
	outFile     = flag.String("out_file", "/Users/why/go/src/i18n-yaml/example/tpl.go", "generated go")
	pkgName     = flag.String("pkg_name", "i18n", "code save to pkg, default i18n")
	funcPrefix  = flag.String("func_prefix", "", "func prefix, help Coding")
	defaultLang = flag.String("default_lang", "en", "i18n default lang")
)

func main() {
	flag.Parse()

	tpls := make([]*internal.Tpl, 0)
	err := internal.NewTplDecoder(*tplFile).Deocde(&tpls)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = internal.NewCodeBuilder(*pkgName, *funcPrefix, *defaultLang).GenerateGoFile(*outFile, tpls...)
	if err != nil {
		fmt.Println(err)
		return
	}
}
