package main

import "github.com/spf13/pflag"

func (f *UserForm) RegisterFlags(flags *pflag.FlagSet) {
	flags.StringVar(&f.Name, "name", "value string", "usage string")
	flags.StringVar(&f.Email, "email", "value string", "usage string")
	flags.IntVar(&f.Age, "Age", 0, "usage string")
}
