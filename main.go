package main

import "socks5"


func main(){
// Create a SOCKS5 server
conf := &socks5.Config{}
server, err := socks5.New(conf)
if err != nil {
  panic(err)
}
// Create SOCKS5 proxy on localhost port 1080
if err := server.ListenAndServe("tcp", "0.0.0.0:1080"); err != nil {
  panic(err)
}
}
