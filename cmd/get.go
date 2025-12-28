package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var parallelDownloads int

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Download NOAA survey data to local path",
	Long:  `Use 'clug get <subcommand> to download a dataset from the Noaa Open Data Dissemination cloud.`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	var length = len(args)
	//	if length <= 1 {
	//		fmt.Println("Please specify a subcommand")
	//		fmt.Println(cmd.UsageString())
	//		return
	//	}
	//},
}

func init() {
	rootCmd.AddCommand(GetCmd)
	GetCmd.PersistentFlags().IntVarP(&parallelDownloads, "parallel-downloads", "p", 3, "Number of parallel downloads")
	err := viper.BindPFlag("parallel-downloads", GetCmd.PersistentFlags().Lookup("parallel-downloads"))
	if err != nil {
		log.Fatal(err)
	}
}
