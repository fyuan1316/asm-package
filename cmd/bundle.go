package cmd

import (
	"flag"
	"fmt"
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
		// bin/asm-package bundle --asmBundleVersion=v3.7-13-ge53b7de --flaggerBundleVersion=v3.7-3-ga0a14d5
		asmVersion := viper.GetString("asmBundleVersion")
		flaggerVersion := viper.GetString("flaggerBundleVersion")
		output := viper.GetString("output")
		dockerCmd := viper.GetString("dockerCmd")

		typ := "bundle"
		params := map[string]string{
			"Registry":             "build-harbor.alauda.cn",
			"DockerBin":            dockerCmd,
			"AsmBundleVersion":     asmVersion,     // "v3.7-13-ge53b7de",
			"FlaggerBundleVersion": flaggerVersion, // "v3.7-3-ga0a14d5",
			"Destination":          output,
			"BundleName":           "asm-bundle.tar",
		}
		err := pkg.Download(typ, params)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Exported %s to %s\n", params["BundleName"], params["Destination"])
	},
}

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
