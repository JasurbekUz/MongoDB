# MongoDB
mongos
## SQL - NoSQL
```
     NoSQL    |     SQL
  ------------|-----------
   database   | database
   collection | table
```
## basic comands
* show dbs  - showing database
```
    > show dbs
    admin    0.000GB
    config   0.000GB
    example  0.000GB
    local    0.000GB
    > 
```
* use <db_name> - switched to bd <db_name>
```
    > use example
    switched to db example
    > 
```
* show collections - showing collections of <db_name>
```
    > show collections
    movie
    myCollection
    >
```
* db.dropDatabase() - delete database
```
    > db.dropDatabase()
    { "dropped" : "example", "ok" : 1 }
    > show dbs
    admin   0.000GB
    config  0.000GB
    local   0.000GB
    > 
```
* use <db_name> - create new database
```
    > use example
    switched to db example
    > 
```
```
    > show dbs
    admin   0.000GB
    config  0.000GB
    local   0.000GB
    > 
```
we can't see database 'example', there aren't any collections in daatabase 'example' 

* db.createCollection('collection_name') - create collection
```
    > use example
    switched to db example
    > db.createCollection('users')
    { "ok" : 1 }
    > 
```
```
    > show dbs
    admin    0.000GB
    config   0.000GB
    example  0.000GB
    local    0.000GB
    > 
```
* db.users.insert({data}) - writing data in collection
```
    > use example
    switched to db example
    > db.createCollection('users')
    { "ok" : 1 }
    > db.users.insert({
    ...         name: 'Jasur',
    ...         age: 22,
    ...         contacts: ['93 106 21 07', '91 106 21 07'],
    ...         relationship: {
    ...             name: 'Firdavs',
    ...             age: 2
    ...         }
    ...     })
    WriteResult({ "nInserted" : 1 })
    > 
```
* db.users.insertMany([{data}, {data}, ...]) - writing many datas in collection
if operation is completed successfull, appear written below:
```
    {
        "acknowledged": true,
        "insertedIds": [
            ObjectId('1qwdfty67uyhtrer56tyuh'),
            ObjectId('1qwdfty67uyhtrer56tyuh'),
            ObjectId('1qwdfty67uyhtrer56tyuh')
        ]
    }
```
ex:
```
    > db.users.insertMany([
    ...     {
    ...         name: 'ALi',
    ...         age: 23,
    ...         contacts: ['93 106 22 07', '90 106 21 07'],
    ...         relationship: {
    ...             name: 'Ziyod',
    ...             age: 2
    ...         }
    ...     },
    ...     {
    ...         name: 'Akbar',
    ...         age: 24,
    ...         contacts: ['93 102 21 07', '91 126 21 07'],
    ...         relationship: {
    ...             name: 'Vali',
    ...             age: 3
    ...         }
    ...     },
    ...     {
    ...         name: 'Aziz',
    ...         age: 26,
    ...         contacts: ['93 106 21 67', '91 606 21 07'],
    ...         relationship: {
    ...             name: 'Islom',
    ...             age: 8
    ...         }
    ...     }
    ... ])
    {
        "acknowledged" : true,
        "insertedIds" : [
            ObjectId("61fa329278743b72db5c95f2"),
            ObjectId("61fa329278743b72db5c95f3"),
            ObjectId("61fa329278743b72db5c95f4")
        ]
    }

    ```
    * db.<collection_name>.find() - showning datas of collection
    ```
        > use example
        switched to db example
        > db.users.find()
        { "_id" : ObjectId("61fa10baf966a6fec741a499"), "name" : "Jasur", "age" : 22, "contacts" : [ "93 106 21 07", "91 106 21 07" ], "relationship" : { "name" : "Firdavs", "age" : 2 } }
        >
```
* db.users.find().pretty() - showning datas of collection
```
    > use example 
    switched to db example
    > db.users.find().pretty()
    {
        "_id" : ObjectId("61fa10baf966a6fec741a499"),
        "name" : "Jasur",
        "age" : 22,
        "contacts" : [
            "93 106 21 07",
            "91 106 21 07"
        ],
        "relationship" : {
            "name" : "Firdavs",
            "age" : 2
        }
    }
    {
        "_id" : ObjectId("61fa329278743b72db5c95f2"),
        "name" : "ALi",
        "age" : 23,
        "contacts" : [
            "93 106 22 07",
            "90 106 21 07"
        ],
        "relationship" : {
            "name" : "Ziyod",
            "age" : 2
        }
    }
    {
        "_id" : ObjectId("61fa329278743b72db5c95f3"),
        "name" : "Akbar",
        "age" : 24,
        "contacts" : [
            "93 102 21 07",
            "91 126 21 07"
        ],
        "relationship" : {
            "name" : "Vali",
            "age" : 3
        }
    }
    {
        "_id" : ObjectId("61fa329278743b72db5c95f4"),
        "name" : "Aziz",
        "age" : 26,
        "contacts" : [
            "93 106 21 67",
            "91 606 21 07"
        ],
        "relationship" : {
            "name" : "Islom",
            "age" : 8
        }
    }
    > 
```
#### searching 
* method: **find({key:value})**
*db.<collection_name>.find({key:'value'})*

```
    > db.users.find({name:'Jasur'})
    { "_id" : ObjectId("61fa10baf966a6fec741a499"), "name" : "Jasur", "age" : 22, "contacts" : [ "93 106 21 07", "91 106 21 07" ], "relationship" : { "name" : "Firdavs", "age" : 2 } }
```
```
    > db.users.find({age:22}).pretty()
    {
        "_id" : ObjectId("61fa10baf966a6fec741a499"),
        "name" : "Jasur",
        "age" : 22,
        "contacts" : [
            "93 106 21 07",
            "91 106 21 07"
        ],
        "relationship" : {
            "name" : "Firdavs",
            "age" : 2
        }
    }
    > 
```
***
## useful methods
#### method: *.sort({key:value})* 
*db.users.find().sort({key:value})*
```
    > db.users.find().sort({name:1}).pretty()
    {
        "_id" : ObjectId("61fa329278743b72db5c95f2"),
        "name" : "ALi",
        "age" : 23,
        "contacts" : [
            "93 106 22 07",
            "90 106 21 07"
        ],
        "relationship" : {
            "name" : "Ziyod",
            "age" : 2
        }
    }
    {
        "_id" : ObjectId("61fa329278743b72db5c95f3"),
        "name" : "Akbar",
        "age" : 24,
        "contacts" : [
            "93 102 21 07",
            "91 126 21 07"
        ],
        "relationship" : {
            "name" : "Vali",
            "age" : 3
        }
    }
    {
        "_id" : ObjectId("61fa329278743b72db5c95f4"),
        "name" : "Aziz",
        "age" : 26,
        "contacts" : [
            "93 106 21 67",
            "91 606 21 07"
        ],
        "relationship" : {
            "name" : "Islom",
            "age" : 8
        }
    }
    {
        "_id" : ObjectId("61fa10baf966a6fec741a499"),
        "name" : "Jasur",
        "age" : 22,
        "contacts" : [
            "93 106 21 07",
            "91 106 21 07"
        ],
        "relationship" : {
            "name" : "Firdavs",
            "age" : 2
        }
    }
    > 
```
#### method: *.count()*

*db.<collection_name>.find().count()*
```
    > db.users.find().count()
    4
    > 
```
```
    > db.users.find({age:22}).count()
    1
    > 
```
#### method: *.limit(arg)*
*db.users.find().limit(2).pretty()*
```
    > db.users.find().limit(2).pretty()
    {
        "_id" : ObjectId("61fa10baf966a6fec741a499"),
        "name" : "Jasur",
        "age" : 22,
        "contacts" : [
            "93 106 21 07",
            "91 106 21 07"
        ],
        "relationship" : {
            "name" : "Firdavs",
            "age" : 2
        }
    }
    {
        "_id" : ObjectId("61fa329278743b72db5c95f2"),
        "name" : "ALi",
        "age" : 23,
        "contacts" : [
            "93 106 22 07",
            "90 106 21 07"
        ],
        "relationship" : {
            "name" : "Ziyod",
            "age" : 2
        }
    }
    > 
```
#### JS functions
* **.forEach(function(arg){ func body })**
```
    > db.users.find().forEach(
    ...     function(doc) { 
    ...         print('user name:' + doc.name)
    ...     }
    ... )
    user name:Jasur
    user name:ALi
    user name:Akbar
    user name:Aziz
    > 
```
***
## update and delete
### updating
#### method:  *.update({1},{2},{3})*

* 1-for search
* 2-for new datas
* 3-
*db.users.update({id:''12wer5te4trehgfdfg},{name:'name'},{upsert: true})*
```
    > db.users.update({
    ...     name: 'Jasur'
    ... },
    ... {
    ...     name: 'Jasurbek',
    ...     age: 23,
    ... },{
    ...     upsert: true
    ... })
    WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
    > 
```
```
"nMatched" : 1 - qidirilgan obektdan 1-ta topildi
"nUpserted" : 0 - yangi yaratilmadi, chunki qidirlgan obektdan bor
"nModified" : 1 - ozgartitildi
```
if don't have searched data in collection:
```
    > db.users.update({
    ...     name: 'Jasur'
    ... },
    ... {
    ...     name: 'Jasurbek',
    ...     age: 23,
    ... },{
    ...     upsert: true
    ... })
    WriteResult({
        "nMatched" : 0,
        "nUpserted" : 1,
        "nModified" : 0,
        "_id" : ObjectId("61fa4be3b5c1a271e431a638")
    })
    > 
```
**'update' method maintened which we changed datas and deleted other datas:**
```
    > db.users.find().pretty()
    {
        "_id" : ObjectId("61fa10baf966a6fec741a499"),
        "name" : "Jasurbek",
        "age" : 23
    }
    {
        "_id" : ObjectId("61fa329278743b72db5c95f2"),
        "name" : "ALi",
        "age" : 23,
        "contacts" : [
            "93 106 22 07",
            "90 106 21 07"
        ],
        "relationship" : {
            "name" : "Ziyod",
            "age" : 2
        }
    }
    {
        "_id" : ObjectId("61fa329278743b72db5c95f3"),
        "name" : "Akbar",
        "age" : 24,
        "contacts" : [
            "93 102 21 07",
            "91 126 21 07"
        ],
        "relationship" : {
            "name" : "Vali",
            "age" : 3
        }
    }
    {
        "_id" : ObjectId("61fa329278743b72db5c95f4"),
        "name" : "Aziz",
        "age" : 26,
        "contacts" : [
            "93 106 21 67",
            "91 606 21 07"
        ],
        "relationship" : {
            "name" : "Islom",
            "age" : 8
        }
    }
    > 
```
**to prevent this, we use '$set' *aggregate* function:**
```
    > db.users.update({
    ...     name: 'Akbar'
    ...     },
    ...     {
    ...         $set: {
    ...             age: 23
    ...         }
    ...     }
    ... )
    WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
    > db.users.find({name:'Akbar'}).pretty()
    {
        "_id" : ObjectId("61fa329278743b72db5c95f3"),
        "name" : "Akbar",
        "age" : 23,
        "contacts" : [
            "93 102 21 07",
            "91 126 21 07"
        ],
        "relationship" : {
            "name" : "Vali",
            "age" : 3
        }
    }
    > 
```
#### some aggregate functions:
##### $inc - increment
```
    > db.users.update(
    ...     {
    ...         name: 'Akbar'
    ...     },
    ...     {
    ...         $inc:{
    ...             age: 7
    ...         }
    ...     }
    ... )
    WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
    > db.users.find({name:'Akbar'}).pretty()
    {
        "_id" : ObjectId("61fa329278743b72db5c95f3"),
        "name" : "Akbar",
        "age" : 30,
        "contacts" : [
            "93 102 21 07",
            "91 126 21 07"
        ],
        "relationship" : {
            "name" : "Vali",
            "age" : 3
        }
    }
    >  
```
##### $rename
```
    > db.users.update(
    ...     {
    ...         name:'Akbar'
    ...     },
    ...     {
    ...         $rename: {
    ...             name: 'fullname'
    ...         }
    ...     }
    ... )
    WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
    > db.users.find({name:'Akbar'}).pretty()
    > db.users.find({fullname:'Akbar'}).pretty()
    {
        "_id" : ObjectId("61fa329278743b72db5c95f3"),
        "age" : 30,
        "contacts" : [
            "93 102 21 07",
            "91 126 21 07"
        ],
        "relationship" : {
            "name" : "Vali",
            "age" : 3
        },
        "fullname" : "Akbar"
    }
    > 
```
### deleteing
#### method: ***.remove()***
```
    > db.users.remove({fullname:'Akbar'})
    WriteResult({ "nRemoved" : 1 })
    > db.users.find({name:'Akbar'}).pretty()
    > 
```
## searching
**step 1 createing index:** ***db.users.createIndex({key:'value'})***
*db.users.createIndex({name:'text'})*

```
    > db.users.createIndex({name:'text'})
    {
        "numIndexesBefore" : 2,
        "numIndexesAfter" : 2,
        "note" : "all indexes already exist",
        "ok" : 1
    }
    >
```
**step 2 createing index:** ***write query***
*db.users.find({ $text:{ $search: 'Ali'}})*
```
    > db.users.find({
    ...     $text:{
    ...         $search: 'Ali'
    ...     }
    ... }).pretty()
    {
        "_id" : ObjectId("61fa73d7e538d83911dace5e"),
        "name" : "ALi",
        "age" : 23,
        "contacts" : [
            "93 106 22 07",
            "90 106 21 07"
        ],
        "relationship" : {
            "name" : "Ziyod",
            "age" : 2
        }
    }
    > 
```
### $gt and $lt
#### $gt >
***db.users.find({age: {$gt:25}}).pretty()***
```
    > db.users.find({age: {$gt:25}}).pretty()
    {
        "_id" : ObjectId("61fa73d7e538d83911dace62"),
        "name" : "Aziz",
        "age" : 26,
        "contacts" : [
            "93 106 21 67",
            "91 606 21 07"
        ],
        "relationship" : {
            "name" : "Islom",
            "age" : 8
        }
    }
    > 
```
#### $lt - <
***db.users.find({age: {$lt:25}}).pretty()***
```
    > db.users.find({age: {$lt:25}}).pretty()
    {
        "_id" : ObjectId("61fa73d7e538d83911dace5e"),
        "name" : "ALi",
        "age" : 23,
        "contacts" : [
            "93 106 22 07",
            "90 106 21 07"
        ],
        "relationship" : {
            "name" : "Ziyod",
            "age" : 2
        }
    }
    {
        "_id" : ObjectId("61fa73d7e538d83911dace5f"),
        "name" : "Akbar",
        "age" : 24,
        "contacts" : [
            "93 102 21 07",
            "91 126 21 07"
        ],
        "relationship" : {
            "name" : "Vali",
            "age" : 3
        }
    }
    {
        "_id" : ObjectId("61fa73d7e538d83911dace60"),
        "name" : "Ilyos",
        "age" : 24,
        "contacts" : [
            "93 102 21 07",
            "91 126 21 07"
        ],
        "relationship" : {
            "name" : "Vali",
            "age" : 3
        }
    }
    {
        "_id" : ObjectId("61fa73d7e538d83911dace61"),
        "name" : "Ilhom",
        "age" : 24,
        "contacts" : [
            "93 102 21 07",
            "91 126 21 07"
        ],
        "relationship" : {
            "name" : "Vali",
            "age" : 3
        }
    }
    > 
```

#### $gte - >=, $lte - <=
