package modules

import "fmt"

type Executor interface {
	fmt.Stringer
	Execute([]string) error
}

var SupportedModules = map[string]func() Executor{
	"backend": NewBackendModule,
}

func ModuleFactory(module string) (Executor, error) {
	if moduleInitFunc, ok := SupportedModules[module]; !ok {
		return nil, fmt.Errorf("module not supported")
	} else {
		return moduleInitFunc(), nil
	}
}
