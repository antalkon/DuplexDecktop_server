# Этап 1: Сборка
FROM golang:1.22 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем все файлы из текущей директории в контейнер
COPY . .

# Скачиваем зависимости и собираем приложение
RUN go mod download
RUN go build -o /app/bin/app cmd/app/main.go

# Этап 2: Минимальный контейнер для выполнения
FROM scratch

# Копируем бинарный файл из первого этапа
COPY --from=builder /app/bin/app /app/bin/app

# Устанавливаем рабочую директорию и указываем команду для запуска
WORKDIR /app/bin
CMD ["./app"]
