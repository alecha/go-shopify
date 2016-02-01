go-shopify
==========

[![GoDoc](https://godoc.org/github.com/rapito/go-shopify/shopify?status.svg)](https://godoc.org/github.com/rapito/go-shopify/shopify)  [![baby-gopher](https://raw.github.com/drnic/babygopher-site/gh-pages/images/babygopher-logo-small.png)](http://www.babygopher.org)

Simple API made with **go** to make **CRUD** request to your **Shopify Store**.

Installation
------------
```
go get github.com/rapito/go-shopify
```

Dependencies
------------
```
go get github.com/parnurzeal/gorequest
```

Test dependencies
------------
```
go get github.com/bitly/go-simplejson
go get github.com/bmizerany/assert
```

Use it
----------


- GET Requests

```go
import "fmt"
import "github.com/rapito/go-shopify/shopify"

shop := shopify.New(storeDomain,apiKey,pass)
result, _ := shop.Get("products")

fmt.Println(string(result))
```

- Check out the *examples* folder for simple usage.
- Read some of the tests at *shopify_test.go* for complete CRUD examples.

- There are some built-in wrappers than you can use, you are welcome to add new ones:
  - `GetOrders()`
  - `GetOrder(orderID)`
  - `GetOrderTransactions(orderID)`
  - `GetOrderTransactionsCount(orderID)`
  - `GetOrdersCount()`
  - `GetProducts()`
  - `GetProduct(productID)`
  - `GetProductImages(productID)`
  - `GetProductVariants(productID)`

Contribution
------------

 - You may fork this library and modify it as you please.
 - You can make a pull request and I will be happy to check it out and merge it.
 - If you find a bug, create an issue and I will do my best to fix it (someday).

Original Work
-------------

While I was looking for something cool to do with this new language im learning
(Go, obviously), I ran into [hammond-bones'](https://github.com/hammond-bones/) **go-shopify**
library. Which inspired me to start creating this one.

- Fork it at: [go-shopify](https://github.com/hammond-bones/go-shopify)
