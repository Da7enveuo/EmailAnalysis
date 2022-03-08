package Formatting

type MasterHtmlBody struct {
	HtmlText []string
	Links    []string
	Domains  []string
	Js       Javascript
}

type Javascript struct {
	imports  []string
	inlineJS []string
}
