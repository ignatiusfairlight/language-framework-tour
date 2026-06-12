# Gin Framework Post Mortem

Gin, gin, gin. Oh wow. This one is a lot to pack. From how to setup it's docker container to using it to make the API application, everything is very very different from what I've did in Django and Laravel.

### Building the docker setup

Unlike the previous two languages, Go is a compiled language (another one in this tour being Rust). So having volume set for Gin is pointless since every changes need to be rebuild before implementing. Which I think is a blessing and a curse. It's a blessing because Go's error message is very very very easy to understand. So when there's an error in compilation, you know exactly what to do. If you need an LLM to explain the error message, then uh, consider going back to primary school.

Here's the cursed part, broken down into steps.
    1. Make a `go.mod` file using `go mod init`
    2. Make a `main.go`. It must contain `package main` and `func main()`
    3. Run `go mod tidy` to spawn `go.sum`

Seems simple right? Under normal circumstances, this is very easy to implement under the assumption that this is a local environment. However, like my previous tour, this is done in a docker container. So all of these commands are required to be executed in a temporary container. So the steps becomes like this:
    1. Make a `go.mod` file using `docker run --rm -v $(pwd):/app -w /app golang:1.26-bookworm go mod init <module-path>`
    2. Make a `main.go`. It must contain `package main` and `func main()` in your local directory
    3. Run `docker run --rm -v $(pwd)/backend/gin:/app -w /app golang:1.26-bookworm go mod tidy` to spawn `go.sum`

Now that I've written it like this, it does look rather simple than I experienced. Maybe the pain was only because it was my first time doing it? Who knows. I'm open to work on another Go/Gin related projects.

### What to expect when installing Gin

Unlike Laravel and Django, Gin is just a package. Get it into the project is as simple as importing it in your `main.go` and run `go mod tidy`. Which is the easy part. The tricky part is that nothing changes after this process. Nothing **obvious** that is. Unlike the other two frameworks, Gin provides most of the features except for the project structure, conventions and the ORM. Initially I thought I need to figure this all by myself, which is fine, this is a simple project to understand one aspect of backend development. But dear Opus 4.8 here pointed me to `gothinkster/golong-gin-realworld-example-app`, in which case I try to emulate without ignoring my own concepts that I carry (like having a service layer). So based on gothinkster, here's what I end up settling with:
    1. Project structure → Domain based (posts, comments, users in their own folder)
    2. Convention → A mix of Laravel and gothinksters. In `router`, I should've strictly follow Laravel's convention (index, show, delete, etc) but I ended up prefixing it with the domain (so it becomes PostIndex, PostDelete, PostUpdate, etc). While in the service layer, I use REST-like semantics (GetPosts, GetPostById, CreatePost, UpdatePost, etc)
    3. ORM → use `gorm.io/gorm`

I wouldn't call it the best way to use Gin, in fact it is far from the perfect setup. But at least it was interesting setting it all up by myself.

### Verdict/Conclusion/What I've learned

Go feels way different from PHP and Python. I don't have the right word for it for now but I know I'm within Data Oriented Programming in this project. Anyways, here are the things I find cool when doing this project:
    1. In a folder ('domain' going forward), functions are not required to be imported between files like in PHP. As long as they have different names and are within the same domain, you can call it.
    2. I think I somewhat applied Domain Driven Design here?
    3. Fowler's Service Layer can still be applied here.
    4. When compiling/building the application after making changes, Go's compiler can catch errors and these errors are very very well written. I have still yet to understand why people love Golang but I think this is one of them.


