## TL;DR
- deploy `kubectl apply -f deploy/`
- add book `curl -X POST http://<node_ip>:30030/books      -H "Content-Type: application/json"      -d '{"title":"Go in Action","authors":["William Kennedy", "Brian Ketelsen"],"publish_date":"2015-11-26"}'`
- list books `curl -s http://<node_ip>:30030/books`
- `curl -is http://<node_ip>:30030/books | grep Hostname`  
В заголовках добавлено поле `Hostname` для отслеживания имени хоста, с которого пришло сообщение

## Task
_Формат сдачи задания:_ git-репозиторий с проектом.  
_Требования к репозиторию:_ все манифесты проекта должны быть в директории репозитория deploy, чтобы команда `kubectl apply -f deploy/` втаскивала сразу весь проект.  
_Платформа:_ Kubernetes версии 1.34, подключение к интернету есть, доступ к Docker Hub есть, образы оттуда можно использовать. IngressController никакой не установлен, доступен StorageClass с идентификатором csi-driver-lvm-linear

#### Normal (на троечку) 3️⃣
Любое веб-приложение, которое использует базу данных для хранения данных. Базу данных желательно брать из кластерных (MongoDB, CockroachDB), но если иначе никак, то можно любую из классических (MySQL, PostgreSQL). Опубликовать порт веб-приложения с помощью NodePort, без использования Ingress. Для хранения данных использовать Volume (emptyDir).

## Description

Приложение `mongobooks` реализует простмотр и добавление книг в список, хранящийся в БД MongoDB

## Structure
 - `/src` - source code of mongobooks project
 - `/deploy` - k8s yaml manifests

## Deploy

`kubectl apply -f deploy/`

## Mongobooks App
### Build
One could build Mongobooks App with `./src/mongobooks/docker-compose.yml` file

### Test
#### List books:
- `curl -is http://localhost:8080/books`  
При запуске сиситема пуста, необходимо проинициализировать её добавив значение через добавление значения
- `curl -is http://localhost:8080/books | grep Hostname`  
В заголовках добавлено поле `Hostname` для отслеживания имени хоста, с которого пришло сообщение

#### Add book:
- `curl -X POST http://localhost:8080/books      -H "Content-Type: application/json"      -d '{"title":"Go in Action","authors":["William Kennedy", "Brian Ketelsen"],"publish_date":"2015-11-26"}'`
- `curl -X POST http://localhost:8080/books      -H "Content-Type: application/json"      -d '{"title":"Kubernetes in Action","authors":["Marko Luksha"],"publish_date":"2017-10-26"}' `