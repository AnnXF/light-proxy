package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proxy",
		Short: "A light proxy CLI tool",
		Long:  `This is a light proxy CLI tool developed by GetcharZp. It is primarily used for proxying internet connections, providing an easy way to manage and configure proxy settings.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Proxy !")
		},
	}

	// proxy 直连、桥接
	cmd.AddCommand(NewRelayCommand())
	// 配置
	cmd.AddCommand(NewConfigCommand())

	return cmd
}
