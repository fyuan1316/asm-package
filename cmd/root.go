package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	registry string
	//username string
	//passwd   string
	output   string
	helmCmd string
	dockerCmd string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "asm-package",
	Short: "asm打包工具",
	Long: `asm-package 是服务网格独立部署的命令行打包工具.
`,
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
	//cobra.OnInitialize(initConfig)
	//kubeCfgDefault := fmt.Sprintf("%s/.kube/config", homedir.HomeDir())
	rootCmd.PersistentFlags().StringVar(&registry, "registry", "build-harbor.alauda.cn", "registry")
	//rootCmd.PersistentFlags().StringVar(&username, "username", "Jian_Liao", "username of registry")
	//rootCmd.PersistentFlags().StringVar(&passwd, "passwd", "Asm@1234", "password of registry")
	rootCmd.PersistentFlags().StringVar(&output, "output", "/tmp", "output directory")

	rootCmd.PersistentFlags().StringVar(&helmCmd, "helmCmd", "helm", "helm v3 binary")
	rootCmd.PersistentFlags().StringVar(&dockerCmd, "dockerCmd", "docker", "docker binary")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	//if err := viper.ReadInConfig(); err == nil {
	//	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	//}
}
