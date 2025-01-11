package file

import "time"

type Path string

func ParsePath(s string) (Path, error) {
	return Path(s), nil
}

func (p Path) String() string {
	return string(p)
}

type Metadata struct {
	Name    string
	Size    int64
	IsDir   bool
	ModTime time.Time
}
