package console

import (
	"fmt"
	"github.com/notblessy/serv/generator"
	"github.com/spf13/cobra"
	"log"
)

var generatorCmd = &cobra.Command{
	Use:   "generate [name]",
	Short: "a service generator",
	Long:  `generate can generate project with one hit, example 'serv generate [name]`,
	Args:  cobra.ExactArgs(1),
	Run:   generateProject,
}

func generateProject(cmd *cobra.Command, args []string) {
	if args[0] != "" {
		s := generator.NewServiceGenerator()
		err := s.GenerateService(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(args[0])
	} else {
		fmt.Println("need argument <service-name>' ")
	}
}
