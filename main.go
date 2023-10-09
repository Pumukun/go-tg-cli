package main

import (
    "fmt"
    "time"
    
    "github.com/Arman92/go-tdlib"
)

func main() {
	tdlib.SetLogVerbosityLevel(1)
	tdlib.SetFilePath("./errors.txt")
	
	client := tdlib.NewClient(tdlib.Config{
		APIID:               "27134036",
		APIHash:             "48e6be1f8cd5fc0f831bf657307f41e8",
		//Phone:               "89031062562",
		DatabaseDirectory:   "./tdlib-db",
		//FilesDirectory:      "./tdlib-files",
		UseMessageDatabase:  true,
		UseSecretChats:      false,
		SystemLanguageCode:  "en",
		DeviceModel:         "Desktop",
		SystemVersion:       "Unknown",
		ApplicationVersion:  "1.0",
		EnableStorageOptimizer: true,
		IgnoreFileNames:       false,
	})
	
	fmt.Println("TDLib client created")
	
	currentState, err := client.Authorize()
	if err != nil {
		fmt.Printf("Error getting current state: %v", err)
		return
	}

	fmt.Printf("Current state: %v\n", currentState.GetAuthorizationStateEnum())
	
	for currentState.GetAuthorizationStateEnum() != tdlib.AuthorizationStateReadyType {
		time.Sleep(2 * time.Second)
		currentState, _ = client.Authorize()
		fmt.Printf("Current state: %v\n", currentState.GetAuthorizationStateEnum())
	}
	
	fmt.Println("Authorization ready! The user is logged in")
}
