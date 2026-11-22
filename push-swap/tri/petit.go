package tri

import "push-swap/piles"

// trierTrois trie une pile a de trois éléments (2 opérations au plus).
func trierTrois(p *piles.Piles) {
	haut := p.A[0]
	milieu := p.A[1]
	bas := p.A[2]

	if haut < milieu && milieu < bas {
		// 1 2 3 : déjà trié
	} else if haut > milieu && milieu < bas && haut < bas {
		p.Sa()
	} else if haut > milieu && milieu < bas && haut > bas {
		p.Ra()
	} else if haut < milieu && milieu > bas && haut > bas {
		p.Rra()
	} else if haut < milieu && milieu > bas && haut < bas {
		p.Sa()
		p.Ra()
	} else {
		p.Sa()
		p.Rra()
	}
}

// trierCinq trie 4 ou 5 éléments : on met de côté les plus petits sur b
// jusqu'à ce qu'il en reste 3 dans a, on trie ces 3, puis on remonte b.
func trierCinq(p *piles.Piles) {
	for len(p.A) > 3 {
		amenerAuSommetA(p, positionDuMinimum(p.A))
		p.Pb()
	}
	trierTrois(p)
	for len(p.B) > 0 {
		p.Pa()
	}
}

// trierSimple : tant que a n'est pas triée, on envoie son minimum sur b,
// puis on remonte tout b. Efficace pour les piles de moins de 20 éléments.
func trierSimple(p *piles.Piles) {
	for !piles.EstTrie(p.A) {
		amenerAuSommetA(p, positionDuMinimum(p.A))
		p.Pb()
	}
	for len(p.B) > 0 {
		p.Pa()
	}
}

// positionDuMinimum renvoie la position du plus petit élément de la liste.
func positionDuMinimum(liste []int) int {
	indexMin := 0
	for i := 1; i < len(liste); i++ {
		if liste[i] < liste[indexMin] {
			indexMin = i
		}
	}
	return indexMin
}

// positionDuMaximum renvoie la position du plus grand élément de la liste.
func positionDuMaximum(liste []int) int {
	indexMax := 0
	for i := 1; i < len(liste); i++ {
		if liste[i] > liste[indexMax] {
			indexMax = i
		}
	}
	return indexMax
}

// amenerAuSommetA fait remonter l'élément de a en position donnée,
// par le chemin le plus court (ra ou rra).
func amenerAuSommetA(p *piles.Piles, position int) {
	if position < len(p.A)/2 {
		for i := 0; i < position; i++ {
			p.Ra()
		}
	} else {
		for i := 0; i < len(p.A)-position; i++ {
			p.Rra()
		}
	}
}

// amenerAuSommetB : même chose sur b (rb ou rrb).
func amenerAuSommetB(p *piles.Piles, position int) {
	if position < len(p.B)/2 {
		for i := 0; i < position; i++ {
			p.Rb()
		}
	} else {
		for i := 0; i < len(p.B)-position; i++ {
			p.Rrb()
		}
	}
}
