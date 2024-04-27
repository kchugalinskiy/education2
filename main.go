package main

import (
	"github.com/kchugalinskiy/education2/chans"
)

// const char *hello = "Hello";
// const char hello[6] = { 'H','e','l', 'l', 'o', '\0' };
// const char world[6] = { 'w','o','r', 'l', 'd', '\0' };
// const char helloworld[11] = ???
// memcpy(helloworld, hello, 5);
// memcpy(helloworld + 5, world, 5);
// helloworld[10] = '\0';
/*
class String {
public:
	String(const char *str) {
		len = strlen(str) + 1;
		m_str = new char[strlen(str) + 1];
		strcpy(m_str, str);
	}
	~String() {
		delete[] m_str;
	}

	[a,b, c, d, e, f, g, h, i, j, 0, 0,0, 0,0,0]
std::vector<int> v = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};


	String &operator+(const String &str) {
		char *tmp = new char[len + str.len - 1];
		strcpy(tmp, m_str);
		strcat(tmp+len, str.m_str);
		delete[] m_str;
		m_str = tmp;
		len = len + str.len - 1;
		return *this;
	}
private:
	char *m_str;
	size_t len;
	size_t cap;
};
*/

//go:generate mockery --name=Vehicle --exported

func main() {
	//arr := []int{1, 2, 3, 4, 5}
	//arr1 := arr[1:3]
	//arr2 := [5]int{1, 2, 3, 4, 5}
	//fmt.Println(arr)
	//fmt.Println(arr1)
	//fmt.Println(arr2)
	//g(&arr)
	//fmt.Println(arr)
	//fmt.Println(arr1)
	//fmt.Println("Hello, World!")
	//var ptrs []*int
	//for i, v := range []int{1, 2, 3, 4, 5} {
	//	println(i, v)
	//	tmp := v
	//	ptrs = append(ptrs, &tmp)
	//}
	//
	//for _, ptr := range ptrs {
	//	println(*ptr)
	//}
	//concurrency.Conc()
	//if err := printFile(); err != nil {
	//	fmt.Printf("%v\n", err)
	//}
	chans.Somefun()
}

//func printFile() error {
//	f, err := os.Open("main.g")
//	if err != nil {
//		return fmt.Errorf("print file: %w", err)
//	}
//	defer f.Close()
//	return nil
//}
