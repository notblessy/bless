package console

import (
	"fmt"
	"github.com/notblessy/serv/generator"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "serv",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(console *cobra.Command, args []string) { },
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
	generatorCmd.AddCommand(generatorCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var generatorCmd = &cobra.Command{
	Use:   "repository [name]",
	Short: "repository generator",
	Long:  `Generate repository`,
	Args:  cobra.ExactArgs(1),
	Run:   generateRepository,
}

func generateRepository(cmd *cobra.Command, args []string) {
	if args[0] != "" {
		r := generator.NewRepositoryGenerator()
		err := r.Generate(args[0])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("please input <name>' ")
	}
}
