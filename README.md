Go Lang vacancy test
This is a task for Go programmer vacancy at Geeks.Team

As a test for your begginer Go lang skills we ask you to create a simple web-server.

Using Gin framework https://github.com/gin-gonic/gin create a web server with a handler /checkText. Handler will listen for POST request with such JSON:

```json
{
   "Site":["https://google.com","https://yahoo.com"],
   "SearchText":"Google"
}
```


Request structure:
```go
type Request struct {
    Site []string // Slice of strings: https://blog.golang.org/go-slices-usage-and-internals
    SearchText string
}
```

After getting request handler must get the body content of each website mentioned in Site variable (this is list of urls) and search in it for a SearchText. You can use default Go http client to get the websites body content.

If the requested SearchText was found on the page of any Site url, webserver must return JSON with the url of the site at which text was found.
Response example:

```json
{
    "FoundAtSite":"https://google.com"
}
```
Response structure:

```go
type Response struct {
    FoundAtSite string
}
```

If text was not found return HTTP Code 204 No Content.
Your test web-server must be provided at your Github repo. Just send as a link.

If you have any questions: max @ geeks.team.

--------------------------------------------------------------------------------------------------------------------------------------


So HTTP request should look like this:
```
POST /checkText HTTP/1.1
Host: somehost:8080
Content-Type: application/json

{
	"Site":["http://yahoo.com","google.com"],
	"SearchText":"google"
}
```
