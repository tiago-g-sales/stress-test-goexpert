/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// executeRequestCmd represents the executeRequest command
var executeRequestCmd = &cobra.Command{
	Use:   "executeRequest",
	Short: "Executa requisições baseado nos parametros de url, quantidade de requisições e número de chamadas simultâneas",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executeRequest called")
	},
}

func init() {
	rootCmd.AddCommand(executeRequestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// executeRequestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// executeRequestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
