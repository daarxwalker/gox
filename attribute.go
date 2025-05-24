package gox

import (
	"fmt"
	"strings"
	
	"golang.org/x/exp/constraints"
)

const (
	attributeAccept          = "accept"
	attributeAcceptCharset   = "accept-charset"
	attributeAccessKey       = "access-key"
	attributeAction          = "action"
	attributeAllow           = "allow"
	attributeAlt             = "alt"
	attributeAsync           = "async"
	attributeAutoCapitalize  = "autocapitalize"
	attributeAutoComplete    = "autocomplete"
	attributeAutoFocus       = "autofocus"
	attributeAutoPlay        = "autoplay"
	attributeBackground      = "background"
	attributeBg              = "bgcolor"
	attributeBorder          = "border"
	attributeBuffered        = "buffered"
	attributeCapture         = "capture"
	attributeCharSet         = "charset"
	attributeChecked         = "checked"
	attributeClass           = "class"
	attributeCols            = "cols"
	attributeColSpan         = "colspan"
	attributeContent         = "content"
	attributeContentEditable = "contenteditable"
	attributeControls        = "controls"
	attributeCoords          = "coords"
	attributeCsp             = "csp"
	attributeDateTime        = "datetime"
	attributeDecoding        = "decoding"
	attributeDefault         = "default"
	attributeDefer           = "defer"
	attributeDir             = "dir"
	attributeDirName         = "dirname"
	attributeDisable         = "disabled"
	attributeDownload        = "download"
	attributeDraggable       = "draggable"
	attributeEncType         = "enctype"
	attributeEnterKeyHint    = "enterkeyhint"
	attributeFor             = "for"
	attributeFormAction      = "formaction"
	attributeFormEncType     = "formenctype"
	attributeFormMethod      = "formmethod"
	attributeFormNoValidate  = "formnovalidate"
	attributeFormTarget      = "formtarget"
	attributeHeaders         = "headers"
	attributeHeight          = "height"
	attributeHidden          = "hidden"
	attributeHigh            = "high"
	attributeHref            = "href"
	attributeHrefLang        = "hreflang"
	attributeHttpEquiv       = "http-equiv"
	attributeId              = "id"
	attributeIntegrity       = "integrity"
	attributeInputMode       = "inputmode"
	attributeIsMap           = "ismap"
	attributeItemId          = "itemid"
	attributeItemProp        = "itemprop"
	attributeItemRef         = "itemref"
	attributeItemScope       = "itemscope"
	attributeItemType        = "itemtype"
	attributeKind            = "kind"
	attributeLang            = "lang"
	attributeLoading         = "loading"
	attributeList            = "list"
	attributeLoop            = "loop"
	attributeLow             = "low"
	attributeMax             = "max"
	attributeMaxLength       = "maxlength"
	attributeMinLength       = "minlength"
	attributeMedia           = "media"
	attributeMethod          = "method"
	attributeMin             = "min"
	attributeMultiple        = "multiple"
	attributeMuted           = "muted"
	attributeName            = "name"
	attributeNonce           = "nonce"
	attributeNoValidate      = "novalidate"
	attributeOpen            = "open"
	attributeOptimum         = "optimum"
	attributePing            = "ping"
	attributePlaceholder     = "placeholder"
	attributePlaysInline     = "playsinline"
	attributePoster          = "poster"
	attributePreload         = "preload"
	attributeReadonly        = "readonly"
	attributeReferrerPolicy  = "referrerpolicy"
	attributeRel             = "rel"
	attributeRequired        = "required"
	attributeReversed        = "reversed"
	attributeRole            = "role"
	attributeRows            = "rows"
	attributeRowSpan         = "rowspan"
	attributeSandbox         = "sandbox"
	attributeScope           = "scope"
	attributeSelected        = "selected"
	attributeShape           = "shape"
	attributeSize            = "size"
	attributeSizes           = "sizes"
	attributeSpellCheck      = "spellcheck"
	attributeSrc             = "src"
	attributeSrcDoc          = "srcdoc"
	attributeSrcLang         = "srclang"
	attributeSrcSet          = "srcset"
	attributeStart           = "start"
	attributeStep            = "step"
	attributeTabIndex        = "tabindex"
	attributeTarget          = "target"
	attributeTranslate       = "translate"
	attributeType            = "type"
	attributeUseMap          = "usemap"
	attributeValue           = "value"
	attributeVersion         = "version"
	attributeWidth           = "width"
	attributeWrap            = "wrap"
)

func Accept(values ...string) Node {
	return createAttribute(attributeAccept, values...)
}

func AcceptCharset(values ...string) Node {
	return createAttribute(attributeAcceptCharset, values...)
}

func AccessKey(values ...string) Node {
	return createAttribute(attributeAccessKey, values...)
}

func Action(values ...string) Node {
	return createAttribute(attributeAction, values...)
}

func Allow(values ...string) Node {
	return createAttribute(attributeAllow, values...)
}

func Alt(values ...string) Node {
	return createAttribute(attributeAlt, values...)
}

func Async() Node {
	return createAttribute[any](attributeAsync)
}

func AutoCapitalize(values ...bool) Node {
	return createAttribute(attributeAutoCapitalize, values...)
}

func AutoComplete(values ...string) Node {
	return createAttribute(attributeAutoComplete, values...)
}

func AutoFocus(values ...bool) Node {
	return createAttribute(attributeAutoFocus, values...)
}

func AutoPlay(values ...bool) Node {
	return createAttribute(attributeAutoPlay, values...)
}

func Background(values ...string) Node {
	return createAttribute(attributeBackground, values...)
}

func Bg(values ...string) Node {
	return createAttribute(attributeBg, values...)
}

func Border(values ...string) Node {
	return createAttribute(attributeBorder, values...)
}

func Buffered(values ...bool) Node {
	return createAttribute(attributeBuffered, values...)
}

func Capture(values ...string) Node {
	return createAttribute(attributeCapture, values...)
}

func CharSet(values ...string) Node {
	return createAttribute(attributeCharSet, values...)
}

func Checked(values ...bool) Node {
	return createAttribute(attributeChecked, values...)
}

func Class(value any) Node {
	switch v := value.(type) {
	case string:
		return createAttribute(attributeClass, v)
	case fmt.Stringer:
		return createAttribute(attributeClass, v.String())
	}
	return createAttribute(attributeClass, fmt.Sprint(value))
}

func Cols(values ...int) Node {
	return createAttribute(attributeCols, values...)
}

func ColSpan(values ...int) Node {
	return createAttribute(attributeColSpan, values...)
}

func Content(values ...string) Node {
	return createAttribute(attributeContent, values...)
}

func ContentEditable(values ...bool) Node {
	return createAttribute(attributeContentEditable, values...)
}

func Controls(values ...bool) Node {
	return createAttribute(attributeControls, values...)
}

func Coords(values ...string) Node {
	return createAttribute(attributeCoords, strings.Join(values, ","))
}

func Csp(values ...string) Node {
	return createAttribute(attributeCsp, values...)
}

func DateTime(values ...string) Node {
	return createAttribute(attributeDateTime, values...)
}

func Decoding(values ...string) Node {
	return createAttribute(attributeDecoding, values...)
}

func Default(values ...bool) Node {
	return createAttribute(attributeDefault, values...)
}

func Defer() Node {
	return createAttribute[any](attributeDefer)
}

func DirName(values ...string) Node {
	return createAttribute(attributeDirName, values...)
}

func Disabled(values ...bool) Node {
	return createAttribute(attributeDisable, values...)
}

func Download(values ...bool) Node {
	return createAttribute(attributeDownload, values...)
}

func Draggable(values ...bool) Node {
	return createAttribute(attributeDraggable, values...)
}

func Dir(values ...string) Node {
	return createAttribute(attributeDir, values...)
}

func EncType(values ...string) Node {
	return createAttribute(attributeEncType, values...)
}

func EnterKeyHint(values ...string) Node {
	return createAttribute(attributeEnterKeyHint, values...)
}

func For(values ...string) Node {
	return createAttribute(attributeFor, values...)
}

func FormAction(values ...string) Node {
	return createAttribute(attributeFormAction, values...)
}

func FormEncType(values ...string) Node {
	return createAttribute(attributeFormEncType, values...)
}

func FormMethod(values ...string) Node {
	return createAttribute(attributeFormMethod, values...)
}

func FormNoValidate(values ...string) Node {
	return createAttribute(attributeFormNoValidate, values...)
}

func FormTarget(values ...string) Node {
	return createAttribute(attributeFormTarget, values...)
}

func Headers(values ...string) Node {
	return createAttribute(attributeHeaders, values...)
}

func Height[T constraints.Ordered](values ...T) Node {
	return createAttribute(attributeHeight, values...)
}

func Hidden(values ...bool) Node {
	return createAttribute(attributeHidden, values...)
}

func High[T constraints.Ordered](values ...T) Node {
	return createAttribute(attributeHigh, values...)
}

func Href(values ...string) Node {
	return createAttribute(attributeHref, strings.Join(values, "/"))
}

func HrefLang(values ...string) Node {
	return createAttribute(attributeHrefLang, values...)
}

func HttpEquiv(values ...string) Node {
	return createAttribute(attributeHttpEquiv, values...)
}

func Id(values ...string) Node {
	return createAttribute(attributeId, values...)
}

func Integrity(values ...string) Node {
	return createAttribute(attributeIntegrity, values...)
}

func InputMode(values ...bool) Node {
	return createAttribute(attributeInputMode, values...)
}

func IsMap(values ...bool) Node {
	return createAttribute(attributeIsMap, values...)
}

func ItemId(values ...string) Node {
	return createAttribute(attributeItemId, values...)
}

func ItemProp(values ...string) Node {
	return createAttribute(attributeItemProp, values...)
}

func ItemRef(values ...string) Node {
	return createAttribute(attributeItemRef, values...)
}

func ItemScope(values ...string) Node {
	return createAttribute(attributeItemScope, values...)
}

func ItemType(values ...string) Node {
	return createAttribute(attributeItemType, values...)
}

func Kind(values ...string) Node {
	return createAttribute(attributeKind, values...)
}

func Lang(values ...string) Node {
	return createAttribute(attributeLang, values...)
}

func Loading(values ...string) Node {
	return createAttribute(attributeLoading, values...)
}

func List(values ...string) Node {
	return createAttribute(attributeList, values...)
}

func Loop(values ...bool) Node {
	return createAttribute(attributeLoop, values...)
}

func Low[T constraints.Ordered](values ...T) Node {
	return createAttribute(attributeLow, values...)
}

func Max[T constraints.Ordered](values ...T) Node {
	return createAttribute(attributeMax, values...)
}

func MaxLength[T constraints.Ordered](values ...T) Node {
	return createAttribute(attributeMaxLength, values...)
}

func Media(values ...string) Node {
	return createAttribute(attributeMedia, values...)
}

func Method(values ...string) Node {
	return createAttribute(attributeMethod, values...)
}

func MinLength[T constraints.Ordered](values ...T) Node {
	return createAttribute(attributeMinLength, values...)
}

func Min[T constraints.Ordered](values ...T) Node {
	return createAttribute(attributeMin, values...)
}

func Multiple(values ...bool) Node {
	return createAttribute(attributeMultiple, values...)
}

func Muted(values ...bool) Node {
	return createAttribute(attributeMuted, values...)
}

func Name(values ...string) Node {
	return createAttribute(attributeName, values...)
}

func Nonce(values ...string) Node {
	return createAttribute(attributeNonce, values...)
}

func NoValidate(values ...bool) Node {
	return createAttribute(attributeNoValidate, values...)
}

func Open(values ...bool) Node {
	return createAttribute(attributeOpen, values...)
}

func Optimum[T constraints.Ordered](values ...T) Node {
	return createAttribute(attributeOptimum, values...)
}

func Ping(values ...string) Node {
	return createAttribute(attributePing, values...)
}

func Placeholder(values ...string) Node {
	return createAttribute(attributePlaceholder, values...)
}

func PlaysInline(values ...bool) Node {
	return createAttribute(attributePlaysInline, values...)
}

func Poster(values ...string) Node {
	return createAttribute(attributePoster, values...)
}

func Preload(values ...bool) Node {
	return createAttribute(attributePreload, values...)
}

func Readonly(values ...bool) Node {
	return createAttribute(attributeReadonly, values...)
}

func ReferrerPolicy(values ...string) Node {
	return createAttribute(attributeReferrerPolicy, values...)
}

func Rel(values ...string) Node {
	return createAttribute(attributeRel, values...)
}

func Required(values ...bool) Node {
	return createAttribute(attributeRequired, values...)
}

func Reversed(values ...bool) Node {
	return createAttribute(attributeReversed, values...)
}

func Role(values ...string) Node {
	return createAttribute(attributeRole, values...)
}

func Rows(values ...int) Node {
	return createAttribute(attributeRows, values...)
}

func RowSpan(values ...int) Node {
	return createAttribute(attributeRowSpan, values...)
}

func Sandbox(values ...string) Node {
	return createAttribute(attributeSandbox, values...)
}

func Scope(values ...string) Node {
	return createAttribute(attributeScope, values...)
}

func Selected(values ...bool) Node {
	return createAttribute(attributeSelected, values...)
}

func Shape(values ...string) Node {
	return createAttribute(attributeShape, values...)
}

func Size[T constraints.Ordered](values ...T) Node {
	return createAttribute(attributeSize, values...)
}

func Sizes(values ...string) Node {
	return createAttribute(attributeSizes, values...)
}

func SpellCheck(values ...bool) Node {
	return createAttribute(attributeSpellCheck, values...)
}

func Src(values ...string) Node {
	return createAttribute(attributeSrc, values...)
}

func SrcDoc(values ...string) Node {
	return createAttribute(attributeSrcDoc, values...)
}

func SrcLang(values ...string) Node {
	return createAttribute(attributeSrcLang, values...)
}

func SrcSet(values ...string) Node {
	return createAttribute(attributeSrcSet, values...)
}

func Start(values ...int) Node {
	return createAttribute(attributeStart, values...)
}

func Step(values ...string) Node {
	return createAttribute(attributeStep, values...)
}

func TabIndex(values ...int) Node {
	return createAttribute(attributeTabIndex, values...)
}

func Target(values ...string) Node {
	return createAttribute(attributeTarget, values...)
}

func Translate(values ...string) Node {
	return createAttribute(attributeTranslate, values...)
}

func Type(values ...string) Node {
	return createAttribute(attributeType, values...)
}

func UseMap(values ...string) Node {
	return createAttribute(attributeUseMap, values...)
}

func Value(values ...any) Node {
	return createAttribute(attributeValue, values...)
}

func Version(values ...string) Node {
	return createAttribute(attributeVersion, values...)
}

func Width[T constraints.Ordered](values ...T) Node {
	return createAttribute(attributeWidth, values...)
}

func Wrap(values ...string) Node {
	return createAttribute(attributeWrap, values...)
}
