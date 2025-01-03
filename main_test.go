package main

import (
	"context"
	"os"
	"testing"

	"sigs.k8s.io/e2e-framework/pkg/env"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/features"
)

var testenv env.Environment

func TestMain(m *testing.M) {
	testenv = env.NewWithConfig(envconf.New().WithFailFast())
	os.Exit(testenv.Run(m))
}

func TestSkip(t *testing.T) {
	feat1 := features.New("skip").
		Assess("Assess 1", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			t.Log("Assess 1 (should be printed)")
			t.Skipf("skipping Assess 1")
			return ctx
		}).
		Assess("Assess 2", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			t.Log("Assess 2 (should be printed)")
			return ctx
		}).
		Feature()

	feat2 := features.New("succeed").
		Assess("Assess 1", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			t.Log("Assess 1 (should be printed)")
			return ctx
		}).
		Assess("Assess 2", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			t.Log("Assess 2 (should be printed)")
			return ctx
		}).
		Feature()

	testenv.Test(t, feat1, feat2)
}
