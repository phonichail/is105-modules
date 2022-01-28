package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	human := initActor("👨", "Human")
	fox := initActor("🦊", "Fox")
	chicken := initActor("🐔", "Chicken")
	corn := initActor("🌽", "Corn")
	actors := [4]actor{*human, *fox, *chicken, *corn}
	fmt.Println(crossRiver(&actors[0], &actors[2]))

	expectedOutcome := false

	if Eats(&actors) == expectedOutcome {
		t.Errorf("Dead")
	}

	fmt.Println(crossRiver(&actors[0], &actors[0]))

	if Eats(&actors) == expectedOutcome {
		t.Errorf("Dead")
	}

	fmt.Println(crossRiver(&actors[0], &actors[1]))

	if Eats(&actors) == expectedOutcome {
		t.Errorf("Dead")
	}

	fmt.Println(crossRiver(&actors[0], &actors[2]))

	if Eats(&actors) == expectedOutcome {
		t.Errorf("Dead")
	}

	fmt.Println(crossRiver(&actors[0], &actors[3]))

	if Eats(&actors) == expectedOutcome {
		t.Errorf("Dead")
	}

	fmt.Println(crossRiver(&actors[0], &actors[0]))

	if Eats(&actors) == expectedOutcome {
		t.Errorf("Dead")
	}

	fmt.Println(crossRiver(&actors[0], &actors[2]))
	/*

			func init() {
			human := initActor("👨", "Human")
			fox := initActor("🦊", "Fox")
			chicken := initActor("🐔", "Chicken")
			corn := initActor("🌽", "Corn")
			actors := [4]actor{*human, *fox, *chicken, *corn}
			crossRiver(&actors[0], &actors[1])
			Eats(&actors)
		}
					if alive !=  {
				        t.Errorf("Dead", got)
				    }
	*/
}
