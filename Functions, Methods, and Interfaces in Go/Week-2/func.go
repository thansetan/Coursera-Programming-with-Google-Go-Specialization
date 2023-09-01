package main

import "fmt"

/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity v_o,
and initial displacement s_o.

s = ½ a t2 + v_o*t + s_o

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity v_o, and initial
displacement s_o. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time,
and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, v_o = 2, s_o = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))
*/

func main() {
	var (
		s0, v0, a, t float64
	)
	what := []string{"acceleration (a)", "initial velocity (v0)", "initial displacement (s0)"}
	values := []*float64{&a, &v0, &s0}
	for i, w := range what {
		fmt.Printf("Enter value for %s :", w)
		fmt.Scan(values[i])
	}
	fmt.Printf("Enter value for time (t): ")
	fmt.Scan(&t)

	displacement := GenDisplaceFn(a, v0, s0)
	fmt.Printf("Displacement after %.1f seconds : %.2f", t, displacement(t))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 1.0/2*a*t*t + v0*t + s0
	}
}
