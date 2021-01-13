package main

import (
	"fmt"
	"net/http"

	"github.com/nurliman/knitto-elearning-video-backend/transcode"

	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
)

func main() {
	store := filestore.FileStore{
		Path: "./uploads",
	}

	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	config := tusd.Config{
		BasePath:                "/files/",
		StoreComposer:           composer,
		NotifyCompleteUploads:   true,
		RespectForwardedHeaders: true,
	}

	handler, err := tusd.NewHandler(config)
	if err != nil {
		panic(fmt.Errorf("Unable to create handler: %s", err))
	}

	go func() {
		for {
			event := <-handler.CompleteUploads
			fmt.Printf("Upload %s finished\n", event.Upload.ID)

			videoPath := fmt.Sprintf("./uploads/%s", event.Upload.ID)
			transcode.Video(&videoPath)
		}
	}()

	http.Handle("/files/", http.StripPrefix("/files/", handler))
	err = http.ListenAndServe(":3002", nil)
	if err != nil {
		panic(fmt.Errorf("Unable to listen: %s", err))
	}

}
