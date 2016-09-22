package html

const (
	Index = `<!DOCTYPE html>
<html>
    <head>

    </head>
    <body>
        <h1>gophurl</h1>

	<table style="width: 100%"> 
	   {{ range $key, $url := . }}
	   <tr>
	       <td>{{ $key }}</td>
	       <td><a href="{{ $url.Url }}">{{ $url.Url }}</a></td>
	       <td>{{ $url.Accessed }}</td>
	       <td>{{ $url.Created }}</td>
	   </tr>
	   {{ end }}
	</table>
    </body>
</html>`
)
