package comb

import (
	"fmt"
	"reflect"
	"strings"
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

		if err := set(arg, rValue, rType); err != nil {
			return Context{}, err
		}

		*arg = ""
	}

	fmt.Printf("%#v\n", unparsedArgs)

	return Context{}, nil
}

func set(arg *string, cliReflectValue reflect.Value, cliReflectType reflect.Type) error {
	for i := 0; i < cliReflectType.NumField(); i++ {
		field := cliReflectType.Field(i)

		if strings.EqualFold(*arg, "-h") || strings.EqualFold(*arg, "-help") {
			fmt.Print("FLAG: <SPECIAL HELP FLAG>\n")
			return nil
		}

		cmdName, ok := field.Tag.Lookup("cmd")
		if ok && strings.EqualFold(*arg, cmdName) {
			fmt.Printf("CMD:   %s\n", cmdName)
			return nil
		}

		flagName, ok := field.Tag.Lookup("flag")
		if ok && strings.EqualFold(*arg, "-"+flagName) {
			fmt.Printf("FLAG: -%s\n", flagName)
			return nil
		}
	}

	return nil
}
