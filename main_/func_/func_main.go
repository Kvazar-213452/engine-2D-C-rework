package func_

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
)

// FindFreePort finds a free port on localhost.
func FindFreePort() int {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		fmt.Println("Error finding free port:", err)
		return 0
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port
}

// RunExe runs the ACWA.exe executable from the core directory.
func RunExe() {
	// Specify the absolute path to ACWA.exe
	corePath := "C:\\Users\\god19\\Desktop\\let\\core\\ACWA.exe"

	cmd := exec.Command(corePath)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting ACWA.exe:", err)
		os.Exit(1)
	}
	cmd.Wait()
}

// WriteStartArticle writes the start.article file with the given port.
func WriteStartArticle(port int) {
	content := fmt.Sprintf(`name = Hot keys
window_h = 800
window_w = 1000
html = <style>iframe{position: fixed;height: 100%%;width: 100%%;top: 0%%;left: 0%%;}</style><iframe src="http://127.0.0.1:%s" frameborder="0"></iframe>
`, strconv.Itoa(port))

	err := os.WriteFile("core/start.article", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing start.article:", err)
		os.Exit(1)
	}
}
