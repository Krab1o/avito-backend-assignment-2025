# Avito Backend Assignment 2025

## Проблемы и выбор решения

### Создание DTO и Конвертеров
При имплементации взаимодействия между слоями на каждом слое были созданы свои 
для упрощения масштабирования продукта. 

Это касается и тривиального эндпоинта `/api/buy/:item` с одним параметром запроса.
Если структура запроса усложнится (например, мы захотим делать запрос, с помощью
которого захотим купить несколько элементов мерча), то можно будет легко добавить
эти поля в структуру и удобнее менять API.

В свою очередь, это требует создания конвертеров, которые копируют данные для
каждой новой модели.

### Целостность транзакций