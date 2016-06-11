.PHONY: build clean

build:
	docker build -t kappa-build -f build/Dockerfile.build .
	docker run --rm kappa-build sh -c "cat /kappa/kappa" > bin/kappa
	docker build -t kappa:latest -f build/Dockerfile .

clean:
	docker rmi kappa-build
	rm bin/*
