package main

import (
	"flag"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"net"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var isRegister = flag.Bool("register", false, "register url scheme")
	flag.Parse()

	if *isRegister {
		registerURLScheme()
		return
	}

	u, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var cmd = []string{ "-C"}

	if u.User.Username() != "" {
		cmd = append(cmd, "-l")
		cmd = append(cmd, u.User.Username())
	}else{
		cmd = append(cmd, "-l")
		cmd = append(cmd, "root")
	}

	p, has := u.User.Password()
	if has {
		cmd = append(cmd, "-pw")
		cmd = append(cmd, p)
	}

	host, port, err := net.SplitHostPort(u.Host)
	if err == nil {
		cmd = append(cmd, "-P")
		cmd = append(cmd, port)
		cmd = append(cmd, host)
	}else{
		cmd = append(cmd, u.Host)
	}

	ex, _ := os.Executable()
	execCmd := exec.Command(filepath.Dir(ex) + "\\putty.exe", cmd...)
	err = execCmd.Start()
	fmt.Println(err)
}

func registerURLScheme() {
	var k registry.Key

	prefix := "SOFTWARE\\Classes\\"
	urlScheme := "ssh"
	basePath := prefix + urlScheme
	permission := uint32(registry.QUERY_VALUE | registry.SET_VALUE)
	baseKey := registry.CURRENT_USER

	ex, _ := os.Executable()
	fmt.Println(ex)

	programLocation := ex

	// create key
	k, _, _ = registry.CreateKey(baseKey, basePath, permission)

	// set description
	k.SetStringValue("", "Putty")
	k.SetStringValue("URL Protocol", "")

	// set icon
	k, _, _ = registry.CreateKey(registry.CURRENT_USER, "lumiere\\DefaultIcon", registry.ALL_ACCESS)
	k.SetStringValue("", programLocation+",1")

	// create tree
	registry.CreateKey(baseKey, basePath+"\\shell", permission)
	registry.CreateKey(baseKey, basePath+"\\shell\\open", permission)
	k, _, _ = registry.CreateKey(baseKey, basePath+"\\shell\\open\\command", permission)

	// set open command
	k.SetStringValue("", programLocation+" \"%1\"")

	fmt.Println("register success.")
}
