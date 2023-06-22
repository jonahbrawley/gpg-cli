package cmd

import (
	//"fmt"
	//"os"

	"github.com/spf13/cobra"
)

var (
	// flags
	// flag1	string
	// flag2	string

	rootCmd = &cobra.Command{
		Use: "pkey",
		Short: "A GPG private key helper",
		Long: `A CLI for checking status of GPG keys on your local machine`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// config
	// cobra.OnInitialize(initConfig) // initConfig = func

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "Jonah Brawley", "author name for copyright attribution")
}