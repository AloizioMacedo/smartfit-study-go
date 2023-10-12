# Introduction

This is a mini-project intended as self-study for HTMX + Go.

It uses as motivation a
[coding challenge](https://github.com/bioritmo/front-end-code-challenge-smartsite)
from a brazilian fitness company called [SmartFit](https://www.smartfit.com.br).

# Technologies

We use Go's native templating functionalities together with [HTMX](https://htmx.org/)
when we update the page. The whole project runs as a single service running
[Gin](https://gin-gonic.com/).

If built as a Docker image, it weighs around 40 MB.

# How to run

You can run the project directly with

```bash
go run
```

in the root of the repo in case you have [Go](https://go.dev/) installed.

The app gets served at port 8080, and the page can be accessed at "/".

If you don't have Go, you can build it as a Docker image and run it directly as a
container. In this case, remember to set the ports, e.g.
`docker run -p 8080:8080 {name}`.

# Commentary about the challenge

It is unclear how the hours filter that the challenge proposes is supposed to work.

The interpretation that at first seems to make most sense is the following:

> If the user marked a time slot X, then an entry should be showed
> as long as X intersects any interval of hours of any schedule of that entry.

The code implements this implementation we just described.

The other filter (the one about showing closed units) has very minor impact,
as only three units seem to be closed in the provided file.

# Rust version

I've also developed a version with [Rust](https://www.rust-lang.org/), take
a look in [this repository](https://github.com/AloizioMacedo/smartfit-study)
in case you are interested.
