package main

import (
	"bytes"
	"html/template"
)

const htmlTemplate = `
<!DOCTYPE html>

<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<meta name="author" content="ying32">

	<title>Markdown 预览</title>
	<script src="https://cdn.jsdelivr.net/npm/jquery@1.11.1/dist/jquery.min.js"></script>

	<!-- markdown -->
	<script src="https://cdn.jsdelivr.net/npm/marked@0.6.2/marked.min.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/moment@2.18.1/min/moment.min.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/moment@2.18.1/locale/zh-cn.min.js"></script>
	<!-- 代码高亮 -->
	<link href="https://cdn.jsdelivr.net/npm/highlight.js@9.15.10/styles/a11y-light.min.css" rel="stylesheet">
	<script src="https://cdn.jsdelivr.net/npm/highlight.js@9.15.1/highlight.min.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/highlight.js@9.15.1/languages/go.min.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/highlight.js@9.15.1/languages/delphi.min.js"></script>
    <!-- markdown css -->
	<link href="https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/3.0.1/github-markdown.min.css" rel="stylesheet">

 
	<style>
		.markdown-body {
			box-sizing: border-box;
			min-width: 200px;
		<!-- max-width: 980px; -->
			margin: 0 auto;
			padding: 45px;
		}
		@media (max-width: 767px) {
			.markdown-body {
				padding: 15px;
			}
		}
	</style>
	<script>
		// 语言设置
		moment.locale('zh-cn');

		hljs.initHighlightingOnLoad();

		marked.setOptions({
			renderer: new marked.Renderer(),
			gfm: true,
			tables: true,
			breaks: false,
			pedantic: false,
			sanitize: false,
			smartLists: true,
			smartypants: false,
			highlight: function (code, lang) {
				return hljs.highlightAuto(code).value;
			}
		});
	</script>
</head>

<body>
	<article class="markdown-body" id="mdcontext">{{.MdData}}</article>
<script>
	$(function() {
	   $("#mdcontext").html(marked($("#mdcontext").text(), {gfm: true, breaks: true, tables:true}));
	}); 
</script>
</body>
</html>
`

func makeHtmlPage(datas map[string]interface{}) *bytes.Buffer {
	if len(datas) == 0 {
		return nil
	}
	buf := bytes.NewBuffer([]byte{})
	t := template.New("index")
	t, err := t.Parse(htmlTemplate)
	if err != nil {
		return nil
	}
	err = t.Execute(buf, datas)
	if err != nil {
		return nil
	}
	return buf
}
