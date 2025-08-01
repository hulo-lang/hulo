// Copyright 2025 The Hulo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package tools

import (
	"os"
	"runtime"
	"time"

	"github.com/caarlos0/log"
	"github.com/magefile/mage/mg"
	"github.com/opencommand/tinge"
)

// Build compiles and packages the project using GoReleaser.
func Build() {
	log.Info(tinge.Styled().Bold("starting build...").String())
	start := time.Now()

	mg.Deps(generateHuloParser)

	elapsed := time.Since(start)
	log.Infof("build completed in %.2fs", elapsed.Seconds())

	err := os.Setenv("GOVERSION", runtime.Version())
	if err != nil {
		log.WithError(err).Fatal("failed to set GOVERSION environment variable")
	}

	// TODO build:all build:single-platform
	err = runCmd("goreleaser", "build", "--snapshot", "--clean")
	if err != nil {
		log.WithError(err).Fatal("goreleaser build failed")
	}

	// copy standard library to the output directory
}
