all: generate

run: generate
	( cd out/dist; python3 -m http.server )

clean:
	rm -rf ./out/

rebuild: clean generate

out:
	mkdir -p $@

out/posts:
	mkdir -p $@

out/dist:
	mkdir -p $@

# blogo program

SRC_BLOGO := $(wildcard bin/*.go)

out/blogo: $(SRC_BLOGO) | out
	go build -o $@ $(SRC_BLOGO)

# assets

out/dist/%: html/% | out/dist
	cp -T -r $^ $@

SRC_HTML := $(wildcard html/**)
OUT_HTML := $(patsubst html/%,out/dist/%,$(SRC_HTML))

# blog contents

SRC_TEMPLATES := $(wildcard templates/*.template)

out/posts/%.post: posts/%.md out/blogo $(SRC_TEMPLATES) | out/posts
	out/blogo postify $<

out/dist/%.html: posts/%.md out/posts/%.post out/blogo | out/dist
	out/blogo gen-post $< > $@

SRC_POSTS := $(wildcard posts/*.md)
GEN_POSTS := $(patsubst posts/%.md,out/posts/%.post,$(SRC_POSTS))
OUT_POSTS := $(patsubst posts/%.md,out/dist/%.html,$(SRC_POSTS))

.PRECIOUS: $(GEN_POSTS)

out/dist/index.html: $(SRC_POSTS) out/blogo $(SRC_TEMPLATES) | out/dist
	out/blogo gen-index > $@

out/dist/all-posts.html: $(SRC_POSTS) out/blogo $(SRC_TEMPLATES) | out/dist
	out/blogo gen-all-posts > $@

out/dist/all-tags.html: $(SRC_POSTS) out/blogo $(SRC_TEMPLATES) | out/dist
	out/blogo gen-all-tags > $@

out/dist/sitemap.xml: $(SRC_POSTS) out/blogo | out/dist
	out/blogo gen-sitemap > $@

out/dist/robots.txt: $(SRC_POSTS) out/blogo | out/dist
	out/blogo gen-robots-txt > $@

out/dist/feed.rss: $(SRC_POSTS) out/blogo $(SRC_TEMPLATES) | out/dist
	out/blogo gen-feed > $@

generate: $(OUT_POSTS) out/blogo out/dist/index.html out/dist/all-posts.html out/dist/all-tags.html out/dist/feed.rss out/dist/sitemap.xml out/dist/robots.txt $(OUT_HTML) | out/dist
	out/blogo gen-tag

print-%  : ; @echo $* = $($*)
