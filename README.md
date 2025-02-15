# Avito Backend Assignment 2025

## Setup

## Проблемы и выбор решения

### Использование API и авторизация
Данное задание предполагает использование конкретного API — поэтому ручек
больше не создавалось. Логика следующая:

1. Если пара username-password такая, что пользователь не найден, то регистрируем
его.
2. Если же мы получаем пару username-password такую, что username существует, то
в этом случае смотрим пароль и решаем, аутентифицировать пользователя или нет.
3. С одной передаваемой строкой в виде токена не реализовать модель с Access-
и Refresh-токенами.

### Создание DTO и Конвертеров
При имплементации взаимодействия между слоями на каждом слое были созданы свои
структуры для упрощения масштабирования продукта. Создавались они для каждого
сервиса, хендлера и репозитория.

Это касается и тривиального эндпоинта `/api/buy/:item` с одним параметром запроса.
Если структура запроса усложнится (например, мы захотим делать запрос, с помощью
которого захотим купить несколько элементов мерча), то можно будет легко добавить
эти поля в структуру и удобнее менять API.

В свою очередь, это требует создания конвертеров, которые копируют данные для
каждой новой модели.

### Последовательность при конвертации моделей между слоями
Тем не менее, в разных методах сервисного слоя я не преобразовывал repo-модель 
обратно в service-модель, потому что это привело бы к ненужному лишнему преобразованию
хэша в пароль. Хоть это и увеличивает зацепленность, это приведёт к более
эффективному коду.

### Понятие транзакций
При создании БД появилась идея создать фиктивного пользователя "SYSTEM", который
при регистрации начислял бы пользователям их стартовую 1000 монет. Посчитал, что
для дизайна это было бы более удобно: если бы условный фронтенд запрашивал данные об истории
изменения счёта пользователя, то он бы мог получить все данные об изменениях баланса,
опираясь на данные с бэкенда. Если же я бы не создавал эту транзакцию, то фронтенду
потребовалось бы генерировать это прибавление тысячи на стороне клиента, что нарушило
бы последовательность всех операций в кошельке.

От идеи отказался, потому что тогда, для того чтобы быть последовательным,
пришлось бы расширить понятие транзакции: получилось бы, что нам требуется 
указывать в транзакциях ещё и историю покупок. И что её и нужно было бы вернуть
в эндпоинте `info`. Поэтому для упрощения оставил сущность транзакций
употребимой лишь в терминах перевода монет от одного пользователя другому (без
покупок). Хотя в иных случаях, наверное, можно было бы расширить это понятие до
истории всех событий с кошельком: изначальное начисление, переводы и покупки.

Надеюсь, будете согласны с целесообразностью моей трактовки :)

### Структуризация методов в слое репозитория
Вариант Info использует все таблицы, поэтому было решено поместить метод, 
позволяющий собрать информацию со всех слоёв, в UserRepository как в центральную
сущность сервиса. Остальные репозитории выполняют сугубо индивидуальную задачу,
то есть не взаимодействуют с остальными таблицами.

### Метод WithTransaction
Для реализации транзакций в слое сервисов была создана функция `WithTransaction`.
В каждом репозитории она скопирована, поскольку логика транзакции у каждого репозитория
может отличаться. В случае, если требовалось бы обобщить реализацию транзакций
над несколькими репозиториями, можно было бы создать структуру, реализующую этот
метод, и инжектить её в структуры реализаций репозиториев в качестве зависимостей.

Но я посчитал, что это также усложнит код, и просто оставил в каждом репозитории
свой метод. Достаточно дискуссионный вопрос.

### Обработка ошибок
Для обработки ошибок реализована обёртка классической ошибки в сущность `AppError`,
которая также передаёт смысл ошибок, происходящих на внутренних слоях. Затем
api-слой обрабатывает коды этих ошибок и на их основе возвращает HTTP-код: это
сделано для меньшей связности и зацепленности относительно HTTP-интерфейса:
если бы нам потребовалось изменить интерфейс, например, на gRPC, нам не пришлось
бы переделывать сущность внутренних ошибок. Кажется, что это более масштабируемый
подход.

Большое количество функций с созданием разных типов в пакете `errors` создано
для лучшей читаемости кода.

### Пароль
Пароль должен быть не пустым.