package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micmonay/keybd_event"
	"log"
	"net"
)

type keyStruct struct {
	Key   string `json:"key" binding:"required"`
	Shift bool   `json:"shift" binding:"required"`
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	httpServer := gin.Default()

	httpServer.POST("/key", func(c *gin.Context) {
		var keyBody keyStruct
		err := c.ShouldBindJSON(&keyBody)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err = pressKey(keyBody)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "success"})
	})

	httpServer.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	ip := GetOutboundIP()

	log.Printf("Server started on %s:8080", ip.String())
	err := httpServer.Run(ip.String() + ":8080")
	if err != nil {
		log.Printf("Error starting server: %s", err.Error())
		return
	}

}
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func pressKey(key keyStruct) error {
	kb, err := keybd_event.NewKeyBonding()

	if err != nil {
		return err
	}

	if key.Shift {
		kb.HasSHIFT(true)
	}

	switch key.Key {
	case "a":
		kb.SetKeys(keybd_event.VK_A)
	case "b":
		kb.SetKeys(keybd_event.VK_B)
	case "c":
		kb.SetKeys(keybd_event.VK_C)
	case "d":
		kb.SetKeys(keybd_event.VK_D)
	case "e":
		kb.SetKeys(keybd_event.VK_E)
	case "f":
		kb.SetKeys(keybd_event.VK_F)
	case "g":
		kb.SetKeys(keybd_event.VK_G)
	case "h":
		kb.SetKeys(keybd_event.VK_H)
	case "i":
		kb.SetKeys(keybd_event.VK_I)
	case "j":
		kb.SetKeys(keybd_event.VK_J)
	case "k":
		kb.SetKeys(keybd_event.VK_K)
	case "l":
		kb.SetKeys(keybd_event.VK_L)
	case "m":
		kb.SetKeys(keybd_event.VK_M)
	case "n":
		kb.SetKeys(keybd_event.VK_N)
	case "o":
		kb.SetKeys(keybd_event.VK_O)
	case "p":
		kb.SetKeys(keybd_event.VK_P)
	case "q":
		kb.SetKeys(keybd_event.VK_Q)
	case "r":
		kb.SetKeys(keybd_event.VK_R)
	case "s":
		kb.SetKeys(keybd_event.VK_S)
	case "t":
		kb.SetKeys(keybd_event.VK_T)
	case "u":
		kb.SetKeys(keybd_event.VK_U)
	case "v":
		kb.SetKeys(keybd_event.VK_V)
	case "w":
		kb.SetKeys(keybd_event.VK_W)
	case "x":
		kb.SetKeys(keybd_event.VK_X)
	case "y":
		kb.SetKeys(keybd_event.VK_Y)
	case "z":
		kb.SetKeys(keybd_event.VK_Z)
	case "1":
		kb.SetKeys(keybd_event.VK_1)
	case "2":
		kb.SetKeys(keybd_event.VK_2)
	case "3":
		kb.SetKeys(keybd_event.VK_3)
	case "4":
		kb.SetKeys(keybd_event.VK_4)
	case "5":
		kb.SetKeys(keybd_event.VK_5)
	case "6":
		kb.SetKeys(keybd_event.VK_6)
	case "7":
		kb.SetKeys(keybd_event.VK_7)
	case "8":
		kb.SetKeys(keybd_event.VK_8)
	case "9":
		kb.SetKeys(keybd_event.VK_9)
	case "0":
		kb.SetKeys(keybd_event.VK_0)
	case "space":
		kb.SetKeys(keybd_event.VK_SPACE)
	case "enter":
		kb.SetKeys(keybd_event.VK_ENTER)
	case "backspace":
		kb.SetKeys(keybd_event.VK_BACKSPACE)
	case "tab":
		kb.SetKeys(keybd_event.VK_TAB)
	case "capslock":
		kb.SetKeys(keybd_event.VK_CAPSLOCK)
	}

	return kb.Launching()
}
