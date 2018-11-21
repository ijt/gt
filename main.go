package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	// Run go test.
	args := []string{"go", "test"}
	args = append(args, os.Args[1:]...)
	gotest := exec.Command(args[0], args[1:]...)
	var buf bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &buf)
	gotest.Stdout = mw
	gotest.Stderr = mw
	_ = gotest.Run()

	// Run vim on each of the errors.
	// TODO(ijt): Show each error at a convenient time.
	rx := regexp.MustCompile(`([\w/~]+\.go):(\d+)`)
	s := bufio.NewScanner(bytes.NewReader(buf.Bytes()))
	for s.Scan() {
		line := s.Text()
		m := rx.FindStringSubmatch(line)
		if len(m) == 3 {
			cmd := exec.Command("vim", m[1], "+"+m[2])
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			if err := cmd.Run(); err != nil {
				fmt.Printf("v: %v\n", err)
			}
		}
	}
}
