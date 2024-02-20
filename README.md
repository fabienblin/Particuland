# Partiucland
A particle simulation based on a rulset of asymetric interactions.
![alt](https://github.com/fabienblin/Particuland/blob/main/particuland.gif)

## Usage
If a particular configuration pleases you, keep in mind you can reproduce a scenario by using the same seed.
```
./particuland --seed 1234
```

## How it works
Particles belong to a predefined set of species and each species has a certain set of attributes. The first and most obvious attribute of a species will be it's color. Then you will need to define it's interaction radius with other particles. The species will also have a maximum velocity.

Each particle belongs to a particle species and is defined in space by it's position and direction and speed.

Interactions between species will be defined by a table initialized with 0 values, meaning there is no attraction or repulsion between particles. When the interaction is positive, particles attract and vice versa, when the interaction value is negative they repel.
The interactions are asymetric wich means particle A can be attracted to particle B, but B can flee from A. This can be compared to a pray/predator situation.

## Basic example
A row is the reference particle's interaction to other species.

Lets imagine we have 3 species of particles, Red, Green and Blue. We can define their interactions such as Red is the apex predator and Blue is at the bottom of the food chain. In that situation we can represent the interaction table as such :

|       | Red   | Green | Blue  |
|-------|-------|-------|-------|
| Red   |   0   |   1   |   0   |
| Green |  -1   |   0   |   1   |
| Blue  |   0   |  -1   |   0   |

In  this example
- Red will be attracted to Green
- Green will run from Red and try to catch Blue
- Blue will run away from Green


## Optimization
The computational cost of moving the particles get increasingly hight as the total number of particles increase, mainly because of the KNN implementation that is very basic. Instead of making a quad tree I decided to parallelize the function for each particle. Having a quad tree would certainly increase efficiency buy also the complexity of the algorithm. This would need some comparative limit testing.

## Define interactions
As a developer you can change species interactions using the SetInteraction function. By default all is random.

Salut la Gophosphère,


Au menu aujourd'hui, un assortiment de particules sur fond noir, sauce newtonienne. Je vous présente Particuland https://lnkd.in/eGVZ-dS7, un monde où tout est particule, chacune avec sa personnalité, son charme et son histoire.

Pour expliquer un peu la magie 🤩 , vous avez des espèces de particules, avec chacune leur couleur et surtout leur forces d’interactions avec les autres espèces, ensuite chaque particule a un vecteur de vélocité. Tout ça est donné à la loi universelle de la gravitation de Newton et BIM ! (ça fait des chocapics)

Dans l'exemple en dessous vous voyez 3 espèces, rouge ❤️ , verte 💚 et bleue 💙 , ce qui fait 9 interactions. Avec des valeurs au hasard je suis tombé sur cette configuration où les particules bleues semblent former un triangle avec du vert . Et puis le rouge se tient loin en formant une sorte de barrière le plus loin possible.

Hésitez pas à tester par vous-même et me dire si vous tombez sur des choses intéressantes !

#golang #generativeart #ebitengine #newton #particlelife