package main

import (
    "log"
    "time"
    "regexp"
    "bufio"
    "os"
    "fmt"
    "strings"
    "github.com/tarm/serial"
    // "github.com/gorilla/websocket"
)

func main() {
    r, _ := regexp.Compile("[0-9]{1,3}.[0-9]{1,4}")

    fmt.Print("Enter serialport name: ")
    reader := bufio.NewReader(os.Stdin)
    port, _ := reader.ReadString('\n')
    port = strings.TrimSuffix(port, "\n")
    fmt.Println("Connecting to " + string(port))
    serialport_config := &serial.Config{Name: port, Baud: 9600}
    s, err := serial.OpenPort(serialport_config)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to " + string(port))

    buf := make([]byte, 128)
    t1 := time.NewTicker(time.Millisecond * 500)
    for t := range t1.C {
        fmt.Println("Ticker start @ ", t)
        n, err := s.Read(buf)
        if err != nil {
            log.Fatal(err)
        }
        str := string(buf)
        log.Printf("%q", buf[:n])
        if r.MatchString(str) {
            matched_string := r.FindString(str)
            fmt.Println(matched_string)
            // websocket_connection, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8899/serialport", nil)
            // if err != nil {
            //     log.Fatal("dial:", err)
            // }
            // defer websocket_connection.Close()
            // err = websocket_connection.WriteMessage(websocket.TextMessage, []byte(matched_string))
            // if err != nil {
            //     log.Println(err)
            //     return
            // }
        }
    }
}
