package list

import (
	"WatchTogether-Server/broadcast"
	"WatchTogether-Server/events"
	"WatchTogether-Server/player"
	"bufio"
	"fmt"
	"net"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
)

var videos []string
var videoPlaying string
var retardSeconds int = 40

var clients []net.Conn

var timer *time.Timer
var mu sync.Mutex

func init() {
	events.SubscribeVideoFinish(VideoComplete)
	events.SubscribeClientDisconnect(handleClientDisconnect)
}
//If we dont do this, if client is in video when disconnect the next video never set
func handleClientDisconnect(conn net.Conn){
	mu.Lock()
	clients = slices.DeleteFunc(clients, func(connIt net.Conn) bool {
		var eliminar bool = connIt == conn
		if eliminar{
			fmt.Println("Cliente eliminado de la lista!")
		}
		return eliminar
	})
	mu.Unlock()
}

func VideoComplete(client net.Conn) {
	mu.Lock()
	clients = slices.DeleteFunc(clients, func(clientSlice net.Conn) bool {
		return clientSlice == client
	})
 	mu.Unlock()
	if len(clients) == 0 {
		NextVideo()
	}
}

func SetList(fileName string) {
	ClearList()
	computeList(fileName)
	NextVideo()
}

func NextVideo() {
	video, err := pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	videoPlaying = video
	clients = broadcast.GetConnections()
	if timer != nil{
		timer.Stop()
	}
	timer = player.SetVideoWithRetard(videoPlaying, retardSeconds)
}

func SetRetardSeconds(retard int) {
	retardSeconds = retard
}

func computeList(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("No existe el fichero %s", fileName)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parseVideoAndAdd(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error leyendo %s: %v", fileName, err)
	}
}

func parseVideoAndAdd(video string) {
	if strings.HasPrefix(video, "+") {
		return
	}
	if after, ok := strings.CutPrefix(video, "-"); ok  {
		computeList(after)
		return
	}
	videos = append(videos, video)
}

func pop() (string, error) {
	if len(videos) == 0 {
		return "", fmt.Errorf("Video List is Empty, can not POP")
	}
	var first = videos[0]
	videos = videos[1:]
	return first, nil
}

func ClearList() {
	player.RemoveVideo()
	initializeList()
}

func initializeList() {
	mu.Lock()
	videos = make([]string, 0)
	clients = make([]net.Conn, 0)
	mu.Unlock()
}
