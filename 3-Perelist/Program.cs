using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using MongoDB.Bson;
using MongoDB.Driver;

namespace Perelist
{
    class Program
    {
        static void Main(string[] args)
        {


            var idiomes = new List<String> { "anglès", "francès" };
            var doc = new BsonDocument()
            .Add("nom", "Pere")
            .Add("cognoms", "Pi")
            .Add("edat", 40)
            .Add("adreça", new BsonDocument()
                .Add("carrer", "Nou")
                .Add("municipi", "Girona")
            )
            .Add("idiomes", new BsonArray(idiomes));

            var nom = doc["nom"];
            System.Console.WriteLine(nom);
            var edat = doc["edat"].AsInt32;
            System.Console.WriteLine(edat);

            System.Console.WriteLine(doc.Contains("nom"));
            System.Console.WriteLine(doc.ContainsValue("Pere"));



            var client = new MongoClient();

            using (var cursor = client.ListDatabases())
            {
                System.Console.WriteLine("--------------- DATABASES ---------- ");
                foreach (var data in cursor.ToEnumerable())
                {
                    // Retorna un BsonDocument =>  `data["name"] pinta només el nom
                    System.Console.WriteLine(data.ToString());
                }
            }

            var db = client.GetDatabase("bonyetes");
            System.Console.WriteLine("-------------- COLECTIONS ---------------- ");
            // Puc agafar els noms directament
            foreach (var col in db.ListCollectionNames().ToEnumerable())
            {
                System.Console.WriteLine(col.ToString());
            }

            var coleccio = db.GetCollection<BsonDocument>("dades");


            var frederic =
                new BsonDocument {
                    {"nom", "Frederic"},
                    {"cognom", "Pou"},
                    {"adreca", new BsonDocument
                        {
                            {"carrer", "Nou"},
                            {"numero", 99},
                            {"poblacio", "X"}
                        }
                    },
                    {"cotxes", new BsonArray
                        {
                            "Fiat", "Seat", "Mercedes"
                        }
                    },
                };

            coleccio.InsertOne(frederic);



            System.Console.WriteLine("--------------- COMPTA ALCALDES -------------------  ");
            // Compta
            var filterAlcaldes = Builders<BsonDocument>.Filter.Exists("alcalde");
            System.Console.WriteLine($"Alcaldes: {coleccio.CountDocuments(filterAlcaldes)}/{coleccio.CountDocuments(new BsonDocument())}");

            // Cerca
            var filter = Builders<BsonDocument>.Filter.Eq("nom", "Pere");
            var peres = coleccio.Find(filter);
            foreach (var pere in peres.ToList())
            {
                System.Console.WriteLine(pere.ToString());
            }
            System.Console.WriteLine("--------------- PERES -------------------  ");
            foreach (var pere in peres.ToList())
            {
                System.Console.WriteLine($"{pere["nom"]} {pere["cognom"]}");
            }


            // --- Elimina Frederic

            var filterFrederic = Builders<BsonDocument>.Filter.Eq("nom", "Frederic");
            var filterPou = Builders<BsonDocument>.Filter.Eq("cognom", "Pou");

            coleccio.DeleteMany(filterFrederic & filterPou);


            // Pobles --
            System.Console.WriteLine(" ------------- ORDENA POBLES per quantitat de gent ---------------- ");

            var grouping = new BsonDocument
            {
                    {"$group", new BsonDocument
                        {
                            { "_id", "$adreca.poblacio" },
                            { "suma", new BsonDocument {
                                        {"$sum", 1}
                                    }
                            },
                        }
                    }
            };

            var sorting = new BsonDocument
            {
                { "$sort", new BsonDocument
                    {
                      { "suma", -1 }
                    }
                }
            };

            var pipeline = new BsonDocument[] {
                grouping,
                sorting,
            };

            var results1 = coleccio.Aggregate<BsonDocument>(pipeline).ToList();

            foreach (var result1 in results1)
            {
                System.Console.WriteLine(result1.ToString());
                // System.Console.WriteLine($@"{result1["_id"]}: {result1["suma"]}");
            }


            System.Console.WriteLine(" ------------- ORDENA POBLES per quantitat de gent 2 ---------------- ");

            var results = coleccio.Aggregate()
                                  .Group(g => g["adreca.poblacio"],
                                          r => new
                                          {
                                              poble = r.Key,
                                              suma = r.Count()
                                          }
                                  )
                                  .Match(c => c.suma > 10)
                                  .SortByDescending(x => x.suma);

            foreach (var result in results.ToList())
            {
                System.Console.WriteLine(result.ToString());
            }
        }
    }
}
