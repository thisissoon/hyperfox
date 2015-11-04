// Copyright (c) 2012-2014 José Carlos Nieto, https://menteslibres.net/xiam
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"

	"github.com/thisissoon/hyperfox/api"
	"github.com/thisissoon/hyperfox/deadpool"
	"github.com/xiam/hyperfox/proxy"
	"github.com/xiam/hyperfox/tools/capture"
	"github.com/xiam/hyperfox/tools/logger"
)

const version = "0.9"

const (
	defaultAddress           = `0.0.0.0`
	defaultPort              = uint(1080)
	defaultSSLPort           = uint(10443)
	defaultCaptureCollection = `capture`
)

var (
	flagDatabase    = flag.String("b", "", "Path to database.")
	flagAddress     = flag.String("l", defaultAddress, "Bind address.")
	flagPort        = flag.Uint("p", defaultPort, "Port to bind to.")
	flagSSLPort     = flag.Uint("s", defaultSSLPort, "Port to bind to (SSL mode).")
	flagSSLCertFile = flag.String("c", "", "Path to root CA certificate.")
	flagSSLKeyFile  = flag.String("k", "", "Path to root CA key.")
	flagDeadpoolUrl = flag.String("d", "", "Path to root CA key.")
)

var (
	enableDatabaseSave = false
)

// Parses flags and initializes Hyperfox tool.
func main() {
	var err error
	var sslEnabled bool

	// Banner.
	log.Info("Hyperfox v%s // https://hyperfox.org", version)
	log.Info("By José Carlos Nieto.")

	// Parsing command line flags.
	flag.Parse()

	// Is SSL enabled?
	if *flagSSLPort > 0 && *flagSSLCertFile != "" {
		sslEnabled = true
	}

	// User requested SSL mode.
	if sslEnabled {
		if *flagSSLCertFile == "" {
			flag.Usage()
			log.Fatal(ErrMissingSSLCert)
		}

		if *flagSSLKeyFile == "" {
			flag.Usage()
			log.Fatal(ErrMissingSSLKey)
		}

		os.Setenv(proxy.EnvSSLCert, *flagSSLCertFile)
		os.Setenv(proxy.EnvSSLKey, *flagSSLKeyFile)
	}

	// Creatig proxy.
	p := proxy.NewProxy()

	// Attaching logger.
	p.AddLogger(logger.Stdout{})

	// Attaching capture tool.
	res := make(chan capture.Response, 256)

	p.AddBodyWriteCloser(capture.New(res))

	// Saving captured data with a goroutine.
	go func() {
		for {
			select {
			case r := <-res:
				res, err := api.SendCapturedObject(deadpool.DeadpoolApiAdapter{}, &r)
				if err != nil {
					log.Error(err)
				} else {
					log.Info(res)
				}
			}
		}
	}()

	// if err = startServices(); err != nil {
	// 	log.Fatal("startServices:", err)
	// }

	fmt.Println("")

	cerr := make(chan error)

	// Starting proxy servers.

	go func() {
		if err := p.Start(fmt.Sprintf("%s:%d", *flagAddress, *flagPort)); err != nil {
			cerr <- err
		}
	}()

	if sslEnabled {
		go func() {
			if err := p.StartTLS(fmt.Sprintf("%s:%d", *flagAddress, *flagSSLPort)); err != nil {
				cerr <- err
			}
		}()
	}

	err = <-cerr

	log.Fatalf(ErrBindFailed.Error(), err)
}
