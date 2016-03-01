BEATNAME=rabbitbeat
BEAT_DIR=github.com/cero-t
SYSTEM_TESTS=false
TEST_ENVIRONMENT=false
ES_BEATS=./vendor/github.com/elastic/beats
GOPACKAGES=$(shell glide novendor)
PREFIX?=.

# Path to the libbeat Makefile
-include $(ES_BEATS)/libbeat/scripts/Makefile

.PHONY: init
init:
	glide update  --no-recursive
	make update
	git init
	git add .

.PHONY: install-cfg
install-cfg:
	mkdir -p $(PREFIX)
	cp etc/rabbitbeat.template.json     $(PREFIX)/rabbitbeat.template.json
	cp etc/rabbitbeat.yml               $(PREFIX)/rabbitbeat.yml
	cp etc/rabbitbeat.yml               $(PREFIX)/rabbitbeat-linux.yml
	cp etc/rabbitbeat.yml               $(PREFIX)/rabbitbeat-binary.yml
	cp etc/rabbitbeat.yml               $(PREFIX)/rabbitbeat-darwin.yml
	cp etc/rabbitbeat.yml               $(PREFIX)/rabbitbeat-win.yml

.PHONY: update-deps
update-deps:
	glide update  --no-recursive
