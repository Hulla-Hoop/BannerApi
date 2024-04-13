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

## Условия

### `1. Используйте этот API `

### `2. Тегов и фичей небольшое количество (до 1000), RPS — 1k, SLI времени ответа — 50 мс, SLI успешности ответа — 99.99%`

Ограничиение на количество тегов и фичей реализует код <https://github.com/Hulla-Hoop/BannerApi/blob/main/internal/service/serviceBanner/insert.go>

``` 
func (c *serviceBanner) validate(reqId string, banner model.BannerHttp) error {

	if banner.Feature_id <= 0 || banner.Feature_id >= 1000 {
		c.logger.WithField("ServiceBanner.validate", reqId).Error("некорректные данные ", banner)
		return ErrIncorrectData{msg: fmt.Sprintf("данные %d некорректны", banner.Feature_id)}
	}

	var errCount int
	for _, v := range banner.Tags_id {

		if v <= 0 || v >= 1000 {
			errCount++
			c.logger.WithField("ServiceBanner.validate", reqId).Error("некорректные тег -- ", v)
		}
	}
	if errCount == len(banner.Tags_id) {
		return ErrIncorrectData{msg: "все теги некорректны должен быть хотя бы один корректный тег"}
	}

	ok, err := regexp.MatchString(`^(https?|ftp):\/\/[^\s\/$.?#].[^\s]*$`, banner.Content.Url)
	if !ok || err != nil {
		c.logger.WithField("ServiceBanner.validate", reqId).Error("некорректные данные ", banner)
		return ErrIncorrectData{msg: fmt.Sprintf("данные %s некорректны в данном поле должна быть ссылка", banner.Content.Url)}
	}

	return nil

}
 ```
- Нагрузочный тест по ручке /banner
  
``` 
curl --location 'http://localhost:8080/user_banner?tag_id=1&feature_id=3&use_last_revision=false' \

--header 'token: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIifQ.IQohLkYqxzY9A6ent4MGs1NyBNcSQyiyAd5ZG_c39CEHbuKwOuMNXhMO5dg01rB9CSV5R7MchcaZHDYZs_k7Bg' 
```

![Нагрузочный](/docs/Нагрузочный%20тест.png)

- SLI  времени ответа до 50мс
- SLI  успешности 100%

## Бонус 

![Напоминания](/docs/Ребята.png)