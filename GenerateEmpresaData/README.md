# Preguntes

## 1.Quants empleats hi ha en cada població?

```mongo
db.dades.aggregate([
    {
        $group: { _id: "$adreça.poblacio", numero: { $sum: 1} }
    }
])
```

## 2.On viu l'Juan Morón?

```mongo
db.dades.aggregate( [
    { $match: { "nom": "Juan Moron" } },
    { $project: {
        _id:0, carrer: "$adreça.carrer",
        numero: "$adreça.numero",
        poblacio: "$adreça.poblacio" }
    }
])
```

## 3.Quantes dones i quants homes té el departament d'"Informàtica"?

```mongo
db.dades.aggregate([
    { $match: { departament: "Informàtica"  } },
    { $group:{ _id: "$sexe", suma: { $sum: 1 } }}
])
```

## 4.Quanta gent hi ha que parla cada idioma?

```mongo
db.dades.aggregate([
    {
        $unwind: "$idiomes"
    },
    {
        $group: { _id: "$idiomes", quantitat: {$sum: 1} }
    }
])
```

## 5.Quin empleat parla més idiomes

```mongo
db.dades.aggregate([
    { $project: { _id: 1, nom: 1, idiomes:{$size:"$idiomes"} } },
    { $sort: { idiomes: -1 } },
    { $limit: 1 },
    { $project: { _id:0, nom:"$nom" } }
])
```

## 6.Quins dels empleats només parlen "català" i "castellà"

```mongo
db.dades.aggregate([
    {
        "idiomes": { $size:2, $all: ["català", "castellà"] }
    }
])
```

> Sense $size dóna també els que parlen altre idiomes
> La opció "idiomes": ["català", "castellà"] només els retorna si estan en el mateix ordre

## 7.En quines adreces hi viuen més de 2 empleats?

```mongo
db.dades.aggregate([
    {$group: { _id: { $concat: ["$adreça.carrer",",", {$toString:"$adreça.numero"} ," ","$adreça.poblacio",]}, count: { $sum: 1} } },
    { $match: { count: { $gt: 2 } }  }
])
```

## 8.Quin són els cognom que tenen més empleats?

```mongo
db.dades.aggregate([
    { $group: { _id: { $arrayElemAt:[ { $split: ["$nom"," "] }, -1 ]}, suma: { $sum:1 }}},
    { $group: { _id: "$suma", cognoms: { $push: "$_id" } } },
    { $sort: {_id: -1}},
    { $limit: 1 }
])
```

## 9. Quanta gent parla "xinès" en cada departament?

```mongo
db.dades.aggregate([
    { $match: { idiomes: "xinès" } },
    { $group: { _id: "$departament", suma: { $sum: +1 }}}
])
```

## 10. Què més?

```mongo

```
