package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMain(t *testing.T) {
	resp, err := http.Get("http://localhost:12345/meeting/?id=5f4dcb738b246dc74d8ecd44")
	if err != nil {
		t.Error("Fail")
	}
	if resp == nil {
		t.Error("NO response")
	}
}

func BenchmarkMaingetmeet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		http.Get("http://localhost:12345/meeting/?id=5f4dcc74fa1a4b2011daf69a")
	}
}

func BenchmarkMaingetparticipant(b *testing.B) {
	for n := 0; n < b.N; n++ {
		http.Get("http://localhost:12345/articles/?participant=arnavdixit127@gmail.com")
	}
}

func BenchmarkMaingetpost(b *testing.B) {
	var message Meeting
	var part participant
	part.Name = "Arnav Dixit"
	part.Email = "arnavdixit127@gmail.com"
	part.Rsvp = "No"
	message.Title = "Title"
	message.Participants = append(message.Participants, part)
	message.Starttime = "2021-09-01T09:52:12+05:30"
	message.Endtime = "2021-09-01T10:52:12+05:30"

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Erur")
	}
	for n := 0; n < b.N; n++ {
		resp, err := http.Post("http://localhost:12345/meetings", "application/json", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			b.Error("Fail")
		}
		if resp == nil {
			b.Error("NO response")
		}
	}
}

func BenchmarkParticipantsBusy(b *testing.B) {
	var message Meeting
	var part participant
	part.Name = "Arnav Dixit"
	part.Email = "arnavdixit127@gmail.com"
	part.Rsvp = "No"
	message.Title = "Title"
	message.Participants = append(message.Participants, part)
	message.Starttime = "2021-09-01T09:52:12+05:30"
	message.Endtime = "2021-09-01T10:52:12+05:30"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, _ = mongo.Connect(ctx, clientOptions)
	for n := 0; n < b.N; n++ {
		ParticipantsBusy(message)
	}
}
