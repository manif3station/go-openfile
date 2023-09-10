package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/manif3station/shared_lib"
)

func _select(line string, found ...string) {
	found = shared_lib.Grep(found, func(path string) bool {
		return !regexp.MustCompile(`skip$`).MatchString(path)
	})

	if len(found) == 0 {
		return
	}

	seen := map[string]bool{}

	for_vim := false

	if os.Getenv("FOR_VIM") != "" {
		for_vim = true
	}

	list_file := ""

	if for_vim {
		tmpdir, err := ioutil.TempDir("dir", "prefix")
		shared_lib.CheckErr(err)
		list_file = tmpdir + "/list.txt"
	}

	for pos, file := range found {
		if seen[file] {
			continue
		}
		seen[file] = true

		output := ""

		if len(found) == 1 {
			output = file + "\n"
		} else {
			output = fmt.Sprintf("%d: %s\n", pos+1, file)
		}

		if for_vim {
			err := os.WriteFile(list_file, []byte(output), 0644)
			shared_lib.CheckErr(err)
		} else {
			fmt.Print(output)
		}
	}

	if for_vim {
		fmt.Println(list_file)
		os.Exit(0)
	}

	choices := ""

	if len(seen) == 1 {
		choices = "1"
	} else {
		fmt.Print("> ")
		choices, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		choices = strings.Replace(choices, "\n", "", -1)
	}

	choosen := []string{}

	if choices != "" {
		for _, choice := range shared_lib.Split(`[,\s]`, choices, -1) {
			if choice == "" {
				continue
			}

			if num := shared_lib.Replace(`\D`, choice, ""); num != "" {
				pos, _ := strconv.ParseInt(num, 10, 32)
				choosen = append(choosen, found[pos-1])
			}
		}
	} else {
		choosen = found
	}

	args := []string{"-p"}

	for _, v := range choosen {
		args = append(args, v)
	}

	if line == "" {
		shared_lib.Exec("vim", args...)
	} else {
		if line != "" && !strings.Contains(line, "+") {
			line = "+" + line
		}
		args = append(args, line)
		shared_lib.Exec("vim", args...)
	}

	os.Exit(0)
}
