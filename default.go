package fproto_gen

import (
	"os"
	"path/filepath"
)

//
// FileOutput: default
//

type FileOutput_Default struct {
	OutputPath string
}

func NewFileOutput_Default(outputPath string) *FileOutput_Default {
	return &FileOutput_Default{
		OutputPath: outputPath,
	}
}

func (f *FileOutput_Default) Initialize() error {
	return nil
}

func (f *FileOutput_Default) Finalize() error {
	return nil
}

func (f *FileOutput_Default) Output(g GeneratorFile) error {
	p := filepath.Join(f.OutputPath, g.Filename())

	// create paths
	err := os.MkdirAll(filepath.Dir(p), os.ModePerm)
	if err != nil {
		return err
	}

	// create file
	file, err := os.Create(p)
	if err != nil {
		return err
	}
	defer file.Close()

	// output contents
	err = g.Output(file)
	if err != nil {
		return err
	}

	return nil
}
