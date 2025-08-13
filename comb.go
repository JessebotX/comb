package comb

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Context struct{}

func Parse(args []string, cli any) (Context, error) {
	unparsedArgs := make([]string, len(args))
	copy(unparsedArgs, args)

	rValue := reflect.ValueOf(cli).Elem()
	rType := reflect.TypeOf(cli).Elem()

	for i := range unparsedArgs {
		arg := &unparsedArgs[i]

		if *arg == "" {
			continue
		}

		if err := set(arg, &unparsedArgs, i, rValue, rType); err != nil {
			return Context{}, err
		}
	}

	return Context{}, nil
}

func set(arg *string, unparsed *[]string, currentArgIndex int, cliReflectValue reflect.Value, cliReflectType reflect.Type) error {
	for i := 0; i < cliReflectType.NumField(); i++ {
		field := cliReflectType.Field(i)
		fieldReflectValue := cliReflectValue.FieldByName(field.Name)

		cmdName, ok := field.Tag.Lookup("cmd")
		if ok && strings.EqualFold(*arg, cmdName) {
			fmt.Printf("(TODO) CMD: %s\n", cmdName)
			return nil
		}

		flagName, ok := field.Tag.Lookup("flag")
		if ok && strings.EqualFold(*arg, "-"+flagName) {
			if !fieldReflectValue.CanSet() {
				return fmt.Errorf("flag \"%s\": cannot be given a value", *arg)
			}

			switch fieldReflectValue.Kind() {
			case reflect.Bool:
				fieldReflectValue.SetBool(true)
				*arg = ""
			case reflect.String:
				if (currentArgIndex + 1) >= len(*unparsed) {
					return fmt.Errorf("flag \"%s\": missing argument of type 'string'", *arg)
				}

				fieldReflectValue.SetString((*unparsed)[currentArgIndex+1])
				*arg = ""
				(*unparsed)[currentArgIndex+1] = ""
			case reflect.Int:
				if (currentArgIndex + 1) >= len(*unparsed) {
					return fmt.Errorf("flag \"%s\": missing argument of type 'int'", *arg)
				}

				parsedArg, err := strconv.ParseInt((*unparsed)[currentArgIndex+1], 10, 64)
				if err != nil {
					return fmt.Errorf("flag \"%s\": %w", parsedArg)
				}

				fieldReflectValue.SetInt(parsedArg)
				*arg = ""
				(*unparsed)[currentArgIndex+1] = ""
			case reflect.Float64:
				if (currentArgIndex + 1) >= len(*unparsed) {
					return fmt.Errorf("flag \"%s\": missing argument of type 'int'", *arg)
				}

				parsedArg, err := strconv.ParseFloat((*unparsed)[currentArgIndex+1], 64)
				if err != nil {
					return fmt.Errorf("flag \"%s\": %w", parsedArg)
				}

				fieldReflectValue.SetFloat(parsedArg)
				*arg = ""
				(*unparsed)[currentArgIndex+1] = ""
			default:
				return fmt.Errorf("flag \"%s\" unsupported field type %v", *arg, fieldReflectValue.Type())
			}

			return nil
		}
	}

	return nil
}
