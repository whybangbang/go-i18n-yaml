package i18n

import (
	"mimir/common/zlog"
	"testing"
)

func TestI18nTp2(t *testing.T) {
	result := Tp2("es", "why")
	zlog.Info("result: ", result)
}