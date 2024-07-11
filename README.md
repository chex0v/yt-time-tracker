# Yt-time-tracker

Утилита которая позволяет получить/добавить информацию о треке времени в YouTrack

Использование:

`ytt [command]`

## Команды:
```
add         Добавить время в задачу
time        Узнать информацию о треке времени в задаче
me          Информация о моём профиле
report      Информация о учёте времени за сегодня или за промежуток времени при указании параметров.
```
`task - название задачи. Пример VUZ-1209 или vuz-1209`

`ytt add`:

`ytt add [task] [time] [message] [flags]`

Флаги:
```
-d, --date string   Дата события. Формат YYYY-MM-DD. Пример 2024-06-25
-t, --type string   Тип задачи из настроек TYPES
```

`ytt report [flags]`

Флаги:
```
-d, --date string      Дата начала промежутка времени. Формат YYYY-MM-DD
-t, --date-to string   Дата конца промежутка времени. Формат YYYY-MM-DD
```

В настройка можно задать шаблон для названия задачи, параметер TASK.

Пример:

```yaml
TASKS:
  - key: "home"
    value: "EWHO-27"
```

Теперь можно вызвать команду так, чтобы не вводить частые названия задач

```bash
ytt time home
ytt add home 1m "test"
```

По такому же принципу есть блок с типами времени

Пример:

```yaml
TYPES:
  - key: "dev"
    value:  "116-0"
```
Теперь можно вызвать команду так, чтобы не вводить частые названия задач и тип задачи не выбирать

```bash
ytt time home
ytt add home 1m "test" -t dev
```
