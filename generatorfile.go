package fproto_gen

import "io"

type GeneratorFile interface {
	Generator() Generator
	HaveData() bool
	Output(w io.Writer) error
	Filename() string
	P(str ...interface{})
	In()
	Out()
}
