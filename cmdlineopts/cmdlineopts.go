package cmdlineopts

import (
	"strings"

	"github.com/timvaillancourt/go-mongodb-config/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const serverCommand = "getCmdLineOpts"

// CmdLineOpts is a struct for the response of the 'getCmdLineOpts' MongoDB server command
type CmdLineOpts struct {
	Argv   []string       `bson:"argv"`
	Parsed *config.Config `bson:"parsed"`
	Ok     int            `bson:"ok"`
}

// Cmdline returns a string describing the starting command line of the MongoDB process
func (opts *CmdLineOpts) Cmdline() string {
	return strings.Join(opts.Argv, " ")
}

// Config returns a github.com/timvaillancourt/go-mongodb-config/config.Config struct using the 'getCmdLineOpts' output
func (opts *CmdLineOpts) Config() *config.Config {
	return opts.Parsed
}

// New runs the 'getCmdLineOpts' MongoDB server command on a mgo.Session and returns a pointer to a CmdLineOpts struct and an error
func New(session *mgo.Session) (*CmdLineOpts, error) {
	opts := &CmdLineOpts{}
	err := session.Run(bson.D{{serverCommand, "1"}}, opts)
	return opts, err
}
