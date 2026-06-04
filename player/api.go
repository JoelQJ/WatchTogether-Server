package player

import (
	"WatchTogether-Server/broadcast"
	"WatchTogether-Server/distpacher"
	"time"
)

func SetVideoWithRetard(url string, secondsWait int) {
	broadcast.SenAll(distpacher.WriteSetVideoPacket(url))
	Play(false)
	//Ponemos el video y el play true es mas tarde para esperar a que cargue
	time.AfterFunc(time.Duration(secondsWait)*time.Second, func() {
		SetTime(0)
		Play(true)
	})
}

func RemoveVideo() {
	//Todo WE NEED TO DO THE PACKET
}

func Play(play bool) {
	broadcast.SenAll(distpacher.WritePlayPusePacket(play))
}

func SetTime(time float64) {
	broadcast.SenAll(distpacher.WriteSetTimePacket(time))

}