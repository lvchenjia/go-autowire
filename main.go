package main

func main() {
	if config_debug {
		printlnYellow("Debug Mode")
		test()
	}

	printlnPink("Release")
}