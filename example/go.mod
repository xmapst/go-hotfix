module github.com/xmapst/go-hotfix/example

go 1.23.3

replace github.com/xmapst/go-hotfix => ../

require (
	github.com/pires/go-proxyproto v0.8.0
	github.com/reiver/go-oi v1.0.0
	github.com/reiver/go-telnet v0.0.0-20180421082511-9ff0b2ab096e
	github.com/soheilhy/cmux v0.1.5
	github.com/xmapst/go-hotfix v0.0.0-00010101000000-000000000000
	github.com/xmapst/logx v1.0.4
	go.uber.org/zap v1.27.0
)

require (
	github.com/agiledragon/gomonkey/v2 v2.12.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/traefik/yaegi v0.16.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
)
