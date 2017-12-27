package main

import "fmt"

func main() {
	// scanWithSscan()
	// scanWithSscanf()
	// scanWithSscanln()
	//scanWithScan()
	// scanFromReader()
	// customFormatter()
	scanWithCustomScanner()
}

func scanWithSscan() {
	var d1, d2 int
	var s1 string
	fmt.Printf("Before scanning: %d, %d, %s\n", d1, d2, s1)
	if _, err := fmt.Sscan("5 7 \n9", &d1, &d2, &s1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After scanning with Sscan: %d, %d, %s\n", d1, d2, s1)
}

func scanWithSscanf() {
	var d1, d2 int
	var s1 string
	fmt.Printf("Before scanning: %d, %d, %s\n", d1, d2, s1)
	if _, err := fmt.Sscanf("5, 7 and 9 are my values", "%d, %d and %s", &d1, &d2, &s1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After scanning with Sscanf: %d, %d, %s\n", d1, d2, s1)
}

func scanWithSscanln() {
	var d1, d2 int
	var s1 string
	fmt.Printf("Before scanning: %d, %d, %s\n", d1, d2, s1)
	if _, err := fmt.Sscanln("5 7 9\n", &d1, &d2, &s1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After scanning with Sscanln: %d, %d, %s\n", d1, d2, s1)
}

func scanWithScan() {
	var s1, s2, s3 string
	fmt.Print("You inputs please: ")
	if _, err := fmt.Scan(&s1, &s2, &s3); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After scanning with Scan: %s, %s, %s\n", s1, s2, s3)
}

/*
func scanFromReader() {

}
func customFormatter() {

}
*/

func scanWithCustomScanner() {
	var c coords
	if _, err := fmt.Sscanf("(1, 2, 3)", "%ɸ", &c); err != nil {
		fmt.Printf("Error scanning: %s\n", err)
	}
	//fmt.Printf("%v\n", c)
	fmt.Printf("%ɸ\n", c)
}

type coords struct {
	x, y, z int
}

func (c *coords) Scan(f fmt.ScanState, v rune) error {
	switch v {
	case 'ɸ':
		n, err := fmt.Fscanf(f, "(%v, %v, %v)", &c.x, &c.y, &c.z)
		if err != nil {
			return err
		}
		if n != 3 {
			return fmt.Errorf("3 values expected, got %d", n)
		}
	}
	return nil
}

func (c coords) Format(f fmt.State, v rune) {
	switch v {
	case 'ɸ':
		fmt.Fprintf(f, "ɸ(%v, %v, %v)", c.x, c.y, c.z)
	case 'v':
		fmt.Fprintf(f, "Coords{x=%v, y=%v, z=%v}", c.x, c.y, c.z)
	default:
		fmt.Fprintf(f, "%v, %v, %v", c.x, c.y, c.z)
	}
}
