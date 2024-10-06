package func_

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
)

func FindFreePort() int {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port
}

func RunExe() {
	cmd := exec.Command("./ACWA.exe")
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting ACWA.exe:", err)
		os.Exit(1)
	}
	cmd.Wait()
}

func WriteStartArticle(port int) {
	content := fmt.Sprintf(`name = Hot keys
window_h = 800
window_w = 1000
html = <style>iframe{position: fixed;height: 100%%;width: 100%%;top: 0%%;left: 0%%;}</style><iframe src="http://127.0.0.1:%s" frameborder="0"></iframe>
`, strconv.Itoa(port))

	err := os.WriteFile("start.article", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing start.article:", err)
		os.Exit(1)
	}
}
