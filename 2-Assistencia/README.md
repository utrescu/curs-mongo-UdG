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

## 2. Quins departaments són els 5 treballadors que cobren menys?

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

## 2. Mitjana de sous de cada departament?

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

## 3. Qui són els caps de departament?

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

## 5. Quines dues setmanes han estat les que han tingut més baixes?

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

## 4. Quina assistència tenim cada dia?

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

## 5. Quins dies hi ha algú del departament d'Informàtica?

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
