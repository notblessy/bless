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
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		err = r.checkProjectIfExist(name)
	}

	or := strings.Split(gitOrigin, ".")
	if len(or) == 2 {
		_, ok := utils.SupportedGit[or[0]]
		if !ok {
			log.Fatal("unknown git")
		}
	}

	fmt.Println(utils.DefaultServBumper)

	gomod, err := os.Create(fmt.Sprintf(`%s/go.mod`, name))
	if err != nil {
		log.Fatal("error create go mod")
	}

	goversion := runtime.Version()
	_, err = gomod.WriteString(utils.DefaultGoMod(name, gitOrigin, goversion[len(goversion)-4:]))
	fatalOnError(err)

	main, err := os.Create(fmt.Sprintf(`%s/main.go`, name))
	fatalOnError(err)

	_, err = main.WriteString(utils.DefaultMain)
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

func (r *repository) checkProjectIfExist(name string) error {
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

func fatalOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
