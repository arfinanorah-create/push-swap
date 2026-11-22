// Package tri calcule la liste d'opérations qui trie la pile a.
package tri

import "push-swap/piles"

// seuilBlocs : en dessous, la méthode simple coûte moins cher que les blocs.
const seuilBlocs = 20

// Trier renvoie les opérations qui trient les nombres (vide si déjà trié).
// On choisit la méthode selon le nombre d'éléments.
func Trier(nombres []int) []string {
	if piles.EstTrie(nombres) {
		return []string{}
	}

	// On travaille sur les rangs (0 = le plus petit) : l'ordre est le même,
	// mais les valeurs vont toujours de 0 à n-1.
	p := piles.New(rangs(nombres))

	if len(nombres) == 2 {
		p.Sa()
	} else if len(nombres) == 3 {
		trierTrois(p)
	} else if len(nombres) <= 5 {
		trierCinq(p)
	} else if len(nombres) < seuilBlocs {
		trierSimple(p)
	} else {
		trierParBlocs(p, tailleDeBloc(len(nombres)))
	}
	return p.Operations
}

// rangs renvoie, pour chaque nombre, son rang : le nombre de valeurs
// strictement plus petites. Exemple : [40 10 50 20 30] donne [3 0 4 1 2].
func rangs(nombres []int) []int {
	resultat := make([]int, len(nombres))
	for i, nombre := range nombres {
		for _, autre := range nombres {
			if autre < nombre {
				resultat[i]++
			}
		}
	}
	return resultat
}
