package main

import (
	REPL "skillz_cli/repl"
)

func main() {
	sh := REPL.NewShell()
	sh.LoadModule([]string{"backend"})
	sh.Run()
}
