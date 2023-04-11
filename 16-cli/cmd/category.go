/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// name, err := cmd.Flags().GetString("name")
		/*fmt.Println("category called", category)
		exists, err := cmd.Flags().GetBool("exists")
		fmt.Println("category called", exists, err)
		id, err := cmd.Flags().GetInt16("id")
		fmt.Println("category called", id, err)*/
		cmd.Help()
	},
	/*PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamado antes do run")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamado depois do run")
	},*/
	/*RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("Ocorreu um erro")
	},*/
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
	/*categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "Y", "Category name")
	categoryCmd.PersistentFlags().BoolP("exists", "e", false, "Category exists")
	categoryCmd.PersistentFlags().Int16P("id", "i", 0, "Category id")*/
}
