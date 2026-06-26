package main

func processor(text string) string {
	// conversions (hex/bin)
	text = Conversion(text)

	text = casing(text)

	text = Capitalized(text)

	// articles (a/an)
	text = Articles(text)

	// punctuation spacing
	text = punctuation(text)

	// quotes
	text = Quote(text)

	return text
}
