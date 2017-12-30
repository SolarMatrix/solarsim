# Phase 1: Solar System Modelling

## System Objects

| Object Name | Object Type | Object SubType        |
| ----------- | ----------- | --------------------- |
| Star        | Single      | Yellow                |
|             |             | Red Dwarf             |
|             |             | Blue Giant            |
|             |             | Blackhole             |
|             | Binary      | Yellow - Yellow       |
|             |             | Yellow - Red          |
|             |             | Red - Red             |
|             |             | Blackhole - Red       |
|             |             | Blackhole - Yellow    |
|             |             | Blackhole - Blackhole |
| Planet      | Rocky       | Metallic              |
|             |             | Rocky                 |
|             | Gas         | Hydrogen Giant        |
|             |             | Helium Giant          |
|             |             | Ringed Giant          |
|             | Asteroid    | Rocky                 |
|             |             | Metallic              |
| Artifact    | Comet       | Ice                   |
|             |             | Rock                  |
|             | Moon        | Rocky                 |
|             |             | Phase locked          |

## Calculation Attributes

1. Orbit

   Stars are locked to point (0,0,0), all other objects will require orbit calculating
   
2. Temperature

   Linked to star output, percentage of output falling on a given area of an objects curved surface

3. Rotation

   Initial rotation of all non-star objects, could further this aspect with effect of orbiting moons and sun rotation gradient releasing electrical storms

4. Greenhouse gas effect

   Modifier to temperature based on atmospheric composition

5. Hit counter

   Number of object collisions, given an objet of sufficient size affect orbit, extinction events, planetary cohesion

6. Gravity

   Required for orbit calculations as well as important modifier for surface life simulation

7. Star output

   Power driving the system dynamics, energy available on planet surfaces impossible to be greater than this output

