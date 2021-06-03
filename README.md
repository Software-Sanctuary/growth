# Growth Framework

## Introduction

Growth is a framework that helps a gopher to define APIs in a declarative way.
Growth removes writing boilerplate code and lets the gopher to focus on business logic and data manipulation.

## Example of RESTful Api with Growth

```golang
package main

import (
    "github.com/softwaresanctuary/growth"
)

func main() {
    growth
        .CreateAnAPI()
            .ThatHasRoutes(
                growth.ApiRoute(growth.API.HttpMethod.GET, "/", welcome),
                growth.ApiRoute(growth.API.HttpMethod.GET, "/about", about)
            ).Run("8080")
}

```

## Libraries

- Seed (API) []
- Bark (Containerization) []
- Sprout (DatabaseConnection and Pool) []
- Branches (Query Generator) []
- Leaf (Logging)
- Crown (AccessControl)
- Rings (Testing)
  
