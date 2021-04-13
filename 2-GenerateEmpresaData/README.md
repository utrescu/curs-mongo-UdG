# Preguntes

## 1.Quants empleats hi ha en cada població?

```mongo
db.dades.aggregate([
    {
        $group: { _id: "$adreça.poblacio", numero: { $sum: 1} }
    }
])
```

## 2.On viu en "Filomenu Juanola Masvidal"?

```mongo
db.dades.aggregate( [
    { $match: { "nom": "Filomenu Juanola Masvidal" } },
    { $project: {
        _id:0, carrer: "$adreça.carrer",
        numero: "$adreça.numero",
        poblacio: "$adreça.poblacio" }
    }
])
```

## 3. Quanta gent parla "xinès" en cada departament?

```mongo
db.dades.aggregate([
    { $match: { idiomes: "xinès" } },
    { $group: { _id: "$departament", suma: { $sum: +1 }}}
])
```

## 4.Quantes dones i quants homes té el departament d'"Informàtica"?

```mongo
db.dades.aggregate([
    { $match: { departament: "Informàtica"  } },
    { $group:{ _id: "$sexe", suma: { $sum: 1 } }}
])
```

## 5. Quines dones hi ha en el departament de recursos humans?

```mongo
db.dades.aggregate([
    { $match: { departament: "Recursos Humans", sexe: "D"  } },
    { $group:{ _id: "$sexe", gent: { $addToSet: "$nom" } } }
])
```

## 6.Quanta gent hi ha que parla cada idioma?

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

## 7.Quin empleat parla més idiomes

```mongo
db.dades.aggregate([
    { $project: { _id: 1, nom: 1, idiomes:{$size:"$idiomes"} } },
    { $sort: { idiomes: -1 } },
    { $limit: 1 },
    { $project: { _id:0, nom:"$nom", idiomes: 1 } }
])
```

## 8.Quins dels empleats només parlen "català" i "castellà"

```mongo
db.dades.aggregate([
    {
        $match:
        {
            "idiomes": { $size:2, $all: ["català", "castellà"] }
        }
    },
    {
        $project:
        {
            "nom": 1
        }
    }
])
```

> Sense $size dóna també els que parlen altre idiomes
> La opció "idiomes": ["català", "castellà"] només els retorna si estan en el mateix ordre

## 9.En quines adreces hi viuen més de 2 empleats?

```mongo
db.dades.aggregate([
    {$group: { _id: { $concat: ["$adreça.carrer",",", {$toString:"$adreça.numero"} ," ","$adreça.poblacio",]}, count: { $sum: 1} } },
    { $match: { count: { $gt: 2 } }  }
])
```

## 10.Quin és el cognom que tenen més empleats?

```mongo
db.dades.aggregate([
    { $group: { _id: { $arrayElemAt:[ { $split: ["$nom"," "] }, -2 ]}, suma: { $sum:1 }}},
    { $group: { _id: "$suma", cognoms: { $push: "$_id" } } },
    { $sort: {_id: -1}},
    { $limit: 1 }
])
```

## 11. Quines són les zones on hi ha menys viatjants de Vendes?

```mongo
db.dades.aggregate([
    {
        $match: { "departament": "Vendes" }
    },
    {
        $unwind: "$zones"
    },
    {
        $group: {
            _id: "$zones.nom",
            viatjants: { $sum: 1 }
        }
    },
    {
        $group: {
            _id: "$viatjants",
            zones: { $addToSet: "$_id"  }
        }
    },
    {
        $sort: { "_id": 1 }
    },
    {
        $limit: 1
    }
])
```
