export $(shell cat .env | xargs)

BACKEND=./Cardora1
FRONTEND=./frontend

.PHONY: run-backend run-frontend run