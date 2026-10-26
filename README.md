# go-reloaded

Projet Ynov B1 en Go.

Le programme lit un fichier texte, le corrige et ecrit le resultat dans un autre fichier.

Installer : 

Il faut avoir Go installe (https://go.dev/dl/).

git clone https://github.com/theosugier/go-reloaded.git
cd go-reloaded

Lancer : 

go run . sample.txt result.txt

Ce qu'il fait: 

- `(hex)` et `(bin)` : convertit le mot d'avant en decimal
- `(up)`, `(low)`, `(cap)` : majuscules, minuscules ou premiere lettre en majuscule
- `(up, 2)` : pareil sur les 2 mots d'avant
- la ponctuation `. , ! ? : ;` est collee au mot d'avant
- les `' '` sont colles aux mots qu'ils entourent
- `a` devient `an` devant une voyelle ou un h

Exemples: 


1E (hex) files were added
-> 30 files were added

it was the age of foolishness (cap, 6)
-> It Was The Age Of Foolishness

I was sitting over there ,and then BAMM !!
-> I was sitting over there, and then BAMM!!

There it was. A amazing rock!
-> There it was. An amazing rock!

