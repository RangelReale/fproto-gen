package fproto_gen

type FileOutput interface {
	Initialize() error
	Finalize() error
	Output(gf GeneratorFile) error
}
