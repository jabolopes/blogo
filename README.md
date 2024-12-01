# blogo

A simple program to create blogs.

I created it because I fell in love with
[bashblog](https://github.com/cfenollosa/bashblog) but there were things I
wanted to adapt to my own needs and hacking shell scripting wasn't
straightforward for me.

While bashblog still remains the option that has no (required) external
dependencies, blogo provides a few different features that bashblog doesn't:

* Markdown is first class - blogo generates posts from Markdown only, it
  doesn't generate posts from HTML. This means, in blogo there is no ambiguity
  which is the original source document, unlike in bashblog (Markdown or
  HTML?). Removing this ambiguity makes blog management easier.

* Post dates are stored inside the Markdown post itself - this feature is
  important to me because I have many old posts that were written months or
  years ago. bashblog made it very difficult to manage old post dates because it
  uses a complex system of timestamps, where the timestamp can be managed by the
  filesystem or as a special marker in the generated HTML file. In the case of
  blogo, the post date is stored inside the post itself.

* Clear file organization - blogo organizes the files in clear subdirectories
  and the Markdown post files are clearly separated from the rest of files, in
  particular the output directory, the template files, and the generational
  programs. In bashblog, all files are in the toplevel directory and it becomes
  very difficult to distinguish between what is source and what is generated, in
  particular if one wants to later export to a server only the generated files
  and not the source files.

* Reproducible generation / builds - I wanted to have the generational process
  (or build process) deterministic and reproducible. There's likely still work
  to be done to make the build more deterministic but by separating the source
  and outputs, and by treating the generation as a build process, I managed to
  get very close to that goal.

Unlike bashblog, blogo has an dependency on the [Go](https://go.dev/)
programming language. The [Go toolchain must be
installed](https://go.dev/doc/install). The Markdown script is included in
blogo so it doens't have to be downloaded / installed.

## Usage

1. Download and install the [Go toolchain](https://go.dev/doc/install) (if not
   already installed).

2. Download blogo either by cloning this repository or by downloading the files
   from GitHub.

3. Add a post by creating a file with the extension `.md`, e.g., `mypost.md` in
   the same directory as this file (`README.md`).

4. Run `make` to generate your blog. All the blogs files (that you'd need to
   deploy to a Web server) are written to the `out/dist/` directory.

That's all!

If you'd like to launch a Web server to test your blog, do the following
(requires Python3 installed):

4. Run `make run` and in your browser visit `http://localhost:8000`.

## Documentation

### How to create a post?

1. Create a Markdown file (e.g., `my-post.md`) in the toplevel directory.
2. Run `make rebuild`.

See `my-first-post.md` for a template.

### How to set tags in posts?

1. Edit the Markdown file of a post.
2. Add a line at the end of the post with the format `Tags: tag1,
   tag2, etc`. For example, `Tags: story, scifi`.
3. Run `make rebuild`.

See `my-first-post.md` for a template.

### How to set date in posts?

1. Edit the Markdown file of a post.
2. Add a line after the title with the format `Date: Month Day, Year
   &mdash; Author Name</div>`, e.g., `Date: November 04, 2022 &mdash;
   Jose Lopes</div>`.
3. Run `make rebuild`.

See `my-first-post.md` for a template.

### How to delete a post?

1. Delete the corresponding Markdown file.
2. Run `make rebuild`.

### How to change the blog's configuration, e.g., blog name, etc?

1. Edit the file `bin/main.go`. It contains constants, such as,
   `blogName`, `blogDescription`, `authorName`, among others. These
   parameters can be changed.
2. Run `make rebuild`.

### How to change the blog's appearance?

Edit the CSS files (see `css/` directory) and the HTML templates (see
`templates/` directory), and run `make rebuild`.

Any CSS files stored in the `css/` directory are automatically copied
to the `out/dist/` directory when running `make`.

To link new CSS files to your blog's HTML pages edit the
`templates/index.template` file and include CSS include tags.

### How to add custom HTML pages?

1. Create a directory named `html/`.
2. Add any HTML pages to the `html/` directory.
3. Run `make rebuild`.

### License

The license (see `LICENSE`) covers the blog generation software included in this
repository.

The license does not cover any websites generated using this software. For
example, if you use this software to generate your blog, the posts and the HTML
pages generated are owned by you and this license does not apply to them.

In other words, you retain all the rights of the contents of your blog even if
those contents were generated by this software.
