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

### ` 3. Для авторизации доступов должны использоваться 2 вида токенов: пользовательский и админский.  Получение баннера может происходить с помощью пользовательского или админского токена, а все остальные действия могут выполняться только с помощью админского токена. `

- Используется токен с сигнатурой как на картинке ниже 

![sign](/docs/сигнатура%20токена.png)

- Доступные пользователи 

```
var Users model.Users = []model.User{
	// admin eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIn0.v4A2q3xg9-zFWjP_CTV2HQuteszG7Mx08GUOLUIfOnG1a2P9c2ZU1FKKRiVEpJVMMZvCb4JjlPTNrkzIy1tCbA
	{
		Username: "admin",
		Root:     true,
	},
	// user eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIifQ.IQohLkYqxzY9A6ent4MGs1NyBNcSQyiyAd5ZG_c39CEHbuKwOuMNXhMO5dg01rB9CSV5R7MchcaZHDYZs_k7Bg
	{
		Username: "user",
		Root:     false,
	}}

// chel(для теста доступа) eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImNoZWwifQ.2ZqHPnzoswMpS4WuvSFQdO97KoS0GYYD71kbY9dPTKl4RXFs1TtDYWlOBJn7liCE1eeozVn3Fgew-RxI7ZA-pQ

 ```
- проверка реализована в <https://github.com/Hulla-Hoop/BannerApi/blob/main/internal/endpoint/midllware/aut.go>

### `4.Реализуйте интеграционный или E2E-тест на сценарий получения баннера. `

Реализовано в коллекции postman (тесты на возвращаемые коды)

<https://github.com/Hulla-Hoop/BannerApi/blob/main/docs/April(Avito).postman_collection.json>

### 5. Если при получении баннера передан флаг use_last_revision, необходимо отдавать самую актуальную информацию.  В ином случае допускается передача информации, которая была актуальна 5 минут назад.

Данная механика реализуется пакетом кеш <https://github.com/Hulla-Hoop/BannerApi/blob/main/internal/repo/timecash>

- не стал добавлять удаление и обновление кеша при использовании ручек DELETE /banner/{id} , PATCH /banner/{id} так как в условии сказано только что они должны удаляться из кеша по истечении времени.


### 6. Баннеры могут быть временно выключены. Если баннер выключен, то обычные пользователи не должны его получать, при этом админы должны иметь к нему доступ.
Реализует код <https://github.com/Hulla-Hoop/BannerApi/blob/main/internal/repo/Db/getOne.go>

```
if !root {
		err = p.dB.QueryRow(`SELECT banner.title,banner.text,banner.url FROM banner
	JOIN chains
	ON banner.id=chains.banner_id
	WHERE chains.feature_id=$1 AND chains.tags_id=$2 AND banner.active=true;
	`, feature_id, tag_id).Scan(&b.Title, &b.Text, &b.Url)
	} else {
		err = p.dB.QueryRow(`SELECT banner.title,banner.text,banner.url FROM banner
	JOIN chains
	ON banner.id=chains.banner_id
	WHERE chains.feature_id=$1 AND chains.tags_id=$2;
	`, feature_id, tag_id).Scan(&b.Title, &b.Text, &b.Url)
	}

 ```





## Бонус 

![Напоминания](/docs/Ребята.png)