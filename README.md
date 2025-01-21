# RPG - Jeu de rôle

## Description
Ce projet est un jeu de rôle (RPG) textuel développé en Go. Il permet aux joueurs d'explorer des étages, d'affronter des ennemis, de collecter des objets et de progresser dans l'histoire.

## Fonctionnalités principales
- **Chargement dynamique des données** :
  - Les informations sur les joueurs, ennemis, boss, pièges et sorts sont chargées depuis des fichiers JSON.
- **Système d'étages** :
  - Exploration des différents niveaux où les joueurs peuvent rencontrer des ennemis et des pièges.
- **Gestion des entités** :
  - Création et gestion des personnages joueurs et ennemis.

## Structure du projet
- `main.go` : Point d'entrée du programme. Initialise les données et démarre le jeu.
- `Creation_etre/` : Contient le code pour la création et la gestion des entités du jeu.
- `Etage/` : Contient les mécaniques liées aux étages et à la progression des joueurs.
- `Parser/` : Contient les fonctions pour charger les données JSON (joueurs, boss, etc.).
- `JSON/` : Dossier contenant les fichiers de données JSON utilisés par le jeu.

## Démarrage
### Prérequis
- [Go](https://go.dev/dl/) (version 1.16 ou ultérieure)

## Organisation des fichiers JSON
- `bosses.json` : Définit les caractéristiques des boss.
- `classes.json` : Décrit les classes disponibles pour les joueurs.
- `enemies.json` : Liste les ennemis que les joueurs peuvent rencontrer.
- `players.json` : Contient les informations des joueurs.
- `spells.json` : Détaille les sorts utilisables par les entités.
- `traps.json` : Définit les pièges présents dans les étages.
