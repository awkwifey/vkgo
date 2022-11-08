# Отправка сообщений
## SendMessage
Позволяет отправить сообщение, используя структуру Message
```go
bot := API.Options{Token: "", Version: 5.131}
bot.SendMessage(API.Message{
	ChatID: 542439242,
	Text: "🇷🇺 Hello, Russia!",
	RandomID: 0,
})
```

### SendText
Проще предыдущего метода, отправляет текст
```go
bot.SendText(542439242, "🇷🇺 Hello, Russia!"})
```
Внутри содержит RandomID, построенный на Time
## SendAttachments
### SendPhoto
Просто отправляет фотографию
```go
bot.SendPhoto(542439242, 542439242, 457279110)
```

## SendForwardedText
Вносит свой комментарий в чужое сообщение
```go
bot.SendForwardedText(message.ChatID, message.ID, "🇷🇺 Hello, Russia!")
```
Внимание! _conversation_message_id_ изменён на __ID__, а __MessageID__ используется вместо id. Подробнее об этом улучшении [здесь](https://github.com/botology/vkgo/%F0%9F%93%96%20Docs/%F0%9F%93%9A%20Messages/%F0%9F%87%B7%F0%9F%87%BA%20RU_ru.md)