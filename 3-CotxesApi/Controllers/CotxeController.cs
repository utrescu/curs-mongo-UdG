using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using CotxesApi.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using MongoDB.Bson;
using MongoDB.Driver;

namespace CotxeApi.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class CotxeController : ControllerBase
    {
        private readonly IMongoCollection<Cotxe> _cotxes;
        private readonly ILogger<CotxeController> _logger;

        public CotxeController(ILogger<CotxeController> logger)
        {
            var client = new MongoClient();
            var database = client.GetDatabase("bonyetes");

            _cotxes = database.GetCollection<Cotxe>("dades");
            _logger = logger;
        }

        [HttpGet]
        public IEnumerable<Cotxe> Get()
        {
            return _cotxes.Find<Cotxe>(new BsonDocument()).ToList();
        }

        [HttpGet("alcaldes")]
        public IEnumerable<Cotxe> GetAlcaldes()
        {
            return _cotxes.Find<Cotxe>(b => b.alcalde == true).ToList();
        }

        [HttpGet("pobles")]
        public IEnumerable<String> GetPobles()
        {
            return _cotxes.AsQueryable<Cotxe>().Select(c => c.adreca.poblacio).Distinct();
        }

        [HttpGet("nom/{nom}")]
        public IEnumerable<Cotxe> GetPeres(string nom)
        {
            var filter = Builders<Cotxe>.Filter.Eq("nom", nom);
            return _cotxes.Find(filter).ToList();
        }


        [HttpGet("gentperpoble")]
        public IEnumerable<object> GetNumPersones(string nom)
        {
            return _cotxes.Aggregate()
                                .Group(c => c.adreca.poblacio,
                                       g => new
                                       {
                                           Persones = g.Key,
                                           Total = g.Count()
                                       }
                                )
                                .ToList();
        }

        //
        // He hagut de fer una classe helper perquè no aconsegueixo
        // fer que el unwind+group em detecti la clau d'agregació
        //
        public class CotxeOut
        {

            public string nom { get; set; }
            public string cognom { get; set; }

            public Address adreca { get; set; }
            public string cotxes { get; set; }

            public bool alcalde { get; set; }
        }

        [HttpGet("sumacotxes")]
        public IEnumerable<object> GetNumCotxes(string nom)
        {
            return _cotxes.Aggregate()
                                .Unwind<Cotxe, CotxeOut>(c => c.cotxes)
                                .Group(c => c.cotxes,
                                       g => new
                                       {
                                           Marca = g.Key,
                                           Count = g.Count()
                                       }
                                )
                                .ToList();
        }



    }
}
