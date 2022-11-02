package main

import (
	"fmt"
)

/*
1. Enable operations of entering, changing, viewing and deleting: members, workers, books, literary genre, bookshelves, categories of membership fees, and borrowed books.
2. Enter new items and change existing ones from the keyboard.
3. The desired action (entry, modification, review, deletion) is entered by the user via the keyboard based on the options offered in the terminal.
4. For a new entry, define the required columns and limit the entry if the required data is not entered.
5. To connect the tables, keys will be used, which represent the unique identifier of a certain item, such as:
- worker/member: text record JMBG of 13 digits,
- book: text record of 5 characters, where the 1st character indicates the genre (B1010, K5425, etc.),
- shelf: text record with genre label and shelf numbering;
6. Create project documentation in the form of a Word document. Use this document as a template, and start with Chapter 5 of your project report.
*/

type man struct {
	name     string
	lastname string
	year     string
	JMBG     string
}

type member struct {
	man
	limitUpRaise    uint8
	currentlyRaised uint8
}
type worker struct {
	man
	position string
	shift    string
}
type book struct {
	title  string
	author string
	ID     string
	Status string
}
type genre struct {
	NameZanra string
	ZnakZanra string
}

type shelf struct {
	genre string
	index int
}

func main() {
	Members := []member{
		{
			man: man{name: "Amel", lastname: "Smajic", year: "2002", JMBG: "2102200286041"},

			limitUpRaise:    3,
			currentlyRaised: 0,
		},
		{
			man: man{
				name:     "Armin",
				lastname: "Avdic",
				year:     "1999",
				JMBG:     "120419991234",
			},

			limitUpRaise:    3,
			currentlyRaised: 0,
		},
	}

	Workers := []worker{

		{
			man: man{
				name:     "Team",
				lastname: "tibersom",
				year:     "1982",
				JMBG:     "19820312310293",
			},
			position: "Manager",
			shift:    "First",
		}, {
			man: man{
				name:     "Andrey",
				lastname: "Robines",
				year:     "1999",
				JMBG:     "2102199991234",
			},
			position: "Librarian",
			shift:    "First",
		}, {
			man: man{
				name:     "Michael",
				lastname: "Jeas",
				year:     "1983",
				JMBG:     "120319839213",
			},
			position: "Caretaker",
			shift:    "First",
		},
	}

	Books := []book{
		{
			title:  "The Lion, the Witch and the Wardrobe",
			author: "C.S. Lewis",
			ID:     "",
			Status: "available",
		}, {
			title:  "She: A History of Adventure",
			author: "H. Rider Haggard",
			ID:     "",
			Status: "available",
		}, {
			title:  "The Da Vinci Code",
			author: "Dan Brown",
			ID:     "",
			Status: "available",
		},
	}
	Genres := []genre{
		{
			NameZanra: "Fantasy",
			ZnakZanra: "F",
		},
		{
			NameZanra: "Adventure",
			ZnakZanra: "A",
		},
		{
			NameZanra: "History",
			ZnakZanra: "H",
		},
		{
			NameZanra: "Thriller",
			ZnakZanra: "T",
		},
	}

	shelves := []shelf{
		{
			genre: "Fantasy",
			index: 0000,
		},
		{
			genre: "Adventure",
			index: 0001,
		},
		{
			genre: "History",
			index: 0002,
		},
		{
			genre: "Thriller",
			index: 0003,
		},
	}

	var active = true
	for active {
		var options = 0
		//fmt.Printf("")
		fmt.Printf("Welcome to the library\n")
		fmt.Printf("Please select the database you need to edit\n")
		fmt.Printf("1.) User\n")
		fmt.Printf("2.) Employees\n")
		fmt.Printf("3.) Book\n")
		fmt.Printf("4.) Genre\n")
		fmt.Printf("5.) Shelf\n")
		fmt.Printf("6.) Exit\n")

		fmt.Scan(&options)

		switch options {

		case 1:
			userf(Members)

		case 2:
			workerf(Workers)

		case 3:
			bookf(Books)

		case 4:
			genref(Genres)

		case 5:
			shelff(shelves)

		case 6:
			active = false

		default:
			error()
		}
	}

}

func error() {

	fmt.Printf("Please enter correct data \n")
}

func Manu2(base string) (option2 int) {

	fmt.Printf("Welcome to database %v \n", base)
	fmt.Printf("Please select the desired option\n")
	fmt.Printf("1.) Enter %v\n", base)
	fmt.Printf("2.) Modify %v\n", base)
	fmt.Printf("3.) Browse %v\n", base)
	fmt.Printf("4.) Delete %v\n", base)
	fmt.Printf("5.) Exit")
	fmt.Scan(&option2)

	return
}

func userf(User []member) {

	var active bool = true
	for active {

		var option = Manu2("User")
		fmt.Printf("selection is: %v", option)

		switch option {
		case 1:
			UserInput(User)
		case 2:

		case 3:
			BrowseUsers(User)
		case 4:

		case 5:
			active = false

		default:
			error()
		}
	}
}
func line() {
	fmt.Printf("\n------------------------------------------------ ------------------------------------\n")

}

func BrowseUsers(Members []member) {
	for i, val := range Members {
		line()
		fmt.Printf("\nUser name is: %v \n", val.name)
		fmt.Printf("The user's last name is: %v \n", val.lastname)
		fmt.Printf("The user's age is: %v \n", val.year)
		fmt.Printf("User jmbg is: %v \n", val.JMBG)
		if val.currentlyRaised == 0 {
			fmt.Printf("The user has no uploaded books and can upload %v books\n\n", val.limitUpRaise)

		} else {
			fmt.Printf("User has %v of books uploaded and can upload %v more\n\n ", val.currentlyRaised, (val.limitUpRaise - val.currentlyRaised))

		}
		line()
		i++
	}

}

func UserInput(Members []member) {
	var n = len(Members)
	fmt.Printf("\nPlease enter the information of the new User : \n")
	line()
	var name string
	fmt.Printf("User name : \n")
	fmt.Scanf(name)
	Members[n].name = name
	fmt.Printf("User last name : \n")
	fmt.Scanf(Members[n].lastname)
	fmt.Printf("User age :\n")
	fmt.Scanf(Members[n].year)
	fmt.Printf("User JMBG : \n")
	fmt.Scanf(Members[n].JMBG)
	fmt.Printf("Enter lifting limit (int) : \n")
	fmt.Scan(Members[n].limitUpRaise)

}
func workerf(Worker []worker) {

	var active = true
	for active {
		var option = Manu2("Employee")

		switch option {

		case 1:

		case 2:

		case 3:
			viewWorker(Worker)
		case 4:

		case 5:
			active = false

		default:
			error()
		}
	}
}

func viewWorker(Workers []worker) {
	fmt.Printf("\nList of workers:\n")
	line()

	for i, val := range Workers {
		line()
		fmt.Printf("\nName : %v\n", val.name)
		fmt.Printf("Surname : %v\n", val.lastname)
		fmt.Printf("Year : %v\n", val.year)
		fmt.Printf("Jmbg : %v\n", val.JMBG)
		fmt.Printf("Position : %v\n", val.position)
		fmt.Printf("Shift : %v\n", val.shift)
		line()
		i++
	}

}

func bookf(Books []book) {

	var active = true
	for active {

		var option = Manu2("Book")
		switch option {
		case 1:

		case 2:

		case 3:
			ViewBook(Books)
		case 4:

		case 5:
			active = false

		default:
			error()
		}
	}
}

func ViewBook(Books []book) {
	fmt.Printf("\nList of books : \n")
	line()
	for i, val := range Books {

		fmt.Printf("\nThe name of the book is : %v\n", val.title)
		fmt.Printf("The author of the book is : %v\n", val.author)
		fmt.Printf("Id of the book is : %v\n", val.ID)
		fmt.Printf("Book status is : %v\n", val.Status)
		line()
		i++
	}

}

func genref(Genres []genre) {

	var active bool = true
	for active {
		var option = Manu2("Genre")

		switch option {
		case 1:

		case 2:

		case 3:
			GenreView(Genres)
		case 4:

		case 5:
			active = false

		default:
			error()
		}
	}
}

func GenreView(Genres []genre) {
	fmt.Printf("\nView of genres\n")
	line()

	for i, val := range Genres {
		fmt.Printf("Genre NAME is : %v \n", val.NameZanra)
		fmt.Printf("Genre code is : %v \n\n", val.ZnakZanra)
		line()

		i++
	}

}
func shelff(Shelf []shelf) {

	var active = true
	for active {
		var option = Manu2("Shelf")

		switch option {

		case 1:

		case 2:

		case 3:
			ViewShelf(Shelf)
		case 4:

		case 5:
			active = false

		default:
			error()
		}
	}
}
func ViewShelf(Shelf []shelf) {

	fmt.Printf("\n\nList of shelves\n")
	line()
	for i, val := range Shelf {
		fmt.Printf("The genre of the shelf is: %v\n", val.genre)
		fmt.Printf("Shelf index is : %v\n", val.index)
		line()
		i++
	}
}
