# push-swap

Deux programmes en Go (bibliothèque standard uniquement).

## Lancer
    go run ./cmd/push-swap "4 67 3 87 23"
    ARG="4 67 3 87 23"; go run ./cmd/push-swap "$ARG" | wc -l
    ARG="4 67 3 87 23"; go run ./cmd/push-swap "$ARG" | go run ./cmd/checker "$ARG"   # OK

## Tests
    go test ./...

## Organisation
- piles/   : les 2 piles + les 11 instructions (utilisé par les deux programmes)
- arguments/: lecture des arguments et vérification des doublons
- tri/     : 2-3 nombres, 4-5 nombres, méthode simple (< 20), stratégie par blocs (>= 20)
- cmd/push-swap/ : programme push-swap
- cmd/checker/   : programme checker
