# go-i18n-yaml
generate i18n code by yaml file

compile

	go build -gcflags "-N -l" main.go
	mv main i18n-yaml
	mv i18n-yaml ~/go/bin

build

	~/go/bin/i18n-yaml --help
	~/go/bin/i18n-yaml -default_lang=en -func_prefix=I18n -out_file=/Users/why/go/src/i18n-yaml/example/tpl.go -pkg_name=i18n -tpl_file=/Users/why/go/src/i18n-yaml/example/tpl.yaml
 
优点

* 简单
* 实用
*  类proto 实现方法，yaml生成调用代码
* github.com/nicksnyder/go-i18n/v2/i18n 提供了很多中看不中用的方法，我感觉应该这样实现

