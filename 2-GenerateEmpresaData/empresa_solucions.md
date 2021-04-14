# Preguntes

## 1.Quants empleats hi ha en cada població?

```
db.dades.aggregate([

    {
        $group: {
            _id: "$adreça.poblacio",
            numero: { $sum: 1 }
        }
    }
])
```

## 2.On viu en "Filomenu Juanola Masvidal"?

```
db.dades.aggregate([
    { 
        $match: { "nom": "Filomenu Juanola Masvidal"}
    },
    {
        $project: {
            _id:0,
            carrer: "$adreça.carrer",
            poblacio: "$adreça.poblacio",
            numero:"$adreça.numero",
            missatge: "vegilar-lo"
        }
    }
])
```

## 3. Quanta gent parla "xinès" en cada departament?

```
db.dades.aggregate([

    {
        $match: {
            "idiomes": "xinès"
        },
    },
    {
        $group: {
            _id: "$departament",
            numero: { $sum: 1 }
        }
    }
])
```

## 4.Quantes dones i quants homes té el departament d'"Informàtica"?

```
db.dades.aggregate([
    {
        $match: { departament: "Informàtica" }
    },
    {
        $group: {
            _id: "$sexe",
            suma: { $sum: 1 }
        }
    }
])
```

## 5. Quantes dones i quines  hi ha en el departament de recursos humans?

```
db.dades.aggregate([
    {
        $match: {
            departament: "Recursos Humans",
            sexe: "D"
        }
    },
    {
        $group: {
            _id: "$sexe",
            suma: { $sum: 1 },
            noms: { $addToSet: "$nom" }
        }
    }
])
```

## 6.Quanta gent hi ha que parla cada idioma?

```
db.dades.aggregate([
    {
        $unwind: "$idiomes"
    },
    {
        $group: {
            _id: "$idiomes",
            quantitat: {$sum: 1}
        }
    }
])
```

## 7.Quin empleat parla més idiomes

```
db.dades.aggregate([
    {
        $project: { 
            _id: 1, nom: 1, idiomes: { $size: "$idiomes" }
        }
    },
    {
        $group: {
            _id: "$idiomes",
            noms: { $addToSet: "$nom" }
        }
    },
    {
        $sort: { "_id": -1 }
    },
    {
        $limit: 1
    }
])
```

## 8.Quins dels empleats només parlen "català" i "castellà"

```
db.dades.aggregate([
    {
        $match: 
        {
            "idiomes": { $all: ["català", "castellà"] , $size: 2 }
        }
    },
    {
        $project: { "nom": 1 }
    }
])
```

## 9.En quines adreces hi viuen més de 2 empleats?

```
db.dades.aggregate([
    {
        $group: { 
            _id: { $concat: ["$adreça.carrer",",", {$toString:"$adreça.numero"} ," ","$adreça.poblacio",]}, 
            count: { $sum: 1} 
        } 
    },
    { 
        $match: { count: { $gt: 2 } }  
    }
])
```

## 10.Quin és el cognom que tenen més empleats?

```
db.dades.aggregate([
    { 
       $group: { 
          _id: { 
             $arrayElemAt:[ { $split: ["$nom"," "] }, -2 ]
          }, 
          suma: { $sum:1 }
       }
    },
    { 
       $group: { 
          _id: "$suma", 
          cognoms: { $push: "$_id" } 
       } 
    },
    { $sort: {_id: -1}},
    { $limit: 1 }
])
```

## 11. Quines són les zones on hi ha menys viatjants de Vendes?

```
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
            zones: { $push: "$_id" }
        }
    },
    {
        $sort: { _id: 1 }
    },
    {
        $limit: 1
    }

])
```
