package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"net"
	"os"
	"time"
	"wails_rover/rover"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed frontend/dist
var assets embed.FS

func handleConnection(conn net.Conn, app *application.App) {
	defer conn.Close()

	// Create a reader to read data from the connection.
	reader := bufio.NewReader(conn)
	for {
		// Read data from the connection.
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		// Print the received message.
		fmt.Print("Message received: ", message)

		app.Events.Emit(&application.WailsEvent{
			Name: "message",
			Data: message,
		})

		// Send a response back to the client.
		// _, err = conn.Write([]byte("Message received\n"))
		// if err != nil {
		// 	fmt.Println("Error writing:", err.Error())
		// 	return
		// }
	}
}

func startTcpServer(app *application.App) {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error starting server:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")

	for {
		// Accept an incoming connection.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}

		// Handle the connection in a new goroutine.
		go handleConnection(conn, app)
	}
}

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {
	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "wails-proba",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(&GreetService{}),
			application.NewService(&Proba{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	r := rover.Rover{}

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:     "Window 1",
		Frameless: true,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
	})

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.Events.Emit(&application.WailsEvent{
				Name: "time",
				Data: now,
			})
			time.Sleep(time.Second)
		}
	}()

	go func() {
		startTcpServer(app)
	}()

	// Run the application. This blocks until the application has been exited.
	err := app.Run()
	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
