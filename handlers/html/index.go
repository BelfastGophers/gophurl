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
	       <td><a href="{{ $url.URL }}">{{ $url.URL }}</a></td>
	       <td><a href="{{ $url.Code }}">{{ $url.Code }}</a></td>
	       <td>{{ $url.AccessCount }}</td>
	       <td>{{ $url.Created }}</td>
	   </tr>
	   {{ end }}
	</table>
    </body>
</html>`
)
