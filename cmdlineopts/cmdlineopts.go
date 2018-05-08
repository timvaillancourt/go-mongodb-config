package cmdlineopts

import (
	"strings"

	"github.com/timvaillancourt/go-mongodb-config/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CmdLineOpts struct {
	Argv   []string       `bson:"argv"`
	Parsed *config.Config `bson:"parsed"`
	Ok     int            `bson:"ok"`
}

func (clo *CmdLineOpts) Cmdline() string {
	return strings.Join(clo.Argv, " ")
}

func (clo *CmdLineOpts) Config() *config.Config {
	return clo.Parsed
}

func Get(session *mgo.Session) (*CmdLineOpts, error) {
	cmdLineOpts := &CmdLineOpts{}
	err := session.Run(bson.D{{"getCmdLineOpts", "1"}}, cmdLineOpts)
	return cmdLineOpts, err
}
