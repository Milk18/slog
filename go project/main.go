package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Simple Go Web Page</title>
			<style>
				body {
					font-family: Arial, sans-serif;
				}
				img {
					max-width: 200px;
					height: auto;
				}
				#messages {
					margin-top: 10px;
				}
			</style>
			<script>
				function addMessage() {
					var inputBox = document.getElementById("textbox");
					var message = inputBox.value;
					if (message.trim() !== "") {
						var messagesDiv = document.getElementById("messages");
						var messageElement = document.createElement("p");
						var now = new Date();
						var timestamp = now.toLocaleString();
						messageElement.textContent = timestamp + " - " + message;
						messagesDiv.appendChild(messageElement);
						inputBox.value = "";
					}
				}
			</script>
		</head>
		<body>
			<h1>Welcome to SLOG</h1>
 			<h2>POST YOUR ISH</h2>
			<img src="https://th.bing.com/th/id/R.78f3df92984a4618118670cd017cc364?rik=DPnO4vs%2bkCXFZQ&riu=http%3a%2f%2fvignette2.wikia.nocookie.net%2fadventuretimewithfinnandjake%2fimages%2f3%2f3f%2fSnorlock.png%2frevision%2flatest%3fcb%3d20120912030125&ehk=fKYGg1Lq5MW6CzliICR7GhEt9PQ22BBXCtu%2b2pMyLEE%3d&risl=&pid=ImgRaw&r=0" alt="Placeholder Image">
			<br>
			<br>
			<label for="textbox">Enter Something:</label>
			<input type="text" id="textbox" name="textbox">
			<button onclick="addMessage()">Add Message</button>
			<div id="messages"></div>
		</body>
		</html>
		`
		t, err := template.New("webpage").Parse(tmpl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	})

	http.ListenAndServe(":5000", nil)
}
