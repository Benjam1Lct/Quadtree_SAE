# Project Quadtree

Code source initial pour le projet d'introduction au développement (R1.01) et de SAÉ implémentation d'un besoin client (SAE1.01), année 2023-2024.

## Introduction

Le projet Quadtree est le projet final et SAÉ du cours d'initiation au développement de la première année du BUT informatique. L'objectif de ce projet est de matérialiser les déplacements d'un personnage sur un terrain potentiellement infini, généré au fur et à mesure de son exploration. Le terrain est stocké en mémoire à l'aide d'une structure de données appelée quadtree, ou arbre quaternaire.

## Dates d'évaluation

- La partie 5 du projet doit être rendue au plus tard le vendredi 15 décembre 2023 (semaine 50).
- La partie 6 (SAÉ implémentation d'un besoin client) doit être rendue au plus tard le mercredi 10 janvier 2023 (semaine 2).
- Évaluation du code d'un autre groupe : vendredi 12 janvier 2023 (semaine 2).
- Contrôle sur table en semaine 2 (date à déterminer).

## Critères d'évaluation

Les critères d'évaluation comprennent le bon fonctionnement du code, le choix et la qualité des cas de tests, la qualité du code (noms de variables et fonctions, organisation des fichiers, documentation), et l'efficacité du code. Pour la SAÉ, le nombre et la complexité des extensions réalisées seront également évalués.

## Organisation des Sources du Projet

Le projet est organisé en différents paquets, chacun ayant une responsabilité spécifique dans le fonctionnement du jeu.

### Paquets principaux

- **assets**: Contient les images utilisées dans le projet.
- **camera**: Gestion de la caméra.
- **character**: Gestion du personnage.
- **cmd**: Construction de l'exécutable et fichier de configuration.
- **configuration**: Lecture des fichiers de configuration.
- **floor**: Gestion du terrain.
- **game**: Implantation de l'interface ebiten.Game.
- **quadtree**: Bibliothèque pour les quadtree.

## Utilisation de la Bibliothèque Ebitengine

La bibliothèque Ebitengine simplifie le développement de jeux en 2D en offrant une API pour l'affichage d'images, la détection des interactions utilisateur et le cadencement des calculs. Veuillez consulter la [documentation d'Ebitengine](https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2) pour plus de détails.

## Informations sur les Quadtree

Un quadtree est une structure de données arborescente où chaque nœud (non feuille) a exactement quatre enfants. Dans le cadre de ce projet, les quadtrees sont utilisés pour représenter le terrain de manière compacte.

## Principe des Quadtree

Un quadtree divise le terrain en quatre zones, chaque zone étant représentée par un nœud de l'arbre. Chaque nœud non feuille a quatre enfants, représentant les quatre sous-zones de la zone parente. Les zones qui contiennent un seul type de terrain sont représentées par des feuilles qui stockent l'information sur la taille de la zone et le type de terrain.

## Extensions ajoutées
- Génération aléatoire de terrain
- Interdiction de marcher sur l’eau
- Téléporteurs
- Caméra bloquée aux bords du terrain (CameraMode 2)
- Texture ameliorées
- Animation
- Camera cinématique (CameraMode 3)
- Sauvegarde de la generation aleatoire du terrain

