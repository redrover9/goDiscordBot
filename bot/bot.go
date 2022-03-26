package bot

import (
	"discordBot/config"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
)

var BotId string
var goBot *discordgo.Session
var qAndA *qAndAStruct

type qAndAStruct struct {
	Question string `json:"Question"`
	Answer   string `json:"Answer"`
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
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}

func getQuestion() {
	fmt.Println("Reading question file...")
	file, err := ioutil.ReadFile("C:/Users/grace/GolandProjects/discordBot/bot/questions.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(file, &qAndA)

	if err != nil {
		fmt.Println(err.Error())
	}

	Question = qAndA.Question
	Answer = qAndA.Answer

}
