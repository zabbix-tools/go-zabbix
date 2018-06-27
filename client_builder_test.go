package zabbix

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const (
	fakeURL        = "http://localhost/api_jsonrpc.php"
	fakeToken      = "0424bd59b807674191e7d77572075f33"
	fakeAPIVersion = "2.0"
)

func prepareTemporaryDir(t *testing.T) (dir string, success bool) {
	tempDir, err := ioutil.TempDir("", "zabbix-session-test")

	if err != nil {
		t.Fatalf("cannot create a temporary dir for session cache: %v", err)
		return "", false
	}

	t.Logf("used %s directory as temporary dir", tempDir)

	return tempDir, true
}

func getTestFileCache(baseDir string) SessionAbstractCache {
	sessionFilePath := baseDir + "/" + ".zabbix_session"
	return NewSessionFileCache().SetFilePath(sessionFilePath)
}

func TestSessionCache(t *testing.T) {
	// Create a fake session for r/w test
	fakeSession := &Session{
		URL:        fakeURL,
		Token:      fakeToken,
		APIVersion: fakeAPIVersion,
	}

	tempDir, success := prepareTemporaryDir(t)

	if !success {
		return
	}

	cache := getTestFileCache(tempDir)

	if err := cache.SaveSession(fakeSession); err != nil {
		t.Errorf("failed to save mock session - %v", err)
		return
	}

	if !cache.HasSession() {
		t.Errorf("session was saved but not detected again by cache")
		return
	}

	// Try to get a cached session
	cachedSession, err := cache.GetSession()

	if err != nil {
		t.Error(err)
		return
	}

	// Check session integrity
	if err := compareSessionWithMock(cachedSession); err != nil {
		t.Error(err)
	}

	testClientBuilder(t, cache)

	if err := cache.Flush(); err != nil {
		t.Error("failed to remove a cached session file")
	}
}

func compareSessionWithMock(session *Session) error {
	if session.URL != fakeURL {
		return fmt.Errorf("Session URL '%s' is not equal to '%s'", session.URL, fakeURL)
	}

	if session.Token != fakeToken {
		return fmt.Errorf("Session token '%s' is not equal to '%s'", session.Token, fakeToken)
	}

	if session.APIVersion != fakeAPIVersion {
		return fmt.Errorf("Session version '%s' is not equal to '%s'", session.APIVersion, fakeAPIVersion)
	}

	return nil
}

// should started by TestSessionCache
func testClientBuilder(t *testing.T, cache SessionAbstractCache) {
	username, password, url := GetTestCredentials()

	if !cache.HasSession() {
		t.Errorf("ManualTestClientBuilder test requires a cached session, run TestSessionCache before running this test case")
		return
	}

	// Try to build a session using the session builder
	client, err := CreateClient(url).WithCache(cache).WithCredentials(username, password).Connect()

	if err != nil {
		t.Errorf("failed to create a session using cache - %s", err)
		return
	}

	if err := compareSessionWithMock(client); err != nil {
		t.Error(err)
	}
}
