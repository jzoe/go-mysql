/*
	Copyright (c) 2014, Percona LLC and/or its affiliates. All rights reserved.

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package log

// Options encapsulate common options for making a new log parser.
type Options struct {
	StartOffset        uint64          // byte offset in file at which to start parsing
	FilterAdminCommand map[string]bool // admin commands to ignore
	Debug              bool            // print trace info to STDOUT
}

// A LogParser sends log events to a channel.
type LogParser interface {
	Start() error
	Stop()
	EventChan() <-chan *Event
}
