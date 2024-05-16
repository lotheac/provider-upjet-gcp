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
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *LiteSubscription) ResolveReferences( // ResolveReferences of this LiteSubscription.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "LiteTopic", "LiteTopicList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Topic),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.TopicRef,
			Selector:     mg.Spec.ForProvider.TopicSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Topic")
	}
	mg.Spec.ForProvider.Topic = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "LiteTopic", "LiteTopicList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Topic),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.TopicRef,
			Selector:     mg.Spec.InitProvider.TopicSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Topic")
	}
	mg.Spec.InitProvider.Topic = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TopicRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LiteTopic.
func (mg *LiteTopic) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.ReservationConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta1", "LiteReservation", "LiteReservationList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ReservationConfig.ThroughputReservation),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.ReservationConfig.ThroughputReservationRef,
				Selector:     mg.Spec.ForProvider.ReservationConfig.ThroughputReservationSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ReservationConfig.ThroughputReservation")
		}
		mg.Spec.ForProvider.ReservationConfig.ThroughputReservation = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ReservationConfig.ThroughputReservationRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.ReservationConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta1", "LiteReservation", "LiteReservationList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ReservationConfig.ThroughputReservation),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.ReservationConfig.ThroughputReservationRef,
				Selector:     mg.Spec.InitProvider.ReservationConfig.ThroughputReservationSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ReservationConfig.ThroughputReservation")
		}
		mg.Spec.InitProvider.ReservationConfig.ThroughputReservation = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ReservationConfig.ThroughputReservationRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Subscription.
func (mg *Subscription) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.DeadLetterPolicy != nil {
		{
			m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DeadLetterPolicy.DeadLetterTopic),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.DeadLetterPolicy.DeadLetterTopicRef,
				Selector:     mg.Spec.ForProvider.DeadLetterPolicy.DeadLetterTopicSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DeadLetterPolicy.DeadLetterTopic")
		}
		mg.Spec.ForProvider.DeadLetterPolicy.DeadLetterTopic = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DeadLetterPolicy.DeadLetterTopicRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "Topic", "TopicList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Topic),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.TopicRef,
			Selector:     mg.Spec.ForProvider.TopicSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Topic")
	}
	mg.Spec.ForProvider.Topic = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.DeadLetterPolicy != nil {
		{
			m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DeadLetterPolicy.DeadLetterTopic),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.DeadLetterPolicy.DeadLetterTopicRef,
				Selector:     mg.Spec.InitProvider.DeadLetterPolicy.DeadLetterTopicSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DeadLetterPolicy.DeadLetterTopic")
		}
		mg.Spec.InitProvider.DeadLetterPolicy.DeadLetterTopic = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DeadLetterPolicy.DeadLetterTopicRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "Topic", "TopicList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Topic),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.TopicRef,
			Selector:     mg.Spec.InitProvider.TopicSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Topic")
	}
	mg.Spec.InitProvider.Topic = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TopicRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubscriptionIAMMember.
func (mg *SubscriptionIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "Subscription", "SubscriptionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Subscription),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.SubscriptionRef,
			Selector:     mg.Spec.ForProvider.SubscriptionSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Subscription")
	}
	mg.Spec.ForProvider.Subscription = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubscriptionRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "Subscription", "SubscriptionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Subscription),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.SubscriptionRef,
			Selector:     mg.Spec.InitProvider.SubscriptionSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Subscription")
	}
	mg.Spec.InitProvider.Subscription = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubscriptionRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Topic.
func (mg *Topic) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta2", "CryptoKey", "CryptoKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KMSKeyNameRef,
			Selector:     mg.Spec.ForProvider.KMSKeyNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyName")
	}
	mg.Spec.ForProvider.KMSKeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta2", "CryptoKey", "CryptoKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.KMSKeyNameRef,
			Selector:     mg.Spec.InitProvider.KMSKeyNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyName")
	}
	mg.Spec.InitProvider.KMSKeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TopicIAMMember.
func (mg *TopicIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "Topic", "TopicList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Topic),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.TopicRef,
			Selector:     mg.Spec.ForProvider.TopicSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Topic")
	}
	mg.Spec.ForProvider.Topic = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "Topic", "TopicList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Topic),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.TopicRef,
			Selector:     mg.Spec.InitProvider.TopicSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Topic")
	}
	mg.Spec.InitProvider.Topic = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TopicRef = rsp.ResolvedReference

	return nil
}
