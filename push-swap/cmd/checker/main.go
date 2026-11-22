// checker lit des opérations sur l'entrée standard, les exécute sur la
// pile a reçue en argument, puis affiche OK ou KO.
package main

import (
	"bufio"
	"fmt"
	"os"

	"push-swap/arguments"
	"push-swap/piles"
)

// executer fait tout le travail et renvoie une erreur au lieu d'arrêter le programme.
func executer(args []string) error {
	if len(args) == 0 {
		return nil // pas d'argument : on n'affiche rien
	}

	nombres, err := arguments.LireNombres(args)
	if err != nil {
		return err
	}
	err = arguments.VerifierDoublons(nombres)
	if err != nil {
		return err
	}
	p := piles.New(nombres)

	// Lecture de l'entrée standard, une opération par ligne.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ligne := scanner.Text()
		if ligne == "" {
			continue // on ignore les lignes vides
		}
		err := p.Appliquer(ligne)
		if err != nil {
			return err // opération inconnue
		}
	}

	if p.EstTriee() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
	return nil
}

func main() {
	err := executer(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error") // format imposé par le sujet
		os.Exit(1)
	}
}
