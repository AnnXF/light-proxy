package cmd

import (
	"github.com/spf13/cobra"
	"runtime"
)

var supportOS = map[string]struct{}{
	"linux":   {},
	"windows": {},
	"darwin":  {},
}

func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Use config CLI to configure proxy settings.",
		Run: func(cmd *cobra.Command, args []string) {
			set, _ := cmd.Flags().GetString("set")
			config(set)
		},
	}

	cmd.Flags().StringP("set", "s", "0", `
--set 0 -> use 0 to clear proxy config
--set 127.0.0.1:8080 -> with domain param to set proxy config
`)

	return cmd
}

// config 代理配置
//
// set: 0 , 清除配置的代理
// set: 127.0.0.1:8080, 设置代理地址
func config(set string) {
	if _, ok := supportOS[runtime.GOOS]; !ok {
		println("auto config not support for os:" + runtime.GOOS)
		return
	}
	if set == "0" {
		configClear()
	} else {
		configClear()
		configSet(set)
	}
}
