# blogo

A simple program to create blogs.

![Main page screenshot](https://github.com/user-attachments/assets/ab20252f-3734-4d1c-9da5-fba9285ae0cf)

![All posts screenshot](https://github.com/user-attachments/assets/27bfd327-39d3-4ae7-9c33-14cbb41eec8c)

![All tags screenshot](https://github.com/user-attachments/assets/636cfed9-b10c-49d8-9d65-c9b8fc6512e2)

![A tag screenshot](https://github.com/user-attachments/assets/2aec10cb-5017-4a86-8523-4c03a383cbc2)

Inspired by [bashblog](https://github.com/cfenollosa/bashblog).

To create comic blogs see [comico](http://github.com/jabolopes/comico).

## Usage

1. Download and install the [Go toolchain](https://go.dev/doc/install) (if not
   already installed).

2. Download blogo either by cloning this repository or by downloading the files
   from GitHub.

3. Add post to the `posts/` directory. See existing examples in that directory.

4. Run `make` to generate your blog. All the blogs files (that you'd need to
   deploy to a Web server) are written to the `out/dist/` directory.

That's all!

If you'd like to launch a Web server to test your blog, do the following
(requires Python3 installed):

5. Run `make run` and in your browser visit `http://localhost:8000`.

## Documentation

### How to create a post?

1. Create a Markdown file, e.g., `posts/my-post.md`.
2. Run `make rebuild`.

See the `posts/` directory for examples.

### How to set tags in posts?

Tags are set directly in the post via the `Tags:` field.

See the `posts/` directory for examples.

### How to set the date in posts?

Dates are set directly in the post via the `Date:` field.

See the `posts/` directory for examples.

### How to delete a post?

1. Delete the file from the `posts/` directory.
2. Run `make rebuild`.

### How to change the blog's configuration, e.g., blog name, etc?

1. Edit the file `bin/main.go`.
2. Change the `blogName`, `blogDescription`, `authorName`, etc.
3. Run `make rebuild`.

### How to change the blog's appearance?

1. Edit the CSS files (see the `html/css/` directory)
2. Edit the HTML templates (see `templates/` directory)
3. Run `make rebuild`.

Any CSS files stored in the `html/css/` directory are automatically copied
to the `out/dist/` directory when running `make`.

To link new CSS files to your blog's HTML pages edit the
`templates/index.template` file and include CSS include tags.

### How to add custom HTML pages or custom files?

Any files in the `html/` directory are copied directly to the output.

Add any files or directories to the `html/` directory to have them automatically copied to the output when running `make`.

### License

The license (see `LICENSE`) covers the blog generation software included in this
repository.

The license does not cover any websites generated using this software. For
example, if you use this software to generate your blog, the posts and the HTML
pages generated are owned by you and this license does not apply to them.

In other words, you retain all the rights of the contents of your blog even if
those contents were generated by this software.
