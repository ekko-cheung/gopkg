package crypto

import "testing"

type child struct {
	Value string `crypto:"-"`
}

type config struct {
	Value string
	Child child
}

func TestDecodeConfig(t *testing.T) {
	key := "abcdefgh"
	c := &config{
		Value: "a//xM5D6Q3XwT14eXPsa/A==",
		Child: child{
			Value: "a//xM5D6Q3XwT14eXPsa/A==",
		},
	}
	err := DecodeConfig(c, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", c)
}

func TestEncodeConfig(t *testing.T) {
	key := "abcdefgh"
	c := &config{
		Value: "helloworld",
		Child: child{
			Value: "helloworld",
		},
	}
	err := EncodeConfig(c, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", c)
}
