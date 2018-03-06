package main
import (
    "fmt" 
    "net"  
    "strconv"
//    "strings"
)


func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
    _,err := conn.WriteToUDP([]byte("From server: Hello I got your mesage "), addr)
    if err != nil {
        fmt.Printf("Couldn't send response %v", err)
    }
}


func main() {
    var counter int
    counter=1
    p := make([]byte, 2048)
    addr := net.UDPAddr{
        Port: 9000,
        IP: net.ParseIP("0.0.0.0"),
    }
    ser, err := net.ListenUDP("udp", &addr)
    if err != nil {
        fmt.Printf("Some error %v\n", err)
        return
    }
    for {
        _,remoteaddr,err := ser.ReadFromUDP(p)
        if err !=  nil {
            fmt.Printf("Some error  %v", err)
            continue
        }
        fmt.Println("received")
        received,err := strconv.Atoi(string(p[0]))
if err != nil {
        fmt.Println(err)
}
        fmt.Println(received)
        expecting,err:=strconv.Atoi(strconv.Itoa(counter))
if err !=nil {
        fmt.Println(err)
}
        fmt.Println(expecting)


        if !(received == expecting){
           fmt.Println("lost")
        }
        counter = counter +1
        go sendResponse(ser, remoteaddr)
    }
}
