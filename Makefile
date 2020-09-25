.PHONY: generate
generate:
	rm -rf work && mkdir work
	cat sample/*.md > work/sample.md
	npx md-to-pdf work/sample.md

