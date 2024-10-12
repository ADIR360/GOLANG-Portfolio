package main

import (
    "fmt"
    "golang.org/x/crypto/ssh"
    "io"
    "log"
    "net"
    "os/exec"
)

func main() {
    config := &ssh.ServerConfig{
        NoClientAuth: true,
    }

    privateBytes := []byte(`YOUR_PRIVATE_KEY_HERE`) // Load your private key here
    private, err := ssh.ParsePrivateKey(privateBytes)
    if err != nil {
        log.Fatalf("unable to parse private key: %v", err)
    }
    config.AddHostKey(private)

    listener, err := net.Listen("tcp", "0.0.0.0:2222")
    if err != nil {
        log.Fatalf("failed to listen for connection: %v", err)
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatalf("failed to accept connection: %v", err)
        }

        go handleConn(conn, config)
    }
}

func handleConn(conn net.Conn, config *ssh.ServerConfig) {
    sshConn, _, _, err := ssh.NewServerConn(conn, config)
    if err != nil {
        log.Printf("failed to establish SSH connection: %v", err)
        conn.Close()
        return
    }
    log.Printf("logged in: %s", sshConn.User())

    // Run your application
    cmd := exec.Command("go", "run", "main.go") // Change this to your actual command
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Printf("failed to create stdout pipe: %v", err)
        return
    }
    stderr, err := cmd.StderrPipe()
    if err != nil {
        log.Printf("failed to create stderr pipe: %v", err)
        return
    }

    if err := cmd.Start(); err != nil {
        log.Printf("failed to start command: %v", err)
        return
    }

    go io.Copy(sshConn.Stdout(), stdout)
    go io.Copy(sshConn.Stderr(), stderr)

    if err := cmd.Wait(); err != nil {
        log.Printf("command finished with error: %v", err)
    }
}
