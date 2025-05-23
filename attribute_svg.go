package gox

import "strings"

const (
	attributeAccumulate                = "accumulate"
	attributeAdditive                  = "additive"
	attributeAlignmentBaseline         = "alignment-baseline"
	attributeAmplitude                 = "amplitude"
	attributeAttributeName             = "attributeName"
	attributeAzimuth                   = "azimuth"
	attributeBaseFrequency             = "baseFrequency"
	attributeBaselineShift             = "baseline-shift"
	attributeBegin                     = "begin"
	attributeBias                      = "bias"
	attributeBy                        = "by"
	attributeCalcMode                  = "calcMode"
	attributeClipPathUnits             = "clipPathUnits"
	attributeClipRule                  = "clip-rule"
	attributeColor                     = "color"
	attributeColorInterpolation        = "color-interpolation"
	attributeColorInterpolationFilters = "color-interpolation-filters"
	attributeCrossorigin               = "crossorigin"
	attributeCursor                    = "cursor"
	attributeCx                        = "cx"
	attributeCy                        = "cy"
	attributeD                         = "d"
	attributeDiffuseConstant           = "diffuseConstant"
	attributeDirection                 = "direction"
	attributeDisplay                   = "display"
	attributeDivisor                   = "divisor"
	attributeDominantBaseline          = "dominant-baseline"
	attributeDur                       = "dur"
	attributeDx                        = "dx"
	attributeDy                        = "dy"
	attributeEdgeMode                  = "edgeMode"
	attributeElevation                 = "elevation"
	attributeEnd                       = "end"
	attributeExponent                  = "exponent"
	attributeFill                      = "fill"
	attributeFillOpacity               = "fill-opacity"
	attributeFillRule                  = "fill-rule"
	attributeFilterUnits               = "fitlerUnits"
	attributeFloodColor                = "flood-color"
	attributeFloodOpacity              = "flood-opacity"
	attributeFontFamily                = "font-family"
	attributeFontSize                  = "font-size"
	attributeFontSizeAdjust            = "font-size-adjust"
	attributeFontStretch               = "font-stretch"
	attributeFontStyle                 = "font-style"
	attributeFontVariant               = "font-variant"
	attributeFontWeight                = "font-weight"
	attributeFrom                      = "from"
	attributeFr                        = "fr"
	attributeFx                        = "fx"
	attributeFy                        = "fy"
	attributeGradientTransform         = "gradientTransform"
	attributeGradientUnits             = "gradientUnits"
	attributeImageRendering            = "image-rendering"
	attributeIn                        = "in"
	attributeIn2                       = "in2"
	attributeIntercept                 = "intercept"
	attributeK1                        = "k1"
	attributeK2                        = "k2"
	attributeK3                        = "k3"
	attributeK4                        = "k4"
	attributeKernelMatrix              = "kernelMatrix"
	attributeKeyPoints                 = "keyPoints"
	attributeKeySplines                = "keySplines"
	attributeKeyTimes                  = "keyTimes"
	attributeLengthAdjust              = "lengthAdjust"
	attributeLetterSpacing             = "letter-spacing"
	attributeLightingColor             = "lighting-color"
	attributeLimitingConeAngle         = "limitingConeAngle"
	attributeMarkerEnd                 = "marker-end"
	attributeMarkerMid                 = "marker-mid"
	attributeMarkerStart               = "marker-start"
	attributeMarkerHeight              = "markerHeight"
	attributeMarkerUnits               = "markerUnits"
	attributeMarkerWidth               = "markerWidth"
	attributeMaskContentUnits          = "maskContentUnits"
	attributeMaskUnits                 = "maskUnits"
	attributeMode                      = "mode"
	attributeNumOctaves                = "numOctaves"
	attributeOpacity                   = "opacity"
	attributeOperator                  = "operator"
	attributeOrder                     = "order"
	attributeOrient                    = "orient"
	attributeOrigin                    = "origin"
	attributeOverflow                  = "overflow"
	attributeOverlinePosition          = "overlinePosition"
	attributeOverlineThickness         = "overlineThickness"
	attributePaintOrder                = "paint-order"
	attributePathLength                = "pathLength "
	attributePatternContentUnits       = "patternContentUnits"
	attributePatternTransform          = "patternTransform"
	attributePatternUnits              = "patternUnits"
	attributePointerEvents             = "pointer-events"
	attributePoints                    = "points"
	attributePointsAtX                 = "pointsAtX"
	attributePointsAtY                 = "pointsAtY"
	attributePointsAtZ                 = "pointsAtZ"
	attributePreserveAlpha             = "preserveAlpha"
	attributePreserveAspectRatio       = "preserveAspectRatio"
	attributePrimitiveUnits            = "primitiveUnits"
	attributeR                         = "r"
	attributeRadius                    = "radius"
	attributeRefX                      = "refX"
	attributeRefY                      = "refY"
	attributeRepeatCount               = "repeatCount"
	attributeRepeatDur                 = "repeatDur"
	attributeRestart                   = "restart"
	attributeResult                    = "result"
	attributeRotate                    = "rotate"
	attributeRx                        = "rx"
	attributeRy                        = "ry"
	attributeScale                     = "scale"
	attributeSeed                      = "seed"
	attributeShapeRendering            = "shape-rendering"
	attributeSpacing                   = "spacing"
	attributeSpecularConstant          = "specularConstant"
	attributeSpecularExponent          = "specularExponent"
	attributeSpreadMethod              = "spreadMethod"
	attributeStartOffset               = "startOffset"
	attributeStdDeviation              = "stdDeviation"
	attributeStitchTiles               = "stitchTiles"
	attributeStopColor                 = "stop-color"
	attributeStopOpacity               = "stop-opacity"
	attributeStrikethroughPosition     = "strikethrough-position"
	attributeStrikethroughThickness    = "strikethrough-thickness"
	attributeStroke                    = "stroke"
	attributeStrokeDasharray           = "stroke-dasharray"
	attributeStrokeDashoffset          = "stroke-dashoffset"
	attributeStrokeLinecap             = "stroke-linecap"
	attributeStrokeLinejoin            = "stroke-linejoin"
	attributeStrokeMiterlimit          = "stroke-miterlimit"
	attributeStrokeOpacity             = "stroke-opacity"
	attributeStrokeWidth               = "stroke-width"
	attributeSurfaceScale              = "surfaceScale"
	attributeSystemLanguage            = "systemLanguage"
	attributeTableValues               = "tableValues"
	attributeTargetX                   = "targetX"
	attributeTargetY                   = "targetY"
	attributeTextAnchor                = "text-anchor"
	attributeTextDecoration            = "text-decoration"
	attributeTextRendering             = "text-rendering"
	attributeTextLength                = "textLength"
	attributeTo                        = "to"
	attributeTransform                 = "transform"
	attributeTransformOrigin           = "transform-origin"
	attributeUnderlinePosition         = "underline-position"
	attributeUnderlineThickness        = "underline-thickness"
	attributeUnicodeBidi               = "unicode-bidi"
	attributeValues                    = "values"
	attributeVectorEffect              = "vector-effect"
	attributeViewBox                   = "viewBox"
	attributeVisibility                = "visibility"
	attributeWordSpacing               = "word-spacing"
	attributeWritingMode               = "writing-mode"
	attributeX                         = "x"
	attributeX1                        = "x1"
	attributeX2                        = "x2"
	attributeXChannelSelector          = "xChannelSelector"
	attributeXmlns                     = "xmlns"
	attributeY                         = "y"
	attributeY1                        = "y1"
	attributeY2                        = "y2"
	attributeYChannelSelector          = "yChannelSelector"
	attributeZ                         = "z"
)

func Accumulate(values ...string) Node {
	return createAttribute(attributeAccumulate, values...)
}

func Additive(values ...string) Node {
	return createAttribute(attributeAdditive, values...)
}

func AlignmentBaseline(values ...string) Node {
	return createAttribute(attributeAlignmentBaseline, values...)
}

func Amplitude(values ...string) Node {
	return createAttribute(attributeAmplitude, values...)
}

func AttributeName(values ...string) Node {
	return createAttribute(attributeAttributeName, values...)
}

func Azimuth(values ...string) Node {
	return createAttribute(attributeAzimuth, values...)
}

func BaseFrequency(values ...string) Node {
	return createAttribute(attributeBaseFrequency, values...)
}

func BaselineShift(values ...string) Node {
	return createAttribute(attributeBaselineShift, values...)
}

func Begin(values ...string) Node {
	return createAttribute(attributeBegin, values...)
}

func Bias(values ...string) Node {
	return createAttribute(attributeBias, values...)
}

func By(values ...string) Node {
	return createAttribute(attributeBy, values...)
}

func CalcMode(values ...string) Node {
	return createAttribute(attributeCalcMode, values...)
}

func ClipPathUnits(values ...string) Node {
	return createAttribute(attributeClipPathUnits, values...)
}

func ClipRule(values ...string) Node {
	return createAttribute(attributeClipRule, values...)
}

func Color(values ...string) Node {
	return createAttribute(attributeColor, values...)
}

func ColorInterpolation(values ...string) Node {
	return createAttribute(attributeColorInterpolation, values...)
}

func ColorInterpolationFilters(values ...string) Node {
	return createAttribute(attributeColorInterpolationFilters, values...)
}

func CrossOrigin(values ...string) Node {
	return createAttribute(attributeCrossorigin, values...)
}

func Cursor(values ...string) Node {
	return createAttribute(attributeCursor, values...)
}

func Cx(values ...string) Node {
	return createAttribute(attributeCx, values...)
}

func Cy(values ...string) Node {
	return createAttribute(attributeCy, values...)
}

func D(values ...string) Node {
	return createAttribute(attributeD, values...)
}

func DiffuseConstant(values ...string) Node {
	return createAttribute(attributeDiffuseConstant, values...)
}

func Direction(values ...string) Node {
	return createAttribute(attributeDirection, values...)
}

func Display(values ...string) Node {
	return createAttribute(attributeDisplay, values...)
}

func Divisor(values ...string) Node {
	return createAttribute(attributeDivisor, values...)
}

func DominantBaseline(values ...string) Node {
	return createAttribute(attributeDominantBaseline, values...)
}

func Dur(values ...string) Node {
	return createAttribute(attributeDur, values...)
}

func Dx(values ...string) Node {
	return createAttribute(attributeDx, values...)
}

func Dy(values ...string) Node {
	return createAttribute(attributeDy, values...)
}

func EdgeMode(values ...string) Node {
	return createAttribute(attributeEdgeMode, values...)
}

func Elevation(values ...string) Node {
	return createAttribute(attributeElevation, values...)
}

func End(values ...string) Node {
	return createAttribute(attributeEnd, values...)
}

func Exponent(values ...string) Node {
	return createAttribute(attributeExponent, values...)
}

func Fill(values ...string) Node {
	return createAttribute(attributeFill, values...)
}

func FillOpacity(values ...string) Node {
	return createAttribute(attributeFillOpacity, values...)
}

func FillRule(values ...string) Node {
	return createAttribute(attributeFillRule, values...)
}

func FilterUnits(values ...string) Node {
	return createAttribute(attributeFilterUnits, values...)
}

func FloodColor(values ...string) Node {
	return createAttribute(attributeFloodColor, values...)
}

func FloodOpacity(values ...string) Node {
	return createAttribute(attributeFloodOpacity, values...)
}

func FontFamily(values ...string) Node {
	return createAttribute(attributeFontFamily, values...)
}

func TextSize(values ...string) Node {
	return createAttribute(attributeFontSize, values...)
}

func FontSizeAdjust(values ...string) Node {
	return createAttribute(attributeFontSizeAdjust, values...)
}

func FontStretch(values ...string) Node {
	return createAttribute(attributeFontStretch, values...)
}

func FontStyle(values ...string) Node {
	return createAttribute(attributeFontStyle, values...)
}

func FontVariant(values ...string) Node {
	return createAttribute(attributeFontVariant, values...)
}

func FontWeight(values ...string) Node {
	return createAttribute(attributeFontWeight, values...)
}

func From(values ...string) Node {
	return createAttribute(attributeFrom, values...)
}

func Fr(values ...string) Node {
	return createAttribute(attributeFr, values...)
}

func Fx(values ...string) Node {
	return createAttribute(attributeFx, values...)
}

func Fy(values ...string) Node {
	return createAttribute(attributeFy, values...)
}

func GradientTransform(values ...string) Node {
	return createAttribute(attributeGradientTransform, values...)
}

func GradientUnits(values ...string) Node {
	return createAttribute(attributeGradientUnits, values...)
}

func ImageRendering(values ...string) Node {
	return createAttribute(attributeImageRendering, values...)
}

func In(values ...string) Node {
	return createAttribute(attributeIn, values...)
}

func In2(values ...string) Node {
	return createAttribute(attributeIn2, values...)
}

func Intercept(values ...string) Node {
	return createAttribute(attributeIntercept, values...)
}

func K1(values ...string) Node {
	return createAttribute(attributeK1, values...)
}

func K2(values ...string) Node {
	return createAttribute(attributeK2, values...)
}

func K3(values ...string) Node {
	return createAttribute(attributeK3, values...)
}

func K4(values ...string) Node {
	return createAttribute(attributeK4, values...)
}

func KernelMatrix(values ...string) Node {
	return createAttribute(attributeKernelMatrix, values...)
}

func KeyPoints(values ...string) Node {
	return createAttribute(attributeKeyPoints, values...)
}

func KeySplines(values ...string) Node {
	return createAttribute(attributeKeySplines, values...)
}

func KeyTimes(values ...string) Node {
	return createAttribute(attributeKeyTimes, values...)
}

func LengthAdjust(values ...string) Node {
	return createAttribute(attributeLengthAdjust, values...)
}

func LetterSpacing(values ...string) Node {
	return createAttribute(attributeLetterSpacing, values...)
}

func LightingColor(values ...string) Node {
	return createAttribute(attributeLightingColor, values...)
}

func LimitingConeAngle(values ...string) Node {
	return createAttribute(attributeLimitingConeAngle, values...)
}

func MarkerEnd(values ...string) Node {
	return createAttribute(attributeMarkerEnd, values...)
}

func MarkerMid(values ...string) Node {
	return createAttribute(attributeMarkerMid, values...)
}

func MarkerStart(values ...string) Node {
	return createAttribute(attributeMarkerStart, values...)
}

func MarkerHeight(values ...string) Node {
	return createAttribute(attributeMarkerHeight, values...)
}

func MarkerUnits(values ...string) Node {
	return createAttribute(attributeMarkerUnits, values...)
}

func MarkerWidth(values ...string) Node {
	return createAttribute(attributeMarkerWidth, values...)
}

func MaskContentUnits(values ...string) Node {
	return createAttribute(attributeMaskContentUnits, values...)
}

func MaskUnits(values ...string) Node {
	return createAttribute(attributeMaskUnits, values...)
}

func Mode(values ...string) Node {
	return createAttribute(attributeMode, values...)
}

func NumOctaves(values ...string) Node {
	return createAttribute(attributeNumOctaves, values...)
}

func Opacity(values ...string) Node {
	return createAttribute(attributeOpacity, values...)
}

func Operator(values ...string) Node {
	return createAttribute(attributeOperator, values...)
}

func Order(values ...string) Node {
	return createAttribute(attributeOrder, values...)
}

func Orient(values ...string) Node {
	return createAttribute(attributeOrient, values...)
}

func Origin(values ...string) Node {
	return createAttribute(attributeOrigin, values...)
}

func Overflow(values ...string) Node {
	return createAttribute(attributeOverflow, values...)
}

func OverlinePosition(values ...string) Node {
	return createAttribute(attributeOverlinePosition, values...)
}

func OverlineThickness(values ...string) Node {
	return createAttribute(attributeOverlineThickness, values...)
}

func PaintOrder(values ...string) Node {
	return createAttribute(attributePaintOrder, values...)
}

func PathLength(values ...string) Node {
	return createAttribute(attributePathLength, values...)
}

func PatternContentUnits(values ...string) Node {
	return createAttribute(attributePatternContentUnits, values...)
}

func PatternTransform(values ...string) Node {
	return createAttribute(attributePatternTransform, values...)
}

func PatternUnits(values ...string) Node {
	return createAttribute(attributePatternUnits, values...)
}

func PointerEvents(values ...string) Node {
	return createAttribute(attributePointerEvents, values...)
}

func Points(values ...string) Node {
	return createAttribute(attributePoints, values...)
}

func PointsAtX(values ...string) Node {
	return createAttribute(attributePointsAtX, values...)
}

func PointsAtY(values ...string) Node {
	return createAttribute(attributePointsAtY, values...)
}

func PointsAtZ(values ...string) Node {
	return createAttribute(attributePointsAtZ, values...)
}

func PreserveAlpha(values ...string) Node {
	return createAttribute(attributePreserveAlpha, values...)
}

func PreserveAspectRatio(values ...string) Node {
	return createAttribute(attributePreserveAspectRatio, values...)
}

func PrimitiveUnits(values ...string) Node {
	return createAttribute(attributePrimitiveUnits, values...)
}

func R(values ...string) Node {
	return createAttribute(attributeR, values...)
}

func Radius(values ...string) Node {
	return createAttribute(attributeRadius, values...)
}

func RefX(values ...string) Node {
	return createAttribute(attributeRefX, values...)
}

func RefY(values ...string) Node {
	return createAttribute(attributeRefY, values...)
}

func RepeatCount(values ...string) Node {
	return createAttribute(attributeRepeatCount, values...)
}

func RepeatDur(values ...string) Node {
	return createAttribute(attributeRepeatDur, values...)
}

func Restart(values ...string) Node {
	return createAttribute(attributeRestart, values...)
}

func Result(values ...string) Node {
	return createAttribute(attributeResult, values...)
}

func Rotate(values ...string) Node {
	return createAttribute(attributeRotate, values...)
}

func Rx(values ...float64) Node {
	return createAttribute(attributeRx, values...)
}

func Ry(values ...float64) Node {
	return createAttribute(attributeRy, values...)
}

func Scale(values ...string) Node {
	return createAttribute(attributeScale, values...)
}

func Seed(values ...string) Node {
	return createAttribute(attributeSeed, values...)
}

func ShapeRendering(values ...string) Node {
	return createAttribute(attributeShapeRendering, values...)
}

func Spacing(values ...string) Node {
	return createAttribute(attributeSpacing, values...)
}

func SpecularConstant(values ...string) Node {
	return createAttribute(attributeSpecularConstant, values...)
}

func SpecularExponent(values ...string) Node {
	return createAttribute(attributeSpecularExponent, values...)
}

func SpreadMethod(values ...string) Node {
	return createAttribute(attributeSpreadMethod, values...)
}

func StartOffset(values ...string) Node {
	return createAttribute(attributeStartOffset, values...)
}

func StdDeviation(values ...string) Node {
	return createAttribute(attributeStdDeviation, values...)
}

func StitchTiles(values ...string) Node {
	return createAttribute(attributeStitchTiles, values...)
}

func StopColor(values ...string) Node {
	return createAttribute(attributeStopColor, values...)
}

func StopOpacity(values ...string) Node {
	return createAttribute(attributeStopOpacity, values...)
}

func StrikethroughPosition(values ...string) Node {
	return createAttribute(attributeStrikethroughPosition, values...)
}

func StrikethroughThickness(values ...string) Node {
	return createAttribute(attributeStrikethroughThickness, values...)
}

func Stroke(values ...string) Node {
	return createAttribute(attributeStroke, values...)
}

func StrokeDasharray(values ...string) Node {
	return createAttribute(attributeStrokeDasharray, values...)
}

func StrokeDashoffset(values ...string) Node {
	return createAttribute(attributeStrokeDashoffset, values...)
}

func StrokeLinecap(values ...string) Node {
	return createAttribute(attributeStrokeLinecap, values...)
}

func StrokeLinejoin(values ...string) Node {
	return createAttribute(attributeStrokeLinejoin, values...)
}

func StrokeMiterlimit(values ...string) Node {
	return createAttribute(attributeStrokeMiterlimit, values...)
}

func StrokeOpacity(values ...string) Node {
	return createAttribute(attributeStrokeOpacity, values...)
}

func StrokeWidth(values ...string) Node {
	return createAttribute(attributeStrokeWidth, values...)
}

func SurfaceScale(values ...string) Node {
	return createAttribute(attributeSurfaceScale, values...)
}

func SystemLanguage(values ...string) Node {
	return createAttribute(attributeSystemLanguage, values...)
}

func TableValues(values ...string) Node {
	return createAttribute(attributeTableValues, values...)
}

func TargetX(values ...string) Node {
	return createAttribute(attributeTargetX, values...)
}

func TargetY(values ...string) Node {
	return createAttribute(attributeTargetY, values...)
}

func TextAnchor(values ...string) Node {
	return createAttribute(attributeTextAnchor, values...)
}

func TextDecoration(values ...string) Node {
	return createAttribute(attributeTextDecoration, values...)
}

func TextRendering(values ...string) Node {
	return createAttribute(attributeTextRendering, values...)
}

func TextLength(values ...string) Node {
	return createAttribute(attributeTextLength, values...)
}

func To(values ...string) Node {
	return createAttribute(attributeTo, values...)
}

func Transform(values ...string) Node {
	return createAttribute(attributeTransform, values...)
}

func TransformOrigin(values ...string) Node {
	return createAttribute(attributeTransformOrigin, values...)
}

func UnderlinePosition(values ...string) Node {
	return createAttribute(attributeUnderlinePosition, values...)
}

func UnderlineThickness(values ...string) Node {
	return createAttribute(attributeUnderlineThickness, values...)
}

func UnicodeBidi(values ...string) Node {
	return createAttribute(attributeUnicodeBidi, values...)
}

func Values(values ...string) Node {
	return createAttribute(attributeValues, values...)
}

func VectorEffect(values ...string) Node {
	return createAttribute(attributeVectorEffect, values...)
}

func ViewBox(values ...string) Node {
	return createAttribute(attributeViewBox, strings.Join(values, " "))
}

func Visibility(values ...string) Node {
	return createAttribute(attributeVisibility, values...)
}

func WordSpacing(values ...string) Node {
	return createAttribute(attributeWordSpacing, values...)
}

func WritingMode(values ...string) Node {
	return createAttribute(attributeWritingMode, values...)
}

func X(values ...string) Node {
	return createAttribute(attributeX, values...)
}

func X1(values ...string) Node {
	return createAttribute(attributeX1, values...)
}

func X2(values ...string) Node {
	return createAttribute(attributeX2, values...)
}

func XChannelSelector(values ...string) Node {
	return createAttribute(attributeXChannelSelector, values...)
}

func Xmlns(values ...string) Node {
	return createAttribute(attributeXmlns, values...)
}

func Y(values ...string) Node {
	return createAttribute(attributeY, values...)
}

func Y1(values ...string) Node {
	return createAttribute(attributeY1, values...)
}

func Y2(values ...string) Node {
	return createAttribute(attributeY2, values...)
}

func YChannelSelector(values ...string) Node {
	return createAttribute(attributeYChannelSelector, values...)
}

func Z(values ...string) Node {
	return createAttribute(attributeZ, values...)
}
