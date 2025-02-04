/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/spf13/viper"
	"github.com/tiago-g-sales/stress-test-goexpert/cmd"
)

// load env vars cfg
func init() {
	viper.AutomaticEnv()
}



func main() {
	cmd.Execute()
}
