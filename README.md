# estiam_golang


16/11/2023
Exercice : Gestion de Données avec Fichier en Go

Contexte: Vous avez implémenté un dictionnaire en utilisant une map dans un exercice précédent. Maintenant, vous allez modifier votre implémentation pour stocker les données dans un fichier plutôt que dans une map.

Tâches:	
Modifiez le package dictionary pour utiliser un fichier au lieu d'une map pour stocker les entrées du dictionnaire.
Utilisez les méthodes Add, Get, Remove, et List du package dictionary dans main.go.
Assurez-vous que les opérations sur les données fonctionnent correctement après ces modifications.

Instructions:
Modifiez le code dans dictionary.go pour utiliser un fichier (au format de votre choix) au lieu d'une map pour stocker les entrées du dictionnaire.
Adaptez les méthodes Add, Get, Remove, et List en conséquence.
Testez les opérations dans main.go pour garantir que tout fonctionne correctement.

Consignes supplémentaires:	
Ajoutez et validez (git add, git commit) vos modifications pour chaque étape.
Poussez (git push) régulièrement vos modifications sur GitHub.
Assurez-vous que votre programme fonctionne correctement avec les données stockées dans un fichier.


**Exercice Avancé :**
*Gestion Concurrente avec Channels en Go*
Contexte : Vous avez maintenant une version fonctionnelle de votre dictionnaire qui stocke les entrées dans un fichier. Pour rendre votre programme plus performant, vous allez introduire la gestion concurrente en utilisant des channels pour les opérations d'ajout et de suppression.

**Tâches:**
1. Ajoutez des channels pour les opérations d'ajout (Add) et de suppression (Remove) dans le package dictionary.
2. Utilisez la concurrence pour gérer simultanément les opérations d'ajout et de suppression dans main.go.
3. Assurez-vous que les opérations sur les données fonctionnent correctement en utilisant la gestion concurrente.


**Instructions :**
* Modifiez dictionary.go pour ajouter des channels pour les opérations d'ajout et de suppression.
* Adaptez les méthodes Add et Remove pour utiliser des channels.
* Dans main.go, utilisez la concurrence pour effectuer simultanément des opérations d'ajout et de suppression.
* Testez attentivement pour vous assurer que la gestion concurrente fonctionne correctement.


**Consignes supplémentaires :**
1. Ajoutez et validez (git add, git commit) vos modifications à chaque étape.
2. Poussez (git push) régulièrement vos modifications sur GitHub.
3. Documentez vos changements dans les messages de commit.