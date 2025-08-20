package comb

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Context struct {
	Rest []string
}

func Parse(args []string, cli any) (Context, error) {
	unparsedArgs := make([]string, len(args))
	copy(unparsedArgs, args)
	rValue := reflect.ValueOf(cli).Elem()
	rType := reflect.TypeOf(cli).Elem()

	if err := parse(&unparsedArgs, 0, rValue, rType); err != nil {
		return Context{Rest: unparsedArgs}, err
	}

	return Context{Rest: unparsedArgs}, nil
}

func parse(unparsedArgs *[]string, currentArgIndex int, cliReflectValue reflect.Value, cliReflectType reflect.Type) error {
	for i := range *unparsedArgs {
		arg := &((*unparsedArgs)[i])

		if *arg == "" {
			continue
		}

		if err := set(arg, unparsedArgs, i, cliReflectValue, cliReflectType); err != nil {
			return err
		}
	}

	return nil
}

func set(arg *string, unparsedArgs *[]string, currentArgIndex int, cliReflectValue reflect.Value, cliReflectType reflect.Type) error {
	for i := 0; i < cliReflectType.NumField(); i++ {
		field := cliReflectType.Field(i)
		fieldReflectValue := cliReflectValue.FieldByName(field.Name)

		cmdName, ok := field.Tag.Lookup("cmd")
		if ok && strings.EqualFold(*arg, cmdName) {
			// fmt.Printf("(TODO) CMD: %s\n", cmdName)
			subcommandField := fieldReflectValue.Addr().Interface()
			rValue := reflect.ValueOf(subcommandField).Elem()
			rType := reflect.TypeOf(subcommandField).Elem()

			*arg = ""
			if err := parse(unparsedArgs, currentArgIndex, rValue, rType); err != nil {
				return err
			}
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
				if (currentArgIndex + 1) >= len(*unparsedArgs) {
					return fmt.Errorf("flag \"%s\": missing argument of type 'string'", *arg)
				}

				fieldReflectValue.SetString((*unparsedArgs)[currentArgIndex+1])
				*arg = ""
				(*unparsedArgs)[currentArgIndex+1] = ""
			case reflect.Int:
				if (currentArgIndex + 1) >= len(*unparsedArgs) {
					return fmt.Errorf("flag \"%s\": missing argument of type 'int'", *arg)
				}

				parsedArg, err := strconv.ParseInt((*unparsedArgs)[currentArgIndex+1], 10, 64)
				if err != nil {
					return fmt.Errorf("flag \"%s\": %w", arg, err)
				}

				fieldReflectValue.SetInt(parsedArg)
				*arg = ""
				(*unparsedArgs)[currentArgIndex+1] = ""
			case reflect.Float64:
				if (currentArgIndex + 1) >= len(*unparsedArgs) {
					return fmt.Errorf("flag \"%s\": missing argument of type 'int'", *arg)
				}

				parsedArg, err := strconv.ParseFloat((*unparsedArgs)[currentArgIndex+1], 64)
				if err != nil {
					return fmt.Errorf("flag \"%s\": %w", arg, err)
				}

				fieldReflectValue.SetFloat(parsedArg)
				*arg = ""
				(*unparsedArgs)[currentArgIndex+1] = ""
			default:
				return fmt.Errorf("flag \"%s\" unsupported field type %v", *arg, fieldReflectValue.Type())
			}

			return nil
		}
	}

	return nil
}
