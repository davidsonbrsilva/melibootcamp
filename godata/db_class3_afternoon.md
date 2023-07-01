# Para começar

1. Quantas coleções tem o banco de dados?
   1 coleção (restaurantes)

2. Quantos documentos em cada coleção? Quanto ocupa cada coleção?
  25.359 documentos na coleção restaurantes. 4,37MB.

3. Quantos índices em cada coleção? Quanto espaço os índices de cada coleção ocupam?
  Tem 1 índice na coleção de restaurantes. 909,3KB.

4. Traga um documento de exemplo de cada coleção. db.collection.find(...).pretty() nos dá um formato mais legível.
  ```json
{
  "direccion": {
    "edificio": "625",
    "coord": [
      -73.990494,
      40.7569545
    ],
    "calle": "8 Avenue",
    "codigo_postal": "10018"
  },
  "barrio": "Manhattan",
  "tipo_cocina": "American",
  "grados": [
    {
      "date": {
        "$date": "2014-06-09T00:00:00.000Z"
      },
      "grado": "A",
      "puntaje": 12
    },
    {
      "date": {
        "$date": "2014-01-10T00:00:00.000Z"
      },
      "grado": "A",
      "puntaje": 9
    },
    {
      "date": {
        "$date": "2012-12-07T00:00:00.000Z"
      },
      "grado": "A",
      "puntaje": 4
    },
    {
      "date": {
        "$date": "2011-12-13T00:00:00.000Z"
      },
      "grado": "A",
      "puntaje": 9
    },
    {
      "date": {
        "$date": "2011-09-09T00:00:00.000Z"
      },
      "grado": "A",
      "puntaje": 13
    }
  ],
  "nombre": "Cafe Metro",
  "restaurante_id": "40363298"
}
  ```

5. Para cada coleção, liste os campos de nível raiz (ignore campos em documentos
aninhados) e seus tipos de dados.
restaurantes: direction, barrio, tipo_cocina, grados, nombre, restaurante_id.

# Exercício 1

1. Devolver restaurante_id, nombre, barrio e tipo_cocina mas excluindo _id para um documento (o primeiro):
```
db.restaurantes.find({},{restaurante_id: 1, nombre: 1, barrio: 1, tipo_cocina: 1, _id: 0}).limit(1)
```

2. Devolver restaurante_id, nombre, barrio e tipo_cocina para os primeiros 3 restaurantes que contenham 'Bake' em alguma parte do seu nome.
```
db.restaurantes.find({nombre: /Bake/},{restaurante_id: 1, nombre: 1, barrio: 1, tipo_cocina: 1, _id: 0}).limit(3)
```

3. Contar os restaurantes de comida (tipo_cocina) china (Chinese) o tailandesa (Thai) do bairro (bairro) Bronx.
```
db.restaurantes.find({tipo_cocina: { $in: ['Chinese', 'Thai']} }).count()
```

# Exercício 2

1. Traga 3 restaurantes que receberam pelo menos uma classificação de grau 'A' com pontuação maior que 20. A mesma classificação deve atender às duas condições simultaneamente.
```
db.restaurantes.find({"grados.grado": "A", "grados.puntaje": { $gt: 20 }}).limit(3)

db.restaurantes.find({"grados": { $elemMatch: { grado: "A", puntaje: { $gt: 20 } } }}).limit(3)
```

2. Quantos documentos estão faltando coordenadas geográficas? Em outras palavras, verifique se o tamanho de direccion.coord é 0 e contar.
```
db.restaurantes.find({ "direccion.coord": { $size: 0 } }).count()
```

3. Devolver nombre, barrio, tipo_cocina y grados para os 3 primeiros restaurantes; de cada documento apenas a última avaliação.
```
db.restaurantes.find({},{ nombre: 1, barrio: 1, tipo_cocina: 1, grados: { $slice: -1 }, _id: 0 })
```

# Exercício 3

1. Quais são os 3 principais tipos de culinária (cuisine) que podemos encontrar nos dados? Googlear "mongodb group by field, count it and sort it".
```
db.restaurantes.aggregate([ { $project: { tipo_cocina: 1 } }, { $group: { _id: "$tipo_cocina", total: { $sum: 1 } } }, { $sort: { total: -1 } }, { $limit: 3 } ])

db.restaurantes.aggregate([ { $group: { _id: "$tipo_cocina", total: { $count: {} } } }, { $sort: { total: -1 } }, { $limit: 3 } ])

// renamed _id to cozinha at the end
db.restaurantes.aggregate([ { $group: { _id: "$tipo_cocina", total: { $count: {} } } }, { $sort: { total: -1 } }, { $limit: 3 }, { $project: { _id: 0, cozinha: "$_id", total: 1 } } ])
```

2. Quais são os bairros mais desenvolvidos gastronomicamente? Calcular a média ($avg) a pontuação (grades.score) por bairro; considerando restaurantes com mais de três avaliações; Classifique os bairros com a melhor pontuação.
```
db.restaurantes.aggregate([{ $match: { $expr: { $gt: [{ $size: "$grados" }, 3] } } }, { $unwind: '$grados' },{ $group: { _id: '$barrio', score: { $avg: '$grados.puntaje' } } }, { $sort: { score: -1 } } ])
```

3. Uma pessoa querendo comer está na longitude -73.93414657 e na latitude 40.82302903, Quais opções você tem em um raio de 500 metros?
```
db.restaurantes.find({ 'direccion.coord': { $nearSphere: { $geometry: { type: "Point", coordinates: [-73.93414657, 40.82302903] }, $maxDistance: 500 } } })
```
