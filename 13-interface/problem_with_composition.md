consider the below java code

```java

class Animal {
    boolean hungry;

    Animal(boolean hungry) {
        this.hungry = hungry;
    }

    String makeDefaultSound() {
        return "";
    }

    String makeSound() {
        if (hungry) {
            return "I am hungry!!";
        }
        return makeDefaultSound();
    }
}

class Cat extends Animal {
    Cat(boolean hungry) {
        super(hungry);
    }

    @Override
    String makeDefaultSound() {
        return "mew";
    }
}

class Dog extends Animal {
    Dog(boolean hungry) {
        super(hungry);
    }

    @Override
    String makeDefaultSound() {
        return "Woof";
    }
}

public class Main {
    public static void main(String[] args) {
        Cat cat = new Cat(false);
        Dog dog = new Dog(false);

        System.out.printf("cat says %s%n", cat.makeSound());
        System.out.printf("dog says %s%n", dog.makeSound());
    }
}



```

when you run this code you will get the output

```
cat says mew
dog says Woof
```

if you write the equivalent code in golang  it will look like this

```go

package main

import "fmt"

type Animal struct {
	Hungry bool
}

func (a Animal) MakeDefaultSound() string {
	return ""
}

func (a Animal) MakeSound() string {
	if a.Hungry {
		return "i am hungry!!"
	}
	return a.MakeDefaultSound()
}

type Cat struct {
	Animal
}

func (c Cat) MakeDefaultSound() string {
	return "mew"
}

type Dog struct {
	Animal
}

func (d Dog) MakeDefaultSound() string {
	return "Woof"
}

func main() {
	cat := Cat{}
	cat.Hungry = false
	dog := Dog{}
	dog.Hungry = false

	//unfortunaterly comnposition doesnot work like inheritance and both anmials default sound of animal.
	//they should have made defulat sound of cat and dog.
	fmt.Printf("cat says %s \n", cat.MakeSound())
	fmt.Printf("dog says %s \n", dog.MakeSound())
}

```
unfortunaterly comnposition doesnot work like inheritance and we get output

```
cat says  
dog says 

```

The reason for this is the struct cat

```go 
type Cat struct {
	Animal
}
```
is not inhertiting from animal what it does is it also created a member variable Animal and also promotes all of Animals bound methods so that they are availabe
to call as if they are memebers are of cat struct. But the method signeture doesnot change though. so the when we call MakeSound from cat we are calling mke sound of Animal.

How do we fix it?

```go 

package main

import "fmt"

type Animal struct {
	Hungry bool
}

func (a Animal) MakeDefaultSound() string {
	return ""
}

type SoundMaker interface {
	MakeDefaultSound() string
}

func (a Animal) MakeSound(sm SoundMaker) string {
	if a.Hungry {
		return "i am hungry!!"
	}
	return sm.MakeDefaultSound()
}

type Cat struct {
	Animal
}

func (c Cat) MakeDefaultSound() string {
	return "mew"
}

type Dog struct {
	Animal
}

func (d Dog) MakeDefaultSound() string {
	return "Woof"
}

func main() {
	cat := Cat{}
	cat.Hungry = false
	dog := Dog{}
	dog.Hungry = false

	fmt.Printf("cat says %s \n", cat.MakeSound(cat))
	fmt.Printf("dog says %s \n", dog.MakeSound(dog))
}


```

since passing cat or dog as parameter where ever it is called is bit cluncy we can rewrite the code better.


```go 
package main

import "fmt"

type Animal struct {
	Hungry bool
}

func (a Animal) MakeDefaultSound() string {
	return ""
}

type SoundMaker interface {
	MakeDefaultSound() string
}

func (a Animal) MakeSound(sm SoundMaker) string {
	if a.Hungry {
		return "i am hungry!!"
	}
	return sm.MakeDefaultSound()
}

type Cat struct {
	Animal
}

func (c Cat) MakeDefaultSound() string {
	return "mew"
}

func (c Cat) MakeSound() string {
	return c.Animal.MakeSound(c)
}

type Dog struct {
	Animal
}

func (d Dog) MakeDefaultSound() string {
	return "Woof"
}

func (d Dog) MakeSound() string {
	return d.Animal.MakeSound(d)
}

func main() {
	cat := Cat{}
	cat.Hungry = false
	dog := Dog{}
	dog.Hungry = false

	fmt.Printf("cat says %s \n", cat.MakeSound())
	fmt.Printf("dog says %s \n", dog.MakeSound())
}

```

in case of an animla that does not make sound

```go 

type Spider struct {
    Animal
}


func(s Spider) MakeSound() string {
    return s.Animal.makeSound(s.Animal)
}

```


in go we need to write more code compared to object oriented languages


but there is an advantage in go  

let us go back to java and make the inheritance tree a little more deeper

```java
class Canine extends Animal {
    Canine(boolean hungry) {
        super(hungry);
    }

    @Override
    String makeDefaultSound() {
        return "Woof";
    }
}
class Dog extends Canine {
    Dog(boolean hungry) {
        super(hungry);
    }

    @Override
    String makeDefaultSound() {
        return "Woof";
    }
}

```

No it becomes tricky for the developer to understand which class says woof , the canine or dog

deeper more complex inheritance structures make code duplication difficult to detect and cleanup

in go we can see what parameter is being passed in

```go
func (d Dog) MakeSound() string {
	return d.Animal.MakeSound(d)
}
```

if we want the default sound from canine

```go

func (d Dog) MakeSound() string {
	return d.Animal.MakeSound(d.Canine)
}

```

from https://www.youtube.com/watch?v=4KsgNO2exS8