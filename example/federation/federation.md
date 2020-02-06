---
title: 'Using Apollo federation gqlgen'
description: How federate many services into a single graph using Apollo
linkTitle: Apollo Federation
menu: { main: { parent: 'recipes' } }
---

In this quick guide we are going to implement the example [Apollo Federation](https://www.apollographql.com/docs/apollo-server/federation/introduction/)
server in gqlgen. You can find the finished result in the [examples directory](/example/federation).


### Create the servers

For each server to be federated we will create a new gqlgen project.

```bash
go run github.com/99designs/gqlgen
```

Update the schema to reflect the federated example
```graphql
type Review {
  body: String
  author: User @provides(fields: "username")
  product: Product
}

extend type User @key(fields: "id") {
  id: ID! @external
  reviews: [Review]
}

extend type Product @key(fields: "upc") {
  upc: String! @external
  reviews: [Review]
}
```

and regenerate
```bash
go run github.com/99designs/gqlgen
```
