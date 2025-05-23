package gox

const (
	elementAnimate             = "animate"
	elementAnimateMotion       = "animateMotion"
	elementAnimateTransform    = "animateTransform"
	elementCircle              = "circle"
	elementDefs                = "defs"
	elementDesc                = "desc"
	elementEllipse             = "ellipse"
	elementFeBlend             = "feBlend"
	elementFeColorMatrix       = "feColorMatrix"
	elementFeComponentTransfer = "feComponentTransfer"
	elementFeComposite         = "feComposite"
	elementFeConvolveMatrix    = "feConvolveMatrix"
	elementFeDiffuseLightning  = "feDiffuseLighting"
	elementFeDisplacementMap   = "feDisplacementMap"
	elementFeDistantLight      = "feDistantLight"
	elementFeDropShadow        = "feDropShadow"
	elementFeFlood             = "feFlood"
	elementFeFuncA             = "feFuncA"
	elementFeFuncB             = "feFuncB"
	elementFeFuncG             = "feFuncG"
	elementFeFuncR             = "feFuncR"
	elementFeGaussianBlur      = "feGaussianBlur"
	elementFeImage             = "feImage"
	elementFeMerge             = "feMerge"
	elementFeMergeNode         = "feMergeNode"
	elementFeMorphology        = "feMorphology"
	elementFeOffset            = "feOffset"
	elementFePointLight        = "fePointLight"
	elementFeSpecularLighting  = "feSpecularLighting"
	elementFeSpotlight         = "feSpotlight"
	elementFeTile              = "feTile"
	elementFeTurbulence        = "feTurbulence"
	elementForeignObject       = "foreignObject"
	elementG                   = "g"
	elementImage               = "image"
	elementLine                = "line"
	elementLinearGradient      = "linearGradient"
	elementMarker              = "marker"
	elementMetadata            = "metadata"
	elementMpath               = "mpath"
	elementPolygon             = "polygon"
	elementPolyline            = "polyline"
	elementRadialGradient      = "radialGradient"
	elementRect                = "rect"
	elementSet                 = "set"
	elementStop                = "stop"
	elementSvg                 = "svg"
	elementSwitch              = "switch"
	elementSymbol              = "symbol"
	elementTextPath            = "textPath"
	elementTspan               = "tspan"
	elementView                = "view"
)

func Animate(nodes ...Node) Node {
	return createElement(elementAnimate, nodes...)
}

func AnimateMotion(nodes ...Node) Node {
	return createElement(elementAnimateMotion, nodes...)
}

func AnimateAnimateTransform(nodes ...Node) Node {
	return createElement(elementAnimateTransform, nodes...)
}

func Circle(nodes ...Node) Node {
	return createElement(elementCircle, nodes...)
}

func Defs(nodes ...Node) Node {
	return createElement(elementDefs, nodes...)
}

func Desc(nodes ...Node) Node {
	return createElement(elementDesc, nodes...)
}

func Ellipse(nodes ...Node) Node {
	return createElement(elementEllipse, nodes...)
}

func FeBlend(nodes ...Node) Node {
	return createElement(elementFeBlend, nodes...)
}

func FeColorMatrix(nodes ...Node) Node {
	return createElement(elementFeColorMatrix, nodes...)
}

func FeComponentTransfer(nodes ...Node) Node {
	return createElement(elementFeComponentTransfer, nodes...)
}

func FeComposite(nodes ...Node) Node {
	return createElement(elementFeComposite, nodes...)
}

func FeConvolveMatrix(nodes ...Node) Node {
	return createElement(elementFeConvolveMatrix, nodes...)
}

func FeDiffuseLightning(nodes ...Node) Node {
	return createElement(elementFeDiffuseLightning, nodes...)
}

func FeDisplacementMap(nodes ...Node) Node {
	return createElement(elementFeDisplacementMap, nodes...)
}

func FeDistantLight(nodes ...Node) Node {
	return createElement(elementFeDistantLight, nodes...)
}

func FeDropShadow(nodes ...Node) Node {
	return createElement(elementFeDropShadow, nodes...)
}

func FeFlood(nodes ...Node) Node {
	return createElement(elementFeFlood, nodes...)
}

func FeFuncA(nodes ...Node) Node {
	return createElement(elementFeFuncA, nodes...)
}

func FeFuncB(nodes ...Node) Node {
	return createElement(elementFeFuncB, nodes...)
}

func FeFuncG(nodes ...Node) Node {
	return createElement(elementFeFuncG, nodes...)
}

func FeFuncR(nodes ...Node) Node {
	return createElement(elementFeFuncR, nodes...)
}

func FeGaussianBlur(nodes ...Node) Node {
	return createElement(elementFeGaussianBlur, nodes...)
}

func FeImage(nodes ...Node) Node {
	return createElement(elementFeImage, nodes...)
}

func FeMerge(nodes ...Node) Node {
	return createElement(elementFeMerge, nodes...)
}

func FeMergeNode(nodes ...Node) Node {
	return createElement(elementFeMergeNode, nodes...)
}

func FeMorphology(nodes ...Node) Node {
	return createElement(elementFeMorphology, nodes...)
}

func FeOffset(nodes ...Node) Node {
	return createElement(elementFeOffset, nodes...)
}

func FePointLight(nodes ...Node) Node {
	return createElement(elementFePointLight, nodes...)
}

func FeSpecularLighting(nodes ...Node) Node {
	return createElement(elementFeSpecularLighting, nodes...)
}

func FeSpotlight(nodes ...Node) Node {
	return createElement(elementFeSpotlight, nodes...)
}

func FeTile(nodes ...Node) Node {
	return createElement(elementFeTile, nodes...)
}

func FeTurbulence(nodes ...Node) Node {
	return createElement(elementFeTurbulence, nodes...)
}

func ForeignObject(nodes ...Node) Node {
	return createElement(elementForeignObject, nodes...)
}

func G(nodes ...Node) Node {
	return createElement(elementG, nodes...)
}

func Image(nodes ...Node) Node {
	return createElement(elementImage, nodes...)
}

func Line(nodes ...Node) Node {
	return createElement(elementLine, nodes...)
}

func LinearGradient(nodes ...Node) Node {
	return createElement(elementLinearGradient, nodes...)
}

func Marker(nodes ...Node) Node {
	return createElement(elementMarker, nodes...)
}

func Metadata(nodes ...Node) Node {
	return createElement(elementMetadata, nodes...)
}

func Mpath(nodes ...Node) Node {
	return createElement(elementMpath, nodes...)
}

func Polygon(nodes ...Node) Node {
	return createElement(elementPolygon, nodes...)
}

func Polyline(nodes ...Node) Node {
	return createElement(elementPolyline, nodes...)
}

func RadialGradient(nodes ...Node) Node {
	return createElement(elementRadialGradient, nodes...)
}

func Rect(nodes ...Node) Node {
	return createElement(elementRect, nodes...)
}

func Set(nodes ...Node) Node {
	return createElement(elementSet, nodes...)
}

func Stop(nodes ...Node) Node {
	return createElement(elementStop, nodes...)
}

func Svg(nodes ...Node) Node {
	return createElement(elementSvg, Xmlns("http://www.w3.org/2000/svg"), Fragment(nodes...))
}

func Switch(nodes ...Node) Node {
	return createElement(elementSwitch, nodes...)
}

func Symbol(nodes ...Node) Node {
	return createElement(elementSymbol, nodes...)
}

func TextPath(nodes ...Node) Node {
	return createElement(elementTextPath, nodes...)
}

func Tspan(nodes ...Node) Node {
	return createElement(elementTspan, nodes...)
}

func View(nodes ...Node) Node {
	return createElement(elementView, nodes...)
}
