[![][kuma-logo]][kuma-url]

[![CircleCI](https://circleci.com/gh/Kong/kuma.svg?style=svg&circle-token=e3f6c5429ee47ca0eb4bd2542e4b8801a7856373)](https://circleci.com/gh/Kong/kuma)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/Kong/kuma/blob/master/LICENSE)
[![Twitter](https://img.shields.io/twitter/follow/thekonginc.svg?style=social&label=Follow)](https://twitter.com/intent/follow?screen_name=thekonginc)

Kuma is a universal open source control-plane for Service Mesh and Microservices that can run and be operated natively across both Kubernetes and VM environments, in order to be easily adopted by every team in the organization.

Built on top of Envoy, Kuma can instrument any L4/L7 traffic to secure, observe, route and enhance connectivity between any service or database. It can be used natively in Kubernetes via CRDs or via a RESTful API on VM and Bare Metal environments, and it doesn't require a change to your application's code in order to be used.

Built by Envoy contributors at Kong 🦍.

**Need help?** Installing and using Kuma should be as easy as possible. [Contact and chat](https://kuma.io/community) with the community in real-time if you get stuck or need clarifications. We are here to help.

[Installation](https://kuma.io/install) |
[Documentation](https://kuma.io/docs) |
[Slack Chat](https://chat.kuma.io) |
[Community](https://kuma.io/community) |
[Blog](https://konghq.com/blog) |
[Kong](https://konghq.com)

## Summary

- [**Why Kuma?**](#why-kuma)
- [**Features**](#features)
- [**Distributions**](#distributions)
- [**Development**](#development)
- [**Enterprise Demo**](#enterprise-demo)
- [**License**](#license)

## Why Kuma?

Modern applications will inevitably make requests over a network to communicate to other services, like databases, caches or microservices. But - as we all know - the network is by default unreliable and unsecure, and can introduce significant challenges to any modern environment like security, tracing and routing amongh the others.

Kuma is a better way to build L4/L7 connectivity among your services and applications (Service Mesh) by reducing the code that application teams have to write, enabling to ship products faster and improve the reliability and security of the overall architecture with minimal effort. Therefore, Kuma embraces the sidecar proxy model by leveraging Envoy as its sidecar data-plane technology and by providing a Universal Control Plane that can run on both modern Kubernetes and existing VM/Bare Metal architectures in order to deliver business value across every team in the organization with one comprhensive solution.

[![][kuma-benefits]][kuma-url]

## Features

* **Universal Control Plane**: Easy to use, distributed, runs anywhere on both Kubernetes and VM/Bare Metal.
* **Lightweight Data Plane**: Powered by Envoy to process any L4/L7 traffic, with automatic Envoy bootstrapping.
* **Automatic DP Injection**: No code changes required in K8s. Easy YAML specification for VM and Bare Metal deployments.
* **Multi-Tenancy**: To setup multiple isolated Meshes in one cluster and one Control Plane, lowering OPs cost.
* **mTLS**: Automatic mTLS issuing, identity and encryption with optional support for third-party CA.
* **Traffic Permissions**: To firewall traffic between the services of a Mesh.
* **Traffic Routing**: With dynamic load-balancing for blue/green, canary, versioning and rollback deployments.
* **Traffic Logs**: To log all the activity to a third-party service, like Splunk or ELK.
* **Traffic Metrics**: For every Envoy dataplane managed by Kuma, via Prometheus and other integrations.
* **Proxy Configuration Templating**: The easiest way to run and configure Envoy with low-level configuration.
* **Healthchecks**: Both active and passive.
* **GUI**: Out of the box browser GUI to explore all the Service Meshes configured in the system.
* **Tagging Selectors**: To apply sophisticated regional, cloud-specific and team-oriented policies.
* **Platform-Agnostic**: Support for K8s, VMs, and bare metal.
* **Powerful APIM Ingress**: Support for [Kong Gateway](https://github.com/Kong/kong) integration to provide full-cycle APIM.

## Distributions

Kuma is a platform-agnostic product that comes in many shapes. You can explore the available installation options at [the official website](https://kuma.io/install).

You can use Kuma for modern greenfield applications built on containers as well as existing applications running on more traditional infrastructure. Kuma can be fully configured via CRDs (Custom Resource Definitions) on Kubernetes and via a RESTful HTTP API in other environments that can be easily integrated with CI/CD workflows.

Kuma also provides an easy to use `kumactl` CLI client for every environment.

## Development

Kuma is under active development and production-ready.

See [Developer Guide](DEVELOPER.md) for further details.

## Enterprise Demo

If you are implementing Kuma in a mission-critical environment, visit [Request Demo](https://kuma.io/request-demo/) and get in touch with Kong.

## License

```
Copyright 2019 Kong Inc.

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

[kuma-url]: https://kuma.io/
[kuma-logo]: https://kuma-public-assets.s3.amazonaws.com/kuma-logo.png
[kuma-benefits]: https://kuma-public-assets.s3.amazonaws.com/kuma-benefits.jpg
