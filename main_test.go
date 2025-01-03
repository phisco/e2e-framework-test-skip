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
var testEnvNoFF env.Environment

func TestMain(m *testing.M) {
	testenv = env.NewWithConfig(envconf.New().WithFailFast())
	testEnvNoFF = env.NewWithConfig(envconf.New())
	os.Exit(testenv.Run(m))
}

func TestSkipWithFailFast(t *testing.T) {
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

func TestSkipNoFailFast(t *testing.T) {
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

	testEnvNoFF.Test(t, feat1, feat2)
}

func TestFailNowWithFailFast(t *testing.T) {
	feat1 := features.New("fail now").
		Assess("Assess 1", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			t.Log("Assess 1 (should be printed)")
			t.FailNow()
			return ctx
		}).
		Assess("Assess 2", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			t.Log("Assess 2 (should NOT be printed)")
			return ctx
		}).
		Feature()

	feat2 := features.New("succeed").
		Assess("Assess 1", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			t.Log("Assess 1 (should NOT be printed)")
			return ctx
		}).
		Assess("Assess 2", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			t.Log("Assess 2 (should NOT be printed)")
			return ctx
		}).
		Feature()

	testenv.Test(t, feat1, feat2)
}

func TestFailNowNoFailFast(t *testing.T) {
	feat1 := features.New(t.Name()).
		Assess("Assess 1", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			t.Log("Assess 1 (should be printed)")
			t.FailNow()
			t.Log("Assess 1 post (should NOT be printed)")
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

	testEnvNoFF.Test(t, feat1, feat2)
}

func TestFailWithFailFast(t *testing.T) {
	feat := features.New(t.Name()).
		Assess("Assess 1", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			t.Log("Assess 1 (should be printed)")
			t.Fail()
			t.Log("Assess 1 post (should be printed)")
			return ctx
		}).
		Assess("Assess 2", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			t.Log("Assess 2 (should be printed)")
			return ctx
		}).
		Feature()

	testenv.Test(t, feat)
}

func TestFailNoFailFast(t *testing.T) {
	feat := features.New(t.Name()).
		Assess("Assess 1", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			t.Log("Assess 1 (should be printed)")
			t.Fail()
			return ctx
		}).
		Assess("Assess 2", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			t.Log("Assess 2 (should be printed)")
			return ctx
		}).
		Feature()

	testEnvNoFF.Test(t, feat)
}
