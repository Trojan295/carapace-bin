package gh

import "github.com/rsteube/carapace"

// ActionReleaseFields completes release fields
//
//	name
//	tagName
func ActionReleaseFields() carapace.Action {
	return carapace.ActionValues(
		"name",
		"tagName",
		"isDraft",
		"isLatest",
		"isPrerelease",
		"createdAt",
		"publishedAt",
	)
}
