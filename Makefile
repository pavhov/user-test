# only for sirst start
start: init-dep init-services build-services run-services

init-dep:
	 cd initial-job && make dep
	 cd user-service && make dep

init-services:
	 cd initial-job/src/service && make env && make dep-install
	 cd user-service/src/service && make env && make dep-install

build-services:
	 cd initial-job/src/service && make build
	 cd user-service/src/service && make build

run-services:
	 make run-initial-job
	 make run-user-service

run-initial-job:
	 cd initial-job/src/service && make run

run-user-service:
	 cd user-service/src/service && make run
