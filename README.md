# shortenlinks


## Инкремент 1
### Задание по треку «Сервис сокращения URL»
Чтобы написать сервис, который будет сжимать длинные URL до нескольких символов, для начала вам нужно разработать сервер.

#### Сервер должен быть доступен по адресу http://localhost:8080 и предоставлять два эндпоинта:

1. Эндпоинт с методом POST и путём /. Сервер принимает в теле запроса строку URL как text/plain и возвращает ответ с кодом 201 и сокращённым URL как text/plain.
2. Эндпоинт с методом GET и путём /{id}, где id — идентификатор сокращённого URL (например, /EwHXdJfB). В случае успешной обработки запроса сервер возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.


#### Endpoints
1. Эндпоинт с методом POST и путём /. Сервер принимает в теле запроса строку URL как text/plain и возвращает ответ с кодом 201 и сокращённым URL как text/plain.
Пример запроса к серверу:
```sh
POST / HTTP/1.1
Host: localhost:8080
Content-Type: text/plain

https://practicum.yandex.ru/
```

Пример ответа от сервера:
```sh
HTTP/1.1 201 Created
Content-Type: text/plain
Content-Length: 30

http://localhost:8080/EwHXdJfB 
```

Пример того же самого, но с curl:
```sh
curl http://localhost:8080/ --include --header "Conternt-Type: text/plain" --request "POST" --data "https://practicum.yandex.ru/"
```
Ответ на запрос что выше:
```sh
HTTP/1.1 201 Created
Content-Type: text/plain; charset=utf-8
Date: Thu, 05 Jun 2025 08:27:08 GMT
Content-Length: 30

http://localhost:8080/BmdwhpPk
```


2. Эндпоинт с методом GET и путём /{id}, где id — идентификатор сокращённого URL (например, /EwHXdJfB). В случае успешной обработки запроса сервер возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.

Пример запроса к серверу:
```sh
GET /EwHXdJfB HTTP/1.1
Host: localhost:8080
Content-Type: text/plain 
```
Пример ответа от сервера:
```sh
HTTP/1.1 307 Temporary Redirect
Location: https://practicum.yandex.ru/ 
```

Пример того же самого, но с curl:
```sh
curl http://localhost:8080/BmdwhpPk --include --header "Conternt-Type: text/plain" --request "GET"
```
Ответ на запрос что выше:
```sh
HTTP/1.1 307 Temporary Redirect
Content-Type: text/plain; charset=utf-8
Location: https://practicum.yandex.ru/
Date: Thu, 05 Jun 2025 08:27:12 GMT
Content-Length: 0
```

На любой некорректный запрос сервер должен возвращать ответ с кодом 400.
https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml

Пример с curl:
curl http://localhost:8080/Bmd --include --header "Conternt-Typ
e: text/plain" --request "GET"

HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
Date: Thu, 05 Jun 2025 08:28:18 GMT
Content-Length: 37

Некорректный запрос
