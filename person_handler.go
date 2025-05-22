package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		person := getPerson()
		fmt.Fprintln(w, person.Name+"\n"+strconv.Itoa(person.Age)+"\n"+person.Description)
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func getPerson() Person {
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
	}
	return person
}

type Person struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Description string `json:"description"`
}
