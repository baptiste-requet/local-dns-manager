/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"vm-manager/vmm-core"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !isAdmin() {
			fmt.Println("This command needs elevated privilege. Please retry with admin permissions.")
		}

		ip, _ := cmd.Flags().GetString("ip")
		name, _ := cmd.Flags().GetString("name")

		if ip == "" && name == "" {
			fmt.Println("Both IP and name cannot be simultaneously null.")
			return
		}

		if ip != "" && name != "" {
			fmt.Println("IP and name cannot be both simultaneously defined. Please choose one.")
			return
		}

		if ip != "" {
			vmmcore.RemoveEntryByIp(ip)
		}
		if name != "" {
			vmmcore.RemoveEntryByName(name)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringP("ip", "i", "", "The IP of the VM")
	removeCmd.Flags().StringP("name", "n", "", "The name of the VM")
}
