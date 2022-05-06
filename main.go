package main

import "fmt"

func main() {
	users := []user{
		{
			Name:     "Lexa",
			Age:      28,
			Nickname: "MANEKEN",
			MMR:      5000,
			Cars: []car{
				{
					Model:  "Camry",
					Engine: 2.4,
				},
				{
					Model:  "Gaz",
					Engine: 1.3,
				},
			},
		},
		{
			Name:     "Julus",
			Age:      28,
			Nickname: "BJ",
			MMR:      0,
			Cars: []car{
				{
					Model:  "Camry",
					Engine: 2.4,
				},
			},
		},
		{
			Name:     "Leppieske",
			Age:      31,
			Nickname: "Joe Sky",
			MMR:      2000,
			Cars: []car{
				{
					Model:  "Creta",
					Engine: 2.0,
				},
				{
					Model:  "Premio",
					Engine: 1.8,
				},
			},
		},
		{
			Name:     "Tola",
			Age:      33,
			Nickname: "zuuman",
			MMR:      0,
			Cars: []car{
				{
					Model:  "Noah",
					Engine: 2.4,
				},
				{
					Model:  "Gruzovik",
					Engine: 1.5,
				},
			},
		},
		{
			Name:     "Dima",
			Age:      36,
			Nickname: "-",
			MMR:      0,
			Cars: []car{
				{
					Model:  "Tahoe",
					Engine: 5.3,
				},
				{
					Model:  "Uaz",
					Engine: 2.7,
				},
			},
		},
		{
			Name:     "Dima",
			Age:      36,
			Nickname: "-",
			MMR:      0,
			Cars: []car{
				{
					Model:  "Tahoe",
					Engine: 5.3,
				},
				{
					Model:  "Uaz",
					Engine: 2.7,
				},
				{
					Model:  "Uaz",
					Engine: 2.7,
				},
			},
		},
		{
			Name:     "Dima",
			Age:      36,
			Nickname: "-",
			MMR:      0,
			Cars: []car{
				{
					Model:  "Tahoe",
					Engine: 5.3,
				},
				{
					Model:  "Uaz",
					Engine: 2.7,
				},
				{
					Model:  "Uaz",
					Engine: 2.7,
				},
				{
					Model:  "Uaz",
					Engine: 2.7,
				},
				{
					Model:  "Uaz",
					Engine: 2.7,
				},
			},
		},
	}
	fmt.Println("распечатываем всех пользователей")
	for i, userData := range users {
		fmt.Println(i, userData)
	}
	fmt.Println("распечатываем пользователя с ником zuuman")
	for i, userData := range users { // вывести пользователя с ником zuuman
		if userData.Nickname == "zuuman" {
			fmt.Println(i, userData)
		}
	}
	fmt.Println("распечатываем пользователя старше 30 лет")
	for i, userData := range users {
		if userData.Age > 30 {
			fmt.Println(i, userData)
		}
	}
	fmt.Println("распечатываем пользователя у кого объем больше 1.5")
	for i, userData := range users {
		for _, carData := range userData.Cars {
			if carData.Engine > 1.5 {
				fmt.Println(i, userData)
				break
			}
		}
	}
	fmt.Println("распечатываем у кого камрюха броо")
	for i, userData := range users {
		for _, car := range userData.Cars {
			if car.Model == "Camry" {
				fmt.Println(i, userData)
			}
		}

	}
	fmt.Println("распечатываем у кого больше 3х тачек")
	for i, userData := range users {
		if len(userData.Cars) > 2 {
			fmt.Println(i, userData)
		}

	}
	fmt.Println("распечатываем у кого 1 тачка")
	for i, userData := range users {
		if len(userData.Cars) == 1 {
			fmt.Println(i, userData)
		}

	}
	fmt.Println(users[6])
}

type user struct {
	Name     string
	Age      int
	Nickname string
	MMR      int
	Cars     []car
}
type car struct {
	Model  string
	Engine float64
}
