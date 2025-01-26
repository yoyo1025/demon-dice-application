.PHONY: setup up d b ps node

setup:
	@make up
	@make ps
d:
	docker compose down
up:
	docker compose up -d
ps:
	docker compose ps
demon_dice_app:
	docker compose exec demon_dice_app sh
demon_dice_db:
	docker compose exec demon_dice_db sh
