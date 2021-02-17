package REPL

import (
	"fmt"
	"skillz_cli/modules"
	"strings"
)

type Shell struct {
	*prompt
	module modules.Executor
}

func NewShell() *Shell {
	return &Shell{
		prompt: NewPrompt(""),
		module: nil,
	}
}

func (sh *Shell) Run() {
	defer sh.Close()

	for {
		sh.Display()
		input := strings.TrimSuffix(sh.GetInput(), "\n")
		cmd := strings.Fields(input)

		switch {
		case len(cmd) == 0 || cmd[0] == "exit":
			return
		case cmd[0] == "continue":
		case cmd[0] == "load":
			sh.LoadModule(cmd[1:])
		default:
			sh.execModuleCmd(cmd)
		}
	}
}

func (sh *Shell) execModuleCmd(cmd []string) {
	if sh.module == nil {
		fmt.Println("No module Loaded")
		return
	}
	if err := sh.module.Execute(cmd); err != nil {
		fmt.Println(err)
	}
	//sh.cmdsFunctions[cmd[0]](cmd[1:])
}

func (sh *Shell) LoadModule(cmd []string) {
	if len(cmd) != 1 {
		fmt.Println("bad load command")
	}
	newModule, err := modules.ModuleFactory(cmd[0])
	if err != nil {
		fmt.Println(err)
	}
	sh.module = newModule
	sh.ChangePromptModule(newModule.String())
	fmt.Printf("module '%v' loaded\n", sh.module)
}
