.PHONY: blog clean cleanall fmt gallery generate sourcemaps

SOURCEMAPS := $(shell fd --type file --color never --extension html --no-ignore _templ_sourcemap)
GENERATED := $(shell fd --type file --color never --extension go _templ)

blog:
	templ generate --watch --proxybind="0.0.0.0" --proxy="http://localhost:3000" --cmd="go run ./examples/blog"

clean:
	rm $(SOURCEMAPS)

cleanall: clean
	rm $(GENERATED)

fmt:
	templ fmt .

gallery:
	templ generate --watch --proxy="http://localhost:3000" --cmd="go run ./examples/gallery"

generate:
	templ generate

sourcemaps:
	templ generate --source-map-visualisations

