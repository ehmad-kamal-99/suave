.DEFAULT_GOAL := help

include Makefile.variables

.PHONY: build check codegen format help test todo vendor veryclean

help:
	@echo
	@echo 'Usage: make COMMAND'
	@echo
	@echo 'Commands:'
	@echo '  build           Compile project.'
	@echo '  check           Run linter.'
	@echo '  codegen         Generate code.'
	@echo '  format          Format source code.'
	@echo '  test            Run tests.'
	@echo '  vendor          Install dependencies.'
	@echo '  veryclean       Clean up target directory.'
	@echo

## prefix before other make targets to run in your local dev environment
local: | quiet
	@$(eval DOCKRUN= )
	@mkdir -p tmp
	@touch tmp/dev_image_id
quiet: # this is silly but shuts up 'Nothing to be done for `local`'
	@:

prepare: tmp/dev_image_id
tmp/dev_image_id:
	@mkdir -p tmp
	@docker rmi -f ${DEV_IMAGE} > /dev/null 2>&1 || true
	@docker build -t ${DEV_IMAGE} -f Dockerfile.dev .
	@docker inspect -f "{{ .ID }}" ${DEV_IMAGE} > tmp/dev_image_id

veryclean:
	@rm -rf tmp vendor

vendor: prepare
	@[ "${JENKINS_CI}" ] &&	${DOCKRUN} bash -c 'go mod vendor && chmod -R 777 vendor' || ( go mod vendor )

format: vendor
	${DOCKRUN} bash ./scripts/format.sh

check: format
	${DOCKRUN} bash ./scripts/check.sh

todo:
	${DOCKRUN} bash ./scripts/todo.sh

codegen: prepare
	${DOCKRUN} bash ./scripts/swagger.sh

test: check db_start
	${DOCKTEST} bash ./scripts/test.sh

db_stop:
	bash ./scripts/db_stop.sh

build:
	bash ./scripts/build.sh

db_start: db_stop
	@docker run -p 6379:6379 --name suave-db -d redislabs/rebloom:latest

seed: db_start
	bash ./scripts/seed.sh
