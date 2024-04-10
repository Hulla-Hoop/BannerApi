# Avito 
## База данных

- БД Postgres
- Схема таблиц 

![Схема БД](/docs/schema.png)

- Данная схема посзволяет соблюдать условия 
   
    - 1. Один баннер может быть связан только с одной фичей и несколькими тегами
    - 2. При этом один тег, как и одна фича, могут принадлежать разным баннерам одновременно
    - 3. Фича и тег однозначно определяют баннер
   
- Создается путем миграций при старте приложения 
- Схема соответствует 3 нормальной базе

## Сервис