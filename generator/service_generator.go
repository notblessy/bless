package generator

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/notblessy/serv/utils"
	"log"
	"os"
	"runtime"
	"strings"
)

type (
	//Service define repository generator
	Service interface {
		GenerateService(name string, gitOrigin string) error
	}

	repository struct{}
)

// NewServiceGenerator :nodoc:
func NewServiceGenerator() Service {
	return &repository{}
}

func (r *repository) GenerateService(name string, gitOrigin string) error {
	fmt.Println(utils.DefaultServBumper)

	_, err := os.Stat(name)
	if os.IsExist(err) {
		log.Fatal(err)
	}

	err = r.createNewProjectDirectory(name)
	fatalOnError(err)

	err = r.initGoService(name, gitOrigin)
	fatalOnError(err)

	err = os.Mkdir(fmt.Sprintf("%s/config", name), os.ModePerm)
	fatalOnError(err)

	err = os.Mkdir(fmt.Sprintf("%s/console", name), os.ModePerm)
	fatalOnError(err)

	err = os.Mkdir(fmt.Sprintf("%s/model", name), os.ModePerm)
	fatalOnError(err)

	err = os.Mkdir(fmt.Sprintf("%s/usecase", name), os.ModePerm)
	fatalOnError(err)

	err = os.Mkdir(fmt.Sprintf("%s/repository", name), os.ModePerm)
	fatalOnError(err)

	err = os.Mkdir(fmt.Sprintf("%s/utils", name), os.ModePerm)
	fatalOnError(err)

	err = os.Mkdir(fmt.Sprintf("%s/delivery", name), os.ModePerm)
	fatalOnError(err)

	return nil
}

func (r *repository) createNewProjectDirectory(name string) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(fmt.Sprintf("%s is OK!, want to scaffold a go service (Y/N)? ", name))

	input, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("fail when read user input")
	}

	ans := strings.Contains(strings.ToUpper(input), "Y")
	if !ans {
		return errors.New("scaffold service aborted")
	}

	if err := os.Mkdir(name, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (r *repository) initGoService(name, gitOrigin string) error {
	or := strings.Split(gitOrigin, ".")
	if len(or) == 2 {
		_, ok := utils.SupportedGit[or[0]]
		if !ok {
			log.Fatal("unknown git")
		}
	}

	gomod, err := os.Create(fmt.Sprintf(`%s/go.mod`, name))
	if err != nil {
		return errors.New("error generate go mod")
	}

	goversion := runtime.Version()
	_, err = gomod.WriteString(utils.DefaultGoMod(name, gitOrigin, goversion[len(goversion)-4:]))
	if err != nil {
		return err
	}

	main, err := os.Create(fmt.Sprintf(`%s/main.go`, name))
	if err != nil {
		return err
	}

	_, err = main.WriteString(utils.DefaultMain)
	if err != nil {
		return err
	}

	return nil
}

func fatalOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
