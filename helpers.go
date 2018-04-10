package fproto_gen

import (
	"bytes"
	"fmt"
	"io"
)

type GeneratorFileHelper struct {
	buf      *bytes.Buffer
	indent   string
	havedata bool
}

func NewGeneratorFileHelper() *GeneratorFileHelper {
	return &GeneratorFileHelper{
		buf: new(bytes.Buffer),
	}
}

func (gf *GeneratorFileHelper) HaveData() bool {
	return gf.havedata
}

func (gf *GeneratorFileHelper) P(str ...interface{}) {
	gf.havedata = true

	gf.buf.WriteString(gf.indent)
	for _, v := range str {
		switch s := v.(type) {
		case string:
			gf.buf.WriteString(s)
		case *string:
			gf.buf.WriteString(*s)
		case bool:
			fmt.Fprintf(gf.buf, "%t", s)
		case *bool:
			fmt.Fprintf(gf.buf, "%t", *s)
		case int:
			fmt.Fprintf(gf.buf, "%d", s)
		case *int32:
			fmt.Fprintf(gf.buf, "%d", *s)
		case *int64:
			fmt.Fprintf(gf.buf, "%d", *s)
		case float64:
			fmt.Fprintf(gf.buf, "%g", s)
		case *float64:
			fmt.Fprintf(gf.buf, "%g", *s)
		case fmt.Stringer:
			fmt.Fprintf(gf.buf, "%s", s.String())
		default:
			panic(fmt.Sprintf("unknown type in printer: %T", v))
		}
	}
	gf.buf.WriteByte('\n')
}

func (gf *GeneratorFileHelper) In() {
	gf.indent += "\t"
}

func (gf *GeneratorFileHelper) Out() {
	if len(gf.indent) > 0 {
		gf.indent = gf.indent[1:]
	}
}

func (gf *GeneratorFileHelper) WriteTo(w io.Writer) (int64, error) {
	return gf.buf.WriteTo(w)
}
