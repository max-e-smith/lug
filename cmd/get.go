package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Download NOAA survey data to local path",
	Long: `Use 'clug get <survey(s)> <local path> <options>' to download marine geophysics data to your machine. 

Data is downloaded from the NOAA Open Data Dissemination cloud bucket by default. Use
the config option to change default buckets used. You must specify a data type(s) to 
download by specifying one or more of the data options for this command. View the command
help for more info on those options.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Local flags
	getCmd.Flags().BoolP("bathy", "b", false, "Download bathy data")

	// Planned local flags
	//getCmd.Flags().BoolP("trackline", "t", false, "Download trackline data")
	//getCmd.Flags().BoolP("water-column", "w", false, "Download water column data")
}
