package fproto_gen

import "github.com/RangelReale/fdep"

type GeneratorSyntax int

const (
	GeneratorSyntax_Proto2 GeneratorSyntax = iota
	GeneratorSyntax_Proto3
)

type Generator interface {
	Syntax() GeneratorSyntax
	GetDep() *fdep.Dep
	GetDepFile() *fdep.DepFile
	Generate() error
	Files() []GeneratorFile
}
