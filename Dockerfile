FROM node:20-alpine AS frontend
WORKDIR /app
COPY frontend/package*.json ./frontend/
RUN cd frontend && npm ci
COPY frontend/ ./frontend/
RUN cd frontend && npm run build

FROM golang:1.23-alpine AS backend
WORKDIR /app
RUN apk add --no-cache gcc musl-dev
COPY --from=frontend /app/frontend/dist ./frontend/dist
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o abbygym

FROM alpine:3.19
WORKDIR /app
RUN apk add --no-cache ca-certificates tzdata
COPY --from=backend /app/abbygym .
EXPOSE 8080
CMD ["./abbygym"]
