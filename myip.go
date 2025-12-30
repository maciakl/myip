package main

import (
    "os"
    "io"
    "net"
    "fmt"
    "net/http"
    "crypto/tls"
    "path/filepath"
)

const version = "0.1.1"

func main() {

    if len(os.Args) > 1 {

        switch os.Args[1] {

            case "-p", "--public":
                ip := getPublicIP()
                fmt.Println(ip)
                os.Exit(0)

            case "-v", "--version":
                Version()

            case "-h", "--help":
                fmt.Println("This command prints the default local outbound IP address of the current machine.")
                Usage()

            default:
                fmt.Fprintln(os.Stderr, "Invalid option:", os.Args[1])
                Usage()
        } 
    } else { // no arguments

        ip := getOutboundIP()
        fmt.Println(ip)
        os.Exit(0)
    }
    
}

func Version() {
    fmt.Println(filepath.Base(os.Args[0]), "version", version)
    os.Exit(0)
}

func Usage() {
    fmt.Println("Usage:", filepath.Base(os.Args[0]), "[options]")
    fmt.Println("Options:")
    fmt.Println("  -p, --public     Print the public IP address instead (requires internet access)")
    fmt.Println("  -v, --version    Print version information and exit")
    fmt.Println("  -h, --help       Print this message and exit")
    os.Exit(0)
}


func getOutboundIP() string {

    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr) // type assertion

    return localAddr.IP.String()
}

func getPublicIP() string {

    // don't verify the server's certificate. probably not a good idea but we want to be able to do this
    // in an environment with TLS inspection enabled
    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

    resp, err := http.Get("https://api.ipify.org")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    ip, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    return string(ip)
}
