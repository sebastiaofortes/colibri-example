package main

import "github.com/colibri-project-io/colibri-sdk-go/pkg/di"

func start() int32{
	return 0
}

func main() {
	// lista de dependências (construtores) 
	dependencies := []interface{}{}
	// instancia o container de DI.
	container := di.NewContainer()
	// adiciona dependências ao container
	container.AddGlobalDependencies(dependencies)
	// iniciando aplicação :)
	container.StartApp(start)
}