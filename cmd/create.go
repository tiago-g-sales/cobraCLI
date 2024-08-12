/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (


	"github.com/spf13/cobra"
	"github.com/tiago-g-sales/cobraCLI/internal/database"

)

func newCreateCmd(categoryDb database.Category) *cobra.Command{
	return &cobra.Command{
		Use: "create",
		Short: "Create a new category",
		Long: "Create a new cateory",
		RunE: runCreate(categoryDb),


	}
}



func runCreate(categoryDb database.Category) RunEFunc{
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		_, err := categoryDb.CreateCategory(name, description)
		if err != nil {
			return err
		}
		return nil
	}

}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(getDB()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the category")
	createCmd.Flags().StringP("description", "d", "", "Description of the category")
	createCmd.MarkFlagsRequiredTogether("name", "description")


}
