package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "goto"
	app.Usage = "Various"

	app.Action = func(c *cli.Context) error {

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func executeCommand(command string, args []string) error {
	fmt.Println("this works")

	c := exec.Command(command, args...)
	var out, outErr bytes.Buffer

	c.Stdout = &out
	c.Stderr = &outErr

	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func retrievePidArray(in string) []string {
	var result []string

	iterator := strings.Split(in, " ")

	for _, val := range iterator {
		pid := parseNumber(val)
		if len(pid) >= 4 {
			br := val
			result = append(result, parseNumber(br))
		}
	}
	output := result
	return output
}

func removeNewline(in string) string {
	trimmed := strings.TrimSpace(in)
	inStrAry := strings.Split(trimmed, "\n")

	var result []string
	for i, s := range inStrAry {
		if (s == " ") && (inStrAry[i+1] != " ") {
			result = append(result, s)
		} else if s != " " {
			result = append(result, s)
		}
	}
	return strings.Join(result, " ")
}

// Parse number will remove any non digit character from a string.
func parseNumber(in string) string {
	usethis := in
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(usethis, "")
}
