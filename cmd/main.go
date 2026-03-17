package main

import (
    "github.com/go-echo-server/pkg/echo"
    "log"
)

func main() {
    log.Println("Starting echo server...")
    echo.StartServer()
}