package main

import (
	"os"

	"github.com/manif3station/shared_lib"
)

func main() {
	item := shared_lib.GetArg(1)

	if item == "" {
		os.Exit(0)
	}

	if item == "update" {
		_ = os.Chdir(shared_lib.MyHomeItem("web"))
		_ = shared_lib.Exec("vim", "OpenFile/main.go")
		os.Exit(0)
	} else if item == "grep" {
		_grep(item)
	}

	line := ""

	if found := shared_lib.Split(":", item, -1); len(found) > 0 {
		item = shared_lib.ArrayItem(found, 0)
		line = shared_lib.ArrayItem(found, 1)
	}

	_remove_git_path_prefix(&item)

	new_args := []string{}

	if _perl_pkg_search(item) {
		os.Exit(0)
	} else if found := _tt_include(item); len(found) > 0 {
		new_args = []string{"web/www/views", found[0] + ".tt"}
	} else if found := _view_files(item); len(found) > 0 {
		new_args = []string{"web/www", found[0]}
	} else if found := _css_files(item); len(found) > 0 {
		new_args = []string{"web/www", found[0] + ".(css|scss)"}
	} else if found, line, files := _perl_module_name(item); found {
		_select(line, files...)
	}

	if len(new_args) == 0 {
	}

	if shared_lib.File_exists(item) {
		if line == "" {
			line = shared_lib.GetArg(2)
		}
		_select(line, item)
	}

	if shared_lib.Dir_exists(item) || len(shared_lib.Match(`\s`, item)) > 0 {
		dir, parts := shared_lib.GetArg(1), os.Args[2:]

		dirs := shared_lib.Split(`\s`, dir, -1)

		if len(parts) == 0 {
			_select("0", dirs...)
		} else {
			line := shared_lib.ArrayItem(shared_lib.Grep(parts, func(part string) bool {
				return len(shared_lib.Match(`^\+\d+$`, part)) > 0
			}), 0)

			found := []string{}

			for _, dir := range dirs {
				for _, file := range shared_lib.Find(dir) {
					match := true
					for _, part := range parts {
						if found := shared_lib.Match(part, file); len(found) == 0 {
							match = false
							break
						}
					}
					if match {
						found = append(found, file)
					}
				}
			}

			_select(line, found...)
		}
	} else if match := shared_lib.Match(`(.+)\:(\d+)\:.*`, item); len(match) > 0 {
		file, line := match[0], match[1]
		_select(line, file)
	}
}
