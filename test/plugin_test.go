package test

import (
	"fmt"
	"github.com/GoodHot/TinyPlugin/ip2region"
	"github.com/GoodHot/TinyPlugin/plugin"
	"testing"
	"time"
)

func TestPlugin1(t *testing.T) {
	server := plugin.NewPlugin(plugin.Option{
		Port:    "auto",
		WebPage: "../webpage",
	})

	server.RegisterHook("request", func(ctx plugin.HTTPContext) {
	})

	server.RegisterController("", func(ctx plugin.HTTPContext) {

	})

	server.Start()
}

func TestPI(t *testing.T) {
	region, _ := ip2region.New("ip2region.db")
	begin := time.Now()
	fmt.Println(region.BinarySearch("36.248.208.254"))
	fmt.Println(region.BinarySearch("36.248.208.254"))
	fmt.Println(region.BinarySearch("36.248.208.254"))
	fmt.Println(region.BinarySearch("36.248.208.254"))
	end := time.Now()
	fmt.Println(end.UnixNano() - begin.UnixNano())


	begin = time.Now()
	fmt.Println(region.BtreeSearch("36.248.208.254"))
	fmt.Println(region.BtreeSearch("36.248.208.254"))
	fmt.Println(region.BtreeSearch("36.248.208.254"))
	fmt.Println(region.BtreeSearch("36.248.208.254"))
	end = time.Now()
	fmt.Println(end.UnixNano() - begin.UnixNano())


	begin = time.Now()
	fmt.Println(region.MemorySearch("36.248.208.254"))
	fmt.Println(region.MemorySearch("36.248.208.254"))
	fmt.Println(region.MemorySearch("36.248.208.254"))
	fmt.Println(region.MemorySearch("36.248.208.254"))
	end = time.Now()
	fmt.Println(end.UnixNano() - begin.UnixNano())
}
