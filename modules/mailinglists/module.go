package mailinglists

import (
	"net/http"
	"ocluso/pkg/moduleinterface"
)

type Module struct{}

func BuildModule(context moduleinterface.ModuleContext) (moduleinterface.Module, error) {
	return &Module{}, nil
}

func (m *Module) Name() string {
	return "mailinglists"
}

func (m *Module) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World from mailinglists!"))
}
