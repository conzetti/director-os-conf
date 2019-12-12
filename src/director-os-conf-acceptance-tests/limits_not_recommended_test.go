package os_conf_acceptance_tests_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("LimitsNotRecommended", func() {
	It("allows users to set security configuration", func() {
		session := boshSSH("director-os-conf/0", "sudo cat /etc/security/limits.d/61-bosh-director-os-conf.conf")
		Eventually(session, 30*time.Second).Should(gbytes.Say("\\* soft nofile 60000"))
		Eventually(session, 30*time.Second).Should(gbytes.Say("\\* hard nofile 100000"))
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))

		session = boshSSH("director-os-conf/0", "ulimit -n")
		Eventually(session, 30*time.Second).Should(gbytes.Say("60000"))
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))

		session = boshSSH("director-os-conf/0", "ulimit -n 100000")
		Eventually(session, 30*time.Second).Should(gexec.Exit(0))

		session = boshSSH("director-os-conf/0", "ulimit -n 100001")
		Eventually(session, 30*time.Second).Should(gexec.Exit(1))
	})
})
