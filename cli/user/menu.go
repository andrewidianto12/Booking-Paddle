package user

import (
	"bufio"
	"fmt"

	"github.com/andrewidianto/Paddle-Booking/entity"
)

func Menu(reader *bufio.Reader, user *entity.LoginUser) {
	for {
		fmt.Println("\n=== USER MENU ===")
		//pilih fitur
		fmt.Println("3. Logout")
		fmt.Print("Pilih: ")

		reader.ReadString('\n')
		return
	}
}
