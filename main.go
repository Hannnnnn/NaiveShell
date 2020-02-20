package main

import (

	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"command"
	"reflect"
	"strings"
	"syscall"
)

func parse(s string) interface{} {
	arr := strings.Split(s, " ")
	return command.Command{
		Name: arr[0],
		Args: arr,
	}
}


func process(s string) {

	command := parse(s)
//	arr := strings.Split(s, " ")
	execute(command)
}


func execute(comm interface{}) {
	refVal := reflect.ValueOf(comm)
	commVal := refVal.Convert(refVal.Type())
	
	if refVal.Type() == reflect.TypeOf(command.Command{}) {
		
		path, err := exec.LookPath(commVal.Field(0).Interface().(string))

		if err != nil {
			log.Fatal(err)
		}

		env := os.Environ()

		err = syscall.Exec(path, commVal.Field(1).Interface().([]string), env)

		if err != nil {
			log.Fatal(err)
		}
	}
	

}

func terminal() {
	hostname, _ := os.Hostname()
	username, _ := user.Current()
	for {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

	//	fmt.Printf("\033[1;32m%s\033[0m\n", username.Username)
	//	fmt.Print("\033[1;36mbold red text\033[0m\n")
		fmt.Printf("\033[1;32m%s@%s\033[0m:\033[1;36m%s\033[0m$ ", username.Username, hostname, pwd )
	//	fmt.Print(username.Username + "@" + hostname + ":" + pwd + "$")
		reader := bufio.NewReader(os.Stdin)

		text, _ := reader.ReadString('\n')
		s := strings.Trim(text, "\r\n ")
		if strings.Compare(s, "quit") == 0 {
			os.Exit(0)
		} else if strings.Compare(s, "") == 0 {
			continue
		} else {
			pid, _, _ := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)

			if pid == 0 {
				process(s)
				break

			} else {
				syscall.Wait4(int(pid), nil, 0, nil)

			}
		}
	}

}




func main() {
	terminal()
	
}
