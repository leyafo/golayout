package daemon

import (
	"github.com/spf13/pflag"
	"os"
)

type Flags struct {
	ConfigFile   string
	RegisterAddr string
	PrintVersion bool
	flags        *pflag.FlagSet
}

func ParseFlags() (*Flags, error) {
	opt := &Flags{
		flags: pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError),
	}

	opt.PrintVersion=false
	opt.flags.StringVarP(&opt.ConfigFile, "config", "f", "", "load the config file name")
	opt.flags.StringVarP(&opt.RegisterAddr, "registerAddr", "r", "", "listening the special address")
	opt.flags.BoolVarP(&opt.PrintVersion, "version", "v", false, "print version")
	err := opt.flags.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}
	return opt, nil
}
