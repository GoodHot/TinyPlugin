package test

import (
	"github.com/GoodHot/TinyPlugin/plugin"
	"testing"
)

func TestPlugin1(t *testing.T) {
	server := plugin.NewPlugin(plugin.Option{
		Port:    "auto",
		WebPage: "../webpage",
	})

	server.RegisterHook("request", func(ctx plugin.HTTPContext) {
		ctx.Context.HTML(200, "")
	})

	server.RegisterController("", func(ctx plugin.HTTPContext) {

	})

	server.Start()
}
