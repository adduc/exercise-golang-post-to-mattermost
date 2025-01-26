reset:
	docker compose kill
	sudo rm -rf volumes
	docker compose rm --force --volumes

bootstrap:
	mkdir -p ./volumes/mattermost/{config,data,logs,plugins,client/plugins,bleve-indexes}
	sudo chown -R 2000:2000 ./volumes/mattermost