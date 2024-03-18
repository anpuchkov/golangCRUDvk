# MoviesDelete
### Description

Удаляет фильм по его идентификатору.

### Parameters

movieId (integer): Идентификатор фильма, который нужно удалить.
### Responses

200 OK: Фильм успешно удален.

400 Bad Request: Неправильный запрос.

404 Not Found: Фильм с указанным идентификатором не найден.

500 Internal Server Error: Внутренняя ошибка сервера.

# MoviesGet
### Description

Получает список всех фильмов.

### Responses

200 OK: Список фильмов успешно получен.
500 Internal Server Error: Внутренняя ошибка сервера.


### Request Body

Отсутствует.

### Response Body

Список фильмов в формате JSON.

# MoviesPost

### Description

Добавляет новый фильм.

### Request Body

title (string): Название фильма.
description (string): Описание фильма.
release_date (string): Дата выхода фильма в формате "гггг-мм-дд".
rating (number): Рейтинг фильма.

### Responses

201 Created: Фильм успешно добавлен.
400 Bad Request: Неправильный запрос.
500 Internal Server Error: Внутренняя ошибка сервера.

### Response Body

Отсутствует.


# MoviesSearchGet

### Description

Ищет фильмы по части названия.

### Parameters

partOfTitle (string): Часть названия фильма.

### Responses

200 OK: Список найденных фильмов успешно получен.
400 Bad Request: Неправильный запрос.
500 Internal Server Error: Внутренняя ошибка сервера.

### Request Body

Отсутствует.

### Response Body

Список найденных фильмов в формате JSON.

# MoviesSortGet

### Description

Получает отсортированный список всех фильмов.

### Parameters

sortType (string, optional): Тип сортировки (date, title, rating). По умолчанию сортирует по рейтингу.

### Responses

200 OK: Список фильмов успешно получен.
400 Bad Request: Неправильный запрос.
500 Internal Server Error: Внутренняя ошибка сервера.

### Request Body

Отсутствует.

### Response Body

Список фильмов в формате JSON.


# MoviesUpdatePut

### Description

Обновляет информацию о существующем фильме.

### Request Body

movieId (integer): Идентификатор фильма, который нужно обновить.
title (string): Новое название фильма.
description (string): Новое описание фильма.
release_date (string): Новая дата выхода фильма в формате "гггг-мм-дд".
rating (number): Новый рейтинг фильма.

### Responses

200 OK: Фильм успешно обновлен.
400 Bad Request: Неправильный запрос.
404 Not Found: Фильм с указанным идентификатором не найден.
500 Internal Server Error: Внутренняя ошибка сервера.


### Response Body

Отсутствует.
