package get

import (
	"github.com/spf13/cobra"

	configv1 "github.com/openshift/api/config/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/kubectl/pkg/cmd/get"
	kcmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/scheme"
)

type GetOptions struct {
	*get.GetOptions
}

func NewGetOptions(streams genericclioptions.IOStreams) *GetOptions {
	return &GetOptions{
		GetOptions: get.NewGetOptions("oc", streams),
	}
}

// NewCmdGet creates a new Get command that supports OpenShift resources.
func NewCmdGet(f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	o := NewGetOptions(streams)
	cmd := &cobra.Command{
		Use: "get",
		Run: func(cmd *cobra.Command, args []string) {
			kcmdutil.CheckErr(o.Complete(f, cmd, args))
			kcmdutil.CheckErr(o.Validate(cmd))
			kcmdutil.CheckErr(o.Run(f, cmd, args))
		},
	}

	o.GetOptions.AddFlags(cmd)

	return cmd
}

func (o *GetOptions) Complete(f kcmdutil.Factory, cmd *cobra.Command, args []string) error {
	if err := o.GetOptions.Complete(f, cmd, args); err != nil {
		return err
	}
	// 1. force sorting to get full object
	o.SortBy = ".metadata.name"
	originalPrinterFunc := o.GetOptions.ToPrinter
	o.GetOptions.ToPrinter = func(mapping *meta.RESTMapping, outputObjects *bool, withNamespace bool, withKind bool) (printers.ResourcePrinterFunc, error) {
		// below is copied from ToPrinter implementation in vendor/k8s.io/kubectl/pkg/cmd/get/get.go
		if mapping.GroupVersionKind != configv1.GroupVersion.WithKind("ClusterOperator") {
			return originalPrinterFunc(mapping, outputObjects, withNamespace, withKind)
		}

		printFlags := o.PrintFlags.Copy()

		if mapping != nil {
			if kcmdutil.GetFlagString(cmd, "output") == "" && o.PrintWithOpenAPICols {
				if apiSchema, err := f.OpenAPISchema(); err == nil {
					printFlags.UseOpenAPIColumns(apiSchema, mapping)
				}
			}
			printFlags.SetKind(mapping.GroupVersionKind.GroupKind())
		}
		if withNamespace {
			printFlags.EnsureWithNamespace()
		}
		if withKind {
			printFlags.EnsureWithKind()
		}

		printer, err := printFlags.ToPrinter()
		if err != nil {
			return nil, err
		}
		printer, err = printers.NewTypeSetter(scheme.Scheme).WrapToPrinter(printer, nil)
		if err != nil {
			return nil, err
		}

		if len(o.SortBy) > 0 {
			printer = &get.SortingPrinter{Delegate: printer, SortField: o.SortBy}
		}
		if outputObjects != nil {
			printer = &skipPrinter{delegate: printer, output: outputObjects}
		}
		printer = &ClusterOperatorPrinter{Delegate: printer}
		return printer.PrintObj, nil
	}
	return nil
}

func (o *GetOptions) Validate(cmd *cobra.Command) error {
	return o.GetOptions.Validate(cmd)
}

func (o *GetOptions) Run(f kcmdutil.Factory, cmd *cobra.Command, args []string) error {
	return o.GetOptions.Run(f, args)
}
