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

const DefaultPullerScript string = `#!/usr/bin/env zsh
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

const DefaultGeneratorScript string = `#!/usr/bin/env bash
name=$1;
origin=$2;
find $name -type f -exec sed -i '' "s!skelago!$name!g" {} \;
find $name -type f -exec sed -i '' "s!github.com/notblessy!$origin!g" {} \;
cp $name/.env.sample $name/.env;
cd $name;
go mod tidy;
`
