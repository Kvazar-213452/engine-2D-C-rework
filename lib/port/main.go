package main

import (
	"C"
	"net"
)

//export FindFreePort
func FindFreePort() int {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0 // Можна повернути 0 або інше значення, якщо не вдалося
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port
}

// main просто для компіляції, тут нічого не потрібно
func main() {}
