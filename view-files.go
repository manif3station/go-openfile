package main

import "regexp"

func _tt_include(item string) []string {
	return regexp.MustCompile(`(?:template|PROCESS|INCLUDE)[\s\(]['"]?([\w\-\/]+)`).FindAllString(item, -1)
}

func _view_files(item string) []string {
	return regexp.MustCompile(`([\w\-\/]+\.(?:js|tt|scss))`).FindAllString(item, -1)
}

func _css_files(item string) []string {
	return regexp.MustCompile(`([\w\-\/]+\.)css`).FindAllString(item, -1)
}
