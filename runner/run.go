package runner

import (
	"fmt"
	"sort"
)

func Run() (text string, err error) {
	domains, err := LoadFile("")
	if err != nil {
		return "", err
	}

	addrMap := ParseAddr(domains)
	var keys []string
	for k := range addrMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if v, ok := addrMap[k]; ok {
			text = fmt.Sprintf("%s%s", text, fmt.Sprintf("%s %s\n", v, k))
		}
	}

	cmd, err := FlushDNS()
	if err != nil {
		return
	}

	text = fmt.Sprintf("%s\n%s", text, cmd)

	return
}
