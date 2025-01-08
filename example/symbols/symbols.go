package symbols

import (
	"reflect"
)

var Symbols = map[string]map[string]reflect.Value{}

//go:generate go install github.com/traefik/yaegi/cmd/yaegi@latest
//go:generate yaegi extract github.com/xmapst/go-hotfix
//go:generate yaegi extract github.com/xmapst/logx

//go:generate yaegi extract github.com/xmapst/go-hotfix/example/handler
