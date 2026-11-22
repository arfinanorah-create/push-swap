// push-swap affiche les opérations qui trient les nombres reçus en argument.
package main

import (
	"fmt"
	"os"

	"push-swap/arguments"
	"push-swap/tri"
)

// executer fait tout le travail et renvoie une erreur au lieu d'arrêter le programme.
func executer(args []string) error {
	if len(args) == 0 {
		return nil // rien à trier : on n'affiche rien
	}

	nombres, err := arguments.LireNombres(args)
	if err != nil {
		return err
	}
	err = arguments.VerifierDoublons(nombres)
	if err != nil {
		return err
	}

	for _, operation := range tri.Trier(nombres) {
		fmt.Println(operation)
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
