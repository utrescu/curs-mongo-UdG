# Bonyetès - dades

## Normal

### Importar

``

### 1.Quines de les persones viuen a Bonyeta?

`db.dades.find({"adreca.poblacio": "Bonyeta"}).pretty()`

### 2. Quantes persones es diuen de cognom “Serra”?

`db.dades.count({"cognom": "Serra"})`

### 3. A quina població hi ha el carrer “Moll”?

`db.dades.find({"adreca.carrer": "Moll"}, {_id:0, "adreca.poblacio":1}).limit(1)`

### 4. Quines marques de cotxes ha tingut en Manel Serra?

`db.dades.find({nom: "Manel", cognom: "Serra"}, {_id: 0, cotxes:1})`

### 5. Quins dels habitants de la comarca només tenen un Seat i un Fiat ?

`db.dades.find({cotxes: { $all: ["Seat", "Fiat"] }})`

### 6.Quines persones no han tingut mai cotxe?

`db.dades.find({ "cotxes": { $exists: false } })`

### 7. Quins dels habitants tenen cotxe de la marca Ferrari?

`db.dades.find({cotxes: "Ferrari" })`

### 8. Quines de les persones que no han tingut mai cotxe viuen a Sant Ficus?

`db.dades.find({ "cotxes": { $exists: false }, "adreca.poblacio": "Sant Ficus" })`

### 9. Per fer-se una idea del gran que són els carrers el millor és mirar els números de les cases. Quin és el carrer més llarg? (Tip: ordena que alguna cosa queda)

`db.dades.find({}, {_id:0, "adreca.carrer":1, "adreca.poblacio":1}).sort({ "adreca.numero": -1 }).limit(1)`

## Aggregate

### 10. Quants habitants té cada poble?

`db.dades.aggregate( [ { $group: { _id: "$adreca.poblacio", habitants: { $sum: 1} } } ] )`
