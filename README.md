# Старт

Для старта вам нужно иметь установленный Go и sqlite. Перейдите в корневую папку проекта и выполните команду `make run`, чтобы запустить проект.

# Установка и запуск

Проект запускается на локальной машине, и Makefile содержит команды для удобства:

- `make run` — запускает приложение.
- `make clean-data` — очищает данные из базы данных и кеша.
- `make stop` — останавливает работу приложения.
- `make server-logs` — выводит логи работы сервера.
- `make database-logs` — выводит логи работы базы данных.
- `make cache-logs` — выводит логи работы кеша.
- `make all-logs` — выводит все логи вместе.

# О проекте

Проект **PET_Order_Pizza_App** — это учебное приложение для заказа пиццы, разработанное с использованием gRPC. Оно предполагает использование PostgreSQL для хранения данных и Redis для кеширования.

Проект придерживается чистой архитектуры и принципов SOLID, что делает его легко масштабируемым и поддерживаемым.

## Особенности проекта

### gRPC

Приложение использует gRPC для коммуникации между клиентом и сервером. gRPC обеспечивает высокую производительность и сжатие данных, что делает его отличным выбором для распределённых систем.

### Migrations

Для управления схемой базы данных используются миграции. Это позволяет легко обновлять и изменять структуру базы данных по мере развития проекта. Миграции применяются автоматически при запуске приложения.

### Tests

В проекте реализованы автоматические тесты, которые покрывают основные сценарии использования приложения. Тесты выполняются при каждом изменении кода, что помогает поддерживать высокое качество и надёжность кода.

### Deploy

Проект готов к деплою на любой инфраструктуре, поддерживающей Go. Все зависимости управляются через Go Modules, что упрощает развертывание на любой машине.

# API

## gRPC Server Information

- **Host:** `localhost`
- **Port:** `44044`
- **Protocol:** gRPC

## Service: `POrder`

### Method: `PlaceOrder`

Places a new order for a pizza.

**Request:**

- **Customer Name:** `string`
- **Pizza Type:** `string` (e.g., "Margherita", "Pepperoni")
- **Size:** `PizzaSize` (enum: `SMALL`, `MEDIUM`, `LARGE`)
- **Toppings:** `repeated string` (Optional - extra toppings)

**Example Command:**

```bash
grpcurl -plaintext -d '{
  "customer_name": "Dmitriy",
  "pizza_type": "Margherita",
  "size": "MEDIUM",
  "toppings": ["extra cheese", "olives"]
}' localhost:44044 pizza_order.POrder/PlaceOrder
```
**Response:**

- **Order ID:** `string` - Unique identifier for the order.
- **Message:** `string` - Confirmation message.

### Method: `CheckOrderStatus`

Checks the status of an existing order.

**Request:**

- **Order ID:** `string`

**Example Command:**

```bash
grpcurl -plaintext -d '{
  "order_id": "12345"
}' localhost:44044 pizza_order.POrder/CheckOrderStatus
```

**Response:**

- **Order ID:** `string` 
- **Status:**  OrderStatus (enum: PREPARING, ON_THE_WAY, DELIVERED, CANCELLED)

### Method: `CancelOrder`

Cancels an existing order.

**Request:**

- **Order ID:** `string`

**Example Command:**

```bash
grpcurl -plaintext -d '{
  "order_id": "12345"
}' localhost:44044 pizza_order.POrder/CancelOrder
```
**Response:**

- **Order ID:** `string` 
- **Message:** `string` - Confirmation message.

## Prerequisites
- Ensure that the gRPC server is running on `localhost:44044`
- Install `grpcurl` if it is not already installed:

```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```