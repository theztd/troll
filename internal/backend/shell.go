package backend

type Shell struct {
	Root string
}

func NewShell(dsn string) (*Shell, error) {
	return &Shell{Root: dsn}, nil
}

func (red *Shell) RunAndRenderTpl(query string, tmpl string) ([]byte, error) {
	// cmd := exec.Command(query)

	// var outBuf, errBuf bytes.Buffer
	// err := cmd.Run()
	return []byte{}, nil
}

func (r *Shell) Close() error { return nil }
