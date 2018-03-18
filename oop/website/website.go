package website

import "fmt"

type author struct {
    firstName string
    lastName  string
    bio       string
}

// Composition!
type Post struct {
    title   string
    content string
    author
}

// Can't anonymously embed slices, have to name them
type Website struct {
    Posts []Post
}

func NewPost(t string, c string, a author) Post {
    return Post{t, c, a}
}

func NewAuthor(first string, last string, bio string) author {
    return author{first, last, bio}
}

func NewWebsite(posts ...Post) Website {
    return Website{posts}
}

func (a author) FullName() string {
    return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

func (p Post) Details() {
    fmt.Println("title: ", p.title)
    fmt.Println("content: ", p.content)
    fmt.Println("author: ", p.FullName())
    fmt.Println("bio: ", p.bio)
}

func (w Website) Contents() {
    fmt.Println("Contents of Website:")
    for _, v := range w.Posts {
        v.Details()
        fmt.Println()
    }
}
