# Region Fixer
Веб приложение для генерации жалоб на инфраструктуру.

region-fixer/

- cmd/ # Точка входа приложения

- internal/ # Внутренние пакеты

- config/ # Конфигурация

- handlers/ # HTTP handlers

- infrastructure/ # Инфраструктура (БД, хранилище)

- middleware/ # Middleware

- models/ # Модели данных

- repository/ # Интерфейсы репозиториев

- router/ # Роутинг

- service/ # Бизнес-логика

- web/ # Frontend (шаблоны, статика)

- migrations/ # Миграции БД

## Требования

- Go 1.26+
- Docker и Docker Compose

## Установка

# Клонировать репозиторий

- git clone <url>

# Перейти в директорию

- cd region-fixer

# Установить зависимости

- go mod download