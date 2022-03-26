package bot

import (
	"discordBot/config"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
)

var a string
var BotId string
var goBot *discordgo.Session
var qAndA *qAndAStruct

type qAndAStruct struct {
	Question string `json:"Which command enables you to redirect stdout and stderr to a file?"`
	Answer   string `json:"> filename 2>&1"`
}

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}
	if m.Content == "!question" {
		quest, ans := getQuestion()
		_, _ = s.ChannelMessageSend(m.ChannelID, quest)
		a = ans
	}

	if m.Content == "!answer" {
		_, _ = s.ChannelMessageSend(m.ChannelID, a)
	}
}

func getQuestion() (string, string) {
	jsonFile, err := ioutil.ReadFile("C:/Users/grace/GolandProjects/discordBot/bot/questions.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	m := map[string]string{}
	err = json.Unmarshal(jsonFile, &m)
	if err != nil {
		fmt.Println(err.Error())
	}
	var quest string
	var ans string
	for q, a := range m { //a := range m {
		quest = q
		ans = a
		break
	}
	return quest, ans
}
