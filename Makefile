docker_sentry_init:
	docker-compose up -d
	docker-compose exec sentry sentry upgrade
	#docker-compose exec sentry pip install sentry-slack
	docker-compose restart sentry
	
creds:
	@echo ""
	@echo "goRest: http://localhost:8080"
	@echo "tabix: http://localhost:8124"
	@echo "clickhouse: http://localhost:8123; login:default; pass: "
	@echo "prometheus: http://localhost:9090"
	@echo "sentry: http://localhost:9010"
	@echo "minio: http://localhost:9001; login: minio-root-user; pass: minio-root-password"