package internal

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"sync"
)

func NewCodeBuilder(pkgName string, funcPrefix string, defaultLang string) *CodeBuilder {
	if defaultLang == "" {
		defaultLang = "en"
	}
	return &CodeBuilder{
		buffer:      bytes.NewBufferString(""),
		funcPrefix:  funcPrefix,
		pkgName:     pkgName,
		defaultLang: defaultLang,
	}
}

type CodeBuilder struct {
	buffer      *bytes.Buffer
	funcPrefix  string
	pkgName     string
	defaultLang string
	once        sync.Once
}

func (d *CodeBuilder) appendComment(tpl *Tpl) error {
	if d.buffer.Len() != 0 {
		d.appendLineTerminator()
	}
	if tpl.Comment == "" {
		return nil
	}
	d.buffer.WriteString("// ")
	_, err := d.buffer.WriteString(tpl.Comment)
	if err != nil {
		return err
	}
	return nil
}

func (d *CodeBuilder) appendFuncTitle(tpl *Tpl) error {
	if d.buffer.Len() != 0 {
		d.appendLineTerminator()
	}
	d.buffer.WriteString("func ")
	d.buffer.WriteString(d.funcPrefix)
	d.buffer.WriteString(strings.ToUpper(tpl.Name[:1]))
	d.buffer.WriteString(tpl.Name[1:])
	d.buffer.WriteString("(lang string, ")
	for i, param := range tpl.Params {
		d.buffer.WriteString(param.name)
		d.buffer.WriteString(" ")
		d.buffer.WriteString(param.typeString)
		d.buffer.WriteString(",")
		if i != len(tpl.Params)-1 {
			d.buffer.WriteString(" ")
		}
	}
	d.buffer.WriteString(") (string) {")
	return nil
}

func (d *CodeBuilder) appendFuncEnd(tpl *Tpl) error {
	d.appendLineTerminator()
	d.buffer.WriteString("}")
	d.appendLineTerminator()
	return nil
}

func (d *CodeBuilder) appendContent(tpl *Tpl) error {
	d.appendLineTerminator()
	d.appendTab()
	d.buffer.WriteString("switch lang {")
	d.appendLineTerminator()
	for _, lang := range tpl.Tpls {
		d.appendTab()
		d.appendTab()
		d.buffer.WriteString("case \"")
		d.buffer.WriteString(lang.lang)
		d.buffer.WriteString("\": ")
		d.appendLangStr(lang, tpl)
	}
	d.appendTab()
	d.buffer.WriteString("}")
	d.appendLineTerminator()
	defaultLang := d.defaultLang
	if tpl.DefaultLang != "" {
		defaultLang = tpl.DefaultLang
	}
	for _, lang := range tpl.Tpls {
		if lang.lang == defaultLang {
			d.appendTab()
			d.appendLangStr(lang, tpl)
			return nil
		}
	}

	return errors.New("not found tpl text for defaultLang")
}

func (d *CodeBuilder) appendLangStr(lang *LangTpl, tpl *Tpl) error {
	if len(tpl.Params) == 0 {
		d.buffer.WriteString("return \"")
		d.buffer.WriteString(lang.text)
		d.buffer.WriteString("\"")
		d.appendLineTerminator()
	} else {
		d.buffer.WriteString("return fmt.Sprintf(\"")
		if lang.text[:1] == "\"" && lang.text[len(lang.text)-1:] == "\"" {
			lang.text = lang.text[1:len(lang.text)-1]
		}
		d.buffer.WriteString(lang.text)
		d.buffer.WriteString("\", ")
		for i, param := range tpl.Params {
			d.buffer.WriteString(param.name)
			if i != len(tpl.Params)-1 {
				d.buffer.WriteString(", ")
			} else {
				d.buffer.WriteString(")")
				d.appendLineTerminator()
			}

		}
	}
	return nil
}

func (d *CodeBuilder) appendLineTerminator() error {
	_, err := d.buffer.WriteString("\n")
	if err != nil {
		return err
	}
	return nil
}

func (d *CodeBuilder) appendTab() error {
	d.buffer.WriteString("\t")
	return nil
}

func (d *CodeBuilder) appendFileTitle() error {
	if d.pkgName == "" {
		return errors.New("pkgName is nil")
	}
	d.buffer.WriteString("package ")
	d.buffer.WriteString(d.pkgName)
	d.appendLineTerminator()
	d.buffer.WriteString("import \"fmt\"")
	d.appendLineTerminator()
	return nil
}

func (d *CodeBuilder) Build(tpls ...*Tpl) ([]byte, error) {
	err := d.build(tpls...)
	if err != nil {
		return nil, err
	}
	return d.buffer.Bytes(), nil
}

func (d *CodeBuilder) build(tpls ...*Tpl) error {
	var err error
	d.once.Do(func() {
		err = d.appendFileTitle()
		if err != nil {
			return
		}
		for _, tpl := range tpls {
			err = d.appendComment(tpl)
			if err != nil {
				return
			}
			err = d.appendFuncTitle(tpl)
			if err != nil {
				return
			}
			err = d.appendContent(tpl)
			if err != nil {
				return
			}
			err = d.appendFuncEnd(tpl)
			if err != nil {
				return
			}
		}
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *CodeBuilder) GenerateGoFile(filePath string, tpls ...*Tpl) error{
	err := d.build(tpls...)
	if err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = d.buffer.WriteTo(file)
	if err != nil {
		return err
	}
	return nil
}
