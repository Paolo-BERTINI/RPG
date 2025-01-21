package etage

import (
	creation_etre "base_game/Creation_etre"
	"fmt"
	"os"
)

func Etage() {
	Perso := creation_etre.Choix_perso()
	Boss := creation_etre.Boss()
	var ln int = 25
	var lb int = 350
	var lp int = 60
	fmt.Println()
	for i := 0; i < 9; i++ {
		fmt.Println("============= Floor", i+1, "=============")
		fmt.Println()
		Enemy := creation_etre.Enemy()
		fmt.Println("Tu te retrouves né à né avec un", Enemy.Name)
		for Enemy.HP > 0 {
			fmt.Println()
			fmt.Println("\033[32m", Perso.Name, "\033[0m")
			for j := 0; j < Perso.HP; j++ {
				fmt.Print("\033[32mI\033[0m")
			}
			fmt.Println("	\033[32m", Perso.HP, " / ", lp, "\033[0m")
			fmt.Println()
			fmt.Println("\033[31m", Enemy.Name, "\033[0m")
			for j := 0; j < Enemy.HP; j++ {
				fmt.Print("\033[31mI\033[0m")
			}
			fmt.Println("	\033[31m", Enemy.HP, " / ", ln, "\033[0m")
			fmt.Println()
			fmt.Println("\033[34m[ 1 ] Attaquer\033[0m / \033[33m[ 2 ] Guérir\033[0m")
			fmt.Print("Choisis une action : ")
			var action int
			fmt.Scan(&action)
			fmt.Println()
			if action == 1 {
				fmt.Println("Tu attaques le", Enemy.Name, ". Il perd", Perso.Str, "HP")
				Enemy.HP -= Perso.Str
			} else if action == 2 {
				fmt.Println("Tu guéris de", Perso.MP, "HP")
				Perso.HP += Perso.MP
				if Perso.HP > lp {
					Perso.HP = lp
				}
			}
			if Enemy.HP < 0 {
				fmt.Println("Tu as vaincu le", Enemy.Name)
				fmt.Println()
				fmt.Println()
			} else {
				fmt.Println("Le", Enemy.Name, "t'attaques. Tu perds", Boss.Str, "HP")
				Perso.HP -= Enemy.Str
				if Perso.HP <= 0 {
					fmt.Println("Tu n'as plus de HP")
					fmt.Println("L'aventure s'arrêtes pour toi, Jeune Héros")
					os.Exit(0)
				}
			}
		}
	}
	fmt.Println("============= Floor 10 =============")
	fmt.Println()
	fmt.Println("Tu entres dans la salle du boss", Boss.Name)
	for Boss.HP > 0 {
		fmt.Println()
		fmt.Println("\033[32m", Perso.Name, "\033[0m")
		for j := 0; j < Perso.HP; j++ {
			fmt.Print("\033[32mI\033[0m")
		}
		fmt.Println("	\033[32m", Perso.HP, " / ", lp, "\033[0m")
		fmt.Println()
		fmt.Println("\033[30m", Boss.Name, "\033[0m")
		for j := 0; j < Boss.HP; j++ {
			fmt.Print("\033[30mI\033[0m")
		}
		fmt.Print("	", Boss.HP, " / ", lb)
		fmt.Println()
		fmt.Println()
		fmt.Println("[ 1 ] Attaquer / [ 2 ] Guérir")
		fmt.Print("Choisis une action : ")
		var action int
		fmt.Scan(&action)
		fmt.Println()
		if action == 1 {
			fmt.Println("Tu attaques ", Boss.Name, ". Il perd", Perso.Str, "HP")
			Boss.HP -= Perso.Str
		} else if action == 2 {
			fmt.Println("Tu guéris de", Perso.MP, "HP")
			Perso.HP += Perso.MP
			if Perso.HP > lp {
				Perso.HP = lp
			}
		}
		if Boss.HP < 0 {
			fmt.Println("Tu as vaincu ", Boss.Name)
		} else {
			fmt.Println(Boss.Name, "t'attaques. Tu perds", Boss.Str, "HP")
			Perso.HP -= Boss.Str
			if Perso.HP <= 0 {
				fmt.Println("Tu n'as plus de HP")
				fmt.Println("L'aventure s'arrêtes pour toi, Jeune Héros")
				os.Exit(0)
			}
		}
	}
	fmt.Println("Tu as sauvé la Princesse Zelda")
}
