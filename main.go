package main

import (
    "fmt"
    "net"
    "os"
)

func main(){

  service := ":2000"
  tcpAddr, err := net.ResolveTCPAddr("tcp", service)
  ckerr(err)

  listener, err := net.ListenTCP("tcp", tcpAddr)
  ckerr(err)

  for {
     conn, err := listener.Accept()
     if err != nil {
        continue
     }

     // run as a go routine
     go handleClient(conn)
  
  } // end forever
}

func handleClient(conn net.Conn) {
    // close connection on exit
    defer conn.Close()
    
    var buf [512]byte
    for {
        // read up to 512 bytes
        n, err := conn.Read(buf[0:])
        if err != nil {
             return
        }
       fmt.Println( string(buf[0:]) )
        // write the n bytes read
       _, err2 := conn.Write(buf[0:n])
       if err2 != nil {
         return
       }
    } // end forever
}

func ckerr(err error) {
    if err != nil {
       fmt.Fprintf(os.Strerr, "Fatal error: %s ", err.Error() )
       os.Exit(1)
    }
}


