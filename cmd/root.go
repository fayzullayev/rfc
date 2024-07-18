package cmd

import (
	"os"
	"rfc/generator"
	"rfc/utils"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "rfc",
	Short: "CLI tool for generating boilerplate code",
	Long:  `CLI tool for generating boilerplate code of Web applications
	Usage:
	rfc templateGenerator -p packageName   // -p or --page
	rfc templateGenerator -c button   // -c or --component
	`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		isPage, _ := cmd.Flags().GetBool("page")
		isComponent, _ := cmd.Flags().GetBool("component")

		if isPage || isComponent {
			var metaData generator.FileMetaData

			if isComponent {
				metaData = generator.GetMetaData("components", args[0])
			} else {
				metaData = generator.GetMetaData("pages", args[0])
			}

			if err := generator.CreateFile(metaData, metaData.FolderName+".tsx", "component"); err != nil {
				utils.LogErrorAndCleanup(err, metaData.FolderPath)
				return
			}

			// creating styles file
			if err := generator.CreateFile(metaData, "index.ts", "index"); err != nil {
				utils.LogErrorAndCleanup(err, metaData.FolderPath)
				return
			}

			// creating index file
			if err := generator.CreateFile(metaData, "style.ts", "style"); err != nil {
				utils.LogErrorAndCleanup(err, metaData.FolderPath)
				return
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("page", "p", false, "Generate page template")
	rootCmd.PersistentFlags().BoolP("component", "c", false, "Generate component template")
}


