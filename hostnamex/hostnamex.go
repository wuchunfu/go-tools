package hostnamex

import "strings"

func ExtractHostnamePort(host string) (hostname, port string) {
	if len(host) == 0 {
		return
	}
	hostname = host
	// not [IPv6]
	if hostname[len(hostname)-1] != ']' {
		// [IPv6]:port
		if colonIndex := strings.LastIndex(hostname, "]:"); colonIndex > 0 {
			hostname = hostname[:colonIndex+1]
		} else if colonIndex = strings.LastIndexByte(hostname, ':'); colonIndex >= 0 {
			hostname = hostname[:colonIndex]
		}
	}
	port = host[len(hostname):]
	return
}

func ExtractListenPort(listen string) string {
	if len(listen) == 0 {
		return ""
	}
	// [IPv6]:port
	if colonIndex := strings.LastIndex(listen, "]:"); colonIndex > 0 {
		return listen[colonIndex+2:]
		// [IPv6]¬
	} else if listen[len(listen)-1] == ']' {
		return ""
	} else if colonIndex = strings.LastIndexByte(listen, ':'); colonIndex >= 0 {
		return listen[colonIndex+1:]
	} else {
		lenListen := len(listen)
		if (lenListen < 5 && IsDigits(listen)) ||
			(lenListen == 5 && listen < "65536") {
			return listen
		}
	}
	return ""
}

func IsDigits(input string) bool {
	for i, length := 0, len(input); i < length; i++ {
		b := input[i]
		if b < '0' || b > '9' {
			return false
		}
	}
	return true
}
