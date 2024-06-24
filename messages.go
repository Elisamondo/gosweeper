package main

var messages = []string{}

func getMessages()([]string){
	return messages
}

func setMessages(newValue string){
	messages = append(messages, newValue)
}

func clearMessages(){
	messages = nil
}