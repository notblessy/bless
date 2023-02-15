package generator

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/notblessy/serv/utils"
	"log"
	"os"
	"strings"
)

type (
	//Service define repository generator
	Service interface {
		GenerateService(name string) error
	}

	repository struct{}
)

// NewServiceGenerator :nodoc:
func NewServiceGenerator() Service {
	return &repository{}
}

func (r *repository) GenerateService(name string) error {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		err = r.checkProjectIfExist(name)
	}

	main, err := os.Create(fmt.Sprintf(`%s/main.go`, name))

	if err != nil {
		log.Fatal("error create main")
	}

	_, err = main.WriteString(utils.MainMap)

	if err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(fmt.Sprintf("%s/config", name), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(fmt.Sprintf("%s/console", name), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(fmt.Sprintf("%s/model", name), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(fmt.Sprintf("%s/usecase", name), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(fmt.Sprintf("%s/repository", name), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(fmt.Sprintf("%s/utils", name), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(fmt.Sprintf("%s/delivery", name), os.ModePerm); err != nil {
		log.Fatal(err)
	}

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
