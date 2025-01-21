package parser

import (
	"encoding/json"
	"errors"
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

type PlayerList struct {
	Etre []Etres `json:"Player"`
}

type EnemyList struct {
	Etre []Etres `json:"Enemy"`
}

func Parser_J(filepath string) ([]Etres, error) {
	file, err := os.Open(filepath)
	fmt.Println(filepath)
	if err != nil {
		return nil, errors.New("Erreur lors de l'ouverture du fichier JSON\n" + err.Error())
	}
	defer file.Close()
	var etresList PlayerList
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&etresList)
	if err != nil {
		return nil, errors.New("Erreur lors de la lecture des données JSON" + err.Error())
	}
	return etresList.Etre, nil
}
/* 
func Parser_E(filepath string) ([]Etres, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, errors.New("Erreur lors de l'ouverture du fichier JSON\n" + err.Error())
	}
	defer file.Close()
	var etresList EnemyList
	fmt.Print(etresList)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&etresList)
	if err != nil {
		return nil, errors.New("Erreur lors de la lecture des données JSON" + err.Error())
	}
	return etresList.Etre, nil
}
 */