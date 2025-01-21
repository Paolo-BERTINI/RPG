package creation_etre

import (
	parser "base_game/Parser"
	"fmt"
	"os"
)

type Etres struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	HP     int    `json:"hp"`
	MP     int    `json:"mp"`
	Str    int    `json:"str"`
	Int    int    `json:"int"`
	Def    int    `json:"def"`
	Res    int    `json:"res"`
	Spd    int    `json:"spd"`
	Luck   int    `json:"luck"`
	Race   int    `json:"race"`
	Class  int    `json:"class"`
	Rarity int    `json:"rarity"`
}

func Enemy() Etres {
	var Enemy = Etres{
		ID:   "12",
		Name: "Bokoblin",
		HP:   25,
		Str:  10,
	}
	return Enemy
}

func Boss() Etres {
	var Boss = Etres{
		ID:   "1",
		Name: "Ganon",
		HP:   350,
		Str:  15,
	}
	return Boss
}

func Choix_perso() parser.Etres {
	persos, err := parser.Parser_J("JSON/players.json")
	if err != nil {
		fmt.Println("Erreur:", err)
		os.Exit(0)
	}
	fmt.Println("Personnages disponible :")
	for i, perso := range persos {
		fmt.Println()
		fmt.Println("[", i+1, "] -", perso.Name)
		fmt.Println("	   Health Points :", perso.HP)
		fmt.Println("	   Mana Points :", perso.MP)
		fmt.Println("	   Strength :", perso.Str)
	}
	fmt.Println()
	fmt.Print("Choisis un personnage : ")
	var id_choisis string
	fmt.Scan(&id_choisis)
	for _, perso := range persos {
		if id_choisis == perso.ID {
			fmt.Println("Tu as choisis", perso.Name)
			return perso
		}
	}
	return parser.Etres{}
}
