package main

import (
	"fmt"
	"math"
	"net/http"
	"runtime"
	"strconv"
)

func bToMB(b uint64) float64 { return float64(b) / 1024.0 / 1024.0 }

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %.2f MiB", bToMB(m.Alloc))
	fmt.Printf("\tSys = %.2f MiB", bToMB(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		for i := 0; i < 100; i++ {
			getPerson()
		}
		person := getPerson()
		//time.Sleep(100 * time.Millisecond)
		val := 1.0
		for i := 0; i < 1000000; i++ {
			val = math.Cos(val)*math.Sin(val) + 1.0
		}
		//fmt.Println(val)
		fmt.Fprintln(w, person.Name+"\n"+strconv.Itoa(person.Age)+"\n"+person.Description)
		printMemUsage()
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func getPerson() Person {
	friends := make([]*Person, 0, 5)
	for i := 0; i < 5; i++ {
		friend := &Person{
			Name: fmt.Sprintf("友達%d", i),
			Age:  20 + i,
			Description: `Hello, my name is Oliver Green, and I'd like to share a bit about myself! 
I'm a 29-year-old software enthusiast with a passion for designing efficient systems 
and learning about emerging technologies. I grew up in a suburban neighborhood, 
surrounded by friendly neighbors and an active local community. As a child, I spent countless hours 
tinkering with computers, exploring coding challenges, and creating small programs that helped 
automate tedious tasks. Over time, this curiosity evolved into a full-blown interest in problem-solving, 
which motivated me to pursue a degree in computer science.

In my spare time, I enjoy hiking through scenic trails, reading thought-provoking books, 
and cooking experimental dishes in my cozy kitchen. I also love attending local meetups 
where fellow developers gather to chat about everything from robust backend infrastructures 
to user-friendly interface design. One of my biggest aspirations is to contribute to open-source projects 
that empower people from various backgrounds to learn programming skills.

I consider persistence, adaptability, and a sense of humor my strongest assets. 
If a project hits a roadblock, I relish the challenge of diagnosing the issue 
and iterating on potential solutions until the bug is squashed. 
Working in collaborative teams energizes me, and I appreciate the chance to exchange ideas 
that push our work closer to brilliance. Ultimately, I believe in continuous learning, 
embracing new experiences, and striving to create tools that can benefit communities both large and small.`,
			Friends: nil,
		}
		friends = append(friends, friend)
	}
	person := Person{
		Name: "WadaTakamichi",
		Age:  26,
		Description: "ぎおんしょうじゃのかねのこえ、しょぎょうむじょうのひびきあり。 " +
			"さらそうじゅのはなのいろ、 じょうしゃひっすいのことわりをあらわす。 " +
			"おごれるひともひさしからず、 ただはるのよのゆめのごとし。 " +
			"たけきものもついにはほろびぬ、 ひとえにかぜのまえのちりにおなじ。" +
			"とおくいちょうをとぶらえば、 しんのちょうこう、かんのおうもう、りょうのしゅうい、とうのろくさん、 " +
			"これらはみな、きゅうしゅせんこうのまつりごとにもしたがわず、 たのしみをきわめ、いさめをもおもいいれず、 " +
			"てんかのみだれんことをさとらずして、 みんかんのうれうるところをしらざつしかば、 ひさしからずしてぼうじにしものどもなり。" +
			"ちかくほんちょうをうかがうに、 じょうへいのまさかど、てんぎょうのすみとも、こうわのぎしん、へいじののぶより、 " +
			"これらはおごれるこころも、たけきこともみなとりどりにこそありしかども、まぢかくは、 " +
			"ろくはらのにゅうどう、さきのだいじょうだいじん、たいらのあそんきよもりこうともうししひとのありさま、 " +
			"つたえうけたまわるこそ、こころもことばもおよばれね。" +
			"そのせんぞをたずぬればかんむてんのうだいごのおうじ、 いっぽんしきぶきょうかずらはらしんのうくだいのこういん、さぬきのかみまさもりがまご、 " +
			"ぎょうぶきょうただもりのあそんのちゃくなんなり。 かのしんのうのみこ、たかみのおう、むかんむいにしてうせたまいぬ。 " +
			"そのみこ、たかもちのおうのとき、はじめてへいのしょうをたまわって、 かずさのすけになりたまいしより、たちまちにおうしをいでてじんしんにつらなる。 " +
			"そのこちんじゅふのしょうぐんよしもち、のちにはくにかとあらたむ。 くにかよりまさもりにいたるろくだいは、しょこくのずりょうたりしかども、 " +
			"てんじょうのせんせきをばいまだゆるされず。",
		Friends: friends,
	}
	return person
}

type Person struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Description string `json:"description"`
	Friends     []*Person
}
