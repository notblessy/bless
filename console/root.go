package console

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "serv",
	Short: "serv is go service generator",
	Long: `
   _____  ______ ____  _    __
  / ___/ / ____// __ \| |  / /
  \__ \ / __/  / /_/ /| | / / 
 ___/ // /___ / _, _/ | |/ /  
/____//_____//_/ |_|  |___/   
                              
serv can scaffold a go service like a toss.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(generatorCmd)
}
