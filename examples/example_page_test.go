package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/daarxwalker/gox"
)

func TestPageExample(t *testing.T) {
	t.Run(
		"render example page", func(t *testing.T) {
			pageContent := Fragment(
				H1(Text("Example page")),
				P(Text("Example paragraph")),
			)
			page := createPage("en", "Example page", pageContent)
			assert.Equal(
				t,
				Minify(
					`
				<!DOCTYPE html>
				<html lang="en">
					<head>
						<title>Example page</title>
						<meta name="viewport" content="width=device-width,initial-scale=1.0" />
					</head>
					<body>
						<h1>Example page</h1>
						<p>Example paragraph</p>
					</body>
				</html>
				`,
				),
				Render(page),
			)
		},
	)
}
