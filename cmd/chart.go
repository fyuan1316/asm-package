package cmd

import (
	"flag"
	"log"

	"github.com/fyuan1316/asm-package/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// bundleCmd represents the list command
var chartCmd = &cobra.Command{
	Use:   "chart",
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
		chartVersion := viper.GetString("chartVersion")
		output := viper.GetString("output")
		params := map[string]string{
			"Registry": "build-harbor.alauda.cn",
			"User":     "Jian_Liao",
			"Password": "Asm@1234",
			// "DockerBin": "docker",
			"HelmBin":      "helm3",
			"ChartVersion": chartVersion, // "v3.7-13-ge53b7de",
			"Destination":  output,
		}
		err := pkg.Download("chart", params)
		if err != nil {
			log.Fatal(err)
		}
	},
}

//func printData(data []model.MeshInfo) {
//	if len(data) == 0 {
//		return
//	}
//	var columns [][]string
//	columns = append(columns, []string{"ServiceMesh", "Cluster"})
//	for _, d := range data {
//		columns = append(columns, []string{d.Name, d.Cluster})
//	}
//	if err := pterm.DefaultTable.WithHasHeader().WithData(columns).Render(); err != nil {
//		panic(err)
//	}
//}

func init() {
	rootCmd.AddCommand(chartCmd)
	chartCmd.Flags().String("chartVersion", "", "chart-global-asm version")
	err := chartCmd.MarkFlagRequired("chartVersion")
	if err != nil {
		panic(err)
	}
}
