package explode

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "launchpad.net/gocheck"
)

type clientSuite struct {
	configPath string
	j          int
}

func (cs *clientSuite) writeTestConfig() {
	cs.j++
	x := strings.Repeat("a", 15+cs.j%32)
	cfgMap := map[string]interface{}{
		"foo":                    x,
		"connect_timeout":        "7ms",
		"exchange_timeout":       "10ms",
		"hosts_cache_expiry":     "1h",
		"expect_all_repaired":    "30m",
		"stabilizing_timeout":    "0ms",
		"connectivity_check_url": "",
		"connectivity_check_md5": "",
		"addr":             ":0",
		"recheck_timeout":  "3h",
		"auth_helper":      "",
		"session_url":      "xyzzy://",
		"registration_url": "reg://",
		"log_level":        "debug",
		"poll_interval":    "5m",
		"poll_settle":      "20ms",
		"poll_net_wait":    "1m",
		"poll_polld_wait":  "3m",
		"poll_done_wait":   "5s",
	}
	_, err := json.Marshal(cfgMap)
	if err != nil {
		panic(err)
	}
}

func (cs *clientSuite) witness() {
	dir := "xxx"
	cs.configPath = filepath.Join(dir, "config")
	cs.writeTestConfig()
}

type Cli2 struct {
	helper string
	ccu    CClickUser
}

func (cli *Cli2) configure() {
	err := cli.ccu.CInit(cli)
	if err != nil {
		panic(err)
	}
}

func (cli *Cli2) getAuth(url string) string {
	auth := "hello " + url
	//err := exec.Command(cli.helper, url).Run()
	p, err := os.StartProcess(cli.helper, []string{cli.helper, url}, new(os.ProcAttr))
	if err == nil {
		_, err = p.Wait()
	}
	if err != nil {
		return ""
	} else {
		return strings.TrimSpace(string(auth))
	}
}

func (cs *clientSuite) poke() {
	cli2 := &Cli2{helper: "dummyauth.sh"}
	cli2.configure()
	cli2.getAuth("xyzzy://")
}

func TestClient(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&clientSuite{})

func (cs *clientSuite) SetUpSuite(c *C) {
	fmt.Println("RUN")
	cs.witness()
	for j := 0; j < 10000; j++ {
		cs.poke()
		for i := 0; i < 61; i++ {
			cs.witness()
		}
	}
}

func (cs *clientSuite) TestNothing(c *C) {
}
