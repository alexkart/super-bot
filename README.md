# Telegram бот

## Основная функциональность

Бот слушает чат Telegram и реагирует на определенные команды и фрагменты текста.

В режиме экспортирования сохраняет лог сообщений в HTML файл.

## Статус

Бот в работе несколько лет и успешно "участвовал" во многих подкастах. 

## Команды бота

| Команда                                   | Описание                                                                                                       |
|-------------------------------------------|----------------------------------------------------------------------------------------------------------------|
| `ping`, `пинг`                            | ответит `pong`, `понг`, см. [basic.data](https://github.com/radio-t/gitter-rt-bot/blob/master/data/basic.data) |
| `анекдот!`, `анкедот!`, `joke!`, `chuck!` | расскажет анекдот с jokesrv.fermyon.app или chucknorris.io                                        |
| `so!`                                     | 1 вопрос со [Stackoverflow](https://stackoverflow.com/questions?tab=Active)                                    |
| `?? <запрос>`, `/ddg <запрос>`            | поискать "<запрос>" на [DuckDuckGo](https://duckduckgo.com)                                                    |
| `chat! <запрос>`                          | задать вопрос для ChatGPT                                                                                      |

## Инструкции по локальной разработке

Для создания тестового бота нужно обратиться к [BotFather](https://t.me/BotFather) и получить от него токен.

После создания бота нужно вручную добавить в группу (Info / Add Members) и дать права администратора (Info / Edit / Administrators / Add Admin).

Приложение ожидает следующие переменные окружения:

* `TELEGRAM_TOKEN` – токен полученный от BotFather
* `TELEGRAM_GROUP` - основная группа в Телеграмме (туда приходят уведомления о новостях, все сообщения сохраняются в лог)
* `MASHAPE_TOKEN` – токен от сервиса [Kong](https://konghq.com/), используется только для DuckDuckGo бота
* `OPENAI_AUTH_TOKEN` – токен от сервиса [OpenAI Platform](https://platform.openai.com/), используется только для получения ChatGPT ответов в OpenAI боте

Дополнительные переменные окружения со значениями по-умолчанию:

* `DEBUG` (false) – включает режим отладки (логируется больше событий)
* `TELEGRAM_LOGS` (logs) - путь к папке куда пишется лог чата
* `SYS_DATA` (data) - путь к папке с *.data файлами и шаблоном для построения HTML отчета
* `TELEGRAM_TIMEOUT` (30s) – HTTP таймаут для скачивания файлов из Telegram при построении HTML отчета
* `RTJC_PORT` (18001) – порт на который приходят уведомления

Запустить бота можно через Docker Compose:

```bash
docker-compose up telegram-bot
```

Или с помощью Make:

```bash
make run ARGS="--super=umputun --super=bobuk --super=grayru --super=ksenks"
```

Для построения HTML отчета необходимо передать дополнительные флаги:

```bash
docker-compose exec telegram-bot ./telegram-rt-bot \
  --super=umputun \
  --super=bobuk \
  --super=grayru \
  --super=ksenks \
  --export-num=688 \
  --export-path=html \
  --export-day=20200208 \
  --export-template=logs.html
```

или

```bash
make run ARGS="--super=umputun --super=bobuk --super=grayru --super=ksenks --export-num=688 --export-path=logs --export-day=20200208 --export-template=data/logs.html"
```
