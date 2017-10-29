# go-mongodb-config
[![Go Report Card](https://goreportcard.com/badge/github.com/timvaillancourt/go-mongodb-config)](https://goreportcard.com/report/github.com/timvaillancourt/go-mongodb-config)

A package for reading/writing MongoDB yaml-based configurations

## Docs
- [github.com/timvaillancourt/go-mongodb-config/config](https://godoc.org/github.com/timvaillancourt/go-mongodb-config/config)

## Usage
Create 'Config' struct from file (YAML-based only):
```
import (
	mongodb_config "github.com/timvaillancourt/go-mongodb-config/config"
)

func main() {
	config, err := mongodb_config.Load("/etc/mongod.conf")
	if err != nil {
		panic(err)
	}
	...
}
```

Create 'Config' struct from uri:
```
import (
        mongodb_config "github.com/timvaillancourt/go-mongodb-config/config"
)

func main() {
        config, err := mongodb_config.LoadUri("http://example.com/etc/mongod.conf")
        if err != nil {
                panic(err)
        }
        ...
}
```

Write 'Config' struct to file:
```
import (
	mongodb_config "github.com/timvaillancourt/go-mongodb-config/config"
)
 
func main() {
	config := mongodb_config.New()
	...
	err := config.Write("/etc/mongod.conf")
	if err != nil {
		panic(err)
	}
}
```
