# GOX
> Write reusable HTML in pure Go
```go
package app

import . "github.com/daarxwalker/gox"

func Page() string {
    return Render(
        Html(
            Lang("en"),
            Head(
                Title(Text("Example app")),
                Meta(Name("viewport"), Content("width=device-width,initial-scale=1.0")),
            ),
            Body(
                H1(Text("Example page")),
                P(Text("Example paragraph")),
            ),
        ),
    )
}

```

<br>

## Why Choose GOX?
> If you love Go, as I do, you love power && simplicity, then GOX is just for you.

GOX offers an innovative approach to building web user interfaces, fully harnessing the power and simplicity of the Go language. Here are its key strengths:

* **Go-Centric Approach:** GOX fully leverages the advantages of the Go language, such as static typing and compilation, for HTML generation. This brings enhanced reliability and performance compared to tools dependent on the JavaScript ecosystem.
* **Declarative Syntax:** The library is inspired by modern frontend frameworks like React and Svelte. Thanks to its declarative UI, you'll enjoy clean and intuitive code that's easily understandable for modern developers.
* **Component-Based Thinking:** With GOX, you can easily create reusable components directly in Go. This is a huge advantage that significantly speeds up and simplifies the development of complex web applications.
* **Clean Code Generation:** GOX generates clean HTML without unnecessary "bloat." This approach aligns with Go's philosophy, which emphasizes efficiency and minimal overhead.
* **Functional Approach:** The use of functions for defining HTML elements and attributes is elegant and adheres to Go idioms. This keeps your code consistent and well-structured.
* **Raw Element:** For situations where you need to embed custom, unaltered HTML, CSS, or JavaScript, GOX provides a special Raw element. This ensures maximum flexibility.
* **Built-in Directives (If, Range):** GOX includes useful tools for conditional rendering and iterations directly within your Go code. This reduces the need for boilerplate code and simplifies dynamic content generation.
* **Extensible Plugin System:** GOX features a simple yet powerful plugin system that allows you to easily integrate custom Go structs directly into your HTML rendering flow, enhancing flexibility and reusability.
* **Seamless Interactivity:** GOX provides dedicated helper functions and utilities for lightweight JavaScript libraries like HTMX, Alpine.js, Datastar (github.com/daarxwalker/gox/pkg/htmx, github.com/daarxwalker/gox/pkg/alpine, github.com/daarxwalker/gox/pkg/datastar). This allows you to build dynamic and interactive UIs directly from your Go code, significantly reducing the need for complex frontend tooling.

<br>

**Feel free to test, submit issues, etc...**

<br>

## Installation
```bash
go get github.com/daarxwalker/gox
```

<br>

## Features
- [Elements](#elements)
- [Attributes](#attributes)
- [Shared](#shared)
- [Raw](#raw)
- [Text](#text)
- [Fragment](#fragment)
- [Modifiers](#modifiers)
- [Functions](#functions)
- [Directives](#directives)
- [Factories](#factories)
- [Plugins](#plugins)
- [Components](#components)
- [Interactivity](#interactivity)
- [FAQ](#faq)

<br>

### Elements
You can create HTML structure with built-in elements or create your own with factory function.<br>
Library should contains most of elements.<br>
```go
Html(
  Head(
    Title()
  ),
  Body(
    Div()
  ),
)
```

<br>

### Attributes
If you need to add attribute to element. You can create your own attributes with factory function.<br>
Library should contains most attributes.<br>
```go
Button(
  Type("submit"),
  Data(Name("track"), Text("submit-button")),
)

// data-*
Data(Name("*"), Text("..."))
```

<br>

### Shared
Some elements and attributes have same names, here comes shared nodes.<br>
All shared nodes have default node type, which can be modified with modifier node.<br>
Because of sharing, you need to use node for content, instead of writing simple string argument.<br>
Shared nodes are: cite, data, form, label, slot, span, style, title.<br>
```go
Label(For("email"), Text("E-mail"))
Style(Element(), Raw("body{background:red;}"))
```

<br>

### Raw
Sometimes you need to write custom HTML, use styles, or anything else...<br>
```go
Raw("body{background:blue;}")
Raw("<div>raw custom</div>")
```

<br>

### Text
Only purpose is to write text content.<br>
```go
Div(Text("Example"))
```

<br>

### Fragment
The fragment has purpose to group nodes.<br>
```go
Fragment(
  Div(),
  Span(),
  H1(),
)
```

<br>

### Modifiers
When you use shared nodes, sometimes you need to change node type, then you have to use modifiers.<br>
```go
Label(Attribute(), Text("example label"))
Style(Element(), Raw(".text{color:red;}"))
```

<br>

### Functions
#### Render
Render function transforms GOX nodes to string.<br>
You can use multiple nodes as arguments, it automatically wraps them with Fragment.<br>
```go
Render(
  Div(Text("example")),
)
```
#### Write
Write function transforms GOX nodes and write them as content to provided writer as bytes.<br>
```go
Write(
  w,
  Div(Text("example")),
)
```
<br>

### Directives
Directives are functions, which can be used for various operations.<br>
#### If
```go
someCondition := ...

If(someCondition, Div(Text("render if condition is true")))
```
#### Range
```go
someSlice := []int{1,2,3}

Range(someSlice, func(item int, _ int) Node {
  return Div(Text(item))
})
```
<br>

### Factories
#### Custom element
Create custom reusable element.<br>
```go
SomeCustomElement := CreateElement("x-custom")
Render(SomeCustomElement(Text("example content")))
```
#### Custom attribute
Create custom reusable attribute.<br>
```go
SomeCustomAttribute := CreateAttribute[string]("x-custom")
Render(Div(SomeCustomAttribute("custom attribute conttent"), Text("example content")))
```
<br>

### Plugins
Custom Go struct can be used as GOX plugin, exists only one rule, must have Node method, which return gox.Node interface.<br>
```go
type CustomPlugin struct {
  Value string
}

func (p CustomPlugin) Node() gox.Node {
  return Div(Text(p.Value))
}

Render(
  Html(
    Body(
      CustomPlugin{ Value: "example" },
    )
  ) 
)
```
<br>

### Components
Power of GOX approach is to create simple reusable components like React, Svelte, etc...<br>
```go
func UiButton(content string) Node {
  return Button(
    Class("bg-blue-400", "text-white", "rounded"),
    Type("button"),
    Data(Name("test"), Text("test-id")),
    Text(content),
  )
}
```

<br>

### Interactivity

While GOX excels at generating reusable HTML in Go, it's also designed to empower you to build **dynamic and interactive user interfaces without the complexity of traditional JavaScript frameworks.**

We provide dedicated packages with helper functions and utilities for seamless integration with lightweight frontend libraries:

* **HTMX:** Easily add dynamic updates, AJAX requests, and interactive behaviors directly from your Go code using HTMX attributes. No need for custom JavaScript. ``github.com/daarxwalker/gox/pkg/htmx``
* **Alpine.js:** Manage local component state, event handling, and conditional rendering with minimal JavaScript, all expressed naturally within your GOX structure. ``github.com/daarxwalker/gox/pkg/alpine``
* **Datastar:** Get the best of both worlds for dynamic UIs, managing state and triggering actions directly within your GOX strucuture. ``github.com/daarxwalker/gox/pkg/datastar``

These integrations mean you can manage your backend, view layer, and much of your frontend interactivity all within the Go ecosystem, streamlining development and reducing context switching.

**Future Plans:** We're continuously working to enhance this ecosystem and expand your options for building reactive UIs with minimal external JavaScript.

<br>

## FAQ
**Why did I create the library?** <br>
> I was tired of JS / TS ecosystem, so much bloat and I wanted something simple, yet powerful, like Go, for frontend.

