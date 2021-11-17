package cmd

import (
	"flag"
	"log"

	"github.com/fyuan1316/asm-package/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// bundleCmd represents the list command
var bundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "打包operator镜像",
	Long: `asm operator包含:
			- asm-operator
            - flagger-operator
`,
	PreRun: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			panic(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		// v3.7-13-ge53b7de
		asmVersion := viper.GetString("asmBundleVersion")
		// v3.7-3-ga0a14d5
		flaggerVersion := viper.GetString("flaggerBundleVersion")

		//pkg.DownloadBundle(asmVersion, flaggerVersion)

		typ := "bundle"
		params := map[string]string{
			"Registry":  "build-harbor.alauda.cn",
			"User":      "Jian_Liao",
			"Password":  "Asm@1234",
			"DockerBin": "docker",
			//"HelmBin":              "helm3",
			"AsmBundleVersion":     asmVersion,     //"v3.7-13-ge53b7de",
			"FlaggerBundleVersion": flaggerVersion, //"v3.7-3-ga0a14d5",
			"Destination":          ".",
		}
		err := pkg.Download(typ, params)
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
	rootCmd.AddCommand(bundleCmd)
	bundleCmd.Flags().String("asmBundleVersion", "", "asm-operator-bundle version")
	bundleCmd.Flags().String("flaggerBundleVersion", "", "flagger-operator-bundle version")
	var err error
	err = bundleCmd.MarkFlagRequired("asmBundleVersion")
	if err != nil {
		panic(err)
	}
	err = bundleCmd.MarkFlagRequired("flaggerBundleVersion")
	if err != nil {
		panic(err)
	}
}
