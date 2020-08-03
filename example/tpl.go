package i18n
import "fmt"


func Tpl1(lang string, ) (string) {
	switch lang {
		case "en": return "tpl1"
		case "zh": return "模版1"
	}
	return "tpl1"

}

// "这是测试模版"
func Tp2(lang string, name string,) (string) {
	switch lang {
		case "en": return fmt.Sprintf("hello %s", name)
		case "zh": return fmt.Sprintf("你好: %s", name)
	}
	return fmt.Sprintf("hello %s", name)

}
