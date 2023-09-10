package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/manif3station/shared_lib"
)

func _perl_pkg_search(item string) bool {
	if found := regexp.MustCompile(`package [\w\:]+`).MatchString(item); found {
		err := shared_lib.Exec("perl", shared_lib.MyHomeItem("bin")+"pkg_search", item)
		shared_lib.CheckErr(err)
		return true
	} else {
		return false
	}
}

func _perl_module_name(item string) (bool, string, []string) {
	found := false
	line := ""
	files := []string{}
	if found := strings.ContainsAny(item, "::"); found {
		item = regexp.MustCompile(`(^::|::$)`).ReplaceAllString(item, "")
		item = regexp.MustCompile(`::/)`).ReplaceAllString(item, "")
		files = shared_lib.Grep(shared_lib.Perl_inc_paths(), func(dir string) bool {
			return shared_lib.File_exists(dir + "/" + item + ".pm")
		})
		if len(files) == 0 {
			dirs := shared_lib.FindAllDir(".")
			files = shared_lib.Grep(shared_lib.Map(dirs, func(dir string) string {
				return dir + "/" + item + ".pm"
			}), func(path string) bool {
				return shared_lib.File_exists(path)
			})
		}
		if len(files) == 0 {
			log.Fatal("$item is invalid package name.\n")
		}
	}
	return found, line, files
}
