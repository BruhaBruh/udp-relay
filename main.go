package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	localPort := os.Getenv("LOCAL_PORT")
	remoteHost := os.Getenv("REMOTE_HOST")
	remotePort := os.Getenv("REMOTE_PORT")

	localAddr, err := net.ResolveUDPAddr("udp", ":"+localPort)
	if err != nil {
		log.Fatal("Ошибка ResolveUDPAddr LOCAL:", err)
	}

	remoteAddr, err := net.ResolveUDPAddr("udp", remoteHost+":"+localPort)
	if err != nil {
		log.Fatal("Ошибка ResolveUDPAddr REMOTE:", err)
	}

	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		log.Fatal("Ошибка ListenUDP:", err)
	}

	defer conn.Close()

	log.Printf("Relay is running: localhost:%s -> %s:%s", localPort, remoteHost, remotePort)

	for {
		buf := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Read error:", err)
			continue
		}
		_, err = conn.WriteToUDP(buf[:n], remoteAddr)
		if err != nil {
			log.Println("Send error:", err)
		}
	}
}
