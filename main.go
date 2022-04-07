package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	// GithubDomainPath GitHub domain file
	GithubDomainPath = "domain/github.txt"

	// Goos
	LINUX   = "linux"
	Darwin  = "darwin"
	WINDOWS = "windows"

	// hosts file path
	WindowsHosts = "C:/Windows/system32/drivers/etc/hosts"
	OtherHosts   = "/etc/hosts"
)

// load static domain.txt file
//go:embed domain/github.txt
var fs embed.FS

func main() {
	domains, err := loadFile()
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("load domain file successfully")

	addrMap, err := resolveIPAddr(domains)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("resolve IP address successfully")

	if err := flushDNS(addrMap); err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("flush DNS successfully")

	return
}

// load GitHub domains from txt file
func loadFile() (domains []string, err error) {
	fr, err := fs.Open(GithubDomainPath)
	if err != nil {
		return nil, fmt.Errorf("open file err：%s", err.Error())
	}

	br := bufio.NewReader(fr)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		domains = append(domains, strings.TrimSpace(string(line)))
	}

	return
}

// resolve ip address of domain
func resolveIPAddr(domains []string) (map[string]string, error) {
	addrMap := make(map[string]string, len(domains))

	for _, domain := range domains {
		addr, err := net.ResolveIPAddr("ip", domain)
		if err != nil {
			return nil, fmt.Errorf("resolve ip addr err：%s", err.Error())
		}
		addrMap[domain] = addr.IP.String()
	}

	return addrMap, nil
}

// flushDNS
func flushDNS(addrMap map[string]string) error {
	var (
		hf   *os.File
		cmd  string
		err  error
		path string
	)

	// choose a goos cmd and hosts file
	switch runtime.GOOS {
	case LINUX:
		path = OtherHosts
		// sudo systemctl restart nscd
		cmd = "sudo systemctl restart nscd"
	case Darwin:
		path = OtherHosts
		// sudo killall -HUP mDNSResponder
		cmd = "sudo killall -HUP mDNSResponder"
	case WINDOWS:
		path = WindowsHosts
		// ipconfig flushdns
		cmd = "ipconfig flushdns"
	default:
		return fmt.Errorf("unsupported OS: %v", runtime.GOOS)
	}

	// open and backup hosts file to host_tmp file
	hf, err = os.OpenFile(OtherHosts, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer hf.Close()

	byteArr, err := io.ReadAll(hf)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s_tmp", path), byteArr, 0666)
	if err != nil {
		return err
	}

	// replace or append github dns info to hosts file
	var delFlag bool
	bufSrc := bytes.NewBuffer(byteArr)
	bufDst := &bytes.Buffer{}

	for {
		line, c := bufSrc.ReadString('\n')
		if c == io.EOF {
			break
		}
		if strings.Contains(line, "# Github Hosts head……") {
			delFlag = true
		}
		if !delFlag {
			bufDst.WriteString(line)
		}
		if strings.Contains(line, "# Github Hosts tail……") {
			delFlag = false
		}
	}

	bufDst.WriteString("# Github Hosts head……\n")
	for domain, ip := range addrMap {
		bufDst.WriteString(fmt.Sprintf("%s %s\n", ip, domain))
	}
	bufDst.WriteString("# Github Hosts tail……\n")
	_, _ = hf.Seek(0, os.SEEK_SET)
	_, _ = hf.Write(bufDst.Bytes())

	command := exec.Command("bash", "-c", cmd)
	if _, err := command.Output(); err != nil {
		return err
	}

	return nil
}
