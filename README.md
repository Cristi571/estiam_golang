# estiam_golang project

## Manipulation de Map en Go

### Contexte:
Vous travaillez sur une application de gestion de mots et de définitions en utilisant Go. Le code actuel utilise une map en mémoire pour stocker les entrées.

### Tâches:
* Créez une fonction main dans le fichier principal (main.go) pour réaliser les tâches suivantes:
* Utilisez la méthode Add pour ajouter quelques mots et définitions à la map.
* Utilisez la méthode Get pour afficher la définition d'un mot spécifique.
* Utilisez la méthode Remove pour supprimer un mot de la map.
* Appelez la méthode List pour obtenir la liste triée des mots et de leurs définitions.
* Exécutez le programme et vérifiez si les opérations sur la map sont correctement implémentées.



## Création d'un Package Dictionary en Go

### Contexte: 
Vous poursuivez le développement de votre application de gestion de mots et de définitions. Pour rendre le code plus modulaire, vous allez créer un package dictionary contenant les méthodes Add, Get, Remove, et List.

### Tâches:
- [ ] Créez un nouveau fichier nommé dictionary.go dans le répertoire dictionary de votre projet.
- [ ] Déplacez les méthodes Add, Get, Remove, et List dans le package dictionary.
- [ ] Importez et utilisez ce package dans le fichier main.go pour effectuer les opérations sur la map.

### Instructions:
1. Créez un fichier dictionary.go dans le répertoire dictionary.
2. Copiez les méthodes Add, Get, Remove, et List de main.go vers dictionary.go.
3. Modifiez le fichier main.go pour importer le package dictionary et utilisez les méthodes du package pour réaliser les opérations sur la map.
4. Exécutez le programme et assurez-vous que tout fonctionne comme prévu.



#### 16/11/2023
## Gestion de Données avec Fichier en Go

### Contexte:
Vous avez implémenté un dictionnaire en utilisant une map dans un exercice précédent. Maintenant, vous allez modifier votre implémentation pour stocker les données dans un fichier plutôt que dans une map.

### Tâches:
- [ ] Modifiez le package dictionary pour utiliser un fichier au lieu d'une map pour stocker les entrées du dictionnaire.
- [ ] Utilisez les méthodes Add, Get, Remove, et List du package dictionary dans main.go.
- [ ] Assurez-vous que les opérations sur les données fonctionnent correctement après ces modifications.

### Instructions:
1. Modifiez le code dans dictionary.go pour utiliser un fichier (au format de votre choix) au lieu d'une map pour stocker les entrées du dictionnaire.
2. Adaptez les méthodes Add, Get, Remove, et List en conséquence.
3. Testez les opérations dans main.go pour garantir que tout fonctionne correctement.

**Consignes supplémentaires:**
1. Ajoutez et validez (git add, git commit) vos modifications pour chaque étape.
2. Poussez (git push) régulièrement vos modifications sur GitHub.
3. Assurez-vous que votre programme fonctionne correctement avec les données stockées dans un fichier.


#### Exercice Avancé :
## Gestion Concurrente avec Channels en Go
### Contexte : 
Vous avez maintenant une version fonctionnelle de votre dictionnaire qui stocke les entrées dans un fichier. Pour rendre votre programme plus performant, vous allez introduire la gestion concurrente en utilisant des channels pour les opérations d'ajout et de suppression.

### Tâches:
- [ ] Ajoutez des channels pour les opérations d'ajout (Add) et de suppression (Remove) dans le package dictionary.
- [ ] Utilisez la concurrence pour gérer simultanément les opérations d'ajout et de suppression dans main.go.
- [ ] Assurez-vous que les opérations sur les données fonctionnent correctement en utilisant la gestion concurrente.


### Instructions :
1. Modifiez dictionary.go pour ajouter des channels pour les opérations d'ajout et de suppression.
2. Adaptez les méthodes Add et Remove pour utiliser des channels.
3. Dans main.go, utilisez la concurrence pour effectuer simultanément des opérations d'ajout et de suppression.
4. Testez attentivement pour vous assurer que la gestion concurrente fonctionne correctement.


**Consignes supplémentaires :**
1. Ajoutez et validez (git add, git commit) vos modifications à chaque étape.
2. Poussez (git push) régulièrement vos modifications sur GitHub.
3. Documentez vos changements dans les messages de commit.





#### 22/11/2023
## API REST avec Gorilla Mux

### Contexte :
Pour étendre votre dictionnaire pour le rendre accessible via une API REST. Pour cela, vous utiliserez Gorilla Mux, un puissant routeur HTTP pour Go.

### Tâches :
- [ ] Intégrez Gorilla Mux dans votre projet 
   installez-le si ce n'est pas déjà fait : 
   `go get -u github.com/gorilla/mux.`
- [ ] Créez trois nouvelles routes :
1. Une pour ajouter une entrée au dictionnaire (POST)
2. Une pour récupérer une définition par mot (GET)
3. Une pour supprimer une entrée par mot (DELETE)
- [ ] Mettez à jour main.go pour utiliser ces nouvelles routes.

### Instructions :
* Installez Gorilla Mux avec go get -u github.com/gorilla/mux.
* Intégrez Gorilla Mux dans votre projet.
* Ajoutez les routes nécessaires dans package route.
* Adaptez main.go pour utiliser les nouvelles routes.






## Middleware Logging avec Gorilla Mux*
### Contexte :
Maintenant que votre API REST fonctionne correctement, vous souhaitez ajouter une fonctionnalité de journalisation (logging) pour enregistrer chaque requête entrante dans une sorte de fichier journal.
### Tâches :
1. Ajoutez un middleware qui enregistre chaque requête dans un fichier journal (un fichier texte simple).
2. Personnalisez le format du message de journalisation pour inclure des informations telles que l'heure, la méthode HTTP et le chemin.
3. Testez votre API en effectuant plusieurs requêtes et vérifiez que les journaux sont correctement enregistrés.
### Instructions :
* Ajoutez un middleware pour la journalisation dans un package middleware.
* Personnalisez le format du message de journalisation.
* Testez votre API en faisant des requêtes et vérifiez que les journaux sont enregistrés correctement.


---
## Ajout d'Authentification avec Gorilla Mux
### Contexte :
Pour renforcer la sécurité de votre API, vous décidez d'ajouter une fonctionnalité d'authentification basique. Les utilisateurs devront fournir un jeton (token) d'authentification dans l'en-tête de leurs requêtes.

### Tâches :
1. Ajoutez un middleware d'authentification qui vérifie la présence d'un jeton dans l'en-tête de la requête.
2. Si le jeton est valide, autorisez l'accès à l'API. Sinon, renvoyez une réponse d'erreur non autorisée.
3. Testez votre API en incluant un jeton d'authentification valide et invalide dans les requêtes.
### Instructions :
* Ajoutez un middleware pour l'authentification.
* Personnalisez la logique d'authentification pour vérifier la présence et la validité du jeton.
* Testez votre API en incluant des jetons valides et invalides dans les requêtes.

### Consignes supplémentaires :
- Ajoutez et validez (git add, git commit) vos modifications à chaque étape.
- Poussez (git push) régulièrement vos modifications sur GitHub.
- Documentez vos changements dans les messages de commit.



#### 11/01/2024
## Suite de l'Exercice : Gestion des Erreurs et Validation des Données
### Contexte : 
Votre API est désormais plus sécurisée grâce à l'authentification, mais pour garantir un fonctionnement robuste, vous souhaitez améliorer la gestion des erreurs et ajouter une validation des données.
### Tâches :
#### A. Gestion des Erreurs :
1. Ajoutez une gestion des erreurs pour tous les endpoints de votre API.
2. Assurez-vous que les erreurs sont correctement loguées et que des réponses d'erreur appropriées sont renvoyées aux clients.

#### B. Validation des Données :
1. Ajoutez une validation des données pour le point de terminaison d'ajout d'une entrée au dictionnaire (méthode POST).
2. Assurez-vous que le mot et la définition respectent certaines règles (par exemple, longueur minimale et maximale).

### Instructions :
* Ajoutez une gestion des erreurs pour tous les endpoints de votre API dans dictionary.go.
* Implémentez une validation des données pour le point de terminaison d'ajout d'une entrée au dictionnaire.
* Testez votre API en incluant des données invalides pour vérifier que les erreurs sont correctement gérées et que la validation des données fonctionne comme prévu.

#### 11/01/2024
## Tests Unitaires pour l'API REST avec Gorilla Mux
### Contexte : Vous avez développé votre API REST avec Gorilla Mux, ajouté la gestion des erreurs, et renforcé la sécurité avec l'authentification et la vérification des rôles. Maintenant, il est temps de créer des tests unitaires approfondis pour assurer la fiabilité de votre application.
### Tâches :
#### A. Tests unitaires pour les fonctions du package Dictionary :
1. Créez des tests unitaires pour chaque fonction du package Dictionary (ajout, récupération, suppression, liste).
2. Assurez-vous de tester différents scénarios, y compris les cas de succès et d'échec.
#### B. Tests unitaires pour les middlewares :
1. Écrivez des tests pour les middlewares d'authentification et de vérification des rôles.
2. Testez différentes situations, telles que des jetons valides et invalides, ainsi que des utilisateurs avec des rôles corrects ou incorrects.
#### C. Tests d'intégration pour les routes :
1. Mettez en place des tests d'intégration pour toutes les routes de votre API.
2. Utilisez un serveur de test pour simuler les requêtes HTTP et assurez-vous que toutes les fonctionnalités sont correctement testées.
#### DCouverture de code :
Utilisez des outils tels que go test avec le paramètre -cover pour mesurer la couverture de code de vos tests.
Identifiez les parties de votre code qui ne sont pas couvertes par les tests et ajoutez des tests pour ces zones.

---
#### 18/01/2024
## Persistance des Données avec une Base de Données
### Contexte :
Vous avez décidé d'ajouter la persistance des données à votre API REST en utilisant une base de données. Dans cette partie, vous explorerez l'utilisation d'une base de données clé-valeur (key-value store) pour stocker les entrées du dictionnaire.

### Tâches :
#### Choix de la Base de Données Clé-Valeur :
Sélectionnez une base de données clé-valeur appropriée pour votre application. Considérez des solutions comme Redis, et configurez-la pour votre application.

#### Configuration de la Base de Données :
* Intégrez les détails de connexion à la base de données clé-valeur dans votre application.
* Adaptez vos fonctions du package Dictionary pour interagir avec la base de données clé-valeur au lieu de la mémoire.

#### Tests avec la Base de Données :
* Mettez à jour vos tests unitaires pour inclure des scénarios où les fonctions interagissent avec la base de données clé-valeur.
* Assurez-vous que vos tests d'intégration fonctionnent correctement avec la nouvelle configuration.