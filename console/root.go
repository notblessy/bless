package console

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bless",
	Short: "bless is go service generator",
	Long: `
	  __    __              
	 / /_  / /__  __________
  / __ \/ / _ \/ ___/ ___/
 / /_/ / /  __(__  |__  ) 
/_.___/_/\___/____/____/                         
bless can scaffold a go service like a toss.`,
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
