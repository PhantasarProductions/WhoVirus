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
	"Bill",
	"Cornelius",
	"Dick",
	"Dennis",
	"Donald",
	"Eduard",
	"Eduardo",
	"Eric",
	"Felix",
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
	"Xavier",
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
	"Bertha",
	"Bella",
	"Belle",
	"Carola",
	"Carolina",
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

