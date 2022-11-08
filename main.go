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
	"strconv"
	"syscall"
	"unsafe"
)

const VERSION = "1.1"

func main() {
	isRegister := flag.Bool("register", false, "register url scheme")
	flag.Parse()

	if *isRegister {
		registerURLScheme()
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("parameter missing.")
		if MessageBox(0, "Need help?\r\n你是否需要一点帮助？如果需要，请点“是”\r\n\r\nBy.Dark495\r\nhttps://github.com/xlch88/putty-url-scheme", "PuTTY URL Scheme | Version "+VERSION, 32+4) == 6 {
			exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://github.com/xlch88/putty-url-scheme").Start()
		}

		return
	}

	u, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	username := ""
	password := ""
	hostname := ""
	port := ""

	if u.User.Username() != "" {
		username = u.User.Username()
	} else {
		username = "root"
	}

	p, has := u.User.Password()
	if has {
		password = p
	}

	_host, _port, err := net.SplitHostPort(u.Host)
	if err == nil {
		port = _port
		hostname = _host
	} else {
		port = "22"
		hostname = _host
	}

	cmd := []string{"-C"}

	tmpSessionsPath := "Software\\SimonTatham\\PuTTY\\Sessions\\PuTTYURLSchemeProxy"
	isRegistry := false
	if u.Query().Get("proxyHost") != "" {
		isRegistry = true
		fmt.Println("Use proxy.")

		permission := uint32(registry.QUERY_VALUE | registry.SET_VALUE)

		registry.DeleteKey(registry.CURRENT_USER, tmpSessionsPath)
		key, _, _ := registry.CreateKey(registry.CURRENT_USER, tmpSessionsPath, permission)

		defaultK, defaultKErr := registry.OpenKey(registry.CURRENT_USER, "Software\\SimonTatham\\PuTTY\\Sessions\\Default%20Settings", registry.READ)

		if defaultKErr == nil {
			value1, _, err := defaultK.GetStringValue("Font")
			if err == nil {
				key.SetStringValue("Font", value1)
			}

			value2, _, err := defaultK.GetIntegerValue("FontCharSet")
			if err == nil {
				key.SetDWordValue("FontCharSet", uint32(value2))
			}

			value3, _, err := defaultK.GetIntegerValue("FontHeight")
			if err == nil {
				key.SetDWordValue("FontHeight", uint32(value3))
			}

			value4, _, err := defaultK.GetIntegerValue("FontQuality")
			if err == nil {
				key.SetDWordValue("FontQuality", uint32(value4))
			}

			value5, _, err := defaultK.GetIntegerValue("FontVTMode")
			if err == nil {
				key.SetDWordValue("FontVTMode", uint32(value5))
			}

			value6, _, err := defaultK.GetIntegerValue("FontIsBold")
			if err == nil {
				key.SetDWordValue("FontIsBold", uint32(value6))
			}
		}

		var proxyPort = 1080
		if u.Query().Get("proxyPort") != "" {
			i, err := strconv.Atoi(u.Query().Get("proxyPort"))
			if err == nil {
				proxyPort = i
			}
		}

		var proxyMethod = 2
		if u.Query().Get("proxyMethod") != "" {
			i, err := strconv.Atoi(u.Query().Get("proxyMethod"))
			if err == nil {
				proxyMethod = i
			}
		}

		key.SetStringValue("ProxyHost", u.Query().Get("proxyHost"))
		key.SetStringValue("ProxyUsername", u.Query().Get("proxyUsername"))
		key.SetStringValue("ProxyPassword", u.Query().Get("proxyPassword"))
		key.SetDWordValue("ProxyPort", uint32(proxyPort))
		key.SetDWordValue("ProxyMethod", uint32(proxyMethod))

		key.SetStringValue("HostName", hostname)
		key.SetStringValue("PortNumber", port)
		key.SetStringValue("UserName", username)
		key.SetStringValue("Protocol", u.Scheme)

		cmd = append(cmd, "-l")
		cmd = append(cmd, username)

		if password != "" {
			cmd = append(cmd, "-pw")
			cmd = append(cmd, password)
		}

		cmd = append(cmd, "-load")
		cmd = append(cmd, "PuTTYURLSchemeProxy")
	}

	cmd = append(cmd, "-l")
	cmd = append(cmd, username)

	if password != "" {
		cmd = append(cmd, "-pw")
		cmd = append(cmd, password)
	}

	cmd = append(cmd, "-P")
	cmd = append(cmd, port)
	cmd = append(cmd, hostname)

	ex, _ := os.Executable()
	puttyBin := filepath.Dir(ex) + "\\putty.exe"
	execCmd := exec.Command(puttyBin, cmd...)

	err = execCmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	if isRegistry {
		execCmd.Wait()
		registry.DeleteKey(registry.CURRENT_USER, tmpSessionsPath)
	}
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

	// create tree
	registry.CreateKey(baseKey, basePath+"\\shell", permission)
	registry.CreateKey(baseKey, basePath+"\\shell\\open", permission)
	k, _, _ = registry.CreateKey(baseKey, basePath+"\\shell\\open\\command", permission)

	// set open command
	k.SetStringValue("", programLocation+" \"%1\"")

	fmt.Println("PuTTY URL Scheme | Version " + VERSION)
	MessageBox(0, "register success.\r\n\r\nBy.Dark495\r\nhttps://github.com/xlch88/putty-url-scheme", "PuTTY URL Scheme | Version "+VERSION, 0)
}

// MessageBox of Win32 API.
func MessageBox(hwnd uintptr, caption, title string, flags uint) int {
	ret, _, _ := syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(flags))

	return int(ret)
}
