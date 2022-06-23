Эмулятор платежного шлюза.

Доступны следующие эндпоинты: добавление новой транзакции, изменение статуса платежа платежной системой, получение статуса платежа по его ID, просмотр всех транзакций пользователя по ID пользователя, просмотр всех транзакций пользователя по его Email, отмена платежа. Статусы платежей на английском, так как пока не решена проблема кодировок, русские буквы не отображаются. Модуль math/random не подключается, поэтому транзакции с ошибками реализованы в отдельном эндпоинте (5 в списке).

Для развертывания присутствуют Dockerfile и docker-compose.

Эндпоинты.

1) /onetrbid/id - получение статуса платежа по его ID. Считывается ID из строки запроса, по нему проводится поиск в базе данных и выводится статус.

   Запрос имеет следующий вид: curl http://localhost:9090/onetrbid/id (число)

2) /alltrbid/userid - получение списка всех транзакций пользователя. Из строки запроса считывается ID пользователя, проводится поиск в БД и выдаются все найденные транзакции.

   Запрос имеет следующий вид: curl http://localhost:9090/alltrbid/userid (число)

3) /alltrbem/email - получение списка всех транзакций пользователя. Из строки запроса считывается Email пользователя, проводится поиск в БД и выдаются все найденные транзакции.

   Запрос имеет следующий вид: curl http://localhost:9090/alltrbem/email (строка)

4) /create - добавление новой транзакции. В запросе передаются следующие параметры: Userid, Email, Price, Currency. Данные передаются в базу данных, где происходит создание записи, поля ID, CreatedOn, UpdatedOn, Status заполняются автоматически (настройка полей в базе данных), статус принимает значение NEW. 

   Запрос имеет следующий вид: curl  http://localhost:9090/create -X POST -d "userid" -d "email" -d "price" -d "currency" (userid - целый тип, price - с плавающей точкой, email - строка, currency - строка)

5) /createfail - добавление новой транзакции с ошибкой. В запросе передаются следующие параметры: Userid, Email, Price, Currency. Данные передаются в базу данных, где происходит создание записи, поля ID, CreatedOn, UpdatedOn, Status заполняются автоматически (настройка полей в базе данных), Status примет значение FAILED. 

   Запрос имеет следующий вид: curl http://localhost:9090/createfail -X POST -d "userid" -d "email" -d "price" -d "currency" (userid - целый тип, price - с плавающей точкой, email - строка, currency - строка)

6) /reject - отмена платежа. В запросе передается ID транзакции. Из базы данных загружается соответствующая запись, после чего проверяется статус. Если он имеет значение SUCCES, UNSUCCES, REJECTED, то будет выдана ошибка что статус изменить нельзя. В противном случае статус будет изменен на REJECTED, параметр UpdatedOn изменится автоматически (настройка поля в базе данныъ).
 
   Запрос имеет следующий вид: curl http://localhost:9090/reject -X PUT -d "id" (id - число, целочисленный тип)

7) /changest - изменение статуса платежной системой. В запросе передается ID транзакции. По нему загружается запись из базы данных, у которой проверяется статус. Есл он имеет значение SUCCESS, UNSUCCESS, REJECTED - вернется ошибка что изменение невозможно. Если статус платежа NEW он будет изменен на SUCCES, UpdatedOn изменится автоматически (настройка поля в базе данных). Если статус платежа FAILED, то он будет изменен на UNSUCCESS, UpdatedON изменится автоматически (настройка поля в базе данных). Запрос должен предварительно пройти авторизацию. Для этого в заголовке запроса передается токен доступа. В данный момент реализовано просто как сравнение с каким-то сохраненным в коде значением. Возможно будет переведено на JWT авторизацию. Токен: 4hbkjdznfk3i27ecb1. Без токена будет выдана ошибка и запрос не будет выполнен.
 
   Запрос имеет следующий вид: curl http://localhost:9090/changest -X PUT -d "id" -H "Authorization: 4hbkjdznfk3i27ecb1" (id - число, целочисленный тип)

8) /changestatvs - изменение статуса платежной системой. В запросе передается ID транзакции и статус, который должен быть установлен. По ID загружается запись из базы данных, у которой проверяется статус. Есл он имеет значение SUCCESS, UNSUCCESS, REJECTED - вернется ошибка что изменение невозможно. Если статус платежа NEW он будет изменен на SUCCES, UpdatedOn изменится автоматически (настройка поля в базе данных). Если статус платежа FAILED, то он будет изменен на UNSUCCESS, UpdatedON изменится автоматически (настройка поля в базе данных). Запрос должен предварительно пройти авторизацию. Для этого в заголовке запроса передается токен доступа. В данный момент реализовано просто как сравнение с каким-то сохраненным в коде значением. Возможно будет переведено на JWT авторизацию. Токен: 4hbkjdznfk3i27ecb1. Без токена будет выдана ошибка и запрос не будет выполнен.

   Запрос имеет следующий вид: curl http://localhost:9090/changestatvs -X PUT -d "status" -d "id" -H "Authorization: 4hbkjdznfk3i27ecb1" (status - строка, id - число, целочисленный тип)
   
База данных имеет следующую структуру: 
БД - MySQL;
Название - constanta;
Таблица - transactions;
Поля:
id - int, primary key, auto increment;
userid - int;
email - varchar;
price - double;
currency - varchar;
createdon - datetime, по умолчанию current timestamp;
updatedon - datetime, по умолчанию current timestamp, on update current timestamp;
status - varchar;

Настройка параметров для подключения базы данных (пользователь, пароль, адрес, название базы данных) производится в каталоге models, файл -  transactions.go, переменная - dbparams. В папке mysql находится файл для создания таблицы необходимой конфигурации, если что-то пойдёт не так при использовании docker-compose.
