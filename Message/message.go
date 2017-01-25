// Copyright 2017 The Yiğit YILDIRIM<yigit@yildirim.me> Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matchmaking

// Request data struct
type Request struct {
	MessageType int
	UserID      string // User ID
	Function    string // Function
	Command     string // Command
	Data        []byte // Command Data JSON
}

// Response data struct
type Response struct {
	MessageType int
	ServerID    string // Server ID
	Function    string // Function
	Command     string // Command
	Data        []byte // Command Data JSON
}
