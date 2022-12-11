SHELL := bash
DAY=`date +%d`
TIMEOUT=120m
TOKEN=''
PDF_VIEWER=''

build:
	@mkdir -p bin
	@for day in $(shell ls | grep -E "^day") ; do go build -trimpath -o bin/$${day} $${day}/$${day}.go ; done

day:
	@mkdir -p day${DAY}
	@mkdir -p inputs/test

	@if ! [ -f day${DAY}/day${DAY}.go ]; then \
		cat templates/template.go.tmpl | sed "s/DAYNUMBER/$$(echo ${DAY} | sed -E 's/^0//g')/g" > day${DAY}/day${DAY}.go; \
		cp templates/tests.go.tmpl day${DAY}/day${DAY}_test.go; \
		echo Created: day${DAY}/day${DAY}.go ; \
		mkdir -p inputs/test/day${DAY}/1 ;\
		touch inputs/test/day${DAY}/1/input.txt ; \
		touch inputs/test/day${DAY}/1/result_p1.txt ; \
		codium inputs/test/day${DAY}/1/input.txt \
		inputs/test/day${DAY}/1/result_p1.txt \
		inputs/test/day${DAY}/1/result_p2.txt \
		day${DAY}/day${DAY}.go ; \
	fi

testday:
	@echo RUNNING TESTS FOR DAY ${DAY}
	@go test day${DAY}/*.go  -v -timeout 0

testall:
	@RC=0; for day in $(shell ls inputs/test/) ; do echo TESTING $${day}; go test -timeout ${TIMEOUT} $${day}/*.go ; RET=$$?; if [ $$RET != 0 ]; then RC=$$RET; fi; done; exit $$RC

testallv:
	@RC=0; for day in $(shell ls inputs/test/) ; do echo TESTING $${day}; go test -v -timeout ${TIMEOUT} $${day}/*.go ; RET=$$?; if [ $$RET != 0 ]; then RC=$$RET; fi; done; exit $$RC

benchmark:
	@if [ ${TOKEN} != '' ]; then time for day in $(shell ls bin/) ; do go build -trimpath -o bin/$${day} $${day}/$${day}.go && time bin/$${day} -token ${TOKEN} ; done; else time for day in $(shell ls bin/) ; do time bin/$${day}; done; fi

profile:
	@mkdir -p profiles
	@TIMESTAMP=`date +%s`; file=`go run day${DAY}/day${DAY}.go 2>&1 | grep "cpu profiling disabled" | grep -Eo "[^ ]+$$"` ; \
		go tool pprof --pdf ./bin/day${DAY} $${file} > profiles/profile_day${DAY}_$${TIMESTAMP}.pdf ; \
		echo profiles/profile_day${DAY}_$${TIMESTAMP}.pdf; \
		if [ ${PDF_VIEWER} != '' ]; then (${PDF_VIEWER} `pwd`/profiles/profile_day${DAY}_$${TIMESTAMP}.pdf&); fi

clean:
	rm -fv bin/* profiles/*
