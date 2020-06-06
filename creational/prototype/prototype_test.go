package prototype

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Received cache was nil")
	}

	itme1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}
	if itme1 == whitePrototype {
		t.Error("item1 cannot be equal to the white prototype")
	}
	shirt1, ok := itme1.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}
	shirt1.SKU = "abbcc"

	itme2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}

	shirt2, ok := itme2.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}
	if shirt1.SKU == shirt2.SKU {
		t.Error("SKU's of shirt1 and shirt2 must be different")
	}

	if shirt1 == shirt2 {
		t.Error("Shirt1 cannot be equal to Shirt 2")
	}
	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())
}
