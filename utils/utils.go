package utils

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func ExecCmd(name string, arg ...string) (err error) {
	log.Printf("[INFO] Executing: %v %v \n", name, strings.Join(arg, " "))
	cmd := exec.Command(name, arg...)
	var bf bytes.Buffer
	cmd.Stdout = &bf
	err = cmd.Run()
	if err != nil {
		log.Println("[ERROR] ", err.Error())
		return
	}
	log.Printf("[INFO] %v\n", bf.String())

	return
}
