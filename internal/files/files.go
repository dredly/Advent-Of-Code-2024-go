package files

import "os"

func MustRead(path string) string {
	bytes, err := os.ReadFile(path)
	if err !=nil {
		panic("error reading file")
	}
	return string(bytes)
}