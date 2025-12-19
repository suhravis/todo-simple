# Build frontend
FROM node:18-alpine AS frontend-builder
WORKDIR /app
COPY frontend/package.json frontend/package-lock.json* ./frontend/
WORKDIR /app/frontend
RUN npm ci
COPY frontend/ .
RUN npm run build

# Build backend
FROM golang:1.21-alpine AS backend-builder
WORKDIR /app/backend
COPY backend/go.mod backend/go.sum* ./
RUN go mod download
COPY backend/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o todo-backend .

# Final image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=backend-builder /app/backend/todo-backend .
COPY --from=frontend-builder /app/frontend/dist ./static

EXPOSE 8080

ENV PORT=8080

CMD ["./todo-backend"]

