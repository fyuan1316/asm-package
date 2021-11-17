package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"

	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "asm-package",
	Short: "asm打包工具",
	Long: `asm-package 是服务网格独立部署的命令行打包工具.
第一步:
使用 --list 命令查看计划升级的服务网格以及所在集群信息.
第二步:
使用 --upgrade 命令执行升级.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		// Allow unknown flags for backward-compatibility.
		UnknownFlags: true,
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	kubeCfgDefault := fmt.Sprintf("%s/.kube/config", homedir.HomeDir())
	rootCmd.PersistentFlags().String("kubeconfig", kubeCfgDefault, "kube config path")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	//if err := viper.ReadInConfig(); err == nil {
	//	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	//}
}
