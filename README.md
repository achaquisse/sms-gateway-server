sms-gateway-server
=====

# Why?

It is needed to use a mobile phone as an SMS gateway.

# How

Server applications will call an HTTP Endpoint requesting a message to be sent, and then an android mobile application
will be pooling an HTTP Endpoint querying messages at the pending state, and schedule them to be sent.

# Solution components

* Server [this project]
* Android client [todo: link here]

# License

Copyright 2022 Américo Chaquisse

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the
License. You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "
AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific
language governing permissions and limitations under the License.