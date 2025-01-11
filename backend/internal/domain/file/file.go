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
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	IsDir   bool      `json:"is_dir"`
	ModTime time.Time `json:"mod_time"`
}
