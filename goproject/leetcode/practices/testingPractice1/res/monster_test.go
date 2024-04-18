package res

import "testing"

func TestStore(t *testing.T) {
	monster := Monster{"sunwukong", 10000}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() run fail, real %v, expect %v", res, true)
	}
	t.Logf("monster.Store() run succ")
}

func TestRestore(t *testing.T) {
	monster := Monster{"sunwukong", 10000}
	newMonster, err := monster.Restore()
	if err != nil {
		t.Fatalf("monster.Store() run fail, real %v, expect %v", err, nil)
	}
	t.Logf("monster.Store() run succ, %v", *newMonster)
}
