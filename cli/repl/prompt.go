package REPL

import (
	"bufio"
	"fmt"
	"github.com/logrusorgru/aurora"
	"io"
	"os"
	"strings"
)

type prompt struct {
	comm       mux
	moduleName string
	reader     *bufio.Reader
}

func checkErr(err error) string {
	if err == io.EOF {
		return "exit"
	} else {
		return ""
	}
}

func (p *prompt) readInput() {
	for {
		if in, err := p.reader.ReadString('\n'); err != nil {
			p.comm.msg <- checkErr(err)
		} else {
			p.comm.msg <- in
		}
	}
}

func NewPrompt(module string) *prompt {
	p := &prompt{moduleName: module}
	p.reader = bufio.NewReader(os.Stdin)

	go p.readInput()
	p.comm.init()
	return p
}

func getCurrentDir() string {
	pwd := os.Getenv("PWD")
	return pwd[strings.LastIndex(pwd, "/")+1:]
}

func (p *prompt) Display() {
	sep := aurora.Green("→")
	brack1 := aurora.Red("[")
	brack2 := aurora.Red("]")
	cwd := aurora.BrightBlue(getCurrentDir())
	if p.moduleName == "" {
		fmt.Printf("%s %s ", cwd, sep)
		return
	}
	module := aurora.BrightBlue(p.moduleName)
	fmt.Printf("%s %s %s %s %s ", cwd, brack1, module, brack2, sep)
}

func (p *prompt) ChangePromptModule(m string) {
	p.moduleName = m
}

func (p *prompt) GetInput() string {
	select {
	case in := <-p.comm.msg:
		return in
	case <-p.comm.interrupt:
		fmt.Println("")
		return ""
	}
}

func (p *prompt) Close() {
	p.comm.Close()
	fmt.Println("Stopping Shell.")
}
