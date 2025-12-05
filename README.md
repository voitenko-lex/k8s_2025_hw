# HowTO

### List books:
- `curl -is http://localhost:8080/books`  
При запуске сиситема пуста, необходимо проинициализировать её добавив значение через добавление значения
- `curl -is http://localhost:8080/books | grep Hostname`  
В заголовках добавлено поле `Hostname` для отслеживания имени хоста, с которого пришло сообщение

### Add book:
- `curl -X POST http://localhost:8080/books      -H "Content-Type: application/json"      -d '{"title":"Go in Action","authors":["William Kennedy", "Brian Ketelsen"],"publish_date":"2015-11-26"}'`
- `curl -X POST http://localhost:8080/books      -H "Content-Type: application/json"      -d '{"title":"Kubernetes in Action","authors":["Marko Luksha"],"publish_date":"2017-10-26"}' `