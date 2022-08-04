package c

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("c-discord.dockerfile", "dockerfiles", botName)
}

func Resources() string {
	return templates.Content("discord/c.md", "resources", "")
}

func MainCContent() string {
	return templates.Content("src/main.c", "discord-c", "")
}

func BWCContent(botName string) string {
	return templates.Content("src/botway.h", "discord-c", botName)
}

func RunPsFileContent() string {
	return templates.Content("run.ps1", "discord-c", "")
}