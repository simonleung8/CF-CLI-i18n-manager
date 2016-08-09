package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/simonleung8/flags"
)

type i18n_entry struct {
	ID          string `json:"id"`
	Translation string `json:"translation"`
}

type i18n_resources []i18n_entry

var (
	fc     flags.FlagContext
	logger io.Writer
)

func init() {
	logger = os.Stdout

	fc = flags.New()
	fc.NewStringFlag("input-file", "i", "path to the input language file")
	fc.Parse(os.Args...) //parse the OS arguments
}

func main() {

	inputFile := fc.String("i")

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Fprintln(logger, "Input file", inputFile, "not found")
		os.Exit(1)
	} else if err != nil {
		fmt.Fprintln(logger, "Error:", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintln(logger, "Error reading file:", inputFile)
		os.Exit(1)
	}

	translations := i18n_resources{}
	err = json.Unmarshal(b, &translations)

	fmt.Printf("%#v\n", len(translations))
}
