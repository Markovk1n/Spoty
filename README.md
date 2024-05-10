# Spoty

# Инструкция по установке и запуску Spoty

## Как установить

1. Скачайте проект Spoty из репозитория на GitHub.
2. Разархивируйте проект в удобное место на вашем компьютере.
3. Убедитесь, что у вас установлен Docker. Если нет, загрузите и установите Docker с [официального сайта](https://www.docker.com/get-started).

## Как запустить

1. Откройте терминал (командную строку).
2. Перейдите в каталог, где вы разархивировали проект Spoty.
3. Введите следующую команду для запуска проекта:


```bash
docker-compose up
```

4. Дождитесь завершения загрузки и инициализации приложения.
5. Откройте браузер и перейдите по адресу [http://localhost:8080](http://localhost:8080)
6. Теперь вы можете пользоваться приложением Spoty!

## Usage
1. Прежде чем начать использовать наше приложение, необходимо пройти процесс авторизации. Для этого, пожалуйста, выберите опцию "Регистрация" (Sign-up).
2. Задайте себе уникальное имя пользователя (не менее 4 символов) и пароль.
3. После успешной регистрации, вы будете перенаправлены на главную страницу.
4. Теперь, для входа в систему, повторно нажмите на кнопку "Авторизация" и выберите "Вход" (Login).
5. Введите свои учетные данные, и система выдаст вам JWT-токен.
6. Поздравляем! Теперь вы готовы использовать все возможности приложения.

### endpoints
Я не успел написать Swagger документацию, по этому временно информация по основным endpoint-там будет здесь.

```python
localhost:8080/search   method:GET 
Введите в формочку имя артиста, название альбома или трека. Вас перенаправит на страницу с результатами.
```
##
```python
localhost:8080/sign-up   method:POST
Принимает запрос на регистрацию пользователя. Отправляет в тело JSON {username :"", password:""}
```

```python
localhost:8080/sign-in   method:POST
Принимает запрос на авторизацию пользователя. Отправляет в тело JSON {username :"", password:""}, Получает: JWT token
```
##
```python
localhost:8080/albums/:id   method:GET
В проекте есть папка "spotifyID", там я написал рандомные ID альбомов. Вы можете вставить свое значение вместо ":id" 
```
```python
localhost:8080/albums/:id   method:POST
Отправляет запрос на создание коментарий под альбомом. Отправляет в тело JSON {comment :""}
```
##
```python
localhost:8080/track/:id   method:GET
В проекте есть папка "spotifyID", там я написал рандомные ID треков. Вы можете вставить свое значение вместо ":id" 
```
```python
localhost:8080/track/:id   method:POST
Отправляет запрос на создание коментарий под треком. Отправляет в тело JSON {comment :""}
```
##
```python
localhost:8080/artists/:id   method:GET
В проекте есть папка "spotifyID", там я написал рандомные ID артистов. Вы можете вставить свое значение вместо ":id"
```



## Description

Проект Spoty – это веб-приложение, предоставляющее пользователям возможность подробного ознакомления с альбомами или отдельными синглами. На платформе отображается информация о каждой песне, включая обзоры и рейтинги. Пользователи имеют возможность добавлять новые обзоры и выставлять оценки песням или альбомам.

Важной функциональностью является возможность поиска, позволяющая пользователям находить альбомы по названию, исполнителю или жанру, что упрощает процесс поиска желаемой музыки.

Spoty также включает в себя главную страницу с новостями о выходе новых альбомов или синглов, а также информацией о текущих трендах в музыкальной индустрии.


## Experience

Разработка приложения Spoty была захватывающим и учебным опытом для меня. Изначально я исследовал список доступных API, однако многие из них были платными или не соответствовали моим требованиям. В конечном итоге я наткнулся на API Spotify, которое оказалось идеальным для моих целей.

Работа с OAuth была новым для меня, и первые попытки получить токен были неудачными. Однако, я нашел альтернативный способ, и в итоге все заработало как задумано.

Следующей преградой было отсутствие в API списка новых и популярных альбомов. Хотя API предоставляло такую информацию, она не менялась часто, и в основном было доступно всего два варианта для тестирования. Решив эту проблему, я просто создал свой собственный список и отобразил его на главном экране приложения.

Из-за ограниченного времени на проект я не успел реализовать все идеи и функциональности, которые не входили в первоначальное задание. Некоторые из них включали в себя возможность пользователей ставить лайки комментариям, альбомам и трекам, а также оценивать их по нескольким критериям по 10-балльной шкале. Я также хотел добавить кнопку для случайного воспроизведения трека или альбома.

Несмотря на эти ограничения, основные функции приложения, такие как вывод информации об артистах, альбомах и треках, оценка альбомов и треков, поисковая система и даже функция прослушивания трека были успешно реализованы.

## Issues 

Если у вас возникнут трудности с запуском или у вас появятся какие-либо вопросы, не стесняйтесь обращаться ко мне через Telegram по адресу @owo_sh

###### created by Aryspayev Dias

