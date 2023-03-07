package generator

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/notblessy/serv/utils"
)

type (
	//Service define repository generator
	Service interface {
		GenerateService(name string, gitOrigin string) error
	}

	repository struct {
		name      string
		gitOrigin string
	}
)

const (
	skelago    = "skelago.sh"
	bash       = "bash"
	scaffolder = "scaffolder.sh"
)

// NewServiceGenerator :nodoc:
func NewServiceGenerator() Service {
	return &repository{}
}

func (r *repository) GenerateService(name string, gitOrigin string) error {
	r.name = name
	r.gitOrigin = gitOrigin

	fmt.Println(utils.DefaultServBumper)

	err := r.checkExistingProjectDirectory(name)
	fatalOnError(err)

	// Create project root directory
	err = os.Mkdir(r.name, os.ModePerm)
	fatalOnError(err)

	r.pullSkelago(r.name)
	r.generateScaffoldScript()
	r.scaffold(r.name)

	return nil
}

func (r *repository) checkExistingProjectDirectory(name string) error {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
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

		return nil
	}

	return errors.New("directory existed")
}

func (r repository) pullSkelago(name string) {
	body := `#!/usr/bin/env zsh
	name=$1;
	cd $name;
	git init;
	git remote add origin git@github.com:notblessy/skelago.git;
	git remote -v;
	git fetch;	
	git pull origin main;
	rm -rf .git;
	cd ..;
`
	bt := []byte(body)
	err := ioutil.WriteFile(skelago, bt, os.ModePerm)
	if err != nil {
		r.clearOnError("fail creating pull script sh")
	}
	defer func() {
		_ = os.Remove(skelago)
	}()

	cmd := exec.Command(bash, skelago, name)
	err = r.runScript(cmd)
	if err != nil {
		log.Fatal("fail running pull script sh")
	}

}

// generateScaffoldScript creates scaffold bash script
func (r repository) generateScaffoldScript() {
	contents := `#!/usr/bin/env bash
name=$1;
find $name -type f -exec sed -i '' "s/skelago/$name/g" {} \;
cp $name/.env.example $name/.env;
cd $name;
go mod tidy;
`
	bt := []byte(contents)
	err := ioutil.WriteFile(scaffolder, bt, 0644)
	if err != nil {
		r.clearOnError("fail when create scaffold script")
	}
}

// scaffold runs scaffold script
func (r repository) scaffold(name string) {
	cmd := exec.Command(bash, scaffolder, name)
	defer func() {
		_ = os.Remove(scaffolder)
	}()
	err := r.runScript(cmd)
	if err != nil {
		fmt.Println(err)
		r.clearOnError("fail scaffolding")
	}
}

func (r repository) runScript(cmd *exec.Cmd) error {
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (r repository) clearOnError(msg string) {
	_ = os.RemoveAll(r.name)
	_ = os.Remove(skelago)
	log.Fatal(msg)
}

func fatalOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
