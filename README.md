# Blinchik-git
Handy Git Automations

[![Go Report Card](https://goreportcard.com/badge/github.com/blinchik/git)](https://goreportcard.com/report/github.com/blinchik/git)

## Cloning Gitlab Projects

This command allow you to clone all Projects under a given Gitlab Group and all of it's Subgroubs. 

The Gitlab private access token should be set as environment variable `"GITLAB_PRIVATE_TOKEN"`


```sh
# Build and move to user available scripts

cd cmd/cloneall
go build
mv cloneall /usr/local/bin/

# Execute by given Gitlab flag "gl", Group ID (i.e. "10101010") 
# and the path to the private key relative to user's home path (i.e. "/.ssh/id_rsa")

cloneall -gl "<Gitlab Group ID>" "<Path to Private Key>"
```



