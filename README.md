# EventMeta

This package is an example on how to version events handled in an event driven architecture.

It's ideas are based on learnings from the [Kubernetes Project](https://github.com/kubernetes/kubernetes) and our internal works.

## General

This library is mainly meant as an demonstration and maybe as playground for further changes to the ideas presented.

The code found here is a partial mirror pulled from our internal repository and changes might get synced from there periodically.

This library provides some basic functionality to create versioned events and allow partial and full parsing based on those versions.

## Use

The main type here is `Event` which defines the interface implemented by `EventMeta`.

An example Event is provided in the v0 folder including api group and versions.
Additionally this library comes with an `Incomplete` type to access metadata from not already identified (unstructured) data.

## Example

There is a very basic example in [example/main.go](example/main.go) (more tbd).

## License

For parts based on Kubernetes, the respective license applies:

```
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

The full license can be found in [LICENSE_k8s](LICENSE_k8s).
Things based on work from the Kubernetes project mainly are adapting mindset and functionality for object versioning and access that then were adapted to different types.

For internal code, the MIT License applies which is located in [LICENSE](LICENSE).
