package tri

import "push-swap/piles"

// tailleDeBloc choisit la taille des blocs selon le nombre d'éléments
// (valeurs trouvées en mesurant : 15 pour 100 éléments, 30 pour 500).
func tailleDeBloc(n int) int {
	if n <= 100 {
		return 15
	}
	return 30
}

// trierParBlocs trie a en deux phases : pousser par blocs, puis remonter.
func trierParBlocs(p *piles.Piles, tailleBloc int) {
	pousserParBlocs(p, tailleBloc)
	remonterLesPlusGrands(p)
}

// pousserParBlocs vide a dans b, en ne poussant que les rangs du bloc en cours.
// Les plus petits partent au fond de b.
func pousserParBlocs(p *piles.Piles, tailleBloc int) {
	pousses := 0 // nombre d'éléments déjà envoyés sur b

	for len(p.A) > 0 {
		rang := p.A[0]
		if rang <= pousses {
			// un des plus petits : au fond de b
			p.Pb()
			p.Rb()
			pousses++
		} else if rang <= pousses+tailleBloc {
			// dans le bloc en cours : au sommet de b
			p.Pb()
			pousses++
		} else {
			// trop grand pour l'instant : il passe en dessous, il reviendra
			p.Ra()
		}
	}
}

// remonterLesPlusGrands remonte b dans a, du plus grand rang au plus petit.
func remonterLesPlusGrands(p *piles.Piles) {
	for len(p.B) > 0 {
		amenerAuSommetB(p, positionDuMaximum(p.B))
		p.Pa()
	}
}
