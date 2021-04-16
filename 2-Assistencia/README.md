# Preguntes

## 1. Quants treballadors té cada departament?

```mongo
db.presencia.aggregate([
    {
        $group: {
            _id: "$departament.nom",
            count: { $sum: 1 }
        }
    }

])
```

## 2. Quant es gasten en sou cada mes?

```mongo
db.presencia.aggregate([
    {
        $group:
        {
            _id:0, sous: { $sum: "$departament.sou" }
        }
    }
])
```

## 3. De quins departaments són els 5 treballadors que cobren menys?

```mongo
db.presencia.aggregate([
    {
        $sort: { "departament.sou": 1 }
    },
    {
        $limit: 5
    },
    {
        $project: {
            _id:0,
            Departament: "$departament.nom"
        }
    }
])
```

## 4. Quina és la mitjana de sous de cada departament?

```mongo
db.presencia.aggregate([
    {
        $group: {
            _id: "$departament.nom",
            sous: { $avg: "$departament.sou" }
        }
    }
])
```

## 5. Qui són els empleats que tenen algun càrrec?

```mongo
db.presencia.aggregate([
    {
        $match: {
            "departament.càrrec": { $exists: true }
        }
    },
    {
        $project: {
            _id: 0,
            nom: "$nom",
            cognoms: "$cognoms",
            departament: "$departament.nom"
        }
    }
])
```

## 6. Quins empleats han tingut "Infecció CoVID" de cada departament?

```mongo
db.presencia.aggregate([
    {
        $match: { "setmanes.justificacio": "Infecció CoVID" }
    },
    {
        $project: {
            _id: 0,
            nom: { $concat: [ "$nom", " ", "$cognoms" ] },
            departament: "$departament.nom"
        }
    },
    {
        $group: {
            _id: "$departament",
            noms: { $addToSet:  "$nom" },
        }
    }
])
```

## 7. Quina assistència tenim cada dia de la setmana?

```mongo
db.presencia.aggregate([
    {
      $project: { "dies": "$setmanes.dies", "_id":0 }
    },
    {
        $unwind: "$dies"
    },
    {
        $unwind: "$dies"
    },
    {
        $group: {
            _id:"$dies",
            personal: { $sum: 1} }
        }
    ])
```

## 8. Quins dies hi ha algú del departament d'Informàtica?

```mongo
db.presencia.aggregate([
    {
        $match: {
            "departament.nom": "Informàtica"
        }
    },
    {
        $project: {
            "dies": "$setmanes.dies" ,
            "_id":0 }
    },
    {
        $unwind: "$dies"
    },
    {
        $unwind: "$dies"
    },
    {
        $group: {
            _id: 0,
            dies: { $addToSet: "$dies"}
        }
    }
])
```

## 9. Quines dues setmanes han estat les que han tingut més baixes?

```mongo
db.presencia.aggregate([
    {
        $unwind: "$setmanes"
    },
    {
        $project: {
            numero: "$setmanes.número",
            faltes : { $subtract: [2, { $size: "$setmanes.dies" }] }
        }
    },
    {
        $group: {
            _id: "$numero",
            suma: { $sum: "$faltes" }
        }
    },
    { $sort: { "suma": -1 } },
    { $limit: 2 }
])
```
