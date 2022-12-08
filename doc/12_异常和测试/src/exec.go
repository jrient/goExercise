// exec.go
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// 1) os.StartProcess //
	/*********************/
	/* Linux: */
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// 1st example: list files
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid)

	/*
		The process id is &{27584 0 0 {{0 0} 0 0 0 0}}total 28
		-rw-r--r-- 1 root root  214 Dec  6 16:31 errors.go
		-rw-r--r-- 1 root root  515 Dec  7 16:07 exec.go
		-rw-r--r-- 1 root root  695 Dec  6 23:51 panic_defer.go
		-rw-r--r-- 1 root root  360 Dec  6 17:37 panic.go
		-rw-r--r-- 1 root root  737 Dec  6 23:12 panic_package.go
		-rw-r--r-- 1 root root  446 Dec  6 22:19 panic_recover.go
		drwxr-xr-x 2 root root 4096 Dec  6 22:28 parse
	*/

	// 2nd example: show all processes
	pid, err = os.StartProcess("/bin/ps", []string{"-e", "-opid,ppid,comm"}, procAttr)

	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}

	fmt.Printf("The process id is %v", pid)

	/*
		The process id is &{29204 0 0 {{0 0} 0 0 0 0}}  PID  PPID COMMAND
		10359 10350 bash
		29163 10359 go <defunct>
		29204     1 ps
	*/

	cmd := exec.Command("gedit") // this opens a gedit-window
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!", err)
		os.Exit(1)
	}
	fmt.Printf("The command is %v", cmd)
}
