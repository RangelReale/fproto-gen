package fproto_gen

import "github.com/RangelReale/fdep"

type CreateGeneratorFunc func(dep *fdep.Dep, depfile *fdep.DepFile) (Generator, error)

type FDepGen struct {
	Dep                 *fdep.Dep
	CreateGeneratorFunc CreateGeneratorFunc
}

func NewFDepGen(dep *fdep.Dep, createGeneratorFunc CreateGeneratorFunc) *FDepGen {
	return &FDepGen{
		Dep:                 dep,
		CreateGeneratorFunc: createGeneratorFunc,
	}
}

func (g *FDepGen) Generate(output FileOutput) error {
	output.Initialize()
	defer output.Finalize()

	for _, df := range g.Dep.Files {
		if df.DepType == fdep.DepType_Own {
			err := g.GenerateFile(df, output)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Output must have been called initialize before calling this
func (g *FDepGen) GenerateFile(depfile *fdep.DepFile, output FileOutput) error {
	output.Initialize()
	defer output.Finalize()

	gn, err := g.CreateGeneratorFunc(g.Dep, depfile)
	if err != nil {
		return err
	}

	err = gn.Generate()
	if err != nil {
		return err
	}

	// write all files
	for _, gf := range gn.Files() {
		if gf != nil && gf.HaveData() {
			err = output.Output(gf)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
