package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var verbose bool
var check bool
var parallel int

var RootCmd = &cobra.Command{
	Use:   "clug",
	Short: "A simple retrieval tool for ocean datasets hosted by the NODD",
	Long: `A CLI library for downloading ocean (bathymetry, trackline, and water column)
	data from the NOAA Open Data Dissemination (NODD) cloud on a survey by survey basis.
	This simplifies s3 object retrieval, which will almost always need to be downloaded 
	in batch groups, avoiding downloading each file object manually. 

	get, given a survey name argument or path, will download all survey files or 
	sub-survey files at a given path.

	glance will summarize all files and file sizes for an equivalent get command

	list will display all files that will be downloaded for an equivalent get command

	config can be used to change default bucket name and download parameters.`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
	RootCmd.PersistentFlags().IntVarP(&parallel, "parallel", "p", 3, "Number of parallel downloads. (default: 3, max: 100)")
	RootCmd.PersistentFlags().BoolVarP(&check, "check", "c", true, "Check local disk space before downloading. (default: true)")

	RootCmd.PersistentFlags().String("source", "s", "A help for foo")

	vErr := viper.BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))
	if vErr != nil {
		log.Fatal(vErr)
	}

	pErr := viper.BindPFlag("parallel", RootCmd.PersistentFlags().Lookup("parallel"))
	if pErr != nil {
		log.Fatal(pErr)
	}

	cErr := viper.BindPFlag("check", RootCmd.PersistentFlags().Lookup("check"))
	if cErr != nil {
		log.Fatal(cErr)
	}

}
