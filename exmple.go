package tsmap

import (
	"fmt"
)

type Person struct {
	name string
	age  int
	sex  string
}

type Some struct {
	something string
	someone   Person
}

func genSome() []Some {
	ss := make([]Some, 0, 10)
	for i := 0; i < 10; i++ {
		ss = append(ss, Some{
			something: "something" + string(rune(i)),
			someone: Person{
				name: "someone" + string(rune(i)),
				age:  i,
				sex:  "male",
			},
		})
	}

	return ss
}

func Example() {
	mp := New[Some, any]()
	ss := genSome()

	for _, some := range ss {
		mp.Set(some, some)
	}

	fmt.Printf("there are %v elements in TSMap \n", mp.Len())

	fn := func(key Some, value any) {
		fmt.Printf("key: %v, value: %v \n", key, value)
	}
	except := func(key Some, value any) bool {
		return key.someone.age > 5
	}

	// print all elements except person.age > 5
	mp.Range(fn, except)
}

func ExampleRange() {
	mp := New[Some, any]()
	ss := genSome()

	for _, some := range ss {
		mp.Set(some, some)
	}

	fmt.Printf("there are %v elements in TSMap \n", mp.Len())

	fn := func(key Some, value any) {
		fmt.Printf("key: %v, value: %v \n", key, value)
	}

	except := func(key Some, value any) bool {
		return key.someone.age > 8
	}
	except = WithUntil(except)

	// print element until first person whose age > 8
	mp.Range(fn, except)
}
