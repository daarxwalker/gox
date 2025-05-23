package hx

const (
	// ClassAdded Applied to a new piece of content before it is swapped, removed after it is settled.
	ClassAdded = "htmx-added"

	// ClassIndicator A dynamically generated class that will toggle visible (opacity:1) when a htmx-request class is present
	ClassIndicator = "htmx-indicator"

	// ClassRequest Applied to either the element or the element specified with hx-indicator while a request is ongoing
	ClassRequest = "htmx-request"

	// ClassSettling Applied to a target after content is swapped, removed after it is settled. The duration can be modified via hx-swap.
	ClassSettling = "htmx-settling"

	// ClassSwapping Applied to a target before any content is swapped, removed after it is swapped. The duration can be modified via hx-swap.
	ClassSwapping = "htmx-swapping"
)
