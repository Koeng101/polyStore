from peewee import *
from playhouse.db_url import connect
from playhouse.postgres_ext import *
import peeweedbevolve
import os

db = connect(os.environ.get('DATABASE_URL'))

class BaseModel(Model):
    class Meta:
        database = db

class Seqhash(BaseModel):
    seqhash = TextField(primary_key=True)
    circular = BooleanField(default=False)
    doubleStranded = BooleanField(default=True)
    sequenceType = TextField(choices=["DNA", "RNA", "PROTEIN"])
    sequence = TextField()

class SeqhashLink(BaseModel):
    child_seqhash = ForeignKeyField(Seqhash)
    parent_seqhash = ForeignKeyField(Seqhash)

class Genbank(BaseModel):
    accession = TextField(primary_key=True)
    seqhash = ForeignKeyField(Seqhash)
    binary_json = BinaryJSONField()
    json_hash = TextField()

class Uniprot(BaseModel):
    accession = TextField(primary_key=True)
    seqhash = ForeignKeyField(Seqhash)
    database = TextField(choices=["Swiss-Prot", "TrEMBL"])
    
models = [Seqhash, SeqhashLink, Genbank, Uniprot]

if __name__ == "__main__":
    db.evolve(models)
