// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/auth/credentials"
	"cloud.google.com/go/auth/oauth2adapt"
	control "cloud.google.com/go/storage/control/apiv2"
	"cloud.google.com/go/storage/control/apiv2/controlpb"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	bucket := "kislayk_bkt"

	scope := "https://www.googleapis.com/auth/devstorage.full_control"

	opts := &credentials.DetectOptions{
		Scopes: []string{scope},
	}

	// Detect credentials using the specified options
	creds, err := credentials.DetectDefault(opts)
	if err != nil {
		log.Fatalf("failed to detect credentials: %v", err)
	}

	// Request token source for required scopes
	ts := oauth2adapt.TokenSourceFromTokenProvider(creds.TokenProvider)

	// Create client options
	//
	//clientOpts := []option.ClientOption{option.WithTokenSource(ts), option.WithEndpoint("storage.googleapis.com:443")} // Succeeds
	clientOpts := []option.ClientOption{option.WithTokenSource(ts), option.WithEndpoint("storage.us-central1.rep.googleapis.com:443")} // Hangs

	controlClient, err := control.NewStorageControlClient(ctx, clientOpts...)
	if err != nil {
		log.Fatalf("failed to create control client: %v", err)
	}

	// Make the GetStorageLayout API call
	req := &controlpb.GetStorageLayoutRequest{
		// Define your request parameters here.  For example:
		Name: fmt.Sprintf("projects/_/buckets/%s/storageLayout", bucket), //  Replace with your actual resource name
	}

	layout, err := controlClient.GetStorageLayout(ctx, req)
	if err != nil {
		log.Fatalf("failed to get storage layout: %v", err)
	}

	fmt.Printf("Storage Layout: %v\n", layout)
}
