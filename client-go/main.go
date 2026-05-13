package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fajarilf/grpc-chat-client/app"
)

func main() {
	model := app.NewModel()
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatalf("there was an error: %v", err)
	}

	// url := flag.String("url", "ws://localhost:3000/ws", "WebSocket server URL")
	// flag.Parse()

	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	// defer cancel()

	// dialCtx, dialCancel := context.WithTimeout(ctx, 10*time.Second)
	// defer dialCancel()

	// conn, _, err := websocket.Dial(dialCtx, *url, nil)
	// if err != nil {
	// 	log.Fatalf("dial %s: %v", *url, err)
	// }
	// defer conn.CloseNow()

	// log.Printf("connected to %s", *url)

	// go readLoop(ctx, conn)

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println("type a message and press Enter (Ctrl+C to quit):")
	// for scanner.Scan() {
	// 	text := strings.TrimRight(scanner.Text(), "\r\n")
	// 	if text == "" {
	// 		continue
	// 	}
	// 	if err := conn.Write(ctx, websocket.MessageText, []byte(text)); err != nil {
	// 		log.Printf("write: %v", err)
	// 		break
	// 	}
	// }
	// if err := scanner.Err(); err != nil {
	// 	log.Printf("stdin: %v", err)
	// }

	// conn.Close(websocket.StatusNormalClosure, "bye")
}

// func readLoop(ctx context.Context, conn *websocket.Conn) {
// 	for {
// 		_, data, err := conn.Read(ctx)
// 		if err != nil {
// 			if ctx.Err() == nil {
// 				log.Printf("read: %v", err)
// 			}
// 			return
// 		}
// 		fmt.Printf("<- %s\n", data)
// 	}
// }
