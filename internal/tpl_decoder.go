package internal

import (
	"errors"
	"github.com/kylelemons/go-gypsy/yaml"
)

func NewTplDecoder(filePath string) *tplDecoder {
	return &tplDecoder{
		filePath: filePath,
	}
}

type tplDecoder struct {
	filePath string
}

func (t *tplDecoder) Deocde(tplSlice *[]*Tpl) error {
	config, err := yaml.ReadFile(t.filePath)
	if err != nil {
		return err
	}
	node, err := yaml.Child(config.Root, "tpls")
	if err != nil {
		return err
	}
	listNode, ok := node.(yaml.List)
	if !ok {
		return errors.New("tpls is not list")
	}
	for i := 0; i< listNode.Len(); i++ {
		tpl := NewTpl()
		node := listNode.Item(i)
		tplNode, ok := node.(yaml.Map)
		if !ok {
			return errors.New("tplNode is not map")
		}
		name, ok := tplNode.Key("name").(yaml.Scalar)
		if !ok {
			return errors.New("name is not string")
		}
		tpl.Name = name.String()
		defaultLang, ok := tplNode.Key("default").(yaml.Scalar)
		if ok {
			tpl.DefaultLang = defaultLang.String()
		}
		comment, ok := tplNode.Key("comment").(yaml.Scalar)
		if ok {
			tpl.Comment = comment.String()
		}
		params, ok := tplNode.Key("params").(yaml.Map)
		if ok {
			for k, v := range params {
				tpl.Params = append(tpl.Params, &Param{
					name:       k,
					typeString: v.(yaml.Scalar).String(),
				})
			}
		}

		tpls, ok := tplNode.Key("tpls").(yaml.Map)
		if !ok {
			return errors.New("tpls is not map")
		}
		for k, v := range tpls {
			tpl.Tpls = append(tpl.Tpls, &LangTpl{
				lang: k,
				text: v.(yaml.Scalar).String(),
			})
		}
		*tplSlice = append(*tplSlice, tpl)
	}

	return nil
}
