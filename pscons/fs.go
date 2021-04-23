package pscons

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func exists(fname string) bool {
	if _, err := os.Stat(fname); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func chkFile(fname string) error {
	if !exists(fname) {
		return errors.New("no such file or directory")
	} else if s, _ := os.Stat(fname); s.IsDir() {
		return errors.New("proposed fname is directory")
	}
	return nil
}

func get(fname string) (string, error) {
	if err := chkFile(fname); err != nil {
		return "", err
	}

	f, err := os.Open(fname)
	if err != nil {
		return "", err
	}

	rdr := bufio.NewReader(f)
	pass, err := rdr.ReadString('\n')
	if err != nil {
		return "", err
	}
	return pass[:len(pass)-1], nil
}

func gen(fname string, force bool) (string, error) {
	if err := chkFile(fname); err != nil {
		if os.IsExist(err) {
			return "", err
		}
	}

	pass, perr := genPassword()
	if perr != nil {
		return "", perr
	}
	f, ferr := os.Create(fname)
	if ferr != nil {
		return "", ferr
	}

	wrt := bufio.NewWriter(f)
	n, err := wrt.WriteString(pass + "\n")
	wrt.Flush()
	if err != nil {
		return "", err
	} else if n != len(pass)+1 {
		return "", errors.New("erroneous length written")
	}

	f.Close()
	return pass, nil
}

func find(ss string) ([]string, error) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, err
	}

	rl := []string{}
	for _, f := range files {
		if strings.Contains(f.Name(), ss) {
			rl = append(rl, f.Name())
		}
	}
	return rl, nil
}

func del(fname string) error {
	if err := chkFile(fname); err != nil {
		return err
	}
	err := os.Remove(fname)
	return err
}
