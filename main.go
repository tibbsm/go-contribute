package main

import "flag"

// scan a given path and its subfolders to search for Git repos
func scan(path string) {
	print("scan")
}

// generates graph of contributions
func stats(email string) {
	print("stats")
}

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories.")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}
	stats(email)
}
