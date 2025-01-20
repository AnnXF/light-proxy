package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proxy",
		Short: "A light proxy CLI tool",
		Long:  `This is a light proxy CLI tool developed by GetcharZp. It is primarily used for proxying internet connections, providing an easy way to manage and configure proxy settings.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				log.Fatalf("[sys] run cmd help error:%s \n", err.Error())
			}
		},
	}

	// proxy 直连
	cmd.AddCommand(NewRelayCommand())
	// proxy 桥接
	cmd.AddCommand(NewBridgeCommand())
	// 配置
	cmd.AddCommand(NewConfigCommand())

	return cmd
}
