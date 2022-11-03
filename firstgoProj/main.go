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
			userf(&Members)

		case 2:
			workerf(&Workers)

		case 3:
			bookf(&Books)

		case 4:
			genref(&Genres)

		case 5:
			shelff(&shelves)

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

func userf(User *[]member) {

	var active bool = true
	for active {

		var option = Manu2("User")
		fmt.Printf("selection is: %v", option)

		switch option {
		case 1:
			UserInput(User)
		case 2:
			var name string
			fmt.Print("JMBG OF the user you wish to modify")
			fmt.Scan(&name)
			UserModify(User, name)
		case 3:
			BrowseUsers(User)
		case 4:
			UserDellet(User)
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

func UserDellet(Members *[]member) {
	var i int
	fmt.Print("\n Index OF the targeted User : ")
	fmt.Scan(&i)
	temp := make([]member, 0)
	temp = append(temp, (*Members)[:i]...)
	temp = append(temp, (*Members)[i+1:]...)
	(*Members) = temp

}

func UserModify(Members *[]member, jmbg string) {

	for i, val := range *Members {
		if val.JMBG == jmbg {

			line()
			fmt.Printf("User name : \n")
			fmt.Scan(&val.name)
			fmt.Printf("User last name : \n")
			fmt.Scan(&val.lastname)
			fmt.Printf("User age :\n")
			fmt.Scan(&val.year)
			fmt.Printf("User JMBG : \n")
			fmt.Scan(&val.JMBG)
			fmt.Printf("Enter lifting limit (int) : \n")
			fmt.Scan(&val.limitUpRaise)
			line()
			(*Members)[i] = val
		}

		i++
	}

}

func BrowseUsers(Members *[]member) {
	for i, val := range *Members {
		line()
		fmt.Printf("\nUser index is: %v \n", i)
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

func UserInput(Members *[]member) {

	fmt.Printf("nPlease enter the information of the new User : n")
	line()

	var name string
	var lastname string
	var age string
	var JMBG string
	var limit uint8

	fmt.Printf("User name : \n")
	fmt.Scan(&name)
	fmt.Printf("User last name : \n")
	fmt.Scan(&lastname)
	fmt.Printf("User age :\n")
	fmt.Scan(&age)
	fmt.Printf("User JMBG : \n")
	fmt.Scan(&JMBG)
	fmt.Printf("Enter lifting limit (int) : \n")
	fmt.Scan(&limit)
	temp := member{
		man: man{

			name:     name,
			lastname: lastname,
			year:     age,
			JMBG:     JMBG,
		},
		limitUpRaise:    limit,
		currentlyRaised: 0,
	}

	*Members = append(*Members, temp)

}
func workerf(Worker *[]worker) {

	var active = true
	for active {
		var option = Manu2("Employee")

		switch option {

		case 1:
			WorkerInput(Worker)
		case 2:
			var name string
			fmt.Printf("\nGet JMBG OF The worker zou want to modify :")
			fmt.Scan(&name)
			WorkerModify(Worker, name)
		case 3:
			viewWorker(Worker)
		case 4:
			WorkerDellet(Worker)
		case 5:
			active = false

		default:
			error()
		}
	}
}
func WorkerDellet(plc *[]worker) {
	var i int
	fmt.Print("\n Index OF the targeted worker : ")
	fmt.Scan(&i)
	temp := make([]worker, 0)
	temp = append(temp, (*plc)[:i]...)
	temp = append(temp, (*plc)[i+1:]...)
	(*plc) = temp

}

func WorkerModify(Workesr *[]worker, jmbg string) {

	for i, val := range *Workesr {
		if val.JMBG == jmbg {

			line()
			fmt.Printf("User name : \n")
			fmt.Scan(&val.name)
			fmt.Printf("User last name : \n")
			fmt.Scan(&val.lastname)
			fmt.Printf("User age :\n")
			fmt.Scan(&val.year)
			fmt.Printf("User JMBG : \n")
			fmt.Scan(&val.JMBG)
			fmt.Printf("Enter workesr shift : \n")
			fmt.Scan(&val.shift)

			fmt.Printf("Enter workers possition : \n")
			fmt.Scan(&val.position)
			line()
			(*Workesr)[i] = val
		}

		i++
	}

}

func WorkerInput(Worker *[]worker) {

	line()
	var temp worker
	fmt.Printf("User name :\n")
	fmt.Scan(&temp.name)

	fmt.Printf("User lastname :\n")
	fmt.Scan(&temp.lastname)

	fmt.Printf("User age :\n")
	fmt.Scan(&temp.year)

	fmt.Printf("User JMBG :\n")
	fmt.Scan(&temp.JMBG)

	fmt.Printf("User position :\n")
	fmt.Scan(&temp.position)

	fmt.Printf("User shift :\n")
	fmt.Scan(&temp.shift)

	*Worker = append(*Worker, temp)
	line()
}

func viewWorker(Workers *[]worker) {
	fmt.Printf("\nList of workers:\n")
	line()

	for i, val := range *Workers {
		line()
		fmt.Printf("\nWorkers index : %v\n", i)
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

func bookf(Books *[]book) {

	var active = true
	for active {

		var option = Manu2("Book")
		switch option {
		case 1:
			BookInput(Books)
		case 2:
			var ID string
			fmt.Print("\nID of the book: ")
			fmt.Scan(ID)

			BookModify(Books, ID)
		case 3:
			ViewBook(Books)
		case 4:
			BookDellet(Books)
		case 5:
			active = false

		default:
			error()
		}
	}
}
func BookDellet(plc *[]book) {
	var i int
	fmt.Print("\n Index OF the targeted book : ")
	fmt.Scan(&i)
	temp := make([]book, 0)
	temp = append(temp, (*plc)[:i]...)
	temp = append(temp, (*plc)[i+1:]...)
	(*plc) = temp

}
func BookModify(Books *[]book, Ident string) {
	for i, val := range *Books {

		if val.ID == Ident {
			line()
			fmt.Print("\nName of Books autor : ")
			fmt.Scan(val.author)
			fmt.Print("\nName of the book : ")
			fmt.Scan(val.title)
			fmt.Print("\nID of the book : ")
			fmt.Scan(val.ID)
			line()
			(*Books)[i] = val

		}

		i++
	}

}
func BookInput(Books *[]book) {
	line()
	var temp book
	fmt.Printf("\nUnesite ime autora knjige \n")
	fmt.Scan(&temp.author)
	fmt.Printf("\nUnesite ime knjige \n")
	fmt.Scan(&temp.title)
	fmt.Printf("\nUnesite ID knjige \n")
	fmt.Scan(&temp.ID)
	temp.Status = "available"
	*Books = append(*Books, temp)
	line()
}

func ViewBook(Books *[]book) {
	fmt.Printf("\nList of books : \n")
	line()
	for i, val := range *Books {
		fmt.Printf("\nThe inex of the book is : %v\n", i)
		fmt.Printf("\nThe name of the book is : %v\n", val.title)
		fmt.Printf("\nThe author of the book is : %v\n", val.author)
		fmt.Printf("\nId of the book is : %v\n", val.ID)
		fmt.Printf("\nBook status is : %v\n", val.Status)
		line()
		i++
	}

}

func genref(Genres *[]genre) {

	var active bool = true
	for active {
		var option = Manu2("Genre")

		switch option {
		case 1:
			GenreInput(Genres)
		case 2:
			var Name string
			print("Name of the genre:")
			fmt.Scan(Name)
			ModifyDenre(Genres, Name)
		case 3:
			GenreView(Genres)
		case 4:
			GanreDellet(Genres)
		case 5:
			active = false

		default:
			error()
		}
	}
}
func GanreDellet(plc *[]genre) {
	var i int
	fmt.Print("\n Index OF the genre worker : ")
	fmt.Scan(&i)
	temp := make([]genre, 0)
	temp = append(temp, (*plc)[:i]...)
	temp = append(temp, (*plc)[i+1:]...)
	(*plc) = temp

}
func ModifyDenre(Genre *[]genre, Name string) {

	for i, val := range *Genre {

		if val.NameZanra == Name {
			line()
			fmt.Print("Name :")
			fmt.Scan(val.NameZanra)
			fmt.Print("\nSign : ")
			fmt.Scan(val.ZnakZanra)
			line()

			(*Genre)[i] = val
		}
		i++
	}

}
func GenreInput(Ganre *[]genre) {
	var temp genre
	line()
	fmt.Print(" \nName of the ganre:\n")
	fmt.Scan(&temp.NameZanra)
	fmt.Print("\n Sign of the ganre is:\n")
	fmt.Scan(&temp.ZnakZanra)
	line()
	*Ganre = append(*Ganre, temp)
}
func GenreView(Genres *[]genre) {
	fmt.Printf("\nView of genres\n")
	line()

	for i, val := range *Genres {
		fmt.Printf("\nGenre index is : %v \n", i)
		fmt.Printf("\nGenre NAME is : %v \n", val.NameZanra)
		fmt.Printf("\nGenre code is : %v \n\n", val.ZnakZanra)
		line()

		i++
	}

}
func shelff(Shelf *[]shelf) {

	var active = true
	for active {
		var option = Manu2("Shelf")

		switch option {

		case 1:
			ShelfInput(Shelf)
		case 2:
			var IND int
			fmt.Print("\nIndex of the shelf pls (int):")
			fmt.Scan(IND)
			ShelfModify(Shelf, IND)
		case 3:
			ViewShelf(Shelf)
		case 4:
			ShelfDellet(Shelf)
		case 5:
			active = false

		default:
			error()
		}
	}
}
func ShelfDellet(plc *[]shelf) {
	var i int
	fmt.Print("\n Index OF the targeted shelf : ")
	fmt.Scan(&i)
	temp := make([]shelf, 0)
	temp = append(temp, (*plc)[:i]...)
	temp = append(temp, (*plc)[i+1:]...)
	(*plc) = temp

}
func ShelfModify(Shelf *[]shelf, IND int) {
	line()

	for i, vla := range *Shelf {

		if vla.index == IND {
			fmt.Print("\nThe genre : ")
			fmt.Scan(vla.genre)
			(*Shelf)[i] = vla
			line()

		}

		i++
	}

}

func ShelfInput(Shelf *[]shelf) {
	line()

	var temp shelf
	fmt.Print("\nShelfs genre is :")
	fmt.Scan(&temp.genre)
	fmt.Print("\nShelfs number is :")
	fmt.Scan(&temp.index)
	*Shelf = append(*Shelf, temp)
}

func ViewShelf(Shelf *[]shelf) {

	fmt.Printf("\n\nList of shelves\n")
	line()
	for i, val := range *Shelf {
		fmt.Printf("The genre of the shelf is: %v\n", val.genre)
		fmt.Printf("Shelf index is : %v\n", val.index)
		line()
		i++
	}
}
