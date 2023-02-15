package generator

import (
	"bufio"
	"errors"
	"fmt"
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

func (r repository) GenerateService(name string) error {
	_, err := os.Stat("config")
	if os.IsNotExist(err) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(fmt.Sprintf("%s is OK!, want to  create (Y/N)? ", name))

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
	}
	return nil
}
