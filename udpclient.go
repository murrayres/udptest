package main
import (
    "fmt"
    "net"
    "bufio"
    "os"
    "strconv"
    "time"
)

func main() {

    p :=  make([]byte, 2048)
    conn, err := net.Dial("udp", os.Getenv("SERVER")+":9000")
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }


    var counter int
    counter = 1
for {
    counterstring := strconv.Itoa(counter)
    fmt.Println("sending: "+counterstring)
    fmt.Fprintf(conn, counterstring)
    _, err = bufio.NewReader(conn).Read(p)
    if err == nil {
        fmt.Printf("%s\n", p)
    } else {
        fmt.Printf("Some error %v\n", err)
    }
    counter = counter +1
    time.Sleep(time.Second*1)
}

    conn.Close()
}

