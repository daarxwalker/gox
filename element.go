package gox

const (
	doctypeHtml = "<!DOCTYPE html>"
)

const (
	elementA          = "a"
	elementAbbr       = "abbr"
	elementAddress    = "address"
	elementArea       = "area"
	elementArticle    = "article"
	elementAside      = "aside"
	elementAudio      = "audio"
	elementB          = "b"
	elementBase       = "base"
	elementBdi        = "bdi"
	elementBdo        = "bdo"
	elementBlockQuote = "blockquote"
	elementBody       = "body"
	elementBr         = "br"
	elementButton     = "button"
	elementCanvas     = "canvas"
	elementCaption    = "caption"
	elementCode       = "code"
	elementCol        = "col"
	elementColGroup   = "colgroup"
	elementDataList   = "datalist"
	elementDd         = "dd"
	elementDel        = "del"
	elementDetails    = "details"
	elementDfn        = "dfn"
	elementDialog     = "dialog"
	elementDl         = "dl"
	elementDt         = "dt"
	elementDiv        = "div"
	elementEm         = "em"
	elementEmbed      = "embed"
	elementFieldSet   = "fieldset"
	elementFigure     = "figure"
	elementFigCaption = "figcaption"
	elementFooter     = "footer"
	elementH1         = "h1"
	elementH2         = "h2"
	elementH3         = "h3"
	elementH4         = "h4"
	elementH5         = "h5"
	elementH6         = "h6"
	elementHead       = "head"
	elementHeader     = "header"
	elementHgroup     = "hgroup"
	elementHr         = "hr"
	elementHtml       = "html"
	elementIframe     = "iframe"
	elementImg        = "img"
	elementInput      = "input"
	elementI          = "i"
	elementIns        = "ins"
	elementKbd        = "kbd"
	elementLegend     = "legend"
	elementLi         = "li"
	elementLink       = "link"
	elementMain       = "main"
	elementMap        = "map"
	elementMark       = "mark"
	elementMath       = "math"
	elementMenu       = "menu"
	elementMeta       = "meta"
	elementMeter      = "meter"
	elementNav        = "nav"
	elementNoScript   = "noscript"
	elementObject     = "object"
	elementOl         = "ol"
	elementOptGroup   = "optgroup"
	elementOption     = "option"
	elementOutput     = "output"
	elementP          = "p"
	elementParam      = "param"
	elementPicture    = "picture"
	elementPortal     = "portal"
	elementPre        = "pre"
	elementProgress   = "progress"
	elementQ          = "q"
	elementRp         = "rp"
	elementRt         = "rt"
	elementRuby       = "ruby"
	elementS          = "s"
	elementSamp       = "samp"
	elementScript     = "script"
	elementSearch     = "search"
	elementSection    = "section"
	elementSelect     = "select"
	elementSmall      = "small"
	elementSource     = "source"
	elementStrong     = "strong"
	elementSub        = "sub"
	elementSummary    = "summary"
	elementSup        = "sup"
	elementTable      = "table"
	elementTbody      = "tbody"
	elementTd         = "td"
	elementTfoot      = "tfoot"
	elementTh         = "th"
	elementThead      = "thead"
	elementTr         = "tr"
	elementTemplate   = "template"
	elementTextArea   = "textarea"
	elementTime       = "time"
	elementTrack      = "track"
	elementU          = "u"
	elementUse        = "use"
	elementVar        = "var"
	elementVideo      = "video"
	elementWbr        = "wbr"
	elementUl         = "ul"
)

var (
	voidElements = []string{
		elementArea, elementBase, elementBr, elementCol, elementEmbed, elementHr, elementImg, elementInput, elementLink,
		elementMeta, elementParam, elementSource, elementTrack, elementWbr, sharedPath,
	}
)

func A(nodes ...Node) Node {
	return createElement(elementA, nodes...)
}

func Abbr(nodes ...Node) Node {
	return createElement(elementAbbr, nodes...)
}

func Address(nodes ...Node) Node {
	return createElement(elementAddress, nodes...)
}

func Area(nodes ...Node) Node {
	return createElement(elementArea, nodes...)
}

func Article(nodes ...Node) Node {
	return createElement(elementArticle, nodes...)
}

func Aside(nodes ...Node) Node {
	return createElement(elementAside, nodes...)
}

func Audio(nodes ...Node) Node {
	return createElement(elementAudio, nodes...)
}

func B(nodes ...Node) Node {
	return createElement(elementB, nodes...)
}

func Base(nodes ...Node) Node {
	return createElement(elementBase, nodes...)
}

func Bdi(nodes ...Node) Node {
	return createElement(elementBdi, nodes...)
}

func Bdo(nodes ...Node) Node {
	return createElement(elementBdo, nodes...)
}

func Blockquote(nodes ...Node) Node {
	return createElement(elementBlockQuote, nodes...)
}

func Body(nodes ...Node) Node {
	return createElement(elementBody, nodes...)
}

func Br() Node {
	return createElement(elementBr)
}

func Button(nodes ...Node) Node {
	return createElement(elementButton, nodes...)
}

func Canvas(nodes ...Node) Node {
	return createElement(elementCanvas, nodes...)
}

func Caption(nodes ...Node) Node {
	return createElement(elementCaption, nodes...)
}

func Code(nodes ...Node) Node {
	return createElement(elementCode, nodes...)
}

func ColGroup(nodes ...Node) Node {
	return createElement(elementColGroup, nodes...)
}

func Datalist(nodes ...Node) Node {
	return createElement(elementDataList, nodes...)
}

func Dd(nodes ...Node) Node {
	return createElement(elementDd, nodes...)
}

func Del(nodes ...Node) Node {
	return createElement(elementDel, nodes...)
}

func Details(nodes ...Node) Node {
	return createElement(elementDetails, nodes...)
}

func Dfn(nodes ...Node) Node {
	return createElement(elementDfn, nodes...)
}

func Dialog(nodes ...Node) Node {
	return createElement(elementDialog, nodes...)
}

func Div(nodes ...Node) Node {
	return createElement(elementDiv, nodes...)
}

func Dl(nodes ...Node) Node {
	return createElement(elementDl, nodes...)
}

func Dt(nodes ...Node) Node {
	return createElement(elementDt, nodes...)
}

func Em(nodes ...Node) Node {
	return createElement(elementEm, nodes...)
}

func Embed(nodes ...Node) Node {
	return createElement(elementEmbed, nodes...)
}

func Fieldset(nodes ...Node) Node {
	return createElement(elementFieldSet, nodes...)
}

func Figure(nodes ...Node) Node {
	return createElement(elementFigure, nodes...)
}

func FigCaption(nodes ...Node) Node {
	return createElement(elementFigCaption, nodes...)
}

func Footer(nodes ...Node) Node {
	return createElement(elementFooter, nodes...)
}

func H1(nodes ...Node) Node {
	return createElement(elementH1, nodes...)
}

func H2(nodes ...Node) Node {
	return createElement(elementH2, nodes...)
}

func H3(nodes ...Node) Node {
	return createElement(elementH3, nodes...)
}

func H4(nodes ...Node) Node {
	return createElement(elementH4, nodes...)
}

func H5(nodes ...Node) Node {
	return createElement(elementH5, nodes...)
}

func H6(nodes ...Node) Node {
	return createElement(elementH6, nodes...)
}

func Head(nodes ...Node) Node {
	return createElement(elementHead, nodes...)
}

func Header(nodes ...Node) Node {
	return createElement(elementHeader, nodes...)
}

func Hgroup(nodes ...Node) Node {
	return createElement(elementHgroup, nodes...)
}

func Html(nodes ...Node) Node {
	return createElement(elementHtml, nodes...)
}

func Hr() Node {
	return createElement(elementHr)
}

func Iframe(nodes ...Node) Node {
	return createElement(elementIframe, nodes...)
}

func Img(nodes ...Node) Node {
	return createElement(elementImg, nodes...)
}

func Input(nodes ...Node) Node {
	return createElement(elementInput, nodes...)
}

func I(nodes ...Node) Node {
	return createElement(elementI, nodes...)
}

func Ins(nodes ...Node) Node {
	return createElement(elementIns, nodes...)
}

func Kbd(nodes ...Node) Node {
	return createElement(elementKbd, nodes...)
}

func Legend(nodes ...Node) Node {
	return createElement(elementLegend, nodes...)
}

func Li(nodes ...Node) Node {
	return createElement(elementLi, nodes...)
}

func Link(nodes ...Node) Node {
	return createElement(elementLink, nodes...)
}

func Main(nodes ...Node) Node {
	return createElement(elementMain, nodes...)
}

func Map(nodes ...Node) Node {
	return createElement(elementMap, nodes...)
}

func Mark(nodes ...Node) Node {
	return createElement(elementMark, nodes...)
}

func Math(nodes ...Node) Node {
	return createElement(elementMath, nodes...)
}

func Menu(nodes ...Node) Node {
	return createElement(elementMenu, nodes...)
}

func Meta(nodes ...Node) Node {
	return createElement(elementMeta, nodes...)
}

func Meter(nodes ...Node) Node {
	return createElement(elementMeter, nodes...)
}

func Nav(nodes ...Node) Node {
	return createElement(elementNav, nodes...)
}

func NoScript(nodes ...Node) Node {
	return createElement(elementNoScript, nodes...)
}

func Object(nodes ...Node) Node {
	return createElement(elementObject, nodes...)
}

func Ol(nodes ...Node) Node {
	return createElement(elementOl, nodes...)
}

func Optgroup(nodes ...Node) Node {
	return createElement(elementOptGroup, nodes...)
}

func Option(nodes ...Node) Node {
	return createElement(elementOption, nodes...)
}

func Output(nodes ...Node) Node {
	return createElement(elementOutput, nodes...)
}

func P(nodes ...Node) Node {
	return createElement(elementP, nodes...)
}

func Param(nodes ...Node) Node {
	return createElement(elementParam, nodes...)
}

func Picture(nodes ...Node) Node {
	return createElement(elementPicture, nodes...)
}

func Portal(nodes ...Node) Node {
	return createElement(elementPortal, nodes...)
}

func Pre(nodes ...Node) Node {
	return createElement(elementPre, nodes...)
}

func Progress(nodes ...Node) Node {
	return createElement(elementProgress, nodes...)
}

func Q(nodes ...Node) Node {
	return createElement(elementQ, nodes...)
}

func Rp(nodes ...Node) Node {
	return createElement(elementRp, nodes...)
}

func Rt(nodes ...Node) Node {
	return createElement(elementRt, nodes...)
}

func Ruby(nodes ...Node) Node {
	return createElement(elementRuby, nodes...)
}

func S(nodes ...Node) Node {
	return createElement(elementS, nodes...)
}

func Samp(nodes ...Node) Node {
	return createElement(elementSamp, nodes...)
}

func Script(nodes ...Node) Node {
	return createElement(elementScript, nodes...)
}

func Search(nodes ...Node) Node {
	return createElement(elementSearch, nodes...)
}

func Section(nodes ...Node) Node {
	return createElement(elementSection, nodes...)
}

func Select(nodes ...Node) Node {
	return createElement(elementSelect, nodes...)
}

func Small(nodes ...Node) Node {
	return createElement(elementSmall, nodes...)
}

func Source(nodes ...Node) Node {
	return createElement(elementSource, nodes...)
}

func Strong(nodes ...Node) Node {
	return createElement(elementStrong, nodes...)
}

func Sub(nodes ...Node) Node {
	return createElement(elementSub, nodes...)
}

func Summary(nodes ...Node) Node {
	return createElement(elementSummary, nodes...)
}

func Sup(nodes ...Node) Node {
	return createElement(elementSup, nodes...)
}

func Table(nodes ...Node) Node {
	return createElement(elementTable, nodes...)
}

func Tbody(nodes ...Node) Node {
	return createElement(elementTbody, nodes...)
}

func Td(nodes ...Node) Node {
	return createElement(elementTd, nodes...)
}

func Tfoot(nodes ...Node) Node {
	return createElement(elementTfoot, nodes...)
}

func Th(nodes ...Node) Node {
	return createElement(elementTh, nodes...)
}

func Thead(nodes ...Node) Node {
	return createElement(elementThead, nodes...)
}

func Tr(nodes ...Node) Node {
	return createElement(elementTr, nodes...)
}

func Template(nodes ...Node) Node {
	return createElement(elementTemplate, nodes...)
}

func Textarea(nodes ...Node) Node {
	return createElement(elementTextArea, nodes...)
}

func Time(nodes ...Node) Node {
	return createElement(elementTime, nodes...)
}

func U(nodes ...Node) Node {
	return createElement(elementU, nodes...)
}

func Use(nodes ...Node) Node {
	return createElement(elementUse, nodes...)
}

func Var(nodes ...Node) Node {
	return createElement(elementVar, nodes...)
}

func Video(nodes ...Node) Node {
	return createElement(elementVideo, nodes...)
}

func Wbr() Node {
	return createElement(elementWbr)
}

func Ul(nodes ...Node) Node {
	return createElement(elementUl, nodes...)
}
