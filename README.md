# jenerate
Generates [Jenkinsfiles](https://jenkins.io/doc/book/pipeline/jenkinsfile) from documents

## How
**jenerate** grabs **Documents** from the [registry](https://github.com/haggis-io/registry) and constructs a Jenkinfile from multiple documents.

### Usage
#### Prerequistities
* A [Registry](https://github.com/haggis-io/registry) is deployed
```bash
./jenerate # Will start the application with defaults
# For help use the -h flag
```

### Development
#### Prerequistities
* Go 1.9.x
* [glide](https://github.com/Masterminds/glide)
* Make

```bash
go get -d github.com/haggis-io/jenerate
cd $GOPATH/src/github.com/haggis-io/jenerate
make build
```
