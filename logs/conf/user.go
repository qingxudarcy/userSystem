package conf

import "github.com/tal-tech/go-zero/core/logx"

var DefaultLog = logx.LogConf{
	ServiceName: "default",
	Mode:        "file",
	Path:        "logs",
}
