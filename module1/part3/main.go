package main

import (
	"fmt"
)

// var boat = "⛵"
var river = "🌊🌊🌊🌊🌊"
var human = "👨"

var fox = "🦊"
var chicken = "🐔"
var corn = "🌽"

type actor struct {
	icon     string
	name     string
	position string
	alive    bool
}

type boat struct {
	icon     string
	name     string
	position string
	carries  *actor
}

func initActor(icon_ string, name_ string) *actor {
	newActor := actor{
		icon:     icon_,
		name:     name_,
		position: "left",
		alive:    true}
	return &newActor
}

func printActorState(myActor *actor) string {
	actorState := "The " + myActor.icon + " is currently at the " + myActor.position + " of the " + river + "."
	return actorState
}

func crossRiver(transporter *actor, transported *actor) string {
	msg := ""
	if transported.alive == false {
		msg = transported.icon + " is dead!"
	} else {

		if transporter.position == transported.position {
			if transporter.position == "left" {
				transporter.position = "right"
				transported.position = "right"

			} else {

				transporter.position = "left"
				transported.position = "left"
			}

			msg = transporter.icon + " transports " + transported.icon + " to " + transported.position
		} else {
			msg = transporter.icon + " is not at the same position as " + transported.icon
		}
	}
	return msg
}

func putInBoat() {

}

func viewState() {

}

func Eats(actors *[4]actor) bool {
	if actors[0].position != actors[1].position && actors[1].position == actors[2].position {
		actors[2].alive = false
		fmt.Println(actors[1].icon, "eats", actors[2].icon)
	}
	if actors[0].position != actors[2].position && actors[2].position == actors[3].position {
		actors[3].alive = false
		fmt.Println(actors[2].icon, "eats", actors[3].icon)
	}
	return true
}

func main() {

	human := initActor("👨", "Human")
	fox := initActor("🦊", "Fox")
	chicken := initActor("🐔", "Chicken")
	corn := initActor("🌽", "Corn")
	actors := [4]actor{*human, *fox, *chicken, *corn}

	fmt.Println(crossRiver(&actors[0], &actors[2]))
	Eats(&actors)
	fmt.Println()

	/*
		boat := boat{icon: "⛵",
			name:     "Boat",
			position: "left",
		}
	*/
}
