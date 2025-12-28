package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var Verbose bool

var rootCmd = &cobra.Command{
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
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
	err := viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	if err != nil {
		log.Fatal(err)
	}
}
