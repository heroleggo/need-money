package exercise

import "fmt"

type Book struct {
	Name   string
	Index  int
	Class  string
	Status bool
}

type Books []Book

// 책의 타입 = 경영-경제/인문-사회/자기계발/여행/외국어/과학/컴퓨터 로 구분
const (
	BusinessEconomy int = iota
	HumanitySociety int = iota
	SelfDevelopment int = iota
	Travel          int = iota
	ForeignLanguage int = iota
	Science         int = iota
	Computer        int = iota
)

// 가지고 있는 기능 : 책 등록하기, 책 조회하기(검색), 책 삭제하기, 책 대여하기, 책 반납하기

func (b *Books) AddBook(book Book) {
	*b = append(*b, book)
}

func (b *Books) SearchBookByName(name string) (Book, bool) {
	for _, book := range *b {
		if book.Name == name {
			return book, true
		}
	}
	return Book{}, false
}

func (b *Books) SearchBookByIndex(index int) (Book, bool) {
	for _, book := range *b {
		if book.Index == index {
			return book, true
		}
	}
	return Book{}, false
}

func (b *Books) DeleteBook(index int) {
	for i, book := range *b {
		if book.Index == index {
			*b = append((*b)[:i], (*b)[i+1:]...)
		}
	}
}

func (b *Books) RentBook(index int) {
	for i, book := range *b {
		if book.Index == index {
			(*b)[i].Status = true
		}
	}
}

func (b *Books) ReturnBook(index int) {
	for i, book := range *b {
		if book.Index == index {
			(*b)[i].Status = false
		}
	}
}

// TODO: add library service
func RunService() {
	defaultBooks := Books{
		Book{
			Name:   "책1",
			Index:  1,
			Class:  "경영-경제",
			Status: false,
		},
		Book{
			Name:   "책2",
			Index:  2,
			Class:  "인문-사회",
			Status: false,
		},
		Book{
			Name:   "책3",
			Index:  3,
			Class:  "자기계발",
			Status: false,
		},
	}

	// read command from user
	var command string
	for command != "exit" {
		fmt.Println("Enter command: ")
		_, err := fmt.Scanln(&command)
		if err != nil {
			return
		}

		switch command {
		case "add":
			var name, class string
			var index int
			fmt.Println("Enter name: ")
			_, err := fmt.Scanln(&name)
			if err != nil {
				return
			}
			fmt.Println("Enter index: ")
			_, err = fmt.Scanln(&index)
			if err != nil {
				return
			}

		}
	}
}
