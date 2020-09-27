.PHONY: genpdf
genpdf:
	rm -rf work && mkdir work
	cat sample/*.md > work/sample.md
	npx md-to-pdf work/sample.md

.PHONY: gentoken
gentoken:
	go run cmd/gentoken/main.go

.PHONY: upload
upload:
	go run cmd/upload/main.go

