package conventionalcommit

import (
	"os"
	"testing"

	. "github.com/onsi/gomega"
)

func TestParseCommitMsg(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		cm, err := ParseCommitMsg(`feat(parser-test): add ability to parse arrays`)
		NewWithT(t).Expect(err).To(BeNil())

		NewWithT(t).Expect(cm.Type).To(Equal("feat"))
		NewWithT(t).Expect(cm.Scope).To(Equal("parser-test"))
		NewWithT(t).Expect(cm.Header).To(Equal("add ability to parse arrays"))
	})

	t.Run("BREAKING CHANGE", func(t *testing.T) {
		cm, err := ParseCommitMsg(`feat: allow provided config object to extend other configs

BREAKING CHANGE: extends key in config file is now used for extending other config files`)
		NewWithT(t).Expect(err).To(BeNil())

		NewWithT(t).Expect(cm.Type).To(Equal("feat"))
		NewWithT(t).Expect(cm.BreakingChange).To(BeTrue())
	})

	t.Run("BREAKING CHANGE", func(t *testing.T) {
		cm, err := ParseCommitMsg(`feat!: allow provided config object to extend other configs`)
		NewWithT(t).Expect(err).To(BeNil())

		NewWithT(t).Expect(cm.Type).To(Equal("feat"))
		NewWithT(t).Expect(cm.BreakingChange).To(BeTrue())
	})
}

func TestParseNewMessage(t *testing.T) {
	msg, _ := os.ReadFile("./testdata/COMMIT_EDITMSG")
	cm, err := ParseCommitMsg(string(msg))
	NewWithT(t).Expect(err).To(BeNil())
	NewWithT(t).Expect(cm.Type).To(Equal("ci"))
	NewWithT(t).Expect(cm.Scope).To(BeEmpty())
	NewWithT(t).Expect(cm.Header).To(Equal("update config"))
	NewWithT(t).Expect(cm.Body).To(Equal("add mysql redis config\n"))
	t.Log(cm.Type)
	t.Log(cm.Scope)
	t.Log(cm.Header)
	t.Log(cm.Body)
}
