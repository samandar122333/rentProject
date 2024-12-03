package main

import (
	"errors"
	"fmt"
)

type Vehicle interface {
	Rent() error
	Return() error
	Details()
	IsRented() bool
}

type Car struct {
	ID     int
	Brand  string
	Rented bool
}

type Bike struct {
	ID     int
	Brand  string
	Rented bool
}

type Skate struct {
	ID     int
	Brand  string
	Rented bool
}

func (c *Car) Rent() error {
	if c.Rented {
		return errors.New("автомобиль уже арендован")
	}
	c.Rented = true
	return nil
}

func (c *Car) Return() error {
	if !c.Rented {
		return errors.New("автомобиль ещё не арендован")
	}
	c.Rented = false
	return nil
}

func (c *Car) Details() {
	fmt.Printf("Автомобиль ID: %d, Бренд: %s, Арендован: %t\n", c.ID, c.Brand, c.Rented)
}

func (c *Car) IsRented() bool {
	return c.Rented
}

func (b *Bike) Rent() error {
	if b.Rented {
		return errors.New("велосипед уже арендован")
	}
	b.Rented = true
	return nil
}

func (b *Bike) Return() error {
	if !b.Rented {
		return errors.New("велосипед ещё не арендован")
	}
	b.Rented = false
	return nil
}

func (b *Bike) Details() {
	fmt.Printf("Велосипед ID: %d, Бренд: %s, Арендован: %t\n", b.ID, b.Brand, b.Rented)
}

func (b *Bike) IsRented() bool {
	return b.Rented
}

func (s *Skate) Rent() error {
	if s.Rented {
		return errors.New("скейтборд уже арендован")
	}
	s.Rented = true
	return nil
}

func (s *Skate) Return() error {
	if !s.Rented {
		return errors.New("скейтборд ещё не арендован")
	}
	s.Rented = false
	return nil
}

func (s *Skate) Details() {
	fmt.Printf("Скейтборд ID: %d, Бренд: %s, Арендован: %t\n", s.ID, s.Brand, s.Rented)
}

func (s *Skate) IsRented() bool {
	return s.Rented
}

func main() {
	var vehicles []Vehicle
	var lastID int = 1

	for {
		var vehicleType string
		fmt.Print("Введите тип транспорта (car/bike/skate или 'exit' для выхода): ")
		fmt.Scanln(&vehicleType)

		if vehicleType == "exit" {
			break
		}

		switch vehicleType {
		case "car":
			var brand string
			fmt.Print("Введите бренд автомобиля: ")
			fmt.Scanln(&brand)
			vehicles = append(vehicles, &Car{ID: lastID, Brand: brand})
			lastID++
		case "bike":
			var brand string
			fmt.Print("Введите бренд велосипеда: ")
			fmt.Scanln(&brand)
			vehicles = append(vehicles, &Bike{ID: lastID, Brand: brand})
			lastID++
		case "skate":
			var brand string
			fmt.Print("Введите бренд скейтборда: ")
			fmt.Scanln(&brand)
			vehicles = append(vehicles, &Skate{ID: lastID, Brand: brand})
			lastID++
		default:
			fmt.Println("Неизвестный тип транспорта.")
		}
	}

	for {
		var action string
		fmt.Print("Введите действие (rent/return/details/list/exit): ")
		fmt.Scanln(&action)

		if action == "exit" {
			break
		}

		if action == "list" {
			listAvailableVehicles(vehicles)
			continue
		}

		var id int
		fmt.Print("Введите ID транспортного средства: ")
		fmt.Scanln(&id)

		var found bool
		for _, v := range vehicles {
			switch v := v.(type) {
			case *Car:
				if v.ID == id {
					found = true
					handleAction(v, action)
				}
			case *Bike:
				if v.ID == id {
					found = true
					handleAction(v, action)
				}
			case *Skate:
				if v.ID == id {
					found = true
					handleAction(v, action)
				}
			}
		}

		if !found {
			fmt.Println("Транспорт не найден.")
		}
	}
	fmt.Println("Система аренды остановлена.")
}

func handleAction(v Vehicle, action string) {
	switch action {
	case "rent":
		if err := v.Rent(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Транспорт арендован.")
		}
	case "return":
		if err := v.Return(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Транспорт возвращен.")
		}
	case "details":
		v.Details()
	default:
		fmt.Println("Неизвестное действие.")
	}
}

func listAvailableVehicles(vehicles []Vehicle) {
	fmt.Println("Доступные транспортные средства:")
	for _, v := range vehicles {
		if !v.IsRented() {
			v.Details()
		}
	}
}
