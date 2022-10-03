package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}
	Token := "Bot " + os.Getenv("APP_BOT_TOKEN")
	BotName := "<@" + os.Getenv("CLIENT_ID") + ">"

	fmt.Println(Token)
	fmt.Println(BotName)

	discord, err := discordgo.New(Token)
	if err != nil {
		fmt.Println("Fail to login.")
		log.Fatal(err)
	}

	discord.AddHandler(onMessageCreate)
	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer discord.Close()

	fmt.Println("Listening...")

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot

	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	clientId := os.Getenv("CLIENT_ID")
	u := m.Author
	fmt.Println(m.ChannelID + " " + u.Username + "(" + u.ID + ")>" + m.Content)

	if u.ID != clientId {
		sendMessage(s, m.ChannelID, u.Mention()+"は？")
		// sendReply(s, m.ChannelID, "testTest", m.Reference())
	}
}

func sendMessage(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)
	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
 }
 
 func sendReply(s *discordgo.Session, channelID string, msg string, reference *discordgo.MessageReference) {
	_, err := s.ChannelMessageSendReply(channelID, msg, reference)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
 }

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Fail to load .env file.")
		return err
	}

	return nil

}
