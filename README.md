# Fin

![alt text](https://avatars1.githubusercontent.com/u/4650108?s=400&v=4)   

A library which build to be a Gin best friend. Fin comes with the abilities to
- Override the default log of Gin and make it becomes Lightnet's log format.
- Manage the correlation id of Composite service and Microservice
- Make a http request with the GET, POST and PUT methods. If the correlation id is in the context, fin will automatically add it to the header of the request.

## Fin in Lightnet project

```go
// main.go in Composite Service
package main

import (
	...
	"gitlab.com/lightnet-thailand/fin"
)

...

func main() {
	ginEngine := gin.New()

	ginEngine.Use(gin.Recovery())
	ginEngine.Use(fin.LogWithCorrelationIdMiddleware(gin.DefaultWriter, []string{"/health"}))
	ginEngine.Use(fin.CompositeCorrelationIdMiddleware())
	ginEngine.Use(fin.FromGinContextToContextMiddleware())

	// Init Boredom Log
	boredom.Init()
    
	// Init usecase and repository
	...

}
```

```go
// /app/modules/stellar_account/repository/microservice_repository.go in Composite Service
package repository

import ...

func (repository *repository) SugarDaddySign(ctx context.Context, xdrOps []string) (*string, error) {
	errContextMsg := "fail to sign sugar daddy xdr ops by requesting to stellar-account-service " + env.StellarServiceUrl

	preSignRequestJSON := struct {
		XdrOps []string `json:"xdrOps"`
	}{xdrOps}
	requestBody, _ := json.Marshal(preSignRequestJSON)

	url := fmt.Sprintf("%s/v1/sugar-daddies.sign", env.StellarServiceUrl)
	resp, _ := fin.Post(ctx, url, constants.ContentTypeJSON, bytes.NewBuffer(requestBody))

	responseBodyData, err := utils.VeloResponseParser(resp, errContextMsg)
	if err != nil {
		return nil, err
	}

	responseModel, err := new(models.SugarDaddyXdrResponse).Parse(responseBodyData)
	if err != nil {
		return nil, errors.Wrap(err, errContextMsg)
	}

	return responseModel.SignedXdr, nil
}
```

```go
// main.go in Microservice Service
package main

import (
	...
	"gitlab.com/lightnet-thailand/fin"
)

...

func main() {
	ginEngine := gin.New()

	ginEngine.Use(gin.Recovery())
	ginEngine.Use(fin.LogWithCorrelationIdMiddleware(gin.DefaultWriter, []string{"/health"}))
	ginEngine.Use(fin.MicroServiceCorrelationIdMiddleware())
	ginEngine.Use(fin.FromGinContextToContextMiddleware())

	// Init Boredom Log
	boredom.Init()
    
	// Init usecase and repository
	...

}
```

## Available API
```go
// converting from gin context to context by placing the gin.Context to a very specific key of context
fin.FromGinContextToContext(gc)
// converting from context to gin context by getting the gin.Context from a very specific key of context
fin.FromContextToGinContext(ctx)
// a very specific key of context which hold gin.Context
fin.GinContextKey
// a key of correlation id in the header
fin.CorrelationIdHeaderKey
```