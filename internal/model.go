package internal

func NewTpl() *Tpl {
	return &Tpl{
		Name:        "",
		Params:      make([]*Param, 0),
		Tpls:        make([]*LangTpl, 0),
		DefaultLang: "",
		Comment:     "",
	}
}

type Tpl struct {
	Name        string     `yaml:"name"`
	Params      []*Param   `yaml:"params"`
	Tpls        []*LangTpl `yaml:"tpls"`
	DefaultLang string     `yaml:"default"`
	Comment     string     `yaml:"default"`
}

type Param struct {
	name       string
	typeString string
}

type LangTpl struct {
	lang string
	text string
}
