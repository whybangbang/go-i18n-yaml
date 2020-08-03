package internal

import (
	"fmt"
	"testing"
)

func TestNewCodeBuilder(t *testing.T) {
	t1 := &Tpl{
		Name:        "tpl1",
		Params:      []*Param{
			{
				name:       "p1",
				typeString: "int64",
			},
			{
				name:       "p2",
				typeString: "string",
			},
		},
		Tpls:        []*LangTpl{
			{
				lang: "en",
				text: "p1 %d p2: %s",
			},
			{
				lang: "zh",
				text: "你好 %d 你不好: %s",
			},
		},
		DefaultLang: "en",
		Comment:     "this is a",
	}

	builder := NewCodeBuilder("i18n", "I18n", "en")
	data, err := builder.Build(t1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}