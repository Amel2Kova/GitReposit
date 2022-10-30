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
	podignuteKnjige []string
	limitNaDizanje  =3
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

	zanr
	index int

}

var Clanovi =[]clan {
	clan{
		covjek:             covjek{ime: "Amel", prezime: "Smajic", godiste: "2002", JMBG: "2102200286041"},
		podignuteKnjige:    []string{""},
		limitNaDizanje:     3,
		trenutnoPodignutih: 0,
	},
	clan{
		covjek:             covjek{
			ime:     "Armin",
			prezime: "Avdic",
			godiste: "1999",
			JMBG:    "120419991234",
		},
		podignuteKnjige:    []string{},
		limitNaDizanje:     3,
		trenutnoPodignutih: 0,
	}



}



func main() {
	var active=1




	for 1==active {
		var opcija = 0
		// fmt.Printf("")
		fmt.Printf("Dobro dosli u biblioteku\n")
		fmt.Printf("Molimo odaberite bazu koju trebate urediti\n")
		fmt.Printf("1.) Korisnika")
		fmt.Printf("2.) Zaposlenika")
		fmt.Printf("3.) Knjiga")
		fmt.Printf("4.) Zanrova")
		fmt.Printf("5.) Polica")
		fmt.Printf("6.) Exit")
		fmt.Scan(opcija)
		switch opcija {
		odluka2 int

		case 1:
		koeisnik()
			break
		case 2:
			radnik()
			break
		case 3:
			knjiga()
			break
		case 4:
			zanr()
			break
		case 5:
			polica()
			break
		case 6:
			active=0
			break
		default		
		errore()
		}
	}
}
func errore(){


	fmt.Printf("Molimo unesite tacan podatak \n")
}

func Manu2 (baza string)(opcija2 int){

	fmt.Printf("Dobro dosli u bazu",baza,"\n")
	fmt.Printf("Molimo odaberite zeljenu opciju\n")
	fmt.Printf("1.) Unesi",baza,"\n")
	fmt.Printf("2.) Izmijeni",baza,"\n")
	fmt.Printf("3.) Pregledaj",baza,"\n")
	fmt.Printf("4.) Izbrisi",baza,"\n")
	fmt.Printf("5.) Exit")
	fmt.Scan(opcija2)

	return opcija2
}

func korisnik(){
	var opcija=Manu2("Korisnika")
	
	var active =1
	switch opcija {
	case 1:

		break
	case 2:
		break
	case 3:
		break
	case 4:
		break
	case 5:
		active=0
		break
	default:
		errore()
	}
}

func radnik(){

	opcija=Manu2("Zaposlenika")
	var active =1
	switch opcija {
	case 1:
		break
	case 2:
		break
	case 3:
		break
	case 4:
		break
	case 5:
		active=0
		break
	default:
		errore()
	}
}

func knjiga(){

	opcija=Manu2("Knjigu")
	var active =1
	switch opcija {
	case 1:
		break
	case 2:
		break
	case 3:
		break
	case 4:
		break
	case 5:
		active=0
		break
	default:
		errore()
	}
}

func zanr(){


	opcija=Manu2("Zanr")
	var active =1
	switch opcija {
	case 1:
		break
	case 2:
		break
	case 3:
		break
	case 4:
		break
	case 5:
		active=0
		break
	default:
		errore()
	}
}

func polica(){


	Manu2("Policu")
	var active =1
	switch opcija {
	case 1:
		break
	case 2:
		break
	case 3:
		break
	case 4:
		break
	case 5:
		active=0
		break
	default:
		errore()
	}
	
}