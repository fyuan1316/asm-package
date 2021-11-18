package cmd

import (
	"flag"

	"github.com/fyuan1316/asm-package/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// bundleCmd represents the list command
var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "打包chart-global-asm及镜像",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			panic(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// v3.7.0-alpha.681
		chartFolder := viper.GetString("chartFolder")
		output := viper.GetString("output")
		pkg.DownloadImages(chartFolder,output)

	},
}

func init() {
	rootCmd.AddCommand(imagesCmd)
	imagesCmd.Flags().String("chartFolder", "/tmp/global-asm", "chart-global-asm folder path")
	err := imagesCmd.MarkFlagRequired("chartFolder")
	if err != nil {
		panic(err)
	}
}
