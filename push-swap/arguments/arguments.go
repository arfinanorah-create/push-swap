// Package arguments lit les arguments du programme et les vérifie.
package arguments

import (
	"errors"
	"strconv"
	"strings"
)

// LireNombres transforme les arguments en liste d'entiers.
// Elle marche avec ./prog "2 1 3" comme avec ./prog 2 1 3
// (chaque argument est découpé sur les espaces).
// Elle renvoie une erreur si un mot n'est pas un nombre entier.
func LireNombres(args []string) ([]int, error) {
	nombres := []int{}
	for _, arg := range args {
		for _, mot := range strings.Fields(arg) {
			nombre, err := strconv.Atoi(mot)
			if err != nil {
				return nil, errors.New("pas un nombre entier : " + mot)
			}
			nombres = append(nombres, nombre)
		}
	}
	return nombres, nil
}

// VerifierDoublons renvoie une erreur dès qu'un nombre apparaît deux fois.
func VerifierDoublons(nombres []int) error {
	vus := map[int]bool{}
	for _, nombre := range nombres {
		if vus[nombre] {
			return errors.New("nombre en double : " + strconv.Itoa(nombre))
		}
		vus[nombre] = true
	}
	return nil
}
