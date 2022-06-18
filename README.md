Эмулятор платежного шлюза.

Доступны следующие эндпоинты: добалвение новой транзакции, изменение статуса платежа платежной системой, получение статуса платежа по его ID, просмотр пвес транзакций пользователя по ID пользователя, просмотр всех транзакций пользователя по его Email, отмена платежа.
Эндпоинты.
1) /onetrbid/id - получение статуса платежа по его ID. Считывается ID из строки запроса, по нему проводится поиск в базе данных и выводится статус.
2) /alltrbid/userid - получение списка всех транзакций пользователя. Из строки запроса считывается ID пользователя, проводится поиск в БД и выдаются все найденные транзакции.
3) /alltrbem/email - получение списка всех транзакций пользователя. Из строки запроса считывается Email пользователя, проводится поиск в БД и выдаются все найденные транзакции.
4) /create - добавление новой транзакции. В запросе передаются следующие параметры: Userid, Email, Price, Currency. ДАнные передаются в базу данных, где происходит создание записи, поля ID, CreatedOn, UpdatedOn, Status заполняются автоматически (настройка полей в базе данных), статус принимает значение NEW. 
Запрос имеет следующий вид: curl  http://localhost:9090/create -X POST -d "123" -d "ex@go.ru" -d "102.23" -d "USD"
5) /reject - отмена платежа. В запросе передается ID транзакции. Из базы данных загружается соответствующая запись, после чего проверяется статус. Если он имеет значение SUCCES, UNSUCCES, REJECTED, то будет выдана ошибка что статус изменить нельзя. В противном случае статус будет изменен на REJECTED, параметр UpdatedOn изменится автоматически (настройка поля в базе данныъ)
6) /changest