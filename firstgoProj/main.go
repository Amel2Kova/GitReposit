package main

import (
	"fmt"
)

/*
1.	Omogućiti operacije unosa, izmjene, pregleda i brisanja:članova, radnika, knjiga, književnog žanra, polica za knjige, kategorija članarina, te posuđenih knjiga.
2.	Unos novih stavki, te izmjenu postojećih vršiti sa tastature.
3.	Željenu radnju (unos, izmjena, pregled, brisanje) korisnik unosi preko tastature na osnovu ponuđenih opcija u terminalu.
4.	Za novi unos definirati potrebne kolone i ograničiti unos ukoliko se ne unesu potrebni podaci.
5.	Za povezivanje tabela koristit će se ključevi, koji predstavljaju jedinstveni identifikator određene stavke, kao npr:
-	radnik/član: tekstualni zapis JMBG od 13 cifara,
-	knjiga: tekstualni zapis od 5 znakova, gdje 1. znak označava žanr (B1010, K5425, itd.),
-	polica: tekstualni zapis sa oznakom žanra i numeracijom police;
6.	Napraviti dokumentaciju projekta u obliku Word dokumenta. Koristiti ovaj dokument kao šablon, te početi sa 5. poglavljem svoj izvještaj o projektu.
*/

type covjek struct {
	ime     string
	prezime string
	godiste string
	JMBG    string
}

type clan struct {
	covjek
	podignuteKnjige    []string
	limitNaDizanje     uint8
	trenutnoPodignutih uint8
}
type radnik struct {
	covjek
	pozicija string
	smjena   string
}
type knjiga struct {
	naziv  string
	autor  string
	ID     string
	Status string
}
type zanr struct {
	NazivZanra string
	ZnakZanra  string
}

type polica struct {
	zanr  string
	index int
}

var Clanovi = []clan{
	{
		covjek:             covjek{ime: "Amel", prezime: "Smajic", godiste: "2002", JMBG: "2102200286041"},
		podignuteKnjige:    []string{""},
		limitNaDizanje:     3,
		trenutnoPodignutih: 0,
	},
	{
		covjek: covjek{
			ime:     "Armin",
			prezime: "Avdic",
			godiste: "1999",
			JMBG:    "120419991234",
		},
		podignuteKnjige:    []string{},
		limitNaDizanje:     3,
		trenutnoPodignutih: 0,
	},
}

var Radnici = []radnik{
	{
		covjek: covjek{
			ime:     "Tim",
			prezime: "tibersom",
			godiste: "1982",
			JMBG:    "19820312310293",
		},
		pozicija: "Upravitelj",
		smjena:   "Prva",
	}, {
		covjek: covjek{
			ime:     "Andrey",
			prezime: "Robines",
			godiste: "1999",
			JMBG:    "2102199991234",
		},
		pozicija: "Bibliotekar",
		smjena:   "Prva",
	}, {
		covjek: covjek{
			ime:     "Michael",
			prezime: "Jeas",
			godiste: "1983",
			JMBG:    "120319839213",
		},
		pozicija: "Domar",
		smjena:   "Prva",
	},
}

var Knjige = []knjiga{
	{
		naziv:  "The Lion, the Witch and the Wardrobe",
		autor:  "C. S. Lewis",
		ID:     "",
		Status: "na raspolaganju",
	}, {
		naziv:  "She: A History of Adventure",
		autor:  "H. Rider Haggard",
		ID:     "",
		Status: "na raspolaganju",
	}, {
		naziv:  "The Da Vinci Code",
		autor:  "Den Brown",
		ID:     "",
		Status: "na raspolaganju",
	},
}
var Zanrovi = []zanr{
	{
		NazivZanra: "Fantazy",
		ZnakZanra:  "F",
	},
	{
		NazivZanra: "Adventure",
		ZnakZanra:  "A",
	},
	{
		NazivZanra: "History",
		ZnakZanra:  "H",
	},
	{
		NazivZanra: "Thriller",
		ZnakZanra:  "T",
	},
}

var Police = []polica{
	{
		zanr:  "Fantazy",
		index: 0000,
	},
	{
		zanr:  "Adventure",
		index: 0001,
	},
	{
		zanr:  "History",
		index: 0002,
	},
	{
		zanr:  "Thriller",
		index: 0003,
	},
}

func main() {
	var active = true

	for active {
		var opcija = 0
		//fmt.Printf("")
		fmt.Printf("Dobro dosli u biblioteku\n")
		fmt.Printf("Molimo odaberite bazu koju trebate urediti\n")
		fmt.Printf("1.) Korisnika\n")
		fmt.Printf("2.) Zaposlenika\n")
		fmt.Printf("3.) Knjiga\n")
		fmt.Printf("4.) Zanrova\n")
		fmt.Printf("5.) Polica\n")
		fmt.Printf("6.) Exit\n")

		fmt.Scan(&opcija)

		switch opcija {

		case 1:
			korisnikf()

		case 2:
			radnikf()

		case 3:
			knjigaf()

		case 4:
			zanrf()

		case 5:
			policaf()

		case 6:
			active = false

		default:
			errore()
		}
	}

}
func errore() {

	fmt.Printf("Molimo unesite tacan podatak \n")
}

func Manu2(baza string) (opcija2 int) {

	fmt.Printf("Dobro dosli u bazu %v \n", baza)
	fmt.Printf("Molimo odaberite zeljenu opciju\n")
	fmt.Printf("1.) Unesi %v\n", baza)
	fmt.Printf("2.) Izmijeni %v\n", baza)
	fmt.Printf("3.) Pregledaj %v\n", baza)
	fmt.Printf("4.) Izbrisi %v\n", baza)
	fmt.Printf("5.) Exit")
	fmt.Scan(&opcija2)

	return
}

func korisnikf() {

	var active bool = true
	for active {

		var opcija = Manu2("Korisnika")
		fmt.Printf("odabir je: %v", opcija)

		switch opcija {
		case 1:
			UnosKorisnika()
		case 2:

		case 3:
			PregledajKorisnike()
		case 4:

		case 5:
			active = false

		default:
			errore()
		}
	}
}
func line() {
	fmt.Printf("\n-------------------------------------------------------------------------------\n")

}
func PregledajKorisnike() {
	for i, val := range Clanovi {
		line()
		fmt.Printf("\nIme korisnika je: %v \n", val.ime)
		fmt.Printf("Prezime korisnika je: %v \n", val.prezime)
		fmt.Printf("Godiste korisnika je: %v \n", val.godiste)
		fmt.Printf("Jmbg korisnika je: %v \n", val.JMBG)
		if val.trenutnoPodignutih == 0 {
			fmt.Printf("Korisnik nema podignutih knjiga i moze da podigne jos %v knjiga\n\n", val.limitNaDizanje)

		} else {
			fmt.Printf("Korisnik ima %v podignutih knjiga i moze podignuti %v jos\n\n ", val.trenutnoPodignutih, (val.limitNaDizanje - val.trenutnoPodignutih))

		}
		line()
		i++
	}

}

func UnosKorisnika() {
	var n = len(Clanovi)
	fmt.Printf("\nMolimo unesite podatke novog Korisnika : \n")
	line()
	fmt.Printf("Ime korisnika : \n")
	fmt.Scanf(Clanovi[n].ime)
	fmt.Printf("Prezime korisnika : \n")
	fmt.Scanf(Clanovi[n].prezime)
	fmt.Printf("Godiste korisnika :\n")
	fmt.Scanf(Clanovi[n].godiste)
	fmt.Printf("JMBG korisnika : \n")
	fmt.Scanf(Clanovi[n].JMBG)
	fmt.Printf("Unesite limit na dizanje (int) : \n")
	fmt.Scan(Clanovi[n].limitNaDizanje)

}
func radnikf() {

	var active = true
	for active {
		var opcija = Manu2("Zaposlenika")

		switch opcija {

		case 1:

		case 2:

		case 3:
			pregledRadnika()
		case 4:

		case 5:
			active = false

		default:
			errore()
		}
	}
}

func pregledRadnika() {
	fmt.Printf("\nLista radnika:\n")
	line()

	for i, val := range Radnici {
		line()
		fmt.Printf("\nIme : %v\n", val.ime)
		fmt.Printf("Prezime : %v\n", val.prezime)
		fmt.Printf("Godiste : %v\n", val.godiste)
		fmt.Printf("Jmbg : %v\n", val.JMBG)
		fmt.Printf("Pozicja : %v\n", val.pozicija)
		fmt.Printf("Smjena : %v\n", val.smjena)
		line()
		i++
	}

}

func knjigaf() {

	var active = true
	for active {

		var opcija = Manu2("Knjigu")
		switch opcija {
		case 1:

		case 2:

		case 3:
			PregledKnjiga()
		case 4:

		case 5:
			active = false

		default:
			errore()
		}
	}
}

func PregledKnjiga() {
	fmt.Printf("\nLista knjiga : \n")
	line()
	for i, val := range Knjige {

		fmt.Printf("\nNaziv knjige je : %v\n", val.naziv)
		fmt.Printf("Autor knjige je : %v\n", val.autor)
		fmt.Printf("Id knjige je : %v\n", val.ID)
		fmt.Printf("Status knjige je : %v\n", val.Status)
		line()
		i++
	}

}

func zanrf() {

	var active bool = true
	for active {
		var opcija = Manu2("Zanr")

		switch opcija {
		case 1:

		case 2:

		case 3:
			PregledZanrova()
		case 4:

		case 5:
			active = false

		default:
			errore()
		}
	}
}

func PregledZanrova() {
	fmt.Printf("\nPregled zanrova\n")
	line()

	for i, val := range Zanrovi {
		fmt.Printf("NAziv zanra je : %v \n", val.NazivZanra)
		fmt.Printf("Kod zanra je : %v \n\n", val.ZnakZanra)
		line()

		i++
	}

}
func policaf() {

	var active = true
	for active {
		var opcija = Manu2("Policu")

		switch opcija {

		case 1:

		case 2:

		case 3:
			PregledPolica()
		case 4:

		case 5:
			active = false

		default:
			errore()
		}
	}
}
func PregledPolica() {

	fmt.Printf("\n\nPregled polica\n")
	line()
	for i, val := range Police {
		fmt.Printf("Zanr police je: %v\n", val.zanr)
		fmt.Printf("Index police je : %v\n", val.index)
		line()
		i++
	}
}
