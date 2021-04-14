# Preguntes

## 0. Importar les dades

## 1. Localitza les dades del nom "JOAN" que hi ha a les comarques de Girona?

`db.noms.find({nom:"JOAN"})`

## 2. Extreu les dades dels noms de dona

`db.noms.find({sexe:"D"})`

## 3. Hi ha algun nom que el portin exactament 666 persones en alguna comarca?

`db.noms.find( { "comarques.quantitat": 666 })`

## 4. Quants noms diferents hi ha en l'Alt Empordà ?

`db.noms.count( { "comarques.comarca": { $in: ["Alt Empordà" ] }})`

`db.noms.count( { "comarques.comarca": "Alt Empordà" })`

## 5. Dades dels noms que porten entre 100 i 101 persones?

`db.noms.find({ total: { $gte: 100 , $lt: 101 } })`

## 6. Quantes "ELNA" hi ha?

`db.noms.find({nom:"ELNA"}, {total:1})`

## 7.Quins són els noms de dona que porten més de 5000 persones?

`db.noms.find({ sexe:"D", total: { $gt: 5000 }}, { nom: 1 })`

## 8. En quantes comarques hi ha gent que es diu `PASCUAL`?

`db.noms.find( { nom: "PASCUAL" }, { num: { $size: "$comarques" }} )`

## 9. Quin és el nom de dona més usat?

`db.noms.find({sexe:"D"}, {_id:0, "nom":1}).sort({"total": -1}).limit(1)`

## 10. Què hi ha més, ALEXANDRE o ALEXANDRA?

`db.noms.find({nom:/^ALEXANDR[EA]$/}, {_id:0, nom:1}).sort({total: -1}).limit(1)`

`db.noms.find({ $or: [ {nom:"ALEXANDRA"},{nom:"ALEXANDRE"} ] }, {_id:0, nom:1}).sort({total: -1}).limit(1)`

## 11. Quins noms contenen "JOAN" en el nom?

`db.noms.find({nom:/\s*JOAN\s*/}, {nom:1}).count()`
