/*
quadtree est le paquet qui fournit la structure de données pour les arbres
quaternaires ainsi que les fonctions et méthodes pour manipuler ces arbres.

# FUNCTIONS

func SearchContent(X, Y int, n *node) int

# TYPES

	type Quadtree struct {
	        // Has unexported fields.
	}

	Quadtree est la structure de données pour les arbres quaternaires. Les
	champs non exportés sont :
	  - width, height : la taille en cases de la zone représentée par l'arbre.
	  - root : le nœud qui est la racine de l'arbre.

func MakeFromArray(floorContent [][]int) (q Quadtree)

	MakeFromArray construit un quadtree représentant un terrain étant donné un
	tableau représentant ce terrain.

func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int)

	GetContent remplit le tableau contentHolder (qui représente un terrain dont
	la case le plus en haut à gauche a pour coordonnées (topLeftX, topLeftY)) à
	partir du quadtree q.
*/
package quadtree
