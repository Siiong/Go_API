# Sneakers API 

## Description

With this API you can visualize a large amount of sneakers with their informations : 

- ID 
- Brand(s)
- Colour(s) 
- Release Date
- Retail price
- Description (history, materials, etc...)

There will be possibility to create a collection and add/delete the shoes you own. This collection can also be visualizable by everyone.

## Installation

...

## Technologies

...

## API References

#### GET all items 
```http
  GET /snkrs
```

#### GET a specific item (by ID)
```http
  GET /snkrs/${id}
```

#### GET your collection
```http
  GET /collections
```

#### ADD item to your collection 
```http
  POST /collections
```

#### DELETE an item (by ID) 
```http
  DELETE /collections/${id_collection}/snkrs/${id_sneakers}
```
