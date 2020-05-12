package main
import (
	"fmt"
	"strings"
)

const ElementColor = "\033[1;33m%s\033[0m"  // Color code to indicate the element in the cast name

func check(name string) {
	// Short codes of elements in periodic table
	pt := [...]string{"H", "Li", "Na", "K", "Rb", "Cs", "Fr", "Be", "Mg", "Ca", "Sr", "Ba", "Ra", "Sc", "Y", 
	"La", "Ac", "Ti", "Zr", "Hf", "Rf", "V", "Nb", "Ta", "Db", "Cr", "Mo", "W", "Sg", "Mn", "Tc",
	"Re", "Bh", "Fe", "Ru", "Os", "Hs", "Co", "Rh", "Ir", "Mt", "Ni", "Pd", "Pt", "Ds", "Cu",
	"Ag", "Au", "Rg", "Zn", "Cd", "Hg", "Cn", "B", "Al", "Ga", "In", "Tl", "Nh", "C", "Si", "Ge",
	"Sn", "Pb", "Fl", "N", "P", "As", "Sb", "Bi", "Mc", "Lv", "F", "Cl", "Br", "I", "At", "Ts",
	"He", "Ne", "Ar", "Kr", "Xe", "Rn", "Og", "Ce", "Pr", "Nd", "Pm", "Sm", "Gd", "Tb", "Dy",
	"Ho", "Er", "Tm", "Yb", "Lu", "Th", "Pa", "U", "Np", "Pu", "Am", "Cm", "Bk", "Cf", "Es", 
	"Fm", "Md", "No", "Lr"}

	c := 0
	for _,e := range pt {
		if strings.Contains(strings.ToLower(name),strings.ToLower(e)) {
			c += 1
			i := strings.Index(strings.ToLower(name),strings.ToLower(e))
			fmt.Printf(name[:i])
			fmt.Printf(ElementColor, e)
			fmt.Println(name[i+len(e):])
			break
			}
		}

	if c == 0 {
		fmt.Println(name)
		}
	}

func main(){
	
	// "Aqaxt Tems" was created as a test name that contains any element in it.
	cast := [...]string{"Bryan Cranston", "Anna Gunn", "Aaron Paul", "Betsy Brandt", "RJ Mitte", "Dean Norris", 
	"Bob Odenkirk", "Jonathan Banks", "Giancarlo Esposito", "Charles Baker", "Jesse Plemons", 
	"Christopher Cousins", "Laura Fraser", "Matt Jones", "Michael Shamus Wiles", "Lavell Crawford", 
	"Ray Campbell", "Krysten Ritter", "Carmen Serano","Aqaxt Tems", "Emily Rios"}
	
	fmt.Println("- Breaking Bad Casting with Element Symbols -\n")

	for _,name := range cast {
		check(name)
	}	
}