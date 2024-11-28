# this is just a documetation to the



```go
type Film struct{
	Title string;
	Director string ;
}



func main(){
	h1 := func(w http.ResponseWriter , r *http.Request){
		templ := template.Must(template.ParseFiles("index.html"));
		films := map[string][]Film{
			"films":{
				{Title: "The GodFather" , Director: "Franc ford"},
				{Title: "Blade Runner" , Director: "Ridley scot"},
				{Title: "The Things" , Director: "Jhon Carpenter"},
			},
		};
		templ.Execute(w , films);
	}
	http.HandleFunc("/" , h1)
	log.Fatal(http.ListenAndServe(":8000" , nil));
}
```

### so we run a server and take templ to take the `index.html` as a template
### and make map that is key  = string and the value = array of Film struct
### so in the template that is `index.html` we can use the {{}} notation because
### go will convert it to template and render it as html to the server on the end
### so we can access the first field of the map `films` by use this syntax 
### `{{.films}}`  and if we want to iterate throw the all array of the `films` array
### we can use {{range .films}} html here to render the films title then {{end}}

## like this 

```html 
<!DOCTYPE html>
<html lang="en">
<head>
       <meta charset="UTF-8">
       <meta name="viewport" content="width=device-width, initial-scale=1.0">
       <title>GoHtmx</title>
</head>
<body>
       {{range .films}}
       <p>{{.Title}} -- {{.Director}}</p>
       {{end}}
</body>
</html>
```

### for more about how you can render the html as template and use go syntax on the 
### html visit this articel https://golangforall.com/en/post/templates.html
