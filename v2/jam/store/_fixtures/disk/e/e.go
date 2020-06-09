package q

import "github.com/capossele/packr/v2"

func init() {
	packr.New("tom", "./petty")
	packr.NewBox("../e/heartbreakers")
}
