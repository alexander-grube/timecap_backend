local:
	go build .
	./backend-bugtrack -l

prod:
	go build .
	./backend-bugtrack