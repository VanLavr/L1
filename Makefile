run:
	@git add .
	@git commit -am "added a task"
	@git push

task:
	@go run task7Map.go