## Contexte : la prod est en feu

Nous sommes le **1er janvier 2026**, il est **1 h 34** du matin et vous êtes d’astreinte. Les couloirs sont vides ; vous lancez la nouvelle saison de votre série préférée.

Mais avant de partir, un dev freelance (TJM **800 € / jour**) a laissé quelques **boulettes** qui ont été **poussées en prod** à la va-vite par Pierre, côté **DevOps**, juste avant son départ en congé.

Tout allait bien jusqu’à une alerte **Kibana** : un conteneur **Docker** redémarre en boucle. La prod est en feu : **à vous d’agir**.

**Objectif :** On n’est pas là pour faire du propre — on verra ça cette semaine. **Déboguez le service** qui **fuit**. Essayez de mettre en place ce que nous avons vu dans les TPs précédents.

**Astuce :** le package [`context`](https://pkg.go.dev/context) est votre meilleur allié.

**Conclusion :** Avec la concurrence, il faut gérer **toute la chaîne des sous-services** pour éviter les **fuites de goroutines** et de mémoire (propagation du `context`, annulation propre, timeouts).
