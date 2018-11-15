package vim

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Edit(prompt string) (string, error) {
	tmpfile, err := ioutil.TempFile("", ".vimtemp")
	if err != nil {
		return "", err
	}

	path := tmpfile.Name()

	if err := writePrompt(path, prompt); err != nil {
		return "", nil
	}

	cmd := exec.Command("vim", path)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return "", err
	}

	content, err := readFile(path)
	if err != nil {
		return "", err
	}

	if err := os.Remove(path); err != nil {
		return "", err
	}

	return content, nil
}

func writePrompt(path, prompt string) error {
	if len(prompt) == 0 {
		return nil
	}
	var fileMode os.FileMode
	return ioutil.WriteFile(path, []byte(prompt), fileMode)
}

func readFile(path string) (string, error) {
	out, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}
	content := string(out)
	return content[:strings.LastIndex(content, "\n")], nil
}

func Prompt(message string) string {
	content, err := Edit(message)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
