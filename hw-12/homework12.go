package main

import (
	"errors"
	"fmt"
	"os/exec"

	//flag "github.com/spf13/pflag"
	//"io"
	"os"
)

func ReadDir(dir string) (map[string]string, error) {
	res := map[string]string{}

	file, err := os.Open(dir)
	if err != nil {
		return res, err
	}

	names, err := file.Readdirnames(0)
	if err != nil {
		return res, err
	}

	for _, arg := range names {
		argFile, err := os.OpenFile(dir+"/"+arg, os.O_RDONLY, 0775)
		if err != nil {
			return res, err
		}

		fileInfo, _ := argFile.Stat()
		argVal := make([]byte, fileInfo.Size())
		_, err = argFile.Read(argVal)
		if err != nil {
			return res, err
		}

		res[arg] = string(argVal)
	}

	return res, nil
}

func RunCmd(cmd []string, env map[string]string) error {

	execCmd := exec.Command(cmd[0])
	for arg, val := range env {
		execCmd.Env = append(execCmd.Env, arg+"="+val)
	}
	out, err := execCmd.Output()

	fmt.Println(string(out))
	return err
}

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(errors.New("you should set env dir and program: f.e. homework12.go env_dir program"))
		os.Exit(1)
	}

	envVariables, err := ReadDir(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = RunCmd(args[2:], envVariables)
	fmt.Println(err)
}
