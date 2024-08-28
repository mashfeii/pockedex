package pockecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
			val: []byte("{'count':1054,'next':'https://pokeapi.co/api/v2/location-area/?offset=20&limit=20','previous':null,'results':[{'name':'canalave-city-area','url':'https://pokeapi.co/api/v2/location-area/1/'},{'name':'eterna-city-area','url':'https://pokeapi.co/api/v2/location-area/2/'},{'name':'pastoria-city-area','url':'https://pokeapi.co/api/v2/location-area/3/'},{'name':'sunyshore-city-area','url':'https://pokeapi.co/api/v2/location-area/4/'},{'name':'sinnoh-pokemon-league-area','url':'https://pokeapi.co/api/v2/location-area/5/'},{'name':'oreburgh-mine-1f','url':'https://pokeapi.co/api/v2/location-area/6/'},{'name':'oreburgh-mine-b1f','url':'https://pokeapi.co/api/v2/location-area/7/'},{'name':'valley-windworks-area','url':'https://pokeapi.co/api/v2/location-area/8/'},{'name':'eterna-forest-area','url':'https://pokeapi.co/api/v2/location-area/9/'},{'name':'fuego-ironworks-area','url':'https://pokeapi.co/api/v2/location-area/10/'},{'name':'mt-coronet-1f-route-207','url':'https://pokeapi.co/api/v2/location-area/11/'},{'name':'mt-coronet-2f','url':'https://pokeapi.co/api/v2/location-area/12/'},{'name':'mt-coronet-3f','url':'https://pokeapi.co/api/v2/location-area/13/'},{'name':'mt-coronet-exterior-snowfall','url':'https://pokeapi.co/api/v2/location-area/14/'},{'name':'mt-coronet-exterior-blizzard','url':'https://pokeapi.co/api/v2/location-area/15/'},{'name':'mt-coronet-4f','url':'https://pokeapi.co/api/v2/location-area/16/'},{'name':'mt-coronet-4f-small-room','url':'https://pokeapi.co/api/v2/location-area/17/'},{'name':'mt-coronet-5f','url':'https://pokeapi.co/api/v2/location-area/18/'},{'name':'mt-coronet-6f','url':'https://pokeapi.co/api/v2/location-area/19/'},{'name':'mt-coronet-1f-from-exterior','url':'https://pokeapi.co/api/v2/location-area/20/'}]}"),
		},
		{
			key: "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
			val: []byte("{'count':1054,'next':'https://pokeapi.co/api/v2/location-area/?offset=40&limit=20','previous':'https://pokeapi.co/api/v2/location-area/?offset=0&limit=20','results':[{'name':'mt-coronet-1f-route-216','url':'https://pokeapi.co/api/v2/location-area/21/'},{'name':'mt-coronet-1f-route-211','url':'https://pokeapi.co/api/v2/location-area/22/'},{'name':'mt-coronet-b1f','url':'https://pokeapi.co/api/v2/location-area/23/'},{'name':'great-marsh-area-1','url':'https://pokeapi.co/api/v2/location-area/24/'},{'name':'great-marsh-area-2','url':'https://pokeapi.co/api/v2/location-area/25/'},{'name':'great-marsh-area-3','url':'https://pokeapi.co/api/v2/location-area/26/'},{'name':'great-marsh-area-4','url':'https://pokeapi.co/api/v2/location-area/27/'},{'name':'great-marsh-area-5','url':'https://pokeapi.co/api/v2/location-area/28/'},{'name':'great-marsh-area-6','url':'https://pokeapi.co/api/v2/location-area/29/'},{'name':'solaceon-ruins-2f','url':'https://pokeapi.co/api/v2/location-area/30/'},{'name':'solaceon-ruins-1f','url':'https://pokeapi.co/api/v2/location-area/31/'},{'name':'solaceon-ruins-b1f-a','url':'https://pokeapi.co/api/v2/location-area/32/'},{'name':'solaceon-ruins-b1f-b','url':'https://pokeapi.co/api/v2/location-area/33/'},{'name':'solaceon-ruins-b1f-c','url':'https://pokeapi.co/api/v2/location-area/34/'},{'name':'solaceon-ruins-b2f-a','url':'https://pokeapi.co/api/v2/location-area/35/'},{'name':'solaceon-ruins-b2f-b','url':'https://pokeapi.co/api/v2/location-area/36/'},{'name':'solaceon-ruins-b2f-c','url':'https://pokeapi.co/api/v2/location-area/37/'},{'name':'solaceon-ruins-b3f-a','url':'https://pokeapi.co/api/v2/location-area/38/'},{'name':'solaceon-ruins-b3f-b','url':'https://pokeapi.co/api/v2/location-area/39/'},{'name':'solaceon-ruins-b3f-c','url':'https://pokeapi.co/api/v2/location-area/40/'}]}"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)

			if !ok {
				t.Errorf("expected value, but not found")
				return
			}

			if string(val) != string(c.val) {
				t.Error("different results")
				return
			}
		})
	}

}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 1*time.Millisecond
	cache := NewCache(baseTime)

	cache.Add("https://pokeapi.co/api/v2/location-area/?offset=20&limit=20", []byte("{'count':1054,'next':'https://pokeapi.co/api/v2/location-area/?offset=40&limit=20','previous':'https://pokeapi.co/api/v2/location-area/?offset=0&limit=20','results':[{'name':'mt-coronet-1f-route-216','url':'https://pokeapi.co/api/v2/location-area/21/'},{'name':'mt-coronet-1f-route-211','url':'https://pokeapi.co/api/v2/location-area/22/'},{'name':'mt-coronet-b1f','url':'https://pokeapi.co/api/v2/location-area/23/'},{'name':'great-marsh-area-1','url':'https://pokeapi.co/api/v2/location-area/24/'},{'name':'great-marsh-area-2','url':'https://pokeapi.co/api/v2/location-area/25/'},{'name':'great-marsh-area-3','url':'https://pokeapi.co/api/v2/location-area/26/'},{'name':'great-marsh-area-4','url':'https://pokeapi.co/api/v2/location-area/27/'},{'name':'great-marsh-area-5','url':'https://pokeapi.co/api/v2/location-area/28/'},{'name':'great-marsh-area-6','url':'https://pokeapi.co/api/v2/location-area/29/'},{'name':'solaceon-ruins-2f','url':'https://pokeapi.co/api/v2/location-area/30/'},{'name':'solaceon-ruins-1f','url':'https://pokeapi.co/api/v2/location-area/31/'},{'name':'solaceon-ruins-b1f-a','url':'https://pokeapi.co/api/v2/location-area/32/'},{'name':'solaceon-ruins-b1f-b','url':'https://pokeapi.co/api/v2/location-area/33/'},{'name':'solaceon-ruins-b1f-c','url':'https://pokeapi.co/api/v2/location-area/34/'},{'name':'solaceon-ruins-b2f-a','url':'https://pokeapi.co/api/v2/location-area/35/'},{'name':'solaceon-ruins-b2f-b','url':'https://pokeapi.co/api/v2/location-area/36/'},{'name':'solaceon-ruins-b2f-c','url':'https://pokeapi.co/api/v2/location-area/37/'},{'name':'solaceon-ruins-b3f-a','url':'https://pokeapi.co/api/v2/location-area/38/'},{'name':'solaceon-ruins-b3f-b','url':'https://pokeapi.co/api/v2/location-area/39/'},{'name':'solaceon-ruins-b3f-c','url':'https://pokeapi.co/api/v2/location-area/40/'}]}"))

	_, ok := cache.Get("https://pokeapi.co/api/v2/location-area/?offset=20&limit=20")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://pokeapi.co/api/v2/location-area/?offset=20&limit=20")
	if ok {
		t.Errorf("expected not to find key")
		return
	}
}
