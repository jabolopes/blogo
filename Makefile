all: generate

run: generate
	( cd out/dist; python3 -m http.server )

clean:
	rm -rf ./out/

rebuild: clean generate

md=bin/Markdown.pl

out:
	mkdir -p $@

out/dist:
	mkdir -p $@

# blogo program

SRC_BLOGO := $(wildcard bin/*.go)

out/blogo: $(SRC_BLOGO) | out
	go build -o $@ $(SRC_BLOGO)

# blog contents

out/%.pre: %.md out/blogo
	( out/blogo postify --titleHref="$(basename $(notdir $@)).html" | $(md) ) < $< > $@

out/dist/%.html: %.md out/%.pre out/blogo | out/dist
	out/blogo gen-post $< > $@

out/dist/%.css: css/%.css | out/dist
	cp $^ $@

SRC_MDS := $(wildcard *.md)
SRC_MDS := $(filter-out README.md, $(SRC_MDS))

SRC_PRES := $(patsubst %.md,out/%.pre,$(SRC_MDS))
OUT_HTMLS := $(patsubst %.md,out/dist/%.html,$(SRC_MDS))

SRC_CSS := $(wildcard css/*.css)
OUT_CSS := $(patsubst css/%,out/dist/%,$(SRC_CSS))

.PRECIOUS: $(SRC_PRES)

out/dist/index.html: $(SRC_MDS) out/blogo $(OUT_CSS) | out/dist
	out/blogo gen-index $(SRC_MDS) > $@

out/dist/all_posts.html: $(SRC_MDS) out/blogo $(OUT_CSS) | out/dist
	out/blogo gen-all-posts $(SRC_MDS) > $@

out/dist/all_tags.html: $(SRC_MDS) out/blogo $(OUT_CSS) | out/dist
	out/blogo gen-all-tags $(SRC_MDS) > $@

out/dist/feed.rss: $(SRC_MDS) out/blogo | out/dist
	out/blogo gen-feed $(SRC_MDS) > $@

generate: $(OUT_HTMLS) out/blogo out/dist/index.html out/dist/all_posts.html out/dist/all_tags.html out/dist/feed.rss $(OUT_CSS) | out/dist
	out/blogo gen-tag --out=out/dist/ $(SRC_MDS)

print-%  : ; @echo $* = $($*)
