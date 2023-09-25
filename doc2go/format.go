package main

import (
	"fmt"
	"io"
)

type formatConfig struct {
	w            io.Writer
	packageName  string
	variableName string
	tjPrefix     string
	tjPackage    string
}

func formatObject(fcfg formatConfig, obj map[string]interface{}) {
	if len(obj) == 0 {
		fmt.Fprintf(fcfg.w, "%s.O{}", fcfg.tjPrefix)
		return
	}
	fmt.Fprintf(fcfg.w, "%s.O{\n", fcfg.tjPrefix)
	for k, v := range obj {
		fmt.Fprintf(fcfg.w, "%q:", k)
		format(fcfg, v)
		fmt.Fprintf(fcfg.w, ",\n")
	}
	fmt.Fprintf(fcfg.w, "}")
}

func formatArray(fcfg formatConfig, arr []interface{}) {
	if len(arr) == 0 {
		fmt.Fprintf(fcfg.w, "%s.A{}", fcfg.tjPrefix)
		return
	}
	if len(arr) == 1 {
		fmt.Fprintf(fcfg.w, "%s.A{", fcfg.tjPrefix)
		format(fcfg, arr[0])
		fmt.Fprintf(fcfg.w, "}")
		return
	}
	fmt.Fprintf(fcfg.w, "%s.A{\n", fcfg.tjPrefix)
	for _, elem := range arr {
		format(fcfg, elem)
		fmt.Fprintf(fcfg.w, ",\n")
	}
	fmt.Fprintf(fcfg.w, "}")
}

func format(fcfg formatConfig, data interface{}) {
	switch v := data.(type) {
	case nil:
		fmt.Fprintf(fcfg.w, "null")
	case bool:
		fmt.Fprintf(fcfg.w, "%t", v)
	case int:
		fmt.Fprintf(fcfg.w, "%d", v)
	case float64:
		fmt.Fprintf(fcfg.w, "%g", v)
	case string:
		fmt.Fprintf(fcfg.w, "%q", v)
	case []interface{}:
		formatArray(fcfg, v)
	case map[string]interface{}:
		formatObject(fcfg, v)
	default:
		panic(fmt.Errorf("unexpected type in JSON deserialization of %v: %T", data, data))
	}
}

func formatHeader(fcfg formatConfig) {
	if fcfg.packageName != "" {
		fmt.Fprintf(fcfg.w, "package %s\n\n", fcfg.packageName)
		if fcfg.tjPrefix != "tj" {
			fmt.Fprintf(fcfg.w, "import %s \"%s\"\n\n", fcfg.tjPrefix, fcfg.tjPackage)
		} else {
			fmt.Fprintf(fcfg.w, "import \"%s\"\n\n", fcfg.tjPackage)
		}
	}
	if fcfg.variableName != "" {
		fmt.Fprintf(fcfg.w, "var %s = ", fcfg.variableName)
	}
}
