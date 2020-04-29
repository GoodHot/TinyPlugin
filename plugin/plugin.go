package plugin

type Option struct {
	Port    string
	WebPage string
}

func NewPlugin(opt Option) *Plugin {
	return nil
}

type Plugin struct {
	Port string `json:"port"` // 插件服务启动的端口，如果配置为'auto'则为自动寻找可用端口
}

func (*Plugin) Start() {
}

// 注册所有请求的回调钩子
func (*Plugin) RegisterHook(name string, hook func(ctx HTTPContext)) {
}

func (*Plugin) RegisterController(uri string, ctrl func(ctx HTTPContext)) {

}
