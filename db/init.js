let dir = pwd()

print(dir)
print(listFiles())

load("/docker-entrypoint-initdb.d/sample_data.js")
print(recipes)

db.myCollection.insertMany(recipes);
