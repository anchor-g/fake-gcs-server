// Copyright 2019 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package backend

import "time"

// BucketCORS represents Cross-Origin Resource Sharing configuration for a bucket.
type BucketCORS struct {
	// MaxAge is the value to return in the Access-Control-Max-Age
	// header used in preflight responses, in seconds.
	MaxAge int64

	// Methods is the list of HTTP methods on which to include CORS response
	// headers (GET, OPTIONS, POST, etc). "*" is permitted and means "any method".
	Methods []string

	// Origins is the list of Origins eligible to receive CORS response
	// headers. "*" is permitted and means "any Origin".
	Origins []string

	// ResponseHeaders is the list of HTTP headers other than the simple
	// response headers to give permission for the user-agent to share
	// across domains.
	ResponseHeaders []string
}

// Bucket represents the bucket that is stored within the fake server.
type Bucket struct {
	Name                  string
	VersioningEnabled     bool
	TimeCreated           time.Time
	DefaultEventBasedHold bool
	CORS                  []BucketCORS
}

const bucketMetadataSuffix = ".bucketMetadata"

type BucketAttrs struct {
	DefaultEventBasedHold bool
	VersioningEnabled     bool
	CORS                  []BucketCORS
}
