# Используем официальный образ Go для сборки приложения
FROM golang:1.18 as builder

# Устанавливаем рабочий каталог внутри контейнера
WORKDIR /app

# Копируем все файлы проекта в контейнер
COPY . .

# Загружаем зависимости
RUN go mod download

# Сборка Go-приложения
RUN go build -o main ./cmd/app/main.go

# Используем минимальный базовый образ для запуска приложения
FROM gcr.io/distroless/base-debian10

# Устанавливаем рабочий каталог внутри контейнера
WORKDIR /app

# Копируем скомпилированное приложение из этапа сборки
COPY --from=builder /app/main .

# Копируем конфигурационный файл
COPY --from=builder /app/config/local.yaml ./config/local.yaml

# Устанавливаем переменные окружения
ENV CONFIG_PATH=/app/config/local.yaml

# Определяем команду для запуска приложения
CMD ["./main"]
