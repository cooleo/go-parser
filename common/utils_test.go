package common

import (
  "testing"
  "fmt"
)

func TestGenerateTextSlug(t *testing.T) {
    result := GenerateTextSlug("Nguyen Dang Hung")
    if result == "nguyen-dang-hung"  {
        t.Log("TestGenerateTextSlug passed!")
    } else {
        t.Error("TestGenerateTextSlug has error")
    }
    fmt.Println("resutl:", result)
}

func TestGenerateTextUnicodeSlug(t *testing.T) {
    result := GenerateTextSlug("Nguyễn Đăng Hùng")
    if result == "nguyễn-đăng-hùng"  {
        t.Log("TestGenerateTextSlug passed!")
    } else {
        t.Error("TestGenerateTextSlug has error")
    }
    fmt.Println("resutl:", result)
}