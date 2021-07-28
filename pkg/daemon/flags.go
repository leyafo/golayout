package daemon

import (
	"github.com/spf13/pflag"
	"os"
)

type Flags struct{
	ConfigFile string
	flags *pflag.FlagSet
}

func NewCmdFlags()*Flags{
	opt := &Flags{
		flags: pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError),
	}

	opt.flags.StringVarP(&opt.ConfigFile, "config", "f", "", "load the config file name")

	return opt
}

func (f *Flags)Parse()(error){
	err := f.flags.Parse(os.Args[1:])
	if err != nil {
		return err
	}
	return nil
}
