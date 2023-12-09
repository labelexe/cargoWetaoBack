# Используем официальный образ Golang
FROM golang:latest

ENV app "api_back"

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем все файлы проекта в текущую директорию в контейнере
COPY . .

# Собираем приложение, учитывая структуру проекта
RUN go build -o ./bin/${app} ./cmd/${app}/main.go

# Определяем порт, который будет открыт в контейнере
EXPOSE 9761

# Команда, которая будет выполнена при запуске контейнера
CMD ["/bin/api_back"]