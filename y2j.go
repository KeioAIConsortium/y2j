package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	flag.Parse()

	switch flag.NArg() {
	case 0:
		yamldata, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		jsondata, err := yaml.YAMLToJSON(yamldata)
		if err != nil {
			panic(err)
		}

		fmt.Print(string(jsondata))
		break
	case 1:
		yamldata, err := ioutil.ReadFile(flag.Arg(0))
		if err != nil {
			panic(err)
		}

		jsondata, err := yaml.YAMLToJSON(yamldata)
		if err != nil {
			panic(err)
		}

		fmt.Print(string(jsondata))
		break
	default:
		fmt.Fprintf(os.Stderr, "input must be from stdin or file\n")
		os.Exit(1)
	}
}
