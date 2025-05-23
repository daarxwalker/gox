# GOX
> Write reusable HTML in pure Go
```go
package app

import . "github.com/creamsensation/gox"

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

## Main reason to use GOX
> If you love Go, as I do, you love power && simplicity, then GOX is just for you.


<br>

**Feel free to test, submit issues, etc...**

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
  CustomData("track", "submit-button"),
)

// data-*
CustomData("*", "...")
```

<br>

### Shared
Some elements and attributes have same names, here comes shared nodes.<br>
All shared nodes have default node type, which can be modified with modifier node.<br>
Because of sharing, you need to use node for content, instead of writing simple string argument.<br>
Shared nodes are: cite, data, form, label, slot, span, style, title.<br>
```go
Label(Text("E-mail"))
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

Range(someSlice, func(item int) Node {
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
    CustomData("test", "test-id"),
    Text(content),
  )
}
```
<br>

## FAQ
**Why did I create the library?** <br>
> I was tired of JS / TS ecosystem, so much bloat and I wanted something simple, yet powerful, like Go, for frontend.

**What will be next steps?** <br>
> I want to create the whole ecosystem, which will use GOX as primary view system
