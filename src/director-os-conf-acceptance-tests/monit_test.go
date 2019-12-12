package os_conf_acceptance_tests_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Monit", func() {
	It("reloads monit at least once after start", func() {
		Eventually(func() string {
			session := boshSSH("director-os-conf/0", "sudo cat /var/vcap/monit/monit.log")
			Eventually(session, 30*time.Second).Should(gexec.Exit(0))
			return string(session.Out.Contents())
		}, 60*time.Second).Should(ContainSubstring("Reinitializing monit daemon"))
	})
})
