docker_sentry_init:
	docker-compose up -d
	docker-compose exec sentry sentry upgrade
	#docker-compose exec sentry pip install sentry-slack
	docker-compose restart sentry