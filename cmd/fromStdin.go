package cmd

import (
	"os"
	"strings"

	pipeline "github.com/mattn/go-pipeline"
)

func getPathFromStdin() string {
	path := "/tmp/cdxStdin"
	if _, err := os.Stat(path); err != nil {
		Fatal(err)
	}
	out, err := pipeline.Output(
		[]string{"cat", "/tmp/cdxStdin"},
		config.FuzzyFinder.GetCommand(),
		[]string{"awk", "{print $NF}"},
	)

	if err != nil {
		Fatal(err)
	}

	return strings.Trim(string(out), "\n")
}
