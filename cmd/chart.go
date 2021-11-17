package cmd

import (
	"flag"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//"gitlab-ce.alauda.cn/micro-service/asm-global-controller/cmd/upgrade/upgrade_from_3.4.x/k8s"
	//"gitlab-ce.alauda.cn/micro-service/asm-global-controller/cmd/upgrade/upgrade_from_3.4.x/k8s/model"
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

		//client, err := k8s.NewClient(k8s.ClientArgs{
		//	KubeConfig: viper.GetString("kubeconfig"),
		//})
		//if err != nil {
		//	_, _ = fmt.Fprint(os.Stderr, err.Error())
		//	return
		//}
		//list, err := client.ListMesh()
		//if err != nil {
		//	_, _ = fmt.Fprint(os.Stderr, err.Error())
		//	return
		//}
		//printData(list)
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
	chartCmd.Flags().String("chartVersion", "", "chart-global-asm的version")

}
