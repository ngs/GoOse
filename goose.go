/*
This is a golang port of "Goose" originaly licensed to Gravity.com
under one or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.

Golang port was written by Antonio Linari

Gravity.com licenses this file
to you under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package goose

type Goose struct {
	config configuration
}

func New(args ...string) Goose {

	return Goose{
		config: GetDefualtConfiguration(args...),
	}
}

func (this Goose) ExtractFromUrl(url string) *Article {
	cc := NewCrawler(this.config, url, "")
	return cc.Crawl()
}

func (this Goose) ExtractFromRawHtml(url string, rawHtml string) *Article {
	cc := NewCrawler(this.config, url, rawHtml)
	return cc.Crawl()
}
