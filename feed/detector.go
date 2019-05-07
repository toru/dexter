package feed

import(
	"bytes"
)

// Heuristically determines if the document is an Atom feed by searching
// for the format namespace. Search is given up after 1024 bytes.
func IsAtomFeed(doc []byte) bool {
	upper := len(doc)
	if upper > detectionLimit {
		upper = detectionLimit
	}

	// TODO: Sadly this is not enough, as shown by the NASA feeds.
	// Partially parse the given document and evaluate from there.
	return bytes.Contains(doc[:upper], []byte(atomHint))
}
