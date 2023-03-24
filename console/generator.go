package console

import (
	"fmt"
	"log"

	"github.com/notblessy/bless/generator"
	"github.com/spf13/cobra"
)

var generatorCmd = &cobra.Command{
	Use:   "tolong-generate [name] [origin]",
	Short: "a service generator",
	Long: `
	  __    __              
	 / /_  / /__  __________
  / __ \/ / _ \/ ___/ ___/
 / /_/ / /  __(__  |__  ) 
/_.___/_/\___/____/____/                    
"tolong-generate" can generate project with one hit, example 'bless tolong-generate go-service github.com/notblessy`,
	Args: cobra.ExactArgs(3),
	Run:  generateProject,
}

func generateProject(cmd *cobra.Command, args []string) {
	if args[0] != "" && args[1] != "" && args[2] != "" {
		s := generator.NewServiceGenerator()
		err := s.GenerateService(args[0], args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("scaffolding finished with name: %s", args[0])
	} else {
		log.Fatal("expected argument <service-name> <git-origin> <git-project>' ")
	}
}
