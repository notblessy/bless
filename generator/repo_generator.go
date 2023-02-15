package generator

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type (
	//Repository define repository generator
	Repository interface {
		Generate(name string) error
	}

	repository struct{}
)

// NewRepositoryGenerator :nodoc:
func NewRepositoryGenerator() Repository {
	return &repository{}
}

func (r repository) Generate(name string) error {
	_, err := os.Stat("repository/model/")
	if os.IsNotExist(err) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("repository/model/ folder not found, want to  create (Y/N)? ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return errors.New("fail when read user input")
		}
		ans := strings.Contains(strings.ToUpper(input), "Y")
		if !ans {
			return errors.New("cancel create repository & model")
		}

		fmt.Println("SUKSES", name)
	}
	return nil
}
