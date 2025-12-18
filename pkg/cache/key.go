package cache

import "fmt"

func UserKey(id uint) string {
	return fmt.Sprintf("user:%d", id)
}
