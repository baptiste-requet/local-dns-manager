/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"vm-manager/vmm-core"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
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
		comment, _ := cmd.Flags().GetString("comment")

		if ip != "" && name != "" {
			vmmcore.AddEntry(ip, name, comment)
		} else {
			fmt.Println("Both IP and name cannot be simultaneously null.")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("ip", "i", "", "The IP of the VM")
	addCmd.Flags().StringP("name", "n", "", "The name of the VM")
	addCmd.Flags().StringP("comment", "c", "", "The comment associated with the VM")
}
