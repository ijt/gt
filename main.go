package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	// Run go test
	gotest := exec.Command("go", "test")
	out, _ := gotest.CombinedOutput()
	fmt.Printf("%s\n", out)

	rx := regexp.MustCompile(`(\w+\.go):(\d+)`)
	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		line := s.Text()
		m := rx.FindStringSubmatch(line)
		if len(m) == 3 {
			cmd := exec.Command("vim", m[1], "+"+m[2])
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			if err := cmd.Run(); err != nil {
				fmt.Printf("v: %v: %s\n", err, out)
			}
		}
	}
}
