package utils

import "fmt"

var SupportedGit = map[string]interface{}{
	"github":    true,
	"gitlab":    true,
	"bitbucket": true,
}

const DefaultServBumper string = `
   _____  ______ ____  _    __
  / ___/ / ____// __ \| |  / /
  \__ \ / __/  / /_/ /| | / / 
 ___/ // /___ / _, _/ | |/ /  
/____//_____//_/ |_|  |___/   
  serv service scaffolder
`

const DefaultMain string = `package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
`

func DefaultGoMod(name, origin, gover string) string {
	return fmt.Sprintf(`module %s/%s

go %s
`, origin, name, gover)
}
