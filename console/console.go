package console

import (
	"WatchTogether-Server/list"
	"WatchTogether-Server/player"
	"fmt"
	"strconv"
	"strings"

	"github.com/abiosoft/ishell/v2"
)

func Start() {
	shell := ishell.New()
	shell.SetPrompt("\033[32m>\033[0m ")

	commands(shell)

	shell.Run()
}

func commands(shell *ishell.Shell) {

	shell.AddCmd(setVideoCommand())
	shell.AddCmd(playCommand())
	shell.AddCmd(setTimeCommand())
	shell.AddCmd(setListCommand())
	shell.AddCmd(cancelVideo())
	shell.AddCmd(cancelList())
}

func setVideoCommand() *ishell.Cmd {
	return &ishell.Cmd{
		Name: "SetVideo",
		Help: "Pone un video a todos los clientes SetVideo <time> <name>",
		Func: func(c *ishell.Context) {
			if len(c.Args) < 2 {
				fmt.Println("Uso Correcto SetVideo <time> <name>")
				return
			}

			timeString := c.Args[0]
			url := strings.Join(c.Args[1:], " ")
			timeWait, _ := strconv.Atoi(timeString)
			player.SetVideoWithRetard(url, timeWait)

		},
	}
}

func playCommand() *ishell.Cmd {
	return &ishell.Cmd{
		Name: "Play",
		Help: "Setea el valor Play en todos los clientes Play <false/true>",
		Completer: func(args []string) []string {
			return []string{"true", "false"}
		},
		Func: func(c *ishell.Context) {
			if len(c.Args) < 1 {
				fmt.Println("Uso Correcto Play <false/true>")
				return
			}

			playString := c.Args[0]
			play, _ := strconv.ParseBool(playString)
			player.Play(play)

		},
	}
}

func setTimeCommand() *ishell.Cmd {
	return &ishell.Cmd{
		Name: "SetTime",
		Help: "Mueve el tiempo del video en todos los clientes",
		Func: func(c *ishell.Context) {
			if len(c.Args) < 1 {
				fmt.Println("Uso Correcto SetTime <segundos>")
				return
			}
			timeString := c.Args[0]
			time, _ := strconv.ParseFloat(timeString, 64)
			player.SetTime(time)
		},
	}
}

func setListCommand() *ishell.Cmd {
	return &ishell.Cmd{
		Name: "SetList",
		Help: "Setea una lista de reproduccion",
		Func: func(c *ishell.Context) {
			if len(c.Args) < 1 {
				fmt.Println("Uso Correcto SetList <fileName> <secRetard> <- Opcional")
				return
			}
		
			if len(c.Args) > 1 {
			    var retard, _ = strconv.Atoi(c.Args[1])
			    list.SetRetardSeconds(retard)

			}
			list.SetList(c.Args[0])

		},
	}
}

func cancelVideo() *ishell.Cmd {
	return &ishell.Cmd{
		Name: "CancelVideo",
		Help: "Manda a todos los clientes el CancelVideoPacket",
		Func: func(c *ishell.Context) {
			fmt.Println("Enviando cancelar video...")
			player.RemoveVideo()
		},
	}
}



func cancelList() *ishell.Cmd {
	return &ishell.Cmd{
		Name: "CancelList",
		Help: "Cancela la lista y para el video actual",
		Func: func(c *ishell.Context) {
			fmt.Println("Cancelando Lista...")
			list.ClearList()
		},
	}
}
