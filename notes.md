### Keywords
* `package`
> Declares the package the code belongs to. Each package corresponds to a single idea
> `package <name of package>`
* `import` 
> Specifies the packages this code will use
> Packages specify functions
```{go}
import(
    "<package name>"
)
````
* `func`
> Declares a function 
> Body enclosed in curly braces (See Curly Braces)

### Curly Braces
Go is picky about the placement of curly braces 
Below is an example of the correct placement of curly braces
```{go}
func main(){
}
```
### Comments
Single line comments 
> Use `//`
Multi line Comments
> Use `/* */`

### Printing
`fmt.Print`
`fmt.Printf`
> Allows you to insert values anywhere
> First argument always text
> Allows you to specify width
`fmt.Println`
> Automatically inserts a newline at the end

### Constants
`const`

### Variables
`var`
> Can declare variables individually 
```{go}
var foo = 42
var bar = 1337
```
> Can declare them as a group
```{go}
var(
    foo = 42
    bar = 1337
```
> Can also declare multiple variable on a single line
```{go}
var foo, bar = 42, 1337
```

:
