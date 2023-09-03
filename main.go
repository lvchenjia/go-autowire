package main

func main() {
	getConfig()

	if config_debug {
		printlnYellow("Debug Mode")
		test()
	}

	printlnPink("Release")
}
