package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/gagliardetto/solana-go"
)

var BotToken string
var Channel string

var publicKey solana.PublicKey

func Run() {
	fmt.Println("channel", Channel)
	fmt.Println("bot token", BotToken)
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(askAddress)
	dg.ChannelMessageSend(Channel, "What address do you want to watch?")

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func askAddress(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.ChannelID != Channel || m.Author.ID == s.State.User.ID || publicKey.String() != "11111111111111111111111111111111" {
		return
	}

	//check if it is a valid address
	pk, err := solana.PublicKeyFromBase58(m.Content)
	if err != nil {
		fmt.Println("error saving sol address")
		s.ChannelMessageSend(Channel, "Not a valid solana address, try again!")
		return
	}
	publicKey = pk
	s.ChannelMessageSend(Channel, "Address saved, "+publicKey.String())
	s.ChannelMessageSend(Channel, "Watching for trades")
}
