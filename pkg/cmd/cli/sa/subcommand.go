package sa

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/openshift/origin/pkg/cmd/templates"
	cmdutil "github.com/openshift/origin/pkg/cmd/util"
	"github.com/openshift/origin/pkg/cmd/util/clientcmd"
)

const ServiceAccountsRecommendedName = "serviceaccounts"

var serviceAccountsLong = templates.LongDesc(`Manage service accounts in your project.

Service accounts allow system components to access the API.`)

const (
	serviceAccountsShort = `Manage service accounts in your project.`
)

func NewCmdServiceAccounts(name, fullName string, f *clientcmd.Factory, out io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:     name,
		Short:   serviceAccountsShort,
		Long:    serviceAccountsLong,
		Aliases: []string{"sa"},
		Run:     cmdutil.DefaultSubCommandRun(out),
	}

	cmds.AddCommand(NewCommandGetServiceAccountToken(GetServiceAccountTokenRecommendedName, fullName+" "+GetServiceAccountTokenRecommendedName, f, out))
	cmds.AddCommand(NewCommandNewServiceAccountToken(NewServiceAccountTokenRecommendedName, fullName+" "+NewServiceAccountTokenRecommendedName, f, out))

	return cmds
}
