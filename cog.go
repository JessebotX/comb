package cog

import (
	"fmt"
	"reflect"
)

type Context struct{}

func Parse(args []string, cli any) (Context, error) {
	unparsedArgs := make([]string, len(args))
	copy(unparsedArgs, args)

	fmt.Printf("%#v\n", unparsedArgs)

	rValue := reflect.ValueOf(cli).Elem()
	rType := reflect.TypeOf(cli).Elem()

	for i := range unparsedArgs {
		arg := &unparsedArgs[i]

		if *arg == "" {
			continue
		}

		set(arg, rValue, rType)

		*arg = ""
	}

	fmt.Printf("%#v\n", unparsedArgs)

	return Context{}, nil
}

func set(arg *string, cliReflectValue reflect.Value, cliReflectType reflect.Type) {
	for i := 0; i < cliReflectType.NumField(); i++ {
		field := cliReflectType.Field(i)

		flagName, ok := field.Tag.Lookup("flag")
		if ok {
			if *arg == "-" + flagName {
				fmt.Printf("found flag -%s\n", flagName)
				return
			}
		}
	}
}
