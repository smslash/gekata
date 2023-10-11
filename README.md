# Gekata - Техническое задание

## Запуск

```console
go run ./cmd/main.go
```

Введите все ссылку с новой строки. После ввода всех ссылок введите пустую строку или нажмите на ```Enter```

## Входные данные

На стандартный вход программе подаётся список URL, по одному в каждой строке, которая терминируется символом новой строки. Обработка (выполнение запросов) должна быть распараллелена по числу присутствующих в системе вычислительных ядер.

## Выходные данные

Необходимо для каждого URL определить код ответа HTTP на запрос GET, размер документа и время между отправкой запроса и получением ответа. Печать результатов должен производиться на стандартный выход в формате CSV, пример: **http://example.com/;200;1560;120ms**.

Возникающие ошибки при обращении к сайту, нужно отображать на стандартный выход.

По завершении программы либо через ```Ctrl+C``` либо при закрытии входа, необходимо вывести статистику работы программы в виде количества отработанных запросов вs разрезе параллельного исполнения, в виде *<порядковый номер параллельного потока запросов>:<число запросов>*.


## Примеры URL

- [https://habr.com/ru/post/215117/](https://habr.com/ru/post/215117/)
- [https://ru.wikipedia.org/wiki/HTTP](https://ru.wikipedia.org/wiki/HTTP)
- [https://developer.mozilla.org/ru/docs/Web/HTTP](https://developer.mozilla.org/ru/docs/Web/HTTP)
- [https://ru.bmstu.wiki/HTTP_(Hypertext_Transfer_Protocol)](https://ru.bmstu.wiki/HTTP_(Hypertext_Transfer_Protocol))
- [https://proselyte.net/tutorials/http-tutorial/introduction/](https://proselyte.net/tutorials/http-tutorial/introduction/)
- [https://www.speedcheck.org/ru/wiki/http/](https://www.speedcheck.org/ru/wiki/http/)
- [http://pki.gov.kz/index.php/ru/ncalayer](http://pki.gov.kz/index.php/ru/ncalayer)
- [http://www.edu.gov.kz/](http://www.edu.gov.kz/)
- [http://adilet.zan.kz/rus](http://adilet.zan.kz/rus)
- [https://wiki.merionet.ru/servernye-resheniya/3/protocol-http/](https://wiki.merionet.ru/servernye-resheniya/3/protocol-http/)
- [https://www.opennet.ru/docs/RUS/http/](https://www.opennet.ru/docs/RUS/http/)
- [https://www.w3.org/Protocols/HTTP/1.1/rfc2616bis/draft-lafon-rfc2616bis-03.html](https://www.w3.org/Protocols/HTTP/1.1/rfc2616bis/draft-lafon-rfc2616bis-03.html)
- [https://flaviocopes.com/http/](https://flaviocopes.com/http/)
- [https://www.extrahop.com/resources/protocols/http/](https://www.extrahop.com/resources/protocols/http/)

> ПРИМЕЧАНИЕ: номер потока - уникальный id потока. Номер параллельного потока. 