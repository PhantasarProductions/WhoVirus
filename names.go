/*
	Who is the Virus?
	Names of files
	
	
	
	(c) Jeroen P. Broks, 2018, All rights reserved
	
		This program is free software: you can redistribute it and/or modify
		it under the terms of the GNU General Public License as published by
		the Free Software Foundation, either version 3 of the License, or
		(at your option) any later version.
		
		This program is distributed in the hope that it will be useful,
		but WITHOUT ANY WARRANTY; without even the implied warranty of
		MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
		GNU General Public License for more details.
		You should have received a copy of the GNU General Public License
		along with this program.  If not, see <http://www.gnu.org/licenses/>.
		
	Exceptions to the standard GNU license are available with Jeroen's written permission given prior 
	to the project the exceptions are needed for.
Version: 18.05.08
*/
package main

var boys = []string{
	"Achmed",
	"Anders",
	"Albus",
	"Adrianus",
	"Andre",
	"Andreas",
	"Bernhard",
	"Bernard",
	"Ben",
	"Bennie",
	"Bart",
	"Bill",
	"Bruce",
	"Chuck",
	"Chuckie",
	"Cornelius",
	"Dan",
	"Daniel",
	"Dick",
	"Dirk",
	"Dennis",
	"Donald",
	"Eduard",
	"Eduardo",
	"Eric",
	"Felix",
	"Ferdinand",
	"Fernando",
	"Fred",
	"Freddie",
	"Gerardus",
	"Gene",
	"Geoffrey",
	"George",
	"Han",
	"Hans",
	"Hank",
	"Harry",
	"Hendrik",
	"Henri",
	"Igor",
	"Isaac",
	"Ivo",
	"James",
	"Joseph",
	"Jeroen",
	"John",
	"Johannes",
	"Jacob",
	"Jake",
	"Jack",
	"Joris",
	"Jim",
	"Jimmy",
	"Jude",
	"Karl",
	"Kees",
	"Ken",
	"Kenny",
	"Klaas",
	"Kyle",
	"Leo",
	"Leonard",
	"Laurens",
	"Laurel",
	"Mario",
	"Marinus",
	"Marcus",
	"Marco",
	"Marc",
	"Marcel",
	"Mustafa",
	"Neville",
	"Nino",
	"Nils",
	"Nick",
	"Nico",
	"Nicholas",
	"Oliver",
	"Otto",
	"Peter",
	"Paul",
	"Paulus",
	"Pedro",
	"Quint",
	"Rhemus",
	"Ron",
	"Ronald",
	"Rudolf",
	"Roland",
	"Severus",
	"Samuel",
	"Simon",
	"Stephen",
	"Stephano",
	"Sirius",
	"Sjoerd",
	"Stan",
	"Theodore",
	"Theo",
	"Tinus",
	"Uli",
	"Victor",
	"Vladimir",
	"Valentino",
	"Wilhelmus",
	"Willem",
	"William",
	"Waldo",
	"Wally",
	"Wim",
	"Wout",
	"Wouter",
	"Xavier",
	"Youp",
	"Zachery",
	"Zack",
}

var girls = []string{
	"Anna",
	"Astrid",
	"Annelies",
	"Amanda",
	"Angela",
	"Angelina",
	"Angelique",
	"Antoinette",
	"Bertha",
	"Bella",
	"Belle",
	"Carola",
	"Carolina",
	"Cecilia",
	"Celestine",
	"Celestina",
	"Celeste",
	"Carla",
	"Carleyn",
	"Candy",
	"Denise",
	"Donna",
	"Daisy",
	"Ellen",
	"Ella",
	"Emma",
	"Eve",
	"Eva",
	"Elena",
	"Esmeralda",
	"Esther",
	"Fleur",
	"Frederica",
	"Francesca",
	"Franka",
	"Felicia",
	"Georgia",
	"Georgina",
	"Gerardina",
	"Gerda",
	"Gina",
	"Hanna",
	"Hermione",
	"Ilse",
	"Inge",
	"Jill",
	"Josephine",
	"Johanna",
	"Jolanda",
	"Judith",
	"Judy",
	"Karla",
	"Laverne",
	"Lea",
	"Lia",
	"Maria",
	"Mary",
	"Mindy",
	"Mandy",
	"Merel",
	"Mina",
	"Madelief",
	"Monica",
	"Monique",
	"Nancy",
	"Nel",
	"Nicole",
	"Nicolette",
	"Nana",
	"Nina",
	"Olga",
	"Penelope",
	"Petra",
	"Paula",
	"Quinty",
	"Rebecca",
	"Rina",
	"Ria",
	"Renata",
	"Renate",
	"Roberta",
	"Sabrine",
	"Sabrina",
	"Samantha",
	"Shelly",
	"Simone",
	"Stephanie",
	"Theresa",
	"Thea",
	"Tina",
	"Ursula",
	"Ulla",
	"Valentina",
	"Victoria",
	"Vicky",
	"Vicxy",
	"Wanda",
	"Walda",
	"Wilhelmina",
	"Wilma",
	"Wendy",
	"Xantippe",
	"Yolanthe",
	"Yara",
	"Zelda",
}


func namelist()([]string,map[string]string){
	sexlist:=[]string{"M","F"} //[][]string{boys,girls}
	allnames:=[]string{}
	sexnames:=map[string]string{}
	for _,sex:=range sexlist{
		asexlist:=boys
		if sex=="F" { asexlist=girls }
		for _,name:=range asexlist{
			allnames=append(allnames,name)
			sexnames[name]=sex
		}
	}
	return allnames,sexnames
}

var Names,Sex = namelist()

func init(){
   if len(Names)<100 { panic("I need at least 100 names") }
}

