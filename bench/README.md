# Bench

The higher the faster
| Name                                      | Time per request [ms]        |
| -------------                             |:-------------:|
| [chi](https://github.com/go-chi/chi)| 0.410|
| [fiber](https://github.com/gofiber/fiber) | 0.414| 
| [echo](https://github.com/labstack/echo)| 0.416 | 
| [fasthttp](https://github.com/valyala/fasthttp)| 0.419| 
| [Gin](https://github.com/gin-gonic/gin)   | 0.420          |
| [kratos](https://github.com/go-kratos/kratos)| 0.421 | 
| [beego](https://github.com/beego/beego)   | 0.463     | 
| [kit](https://github.com/go-kit/kit) | 0.455     | 
| [httprouter](https://github.com/julienschmidt/httprouter)| 0.472|  
| [go-zero](https://github.com/zeromicro/go-zero) | 0.495| 

More info is in [bench.txt](https://github.com/forChin/go-frameworks-pros-and-cons/blob/main/bench/fiber-crud/bench.txt) file in folders of implementations.

Endpoints were tested by [ApacheBench](https://httpd.apache.org/docs/2.4/programs/ab.html):
```
ab -n 1000 -p user_json.txt -T application/json http://localhost:8080/
```
Content of user_json.txt:
```
{
  "name":"chingiz",
  "age":22
}
```
Endpoint logic for all implementations is similar (current example is written in [chi](https://github.com/go-chi/chi)):
```
r := chi.NewRouter()
r.Post("/", func(w http.ResponseWriter, r *http.Request) {
  var u model.User
  if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
    log.Fatal(err)
  }
  err = repo.SaveUser(u)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("-- saved", u.Name, u.Age)
})
log.Fatal(http.ListenAndServe(":8080", r))
```
