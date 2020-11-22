https://img.shields.io/badge/Language-Go-blue

# chef-interface-example
An example to illustrate how to work with interfaces in Go.

The Cutter interface defines a cut function by implementing this function knife and chainsaw become implementers of Cutter.

``` 
type Cutter interface {
    	Cut()
    }
```

Similarly the Cooker interface defines the cook function and implementing this function pan and oven become implementers of Cooker.

``` 
type Cooker interface {
        Cook()
    }   
```

The chef has the ability to cut and cook irrespective of the implementers of Cutter and Cooker interface.

```
type chef struct {
        Cutter
        Cooker
    }
``` 

Shannon Bennett, one of my favorite chefs loves to play with his knife and pan to create his dishes.
Another chef  maybe fatDude loves to chainsaw and throw things in a oven to roast them up... So yeah you get the concept.

```
func main() {
	shannon := chef{
		Cutter: knife{},
		Cooker: pan{},
	}

	shannon.Do()
}
```

We can create a chef in the main and pass the implementers to the respective interfaces that define the chef's behaviors.


