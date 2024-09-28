package functions

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	welcomeMessage = strings.Join([]string{
		"Welcome to TCP-Chat!",
		"         _nnnn_",
		"        dGGGGMMb",
		"       @p~qp~~qMb",
		"       M|@||@) M|",
		"       @,----.JM|",
		"      JS^\\__/  qKL",
		"     dZP        qKRb",
		"    dZP          qKKb",
		"   fZP            SMMb",
		"   HZM            MMMM",
		"   FqM            MMMM",
		" __| \".        |\\dS\"qML",
		" |    `.       | `' \\Zq",
		"_)      \\.___.,|     .'",
		"\\____   )MMMMMP|   .'",
		"     `-'       `--'\n",
	}, "\n")
	mutex sync.Mutex
)

func HandleConnection(conn net.Conn, clients map[string]net.Conn, connectedClients chan struct{}, messageHistory *string) {
	defer conn.Close()

	_, err := conn.Write([]byte(welcomeMessage))
	if err != nil {
		return
	}

	var name string

	message := make([]byte, 1024)
	userName := true

userName:
	_, err = conn.Write([]byte("[ENTER YOUR NAME]:"))
	if err != nil {
		return
	}
	for {
		n, err := conn.Read(message)
		if err != nil {
			break
		}
		if userName {
			if strings.TrimSpace(string(message[:n-1])) == "" || !printable([]byte(strings.TrimSpace(string(message[:n-1])))) {
				goto userName
			}

			name = string(strings.TrimSpace(string(message[:n-1])))

			mutex.Lock()
			if _, exist := clients[name]; exist {
				mutex.Unlock()
				_, err = conn.Write([]byte("*_* THIS NAME IS IN USE *_*\n"))
				if err != nil {
					break
				}
				goto userName
			} else {
				mutex.Unlock()
			}

			if len(connectedClients) == 2 {
				_, err = conn.Write([]byte("   THE CHAT IS FULL, TRY AGAIN LATER\n"))
				if err != nil {
					break
				}
			}
			connectedClients <- struct{}{}

			mutex.Lock()
			userName = false
			clients[name] = conn
			mutex.Unlock()
			prompt := fmt.Sprintf("[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), name)
			_, err = conn.Write([]byte(*messageHistory + prompt))
			if err != nil {
				break
			}

			joiningMsg := fmt.Sprintf("-----%s HAS JOINED THE CHAT-----", name)
			mutex.Lock()
			*messageHistory += joiningMsg + "\n"
			mutex.Unlock()

			broadcast(name, clients, joiningMsg)

		} else {
			// Formatted message with current timestamp and sender name
			if len(strings.TrimSpace(string(message[:n-1]))) == 0 || !printable([]byte(strings.TrimSpace(string(message[:n-1])))) {
				// Optionally, you can send a prompt back to the client if needed
				prompt := fmt.Sprintf("[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), name)
				_, err := conn.Write([]byte(prompt))
				if err != nil {
					break
				}
				continue // Skip further processing for empty messages
			}
			formattedMessage := fmt.Sprintf("[%s][%s]:%s", time.Now().Format("2006-01-02 15:04:05"), name, string(message[:n-1]))
			mutex.Lock()
			*messageHistory += formattedMessage + "\n"
			mutex.Unlock()

			// Broadcast the message to other clients
			broadcast(name, clients, formattedMessage)

			// Re-display the prompt for the sender after their message
			prompt := fmt.Sprintf("[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), name)
			_, err := conn.Write([]byte(prompt))
			if err != nil {
				break
			}
		}
	}

	// Broadcast to other clients that the user has left the chat
	mutex.Lock()
	if _, exist := clients[name]; exist {
		leaveMessage := fmt.Sprintf("%s HAS LEFT OUR CHAT...", name)
		*messageHistory += leaveMessage + "\n"
		delete(clients, name)
		<-connectedClients
		 mutex.Unlock()
		broadcast(name, clients, leaveMessage)
	} else {
		 mutex.Unlock()
	 }
}

func printable(data []byte) bool {
	for _, c := range data {
		if c < 32 {
			return false
		}
	}
	return true
}

func broadcast(sender string, clients map[string]net.Conn, msg string) {
	defer mutex.Unlock()
	mutex.Lock()
	for clientName, conn := range clients {
		if clientName != sender {
			// Send the broadcast message to other clients
			prompt := fmt.Sprintf("[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), clientName)
			_, err := conn.Write([]byte("\n" + msg + "\n" + prompt))
			if err != nil {
				continue
			}
		}
	}
}
