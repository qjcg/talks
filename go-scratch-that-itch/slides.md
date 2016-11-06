title: Go Scratch That Itch
author:
    name: John Gosset
    twitter: jgosset_
    github: qjcg
theme: ./themes/select
controls: true
progress: false
output: slides.html
style: style.css

-- cover

# Go Scratch That Itch
## John Gosset

---

### What?

- "There should be code that does that!"
- Here to convince you to do, instead of not do
- Your problems matter
- It's a simple message, but simple can be hard

---

### How?

- Small is different
- Scratching itches
- Some well-known itches
- Why go philosophy and features are well-suited to programming in-the-small

---

# Small is different

--- cat

## Schrodinger's Cat

--- tdist

---

### Stephen J. Gould: Size and Shape

> The geometry of space itself is the major reason for correlations between size
> and shape. **Simply by growing larger, an object that keeps the same shape will
> suffer a continual decrease in relative surface area**. The decrease occurs
> because volume increases as the cube of length (length x length x length), while
> surface increases only as the square (length x length): in other words, **volume
> grows more rapidly than surface**.

--- .sizeandshape

### Stephen J. Gould: Size and Shape (cont'd)

> I once overheard a children’s conversation in a New York playground. Two young
> girls were discussing the size of dogs. One asked: “Can a dog be as large as
> an elephant?” Her friend responded: “No, if it were as big as an elephant, it
> would look like an elephant.” I wonder if she realized how truly she spoke.

---

### Stephen J. Gould: Size and Shape (cont'd)

- As you scale up software, surface area (ex: public API) can be dragged down
by supporting utility functions, tooling, boilerplate code, etc
- Ask: "Can we keep this small?"

---

### In-the-large versus In-the-small

- [Programming in the large and programming in the small](https://en.wikipedia.org/wiki/Programming_in_the_large_and_programming_in_the_small) (DeRemer & Kron, 1975)
- **in-the-small**: does one thing well, often coded by single developer, complete in itself
- **in-the-large**: industrial teams working on systems project, months to years, often broken up into submodules

---

### Scratching Itches

> Every good work of software starts by scratching a developer's personal itch.
> -- Eric S. Raymond, The Cathedral and the Bazaar

---

### The story of Unix

- [Space Travel] developed by Ken Thompson on a GE635 system

[Space Travel]: https://en.wikipedia.org/wiki/Space_Travel_(video_game)

---

### The story of Unix

- Wanted to port to PDP7 to avoid Bell Labs' internal usage fees, but OS did
not satisfy...

<img height=400px src="img/pdp7.jpeg" />

--- gsti

# Go Scratch That Itch!

---

### Jennifer Dewalt

- Art student with no programming experience, wanted to learn how to turn her
creative ideas into code...

--- gsti

# Go Scratch That Itch!

---

### Jennifer Dewalt

- [180 websites in 180 days](http://jenniferdewalt.com/)
- Each day small, self-contained, open-sourced on GitHub & blogged
- Approach: JFDI (Just F'n Do It!)

---

### One of my own

- Wanted to create something to generate placeholder text and learn more about
Unicode...

--- gsti

# Go Scratch That Itch!

---

### One of my own

- [horeb](https://github.com/qjcg/horeb)

---

### Scratching Itches: Everybody has problems

- Everybody has problems
- Your problems are important!
- You have **unique local knowledge**
- Contrast with "Centrally Planned Software" / design by committee

---

### A few principles for programming in the small

- Solve the problem in the concrete first
- Consider existing tools
- The [Unix Philosophy](https://en.wikipedia.org/wiki/Unix_philosophy)
- YAGNI
- Make your objectives clear (see eg.: [README driven development](http://tom.preston-werner.com/2010/08/23/readme-driven-development.html)

---

### Unix Philosophy

> This is the Unix philosophy: Write programs that do one thing and do it well.
> Write programs to work together. Write programs to handle text streams, because
> that is a universal interface. -- Doug McIlroy

- See also: [Rob Pike's Five Rules of Programming](http://users.ece.utexas.edu/~adnan/pike.html)
	- Acme, an "Integrating" Development Environment ([rsc screencast](https://www.youtube.com/watch?v=dP1xVpMPn8M))
- [Playground: Line filter](https://play.golang.org/p/mpYwOHj2ma)

---

### YAGNI

You Ain't Gonna Need It

<img height=350px src="img/yagni.png" />

---

# Go is good for programming in the small

---

### Go Spec

<img height=220px width="50px" src="img/gomoving2.png" />

- Consider the name "Go": short, simple, moving and doing
- [The Go Spec](https://golang.org/ref/spec) is short and readable

---

### Go Says "No"

- [RSC on Generics](http://research.swtch.com/generic)
- [Go FAQ: Exceptions](https://golang.org/doc/faq#exceptions)
- Go is implemented as a procedural language (not e.g. primarily
functional)

---

### Easy Deployment

- Simple, minimalistic tooling (cf. Makefiles, venvs, package.json...)
- Cross compilation is trivial (`GOOS`, `GOARCH`)
- Static binaries

---

### Polymorphism Without Pain

- [Interfaces](https://gobyexample.com/interfaces) allow reusing code as in the
OO-style polymorphism, but via composition instead of requiring rigid object
hierarchies

---

### Go is a shop-built jig

- [Blog post by Rob Napier](http://robnapier.net/go-is-a-shop-built-jig)

> Go feels under-engineered because it only solves real problems.

---

### Go is a shop-built jig (cont'd)

> [...] little pieces of wood that help you hold plywood while you cut it, or spacers
> that tell you where to put the guide bar for a specific tool, or hold-downs that
> keep a board in place while you’re working on it. They’re not always pretty.
> They often solve hyper-specific problems and work only with your specific tools.
> And when you look at ones that have been used a lot, they sometimes seem a
> little weird.

---

### Go is a shop-built jig (cont'd)

> There might be a random cutout in the middle. Or some little piece that sticks
> off at an angle. Or the corner might be missing a piece. And when you compare
> them to “real” tools, “general” tools like you’d buy from a catalog, they’re
> pretty homey or homely depending on how you’re thinking about it.
> [...] really good at solving [...] very common problems for people who need
> to ship software

---

### Key Takeaways

Principles when programming in the small:
- Solve the problem in the concrete first
- YAGNI
- When in doubt, use brute force
- Write down what you want first (ex: [README driven development])

---

### Key Takeaways (cont'd)

Go features good for programming in the small:
- small spec
- says "no" to 
- Easy to deploy
- interfaces: polymorphism / code reuse without rigid object hierarchies
- community is vibrant, 3rd party libraries are good compared to competing
language ecosystems (and getting better)

---

### Conclusion

- Programming in the small is good
- Go is good for programming in the small
- Your problems matter
- You should...

--- gsti

# Go Scratch That Itch!

---

### References

- [README Driven Development](http://tom.preston-werner.com/2010/08/23/readme-driven-development.html)
- [Space Travel](https://www.bell-labs.com/usr/dmr/www/spacetravel.html)
- [Programming in the large and programming in the small]
- [Worse is Better](https://en.wikipedia.org/wiki/Worse_is_better)
- [Stephen J Gould: Size and Shape](http://www.naturalhistorymag.com/editors_pick/1974_01_pick.html)
- [Jennifer Dewalt: 180 Websites in 180 days](https://jenniferdewalt.com/)
- [The Art of Unix Programming: The Right Size of Software](http://www.catb.org/esr/writings/taoup/html/ch13s04.html)
- [Go Proverbs](https://go-proverbs.github.io/)
- [Unix Philosophy](https://en.wikipedia.org/wiki/Unix_philosophy)
- [When in doubt, use brute force](http://www.catb.org/jargon/html/B/brute-force.html)
- [Go is a shop-built jig](http://robnapier.net/go-is-a-shop-built-jig)

[Programming in the large and programming in the small]: https://en.wikipedia.org/wiki/Programming_in_the_large_and_programming_in_the_small
