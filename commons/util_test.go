package commons

import (
	"errors"
	"log"
	"reflect"
	"sort"
	"testing"
	"time"
	"net"
	"regexp"
)

func TestKeySet(t *testing.T) {
	mp := map[string]interface{}{
		"one": struct{}{},
		"two": struct{}{},
	}

	actual := KeySet(mp)

	//make sure they are sorted before validation
	sort.Strings(actual)

	expected := []string{"one", "two"}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Arrays are not equal. Expected:%s, Actual:%s", expected, actual)
	}
}

func TestRetryAttempts(t *testing.T) {
	i := 0
	Retry(2, 1*time.Second, func() error {
		i++
		return errors.New("some error")
	})

	log.Println(i)
	if 2 != i {
		t.Errorf("Incorrect attempts count: %d", i)
	}
}

func TestGetLocalIP(t *testing.T) {

	ip := net.ParseIP(GetLocalIP())
	if ip.IsLoopback() {
		t.Errorf("IP is loopback: %s", ip.String())
	}

	if !regexp.MustCompile("\\d+\\.\\d+\\.\\d+\\.\\d+").MatchString(ip.String()) {
		t.Errorf("Incorrect IP format: %s", ip.String())
	}

	print(ip.String())
	print()
}