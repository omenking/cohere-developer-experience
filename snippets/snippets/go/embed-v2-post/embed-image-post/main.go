package main

import (
	"context"
	"log"

	cohere "github.com/cohere-ai/cohere-go/v2"
	client "github.com/cohere-ai/cohere-go/v2/client"
)

func main() {
	co := client.NewClient(client.WithToken("<<apiKey>>"))

	resp, err := co.V2.Embed(
		context.TODO(),
		&cohere.V2EmbedRequest{
			Images:         []string{"data:image/jpeg;base64,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"},
			Model:          "embed-english-v3.0",
			InputType:      cohere.EmbedInputTypeImage,
			EmbeddingTypes: []cohere.EmbeddingType{cohere.EmbeddingTypeFloat},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", resp)
}
