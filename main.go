package main

import (
	"fmt"
	"os"
	"strings"
)

type person struct {
	Name string
	Info string
	Url  string
}

var people = map[string]*person{
	"albattani":     &person{"Muhammad ibn Jābir al-Ḥarrānī al-Battānī", "Founding father of astronomy", "https://en.wikipedia.org/wiki/Mu%E1%B8%A5ammad_ibn_J%C4%81bir_al-%E1%B8%A4arr%C4%81n%C4%AB_al-Batt%C4%81n%C4%AB"},
	"almeida":       &person{"June Almeida", "Scottish virologist who took the first pictures of the rubella virus", "https://en.wikipedia.org/wiki/June_Almeida"},
	"archimedes":    &person{"Archimedes", "Physicist, engineer and mathematician who invented too many things to list them here", "https://en.wikipedia.org/wiki/Archimedes"},
	"ardinghelli":   &person{"Maria Ardinghelli", "Italian translator, mathematician and physicist", "https://en.wikipedia.org/wiki/Maria_Ardinghelli"},
	"babbage":       &person{"Charles Babbage", "invented the concept of a programmable computer", "https://en.wikipedia.org/wiki/Charles_Babbage"},
	"banach":        &person{"Stefan Banach", "Polish mathematician, was one of the founders of modern functional analysis", "https://en.wikipedia.org/wiki/Stefan_Banach"},
	"barden":        &person{"John Bardeen", "Co-invented the transistor with William Shockley and Walter Houser Brattain", "https://en.wikipedia.org/wiki/John_Bardeen"},
	"brattain":      &person{"Walter Houser Brattain", "Co-invented the transistor with William Shockley and Walter Houser", "https://en.wikipedia.org/wiki/Walter_Houser_Brattain"},
	"shockley":      &person{"William Shockley", "Co-invented the transistor with William Shockley and Walter Houser", "https://en.wikipedia.org/wiki/William_Shockley"},
	"bartik":        &person{"Jean Bartik", "Born Betty Jean Jennings, was one of the original programmers for the ENIAC computer", "https://en.wikipedia.org/wiki/Jean_Bartik"},
	"bell":          &person{"Alexander Graham Bell", "Eminent Scottish-born scientist, inventor, engineer and innovator who is credited with inventing the first practical telephone", "https://en.wikipedia.org/wiki/Alexander_Graham_Bell"},
	"blackwell":     &person{"Elizabeth Blackwell", "American doctor and first American woman to receive a medical degree", "https://en.wikipedia.org/wiki/Elizabeth_Blackwell"},
	"boh":           &person{"Niels Bohr", "The father of quantum theory", "https://en.wikipedia.org/wiki/Niels_Bohr"},
	"brown":         &person{"Emmet Brown", "Invented time travel", "https://en.wikipedia.org/wiki/Emmett_Brown"},
	"carson":        &person{"Rachel Carson", "American marine biologist and conservationist, her book Silent Spring and other writings are credited with advancing the global environmental movement", "https://en.wikipedia.org/wiki/Rachel_Carson"},
	"colden":        &person{"Jane Colden", "American botanist widely considered the first female American botanist", "https://en.wikipedia.org/wiki/Jane_Colden"},
	"cori":          &person{"Gerty Theresa Cori", "American biochemist who became the third woman—and first American woman—to win a Nobel Prize in science, and the first woman to be awarded the Nobel Prize in Physiology or Medicine. Cori was born in Prague", "https://en.wikipedia.org/wiki/Gerty_Cori"},
	"curie":         &person{"Marie Curie", "Discovered radioactivity", "https://en.wikipedia.org/wiki/Marie_Curie."},
	"darwin":        &person{"Charles Darwin", "Established the principles of natural evolution", "https://en.wikipedia.org/wiki/Charles_Darwin"},
	"einstein":      &person{"Albert Einstein", "invented the general theory of relativity", "https://en.wikipedia.org/wiki/Albert_Einstein"},
	"elion":         &person{"Gertrude Elion", "American biochemist, pharmacologist and the 1988 recipient of the Nobel Prize in Medicine", "https://en.wikipedia.org/wiki/Gertrude_Elion"},
	"engelbart":     &person{"Douglas Engelbart", "Gave the mother of all demos", "https://en.wikipedia.org/wiki/Douglas_Engelbart"},
	"euclid":        &person{"Euclid", "Invented geometry", "https://en.wikipedia.org/wiki/Euclid"},
	"fermat":        &person{"Pierre de Fermat", "pioneered several aspects of modern mathematics", "https://en.wikipedia.org/wiki/Pierre_de_Fermat"},
	"fermi":         &person{"Enrico Fermi", "Invented the first nuclear reactor", "https://en.wikipedia.org/wiki/Enrico_Fermi"},
	"feynman":       &person{"Richard Feynman", "Key contributor to quantum mechanics and particle physics", "https://en.wikipedia.org/wiki/Richard_Feynman"},
	"franklin":      &person{"Benjamin Franklin", "Famous for his experiments in electricity and the invention of the lightning rod", "http://en.wikipedia.org/wiki/Benjamin_Franklin"},
	"galileo":       &person{"Galileo Galilei", "Founding father of modern astronomy, and faced politics and obscurantism to establish scientific truth", "https://en.wikipedia.org/wiki/Galileo_Galilei"},
	"goldstine":     &person{"Adele Goldstine", "Born Adele Katz, wrote the complete technical description for the first electronic digital computer, ENIAC", "https://en.wikipedia.org/wiki/Adele_Goldstine"},
	"goodall":       &person{"Jane Goodall", "British primatologist, ethologist, and anthropologist who is considered to be the world's foremost expert on chimpanzees", "https://en.wikipedia.org/wiki/Jane_Goodall"},
	"hawking":       &person{"Stephen Hawking", "Pioneered the field of cosmology by combining general relativity and quantum mechanics", "https://en.wikipedia.org/wiki/Stephen_Hawking"},
	"heisenberg":    &person{"Werner Heisenberg", "Founding father of quantum mechanics", "https://en.wikipedia.org/wiki/Werner_Heisenberg"},
	"hodgkin":       &person{"Dorothy Hodgkin", "British biochemist, credited with the development of protein crystallography. She was awarded the Nobel Prize in Chemistry in 1964", "https://en.wikipedia.org/wiki/Dorothy_Hodgkin"},
	"hoover":        &person{"Erna Schneider Hoover", "Revolutionized modern communication by inventing a computerized telephon switching method", "https://en.wikipedia.org/wiki/Erna_Schneider_Hoover"},
	"hopper":        &person{"Grace Hopper", "Developed the first compiler for a computer programming language and is credited with popularizing the term 'debugging' for fixing computer glitches", "https://en.wikipedia.org/wiki/Grace_Hopper"},
	"hypatia":       &person{"Hypatia", "Greek Alexandrine Neoplatonist philosopher in Egypt who was one of the earliest mothers of mathematics", "https://en.wikipedia.org/wiki/Hypatia"},
	"jang":          &person{"Yeong-Sil Jang", "Korean scientist and astronomer during the Joseon Dynasty; he invented the first metal printing press and water gauge", "https://en.wikipedia.org/wiki/Jang_Yeong-sil"},
	"jones":         &person{"Karen Spärck Jones", "Came up with the concept of inverse document frequency, which is used in most search engines today", "https://en.wikipedia.org/wiki/Karen_Sp%C3%A4rck_Jones"},
	"kirch":         &person{"Maria Kirch", "German astronomer and first woman to discover a comet", "https://en.wikipedia.org/wiki/Maria_Margarethe_Kirch"},
	"kowalevski":    &person{"Sophie Kowalevski", "Russian mathematician responsible for important original contributions to analysis, differential equations and mechanics", "https://en.wikipedia.org/wiki/Sofia_Kovalevskaya"},
	"lalande":       &person{"Marie-Jeanne de Lalande", "French astronomer, mathematician and cataloguer of stars", "https://en.wikipedia.org/wiki/Marie-Jeanne_de_Lalande"},
	"leakey":        &person{"Mary Leakey", "British paleoanthropologist who discovered the first fossilized Proconsul skull", "https://en.wikipedia.org/wiki/Mary_Leakey"},
	"lovelace":      &person{"Ada Lovelace", "Invented the first algorithm", "https://en.wikipedia.org/wiki/Ada_Lovelace"},
	"lumiere":       &person{"Auguste and Louis Lumière", "The first filmmakers in history", "https://en.wikipedia.org/wiki/Auguste_and_Louis_Lumi%C3%A8re"},
	"mayer":         &person{"Maria Mayer", "American theoretical physicist and Nobel laureate in Physics for proposing the nuclear shell model of the atomic nucleus", "https://en.wikipedia.org/wiki/Maria_Mayer"},
	"mccarthy":      &person{"John McCarthy", "Invented LISP", "https://en.wikipedia.org/wiki/John_McCarthy_(computer_scientist)"},
	"mcclintock":    &person{"Barbara McClintock", "A distinguished American cytogeneticist, 1983 Nobel Laureate in Physiology or Medicine for discovering transposons", "https://en.wikipedia.org/wiki/Barbara_McClintock"},
	"mclean":        &person{"Malcolm McLean", "Invented the modern shipping container", "https://en.wikipedia.org/wiki/Malcom_McLean"},
	"meitner":       &person{"Lise Meitner", "Austrian/Swedish physicist who was involved in the discovery of nuclear fission. The element meitnerium is named after her", "https://en.wikipedia.org/wiki/Lise_Meitner"},
	"mestorf":       &person{"Johanna Mestorf", "German prehistoric archaeologist and first female museum director in Germany", "https://en.wikipedia.org/wiki/Johanna_Mestorf"},
	"morse":         &person{"Samuel Morse", "contributed to the invention of a single-wire telegraph system based on European telegraphs and was a co-developer of the Morse code", "https://en.wikipedia.org/wiki/Samuel_Morse"},
	"newton":        &person{"Isaac Newton", "Invented classic mechanics and modern optics", "https://en.wikipedia.org/wiki/Isaac_Newton"},
	"nobel":         &person{"Alfred Nobel", "Swedish chemist, engineer, innovator, and armaments manufacturer (inventor of dynamite)", "https://en.wikipedia.org/wiki/Alfred_Nobel"},
	"pare":          &person{"Ambroise Pare", "Invented modern surgery", "https://en.wikipedia.org/wiki/Ambroise_Par%C3%A9"},
	"pasteur":       &person{"Louis Pasteur", "Discovered vaccination, fermentation and pasteurization", "https://en.wikipedia.org/wiki/Louis_Pasteur"},
	"perlman":       &person{"Radia Perlman", "Software designer and network engineer and most famous for her invention of the spanning-tree protocol (STP)", "https://en.wikipedia.org/wiki/Radia_Perlman"},
	"pike":          &person{"Rob Pike", "Key contributor to Unix, Plan 9, the X graphic system, utf-8, and the Go programming language", "https://en.wikipedia.org/wiki/Rob_Pike"},
	"poincare":      &person{"Henri Poincaré", "Made fundamental contributions in several fields of mathematics", "https://en.wikipedia.org/wiki/Henri_Poincar%C3%A9"},
	"poitras":       &person{"Laura Poitras", "Director and producer whose work, made possible by open source crypto tools, advances the causes of truth and freedom of information by reporting disclosures by whistleblowers such as Edward Snowden", "https://en.wikipedia.org/wiki/Laura_Poitras"},
	"ptolemy":       &person{"Claudius Ptolemy", "Greco-Egyptian writer of Alexandria, known as a mathematician, astronomer, geographer, astrologer, and poet of a single epigram in the Greek Anthology", "https://en.wikipedia.org/wiki/Ptolemy"},
	"ritchie":       &person{"Dennis Ritchie", "Co-created created UNIX and the C programming language with Ken Thompson", "https://en.wikipedia.org/wiki/Dennis_Ritchie"},
	"thompson":      &person{"Ken Thompson", "Co-created created UNIX and the C programming language with Dennis Ritchie", "https://en.wikipedia.org/wiki/Ken_Thompson"},
	"rosalind":      &person{"Rosalind Franklin", "British biophysicist and X-ray crystallographer whose research was critical to the understanding of DNA", "https://en.wikipedia.org/wiki/Rosalind_Franklin"},
	"sammet":        &person{"Jean E. Sammet", "Developed FORMAC, the first widely used computer language for symbolic manipulation of mathematical formulas", "https://en.wikipedia.org/wiki/Jean_E._Sammet"},
	"sinoussi":      &person{"Françoise Barré-Sinoussi", "French virologist and Nobel Prize Laureate in Physiology or Medicine; her work was fundamental in identifying HIV as the cause of AIDS", "https://en.wikipedia.org/wiki/Fran%C3%A7oise_Barr%C3%A9-Sinoussi"},
	"stallman":      &person{"Richard Matthew Stallman", "Founder of the Free Software movement, the GNU project, the Free Software Foundation, and the League for Programming Freedom. He also invented the concept of copyleft to protect the ideals of this movement, and enshrined this concept in the widely-used GPL (General Public License) for software", "https://en.wikiquote.org/wiki/Richard_Stallman"},
	"swartz":        &person{"Aaron Swartz", "Influential in creating RSS, Markdown, Creative Commons, Reddit, and much of the internet as we know it today. He was devoted to freedom of information on the web", "https://en.wikiquote.org/wiki/Aaron_Swartz"},
	"tesla":         &person{"Nikola Tesla", "Invented the AC electric system and every gadget ever used by a James Bond villain", "https://en.wikipedia.org/wiki/Nikola_Tesla"},
	"torvalds":      &person{"Linus Torvalds", "Invented Linux and Git", "https://en.wikipedia.org/wiki/Linus_Torvalds"},
	"turing":        &person{"Alan Turing", "Founding father of computer science", "https://en.wikipedia.org/wiki/Alan_Turing"},
	"wilson":        &person{"Sophie Wilson", "Designed the first Acorn Micro-Computer and the instruction set for ARM processors", "https://en.wikipedia.org/wiki/Sophie_Wilson"},
	"wozniak":       &person{"Steve Wozniak", "Invented the Apple I and Apple II", "https://en.wikipedia.org/wiki/Steve_Wozniak"},
	"wright":        &person{"The Wright brothers, Orville and Wilbur", "Credited with inventing and building the world's first successful airplane and making the first controlled, powered and sustained heavier-than-air human flight", "https://en.wikipedia.org/wiki/Wright_brothers"},
	"yalow":         &person{"Rosalyn Sussman Yalow", "American medical physicist, and a co-winner of the 1977 Nobel Prize in Physiology or Medicine for development of the radioimmunoassay technique", "https://en.wikipedia.org/wiki/Rosalyn_Sussman_Yalow"},
	"yonath":        &person{"Ada Yonath", "Israeli crystallographer, the first woman from the Middle East to win a Nobel prize in the sciences", "https://en.wikipedia.org/wiki/Ada_Yonath"},
	"payne":         &person{"Cecilia Payne-Gaposchkin", "Astronomer and astrophysicist who, in 1925, proposed in her Ph.D. thesis an explanation for the composition of stars in terms of the relative abundances of hydrogen and helium.", "https://en.wikipedia.org/wiki/Cecilia_Payne-Gaposchkin"},
	"bose":          &person{"Satyendra Nath Bose", "Pprovided the foundation for Bose–Einstein statistics and the theory of the Bose–Einstein condensate", "https://en.wikipedia.org/wiki/Satyendra_Nath_Bose"},
	"raman":         &person{"C. V. Raman", "Indian physicist who won the Nobel Prize in 1930 for proposing the Raman effect", "https://en.wikipedia.org/wiki/C._V._Raman"},
	"ramanujan":     &person{"Srinivasa Ramanujan", "Indian mathematician and autodidact who made extraordinary contributions to mathematical analysis, number theory, infinite series, and continued fractions", "https://en.wikipedia.org/wiki/Srinivasa_Ramanujan"},
	"chandrasekhar": &person{"Subrahmanyan Chandrasekhar", "Astrophysicist known for his mathematical theory on different stages and evolution in structures of the stars. He has won nobel prize for physics", "https://en.wikipedia.org/wiki/Subrahmanyan_Chandrasekhar"},
	"khorana":       &person{" HarGobind Khorana", "Indian-American biochemist who shared the 1968 Nobel Prize for Physiology", "https://en.wikipedia.org/wiki/Har_Gobind_Khorana"},
	"saha":          &person{"Meghnad Saha", "Indian astrophysicist best known for his development of the Saha equation, used to describe chemical and physical conditions in stars", "https://en.wikipedia.org/wiki/Meghnad_Saha"},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "missing required argument, you must pass in a name\n")
		os.Exit(1)
	}
	name := os.Args[1]
	if strings.Contains(name, "_") {
		nameArr := strings.SplitN(name, "_", 2)
		if len(nameArr) < 2 {
			fmt.Fprintf(os.Stderr, "invalid name: %s", name)
		}
		name = nameArr[1]
	}
	person := people[name]
	if person == nil {
		fmt.Fprintf(os.Stderr, "Could not find anyone for %s\nPerhaps we are missing a name? Make a PR at github.com/cpuguy83/docker-who\nFor the official list see https://github.com/docker/docker/blob/master/pkg/namesgenerator/names-generator.go\n", name)
		os.Exit(1)
	}

	fmt.Printf("%s\n%s\nFind out more here: %s\n", person.Name, person.Info, person.Url)
}
