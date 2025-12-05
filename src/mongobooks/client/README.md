#HowTO

List books:
`curl -i http://localhost:8080/books`
`curl -is http://localhost:8080/books | grep Hostname`

Add book:
`curl -X POST http://localhost:8080/books      -H "Content-Type: application/json"      -d '{"title":"Go in Action","authors":["William Kennedy", "Brian Ketelsen"],"publish_date":"2015-11-26"}'`
`curl -X POST http://localhost:8080/books      -H "Content-Type: application/json"      -d '{"title":"Kubernetes in Action","authors":["Marko Luksha"],"publish_date":"2017-10-26"}' `