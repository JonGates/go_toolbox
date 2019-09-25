package go_toolbox

import (
	"fmt"
	"testing"
)

func TestAuthCode(t *testing.T) {
	key := "JonGates#fjdsa5/#$a,j&%^DG1@!@fdsj:<v#"
	str := "I love JonGates"
	encode := AuthCode(str, "ENCODE", key, 0, 8)
	fmt.Println(encode)

	decode := AuthCode(encode, "DECODE", key, 0, 8)
	fmt.Println(decode)
}
