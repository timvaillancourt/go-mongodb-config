# go-mongodb-config
A package for reading/writing MongoDB yaml-based configurations

## Usage
Create 'Config' from file:
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

Write 'Config' to file:

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
