package main

import "regexp"

func _remove_git_path_prefix(item *string) {
	*item = regexp.MustCompile(`^[ab]/`).ReplaceAllString(*item, "")
}
