package player

import (
	"WatchTogether-Server/broadcast"
	"WatchTogether-Server/distpacher"
	"time"
)


func SetVideoWithRetard(url string, secondsWait int){
	broadcast.SenAll(distpacher.WriteSetVideoPacket(url))
	Play(false)
	//Ponemos el video y el play true es mas tarde para esperar a que cargue
	time.AfterFunc(time.Duration(secondsWait)*time.Second, func() {
		SetTime(0)
		Play(true)
	})
}


func Play(play bool){
	broadcast.SenAll(distpacher.WritePlayPusePacket(play))

}

func SetTime(time float64){
	broadcast.SenAll(distpacher.WriteSetTimePacket(time))

}

func FinishedEvent(){
	
}

//This function need to configure a list of videos, when complete server call All clients finished and this ned to set the next
// respect special command, nothin is a name of a video
//  "-" its a new list file <- add this videos to the current listName
//  "+" ignore entry
func SetList(listName string){
	
}

