package packr

import "github.com/capossele/packr/v2/file"

func qfile(name string, body string) File {
	f, err := file.NewFile(name, []byte(body))
	if err != nil {
		panic(err)
	}
	return f
}
