package utils

var SupportedGit = map[string]interface{}{
	"github":    true,
	"gitlab":    true,
	"bitbucket": true,
}

const DefaultServBumper string = `
    __    __              
   / /_  / /__  __________
  / __ \/ / _ \/ ___/ ___/
 / /_/ / /  __(__  |__  ) 
/_.___/_/\___/____/____/    
  bless service scaffolder
`
