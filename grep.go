package main

import (
	"log"
	"os"
	"strings"

	"github.com/manif3station/shared_lib"
)

func _grep(item string) {
	seen := map[string]bool{}
	files := []string{"-p"}
	lno := ""

	for _, line := range strings.Split(shared_lib.SystemOutputOnly("grep", os.Args[2:]...), "\n") {
		result := strings.Split(line, ":")
		if len(result) == 0 {
			continue
		} else if len(result) > 1 {
			file := result[0]
			lno = result[1]
			if seen[file] {
				continue
			} else {
				seen[file] = true
			}
			files = append(files, file)
		}
	}

	if len(files) == 2 {
		files = append(files, "+"+lno)
	}

	if len(files) > 1 {
		shared_lib.Exec("vim", files...)
	} else {
		log.Fatal("Found nothing.")
	}

	os.Exit(0)
}
