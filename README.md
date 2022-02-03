# Sneakers API 

## Description

With this API you can visualize a large amount of sneakers with their informations : 

- ID 
- Brand(s)
- Colorway(s) 
- Release Date
- Retail price
- Description (history, materials, etc...)

There will be possibility to create a collection and add/delete the shoes you own. This collection can also be visualizable by everyone.

## Installation

...

## Technologies

...

## API References

#### Retrives all sneakers
```http
  GET /snkrs
```

#### Retrieves a sneakers by its ID
```http
  GET /snkrs/${id}
```

#### Retrieves a sneakers by its ID with different filters (name, brand, colour, price, etc...)
```http
  GET /snkrs/${id}?name=&brand=&cw=&price=
```

#### Create a sneakers
```http
  POST /snkrs
```

#### Create a collection
```http
  POST /collections
```

#### Retrieves all collections
```http
  GET /collections
```

#### Retrieves your collection by its ID
```http
  GET /collections/${id_collection}
```

#### Add a sneakers to your collection 
```http
  POST /collections/${id_collection}
```

#### Delete a sneakers of a collection by its ID
```http
  DELETE /collections/${id_collection}/snkrs/${id_sneakers}
```
