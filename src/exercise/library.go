package exercise

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
