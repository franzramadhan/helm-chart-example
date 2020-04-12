.PHONY: frontend backend

docker-login:
	docker login -u "$(username)"

minikube:
	minikube start

backend:
	go build -ldflags="-s -w" -o backend/main backend/main.go

frontend:
	pip install -r frontend/requirements.txt

run-backend:
	./backend/main

run-frontend:
	python ./frontend/main.py

publish-backend: docker-login backend
	cd backend && docker build --rm  --no-cache -t ${username}/backend:${version} .
	docker push ${username}/backend:${version}

publish-frontend: docker-login frontend
	cd frontend && docker build --rm --no-cache -t ${username}/frontend:${version} .
	docker push ${username}/frontend:${version}

helm-lint:
	helm lint ./chart

helm-pack:
	helm package ./chart

helm-check-nonprod: minikube helm-lint
	helm install --dry-run --debug --generate-name ./chart

helm-check-prod: minikube helm-lint
	helm install --dry-run --debug --generate-name --set isProduction=true ./chart

helm-install-nonprod: minikube helm-lint
	helm install --generate-name ./chart

helm-install-prod: minikube helm-lint
	helm install --generate-name --set isProduction=true ./chart
