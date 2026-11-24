# Push-swap
 
Push-swap est un projet d'algorithmique écrit en **Go**. Il consiste à trier une liste de nombres en n'utilisant que **deux piles** et **onze instructions**, avec le moins d'instructions possible.
 
Le projet contient deux programmes :
 
- **`push-swap`** : reçoit des nombres et affiche la liste d'instructions qui les trie.
- **`checker`** : reçoit des nombres, lit des instructions, les exécute, puis affiche `OK` si les nombres sont triés, `KO` sinon.
## Sommaire
 
1. [Le problème](#le-problème)
2. [Les instructions](#les-instructions)
3. [Prérequis](#prérequis)
4. [Utilisation](#utilisation)
5. [Gestion des erreurs](#gestion-des-erreurs)
6. [Comment ça marche](#comment-ça-marche)
7. [Organisation du projet](#organisation-du-projet)
8. [Tests](#tests)
9. [Performances](#performances)
## Le problème
 
On a deux piles, **a** et **b**. Au départ, `a` contient les nombres dans le désordre et `b` est vide. Le but est d'obtenir `a` **triée dans l'ordre croissant** (le plus petit nombre au sommet) et `b` **vide**.
 
Dans une pile, on ne peut toucher qu'au **sommet**. On ne peut pas échanger deux éléments éloignés l'un de l'autre. C'est ce qui rend le problème intéressant : un algorithme de tri classique ne s'applique pas directement.
 
Le résultat attendu n'est pas la pile triée, mais **la liste des instructions** qui permet de la trier. C'est cette liste que `push-swap` affiche, et c'est son nombre de lignes qu'on cherche à réduire.
 
## Les instructions
 
Le sommet d'une pile est son premier élément.
 
| Instruction | Effet |
|---|---|
| `sa` | Échange les 2 premiers éléments de `a` |
| `sb` | Échange les 2 premiers éléments de `b` |
| `ss` | Fait `sa` et `sb` en même temps |
| `pa` | Prend le sommet de `b` et le met au sommet de `a` |
| `pb` | Prend le sommet de `a` et le met au sommet de `b` |
| `ra` | Décale `a` vers le haut : le premier élément devient le dernier |
| `rb` | Décale `b` vers le haut |
| `rr` | Fait `ra` et `rb` en même temps |
| `rra` | Décale `a` vers le bas : le dernier élément devient le premier |
| `rrb` | Décale `b` vers le bas |
| `rrr` | Fait `rra` et `rrb` en même temps |
 
Une instruction impossible (par exemple `pa` quand `b` est vide) ne fait rien.
 
### Exemple
 
On part de `a` = `2 1 3 6 8 5` (le premier est au sommet) et de `b` vide.
 
| Instructions | Pile `a` | Pile `b` |
|---|---|---|
| (départ) | 2 1 3 6 8 5 | vide |
| `sa` | 1 2 3 6 8 5 | vide |
| `pb` `pb` `pb` | 6 8 5 | 3 2 1 |
| `rb` | 6 8 5 | 2 1 3 |
| `rra` `rrb` (ou `rrr`) | 5 6 8 | 3 2 1 |
| `pa` `pa` `pa` | 1 2 3 5 6 8 | vide |
 
## Prérequis
 
- [Go](https://go.dev/dl/) version 1.18 ou plus récente.
- Aucun package externe : seule la bibliothèque standard de Go est utilisée.
Pour vérifier que Go est installé :
 
```
go version
```
 
## Utilisation
 
Toutes les commandes se tapent **à la racine du projet**, c'est-à-dire dans le dossier qui contient le fichier `go.mod`.
 
### push-swap
 
Les nombres se donnent dans un seul argument entre guillemets, ou dans plusieurs arguments :
 
```
go run ./cmd/push-swap "4 67 3 87 23"
go run ./cmd/push-swap 4 67 3 87 23
```
 
Résultat (une instruction par ligne) :
 
```
rra
rra
rra
pb
rra
rra
pb
rra
pa
pa
```
 
Pour compter les instructions :
 
- Mac, Linux ou Git Bash : `go run ./cmd/push-swap "4 67 3 87 23" | wc -l`
- PowerShell : `go run ./cmd/push-swap "4 67 3 87 23" | Measure-Object -Line`
### checker
 
Le checker reçoit les nombres en argument et lit les instructions sur l'entrée standard, une par ligne.
 
Mac, Linux ou Git Bash :
 
```
printf "sa\nrra\npb\n" | go run ./cmd/checker "3 2 1 0"
```
 
Résultat : `KO`, car après ces instructions la pile `a` n'est pas triée et `b` n'est pas vide.
 
### Les deux ensemble
 
C'est la façon normale de vérifier que `push-swap` trie correctement.
 
Mac, Linux ou Git Bash :
 
```
ARG="4 67 3 87 23"; go run ./cmd/push-swap "$ARG" | go run ./cmd/checker "$ARG"
```
 
PowerShell :
 
```
$ARG = "4 67 3 87 23"
go run ./cmd/push-swap $ARG | go run ./cmd/checker $ARG
```
 
Résultat : `OK`.
 
### Cas particuliers
 
| Commande | Résultat |
|---|---|
| `go run ./cmd/push-swap` (sans argument) | Rien n'est affiché |
| `go run ./cmd/push-swap "1 2 3"` (déjà trié) | Rien n'est affiché |
| `go run ./cmd/push-swap "42"` (un seul nombre) | Rien n'est affiché |
| `go run ./cmd/checker` (sans argument) | Rien n'est affiché |
 
### Créer les exécutables (facultatif)
 
`go run` compile et lance en une fois. Pour obtenir de vrais fichiers exécutables nommés `push-swap` et `checker` :
 
```
go build -o push-swap ./cmd/push-swap
go build -o checker ./cmd/checker
```
 
On les lance ensuite avec `./push-swap` et `./checker` (sous Windows, `push-swap.exe` et `checker.exe`).
 
## Gestion des erreurs
 
En cas d'erreur, le programme affiche `Error` suivi d'un retour à la ligne sur la **sortie d'erreur**, puis s'arrête avec le code de sortie 1. Il ne plante jamais.
 
| Situation | Exemple |
|---|---|
| Un argument n'est pas un entier | `go run ./cmd/push-swap "0 one 2 3"` |
| Un nombre apparaît deux fois | `go run ./cmd/push-swap "1 2 2"` |
| (checker) Une instruction n'existe pas ou est mal écrite | `echo foo \| go run ./cmd/checker "1 2"` |
 
Les lignes vides lues par le checker sont ignorées.
 
## Comment ça marche
 
### Le principe de `push-swap`
 
1. Il lit les arguments et vérifie qu'ils sont valides.
2. Si les nombres sont déjà triés, il n'affiche rien.
3. Il remplace chaque nombre par son **rang** : 0 pour le plus petit, 1 pour le suivant, jusqu'à n − 1 pour le plus grand. L'ordre est le même, mais les valeurs vont toujours de 0 à n − 1, même avec des nombres négatifs.
4. Il choisit un algorithme selon le nombre d'éléments, exécute les instructions sur les piles en les enregistrant, puis les affiche.
### Les algorithmes
 
| Nombres | Méthode | Idée |
|---|---|---|
| 2 | `sa` | Un simple échange |
| 3 | `trierTrois` | Les six cas possibles, 2 instructions au maximum |
| 4 ou 5 | `trierCinq` | On envoie les plus petits sur `b` jusqu'à ce qu'il en reste 3 dans `a`, on trie ces 3, puis on remonte `b` |
| 6 à 19 | `trierSimple` | Tant que `a` n'est pas triée, on amène son minimum au sommet et on l'envoie sur `b`. À la fin, on remonte tout `b` |
| 20 et plus | `trierParBlocs` | Stratégie par blocs, expliquée ci-dessous |
 
Pour amener un élément au sommet, on choisit le sens de rotation le plus court : `ra` si l'élément est dans la première moitié de la pile, `rra` sinon (de même `rb` et `rrb` pour `b`).
 
### La stratégie par blocs
 
Elle se fait en deux phases.
 
**Phase 1 : vider `a` dans `b`.** On garde un compteur `pousses` (le nombre d'éléments déjà envoyés sur `b`) et une taille de bloc (15, ou 30 au-delà de 100 nombres). Tant que `a` n'est pas vide, on regarde le rang au sommet de `a` :
 
| Si le rang au sommet est… | On fait | Pourquoi |
|---|---|---|
| inférieur ou égal à `pousses` | `pb` puis `rb` | C'est un des plus petits : on l'envoie au fond de `b` |
| inférieur ou égal à `pousses + taille du bloc` | `pb` | Il est dans le bloc en cours : on le pose au sommet de `b` |
| plus grand | `ra` | Trop grand pour l'instant : il passe en dessous et reviendra |
 
**Phase 2 : remonter `b` dans `a`.** Tant que `b` n'est pas vide : on cherche le plus grand rang de `b`, on l'amène au sommet par le chemin le plus court (`rb` ou `rrb`), puis on fait `pa`. On remonte le plus grand en premier, donc `a` se remplit dans le bon ordre.
 
Après la phase 1, le plus grand rang restant est presque toujours près d'une extrémité de `b`, ce qui rend la phase 2 rapide.
 
### Le principe du `checker`
 
1. Il lit les nombres en argument et crée la pile `a`.
2. Il lit chaque ligne de l'entrée standard et exécute l'instruction correspondante. Une instruction inconnue provoque `Error`.
3. À la fin, il affiche `OK` si `a` est triée et `b` vide, sinon `KO`.
## Organisation du projet
 
```
push-swap/
├── go.mod                  Nom du module Go ("push-swap")
├── README.md               Ce fichier
├── cmd/
│   ├── push-swap/main.go   Programme push-swap
│   └── checker/main.go     Programme checker
├── piles/
│   ├── piles.go            Les piles a et b et les 11 instructions
│   └── piles_test.go
├── arguments/
│   ├── arguments.go        Lecture des arguments et vérification des doublons
│   └── arguments_test.go
└── tri/
    ├── tri.go              Choix de l'algorithme et calcul des rangs
    ├── petit.go            Tri de 2 à 19 nombres
    ├── blocs.go            Stratégie par blocs
    └── tri_test.go
```
 
Chaque dossier a une seule responsabilité :
 
- **`cmd`** : les deux programmes (un dossier chacun, car Go n'autorise qu'un seul `main` par dossier). Leur code est très court : ils appellent les autres dossiers.
- **`piles`** : les deux piles et les onze instructions. Les deux programmes l'utilisent, donc les instructions ne sont écrites qu'une fois. Ce dossier enregistre aussi chaque instruction exécutée.
- **`arguments`** : transforme le texte tapé en nombres (`LireNombres`) et refuse les doublons (`VerifierDoublons`).
- **`tri`** : les algorithmes qui décident quelles instructions utiliser. Seul `push-swap` s'en sert.

## Tests
 
Les commandes se tapent à la racine du projet, c'est-à-dire dans le dossier `push-swap/` qui contient le fichier `go.mod`.
 
Tous les tests :
 
```bash
go test ./...
```
 
Avec le détail de chaque test (plus bavard) :
 
```bash
go test -v ./...
```
 
Un seul dossier :
 
```bash
go test ./piles
go test ./arguments
go test ./tri
```
 
Un seul test, par son nom :
 
```bash
go test ./tri -run NomDuTest
```
 
Si on relance les tests sans avoir rien modifié, Go affiche `(cached)` et réutilise le résultat précédent. Pour forcer leur exécution :
 
```bash
go test -count=1 ./...
```
 
Un résultat réussi ressemble à ceci :
 
```text
ok      push-swap/piles       0.003s
ok      push-swap/arguments   0.002s
ok      push-swap/tri         0.150s
?       push-swap/cmd/checker [no test files]
```
 
Les lignes `[no test files]` sont normales : les dossiers `cmd` n'ont pas de tests.
 
Les tests vérifient notamment :
 
- l'exemple du sujet, instruction par instruction ;
- les instructions sur des piles vides ou trop petites ;
- les erreurs (non-entiers, doublons, instructions inconnues) ;
- que les instructions générées trient bien des centaines de piles aléatoires, de 2 à 500 nombres, y compris avec des nombres négatifs.

## Performances
 
Nombre moyen d'instructions mesuré sur des listes aléatoires (moyenne de 10 essais) :
 
| Nombres | Instructions (environ) |
|---|---|
| 10 | 29 |
| 50 | 248 |
| 100 | 580 |
| 500 | 5 050 |
 
 
