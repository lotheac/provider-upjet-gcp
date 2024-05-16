// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Instance.
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
)

func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("servicenetworking.gcp.upbound.io", "v1beta1", "Connection", "ConnectionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AuthorizedNetwork),
			Extract:      resource.ExtractParamPath("network", false),
			Reference:    mg.Spec.ForProvider.AuthorizedNetworkRef,
			Selector:     mg.Spec.ForProvider.AuthorizedNetworkSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AuthorizedNetwork")
	}
	mg.Spec.ForProvider.AuthorizedNetwork = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AuthorizedNetworkRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("servicenetworking.gcp.upbound.io", "v1beta1", "Connection", "ConnectionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AuthorizedNetwork),
			Extract:      resource.ExtractParamPath("network", false),
			Reference:    mg.Spec.InitProvider.AuthorizedNetworkRef,
			Selector:     mg.Spec.InitProvider.AuthorizedNetworkSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AuthorizedNetwork")
	}
	mg.Spec.InitProvider.AuthorizedNetwork = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AuthorizedNetworkRef = rsp.ResolvedReference

	return nil
}
