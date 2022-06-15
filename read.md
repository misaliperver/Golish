# Golish Fiber Server

# Handbook
- https://github.com/gofiber/recipes/blob/master/auth-jwt/handler/product.go
- https://docs.gofiber.io/guide/grouping
- https://go.kaanksc.com/boeluem-2/fonksiyonlar

## Pre-Development
First of all go to https://docs.gofiber.io/ link for fiber docs.

run command on shell.
```
  go get github.com/gofiber/fiber/v2
```
And then, go.mod file change with module name with github.com/misaliperver/golish

You should move main.go file to under the **app** folder.

And than .replit file change run command pathing with app/main.go

## Let's Use The Fiber

Create a folder structer;
- **app**
- constants
- **controllers**
- **helpers**
- middlewares
- models
- **routes**
- tmp

You can change them all as your needs.

## Write 3 important file
- routes/api.go
- controllers/api.controller.go
- helpers/responseTransform.go

We need to make a relational map for request with controller. Controller is a handler function, you can think as bussiness layer for your api.

Helpers are more different other functionality. Helpers provide simple func (example: transform data, text replacement, etc.) for our bussiness layer.

Package names and uses list;
- controllers ---> routes/api.go
- helpers ------> controllers/api.controller.go
- routes -------> app/main.go

## Lets Write File Upload Api
- Add uuid module: go mod "github.com/google/uuid"
- We use tmp folder: uploaded file serve file
- Create common.go file into constants folder: BaseUrl settings

```
package constants

var ServerUrl string = "https://golish.misaliperver.repl.co"
```

We should add main.go static file settings ;
```
  app.Static("/images", "./tmp")
```

Routes add to upload route.

We should write new controller into api.controller.go file.

Below code will get skill for use uploaded file.
```
  file, err := c.FormFile("image")
```
Create new name for file and get .ext name. Save file to ./tmp folder.

Return img adress and other info.

## Lets Write Resize to Fly Api

## Lets Write Reaction Api

## Lets Write Stroy Api