package net

import (
	"fmt"
	"log"
	"net"
	"time"
)

var effectiveConns chan net.Conn
var retryConns chan string

var globalTimeOut time.Duration

//init connect pool
func InitConnectPool(servList []string, timeOut time.Duration, minConnNum int) {
	globalTimeOut = timeOut
	//init channel
	effectiveConns = make(chan net.Conn, len(servList)*minConnNum)
	retryConns = make(chan string, len(servList)*minConnNum)
	//run retry
	go Retry()
	//range server list
	for _, addr := range servList {
		//range connect number
		connCount := 0
		for connCount < minConnNum {
			connCount = connCount + 1
			connectServer(addr, globalTimeOut)
		}
	}
}

//create connect
func connectServer(addr string, globalTimeOut time.Duration) {
	conn, err := net.DialTimeout("tcp", addr, globalTimeOut*time.Second)
	if err != nil {
		log.Println("error connecting", err)
		//write connect to retry channel
		retryConns <- addr
	} else {
		log.Println("connect to", conn.RemoteAddr())
		//write connect to effective channel
		effectiveConns <- conn
	}
}

//get one connect from channel
func Get() net.Conn {
CREATECONN:
	conn := <-effectiveConns
	if conn == nil {
		//retry for multi read
		goto CREATECONN
	}
	return conn
}

//write connect to effective channel
func Put(conn net.Conn) {
	effectiveConns <- conn
}

//write connect to retry channel
func Drop(conn net.Conn) {
	retryConns <- fmt.Sprintf("%v", conn.RemoteAddr())
}

//Retry tcp connect
func Retry() {
	for addr := range retryConns {
		log.Println(addr)
		connectServer(addr, globalTimeOut)
		time.Sleep(5 * time.Second)
	}
}

//close tcp connect
func Close() {
	for conn := range effectiveConns {
		conn.Close()
	}
	close(effectiveConns)
	close(retryConns)
}
