using System.Collections.Generic;
using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace CotxesApi.Models
{

    public class Address
    {
        public string carrer { get; set; }
        public int numero { get; set; }
        public string poblacio { get; set; }
    }

    public class Cotxe
    {
        [BsonId]
        [BsonRepresentation(BsonType.ObjectId)]
        public string Id { get; set; }

        public string nom { get; set; }
        public string cognom { get; set; }

        public Address adreca { get; set; }
        public List<string> cotxes { get; set; }

        public bool alcalde { get; set; }
    }
}