package utils

import (
	"strings"
	"sync"
)

var l sync.Mutex

func OfficeToPdf(in, out string) (err error) {
	l.Lock()
	defer l.Unlock()

	cmd := "unoconv"
	arg := []string{"-vvv", "-f", "pdf", "-o", out, in}
	err = ExecCmd(cmd, arg...)
	return
}

func PdfToJpg(in, out string) (err error) {
	l.Lock()
	defer l.Unlock()

	cmd := "convert"
	arg := []string{"-density", "130", "-alpha", "remove", in, out}
	if err = ExecCmd(cmd, arg...); err != nil {
		return
	}
	if FileExist(out) {
		return
	}

	// convert -append output-*.png out.png
	in = strings.Replace(out, ".", "-*.", 1)
	arg = []string{"-append", in, out}
	err = ExecCmd(cmd, arg...)
	return
}
