TEST_FLAGS? =
DB_USERNAME = root
DB_PASSWORD = password
DB_NAME     = vebra-test
DB_HOST     = 127.0.0.1
DB_PORT     = 3663
DB_NAME     = vebra-test

all:

test-full:
	go test ${TEST_FLAGS}

test:
	go test ${TEST_FLAGS} -short

test-integration:
	$(eval CONTAINER_ID := $(shell docker run --rm --name test-mysql -e MYSQL_ROOT_PASSWORD=${DB_PASSWORD} -e MYSQL_DATABASE=vebra-test -d -p ${DB_PORT}:3306 mysql))
	-$(shell while ! mysqladmin ping --host=127.0.0.1 --port=${DB_PORT} -u root --password=${DB_PASSWORD} --silent 2>/dev/null; do  sleep 1; done)
	export DB_USERNAME=${DB_USERNAME} && \
	export DB_PASSWORD=${DB_PASSWORD} && \
	export DB_HOST=${DB_HOST}         && \
	export DB_PORT=${DB_PORT}         && \
	export DB_NAME=${DB_NAME}         && \
	go test ${TEST_FLAGS} -run Integration
	docker stop test-mysql
