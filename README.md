# ASCII-ART-WEB 

## Description 
`ascii-art-web` is a web application consists in creating and running a server, in which it will be possible to use a web GUI version of the terminal base `ascii-art`
- The code ignores the unprintable characters if they are in the text.
- Max legth of the text is 2000 characters.

## Authors
    - aouchcha
    - maynaou
    - yakhaldy
## Usage
    to start the server you need to run the bash command 
```bash
        'go run main.go'
```

    and then fire the browser on http://localhost on port 8080
```
        'http://localhost:8080'
```
## Implementation details
    - the projects code is devided into the 'main', 'fs', 'functions', 'statics' and 'syles'  folders. 
    - the package 'functions' contains the 'handlers' for the server,
    - the 'fs' package  contains the 'ascii-art' code.
    - banners are located withing the static directory. 
    - templates are located withing the templates directory.
    - CSS files for styling the pages are located withing the styles directory.
    - the code ignore the imprintable characters as the example in the intra do.
 
```
    Method               URLPattern                        Handler                                    Action
    GET                    /                                Welcom                              Show The home page front end
    POST                  /ascii-art                         Last                         send the request body and disply the output 
```